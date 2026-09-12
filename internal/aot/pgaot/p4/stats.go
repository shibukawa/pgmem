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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	v6 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(240)
	m.G0 = v18
	v20 = int32(112)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v21 == v6 {
		v54 = v6
		v58 = v20
		v59 = v6
	} else {
		v29 = v21
		v30 = v6
		v34 = v20
		v35 = v6
		for {
			v40 = v35 + int32(1)
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
			v43 = v41 + (v34 - v29)
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
			v46 = v30 + v41 - v45
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
			if v47 != 0 {
				v29 = v47
				v30 = v46
				v34 = v43
				v35 = v40
				continue
			} else {
				break
			}
			break
		}
		v54 = v46
		v58 = v43
		v59 = v40
	}
	v71 = v54
	v73 = v6
	v74 = v6
	for {
		v82 = l0 + int32(48) + v74<<(uint(int32(2))%32)
		v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
		if v83 != 0 {
			v84 = int32(8)
			v93 = v83
			v94 = v71
			v96 = v73
			for {
				v103 = v94 + (v84<<(uint(v74)%32) + v84)
				v105 = v96 + int32(1)
				v106 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
				if v106 != 0 {
					v93 = v106
					v94 = v103
					v96 = v105
					continue
				} else {
					break
				}
				break
			}
			v113 = v103
			v115 = v105
		} else {
			v113 = v71
			v115 = v73
		}
		if v74 != int32(10) {
			v124 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
			if v124 != 0 {
				v134 = v124
				v135 = v113
				v137 = v115
				for {
					v144 = v135 + (int32(16)<<(uint(v74)%32) | int32(8))
					v146 = v137 + int32(1)
					v147 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
					if v147 != 0 {
						v134 = v147
						v135 = v144
						v137 = v146
						continue
					} else {
						break
					}
					break
				}
				v154 = v144
				v156 = v146
			} else {
				v154 = v113
				v156 = v115
			}
			v71 = v154
			v73 = v156
			v74 = v74 + int32(2)
			continue
		} else {
			break
		}
		break
	}
	if l1 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v113
		*(*int32)(unsafe.Add(mBase, uint32(v18))) = v58
		*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v58 - v113
		*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v59
		*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v115
		v175 = F_pg_snprintf(m, v18+int32(32), int32(200), int32(471876), v18)
		mBase = m.M
		v176 = m.ExcPending
		if v176 != 0 {
			return
		} else {
			m.T0[l1].(func(*base.Module, int32, int32, int32, int32))(m, l0, l2, v18+int32(32), l4)
			mBase = m.M
			v180 = m.ExcPending
			if v180 != 0 {
				return
			} else {
				if l3 != 0 {
					v181 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v181 + v59
					v184 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v184 + v115
					v187 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v187 + v58
					v190 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v190 + v113
				} else {
				}
				m.G0 = v18 + int32(240)
				return
			}
		}
	} else {
		if l3 != 0 {
			v181 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v181 + v59
			v184 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v184 + v115
			v187 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v187 + v58
			v190 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v190 + v113
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
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v44<<(uint(int32(3))%32))+uint32(_consts[899])))
						if v21 != 0 {
							v50 = l2
						} else {
							v50 = l1
						}
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v50<<(uint(int32(3))%32))+uint32(_consts[899])))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v55
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v49
						F_errmsg(m, int32(480516), v9)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(519522), int32(116), int32(225064))
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
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v44<<(uint(int32(3))%32))+uint32(_consts[899])))
						if v21 != 0 {
							v50 = l2
						} else {
							v50 = l1
						}
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v50<<(uint(int32(3))%32))+uint32(_consts[899])))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v55
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v49
						F_errmsg(m, int32(480516), v9)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(519522), int32(116), int32(225064))
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
