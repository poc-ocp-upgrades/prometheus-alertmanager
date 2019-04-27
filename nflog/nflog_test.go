package nflog

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
	pb "github.com/prometheus/alertmanager/nflog/nflogpb"
	"github.com/stretchr/testify/require"
)

func TestLogGC(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	now := utcNow()
	newEntry := func(ts time.Time) *pb.MeshEntry {
		return &pb.MeshEntry{ExpiresAt: ts}
	}
	l := &Log{st: state{"a1": newEntry(now), "a2": newEntry(now.Add(time.Second)), "a3": newEntry(now.Add(-time.Second))}, now: func() time.Time {
		return now
	}, metrics: newMetrics(nil)}
	n, err := l.GC()
	require.NoError(t, err, "unexpected error in garbage collection")
	require.Equal(t, 2, n, "unexpected number of removed entries")
	expected := state{"a2": newEntry(now.Add(time.Second))}
	require.Equal(t, l.st, expected, "unepexcted state after garbage collection")
}
func TestLogSnapshot(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	now := utcNow()
	cases := []struct{ entries []*pb.MeshEntry }{{entries: []*pb.MeshEntry{{Entry: &pb.Entry{GroupKey: []byte("d8e8fca2dc0f896fd7cb4cb0031ba249"), Receiver: &pb.Receiver{GroupName: "abc", Integration: "test1", Idx: 1}, GroupHash: []byte("126a8a51b9d1bbd07fddc65819a542c3"), Resolved: false, Timestamp: now}, ExpiresAt: now}, {Entry: &pb.Entry{GroupKey: []byte("d8e8fca2dc0f8abce7cb4cb0031ba249"), Receiver: &pb.Receiver{GroupName: "def", Integration: "test2", Idx: 29}, GroupHash: []byte("122c2331b9d1bbd07fddc65819a542c3"), Resolved: true, Timestamp: now}, ExpiresAt: now}, {Entry: &pb.Entry{GroupKey: []byte("aaaaaca2dc0f896fd7cb4cb0031ba249"), Receiver: &pb.Receiver{GroupName: "ghi", Integration: "test3", Idx: 0}, GroupHash: []byte("126a8a51b9d1bbd07fddc6e3e3e542c3"), Resolved: false, Timestamp: now}, ExpiresAt: now}}}}
	for _, c := range cases {
		f, err := ioutil.TempFile("", "snapshot")
		require.NoError(t, err, "creating temp file failed")
		l1 := &Log{st: state{}, metrics: newMetrics(nil)}
		for _, e := range c.entries {
			l1.st[stateKey(string(e.Entry.GroupKey), e.Entry.Receiver)] = e
		}
		_, err = l1.Snapshot(f)
		require.NoError(t, err, "creating snapshot failed")
		require.NoError(t, f.Close(), "closing snapshot file failed")
		f, err = os.Open(f.Name())
		require.NoError(t, err, "opening snapshot file failed")
		l2 := &Log{}
		err = l2.loadSnapshot(f)
		require.NoError(t, err, "error loading snapshot")
		require.Equal(t, l1.st, l2.st, "state after loading snapshot did not match snapshotted state")
		require.NoError(t, f.Close(), "closing snapshot file failed")
	}
}
func TestReplaceFile(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dir, err := ioutil.TempDir("", "replace_file")
	require.NoError(t, err, "creating temp dir failed")
	origFilename := filepath.Join(dir, "testfile")
	of, err := os.Create(origFilename)
	require.NoError(t, err, "creating file failed")
	nf, err := openReplace(origFilename)
	require.NoError(t, err, "opening replacement file failed")
	_, err = nf.Write([]byte("test"))
	require.NoError(t, err, "writing replace file failed")
	require.NotEqual(t, nf.Name(), of.Name(), "replacement file must have different name while editing")
	require.NoError(t, nf.Close(), "closing replacement file failed")
	require.NoError(t, of.Close(), "closing original file failed")
	ofr, err := os.Open(origFilename)
	require.NoError(t, err, "opening original file failed")
	defer ofr.Close()
	res, err := ioutil.ReadAll(ofr)
	require.NoError(t, err, "reading original file failed")
	require.Equal(t, "test", string(res), "unexpected file contents")
}
func TestStateMerge(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	now := utcNow()
	newEntry := func(name string, ts, exp time.Time) *pb.MeshEntry {
		return &pb.MeshEntry{Entry: &pb.Entry{Timestamp: ts, GroupKey: []byte("key"), Receiver: &pb.Receiver{GroupName: name, Idx: 1, Integration: "integr"}}, ExpiresAt: exp}
	}
	exp := now.Add(time.Minute)
	cases := []struct {
		a, b	state
		final	state
	}{{a: state{"key:a1/integr/1": newEntry("a1", now, exp), "key:a2/integr/1": newEntry("a2", now, exp), "key:a3/integr/1": newEntry("a3", now, exp)}, b: state{"key:b1/integr/1": newEntry("b1", now, exp), "key:b2/integr/1": newEntry("b2", now.Add(-time.Minute), now.Add(-time.Millisecond)), "key:a2/integr/1": newEntry("a2", now.Add(-time.Minute), exp), "key:a3/integr/1": newEntry("a3", now.Add(time.Minute), exp)}, final: state{"key:a1/integr/1": newEntry("a1", now, exp), "key:a2/integr/1": newEntry("a2", now, exp), "key:a3/integr/1": newEntry("a3", now.Add(time.Minute), exp), "key:b1/integr/1": newEntry("b1", now, exp)}}}
	for _, c := range cases {
		ca, cb := c.a.clone(), c.b.clone()
		res := c.a.clone()
		for _, e := range cb {
			res.merge(e, now)
		}
		require.Equal(t, c.final, res, "Merge result should match expectation")
		require.Equal(t, c.b, cb, "Merged state should remain unmodified")
		require.NotEqual(t, c.final, ca, "Merge should not change original state")
	}
}
func TestStateDataCoding(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	now := utcNow()
	cases := []struct{ entries []*pb.MeshEntry }{{entries: []*pb.MeshEntry{{Entry: &pb.Entry{GroupKey: []byte("d8e8fca2dc0f896fd7cb4cb0031ba249"), Receiver: &pb.Receiver{GroupName: "abc", Integration: "test1", Idx: 1}, GroupHash: []byte("126a8a51b9d1bbd07fddc65819a542c3"), Resolved: false, Timestamp: now}, ExpiresAt: now}, {Entry: &pb.Entry{GroupKey: []byte("d8e8fca2dc0f8abce7cb4cb0031ba249"), Receiver: &pb.Receiver{GroupName: "def", Integration: "test2", Idx: 29}, GroupHash: []byte("122c2331b9d1bbd07fddc65819a542c3"), Resolved: true, Timestamp: now}, ExpiresAt: now}, {Entry: &pb.Entry{GroupKey: []byte("aaaaaca2dc0f896fd7cb4cb0031ba249"), Receiver: &pb.Receiver{GroupName: "ghi", Integration: "test3", Idx: 0}, GroupHash: []byte("126a8a51b9d1bbd07fddc6e3e3e542c3"), Resolved: false, Timestamp: now}, ExpiresAt: now}}}}
	for _, c := range cases {
		in := state{}
		for _, e := range c.entries {
			in[stateKey(string(e.Entry.GroupKey), e.Entry.Receiver)] = e
		}
		msg, err := in.MarshalBinary()
		require.NoError(t, err)
		out, err := decodeState(bytes.NewReader(msg))
		require.NoError(t, err, "decoding message failed")
		require.Equal(t, in, out, "decoded data doesn't match encoded data")
	}
}
func TestQuery(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nl, err := New(WithRetention(time.Second))
	if err != nil {
		require.NoError(t, err, "constructing nflog failed")
	}
	recv := new(pb.Receiver)
	_, err = nl.Query(QGroupKey("key"))
	require.EqualError(t, err, "no query parameters specified")
	_, err = nl.Query(QReceiver(recv))
	require.EqualError(t, err, "no query parameters specified")
	_, err = nl.Query(QGroupKey("nonexistingkey"), QReceiver(recv))
	require.EqualError(t, err, "not found")
	firingAlerts := []uint64{1, 2, 3}
	resolvedAlerts := []uint64{4, 5}
	err = nl.Log(recv, "key", firingAlerts, resolvedAlerts)
	require.NoError(t, err, "logging notification failed")
	entries, err := nl.Query(QGroupKey("key"), QReceiver(recv))
	require.NoError(t, err, "querying nflog failed")
	entry := entries[0]
	require.EqualValues(t, firingAlerts, entry.FiringAlerts)
	require.EqualValues(t, resolvedAlerts, entry.ResolvedAlerts)
}
func TestStateDecodingError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := state{"": &pb.MeshEntry{}}
	msg, err := s.MarshalBinary()
	require.NoError(t, err)
	_, err = decodeState(bytes.NewReader(msg))
	require.Equal(t, ErrInvalidState, err)
}
