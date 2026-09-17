package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitPostmasterChildSlots(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[0])) = int32(32)
	v13 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[1])) = v13
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[2])) = v13
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[3])) = v13
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[4])) = v13
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[5])) = v13
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[6])) = v13
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[7])) = v13
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[8])) = v13
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[9])) = v13
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[10])) = v13
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[12])) = v44
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[13]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[14])) = v48
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[15]))
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[16]))
	v57 = (v52 + v54) << (uint(v13) % 32)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[17])) = v57
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[18]))
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[19]))
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[20]))
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[21]))
	v75 = v61 + (v63 + (v48 + (v44 + (v65 + (v67 + v57))))) + int32(42)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[22])) = v75
	v79 = F_palloc(m, v75*int32(28))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		return
	} else {
		v84 = int32(0)
		v86 = int32(0)
		for {
			v91 = v86 << (uint(int32(4)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[23]))) = v84 + int32(1)
			v96 = v91 + int32(_a_F_InitPostmasterChildSlots_0)
			*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[24]))) = v96
			*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[25]))) = v96
			v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[21])))
			if int32(0) < v99 {
				v107 = v84
				v110 = int32(0)
				for {
					v115 = v79 + v107*int32(28)
					*(*int64)(unsafe.Add(mBase, uint32(v115)+8)) = int64(0)
					v119 = v107 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v115)+4)) = v119
					v121 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v115))) = v121
					*(*uint8)(unsafe.Add(mBase, uint32(v115)+16)) = uint8(v121)
					v125 = *(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[24])))
					if v125 == v121 {
						*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[24]))) = v96
						*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[25]))) = v91 + int32(_a_F_InitPostmasterChildSlots_0)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v115)+24)) = v96
					v133 = *(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[25])))
					*(*int32)(unsafe.Add(mBase, uint32(v115)+20)) = v133
					v136 = v115 + int32(20)
					*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = v136
					*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[25]))) = v136
					v140 = v110 + int32(1)
					v141 = *(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[21])))
					if v140 < v141 {
						v107 = v119
						v110 = v140
						continue
					} else {
						break
					}
					break
				}
				v145 = v119
			} else {
				v145 = v84
			}
			v152 = v86 + int32(1)
			if v152 != int32(18) {
				v84 = v145
				v86 = v152
				continue
			} else {
				break
			}
			break
		}
		v156 = int32(_a_F_InitPostmasterChildSlots_1)
		*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[26])) = v156
		*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[27])) = v156
		return
	}
}
func F_ReleasePostmasterChildSlot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if v15 == int32(2) {
			if v18 != 0 {
				F_errmsg_internal(m, int32(_a_F_ReleasePostmasterChildSlot_0), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_ReleasePostmasterChildSlot_1), int32(241), int32(_a_F_ReleasePostmasterChildSlot_2))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_pfree(m, l0)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v93 = int32(1)
							m.G0 = v8 + int32(32)
							return v93
						}
					}
				}
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v93 = int32(1)
					m.G0 = v8 + int32(32)
					return v93
				}
			}
		} else {
			if v18 != 0 {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v36
				F_errmsg_internal(m, int32(_a_F_ReleasePostmasterChildSlot_3), v8+int32(16))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_ReleasePostmasterChildSlot_1), int32(249), int32(_a_F_ReleasePostmasterChildSlot_2))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v50 == int32(6) {
							v57 = int32(_a_F_ReleasePostmasterChildSlot_4)
						} else {
							v57 = v50<<(uint(int32(4))%32) + int32(_a_F_ReleasePostmasterChildSlot_5)
						}
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
						if v48 < v58 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v102
								F_errmsg_internal(m, int32(_a_F_ReleasePostmasterChildSlot_6), v8)
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ReleasePostmasterChildSlot_1), int32(262), int32(_a_F_ReleasePostmasterChildSlot_2))
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
							if v60+v58 <= v48 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v102
									F_errmsg_internal(m, int32(_a_F_ReleasePostmasterChildSlot_6), v8)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_ReleasePostmasterChildSlot_1), int32(262), int32(_a_F_ReleasePostmasterChildSlot_2))
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v64 = v57 + int32(8)
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
								if v65 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v64))) = v64
									v69 = v64
								} else {
									v69 = v65
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v64
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v69
								v73 = l0 + int32(20)
								*(*int32)(unsafe.Add(mBase, uint32(v69))) = v73
								*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v73
								v77 = *(*int32)(unsafe.Add(mBase, _c_F_ReleasePostmasterChildSlot[0]))
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v83 = v77 + v78<<(uint(int32(2))%32) + int32(44)
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
								*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
								v93 = base.B2i32(v84 == int32(1))
								m.G0 = v8 + int32(32)
								return v93
							}
						}
					}
				}
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v50 == int32(6) {
					v57 = int32(_a_F_ReleasePostmasterChildSlot_4)
				} else {
					v57 = v50<<(uint(int32(4))%32) + int32(_a_F_ReleasePostmasterChildSlot_5)
				}
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
				if v48 < v58 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
					} else {
						v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v102
						F_errmsg_internal(m, int32(_a_F_ReleasePostmasterChildSlot_6), v8)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_ReleasePostmasterChildSlot_1), int32(262), int32(_a_F_ReleasePostmasterChildSlot_2))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
					if v60+v58 <= v48 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v102
							F_errmsg_internal(m, int32(_a_F_ReleasePostmasterChildSlot_6), v8)
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_ReleasePostmasterChildSlot_1), int32(262), int32(_a_F_ReleasePostmasterChildSlot_2))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v64 = v57 + int32(8)
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
						if v65 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v64))) = v64
							v69 = v64
						} else {
							v69 = v65
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v64
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v69
						v73 = l0 + int32(20)
						*(*int32)(unsafe.Add(mBase, uint32(v69))) = v73
						*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v73
						v77 = *(*int32)(unsafe.Add(mBase, _c_F_ReleasePostmasterChildSlot[0]))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v83 = v77 + v78<<(uint(int32(2))%32) + int32(44)
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
						*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
						v93 = base.B2i32(v84 == int32(1))
						m.G0 = v8 + int32(32)
						return v93
					}
				}
			}
		}
	}
}
func F_postmaster_child_launch(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	switch l0 - int32(1) {
	case 0, 5:
		v8 = m.G0
		v9 = int32(16)
		v10 = v8 - v9
		m.G0 = v10
		F_gettimeofday(m, v10)
		mBase = m.M
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		v14 = int64(*(*int32)(unsafe.Add(mBase, uint32(v10)+8)))
		m.G0 = v10 + v9
		*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v14 + v13*int64(1000000) - int64(946684800000000)
	default:
	}
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	v29 = F_fflush(m, int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return int32(0)
	} else {
		v35 = v26 + int32(8)
		F_sigprocmask(m, int32(_a_F_postmaster_child_launch_0), v35)
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[0])) = int32(52)
			F_sigprocmask(m, v35, int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				m.G0 = v26 + int32(16)
				return int32(-1)
			}
		}
	}
}
