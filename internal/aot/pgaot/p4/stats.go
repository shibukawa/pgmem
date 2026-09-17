package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AllocSetStats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	v6 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(240)
	m.G0 = v18
	v20 = int32(112)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v21 != 0 {
		v27 = v21
		v29 = v6
		v30 = v20
		v33 = v6
		for {
			v38 = v33 + int32(1)
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
			v41 = v39 + (v30 - v27)
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
			v44 = v29 + v39 - v43
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
			if v45 != 0 {
				v27 = v45
				v29 = v44
				v30 = v41
				v33 = v38
				continue
			} else {
				break
			}
			break
		}
		v53 = v44
		v54 = v41
		v57 = v38
	} else {
		v53 = v6
		v54 = v20
		v57 = v6
	}
	v70 = v53
	v72 = v6
	v73 = v6
	for {
		v80 = l0 + int32(48) + v72<<(uint(int32(2))%32)
		v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
		if v81 != 0 {
			v82 = int32(8)
			v91 = v81
			v93 = v70
			v96 = v73
			for {
				v101 = v93 + (v82<<(uint(v72)%32) + v82)
				v103 = v96 + int32(1)
				v104 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
				if v104 != 0 {
					v91 = v104
					v93 = v101
					v96 = v103
					continue
				} else {
					break
				}
				break
			}
			v112 = v101
			v115 = v103
		} else {
			v112 = v70
			v115 = v73
		}
		if v72 != int32(10) {
			v122 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
			if v122 != 0 {
				v132 = v122
				v134 = v112
				v137 = v115
				for {
					v142 = v134 + (int32(16)<<(uint(v72)%32) | int32(8))
					v144 = v137 + int32(1)
					v145 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
					if v145 != 0 {
						v132 = v145
						v134 = v142
						v137 = v144
						continue
					} else {
						break
					}
					break
				}
				v153 = v142
				v156 = v144
			} else {
				v153 = v112
				v156 = v115
			}
			v70 = v153
			v72 = v72 + int32(2)
			v73 = v156
			continue
		} else {
			break
		}
		break
	}
	if l1 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v112
		*(*int32)(unsafe.Add(mBase, uint32(v18))) = v54
		*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v54 - v112
		*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v57
		*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v115
		v170 = v18 + int32(32)
		v173 = F_pg_snprintf(m, v170, int32(200), int32(_a_F_AllocSetStats_0), v18)
		mBase = m.M
		v174 = m.ExcPending
		if v174 != 0 {
			return
		} else {
			m.T0[l1].(func(*base.Module, int32, int32, int32, int32))(m, l0, l2, v170, l4)
			mBase = m.M
			v176 = m.ExcPending
			if v176 != 0 {
				return
			} else {
				if l3 != 0 {
					v178 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v178 + v57
					v181 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v181 + v115
					v184 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v184 + v54
					v187 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v187 + v112
				} else {
				}
				m.G0 = v18 + int32(240)
				return
			}
		}
	} else {
		if l3 != 0 {
			v178 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v178 + v57
			v181 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v181 + v115
			v184 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v184 + v54
			v187 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v187 + v112
		} else {
		}
		m.G0 = v18 + int32(240)
		return
	}
}
func F_stats_check_arg_pair(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(1)
	v13 = l0 + int32(20)
	v14 = int32(3)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+l2<<(uint(v14)%32))+4)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+l1<<(uint(v14)%32))+4)))
	if v21 == v11 {
		if v17&int32(1) == int32(0) {
			v32 = int32(0)
			v35 = F_errstart(m, int32(19), v32)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				if v35 == int32(0) {
					v67 = v32
					m.G0 = v9 + int32(16)
					return v67
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						if v21 != 0 {
							v44 = l1
						} else {
							v44 = l2
						}
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v44<<(uint(int32(3))%32))+uint32(_c_F_stats_check_arg_pair[0])))
						if v21 != 0 {
							v50 = l2
						} else {
							v50 = l1
						}
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v50<<(uint(int32(3))%32))+uint32(_c_F_stats_check_arg_pair[0])))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v55
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v49
						F_errmsg(m, int32(_a_F_stats_check_arg_pair_0), v9)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_stats_check_arg_pair_1), int32(116), int32(_a_F_stats_check_arg_pair_2))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								v67 = v32
								m.G0 = v9 + int32(16)
								return v67
							}
						}
					}
				}
			}
		} else {
			v67 = v11
			m.G0 = v9 + int32(16)
			return v67
		}
	} else {
		if v17&int32(1) == int32(0) {
			v67 = v11
			m.G0 = v9 + int32(16)
			return v67
		} else {
			v32 = int32(0)
			v35 = F_errstart(m, int32(19), v32)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				if v35 == int32(0) {
					v67 = v32
					m.G0 = v9 + int32(16)
					return v67
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						if v21 != 0 {
							v44 = l1
						} else {
							v44 = l2
						}
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v44<<(uint(int32(3))%32))+uint32(_c_F_stats_check_arg_pair[0])))
						if v21 != 0 {
							v50 = l2
						} else {
							v50 = l1
						}
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v50<<(uint(int32(3))%32))+uint32(_c_F_stats_check_arg_pair[0])))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v55
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v49
						F_errmsg(m, int32(_a_F_stats_check_arg_pair_0), v9)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_stats_check_arg_pair_1), int32(116), int32(_a_F_stats_check_arg_pair_2))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								v67 = v32
								m.G0 = v9 + int32(16)
								return v67
							}
						}
					}
				}
			}
		}
	}
}
