package nflogpb

func (m *Entry) IsFiringSubset(subset map[uint64]struct{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	set := map[uint64]struct{}{}
	for i := range m.FiringAlerts {
		set[m.FiringAlerts[i]] = struct{}{}
	}
	return isSubset(set, subset)
}
func (m *Entry) IsResolvedSubset(subset map[uint64]struct{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	set := map[uint64]struct{}{}
	for i := range m.ResolvedAlerts {
		set[m.ResolvedAlerts[i]] = struct{}{}
	}
	return isSubset(set, subset)
}
func isSubset(set, subset map[uint64]struct{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for k := range subset {
		_, exists := set[k]
		if !exists {
			return false
		}
	}
	return true
}
