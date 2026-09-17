package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_DCH_datetime_type(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	v3 = l0
	v4 = int32(0)
wl1:
	for {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
		switch v5 - int32(1) {
		case 0:
			break wl1
		case 1:
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			switch v9 {
			case 0, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 24, 25, 27, 28, 29, 30, 31, 33, 34, 35, 37, 38, 42, 43, 51, 52, 53, 54, 55, 56, 57, 58, 60, 62, 63, 65, 68, 90, 91, 97:
				v20 = v4 | int32(1)
				v3 = v3 + int32(12)
				v4 = v20
				continue
			case 1, 3, 14, 15, 16, 17, 18, 19, 21, 22, 23, 32, 36, 40, 41, 45, 46, 50, 59, 61, 94, 95:
				v3 = v3 + int32(12)
				v4 = v4 | int32(2)
				continue
			default:
				v20 = v4
				v3 = v3 + int32(12)
				v4 = v20
				continue
			case 39, 47, 48, 49, 103:
				v3 = v3 + int32(12)
				v4 = v4 | int32(4)
				continue
			}
			continue
		default:
			v20 = v4
			v3 = v3 + int32(12)
			v4 = v20
			continue
		}
		break
	}
	return v4
}
func F_DefineSavepoint(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_DefineSavepoint[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+72))
	if v11 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v63 = m.ExcPending
		if v63 != 0 {
			return
		} else {
			F_errcode(m, int32(322))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_DefineSavepoint_0), int32(0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_DefineSavepoint_1), int32(_a_F_DefineSavepoint_2), int32(_a_F_DefineSavepoint_3))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+76)))
		if v12 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return
			} else {
				F_errcode(m, int32(322))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_DefineSavepoint_0), int32(0))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_DefineSavepoint_1), int32(_a_F_DefineSavepoint_2), int32(_a_F_DefineSavepoint_3))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_DefineSavepoint[1]))
			if int32(0) <= v14 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_errcode(m, int32(322))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_DefineSavepoint_0), int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_DefineSavepoint_1), int32(_a_F_DefineSavepoint_2), int32(_a_F_DefineSavepoint_3))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
				if base.Ui32(int32(19)) < base.Ui32(v17) {
					m.G0 = v7 + int32(32)
					return
				} else {
					v21 = int32(1) << (uint(v17) % 32)
					if v21&int32(_a_F_DefineSavepoint_4) == int32(0) {
						if v21&int32(_a_F_DefineSavepoint_5) != 0 {
							F_errstart_cold(m, int32(22), int32(0))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
								if base.Ui32(v80) <= base.Ui32(int32(19)) {
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v80<<(uint(int32(2))%32))+uint32(_c_F_DefineSavepoint[2])))
									v87 = v85
								} else {
									v87 = int32(_a_F_DefineSavepoint_6)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v87
								F_errmsg_internal(m, int32(_a_F_DefineSavepoint_7), v7+int32(16))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_DefineSavepoint_1), int32(_a_F_DefineSavepoint_8), int32(_a_F_DefineSavepoint_3))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								F_errcode(m, int32(16908610))
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_DefineSavepoint_9)
									F_errmsg(m, int32(_a_F_DefineSavepoint_10), v7)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_DefineSavepoint_1), int32(_a_F_DefineSavepoint_11), int32(_a_F_DefineSavepoint_3))
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					} else {
						F_PushTransaction(m)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							if l0 == int32(0) {
								m.G0 = v7 + int32(32)
								return
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, _c_F_DefineSavepoint[0]))
								v52 = *(*int32)(unsafe.Add(mBase, _c_F_DefineSavepoint[3]))
								v53 = F_MemoryContextStrdup(m, v52, l0)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = v53
									m.G0 = v7 + int32(32)
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_DescribeLockTag(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	v7 = m.G0
	v9 = v7 - int32(208)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)))
	switch v11 {
	case 0:
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = base.I64_rotl(v12, int64(32))
		F_appendStringInfo(m, l0, int32(_a_F_DescribeLockTag_0), v9+int32(16))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			m.G0 = v9 + int32(208)
			return
		}
	case 1:
		v21 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = base.I64_rotl(v21, int64(32))
		F_appendStringInfo(m, l0, int32(_a_F_DescribeLockTag_1), v9+int32(32))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			m.G0 = v9 + int32(208)
			return
		}
	case 2:
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v30
		F_appendStringInfo(m, l0, int32(_a_F_DescribeLockTag_2), v9+int32(48))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			m.G0 = v9 + int32(208)
			return
		}
	case 3:
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+72)) = v39
		*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v38
		*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v37
		F_appendStringInfo(m, l0, int32(_a_F_DescribeLockTag_3), v9-int32(-64))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return
		} else {
			m.G0 = v9 + int32(208)
			return
		}
	case 4:
		v48 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v50
		*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v49
		*(*int64)(unsafe.Add(mBase, uint32(v9)+88)) = base.I64_rotl(v48, int64(32))
		F_appendStringInfo(m, l0, int32(_a_F_DescribeLockTag_4), v9+int32(80))
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return
		} else {
			m.G0 = v9 + int32(208)
			return
		}
	case 5:
		v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v61
		F_appendStringInfo(m, l0, int32(_a_F_DescribeLockTag_5), v9+int32(96))
		mBase = m.M
		v67 = m.ExcPending
		if v67 != 0 {
			return
		} else {
			m.G0 = v9 + int32(208)
			return
		}
	case 6:
		v68 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		*(*int64)(unsafe.Add(mBase, uint32(v9)+112)) = v68
		F_appendStringInfo(m, l0, int32(_a_F_DescribeLockTag_6), v9+int32(112))
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return
		} else {
			m.G0 = v9 + int32(208)
			return
		}
	case 7:
		v75 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		*(*int64)(unsafe.Add(mBase, uint32(v9)+128)) = base.I64_rotl(v75, int64(32))
		F_appendStringInfo(m, l0, int32(_a_F_DescribeLockTag_7), v9+int32(128))
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return
		} else {
			m.G0 = v9 + int32(208)
			return
		}
	case 8:
		v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+152)) = v86
		*(*int32)(unsafe.Add(mBase, uint32(v9)+148)) = v85
		*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v84
		F_appendStringInfo(m, l0, int32(_a_F_DescribeLockTag_8), v9+int32(144))
		mBase = m.M
		v94 = m.ExcPending
		if v94 != 0 {
			return
		} else {
			m.G0 = v9 + int32(208)
			return
		}
	case 9:
		v95 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+168)) = v96
		*(*int64)(unsafe.Add(mBase, uint32(v9)+160)) = v95
		F_appendStringInfo(m, l0, int32(_a_F_DescribeLockTag_9), v9+int32(160))
		mBase = m.M
		v103 = m.ExcPending
		if v103 != 0 {
			return
		} else {
			m.G0 = v9 + int32(208)
			return
		}
	case 10:
		v104 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+188)) = v106
		*(*int32)(unsafe.Add(mBase, uint32(v9)+184)) = v105
		*(*int64)(unsafe.Add(mBase, uint32(v9)+176)) = v104
		F_appendStringInfo(m, l0, int32(_a_F_DescribeLockTag_10), v9+int32(176))
		mBase = m.M
		v114 = m.ExcPending
		if v114 != 0 {
			return
		} else {
			m.G0 = v9 + int32(208)
			return
		}
	case 11:
		v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+200)) = v117
		*(*int32)(unsafe.Add(mBase, uint32(v9)+196)) = v116
		*(*int32)(unsafe.Add(mBase, uint32(v9)+192)) = v115
		F_appendStringInfo(m, l0, int32(_a_F_DescribeLockTag_11), v9+int32(192))
		mBase = m.M
		v125 = m.ExcPending
		if v125 != 0 {
			return
		} else {
			m.G0 = v9 + int32(208)
			return
		}
	default:
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
		F_appendStringInfo(m, l0, int32(_a_F_DescribeLockTag_12), v9)
		mBase = m.M
		v129 = m.ExcPending
		if v129 != 0 {
			return
		} else {
			m.G0 = v9 + int32(208)
			return
		}
	}
}
func F_DetermineTimeZoneOffsetInternal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int64
	_ = v75
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v15 <= int32(-4713) {
		if v15 != int32(-4713) {
			v150 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v150
			*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
			v163 = v150
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if int32(10) < v20 {
				v31 = v20
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v35 = int32(60)
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v46 = base.B2i32(int32(2) < v31)
				if int32(2) < v31 {
					v47 = int32(_a_F_DetermineTimeZoneOffsetInternal_0)
				} else {
					v47 = int32(_a_F_DetermineTimeZoneOffsetInternal_1)
				}
				v48 = v47 + v15
				v56 = base.I32_div_u_s(v48, int32(100))
				v59 = base.I32_div_u_s(v48, int32(400))
				if int32(2) < v31 {
					v63 = int32(1)
				} else {
					v63 = int32(13)
				}
				v68 = base.I32_div_s((v63+v31)*int32(_a_F_DetermineTimeZoneOffsetInternal_2), int32(256))
				v71 = v42 + v48*int32(365) + int32(base.Ui32(v48)>>(uint(int32(2))%32)) - v56 + v59 + v68 - int32(_a_F_DetermineTimeZoneOffsetInternal_3)
				v75 = base.I64_extend_i32_s(v32+(v33+v34*v35)*v35) + base.I64_extend_i32_s(v71)*int64(86400)
				if base.B2i32(int32(0) < v71)&base.B2i32(v75 < int64(0)) != 0 {
					v150 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v150
					*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
					v163 = v150
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v75 - int64(86400)
					v94 = F_pg_next_dst_boundary(m, v13+int32(24), v13+int32(12), v13+int32(4), v13+int32(16), v13+int32(8), v13, l1)
					mBase = m.M
					if v94 < int32(0) {
						v150 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v150
						*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
						v163 = v150
					} else {
						if v94 == int32(0) {
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v99
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
							*(*int64)(unsafe.Add(mBase, uint32(l2))) = v75 - base.I64_extend_i32_s(v101)
							v163 = int32(0) - v101
						} else {
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
							v109 = v75 - base.I64_extend_i32_s(v107)
							v110 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
							v112 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
							v114 = v75 - base.I64_extend_i32_s(v112)
							if base.B2i32(v110 <= v109)|base.B2i32(v110 <= v114) == int32(0) {
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v119
								*(*int64)(unsafe.Add(mBase, uint32(l2))) = v109
								v163 = int32(0) - v107
							} else {
								if base.B2i32(v114 < v110)|base.B2i32(v109 <= v110) == int32(0) {
									v129 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v129
									*(*int64)(unsafe.Add(mBase, uint32(l2))) = v114
									v163 = int32(0) - v112
								} else {
									if v107 < v112 {
										v135 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v135
										*(*int64)(unsafe.Add(mBase, uint32(l2))) = v109
										v163 = int32(0) - v107
									} else {
										v140 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v140
										*(*int64)(unsafe.Add(mBase, uint32(l2))) = v114
										v163 = int32(0) - v112
									}
								}
							}
						}
					}
				}
			} else {
				v150 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v150
				*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
				v163 = v150
			}
		}
	} else {
		if v15 <= int32(_a_F_DetermineTimeZoneOffsetInternal_4) {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v31 = v25
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v35 = int32(60)
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v46 = base.B2i32(int32(2) < v31)
			if int32(2) < v31 {
				v47 = int32(_a_F_DetermineTimeZoneOffsetInternal_0)
			} else {
				v47 = int32(_a_F_DetermineTimeZoneOffsetInternal_1)
			}
			v48 = v47 + v15
			v56 = base.I32_div_u_s(v48, int32(100))
			v59 = base.I32_div_u_s(v48, int32(400))
			if int32(2) < v31 {
				v63 = int32(1)
			} else {
				v63 = int32(13)
			}
			v68 = base.I32_div_s((v63+v31)*int32(_a_F_DetermineTimeZoneOffsetInternal_2), int32(256))
			v71 = v42 + v48*int32(365) + int32(base.Ui32(v48)>>(uint(int32(2))%32)) - v56 + v59 + v68 - int32(_a_F_DetermineTimeZoneOffsetInternal_3)
			v75 = base.I64_extend_i32_s(v32+(v33+v34*v35)*v35) + base.I64_extend_i32_s(v71)*int64(86400)
			if base.B2i32(int32(0) < v71)&base.B2i32(v75 < int64(0)) != 0 {
				v150 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v150
				*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
				v163 = v150
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v75 - int64(86400)
				v94 = F_pg_next_dst_boundary(m, v13+int32(24), v13+int32(12), v13+int32(4), v13+int32(16), v13+int32(8), v13, l1)
				mBase = m.M
				if v94 < int32(0) {
					v150 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v150
					*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
					v163 = v150
				} else {
					if v94 == int32(0) {
						v99 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v99
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
						*(*int64)(unsafe.Add(mBase, uint32(l2))) = v75 - base.I64_extend_i32_s(v101)
						v163 = int32(0) - v101
					} else {
						v107 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
						v109 = v75 - base.I64_extend_i32_s(v107)
						v110 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
						v112 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
						v114 = v75 - base.I64_extend_i32_s(v112)
						if base.B2i32(v110 <= v109)|base.B2i32(v110 <= v114) == int32(0) {
							v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v119
							*(*int64)(unsafe.Add(mBase, uint32(l2))) = v109
							v163 = int32(0) - v107
						} else {
							if base.B2i32(v114 < v110)|base.B2i32(v109 <= v110) == int32(0) {
								v129 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v129
								*(*int64)(unsafe.Add(mBase, uint32(l2))) = v114
								v163 = int32(0) - v112
							} else {
								if v107 < v112 {
									v135 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v135
									*(*int64)(unsafe.Add(mBase, uint32(l2))) = v109
									v163 = int32(0) - v107
								} else {
									v140 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v140
									*(*int64)(unsafe.Add(mBase, uint32(l2))) = v114
									v163 = int32(0) - v112
								}
							}
						}
					}
				}
			}
		} else {
			if v15 != int32(_a_F_DetermineTimeZoneOffsetInternal_5) {
				v150 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v150
				*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
				v163 = v150
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if int32(5) < v28 {
					v150 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v150
					*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
					v163 = v150
				} else {
					v31 = v28
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v35 = int32(60)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v46 = base.B2i32(int32(2) < v31)
					if int32(2) < v31 {
						v47 = int32(_a_F_DetermineTimeZoneOffsetInternal_0)
					} else {
						v47 = int32(_a_F_DetermineTimeZoneOffsetInternal_1)
					}
					v48 = v47 + v15
					v56 = base.I32_div_u_s(v48, int32(100))
					v59 = base.I32_div_u_s(v48, int32(400))
					if int32(2) < v31 {
						v63 = int32(1)
					} else {
						v63 = int32(13)
					}
					v68 = base.I32_div_s((v63+v31)*int32(_a_F_DetermineTimeZoneOffsetInternal_2), int32(256))
					v71 = v42 + v48*int32(365) + int32(base.Ui32(v48)>>(uint(int32(2))%32)) - v56 + v59 + v68 - int32(_a_F_DetermineTimeZoneOffsetInternal_3)
					v75 = base.I64_extend_i32_s(v32+(v33+v34*v35)*v35) + base.I64_extend_i32_s(v71)*int64(86400)
					if base.B2i32(int32(0) < v71)&base.B2i32(v75 < int64(0)) != 0 {
						v150 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v150
						*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
						v163 = v150
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v75 - int64(86400)
						v94 = F_pg_next_dst_boundary(m, v13+int32(24), v13+int32(12), v13+int32(4), v13+int32(16), v13+int32(8), v13, l1)
						mBase = m.M
						if v94 < int32(0) {
							v150 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v150
							*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
							v163 = v150
						} else {
							if v94 == int32(0) {
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v99
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
								*(*int64)(unsafe.Add(mBase, uint32(l2))) = v75 - base.I64_extend_i32_s(v101)
								v163 = int32(0) - v101
							} else {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
								v109 = v75 - base.I64_extend_i32_s(v107)
								v110 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
								v112 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
								v114 = v75 - base.I64_extend_i32_s(v112)
								if base.B2i32(v110 <= v109)|base.B2i32(v110 <= v114) == int32(0) {
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v119
									*(*int64)(unsafe.Add(mBase, uint32(l2))) = v109
									v163 = int32(0) - v107
								} else {
									if base.B2i32(v114 < v110)|base.B2i32(v109 <= v110) == int32(0) {
										v129 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v129
										*(*int64)(unsafe.Add(mBase, uint32(l2))) = v114
										v163 = int32(0) - v112
									} else {
										if v107 < v112 {
											v135 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v135
											*(*int64)(unsafe.Add(mBase, uint32(l2))) = v109
											v163 = int32(0) - v107
										} else {
											v140 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v140
											*(*int64)(unsafe.Add(mBase, uint32(l2))) = v114
											v163 = int32(0) - v112
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	m.G0 = v13 + int32(32)
	return v163
}
func F_DoLockModesConflict(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_DoLockModesConflict[0])))
	return int32(base.Ui32(v7)>>(uint(l1)%32)) & int32(1)
}
func F_dacos(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v17 int64
	_ = v17
	var v22 int32
	_ = v22
	var v35 float64
	_ = v35
	var v46 float64
	_ = v46
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v60 float64
	_ = v60
	var v65 float64
	_ = v65
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v77 float64
	_ = v77
	var v83 float64
	_ = v83
	var v87 float64
	_ = v87
	var v90 float64
	_ = v90
	var v94 float64
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v8 = base.F64_abs(v7)
	if base.Ui64(base.I64_reinterpret_f64(v8)) <= base.Ui64(int64(9218868437227405312)) {
		if base.F64_gt(v8, float64(1)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_dacos_0), int32(0))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_dacos_1), int32(1772), int32(_a_F_dacos_2))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v17 = base.I64_reinterpret_f64(v7)
			v22 = base.I32_wrap_i64(int64(base.Ui64(v17)>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1072693248)) <= base.Ui32(v22) {
				if base.I32_wrap_i64(v17)|(v22-int32(1072693248)) == int32(0) {
					if int64(0) <= v17 {
						v35 = float64(0)
					} else {
						v35 = float64(3.141592653589793)
					}
					v90 = v35
				} else {
					v90 = base.F64_div(float64(0), base.F64_sub(v7, v7))
				}
			} else {
				if base.Ui32(v22) <= base.Ui32(int32(1071644671)) {
					if base.Ui32(v22) < base.Ui32(int32(1012924417)) {
						v87 = float64(1.5707963267948966)
						v90 = v87
					} else {
						v46 = F_R(m, base.F64_mul(v7, v7))
						mBase = m.M
						v90 = base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(v7, v46)), v7), float64(1.5707963267948966))
					}
				} else {
					if v17 < int64(0) {
						v58 = base.F64_mul(base.F64_add(v7, float64(1)), float64(0.5))
						v59 = base.F64_sqrt(v58)
						v60 = F_R(m, v58)
						mBase = m.M
						v65 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v59, base.F64_add(base.F64_mul(v59, v60), float64(-6.123233995736766e-17))))
						v90 = base.F64_add(v65, v65)
					} else {
						v70 = base.F64_mul(base.F64_sub(float64(1), v7), float64(0.5))
						v71 = base.F64_sqrt(v70)
						v72 = F_R(m, v70)
						mBase = m.M
						v77 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v71) & int64(-4294967296))
						v83 = base.F64_add(base.F64_add(base.F64_mul(v71, v72), base.F64_div(base.F64_sub(v70, base.F64_mul(v77, v77)), base.F64_add(v71, v77))), v77)
						v87 = base.F64_add(v83, v83)
						v90 = v87
					}
				}
			}
			if base.F64_eq(base.F64_abs(v90), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v94 = v90
				v95 = F_Float8GetDatum(m, v94)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int32(0)
				} else {
					return v95
				}
			}
		}
	} else {
		v94 = math.Float64frombits(uint64(0x7ff8000000000000))
		v95 = F_Float8GetDatum(m, v94)
		mBase = m.M
		v98 = m.ExcPending
		if v98 != 0 {
			return int32(0)
		} else {
			return v95
		}
	}
}
func F_dacosd(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v18 int32
	_ = v18
	var v33 int64
	_ = v33
	var v38 int32
	_ = v38
	var v61 float64
	_ = v61
	var v68 float64
	_ = v68
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v75 float64
	_ = v75
	var v80 float64
	_ = v80
	var v84 float64
	_ = v84
	var v93 float64
	_ = v93
	var v102 float64
	_ = v102
	var v106 float64
	_ = v106
	var v107 float64
	_ = v107
	var v115 float64
	_ = v115
	var v119 float64
	_ = v119
	var v128 int64
	_ = v128
	var v133 int32
	_ = v133
	var v146 float64
	_ = v146
	var v157 float64
	_ = v157
	var v169 float64
	_ = v169
	var v170 float64
	_ = v170
	var v171 float64
	_ = v171
	var v176 float64
	_ = v176
	var v181 float64
	_ = v181
	var v182 float64
	_ = v182
	var v183 float64
	_ = v183
	var v188 float64
	_ = v188
	var v194 float64
	_ = v194
	var v198 float64
	_ = v198
	var v201 float64
	_ = v201
	var v205 float64
	_ = v205
	var v209 float64
	_ = v209
	var v217 int64
	_ = v217
	var v222 int32
	_ = v222
	var v245 float64
	_ = v245
	var v252 float64
	_ = v252
	var v253 float64
	_ = v253
	var v254 float64
	_ = v254
	var v259 float64
	_ = v259
	var v264 float64
	_ = v264
	var v268 float64
	_ = v268
	var v277 float64
	_ = v277
	var v286 float64
	_ = v286
	var v290 float64
	_ = v290
	var v291 float64
	_ = v291
	var v299 float64
	_ = v299
	var v303 float64
	_ = v303
	var v310 int64
	_ = v310
	var v315 int32
	_ = v315
	var v328 float64
	_ = v328
	var v339 float64
	_ = v339
	var v351 float64
	_ = v351
	var v352 float64
	_ = v352
	var v353 float64
	_ = v353
	var v358 float64
	_ = v358
	var v363 float64
	_ = v363
	var v364 float64
	_ = v364
	var v365 float64
	_ = v365
	var v370 float64
	_ = v370
	var v376 float64
	_ = v376
	var v380 float64
	_ = v380
	var v383 float64
	_ = v383
	var v387 float64
	_ = v387
	var v393 float64
	_ = v393
	var v397 float64
	_ = v397
	var v401 float64
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	v13 = base.F64_abs(v12)
	if base.Ui64(base.I64_reinterpret_f64(v13)) <= base.Ui64(int64(9218868437227405312)) {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dacosd[0])))
		if v18 == int32(0) {
			F_init_degree_constants(m)
			mBase = m.M
		} else {
		}
		if base.F64_gt(v13, float64(1)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v413 = m.ExcPending
			if v413 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v416 = m.ExcPending
				if v416 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_dacosd_0), int32(0))
					mBase = m.M
					v420 = m.ExcPending
					if v420 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_dacosd_1), int32(2127), int32(_a_F_dacosd_2))
						mBase = m.M
						v425 = m.ExcPending
						if v425 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			if base.F64_ge(v12, float64(0)) != 0 {
				if base.F64_le(v12, float64(0.5)) != 0 {
					v33 = base.I64_reinterpret_f64(v12)
					v38 = base.I32_wrap_i64(int64(base.Ui64(v33)>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072693248)) <= base.Ui32(v38) {
						if base.I32_wrap_i64(v33)|(v38-int32(1072693248)) == int32(0) {
							v115 = base.F64_add(base.F64_mul(v12, float64(1.5707963267948966)), float64(7.52316384526264e-37))
						} else {
							v115 = base.F64_div(float64(0), base.F64_sub(v12, v12))
						}
					} else {
						if base.Ui32(v38) <= base.Ui32(int32(1071644671)) {
							if base.Ui32(v38+int32(-1048576)) < base.Ui32(int32(1044381696)) {
								v107 = v12
								v115 = v107
							} else {
								v61 = F_R(m, base.F64_mul(v12, v12))
								mBase = m.M
								v115 = base.F64_add(base.F64_mul(v12, v61), v12)
							}
						} else {
							v68 = base.F64_mul(base.F64_sub(float64(1), base.F64_abs(v12)), float64(0.5))
							v69 = base.F64_sqrt(v68)
							v70 = F_R(m, v68)
							mBase = m.M
							if base.Ui32(int32(1072640819)) <= base.Ui32(v38) {
								v75 = base.F64_add(base.F64_mul(v69, v70), v69)
								v102 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v75, v75), float64(-6.123233995736766e-17)))
							} else {
								v80 = float64(0.7853981633974483)
								v84 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v69) & int64(-4294967296))
								v93 = base.F64_div(base.F64_sub(v68, base.F64_mul(v84, v84)), base.F64_add(v69, v84))
								v102 = base.F64_add(base.F64_sub(base.F64_sub(v80, base.F64_add(v84, v84)), base.F64_sub(base.F64_mul(base.F64_add(v69, v69), v70), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v93, v93)))), v80)
							}
							if v33 < int64(0) {
								v106 = base.F64_neg(v102)
							} else {
								v106 = v102
							}
							v107 = v106
							v115 = v107
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v115
					v119 = *(*float64)(unsafe.Add(mBase, _c_F_dacosd[1]))
					v397 = base.F64_add(base.F64_mul(base.F64_div(v115, v119), float64(-30)), float64(90))
				} else {
					v128 = base.I64_reinterpret_f64(v12)
					v133 = base.I32_wrap_i64(int64(base.Ui64(v128)>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072693248)) <= base.Ui32(v133) {
						if base.I32_wrap_i64(v128)|(v133-int32(1072693248)) == int32(0) {
							if int64(0) <= v128 {
								v146 = float64(0)
							} else {
								v146 = float64(3.141592653589793)
							}
							v201 = v146
						} else {
							v201 = base.F64_div(float64(0), base.F64_sub(v12, v12))
						}
					} else {
						if base.Ui32(v133) <= base.Ui32(int32(1071644671)) {
							if base.Ui32(v133) < base.Ui32(int32(1012924417)) {
								v198 = float64(1.5707963267948966)
								v201 = v198
							} else {
								v157 = F_R(m, base.F64_mul(v12, v12))
								mBase = m.M
								v201 = base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(v12, v157)), v12), float64(1.5707963267948966))
							}
						} else {
							if v128 < int64(0) {
								v169 = base.F64_mul(base.F64_add(v12, float64(1)), float64(0.5))
								v170 = base.F64_sqrt(v169)
								v171 = F_R(m, v169)
								mBase = m.M
								v176 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v170, base.F64_add(base.F64_mul(v170, v171), float64(-6.123233995736766e-17))))
								v201 = base.F64_add(v176, v176)
							} else {
								v181 = base.F64_mul(base.F64_sub(float64(1), v12), float64(0.5))
								v182 = base.F64_sqrt(v181)
								v183 = F_R(m, v181)
								mBase = m.M
								v188 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v182) & int64(-4294967296))
								v194 = base.F64_add(base.F64_add(base.F64_mul(v182, v183), base.F64_div(base.F64_sub(v181, base.F64_mul(v188, v188)), base.F64_add(v182, v188))), v188)
								v198 = base.F64_add(v194, v194)
								v201 = v198
							}
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v201
					v205 = *(*float64)(unsafe.Add(mBase, _c_F_dacosd[2]))
					v397 = base.F64_mul(base.F64_div(v201, v205), float64(60))
				}
			} else {
				v209 = base.F64_neg(v12)
				if base.F64_ge(v12, float64(-0.5)) != 0 {
					v217 = base.I64_reinterpret_f64(v209)
					v222 = base.I32_wrap_i64(int64(base.Ui64(v217)>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072693248)) <= base.Ui32(v222) {
						if base.I32_wrap_i64(v217)|(v222-int32(1072693248)) == int32(0) {
							v299 = base.F64_add(base.F64_mul(v209, float64(1.5707963267948966)), float64(7.52316384526264e-37))
						} else {
							v299 = base.F64_div(float64(0), base.F64_sub(v209, v209))
						}
					} else {
						if base.Ui32(v222) <= base.Ui32(int32(1071644671)) {
							if base.Ui32(v222+int32(-1048576)) < base.Ui32(int32(1044381696)) {
								v291 = v209
								v299 = v291
							} else {
								v245 = F_R(m, base.F64_mul(v209, v209))
								mBase = m.M
								v299 = base.F64_add(base.F64_mul(v209, v245), v209)
							}
						} else {
							v252 = base.F64_mul(base.F64_sub(float64(1), base.F64_abs(v209)), float64(0.5))
							v253 = base.F64_sqrt(v252)
							v254 = F_R(m, v252)
							mBase = m.M
							if base.Ui32(int32(1072640819)) <= base.Ui32(v222) {
								v259 = base.F64_add(base.F64_mul(v253, v254), v253)
								v286 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v259, v259), float64(-6.123233995736766e-17)))
							} else {
								v264 = float64(0.7853981633974483)
								v268 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v253) & int64(-4294967296))
								v277 = base.F64_div(base.F64_sub(v252, base.F64_mul(v268, v268)), base.F64_add(v253, v268))
								v286 = base.F64_add(base.F64_sub(base.F64_sub(v264, base.F64_add(v268, v268)), base.F64_sub(base.F64_mul(base.F64_add(v253, v253), v254), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v277, v277)))), v264)
							}
							if v217 < int64(0) {
								v290 = base.F64_neg(v286)
							} else {
								v290 = v286
							}
							v291 = v290
							v299 = v291
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v299
					v303 = *(*float64)(unsafe.Add(mBase, _c_F_dacosd[1]))
					v393 = base.F64_mul(base.F64_div(v299, v303), float64(30))
				} else {
					v310 = base.I64_reinterpret_f64(v209)
					v315 = base.I32_wrap_i64(int64(base.Ui64(v310)>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072693248)) <= base.Ui32(v315) {
						if base.I32_wrap_i64(v310)|(v315-int32(1072693248)) == int32(0) {
							if int64(0) <= v310 {
								v328 = float64(0)
							} else {
								v328 = float64(3.141592653589793)
							}
							v383 = v328
						} else {
							v383 = base.F64_div(float64(0), base.F64_sub(v209, v209))
						}
					} else {
						if base.Ui32(v315) <= base.Ui32(int32(1071644671)) {
							if base.Ui32(v315) < base.Ui32(int32(1012924417)) {
								v380 = float64(1.5707963267948966)
								v383 = v380
							} else {
								v339 = F_R(m, base.F64_mul(v209, v209))
								mBase = m.M
								v383 = base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(v209, v339)), v209), float64(1.5707963267948966))
							}
						} else {
							if v310 < int64(0) {
								v351 = base.F64_mul(base.F64_add(v209, float64(1)), float64(0.5))
								v352 = base.F64_sqrt(v351)
								v353 = F_R(m, v351)
								mBase = m.M
								v358 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v352, base.F64_add(base.F64_mul(v352, v353), float64(-6.123233995736766e-17))))
								v383 = base.F64_add(v358, v358)
							} else {
								v363 = base.F64_mul(base.F64_sub(float64(1), v209), float64(0.5))
								v364 = base.F64_sqrt(v363)
								v365 = F_R(m, v363)
								mBase = m.M
								v370 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v364) & int64(-4294967296))
								v376 = base.F64_add(base.F64_add(base.F64_mul(v364, v365), base.F64_div(base.F64_sub(v363, base.F64_mul(v370, v370)), base.F64_add(v364, v370))), v370)
								v380 = base.F64_add(v376, v376)
								v383 = v380
							}
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v383
					v387 = *(*float64)(unsafe.Add(mBase, _c_F_dacosd[2]))
					v393 = base.F64_add(base.F64_mul(base.F64_div(v383, v387), float64(-60)), float64(90))
				}
				v397 = base.F64_add(v393, float64(90))
			}
			if base.F64_eq(base.F64_abs(v397), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v427 = m.ExcPending
				if v427 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v401 = v397
				v402 = F_Float8GetDatum(m, v401)
				mBase = m.M
				v405 = m.ExcPending
				if v405 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(16)
					return v402
				}
			}
		}
	} else {
		v401 = math.Float64frombits(uint64(0x7ff8000000000000))
		v402 = F_Float8GetDatum(m, v401)
		mBase = m.M
		v405 = m.ExcPending
		if v405 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v402
		}
	}
}
func F_danish_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v357 int32
	_ = v357
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = v11 + int32(3)
	if v9 < v13 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v134
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v134 < v137 {
		goto L38
	} else {
		goto L39
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v11
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v25 < v11 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v65 < int32(0) {
		goto L1
	} else {
		goto L18
	}
L4:
	;
	v27 = v11
	goto L6
L5:
	;
	v27 = v25
	goto L6
L6:
	;
	v34 = v11
	goto L8
L7:
	;
	v65 = v45
	goto L3
L8:
	;
	if v34 == v27 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v65 = int32(-1)
	goto L3
L11:
	;
	goto L12
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v34))))
	if int32(248) < v40 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v57 = v34 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v57
	v34 = v57
	goto L8
L14:
	;
	v42 = v40 - int32(97)
	if v42 < int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v45 = int32(1)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v42)>>(uint(int32(3))%32)))+uint32(_c_F_danish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v49)>>(uint(v42&int32(7))%32))&v45 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	goto L13
L18:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v77 < v76 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v120 < int32(0) {
		goto L1
	} else {
		goto L33
	}
L20:
	;
	v79 = v76
	goto L22
L21:
	;
	v79 = v77
	goto L22
L22:
	;
	v85 = v76
	goto L24
L23:
	;
	v120 = int32(1)
	goto L19
L24:
	;
	if v85 == v79 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v120 = int32(-1)
	goto L19
L27:
	;
	goto L28
L28:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v85))))
	if int32(248) < v94 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v96 = v94 - int32(97)
	if v96 < int32(0) {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v96)>>(uint(int32(3))%32)))+uint32(_c_F_danish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v102)>>(uint(v96&int32(7))%32))&int32(1) == int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v111 = v85 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v111
	v85 = v111
	goto L24
L33:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v124 = v123 + v120
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v127 < v124 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v129 = v124
	goto L36
L35:
	;
	v129 = v127
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v129
	goto L1
L37:
	;
	return v445
L38:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v235
	v237 = F_r_consonant_pair_1(m, l0)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L43
	} else {
		goto L64
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v137
	if v134 <= v137 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	goto L38
L41:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v144 = int32(1)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142+v134-v144))))
	if base.B2i32(v146&int32(224) != int32(96))|base.B2i32(v144<<(uint(v146)%32)&int32(_a_F_danish_ISO_8859_1_stem_0) == int32(0)) != 0 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v160 = F_find_among_b(m, l0, int32(_a_F_danish_ISO_8859_1_stem_1), int32(32))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	return int32(0)
L44:
	;
	if v160 == int32(0) {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v167
	switch v160 - int32(1) {
	case 0:
		goto L47
	case 1:
		goto L46
	default:
		goto L38
	}
L46:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L52
L47:
	;
	v171 = F_slice_del(m, l0)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L43
	} else {
		goto L48
	}
L48:
	;
	if int32(0) <= v171 {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	v445 = v171
	goto L37
L50:
	;
	if v227 != 0 {
		goto L38
	} else {
		goto L61
	}
L51:
	;
	v227 = v223
	goto L50
L52:
	;
	if v183 <= v184 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v223 = int32(0)
	goto L51
L54:
	;
	v227 = int32(-1)
	goto L50
L55:
	;
	goto L56
L56:
	;
	v196 = int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197+v183-v196))))
	if int32(229) < v201 {
		v223 = v196
		goto L51
	} else {
		goto L57
	}
L57:
	;
	v203 = v201 - int32(97)
	if v203 < int32(0) {
		v223 = v196
		goto L51
	} else {
		goto L58
	}
L58:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v203)>>(uint(int32(3))%32)))+uint32(_c_F_danish_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v209)>>(uint(v203&int32(7))%32))&int32(1) == int32(0) {
		v223 = v196
		goto L51
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v183 - int32(1)
	goto L60
L60:
	;
	goto L53
L61:
	;
	v228 = F_slice_del(m, l0)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L43
	} else {
		goto L62
	}
L62:
	;
	if int32(0) <= v228 {
		goto L38
	} else {
		goto L63
	}
L63:
	;
	v445 = v228
	goto L37
L64:
	;
	if v237 < int32(0) {
		v445 = v237
		goto L37
	} else {
		goto L65
	}
L65:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v241
	v244 = int32(2)
	v246 = int32(0)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v241-v249 < v244 {
		v259 = v246
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v287
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	if v287 < v290 {
		goto L79
	} else {
		goto L80
	}
L67:
	;
	if v259 == int32(0) {
		goto L66
	} else {
		goto L71
	}
L68:
	;
	goto L67
L69:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v255 = F_memcmp(m, v252+v241-v244, int32(_a_F_danish_ISO_8859_1_stem_2), v244)
	mBase = m.M
	if v255 != 0 {
		v259 = v246
		goto L68
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v241 - v244
	v259 = int32(1)
	goto L68
L71:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v262
	v264 = int32(2)
	v266 = int32(0)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v262-v269 < v264 {
		v279 = v266
		goto L73
	} else {
		goto L74
	}
L72:
	;
	if v279 == int32(0) {
		goto L66
	} else {
		goto L76
	}
L73:
	;
	goto L72
L74:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v275 = F_memcmp(m, v272+v262-v264, int32(_a_F_danish_ISO_8859_1_stem_3), v264)
	mBase = m.M
	if v275 != 0 {
		v279 = v266
		goto L73
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v262 - v264
	v279 = int32(1)
	goto L73
L76:
	;
	v282 = F_slice_del(m, l0)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L43
	} else {
		goto L77
	}
L77:
	;
	if v282 < int32(0) {
		v445 = v282
		goto L37
	} else {
		goto L78
	}
L78:
	;
	goto L66
L79:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v341
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+4))
	if v341 < v344 {
		goto L94
	} else {
		goto L95
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v287
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v290
	v296 = v287 - int32(1)
	if v296 <= v290 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v293
	goto L79
L82:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298+v296))))
	if base.B2i32(v300&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v300)%32)&int32(_a_F_danish_ISO_8859_1_stem_4) == int32(0)) != 0 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v314 = F_find_among_b(m, l0, int32(_a_F_danish_ISO_8859_1_stem_5), int32(5))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L43
	} else {
		goto L84
	}
L84:
	;
	if v314 == int32(0) {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v293
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v319
	switch v314 - int32(1) {
	case 0:
		goto L87
	case 1:
		goto L86
	default:
		goto L79
	}
L86:
	;
	v333 = F_slice_from_s(m, l0, int32(3), int32(_a_F_danish_ISO_8859_1_stem_6))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L43
	} else {
		goto L92
	}
L87:
	;
	v323 = F_slice_del(m, l0)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L43
	} else {
		goto L88
	}
L88:
	;
	if v323 < int32(0) {
		v445 = v323
		goto L37
	} else {
		goto L89
	}
L89:
	;
	v327 = F_r_consonant_pair_1(m, l0)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L43
	} else {
		goto L90
	}
L90:
	;
	if int32(0) <= v327 {
		goto L79
	} else {
		goto L91
	}
L91:
	;
	v445 = v327
	goto L37
L92:
	;
	if int32(0) <= v333 {
		goto L79
	} else {
		goto L93
	}
L93:
	;
	v445 = v333
	goto L37
L94:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v442
	v445 = int32(1)
	goto L37
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v341
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L98
L96:
	;
	if v401 != 0 {
		goto L107
	} else {
		goto L108
	}
L97:
	;
	v401 = v397
	goto L96
L98:
	;
	if v357 <= v344 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v397 = int32(0)
	goto L97
L100:
	;
	v401 = int32(-1)
	goto L96
L101:
	;
	goto L102
L102:
	;
	v370 = int32(1)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371+v357-v370))))
	if int32(122) < v375 {
		v397 = v370
		goto L97
	} else {
		goto L103
	}
L103:
	;
	v377 = v375 - int32(98)
	if v377 < int32(0) {
		v397 = v370
		goto L97
	} else {
		goto L104
	}
L104:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v377)>>(uint(int32(3))%32)))+uint32(_c_F_danish_ISO_8859_1_stem[2]))))
	if int32(base.Ui32(v383)>>(uint(v377&int32(7))%32))&int32(1) == int32(0) {
		v397 = v370
		goto L97
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v357 - int32(1)
	goto L106
L106:
	;
	goto L99
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v347
	goto L94
L108:
	;
	goto L109
L109:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v403
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	v407 = F_slice_to(m, l0, v406)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L43
	} else {
		goto L110
	}
L110:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v409))) = v407
	if v407 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	return int32(-1)
L112:
	;
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v347
	v416 = int32(0)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v407-int32(4))))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v422-v347 < v421 {
		v433 = v416
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if v433 == int32(0) {
		goto L94
	} else {
		goto L118
	}
L115:
	;
	goto L114
L116:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v429 = F_memcmp(m, v426+v422-v421, v407, v421)
	mBase = m.M
	if v429 != 0 {
		v433 = v416
		goto L115
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v422 - v421
	v433 = int32(1)
	goto L115
L118:
	;
	v436 = F_slice_del(m, l0)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L43
	} else {
		goto L119
	}
L119:
	;
	if v436 < int32(0) {
		v445 = v436
		goto L37
	} else {
		goto L120
	}
L120:
	;
	goto L94
}
func F_datan2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v27 int64
	_ = v27
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 float64
	_ = v50
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v86 float64
	_ = v86
	var v103 float64
	_ = v103
	var v104 float64
	_ = v104
	var v118 float64
	_ = v118
	var v119 float64
	_ = v119
	var v128 float64
	_ = v128
	var v135 int32
	_ = v135
	var v136 float64
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	v5 = math.Float64frombits(uint64(0x7ff8000000000000))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v7)&int64(9223372036854775807)) {
		v136 = v5
		v138 = F_Float8GetDatum(m, v136)
		mBase = m.M
		v139 = m.ExcPending
		if v139 != 0 {
			return int32(0)
		} else {
			return v138
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v13)))
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v14)&int64(9223372036854775807)) {
			v136 = v5
			v138 = F_Float8GetDatum(m, v136)
			mBase = m.M
			v139 = m.ExcPending
			if v139 != 0 {
				return int32(0)
			} else {
				return v138
			}
		} else {
			v27 = int64(9223372036854775807)
			if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v7)&v27) < base.Ui64(int64(9218868437227405313)))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v14)&v27) <= base.Ui64(int64(9218868437227405312))) == int32(0) {
				v128 = base.F64_add(v7, v14)
			} else {
				v40 = base.I64_reinterpret_f64(v14)
				v43 = base.I32_wrap_i64(int64(base.Ui64(v40) >> (uint(int64(32)) % 64)))
				v46 = base.I32_wrap_i64(v40)
				if v43-int32(1072693248)|v46 == int32(0) {
					v50 = F_atan(m, v7)
					mBase = m.M
					v128 = v50
				} else {
					v54 = int32(base.Ui32(v43)>>(uint(int32(30))%32)) & int32(2)
					v55 = base.I64_reinterpret_f64(v7)
					v59 = v54 | base.I32_wrap_i64(int64(base.Ui64(v55)>>(uint(int64(63))%64)))
					v64 = base.I32_wrap_i64(int64(base.Ui64(v55)>>(uint(int64(32))%64))) & int32(2147483647)
					if v64|base.I32_wrap_i64(v55) == int32(0) {
						switch v59 - int32(2) {
						case 0:
							v128 = float64(3.141592653589793)
						case 1:
							v128 = float64(-3.141592653589793)
						default:
							v119 = v7
							v128 = v119
						}
					} else {
						v74 = v43 & int32(2147483647)
						if v74|v46 == int32(0) {
							v128 = base.F64_copysign(float64(1.5707963267948966), v7)
						} else {
							if v74 == int32(2146435072) {
								if v64 != int32(2146435072) {
									v118 = *(*float64)(unsafe.Add(mBase, uint32(v59<<(uint(int32(3))%32))+uint32(_c_F_datan2[0])))
									v119 = v118
									v128 = v119
								} else {
									v86 = *(*float64)(unsafe.Add(mBase, uint32(v59<<(uint(int32(3))%32))+uint32(_c_F_datan2[1])))
									v128 = v86
								}
							} else {
								if base.B2i32(v64 != int32(2146435072))&base.B2i32(base.Ui32(v64) <= base.Ui32(v74+int32(67108864))) == int32(0) {
									v128 = base.F64_copysign(float64(1.5707963267948966), v7)
								} else {
									if v54 != 0 {
										if base.Ui32(v64+int32(67108864)) < base.Ui32(v74) {
											v104 = float64(0)
										} else {
											v103 = F_atan(m, base.F64_abs(base.F64_div(v7, v14)))
											mBase = m.M
											v104 = v103
										}
									} else {
										v103 = F_atan(m, base.F64_abs(base.F64_div(v7, v14)))
										mBase = m.M
										v104 = v103
									}
									switch v59 - int32(1) {
									case 0:
										v128 = base.F64_neg(v104)
									case 1:
										v128 = base.F64_sub(float64(3.141592653589793), base.F64_add(v104, float64(-1.2246467991473532e-16)))
									case 2:
										v128 = base.F64_add(base.F64_add(v104, float64(-1.2246467991473532e-16)), float64(-3.141592653589793))
									default:
										v119 = v104
										v128 = v119
									}
								}
							}
						}
					}
				}
			}
			if base.F64_ne(base.F64_abs(v128), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v136 = v128
				v138 = F_Float8GetDatum(m, v136)
				mBase = m.M
				v139 = m.ExcPending
				if v139 != 0 {
					return int32(0)
				} else {
					return v138
				}
			} else {
				F_float_overflow_error(m)
				mBase = m.M
				v135 = m.ExcPending
				if v135 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_datand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	var v18 int32
	_ = v18
	var v28 int64
	_ = v28
	var v33 int32
	_ = v33
	var v43 float64
	_ = v43
	var v49 float64
	_ = v49
	var v80 float64
	_ = v80
	var v81 int32
	_ = v81
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v97 float64
	_ = v97
	var v114 float64
	_ = v114
	var v121 int32
	_ = v121
	var v122 float64
	_ = v122
	var v125 float64
	_ = v125
	var v128 float64
	_ = v128
	var v132 float64
	_ = v132
	var v133 float64
	_ = v133
	var v143 float64
	_ = v143
	var v147 float64
	_ = v147
	var v150 float64
	_ = v150
	var v157 int32
	_ = v157
	var v158 float64
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v11)&int64(9223372036854775807)) {
		v158 = math.Float64frombits(uint64(0x7ff8000000000000))
		v159 = F_Float8GetDatum(m, v158)
		mBase = m.M
		v160 = m.ExcPending
		if v160 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v159
		}
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_datand[0])))
		if v18 == int32(0) {
			F_init_degree_constants(m)
			mBase = m.M
		} else {
		}
		v28 = base.I64_reinterpret_f64(v11)
		v33 = base.I32_wrap_i64(int64(base.Ui64(v28)>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(int32(1141899264)) <= base.Ui32(v33) {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v11)&int64(9223372036854775807)) {
				v43 = v11
			} else {
				v43 = base.F64_copysign(float64(1.5707963267948966), v11)
			}
			v143 = v43
		} else {
			if base.Ui32(v33) <= base.Ui32(int32(1071382527)) {
				if base.Ui32(int32(1044381696)) <= base.Ui32(v33) {
					v80 = v11
					v81 = int32(-1)
					v82 = base.F64_mul(v80, v80)
					v83 = base.F64_mul(v82, v82)
					v97 = base.F64_mul(v83, base.F64_add(base.F64_mul(v83, base.F64_add(base.F64_mul(v83, base.F64_add(base.F64_mul(v83, base.F64_add(base.F64_mul(v83, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
					v114 = base.F64_mul(v82, base.F64_add(base.F64_mul(v83, base.F64_add(base.F64_mul(v83, base.F64_add(base.F64_mul(v83, base.F64_add(base.F64_mul(v83, base.F64_add(base.F64_mul(v83, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
					if base.Ui32(v33) <= base.Ui32(int32(1071382527)) {
						v143 = base.F64_sub(v80, base.F64_mul(v80, base.F64_add(v97, v114)))
					} else {
						v121 = v81 << (uint(int32(3)) % 32)
						v122 = *(*float64)(unsafe.Add(mBase, uint32(v121)+uint32(_c_F_datand[1])))
						v125 = *(*float64)(unsafe.Add(mBase, uint32(v121)+uint32(_c_F_datand[2])))
						v128 = base.F64_sub(v122, base.F64_sub(base.F64_sub(base.F64_mul(v80, base.F64_add(v97, v114)), v125), v80))
						if v28 < int64(0) {
							v132 = base.F64_neg(v128)
						} else {
							v132 = v128
						}
						v133 = v132
						v143 = v133
					}
				} else {
					v133 = v11
					v143 = v133
				}
			} else {
				v49 = base.F64_abs(v11)
				if base.Ui32(v33) <= base.Ui32(int32(1072889855)) {
					if base.Ui32(v33) <= base.Ui32(int32(1072037887)) {
						v80 = base.F64_div(base.F64_add(base.F64_add(v49, v49), float64(-1)), base.F64_add(v49, float64(2)))
						v81 = int32(0)
					} else {
						v80 = base.F64_div(base.F64_add(v49, float64(-1)), base.F64_add(v49, float64(1)))
						v81 = int32(1)
					}
				} else {
					if base.Ui32(v33) <= base.Ui32(int32(1073971199)) {
						v80 = base.F64_div(base.F64_add(v49, float64(-1.5)), base.F64_add(base.F64_mul(v49, float64(1.5)), float64(1)))
						v81 = int32(2)
					} else {
						v80 = base.F64_div(float64(-1), v49)
						v81 = int32(3)
					}
				}
				v82 = base.F64_mul(v80, v80)
				v83 = base.F64_mul(v82, v82)
				v97 = base.F64_mul(v83, base.F64_add(base.F64_mul(v83, base.F64_add(base.F64_mul(v83, base.F64_add(base.F64_mul(v83, base.F64_add(base.F64_mul(v83, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
				v114 = base.F64_mul(v82, base.F64_add(base.F64_mul(v83, base.F64_add(base.F64_mul(v83, base.F64_add(base.F64_mul(v83, base.F64_add(base.F64_mul(v83, base.F64_add(base.F64_mul(v83, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
				if base.Ui32(v33) <= base.Ui32(int32(1071382527)) {
					v143 = base.F64_sub(v80, base.F64_mul(v80, base.F64_add(v97, v114)))
				} else {
					v121 = v81 << (uint(int32(3)) % 32)
					v122 = *(*float64)(unsafe.Add(mBase, uint32(v121)+uint32(_c_F_datand[1])))
					v125 = *(*float64)(unsafe.Add(mBase, uint32(v121)+uint32(_c_F_datand[2])))
					v128 = base.F64_sub(v122, base.F64_sub(base.F64_sub(base.F64_mul(v80, base.F64_add(v97, v114)), v125), v80))
					if v28 < int64(0) {
						v132 = base.F64_neg(v128)
					} else {
						v132 = v128
					}
					v133 = v132
					v143 = v133
				}
			}
		}
		*(*float64)(unsafe.Add(mBase, uint32(v7)+8)) = v143
		v147 = *(*float64)(unsafe.Add(mBase, _c_F_datand[3]))
		v150 = base.F64_mul(base.F64_div(v143, v147), float64(45))
		if base.F64_ne(base.F64_abs(v150), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v158 = v150
			v159 = F_Float8GetDatum(m, v158)
			mBase = m.M
			v160 = m.ExcPending
			if v160 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return v159
			}
		} else {
			F_float_overflow_error(m)
			mBase = m.M
			v157 = m.ExcPending
			if v157 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_daterange_canonical(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L39
	}
L4:
	;
	F_range_deserialize(m, v30, v12, v9+int32(24), v9+int32(16), v9+int32(15))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v20 == v17 {
		v30 = v19
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v23 = F_lookup_type_cache(m, v17, int32(2048))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
	if v25 == int32(0) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
	v30 = v23
	goto L4
L11:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
	if v39 == int32(1) {
		v126 = v12
		goto L12
	} else {
		goto L13
	}
L12:
	;
	m.G0 = v9 + int32(32)
	return v126
L13:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
	if v42 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
	if v79 != 0 {
		goto L26
	} else {
		goto L27
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if base.B2i32(v43 == int32(-2147483648))|base.B2i32(v43 == int32(2147483647)) != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
	if v49&int32(1) != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	if base.Ui32(int32(2147483494)) <= base.Ui32(v43+int32(_a_F_daterange_canonical_0)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v56 = int32(0)
	v57 = F_errsave_start(m, v16)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v73 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)) = uint8(v73)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v43 + v73
	goto L14
L21:
	;
	if v57 == int32(0) {
		v126 = v56
		goto L12
	} else {
		goto L22
	}
L22:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errmsg(m, int32(_a_F_daterange_canonical_1), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errsave_finish(m, v16, int32(_a_F_daterange_canonical_2), int32(1648), int32(_a_F_daterange_canonical_3))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v126 = v56
	goto L12
L26:
	;
	v123 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L38
	}
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if base.B2i32(v80 == int32(-2147483648))|base.B2i32(v80 == int32(2147483647)) != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
	if v86&int32(1) == int32(0) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	if base.Ui32(int32(2147483494)) <= base.Ui32(v80+int32(_a_F_daterange_canonical_0)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v95 = int32(0)
	v96 = F_errsave_start(m, v16)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v112)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v80 + int32(1)
	goto L26
L33:
	;
	if v96 == int32(0) {
		v126 = v95
		goto L12
	} else {
		goto L34
	}
L34:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errmsg(m, int32(_a_F_daterange_canonical_1), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errsave_finish(m, v16, int32(_a_F_daterange_canonical_2), int32(1663), int32(_a_F_daterange_canonical_3))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v126 = v95
	goto L12
L38:
	;
	v126 = v123
	goto L12
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
	F_errmsg_internal(m, int32(_a_F_daterange_canonical_4), v9)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_daterange_canonical_2), int32(1776), int32(_a_F_daterange_canonical_5))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_datetime_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v5 == int32(-2147483648) {
		v9 = F_Int64GetDatum(m, int64(-9223372036854775807-1))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v9
		}
	} else {
		if v5 == int32(2147483647) {
			v17 = F_Int64GetDatum(m, int64(9223372036854775807))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v17
			}
		} else {
			if v5 < int32(106751983) {
				v22 = int64(-9223372036854775807 - 1)
				v25 = base.I64_extend_i32_s(v5) * int64(86400000000)
				if v25 != v22 {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
					v30 = v25 + v29
					if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v30+int64(211813488000000000)) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_datetime_timestamp_0), int32(0))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_datetime_timestamp_1), int32(2044), int32(_a_F_datetime_timestamp_2))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v35 = v30
						v36 = F_Int64GetDatum(m, v35)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							return v36
						}
					}
				} else {
					v35 = v22
					v36 = F_Int64GetDatum(m, v35)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						return v36
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_datetime_timestamp_3), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_datetime_timestamp_1), int32(658), int32(_a_F_datetime_timestamp_4))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_datetimetz_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v27 int64
	_ = v27
	var v30 int64
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == int32(-2147483648) {
		v8 = F_Int64GetDatum(m, int64(-9223372036854775807-1))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v8
		}
	} else {
		if v4 == int32(2147483647) {
			v16 = F_Int64GetDatum(m, int64(9223372036854775807))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v16
			}
		} else {
			if v4 < int32(106751983) {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
				v27 = int64(*(*int32)(unsafe.Add(mBase, uint32(v21)+8)))
				v30 = v22 + base.I64_extend_i32_s(v4)*int64(86400000000) + v27*int64(1000000)
				if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v30+int64(211813488000000000)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_datetimetz_timestamptz_0), int32(0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_datetimetz_timestamptz_1), int32(2981), int32(_a_F_datetimetz_timestamptz_2))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v35 = F_Int64GetDatum(m, v30)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						return v35
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_datetimetz_timestamptz_0), int32(0))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_datetimetz_timestamptz_1), int32(2971), int32(_a_F_datetimetz_timestamptz_2))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_dbase_identify(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	v2 = int32(0)
	v4 = l0 & int32(240)
	switch v4 - int32(16) {
	case 0:
		return int32(_a_F_dbase_identify_0)
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		v12 = v2
		return v12
	case 16:
		v12 = int32(_a_F_dbase_identify_1)
		return v12
	default:
		if v4 != 0 {
			v12 = v2
			return v12
		} else {
			return int32(_a_F_dbase_identify_2)
		}
	}
}
func F_dcotd(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 float64
	_ = v43
	var v46 int64
	_ = v46
	var v53 float64
	_ = v53
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v94 int64
	_ = v94
	var v97 int32
	_ = v97
	var v99 int64
	_ = v99
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v110 int32
	_ = v110
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v127 int64
	_ = v127
	var v131 int64
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v142 int64
	_ = v142
	var v145 int32
	_ = v145
	var v160 int64
	_ = v160
	var v168 float64
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 float64
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 float64
	_ = v182
	var v185 int32
	_ = v185
	var v186 float64
	_ = v186
	var v191 float64
	_ = v191
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v211 float64
	_ = v211
	var v215 int32
	_ = v215
	var v216 float64
	_ = v216
	var v217 float64
	_ = v217
	var v222 float64
	_ = v222
	var v224 float64
	_ = v224
	var v226 float64
	_ = v226
	var v229 float64
	_ = v229
	var v233 float64
	_ = v233
	var v237 float64
	_ = v237
	var v241 float64
	_ = v241
	var v250 float64
	_ = v250
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v270 float64
	_ = v270
	var v274 int32
	_ = v274
	var v275 float64
	_ = v275
	var v276 float64
	_ = v276
	var v282 float64
	_ = v282
	var v283 float64
	_ = v283
	var v285 float64
	_ = v285
	var v287 float64
	_ = v287
	var v289 float64
	_ = v289
	var v298 float64
	_ = v298
	var v302 float64
	_ = v302
	var v303 int32
	_ = v303
	var v307 float64
	_ = v307
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v327 float64
	_ = v327
	var v331 int32
	_ = v331
	var v332 float64
	_ = v332
	var v333 float64
	_ = v333
	var v339 float64
	_ = v339
	var v340 float64
	_ = v340
	var v342 float64
	_ = v342
	var v344 float64
	_ = v344
	var v346 float64
	_ = v346
	var v355 float64
	_ = v355
	var v363 float64
	_ = v363
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v383 float64
	_ = v383
	var v387 int32
	_ = v387
	var v388 float64
	_ = v388
	var v389 float64
	_ = v389
	var v394 float64
	_ = v394
	var v396 float64
	_ = v396
	var v398 float64
	_ = v398
	var v401 float64
	_ = v401
	var v405 float64
	_ = v405
	var v409 float64
	_ = v409
	var v413 float64
	_ = v413
	var v419 float64
	_ = v419
	var v420 float64
	_ = v420
	var v422 float64
	_ = v422
	var v425 float64
	_ = v425
	var v428 float64
	_ = v428
	var v431 float64
	_ = v431
	var v437 float64
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	if base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L127
	} else {
		goto L129
	}
L2:
	;
	if base.F64_eq(base.F64_abs(v12), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	v437 = math.Float64frombits(uint64(0x7ff8000000000000))
	goto L4
L4:
	;
	v438 = F_Float8GetDatum(m, v437)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L127
	} else {
		goto L128
	}
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dcotd[0])))
	if v22 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_init_degree_constants(m)
	mBase = m.M
	goto L8
L7:
	;
	goto L8
L8:
	;
	v26 = int32(0)
	v34 = base.I64_reinterpret_f64(v12)
	v38 = int32(2047)
	v39 = base.I32_wrap_i64(int64(base.Ui64(v34)>>(uint(int64(52))%64))) & v38
	if v39 == v38 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v170 = base.F64_lt(v168, float64(0))
	if v170 != 0 {
		goto L50
	} else {
		goto L51
	}
L10:
	;
	v43 = base.F64_mul(v12, float64(360))
	v168 = base.F64_div(v43, v43)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v46 = v34 << (uint(int64(1)) % 64)
	if base.Ui64(v46) <= base.Ui64(int64(-9156662467374350336)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v46 == int64(-9156662467374350336) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	if v39 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v53 = base.F64_mul(v12, float64(0))
	goto L18
L17:
	;
	v53 = v12
	goto L18
L18:
	;
	v168 = v53
	goto L9
L19:
	;
	if int32(1031) < v89 {
		goto L29
	} else {
		goto L30
	}
L20:
	;
	v56 = int32(0)
	v58 = v34 << (uint(int64(12)) % 64)
	if int64(0) <= v58 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v89 = v39
	v90 = v34&int64(4503599627370495) | int64(4503599627370496)
	goto L19
L23:
	;
	v62 = v58
	v65 = v56
	goto L26
L24:
	;
	v76 = v56
	goto L25
L25:
	;
	v89 = v76
	v90 = v34 << (uint(base.I64_extend_i32_u(int32(1)-v76)) % 64)
	goto L19
L26:
	;
	v67 = v65 - int32(1)
	v69 = v62 << (uint(int64(1)) % 64)
	if int64(0) <= v69 {
		v62 = v69
		v65 = v67
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v76 = v67
	goto L25
L28:
	;
	goto L27
L29:
	;
	v94 = v90
	v97 = v89
	goto L32
L30:
	;
	v115 = v90
	v118 = v89
	goto L31
L31:
	;
	v120 = v115 - int64(6333186975989760)
	if v120 < int64(0) {
		v127 = v115
		goto L38
	} else {
		goto L39
	}
L32:
	;
	v99 = v94 - int64(6333186975989760)
	if v99 < int64(0) {
		v106 = v94
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v115 = v108
	v118 = int32(1031)
	goto L31
L34:
	;
	v108 = v106 << (uint(int64(1)) % 64)
	v110 = v97 - int32(1)
	if int32(1031) < v110 {
		v94 = v108
		v97 = v110
		goto L32
	} else {
		goto L37
	}
L35:
	;
	if v99 != int64(0) {
		v106 = v99
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v168 = base.F64_mul(v12, float64(0))
	goto L9
L37:
	;
	goto L33
L38:
	;
	if base.Ui64(v127) <= base.Ui64(int64(4503599627370495)) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	if v120 != int64(0) {
		v127 = v120
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v168 = base.F64_mul(v12, float64(0))
	goto L9
L41:
	;
	v131 = v127
	v134 = v118
	goto L44
L42:
	;
	v142 = v127
	v145 = v118
	goto L43
L43:
	;
	if int32(0) < v145 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v136 = v134 - int32(1)
	v138 = v131 << (uint(int64(1)) % 64)
	if base.Ui64(v131) < base.Ui64(int64(2251799813685248)) {
		v131 = v138
		v134 = v136
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v142 = v138
	v145 = v136
	goto L43
L46:
	;
	goto L45
L47:
	;
	v160 = v142 - int64(4503599627370496) | base.I64_extend_i32_u(v145)<<(uint(int64(52))%64)
	goto L49
L48:
	;
	v160 = int64(base.Ui64(v142) >> (uint(base.I64_extend_i32_u(int32(1)-v145)) % 64))
	goto L49
L49:
	;
	v168 = base.F64_reinterpret_i64(v34&int64(-9223372036854775807-1) | v160)
	goto L9
L50:
	;
	v171 = int32(-1)
	goto L52
L51:
	;
	v171 = int32(1)
	goto L52
L52:
	;
	if v170 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v174 = base.F64_neg(v168)
	goto L55
L54:
	;
	v174 = v168
	goto L55
L55:
	;
	v176 = base.F64_gt(v174, float64(180))
	if v176 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v177 = v26 - v171
	goto L58
L57:
	;
	v177 = v171
	goto L58
L58:
	;
	if v176 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v185 != 0 {
		goto L93
	} else {
		goto L94
	}
L60:
	;
	v182 = base.F64_sub(float64(360), v174)
	goto L62
L61:
	;
	v182 = v174
	goto L62
L62:
	;
	v185 = base.F64_gt(v182, float64(90))
	if v185 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v186 = base.F64_sub(float64(180), v182)
	goto L65
L64:
	;
	v186 = v182
	goto L65
L65:
	;
	if base.F64_le(v186, float64(60)) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v191 = base.F64_mul(v186, float64(0.017453292519943295))
	v195 = m.G0
	v197 = v195 - int32(16)
	m.G0 = v197
	v204 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v191))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v204) <= base.Ui32(int32(1072243195)) {
		goto L71
	} else {
		goto L72
	}
L67:
	;
	goto L68
L68:
	;
	v250 = base.F64_mul(base.F64_sub(float64(90), v186), float64(0.017453292519943295))
	v254 = m.G0
	v256 = v254 - int32(16)
	m.G0 = v256
	v263 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v250))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v263) <= base.Ui32(int32(1072243195)) {
		goto L82
	} else {
		goto L83
	}
L69:
	;
	v237 = base.F64_sub(float64(1), v233)
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v237
	v241 = *(*float64)(unsafe.Add(mBase, _c_F_dcotd[1]))
	v302 = base.F64_add(base.F64_mul(base.F64_div(v237, v241), float64(-0.5)), float64(1))
	goto L59
L70:
	;
	m.G0 = v197 + int32(16)
	goto L69
L71:
	;
	if base.Ui32(v204) < base.Ui32(int32(1044816030)) {
		v233 = float64(1)
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v204) {
		v233 = base.F64_sub(v191, v191)
		goto L70
	} else {
		goto L75
	}
L74:
	;
	v211 = F___cos(m, v191, float64(0))
	mBase = m.M
	v233 = v211
	goto L70
L75:
	;
	v215 = F___rem_pio2(m, v191, v197)
	mBase = m.M
	v216 = *(*float64)(unsafe.Add(mBase, uint32(v197)+8))
	v217 = *(*float64)(unsafe.Add(mBase, uint32(v197)))
	switch v215&int32(3) - int32(1) {
	case 0:
		goto L78
	case 1:
		goto L77
	case 2:
		goto L76
	default:
		goto L79
	}
L76:
	;
	v229 = F___sin(m, v217, v216, int32(1))
	mBase = m.M
	v233 = v229
	goto L70
L77:
	;
	v226 = F___cos(m, v217, v216)
	mBase = m.M
	v233 = base.F64_neg(v226)
	goto L70
L78:
	;
	v224 = F___sin(m, v217, v216, int32(1))
	mBase = m.M
	v233 = base.F64_neg(v224)
	goto L70
L79:
	;
	v222 = F___cos(m, v217, v216)
	mBase = m.M
	v233 = v222
	goto L70
L80:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v289
	v298 = *(*float64)(unsafe.Add(mBase, _c_F_dcotd[2]))
	v302 = base.F64_mul(base.F64_div(v289, v298), float64(0.5))
	goto L59
L81:
	;
	m.G0 = v256 + int32(16)
	goto L80
L82:
	;
	if base.Ui32(v263) < base.Ui32(int32(1045430272)) {
		v289 = v250
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v263) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v270 = F___sin(m, v250, float64(0), int32(0))
	mBase = m.M
	v289 = v270
	goto L81
L86:
	;
	v289 = base.F64_sub(v250, v250)
	goto L81
L87:
	;
	goto L88
L88:
	;
	v274 = F___rem_pio2(m, v250, v256)
	mBase = m.M
	v275 = *(*float64)(unsafe.Add(mBase, uint32(v256)+8))
	v276 = *(*float64)(unsafe.Add(mBase, uint32(v256)))
	switch v274&int32(3) - int32(1) {
	case 0:
		goto L91
	case 1:
		goto L90
	case 2:
		goto L89
	default:
		goto L92
	}
L89:
	;
	v287 = F___cos(m, v276, v275)
	mBase = m.M
	v289 = base.F64_neg(v287)
	goto L81
L90:
	;
	v285 = F___sin(m, v276, v275, int32(1))
	mBase = m.M
	v289 = base.F64_neg(v285)
	goto L81
L91:
	;
	v283 = F___cos(m, v276, v275)
	mBase = m.M
	v289 = v283
	goto L81
L92:
	;
	v282 = F___sin(m, v276, v275, int32(1))
	mBase = m.M
	v289 = v282
	goto L81
L93:
	;
	v303 = v26 - v177
	goto L95
L94:
	;
	v303 = v177
	goto L95
L95:
	;
	if base.F64_le(v186, float64(30)) != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v420 = base.F64_div(v302, v419)
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v420
	v422 = float64(0)
	v425 = *(*float64)(unsafe.Add(mBase, _c_F_dcotd[3]))
	v428 = base.F64_mul(base.F64_div(v420, v425), base.F64_convert_i32_s(v303))
	if base.F64_eq(v428, v422) != 0 {
		goto L124
	} else {
		goto L125
	}
L97:
	;
	v307 = base.F64_mul(v186, float64(0.017453292519943295))
	v311 = m.G0
	v313 = v311 - int32(16)
	m.G0 = v313
	v320 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v307))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v320) <= base.Ui32(int32(1072243195)) {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	goto L99
L99:
	;
	v363 = base.F64_mul(base.F64_sub(float64(90), v186), float64(0.017453292519943295))
	v367 = m.G0
	v369 = v367 - int32(16)
	m.G0 = v369
	v376 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v363))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v376) <= base.Ui32(int32(1072243195)) {
		goto L115
	} else {
		goto L116
	}
L100:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v346
	v355 = *(*float64)(unsafe.Add(mBase, _c_F_dcotd[2]))
	v419 = base.F64_mul(base.F64_div(v346, v355), float64(0.5))
	goto L96
L101:
	;
	m.G0 = v313 + int32(16)
	goto L100
L102:
	;
	if base.Ui32(v320) < base.Ui32(int32(1045430272)) {
		v346 = v307
		goto L101
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v320) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v327 = F___sin(m, v307, float64(0), int32(0))
	mBase = m.M
	v346 = v327
	goto L101
L106:
	;
	v346 = base.F64_sub(v307, v307)
	goto L101
L107:
	;
	goto L108
L108:
	;
	v331 = F___rem_pio2(m, v307, v313)
	mBase = m.M
	v332 = *(*float64)(unsafe.Add(mBase, uint32(v313)+8))
	v333 = *(*float64)(unsafe.Add(mBase, uint32(v313)))
	switch v331&int32(3) - int32(1) {
	case 0:
		goto L111
	case 1:
		goto L110
	case 2:
		goto L109
	default:
		goto L112
	}
L109:
	;
	v344 = F___cos(m, v333, v332)
	mBase = m.M
	v346 = base.F64_neg(v344)
	goto L101
L110:
	;
	v342 = F___sin(m, v333, v332, int32(1))
	mBase = m.M
	v346 = base.F64_neg(v342)
	goto L101
L111:
	;
	v340 = F___cos(m, v333, v332)
	mBase = m.M
	v346 = v340
	goto L101
L112:
	;
	v339 = F___sin(m, v333, v332, int32(1))
	mBase = m.M
	v346 = v339
	goto L101
L113:
	;
	v409 = base.F64_sub(float64(1), v405)
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v409
	v413 = *(*float64)(unsafe.Add(mBase, _c_F_dcotd[1]))
	v419 = base.F64_add(base.F64_mul(base.F64_div(v409, v413), float64(-0.5)), float64(1))
	goto L96
L114:
	;
	m.G0 = v369 + int32(16)
	goto L113
L115:
	;
	if base.Ui32(v376) < base.Ui32(int32(1044816030)) {
		v405 = float64(1)
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v376) {
		v405 = base.F64_sub(v363, v363)
		goto L114
	} else {
		goto L119
	}
L118:
	;
	v383 = F___cos(m, v363, float64(0))
	mBase = m.M
	v405 = v383
	goto L114
L119:
	;
	v387 = F___rem_pio2(m, v363, v369)
	mBase = m.M
	v388 = *(*float64)(unsafe.Add(mBase, uint32(v369)+8))
	v389 = *(*float64)(unsafe.Add(mBase, uint32(v369)))
	switch v387&int32(3) - int32(1) {
	case 0:
		goto L122
	case 1:
		goto L121
	case 2:
		goto L120
	default:
		goto L123
	}
L120:
	;
	v401 = F___sin(m, v389, v388, int32(1))
	mBase = m.M
	v405 = v401
	goto L114
L121:
	;
	v398 = F___cos(m, v389, v388)
	mBase = m.M
	v405 = base.F64_neg(v398)
	goto L114
L122:
	;
	v396 = F___sin(m, v389, v388, int32(1))
	mBase = m.M
	v405 = base.F64_neg(v396)
	goto L114
L123:
	;
	v394 = F___cos(m, v389, v388)
	mBase = m.M
	v405 = v394
	goto L114
L124:
	;
	v431 = v422
	goto L126
L125:
	;
	v431 = v428
	goto L126
L126:
	;
	v437 = v431
	goto L4
L127:
	;
	return int32(0)
L128:
	;
	m.G0 = v9 + int32(16)
	return v438
L129:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L127
	} else {
		goto L130
	}
L130:
	;
	F_errmsg(m, int32(_a_F_dcotd_0), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L127
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_dcotd_1), int32(2390), int32(_a_F_dcotd_2))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L127
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_debugStartup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v90 int32
	_ = v90
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v17 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_pg_printf(m, int32(_a_F_debugStartup_0), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L9
	} else {
		goto L12
	}
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v47 = l2 + v41<<(uint(int32(4))%32) + v30*int32(100)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+102)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+88))
	v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+92)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(28)))) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(20)))) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(16)))) = int32(_a_F_debugStartup_1)
	if v48 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v59 = int32(116)
	goto L8
L7:
	;
	v59 = int32(102)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(32)))) = v59
	v61 = int32(_a_F_debugStartup_1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v47 + int32(24)
	v69 = v30 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v69
	F_pg_printf(m, int32(_a_F_debugStartup_2), v15)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	if v69 != v17 {
		v30 = v69
		goto L4
	} else {
		goto L11
	}
L11:
	;
	goto L5
L12:
	;
	m.G0 = v15 + int32(48)
	return
}
func F_debugtup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v33 int32
	_ = v33
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v138 int32
	_ = v138
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if int32(0) < v20 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_pg_printf(m, int32(_a_F_debugtup_0), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L9
	} else {
		goto L30
	}
L4:
	;
	v47 = v33 + int32(1)
	v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v48 <= v33 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	F_slot_getsomeattrs_int(m, l0, v47)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v33))))
	if v56 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return int32(0)
L10:
	;
	goto L8
L11:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+v33<<(uint(int32(2))%32))))
	v65 = v33 * int32(100)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v65+(v19+v66<<(uint(int32(4))%32)))+88))
	F_getTypeOutputInfo(m, v71, v17+int32(44), v17+int32(43))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v47 != v20 {
		v33 = v47
		goto L4
	} else {
		goto L29
	}
L14:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	v79 = F_OidOutputFunctionCall(m, v78, v63)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v85 = v19 + v81<<(uint(int32(4))%32) + v65
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+102)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v85)+88))
	v88 = int32(*(*int16)(unsafe.Add(mBase, uint32(v85)+92)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(28)))) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(20)))) = v87
	if v79 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v95 = int32(_a_F_debugtup_1)
	goto L18
L17:
	;
	v95 = int32(_a_F_debugtup_2)
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(16)))) = v95
	if v86 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v99 = int32(116)
	goto L21
L20:
	;
	v99 = int32(102)
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(32)))) = v99
	if v79 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v102 = v79
	goto L24
L23:
	;
	v102 = int32(_a_F_debugtup_2)
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v102
	if v79 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v106 = int32(_a_F_debugtup_3)
	goto L27
L26:
	;
	v106 = int32(_a_F_debugtup_2)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v85 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v47
	F_pg_printf(m, int32(_a_F_debugtup_4), v17)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	goto L13
L29:
	;
	goto L5
L30:
	;
	m.G0 = v17 + int32(48)
	return int32(1)
}
func F_decompile_column_index_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_deconstruct_array_builtin(m, v12, int32(21), v10+int32(12), int32(0), v10+int32(8))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v25 <= int32(0) {
		v80 = v25
		goto L4
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v10 + int32(16)
	return v80
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28))))
	v31 = F_get_attname(m, l1, v29, int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v33 = F_quote_identifier(m, v31)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_appendStringInfoString(m, l3, v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v37 < int32(2) {
		v80 = v37
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v40 = int32(1)
	goto L10
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47+v40<<(uint(int32(2))%32)))))
	v53 = F_get_attname(m, l1, v51, int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v80 = v73
	goto L4
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v56 = F_quote_identifier(m, v53)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v56
	if v40 == v55-int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v64 = int32(_a_F_decompile_column_index_array_0)
	goto L16
L15:
	;
	v64 = int32(_a_F_decompile_column_index_array_1)
	goto L16
L16:
	;
	if l2 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v66 = v64
	goto L19
L18:
	;
	v66 = int32(_a_F_decompile_column_index_array_1)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v66
	F_appendStringInfo(m, l3, int32(_a_F_decompile_column_index_array_2), v10)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v72 = v40 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v72 < v73 {
		v40 = v72
		goto L10
	} else {
		goto L21
	}
L21:
	;
	goto L11
}
func F_decompress_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		m.Env.Pgmem_zstream_free(m, v5)
		mBase = m.M
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		if v7 != 0 {
			F_ResourceOwnerForget(m, v7, v4, int32(_a_F_decompress_free_0))
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				F_pfree(m, v4)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					base.MemoryFill(m, l0, int32(0), int32(_a_F_decompress_free_1))
					F_pfree(m, l0)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			F_pfree(m, v4)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				base.MemoryFill(m, l0, int32(0), int32(_a_F_decompress_free_1))
				F_pfree(m, l0)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		base.MemoryFill(m, l0, int32(0), int32(_a_F_decompress_free_1))
		F_pfree(m, l0)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			return
		}
	}
}
func F_deconstruct_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v309 int32
	_ = v309
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v598 int32
	_ = v598
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v757 int32
	_ = v757
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v980 int32
	_ = v980
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	v6 = int32(0)
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v19 = F_palloc0(m, int32(44))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l1
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v25 - int32(63) {
	case 0:
		goto L5
	case 1:
		goto L7
	case 2:
		goto L8
	default:
		goto L6
	}
L3:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v986 = F_lappend(m, v985, v19)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L1
	} else {
		goto L205
	}
L4:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v964 == int32(0) {
		v980 = v959
		goto L3
	} else {
		goto L203
	}
L5:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v930 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v931 = F_bms_add_member(m, v929, v930)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L1
	} else {
		goto L199
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L1
	} else {
		goto L196
	}
L7:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v104 {
	case 0:
		goto L33
	case 1, 5:
		goto L37
	case 2:
		goto L35
	default:
		goto L34
	case 4:
		goto L36
	}
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l2
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v31 == int32(0) {
		v980 = v6
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v34 <= int32(0) {
		v959 = v6
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v44 = v37
	v46 = v6
	v47 = v6
	goto L11
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+v47<<(uint(int32(2))%32))))
	v56 = F_deconstruct_recurse(m, l0, v55, l2, v19, l4)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v959 = v99
	goto L4
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v60+v61<<(uint(int32(2))%32)-int32(4))))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v69 = F_bms_add_members(m, v58, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v69
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v72
	if v56 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v101 = v47 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v101 < v102 {
		v44 = v97
		v46 = v99
		v47 = v101
		goto L11
	} else {
		goto L30
	}
L16:
	;
	v78 = F_list_concat(m, v46, v56)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v81 = v44 - int32(1)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if int32(2) <= v82 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v97 = v44 - int32(1)
	v99 = v78
	goto L15
L20:
	;
	v95 = F_lappend(m, v46, v56)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L29
	}
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_deconstruct_recurse[0]))
	if v46 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v93 = F_list_concat(m, v46, v56)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L28
	}
L24:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v89 = v87
	goto L26
L25:
	;
	v89 = int32(0)
	goto L26
L26:
	;
	if v86 < v89+(v81+v82) {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v97 = v81
	v99 = v93
	goto L15
L29:
	;
	v97 = v81
	v99 = v95
	goto L15
L30:
	;
	goto L12
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v853
	v855 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v855 == int32(2) {
		goto L170
	} else {
		goto L171
	}
L32:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v832)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v835
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v830)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v837
	v846 = v833
	v848 = v834
	v853 = int32(0)
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l2
	v800 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v801 = F_deconstruct_recurse(m, l0, v800, l2, v19, l4)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L167
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L1
	} else {
		goto L164
	}
L35:
	;
	v371 = F_palloc0(m, int32(8))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L84
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l2
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v337 = F_deconstruct_recurse(m, l0, v336, l2, v19, l4)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L80
	}
L37:
	;
	v106 = F_palloc0(m, int32(8))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v106))) = int64(272)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v111 = F_lappend(m, v110, v106)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v106
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v116 = F_deconstruct_recurse(m, l0, v115, l2, v19, l4)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v119+v120<<(uint(int32(2))%32)-int32(4))))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v128 = F_deconstruct_recurse(m, l0, v127, v106, v19, l4)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v131+v132<<(uint(int32(2))%32)-int32(4))))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v141 = F_bms_add_members(m, v139, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v141
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	v146 = F_bms_union(m, v144, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v146
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v149 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	v327 = F_bms_union(m, v325, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L79
	}
L45:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v153 = F_bms_add_member(m, v152, v149)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v153
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v158 = F_bms_add_member(m, v156, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v158
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v163 = F_bms_add_member(m, v161, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v163
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	if v167 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	if v224 <= int32(0) {
		goto L44
	} else {
		goto L60
	}
L50:
	;
	v224 = base.I32_ctz(v210) | v211<<(uint(int32(5))%32)
	goto L49
L51:
	;
	v224 = int32(-2)
	goto L49
L52:
	;
	v177 = base.I32_div_s(int32(0), int32(32))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	if v178 <= v177 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v181 = v167 + int32(8)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v181+v177<<(uint(int32(2))%32))))
	v188 = v185 & int32(-1)
	if v188 != 0 {
		v210 = v188
		v211 = v177
		goto L50
	} else {
		goto L54
	}
L54:
	;
	v190 = v177 + int32(1)
	if v190 == v178 {
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v193 = v190
	goto L56
L56:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v181+v193<<(uint(int32(2))%32))))
	if v200 != 0 {
		v210 = v200
		v211 = v193
		goto L50
	} else {
		goto L58
	}
L57:
	;
	goto L51
L58:
	;
	v202 = v193 + int32(1)
	if v202 != v178 {
		v193 = v202
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v230 = v224
	goto L61
L61:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v230 == v240 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L44
L63:
	;
	if v167 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242+v230<<(uint(int32(2))%32))))
	if v246 == int32(0) {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v246)+96))
	v250 = F_bms_add_member(m, v249, v166)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+96)) = v250
	goto L63
L67:
	;
	if int32(0) < v309 {
		v230 = v309
		goto L61
	} else {
		goto L78
	}
L68:
	;
	v309 = base.I32_ctz(v295) | v296<<(uint(int32(5))%32)
	goto L67
L69:
	;
	v309 = int32(-2)
	goto L67
L70:
	;
	v260 = v230 + int32(1)
	v262 = base.I32_div_s(v260, int32(32))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	if v263 <= v262 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v266 = v167 + int32(8)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v266+v262<<(uint(int32(2))%32))))
	v273 = v270 & (int32(-1) << (uint(v260) % 32))
	if v273 != 0 {
		v295 = v273
		v296 = v262
		goto L68
	} else {
		goto L72
	}
L72:
	;
	v275 = v262 + int32(1)
	if v275 == v263 {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v278 = v275
	goto L74
L74:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v266+v278<<(uint(int32(2))%32))))
	if v285 != 0 {
		v295 = v285
		v296 = v278
		goto L68
	} else {
		goto L76
	}
L75:
	;
	goto L69
L76:
	;
	v287 = v278 + int32(1)
	if v287 != v263 {
		v278 = v287
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	goto L62
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v327
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v330
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v332
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v846 = v128
	v848 = v116
	v853 = v334
	goto L31
L80:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)+12))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v340+v341<<(uint(int32(2))%32)-int32(4))))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v349 = F_deconstruct_recurse(m, l0, v348, l2, v19, l4)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+12))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v353+v354<<(uint(int32(2))%32)-int32(4))))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v360)+12))
	v362 = F_bms_union(m, v351, v361)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v362
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v347)+16))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v360)+16))
	v367 = F_bms_union(m, v365, v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v367
	v830 = v360
	v832 = v347
	v833 = v349
	v834 = v337
	goto L32
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = int32(272)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v376 = F_lappend(m, v375, v371)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v376
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v371
	v381 = F_palloc0(m, int32(8))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v381))) = int64(272)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v386 = F_lappend(m, v385, v381)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v386
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v390 = F_deconstruct_recurse(m, l0, v389, v381, v19, l4)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+12))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v392)+4))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v393+v394<<(uint(int32(2))%32)-int32(4))))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
	v402 = F_bms_copy(m, v401)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371)+4)) = v402
	v406 = F_palloc0(m, int32(8))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v406))) = int64(272)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v411 = F_lappend(m, v410, v406)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v411
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v415 = F_deconstruct_recurse(m, l0, v414, v406, v19, l4)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+12))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v417)+4))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v418+v419<<(uint(int32(2))%32)-int32(4))))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v406)+4))
	v428 = F_bms_add_members(m, v426, v427)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371)+4)) = v428
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v432 = F_bms_add_members(m, v431, v428)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v432
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v400)+12))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v425)+12))
	v437 = F_bms_union(m, v435, v436)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v437
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v442 = F_bms_add_member(m, v440, v441)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v442
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v447 = F_bms_add_member(m, v445, v446)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v447
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v452 = F_bms_add_member(m, v450, v451)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v452
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v400)+12))
	if v456 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	if int32(0) < v513 {
		goto L110
	} else {
		goto L111
	}
L100:
	;
	v513 = base.I32_ctz(v499) | v500<<(uint(int32(5))%32)
	goto L99
L101:
	;
	v513 = int32(-2)
	goto L99
L102:
	;
	v466 = base.I32_div_s(int32(0), int32(32))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v456)+4))
	if v467 <= v466 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v470 = v456 + int32(8)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v470+v466<<(uint(int32(2))%32))))
	v477 = v474 & int32(-1)
	if v477 != 0 {
		v499 = v477
		v500 = v466
		goto L100
	} else {
		goto L104
	}
L104:
	;
	v479 = v466 + int32(1)
	if v479 == v467 {
		goto L101
	} else {
		goto L105
	}
L105:
	;
	v482 = v479
	goto L106
L106:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v470+v482<<(uint(int32(2))%32))))
	if v489 != 0 {
		v499 = v489
		v500 = v482
		goto L100
	} else {
		goto L108
	}
L107:
	;
	goto L101
L108:
	;
	v491 = v482 + int32(1)
	if v491 != v467 {
		v482 = v491
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v519 = v513
	goto L113
L111:
	;
	goto L112
L112:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v425)+12))
	if v615 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L113:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v519 == v529 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	goto L112
L115:
	;
	if v456 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L116:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v531+v519<<(uint(int32(2))%32))))
	if v535 == int32(0) {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v535)+96))
	v539 = F_bms_add_member(m, v538, v455)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v535)+96)) = v539
	goto L115
L119:
	;
	if int32(0) < v598 {
		v519 = v598
		goto L113
	} else {
		goto L130
	}
L120:
	;
	v598 = base.I32_ctz(v584) | v585<<(uint(int32(5))%32)
	goto L119
L121:
	;
	v598 = int32(-2)
	goto L119
L122:
	;
	v549 = v519 + int32(1)
	v551 = base.I32_div_s(v549, int32(32))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v456)+4))
	if v552 <= v551 {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v555 = v456 + int32(8)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v555+v551<<(uint(int32(2))%32))))
	v562 = v559 & (int32(-1) << (uint(v549) % 32))
	if v562 != 0 {
		v584 = v562
		v585 = v551
		goto L120
	} else {
		goto L124
	}
L124:
	;
	v564 = v551 + int32(1)
	if v564 == v552 {
		goto L121
	} else {
		goto L125
	}
L125:
	;
	v567 = v564
	goto L126
L126:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v555+v567<<(uint(int32(2))%32))))
	if v574 != 0 {
		v584 = v574
		v585 = v567
		goto L120
	} else {
		goto L128
	}
L127:
	;
	goto L121
L128:
	;
	v576 = v567 + int32(1)
	if v576 != v552 {
		v567 = v576
		goto L126
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	goto L114
L131:
	;
	if int32(0) < v672 {
		goto L142
	} else {
		goto L143
	}
L132:
	;
	v672 = base.I32_ctz(v658) | v659<<(uint(int32(5))%32)
	goto L131
L133:
	;
	v672 = int32(-2)
	goto L131
L134:
	;
	v625 = base.I32_div_s(int32(0), int32(32))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v615)+4))
	if v626 <= v625 {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v629 = v615 + int32(8)
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v629+v625<<(uint(int32(2))%32))))
	v636 = v633 & int32(-1)
	if v636 != 0 {
		v658 = v636
		v659 = v625
		goto L132
	} else {
		goto L136
	}
L136:
	;
	v638 = v625 + int32(1)
	if v638 == v626 {
		goto L133
	} else {
		goto L137
	}
L137:
	;
	v641 = v638
	goto L138
L138:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v629+v641<<(uint(int32(2))%32))))
	if v648 != 0 {
		v658 = v648
		v659 = v641
		goto L132
	} else {
		goto L140
	}
L139:
	;
	goto L133
L140:
	;
	v650 = v641 + int32(1)
	if v650 != v626 {
		v641 = v650
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	v678 = v672
	goto L145
L143:
	;
	goto L144
L144:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v400)+16))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v425)+16))
	v775 = F_bms_union(m, v773, v774)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L163
	}
L145:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v678 == v688 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	goto L144
L147:
	;
	if v615 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L148:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v690+v678<<(uint(int32(2))%32))))
	if v694 == int32(0) {
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v694)+96))
	v698 = F_bms_add_member(m, v697, v614)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v694)+96)) = v698
	goto L147
L151:
	;
	if int32(0) < v757 {
		v678 = v757
		goto L145
	} else {
		goto L162
	}
L152:
	;
	v757 = base.I32_ctz(v743) | v744<<(uint(int32(5))%32)
	goto L151
L153:
	;
	v757 = int32(-2)
	goto L151
L154:
	;
	v708 = v678 + int32(1)
	v710 = base.I32_div_s(v708, int32(32))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v615)+4))
	if v711 <= v710 {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v714 = v615 + int32(8)
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v714+v710<<(uint(int32(2))%32))))
	v721 = v718 & (int32(-1) << (uint(v708) % 32))
	if v721 != 0 {
		v743 = v721
		v744 = v710
		goto L152
	} else {
		goto L156
	}
L156:
	;
	v723 = v710 + int32(1)
	if v723 == v711 {
		goto L153
	} else {
		goto L157
	}
L157:
	;
	v726 = v723
	goto L158
L158:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v714+v726<<(uint(int32(2))%32))))
	if v733 != 0 {
		v743 = v733
		v744 = v726
		goto L152
	} else {
		goto L160
	}
L159:
	;
	goto L153
L160:
	;
	v735 = v726 + int32(1)
	if v735 != v711 {
		v726 = v735
		goto L158
	} else {
		goto L161
	}
L161:
	;
	goto L159
L162:
	;
	goto L146
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v775
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v400)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v778
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v425)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v780
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v846 = v415
	v848 = v390
	v853 = v782
	goto L31
L164:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v787
	F_errmsg_internal(m, int32(_a_F_deconstruct_recurse_0), v14+int32(-48))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_deconstruct_recurse_1), int32(1405), int32(_a_F_deconstruct_recurse_2))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v803)+12))
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v803)+4))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v804+v805<<(uint(int32(2))%32)-int32(4))))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v813 = F_deconstruct_recurse(m, l0, v812, l2, v19, l4)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v811)+12))
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v816)+12))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v816)+4))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v817+v818<<(uint(int32(2))%32)-int32(4))))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v824)+12))
	v826 = F_bms_union(m, v815, v825)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v826
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v826
	v830 = v824
	v832 = v811
	v833 = v813
	v834 = v801
	goto L32
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v848
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v848
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v846
	v866 = F_list_make2_impl(m, v14+int32(-36), v14+int32(-40))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L1
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v877 = *(*int32)(unsafe.Add(mBase, _c_F_deconstruct_recurse[1]))
	if v848 != 0 {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v866
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v866
	v873 = F_list_make1_impl(m, int32(1), v14+int32(-44))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	v980 = v873
	goto L3
L175:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v848)+4))
	v880 = v878
	goto L177
L176:
	;
	v880 = int32(0)
	goto L177
L177:
	;
	if v846 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v846)+4))
	v883 = v881
	goto L180
L179:
	;
	v883 = int32(0)
	goto L180
L180:
	;
	if v880+v883 <= v877 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v886 = F_list_concat(m, v848, v846)
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L1
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	if v848 == int32(0) {
		v895 = int32(0)
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v980 = v886
	goto L3
L185:
	;
	if v846 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L186:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v848)+4))
	if v890 != int32(1) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v895 = v848
	goto L185
L188:
	;
	goto L189
L189:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v848)+12))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v893)))
	v895 = v894
	goto L185
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v904
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v895
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v895
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v904
	v913 = F_list_make2_impl(m, v14+int32(-28), v14+int32(-32))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L1
	} else {
		goto L195
	}
L191:
	;
	v904 = int32(0)
	goto L190
L192:
	;
	goto L193
L193:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v846)+4))
	if v899 != int32(1) {
		v904 = v846
		goto L190
	} else {
		goto L194
	}
L194:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v846)+12))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v902)))
	v904 = v903
	goto L190
L195:
	;
	v980 = v913
	goto L3
L196:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v919
	F_errmsg_internal(m, int32(_a_F_deconstruct_recurse_3), v16)
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_deconstruct_recurse_1), int32(1446), int32(_a_F_deconstruct_recurse_2))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v931
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l2
	v935 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v936 = F_bms_add_member(m, v935, v930)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v936
	v939 = F_bms_make_singleton(m, v930)
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v939
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = l1
	v949 = F_list_make1_impl(m, int32(1), v14+int32(-52))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	v980 = v949
	goto L3
L203:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v964)+4))
	if v967 <= int32(1) {
		v980 = v959
		goto L3
	} else {
		goto L204
	}
L204:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v970
	v980 = v959
	goto L3
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v986
	m.G0 = v16 - int32(-64)
	return v980
}
func F_deleteObjectsInList(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v344 int32
	_ = v344
	var v357 int32
	_ = v357
	var v374 int32
	_ = v374
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v450 int32
	_ = v450
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v591 int32
	_ = v591
	v4 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(288)
	m.G0 = v20
	v22 = int32(1)
	v24 = F_EventCacheLookup(m, int32(2))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = int32(0)
	if base.B2i32(v34 == v36)|l2&int32(1) == v36 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	return
L3:
	;
	if v24 != 0 {
		v34 = v22
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = F_EventCacheLookup(m, int32(3))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v27 != 0 {
		v34 = v22
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v30 = F_EventCacheLookup(m, int32(1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v34 = base.B2i32(v30 != int32(0))
	goto L1
L8:
	;
	m.G0 = v20 + int32(288)
	return
L9:
	;
	if v35 <= int32(0) {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	v105 = v35
	goto L11
L11:
	;
	if v105 <= int32(0) {
		goto L8
	} else {
		goto L28
	}
L12:
	;
	v54 = v4
	goto L13
L13:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62+v54<<(uint(int32(4))%32))))
	v67 = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v71 = v68 + v54*int32(12)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v72 <= int32(3465) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v105 = v98
	goto L11
L15:
	;
	if v87 != 0 {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	v87 = int32(1)
	goto L15
L17:
	;
	if base.Ui32(v72-int32(1260)) < base.Ui32(int32(3)) {
		v87 = v67
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if base.B2i32(v72 == int32(3466))|base.B2i32(v72 == int32(_a_F_deleteObjectsInList_0)) != 0 {
		v87 = v67
		goto L15
	} else {
		goto L22
	}
L20:
	;
	if v72 != int32(1213) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v87 = v67
	goto L15
L22:
	;
	goto L16
L23:
	;
	F_EventTriggerSQLDropAddObject(m, v71, v66&int32(1), base.B2i32(v66&int32(66) != int32(0)))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v97 = v54 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v97 < v98 {
		v54 = v97
		goto L13
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	goto L14
L28:
	;
	v136 = v105
	v142 = v4
	goto L29
L29:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l2&int32(8) != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L8
L31:
	;
	v591 = v142 + int32(1)
	if v591 < v578 {
		v136 = v578
		v142 = v591
		goto L29
	} else {
		goto L150
	}
L32:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149+v142<<(uint(int32(4))%32)))))
	if v153&int32(1) != 0 {
		v578 = v136
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v158 = v148 + v142*int32(12)
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_deleteObjectsInList[0]))
	if v160 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	F_RunObjectDropHook(m, v161, v162, v163, l2)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L2
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if l2&int32(2) != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L38
L40:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	F_ScanKeyInit(m, v20, int32(1), int32(3), int32(184), v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L48
	}
L41:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_relation_close(m, v166, int32(3))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L2
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	F_doDeletion(m, v158, l2)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L47
	}
L44:
	;
	F_doDeletion(m, v158, l2)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v174 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v174
	goto L40
L47:
	;
	goto L40
L48:
	;
	v185 = int32(2)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	F_ScanKeyInit(m, v20+int32(48), v185, int32(3), int32(184), v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	if v192 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v193 = int32(3)
	F_ScanKeyInit(m, v20+int32(96), v193, v193, int32(65), v192)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L2
	} else {
		goto L53
	}
L51:
	;
	v199 = v185
	goto L52
L52:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v204 = F_systable_beginscan(m, v200, int32(2673), int32(1), int32(0), v199, v20)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L2
	} else {
		goto L54
	}
L53:
	;
	v199 = v193
	goto L52
L54:
	;
	goto L55
L55:
	;
	v223 = F_systable_getnext(m, v204)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L2
	} else {
		goto L57
	}
L56:
	;
	F_systable_endscan(m, v204)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L2
	} else {
		goto L62
	}
L57:
	;
	if v223 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_simple_heap_delete(m, v225, v223+int32(4))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L2
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	goto L56
L61:
	;
	goto L55
L62:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	F_deleteSharedDependencyRecordsFor(m, v232, v233, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v240 = m.G0
	v242 = v240 - int32(144)
	m.G0 = v242
	F_ScanKeyInit(m, v242, int32(1), int32(3), int32(184), v237)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	v249 = int32(2)
	F_ScanKeyInit(m, v242+int32(48), v249, int32(3), int32(184), v238)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	if v239 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v259 = int32(3)
	F_ScanKeyInit(m, v242+int32(96), v259, v259, int32(65), v239)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L2
	} else {
		goto L69
	}
L67:
	;
	v265 = v249
	goto L68
L68:
	;
	v268 = F_table_open(m, int32(2609), int32(3))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L2
	} else {
		goto L70
	}
L69:
	;
	v265 = int32(3)
	goto L68
L70:
	;
	v273 = F_systable_beginscan(m, v268, int32(2675), int32(1), int32(0), v265, v242)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v275 = F_systable_getnext(m, v273)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	if v275 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v284 = v275
	goto L76
L74:
	;
	goto L75
L75:
	;
	F_systable_endscan(m, v273)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L2
	} else {
		goto L81
	}
L76:
	;
	F_simple_heap_delete(m, v268, v284+int32(4))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L2
	} else {
		goto L78
	}
L77:
	;
	goto L75
L78:
	;
	v298 = F_systable_getnext(m, v273)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	if v298 != 0 {
		v284 = v298
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	F_relation_close(m, v268, int32(3))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	v322 = int32(144)
	m.G0 = v242 + v322
	v325 = m.G0
	v327 = v325 - v322
	m.G0 = v327
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v332 = int32(1)
	if v329 <= int32(3591) {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v403 != 0 {
		goto L109
	} else {
		goto L110
	}
L84:
	;
	goto L83
L85:
	;
	v403 = int32(0)
	goto L84
L86:
	;
	if base.B2i32(base.Ui32(v329-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v329-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v403 = v332
		goto L84
	} else {
		goto L107
	}
L87:
	;
	if v329 <= int32(2670) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	if v329 <= int32(_a_F_deleteObjectsInList_1) {
		goto L97
	} else {
		goto L98
	}
L90:
	;
	switch v329 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v403 = v332
		goto L84
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L85
	default:
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v344 = v329 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v344))|base.B2i32(int32(1)<<(uint(v344)%32)&int32(226492515) == int32(0)) != 0 {
		goto L86
	} else {
		goto L95
	}
L93:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v329-int32(2396)) {
		goto L85
	} else {
		goto L94
	}
L94:
	;
	v403 = v332
	goto L84
L95:
	;
	v403 = v332
	goto L84
L96:
	;
	if base.Ui32(v329-int32(3592)) < base.Ui32(int32(2)) {
		v403 = v332
		goto L84
	} else {
		goto L105
	}
L97:
	;
	v357 = v329 - int32(_a_F_deleteObjectsInList_2)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v357))|base.B2i32(int32(1)<<(uint(v357)%32)&int32(963) == int32(0)) != 0 {
		goto L96
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	switch v329 - int32(_a_F_deleteObjectsInList_0) {
	case 0, 1, 2, 3, 4, 59, 60:
		v403 = v332
		goto L84
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L85
	default:
		goto L101
	}
L100:
	;
	v403 = v332
	goto L84
L101:
	;
	if base.Ui32(v329-int32(_a_F_deleteObjectsInList_3)) < base.Ui32(int32(3)) {
		v403 = v332
		goto L84
	} else {
		goto L102
	}
L102:
	;
	v374 = v329 - int32(_a_F_deleteObjectsInList_4)
	if base.Ui32(int32(15)) < base.Ui32(v374) {
		goto L85
	} else {
		goto L103
	}
L103:
	;
	if int32(1)<<(uint(v374)%32)&int32(_a_F_deleteObjectsInList_5) != 0 {
		v403 = v332
		goto L84
	} else {
		goto L104
	}
L104:
	;
	goto L85
L105:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v329-int32(4060)) {
		goto L85
	} else {
		goto L106
	}
L106:
	;
	v403 = v332
	goto L84
L107:
	;
	goto L85
L108:
	;
	m.G0 = v327 + int32(144)
	v510 = F_table_open(m, int32(3394), int32(3))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L2
	} else {
		goto L132
	}
L109:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	F_DeleteSharedSecurityLabel(m, v404, v405)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L2
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	F_ScanKeyInit(m, v327, int32(1), int32(3), int32(184), v404)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L2
	} else {
		goto L113
	}
L112:
	;
	goto L108
L113:
	;
	v413 = int32(2)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	F_ScanKeyInit(m, v327+int32(48), v413, int32(3), int32(184), v419)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	if v422 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v425 = int32(3)
	F_ScanKeyInit(m, v327+int32(96), v425, v425, int32(65), v422)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L2
	} else {
		goto L118
	}
L116:
	;
	v431 = v413
	goto L117
L117:
	;
	v434 = F_table_open(m, int32(3596), int32(3))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L2
	} else {
		goto L119
	}
L118:
	;
	v431 = int32(3)
	goto L117
L119:
	;
	v439 = F_systable_beginscan(m, v434, int32(3597), int32(1), int32(0), v431, v327)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L2
	} else {
		goto L120
	}
L120:
	;
	v441 = F_systable_getnext(m, v439)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	if v441 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v450 = v441
	goto L125
L123:
	;
	goto L124
L124:
	;
	F_systable_endscan(m, v439)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L2
	} else {
		goto L130
	}
L125:
	;
	F_simple_heap_delete(m, v434, v450+int32(4))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L2
	} else {
		goto L127
	}
L126:
	;
	goto L124
L127:
	;
	v464 = F_systable_getnext(m, v439)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L2
	} else {
		goto L128
	}
L128:
	;
	if v464 != 0 {
		v450 = v464
		goto L125
	} else {
		goto L129
	}
L129:
	;
	goto L126
L130:
	;
	F_relation_close(m, v434, int32(3))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L2
	} else {
		goto L131
	}
L131:
	;
	goto L108
L132:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	F_ScanKeyInit(m, v20+int32(144), int32(1), int32(3), int32(184), v517)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L2
	} else {
		goto L133
	}
L133:
	;
	v520 = int32(2)
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	F_ScanKeyInit(m, v20+int32(192), v520, int32(3), int32(184), v524)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L2
	} else {
		goto L134
	}
L134:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	if v527 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v528 = int32(3)
	F_ScanKeyInit(m, v20+int32(240), v528, v528, int32(65), v527)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L2
	} else {
		goto L138
	}
L136:
	;
	v534 = v520
	goto L137
L137:
	;
	v540 = F_systable_beginscan(m, v510, int32(3395), int32(1), int32(0), v534, v20+int32(144))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L2
	} else {
		goto L139
	}
L138:
	;
	v534 = v528
	goto L137
L139:
	;
	goto L140
L140:
	;
	v559 = F_systable_getnext(m, v540)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L2
	} else {
		goto L142
	}
L141:
	;
	F_systable_endscan(m, v540)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L2
	} else {
		goto L147
	}
L142:
	;
	if v559 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	F_simple_heap_delete(m, v510, v559+int32(4))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L2
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	goto L141
L146:
	;
	goto L140
L147:
	;
	F_relation_close(m, v510, int32(3))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L2
	} else {
		goto L148
	}
L148:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L2
	} else {
		goto L149
	}
L149:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v578 = v572
	goto L31
L150:
	;
	goto L30
}
func F_derfc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v61 float64
	_ = v61
	var v80 float64
	_ = v80
	var v85 float64
	_ = v85
	var v90 float64
	_ = v90
	var v93 float64
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v8 = base.I64_reinterpret_f64(v7)
	v11 = base.I32_wrap_i64(int64(base.Ui64(v8) >> (uint(int64(32)) % 64)))
	v13 = v11 & int32(2147483647)
	if base.Ui32(int32(2146435072)) <= base.Ui32(v13) {
		v93 = base.F64_add(base.F64_div(float64(1), v7), base.F64_convert_i32_u(int32(base.Ui32(v11)>>(uint(int32(30))%32))&int32(2)))
	} else {
		if base.Ui32(v13) <= base.Ui32(int32(1072365567)) {
			if base.Ui32(v13) <= base.Ui32(int32(1013972991)) {
				v93 = base.F64_sub(float64(1), v7)
			} else {
				v30 = float64(1)
				v31 = base.F64_mul(v7, v7)
				v61 = base.F64_div(base.F64_add(base.F64_mul(v31, base.F64_add(base.F64_mul(v31, base.F64_add(base.F64_mul(v31, base.F64_add(base.F64_mul(v31, float64(-2.3763016656650163e-05)), float64(-0.005770270296489442))), float64(-0.02848174957559851))), float64(-0.3250421072470015))), float64(0.12837916709551256)), base.F64_add(base.F64_mul(v31, base.F64_add(base.F64_mul(v31, base.F64_add(base.F64_mul(v31, base.F64_add(base.F64_mul(v31, base.F64_add(base.F64_mul(v31, float64(-3.960228278775368e-06)), float64(0.00013249473800432164))), float64(0.005081306281875766))), float64(0.0650222499887673))), float64(0.39791722395915535))), v30))
				if base.B2i32(base.Ui32(int32(1070596095)) < base.Ui32(v13))&base.B2i32(int64(0) <= v8) == int32(0) {
					v93 = base.F64_sub(v30, base.F64_add(base.F64_mul(v7, v61), v7))
				} else {
					v93 = base.F64_sub(float64(0.5), base.F64_add(base.F64_mul(v7, v61), base.F64_add(v7, float64(-0.5))))
				}
			}
		} else {
			if base.Ui32(v13) <= base.Ui32(int32(1077673983)) {
				v80 = F_erfc2(m, v13, v7)
				mBase = m.M
				if int64(0) <= v8 {
					v85 = v80
				} else {
					v85 = base.F64_sub(float64(2), v80)
				}
				v93 = v85
			} else {
				if int64(0) <= v8 {
					v90 = float64(0)
				} else {
					v90 = float64(2)
				}
				v93 = v90
			}
		}
	}
	if base.F64_eq(base.F64_abs(v93), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		F_float_overflow_error(m)
		mBase = m.M
		v100 = m.ExcPending
		if v100 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v101 = F_Float8GetDatum(m, v93)
		mBase = m.M
		v102 = m.ExcPending
		if v102 != 0 {
			return int32(0)
		} else {
			return v101
		}
	}
}
func F_detoast_external_attr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	v4 = l0
	goto L3
L1:
	;
	return v78
L2:
	;
	if v11&int32(254) != int32(2) {
		goto L31
	} else {
		goto L32
	}
L3:
	;
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v7 != int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v22&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	return v4
L6:
	;
	goto L7
L7:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+1)))
	if v11 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if v11 != int32(18) {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v4)+2))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v22 == int32(1) {
		v4 = v21
		goto L3
	} else {
		goto L14
	}
L11:
	;
	v16 = F_toast_fetch_datum(m, v4)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	return v16
L14:
	;
	goto L4
L15:
	;
	v32 = int32(base.Ui32(v22) >> (uint(int32(1)) % 32))
	goto L17
L16:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v32 = int32(base.Ui32(v29) >> (uint(int32(2)) % 32))
	goto L17
L17:
	;
	v33 = F_palloc(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v35 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v60 == int32(0) {
		v78 = v33
		goto L1
	} else {
		goto L30
	}
L20:
	;
	v39 = int32(18)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v41 == v39 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v52 = int32(1)
	if v35&v52 != 0 {
		v60 = int32(base.Ui32(v35) >> (uint(v52) % 32))
		goto L19
	} else {
		goto L29
	}
L23:
	;
	v44 = v39
	goto L25
L24:
	;
	v44 = int32(2)
	goto L25
L25:
	;
	if base.Ui32((v41-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v51 = int32(6)
	goto L28
L27:
	;
	v51 = v44
	goto L28
L28:
	;
	v60 = v51
	goto L19
L29:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v60 = int32(base.Ui32(v56) >> (uint(int32(2)) % 32))
	goto L19
L30:
	;
	base.MemoryCopy(m, v33, v21, v60)
	return v33
L31:
	;
	return v4
L32:
	;
	goto L33
L33:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v4)+2))
	v71 = F_EOH_get_flat_size(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	v73 = F_palloc(m, v71)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	F_EOH_flatten_into(m, v70, v73, v71)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	v78 = v73
	goto L1
}
func F_dfloor(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)))
	v5 = F_Float8GetDatum(m, base.F64_floor(v3))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_digest_block_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+8))
	return v3
}
func F_disassembleLeaf(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	v8 = F_palloc0(m, int32(32))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v15 = l0 + v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+6)))
	if v16&int32(128) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = int32(32)
	v20 = l0 + v19
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v24 = v20 + v21 - v19
	if base.Ui32(v24) <= base.Ui32(v20) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
	if v65 != 0 {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v26 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)) = uint8(v26)
	return v8
L7:
	;
	goto L8
L8:
	;
	v31 = v20
	goto L9
L9:
	;
	v36 = F_palloc(m, int32(32))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v62 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)) = uint8(v62)
	return v8
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v31
	v41 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+8)) = uint8(v41)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v43 == v41 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8
	goto L14
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v8
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v36
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+6)))
	v60 = v31 + (v53+int32(1))&int32(_a_F_disassembleLeaf_0) + int32(8)
	if base.Ui32(v60) < base.Ui32(v24) {
		v31 = v60
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	v67 = F_palloc(m, int32(32))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v96 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)) = uint8(v96)
	return v8
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+20)) = int32(0)
	v71 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+8)) = uint8(v71)
	v74 = v65 * int32(6)
	v75 = F_palloc(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+24)) = v75
	if v74 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	base.MemoryCopy(m, v75, l0+int32(32), v74)
	goto L23
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+28)) = v65
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v82 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8
	goto L26
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v8
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v67
	goto L18
}
func F_do_des(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var __phi150 int32
	_ = __phi150
	var v151 int32
	_ = v151
	var __phi151 int32
	_ = __phi151
	var v154 int32
	_ = v154
	var __phi154 int32
	_ = __phi154
	var v155 int32
	_ = v155
	var __phi155 int32
	_ = __phi155
	var v158 int32
	_ = v158
	var __phi158 int32
	_ = __phi158
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	if l4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1)
L2:
	;
	goto L3
L3:
	;
	v22 = base.B2i32(int32(0) < l4)
	if int32(0) < l4 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = int32(_a_F_do_des_0)
	goto L6
L5:
	;
	v23 = int32(_a_F_do_des_1)
	goto L6
L6:
	;
	if int32(0) < l4 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v26 = int32(_a_F_do_des_2)
	goto L9
L8:
	;
	v26 = int32(_a_F_do_des_3)
	goto L9
L9:
	;
	v27 = int32(6)
	v29 = int32(1020)
	v30 = int32(base.Ui32(l1)>>(uint(v27)%32)) & v29
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_do_des[0])))
	v34 = int32(14)
	v37 = int32(base.Ui32(l1)>>(uint(v34)%32)) & v29
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_do_des[1])))
	v41 = int32(22)
	v44 = int32(base.Ui32(l1)>>(uint(v41)%32)) & v29
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+uint32(_c_F_do_des[2])))
	v51 = int32(base.Ui32(l0)>>(uint(v27)%32)) & v29
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+uint32(_c_F_do_des[3])))
	v58 = int32(base.Ui32(l0)>>(uint(v34)%32)) & v29
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+uint32(_c_F_do_des[4])))
	v65 = int32(base.Ui32(l0)>>(uint(v41)%32)) & v29
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_c_F_do_des[5])))
	v71 = int32(2)
	v74 = l0 << (uint(v71) % 32) & v29
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_do_des[6])))
	v85 = l1 << (uint(v71) % 32) & v29
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+uint32(_c_F_do_des[7])))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v85)+uint32(_c_F_do_des[8])))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_do_des[9])))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_do_des[10])))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v44)+uint32(_c_F_do_des[11])))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_do_des[12])))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v51)+uint32(_c_F_do_des[13])))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v58)+uint32(_c_F_do_des[14])))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_c_F_do_des[15])))
	v122 = l4 >> (uint(int32(31)) % 32)
	v130 = v33 | (v40 | (v47 | (v54 | (v61 | v68) | v77))) | v88
	v131 = v92 | (v95 | (v98 | (v101 | (v104 | (v107 | (v110 | v113))))))
	v132 = l4 ^ v122 - v122
	goto L10
L10:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_do_des[16]))
	if v140 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v263 = int32(6)
	v265 = int32(1020)
	v266 = int32(base.Ui32(v154)>>(uint(v263)%32)) & v265
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v266)+uint32(_c_F_do_des[17])))
	v270 = int32(14)
	v273 = int32(base.Ui32(v154)>>(uint(v270)%32)) & v265
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v273)+uint32(_c_F_do_des[18])))
	v277 = int32(22)
	v280 = int32(base.Ui32(v154)>>(uint(v277)%32)) & v265
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v280)+uint32(_c_F_do_des[19])))
	v287 = int32(base.Ui32(v256)>>(uint(v263)%32)) & v265
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v287)+uint32(_c_F_do_des[20])))
	v294 = int32(base.Ui32(v256)>>(uint(v270)%32)) & v265
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v294)+uint32(_c_F_do_des[21])))
	v301 = int32(base.Ui32(v256)>>(uint(v277)%32)) & v265
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v301)+uint32(_c_F_do_des[22])))
	v307 = int32(2)
	v310 = v256 << (uint(v307) % 32) & v265
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v310)+uint32(_c_F_do_des[23])))
	v321 = v154 << (uint(v307) % 32) & v265
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v321)+uint32(_c_F_do_des[24])))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v269 | (v276 | (v283 | (v290 | (v297 | v304) | v313))) | v324
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v321)+uint32(_c_F_do_des[25])))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v266)+uint32(_c_F_do_des[26])))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v273)+uint32(_c_F_do_des[27])))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v280)+uint32(_c_F_do_des[28])))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v310)+uint32(_c_F_do_des[29])))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v287)+uint32(_c_F_do_des[30])))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v294)+uint32(_c_F_do_des[31])))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v301)+uint32(_c_F_do_des[32])))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v329 | (v332 | (v335 | (v338 | (v341 | (v344 | (v347 | v350))))))
	return int32(0)
L12:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v146 = v132 - int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_do_des[33]))
	__phi150 = v26
	__phi151 = v23
	__phi154 = v130
	__phi155 = v131
	__phi158 = int32(16)
	v150 = __phi150
	v151 = __phi151
	v154 = __phi154
	v155 = __phi155
	v158 = __phi158
	goto L17
L15:
	;
	return int32(0)
L16:
	;
	goto L14
L17:
	;
	v166 = int32(16515072)
	v170 = int32(_a_F_do_des_4)
	v175 = int32(4032)
	v182 = base.I32_rotl(v154, int32(23))&v166 | int32(base.Ui32(v154)>>(uint(int32(11))%32))&v170 | int32(base.Ui32(v154)>>(uint(int32(13))%32))&v175 | int32(base.Ui32(v154)>>(uint(int32(15))%32))&int32(63)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v185 = int32(1)
	v206 = v154<<(uint(v185)%32)&int32(62) | (v154<<(uint(int32(3))%32)&v175 | (v154<<(uint(int32(7))%32)&v166 | (v154<<(uint(int32(5))%32)&v170 | int32(base.Ui32(v154)>>(uint(int32(31))%32)))))
	v208 = v148 & (v182 ^ v206)
	v209 = v182 ^ v183 ^ v208
	v210 = int32(4095)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209&v210)+uint32(_c_F_do_des[34]))))
	v215 = int32(2)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v214<<(uint(v215)%32))+uint32(_c_F_do_des[35])))
	v220 = int32(12)
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v209)>>(uint(v220)%32)))+uint32(_c_F_do_des[36]))))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v224<<(uint(v215)%32))+uint32(_c_F_do_des[37])))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v233 = v206 ^ v231 ^ v208
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v233)>>(uint(v220)%32)))+uint32(_c_F_do_des[38]))))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v238<<(uint(v215)%32))+uint32(_c_F_do_des[39])))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233&v210)+uint32(_c_F_do_des[40]))))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v249<<(uint(v215)%32))+uint32(_c_F_do_des[41])))
	v256 = v155 ^ (v219 | v229 | v243 | v254)
	v257 = int32(4)
	v262 = v158 - v185
	if v262 != 0 {
		__phi150 = v150 + v257
		__phi151 = v151 + v257
		__phi154 = v256
		__phi155 = v154
		__phi158 = v262
		v150 = __phi150
		v151 = __phi151
		v154 = __phi154
		v155 = __phi155
		v158 = __phi158
		goto L17
	} else {
		goto L19
	}
L18:
	;
	if v146 != 0 {
		v130 = v154
		v131 = v256
		v132 = v146
		goto L10
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	goto L11
}
func F_do_truncate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	goto L1
L1:
	;
	v14 = m.Env.X__syscall_truncate64(m, l0, int64(0))
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v14) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	if int32(0) <= v22 {
		goto L11
	} else {
		goto L12
	}
L3:
	;
	if v22 == int32(-1) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_truncate[0])) = int32(0) - v14
	v22 = int32(-1)
	goto L6
L5:
	;
	v22 = v14
	goto L6
L6:
	;
	goto L3
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_do_truncate[0]))
	if v26 == int32(27) {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	goto L2
L10:
	;
	goto L9
L11:
	;
	m.G0 = v7 + int32(16)
	return v22
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_do_truncate[0]))
	if v32 == int32(44) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v37 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	if v37 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_truncate[0])) = v32
	goto L11
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg(m, int32(_a_F_do_truncate_0), v7)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_do_truncate_1), int32(356), int32(_a_F_do_truncate_2))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L14
	} else {
		goto L21
	}
L21:
	;
	goto L18
}
func F_dobyteatrim(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v211 int32
	_ = v211
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v211
L2:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v43 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L3:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v17 == int32(18) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v28 = int32(1)
	if v11&v28 != 0 {
		v40 = int32(base.Ui32(v11)>>(uint(v28)%32)) - v28
		goto L2
	} else {
		goto L12
	}
L6:
	;
	v20 = int32(16)
	goto L8
L7:
	;
	v20 = int32(0)
	goto L8
L8:
	;
	if base.Ui32((v17-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v27 = int32(4)
	goto L11
L10:
	;
	v27 = v20
	goto L11
L11:
	;
	v40 = v27
	goto L2
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = int32(base.Ui32(v34)>>(uint(int32(2))%32)) - int32(4)
	goto L2
L13:
	;
	if base.B2i32(v40 <= int32(0))|base.B2i32(v72 <= int32(0)) != 0 {
		v211 = l0
		goto L1
	} else {
		goto L24
	}
L14:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v49 == int32(18) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v60 = int32(1)
	if v43&v60 != 0 {
		v72 = int32(base.Ui32(v43)>>(uint(v60)%32)) - v60
		goto L13
	} else {
		goto L23
	}
L17:
	;
	v52 = int32(16)
	goto L19
L18:
	;
	v52 = int32(0)
	goto L19
L19:
	;
	if base.Ui32((v49-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v59 = int32(4)
	goto L22
L21:
	;
	v59 = v52
	goto L22
L22:
	;
	v72 = v59
	goto L13
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v72 = int32(base.Ui32(v66)>>(uint(int32(2))%32)) - int32(4)
	goto L13
L24:
	;
	v76 = int32(1)
	if v43&v76 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v80 = v76
	goto L27
L26:
	;
	v80 = int32(4)
	goto L27
L27:
	;
	v81 = l1 + v80
	v83 = int32(1)
	v84 = v81 + v72 - v83
	if v11&v83 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v89 = v83
	goto L30
L29:
	;
	v89 = int32(4)
	goto L30
L30:
	;
	v90 = l0 + v89
	if l2 == int32(0) {
		v128 = v40
		v131 = v90
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v198 = v188 + int32(4)
	v199 = F_palloc(m, v198)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L54
	} else {
		goto L55
	}
L32:
	;
	v188 = int32(0)
	v191 = v180
	goto L31
L33:
	;
	v137 = int32(0)
	if base.B2i32(l3 == v137)|base.B2i32(v128 <= v137) != 0 {
		v188 = v128
		v191 = v131
		goto L31
	} else {
		goto L44
	}
L34:
	;
	v96 = v40
	v99 = v90
	goto L35
L35:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	v106 = v81
	goto L37
L36:
	;
	v180 = l0 + v89 + v40
	goto L32
L37:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v116 != v105 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v121 = int32(1)
	if v121 < v96 {
		v96 = v96 - v121
		v99 = v99 + v121
		goto L35
	} else {
		goto L43
	}
L39:
	;
	v119 = v106 + int32(1)
	if base.Ui32(v119) <= base.Ui32(v84) {
		v106 = v119
		goto L37
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	goto L38
L42:
	;
	v128 = v96
	v131 = v99
	goto L33
L43:
	;
	goto L36
L44:
	;
	v145 = v40 + v90
	v146 = v128
	goto L45
L45:
	;
	if base.Ui32(v84) < base.Ui32(v81) {
		v188 = v128
		v191 = v131
		goto L31
	} else {
		goto L47
	}
L46:
	;
	v180 = v131
	goto L32
L47:
	;
	v155 = v145 - int32(1)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	v157 = v81
	goto L49
L48:
	;
	v172 = int32(1)
	if v172 < v146 {
		v145 = v155
		v146 = v146 - v172
		goto L45
	} else {
		goto L53
	}
L49:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v156 == v167 {
		goto L48
	} else {
		goto L51
	}
L50:
	;
	v188 = v146
	v191 = v131
	goto L31
L51:
	;
	v170 = v157 + int32(1)
	if base.Ui32(v170) <= base.Ui32(v84) {
		v157 = v170
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	goto L46
L54:
	;
	return int32(0)
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199))) = v198 << (uint(int32(2)) % 32)
	if v188 == int32(0) {
		v211 = v199
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.MemoryCopy(m, v199+int32(4), v191, v188)
	v211 = v199
	goto L1
}
func F_double_to_shortest_decimal_buf(m *base.Module, l0 float64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int64
	_ = v53
	var v56 int64
	_ = v56
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v117 int64
	_ = v117
	var v124 int64
	_ = v124
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v147 int64
	_ = v147
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v158 int64
	_ = v158
	var v165 int64
	_ = v165
	var v177 int32
	_ = v177
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v184 int64
	_ = v184
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v207 int64
	_ = v207
	var v211 int64
	_ = v211
	var v212 int64
	_ = v212
	var v217 int32
	_ = v217
	var v218 int64
	_ = v218
	var v222 int64
	_ = v222
	var v228 int64
	_ = v228
	var v229 int64
	_ = v229
	var v231 int64
	_ = v231
	var v234 int64
	_ = v234
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v242 int64
	_ = v242
	var v249 int64
	_ = v249
	var v261 int32
	_ = v261
	var v262 int64
	_ = v262
	var v268 int64
	_ = v268
	var v269 int64
	_ = v269
	var v271 int64
	_ = v271
	var v274 int64
	_ = v274
	var v275 int64
	_ = v275
	var v277 int64
	_ = v277
	var v278 int64
	_ = v278
	var v282 int64
	_ = v282
	var v289 int64
	_ = v289
	var v301 int32
	_ = v301
	var v302 int64
	_ = v302
	var v303 int64
	_ = v303
	var v304 int64
	_ = v304
	var v305 int64
	_ = v305
	var v308 int64
	_ = v308
	var v323 int64
	_ = v323
	var v327 int64
	_ = v327
	var v328 int64
	_ = v328
	var v333 int32
	_ = v333
	var v334 int64
	_ = v334
	var v340 int64
	_ = v340
	var v341 int64
	_ = v341
	var v343 int64
	_ = v343
	var v346 int64
	_ = v346
	var v347 int64
	_ = v347
	var v349 int64
	_ = v349
	var v350 int64
	_ = v350
	var v354 int64
	_ = v354
	var v361 int64
	_ = v361
	var v373 int32
	_ = v373
	var v374 int64
	_ = v374
	var v380 int64
	_ = v380
	var v381 int64
	_ = v381
	var v383 int64
	_ = v383
	var v386 int64
	_ = v386
	var v387 int64
	_ = v387
	var v389 int64
	_ = v389
	var v390 int64
	_ = v390
	var v394 int64
	_ = v394
	var v401 int64
	_ = v401
	var v413 int32
	_ = v413
	var v414 int64
	_ = v414
	var v415 int64
	_ = v415
	var v416 int64
	_ = v416
	var v417 int64
	_ = v417
	var v420 int64
	_ = v420
	var v435 int64
	_ = v435
	var v439 int64
	_ = v439
	var v440 int64
	_ = v440
	var v444 int64
	_ = v444
	var v445 int64
	_ = v445
	var v446 int64
	_ = v446
	var v450 int32
	_ = v450
	var v451 int64
	_ = v451
	var v458 int64
	_ = v458
	var v459 int64
	_ = v459
	var v461 int64
	_ = v461
	var v464 int64
	_ = v464
	var v465 int64
	_ = v465
	var v467 int64
	_ = v467
	var v468 int64
	_ = v468
	var v472 int64
	_ = v472
	var v479 int64
	_ = v479
	var v490 int32
	_ = v490
	var v491 int64
	_ = v491
	var v503 int32
	_ = v503
	var v515 int64
	_ = v515
	var v523 int32
	_ = v523
	var v524 int64
	_ = v524
	var v531 int64
	_ = v531
	var v532 int64
	_ = v532
	var v534 int64
	_ = v534
	var v537 int64
	_ = v537
	var v538 int64
	_ = v538
	var v540 int64
	_ = v540
	var v541 int64
	_ = v541
	var v545 int64
	_ = v545
	var v552 int64
	_ = v552
	var v565 int64
	_ = v565
	var v567 int64
	_ = v567
	var v578 int32
	_ = v578
	var v591 int64
	_ = v591
	var v598 int32
	_ = v598
	var v599 int64
	_ = v599
	var v606 int64
	_ = v606
	var v607 int64
	_ = v607
	var v609 int64
	_ = v609
	var v612 int64
	_ = v612
	var v613 int64
	_ = v613
	var v615 int64
	_ = v615
	var v616 int64
	_ = v616
	var v620 int64
	_ = v620
	var v627 int64
	_ = v627
	var v640 int64
	_ = v640
	var v642 int64
	_ = v642
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v668 int64
	_ = v668
	var v669 int64
	_ = v669
	var v671 int64
	_ = v671
	var v677 int64
	_ = v677
	var v678 int64
	_ = v678
	var v680 int64
	_ = v680
	var v683 int64
	_ = v683
	var v684 int64
	_ = v684
	var v686 int64
	_ = v686
	var v687 int64
	_ = v687
	var v691 int64
	_ = v691
	var v698 int64
	_ = v698
	var v710 int32
	_ = v710
	var v713 int64
	_ = v713
	var v714 int64
	_ = v714
	var v720 int64
	_ = v720
	var v721 int64
	_ = v721
	var v723 int64
	_ = v723
	var v726 int64
	_ = v726
	var v727 int64
	_ = v727
	var v729 int64
	_ = v729
	var v730 int64
	_ = v730
	var v734 int64
	_ = v734
	var v741 int64
	_ = v741
	var v753 int32
	_ = v753
	var v754 int64
	_ = v754
	var v755 int64
	_ = v755
	var v756 int64
	_ = v756
	var v757 int64
	_ = v757
	var v760 int64
	_ = v760
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v782 int64
	_ = v782
	var v786 int64
	_ = v786
	var v787 int64
	_ = v787
	var v792 int32
	_ = v792
	var v793 int64
	_ = v793
	var v797 int64
	_ = v797
	var v803 int64
	_ = v803
	var v804 int64
	_ = v804
	var v806 int64
	_ = v806
	var v809 int64
	_ = v809
	var v810 int64
	_ = v810
	var v812 int64
	_ = v812
	var v813 int64
	_ = v813
	var v817 int64
	_ = v817
	var v824 int64
	_ = v824
	var v836 int32
	_ = v836
	var v837 int64
	_ = v837
	var v843 int64
	_ = v843
	var v844 int64
	_ = v844
	var v846 int64
	_ = v846
	var v849 int64
	_ = v849
	var v850 int64
	_ = v850
	var v852 int64
	_ = v852
	var v853 int64
	_ = v853
	var v857 int64
	_ = v857
	var v864 int64
	_ = v864
	var v876 int32
	_ = v876
	var v877 int64
	_ = v877
	var v878 int64
	_ = v878
	var v879 int64
	_ = v879
	var v880 int64
	_ = v880
	var v883 int64
	_ = v883
	var v898 int64
	_ = v898
	var v902 int64
	_ = v902
	var v903 int64
	_ = v903
	var v908 int32
	_ = v908
	var v909 int64
	_ = v909
	var v915 int64
	_ = v915
	var v916 int64
	_ = v916
	var v918 int64
	_ = v918
	var v921 int64
	_ = v921
	var v922 int64
	_ = v922
	var v924 int64
	_ = v924
	var v925 int64
	_ = v925
	var v929 int64
	_ = v929
	var v936 int64
	_ = v936
	var v948 int32
	_ = v948
	var v949 int64
	_ = v949
	var v955 int64
	_ = v955
	var v956 int64
	_ = v956
	var v958 int64
	_ = v958
	var v961 int64
	_ = v961
	var v962 int64
	_ = v962
	var v964 int64
	_ = v964
	var v965 int64
	_ = v965
	var v969 int64
	_ = v969
	var v976 int64
	_ = v976
	var v988 int32
	_ = v988
	var v989 int64
	_ = v989
	var v990 int64
	_ = v990
	var v991 int64
	_ = v991
	var v992 int64
	_ = v992
	var v995 int64
	_ = v995
	var v1010 int64
	_ = v1010
	var v1014 int64
	_ = v1014
	var v1015 int64
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1020 int64
	_ = v1020
	var v1021 int64
	_ = v1021
	var v1022 int64
	_ = v1022
	var v1029 int64
	_ = v1029
	var v1045 int32
	_ = v1045
	var v1057 int64
	_ = v1057
	var v1058 int64
	_ = v1058
	var v1060 int64
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1064 int64
	_ = v1064
	var v1071 int64
	_ = v1071
	var v1072 int64
	_ = v1072
	var v1074 int64
	_ = v1074
	var v1077 int64
	_ = v1077
	var v1078 int64
	_ = v1078
	var v1080 int64
	_ = v1080
	var v1081 int64
	_ = v1081
	var v1085 int64
	_ = v1085
	var v1092 int64
	_ = v1092
	var v1104 int32
	_ = v1104
	var v1105 int64
	_ = v1105
	var v1112 int64
	_ = v1112
	var v1113 int64
	_ = v1113
	var v1115 int64
	_ = v1115
	var v1118 int64
	_ = v1118
	var v1119 int64
	_ = v1119
	var v1121 int64
	_ = v1121
	var v1122 int64
	_ = v1122
	var v1126 int64
	_ = v1126
	var v1133 int64
	_ = v1133
	var v1144 int32
	_ = v1144
	var v1145 int64
	_ = v1145
	var v1146 int64
	_ = v1146
	var v1147 int64
	_ = v1147
	var v1148 int64
	_ = v1148
	var v1150 int64
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1171 int64
	_ = v1171
	var v1172 int64
	_ = v1172
	var v1173 int64
	_ = v1173
	var v1179 int32
	_ = v1179
	var v1180 int64
	_ = v1180
	var v1187 int64
	_ = v1187
	var v1188 int64
	_ = v1188
	var v1190 int64
	_ = v1190
	var v1193 int64
	_ = v1193
	var v1194 int64
	_ = v1194
	var v1196 int64
	_ = v1196
	var v1197 int64
	_ = v1197
	var v1201 int64
	_ = v1201
	var v1208 int64
	_ = v1208
	var v1220 int32
	_ = v1220
	var v1221 int64
	_ = v1221
	var v1228 int64
	_ = v1228
	var v1229 int64
	_ = v1229
	var v1231 int64
	_ = v1231
	var v1234 int64
	_ = v1234
	var v1235 int64
	_ = v1235
	var v1237 int64
	_ = v1237
	var v1238 int64
	_ = v1238
	var v1242 int64
	_ = v1242
	var v1249 int64
	_ = v1249
	var v1261 int32
	_ = v1261
	var v1262 int64
	_ = v1262
	var v1269 int64
	_ = v1269
	var v1270 int64
	_ = v1270
	var v1272 int64
	_ = v1272
	var v1275 int64
	_ = v1275
	var v1276 int64
	_ = v1276
	var v1278 int64
	_ = v1278
	var v1279 int64
	_ = v1279
	var v1283 int64
	_ = v1283
	var v1290 int64
	_ = v1290
	var v1302 int32
	_ = v1302
	var v1307 int32
	_ = v1307
	var v1308 int64
	_ = v1308
	var v1309 int64
	_ = v1309
	var v1310 int64
	_ = v1310
	var v1314 int32
	_ = v1314
	var v1315 int64
	_ = v1315
	var v1317 int64
	_ = v1317
	var v1318 int64
	_ = v1318
	var v1320 int64
	_ = v1320
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1341 int64
	_ = v1341
	var v1342 int64
	_ = v1342
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1364 int32
	_ = v1364
	var v1376 int64
	_ = v1376
	var v1377 int64
	_ = v1377
	var v1379 int64
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1384 int64
	_ = v1384
	var v1385 int64
	_ = v1385
	var v1392 int64
	_ = v1392
	var v1393 int64
	_ = v1393
	var v1395 int64
	_ = v1395
	var v1398 int64
	_ = v1398
	var v1399 int64
	_ = v1399
	var v1401 int64
	_ = v1401
	var v1402 int64
	_ = v1402
	var v1406 int64
	_ = v1406
	var v1413 int64
	_ = v1413
	var v1425 int32
	_ = v1425
	var v1427 int64
	_ = v1427
	var v1428 int64
	_ = v1428
	var v1435 int64
	_ = v1435
	var v1436 int64
	_ = v1436
	var v1438 int64
	_ = v1438
	var v1441 int64
	_ = v1441
	var v1442 int64
	_ = v1442
	var v1444 int64
	_ = v1444
	var v1445 int64
	_ = v1445
	var v1449 int64
	_ = v1449
	var v1456 int64
	_ = v1456
	var v1468 int64
	_ = v1468
	var v1469 int64
	_ = v1469
	var v1470 int64
	_ = v1470
	var v1471 int64
	_ = v1471
	var v1473 int64
	_ = v1473
	var v1477 int32
	_ = v1477
	var v1479 int64
	_ = v1479
	var v1480 int64
	_ = v1480
	var v1487 int64
	_ = v1487
	var v1488 int64
	_ = v1488
	var v1490 int64
	_ = v1490
	var v1493 int64
	_ = v1493
	var v1494 int64
	_ = v1494
	var v1496 int64
	_ = v1496
	var v1497 int64
	_ = v1497
	var v1501 int64
	_ = v1501
	var v1508 int64
	_ = v1508
	var v1519 int64
	_ = v1519
	var v1521 int64
	_ = v1521
	var v1529 int32
	_ = v1529
	var v1530 int64
	_ = v1530
	var v1531 int64
	_ = v1531
	var v1532 int64
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1536 int64
	_ = v1536
	var v1543 int64
	_ = v1543
	var v1544 int64
	_ = v1544
	var v1546 int64
	_ = v1546
	var v1549 int64
	_ = v1549
	var v1550 int64
	_ = v1550
	var v1552 int64
	_ = v1552
	var v1553 int64
	_ = v1553
	var v1557 int64
	_ = v1557
	var v1564 int64
	_ = v1564
	var v1576 int32
	_ = v1576
	var v1577 int64
	_ = v1577
	var v1584 int64
	_ = v1584
	var v1585 int64
	_ = v1585
	var v1587 int64
	_ = v1587
	var v1590 int64
	_ = v1590
	var v1591 int64
	_ = v1591
	var v1593 int64
	_ = v1593
	var v1594 int64
	_ = v1594
	var v1598 int64
	_ = v1598
	var v1605 int64
	_ = v1605
	var v1616 int64
	_ = v1616
	var v1617 int64
	_ = v1617
	var v1618 int64
	_ = v1618
	var v1619 int64
	_ = v1619
	var v1621 int64
	_ = v1621
	var v1626 int32
	_ = v1626
	var v1638 int64
	_ = v1638
	var v1639 int64
	_ = v1639
	var v1640 int64
	_ = v1640
	var v1646 int32
	_ = v1646
	var v1647 int64
	_ = v1647
	var v1654 int64
	_ = v1654
	var v1655 int64
	_ = v1655
	var v1657 int64
	_ = v1657
	var v1660 int64
	_ = v1660
	var v1661 int64
	_ = v1661
	var v1663 int64
	_ = v1663
	var v1664 int64
	_ = v1664
	var v1668 int64
	_ = v1668
	var v1675 int64
	_ = v1675
	var v1687 int32
	_ = v1687
	var v1688 int64
	_ = v1688
	var v1695 int64
	_ = v1695
	var v1696 int64
	_ = v1696
	var v1698 int64
	_ = v1698
	var v1701 int64
	_ = v1701
	var v1702 int64
	_ = v1702
	var v1704 int64
	_ = v1704
	var v1705 int64
	_ = v1705
	var v1709 int64
	_ = v1709
	var v1716 int64
	_ = v1716
	var v1728 int32
	_ = v1728
	var v1729 int64
	_ = v1729
	var v1736 int64
	_ = v1736
	var v1737 int64
	_ = v1737
	var v1739 int64
	_ = v1739
	var v1742 int64
	_ = v1742
	var v1743 int64
	_ = v1743
	var v1745 int64
	_ = v1745
	var v1746 int64
	_ = v1746
	var v1750 int64
	_ = v1750
	var v1757 int64
	_ = v1757
	var v1769 int32
	_ = v1769
	var v1770 int64
	_ = v1770
	var v1771 int64
	_ = v1771
	var v1772 int64
	_ = v1772
	var v1773 int64
	_ = v1773
	var v1775 int64
	_ = v1775
	var v1776 int64
	_ = v1776
	var v1778 int64
	_ = v1778
	var v1789 int32
	_ = v1789
	var v1802 int64
	_ = v1802
	var v1805 int64
	_ = v1805
	var v1808 int32
	_ = v1808
	var v1814 int32
	_ = v1814
	var v1816 int32
	_ = v1816
	var v1828 int64
	_ = v1828
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1838 int64
	_ = v1838
	var v1847 int32
	_ = v1847
	var v1857 int64
	_ = v1857
	var v1910 int32
	_ = v1910
	var v1916 int32
	_ = v1916
	var v1926 int64
	_ = v1926
	var v1933 int32
	_ = v1933
	var v1936 int32
	_ = v1936
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1945 int32
	_ = v1945
	var v1955 int64
	_ = v1955
	var v1960 int32
	_ = v1960
	var v1964 int64
	_ = v1964
	var v1971 int64
	_ = v1971
	var v1972 int64
	_ = v1972
	var v1974 int64
	_ = v1974
	var v1977 int64
	_ = v1977
	var v1978 int64
	_ = v1978
	var v1980 int64
	_ = v1980
	var v1981 int64
	_ = v1981
	var v1985 int64
	_ = v1985
	var v1992 int64
	_ = v1992
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2007 int64
	_ = v2007
	var v2009 int64
	_ = v2009
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2022 int32
	_ = v2022
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2032 int32
	_ = v2032
	var v2035 int32
	_ = v2035
	var v2046 int32
	_ = v2046
	var v2057 int32
	_ = v2057
	var v2061 int32
	_ = v2061
	var v2065 int64
	_ = v2065
	var v2066 int32
	_ = v2066
	var v2072 int32
	_ = v2072
	var v2074 int32
	_ = v2074
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2097 int32
	_ = v2097
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2105 int32
	_ = v2105
	var v2114 int32
	_ = v2114
	var v2117 int32
	_ = v2117
	var v2121 int32
	_ = v2121
	var v2123 int32
	_ = v2123
	var v2147 int32
	_ = v2147
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2160 int32
	_ = v2160
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2175 int32
	_ = v2175
	var v2179 int32
	_ = v2179
	var v2181 int32
	_ = v2181
	var v2186 int64
	_ = v2186
	var v2189 int32
	_ = v2189
	var v2192 int32
	_ = v2192
	var v2195 int32
	_ = v2195
	var v2200 int32
	_ = v2200
	var v2203 int32
	_ = v2203
	var v2206 int32
	_ = v2206
	var v2211 int32
	_ = v2211
	var v2214 int32
	_ = v2214
	var v2217 int32
	_ = v2217
	var v2221 int32
	_ = v2221
	var v2229 int32
	_ = v2229
	var v2232 int32
	_ = v2232
	var v2244 int32
	_ = v2244
	var v2252 int64
	_ = v2252
	var v2263 int32
	_ = v2263
	var v2264 int64
	_ = v2264
	var v2271 int64
	_ = v2271
	var v2272 int64
	_ = v2272
	var v2274 int64
	_ = v2274
	var v2277 int64
	_ = v2277
	var v2278 int64
	_ = v2278
	var v2280 int64
	_ = v2280
	var v2281 int64
	_ = v2281
	var v2285 int64
	_ = v2285
	var v2292 int64
	_ = v2292
	var v2305 int64
	_ = v2305
	var v2307 int64
	_ = v2307
	var v2320 int32
	_ = v2320
	var v2331 int64
	_ = v2331
	var v2341 int32
	_ = v2341
	var v2342 int64
	_ = v2342
	var v2349 int64
	_ = v2349
	var v2350 int64
	_ = v2350
	var v2352 int64
	_ = v2352
	var v2355 int64
	_ = v2355
	var v2356 int64
	_ = v2356
	var v2358 int64
	_ = v2358
	var v2359 int64
	_ = v2359
	var v2363 int64
	_ = v2363
	var v2370 int64
	_ = v2370
	var v2382 int32
	_ = v2382
	var v2385 int64
	_ = v2385
	var v2387 int64
	_ = v2387
	var v2391 int32
	_ = v2391
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2398 int32
	_ = v2398
	var v2400 int32
	_ = v2400
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2410 int32
	_ = v2410
	var v2413 int32
	_ = v2413
	var v2424 int32
	_ = v2424
	var v2435 int32
	_ = v2435
	var v2439 int32
	_ = v2439
	var v2443 int64
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2450 int32
	_ = v2450
	var v2452 int32
	_ = v2452
	var v2471 int32
	_ = v2471
	var v2475 int32
	_ = v2475
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2483 int32
	_ = v2483
	var v2492 int32
	_ = v2492
	var v2495 int32
	_ = v2495
	var v2499 int32
	_ = v2499
	var v2501 int32
	_ = v2501
	var v2527 int32
	_ = v2527
	var v2529 int32
	_ = v2529
	var v2530 int32
	_ = v2530
	var v2538 int32
	_ = v2538
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2545 int32
	_ = v2545
	var v2546 int32
	_ = v2546
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2556 int32
	_ = v2556
	var v2560 int32
	_ = v2560
	var v2564 int32
	_ = v2564
	var v2569 int32
	_ = v2569
	var v2570 int32
	_ = v2570
	var v2571 int32
	_ = v2571
	var v2572 int32
	_ = v2572
	var v2577 int32
	_ = v2577
	var v2578 int32
	_ = v2578
	var v2581 int32
	_ = v2581
	var v2584 int32
	_ = v2584
	var v2588 int32
	_ = v2588
	var v2589 int32
	_ = v2589
	var v2592 int32
	_ = v2592
	var v2598 int32
	_ = v2598
	var v2607 int32
	_ = v2607
	var v2613 int32
	_ = v2613
	var v2618 int32
	_ = v2618
	var v2627 int32
	_ = v2627
	var v2628 int32
	_ = v2628
	var v2634 int32
	_ = v2634
	var v2657 int32
	_ = v2657
	var v2662 int32
	_ = v2662
	v3 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(592)
	m.G0 = v25
	v27 = base.I64_reinterpret_f64(l0)
	v29 = v27 & int64(4503599627370495)
	v33 = int32(2047)
	v34 = base.I32_wrap_i64(int64(base.Ui64(v27)>>(uint(int64(52))%64))) & v33
	if v34 != v33 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v25 + int32(592)
	v2662 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2657+l1))) = uint8(v2662)
	return
L2:
	;
	if v27 < int64(0) {
		goto L229
	} else {
		goto L230
	}
L3:
	;
	if base.Ui32(int32(52)) < base.Ui32(v34-int32(1023)) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	if base.B2i32(v29 != int64(0))|v34 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if v29 == int64(0) {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	goto L2
L8:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_double_to_shortest_decimal_buf[0])))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)) = uint8(v43)
	v46 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_double_to_shortest_decimal_buf[1])))
	*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v46)
	v2657 = int32(3)
	goto L1
L9:
	;
	if v27 < int64(0) {
		goto L140
	} else {
		goto L141
	}
L10:
	;
	if base.Ui64(int64(999999999999999)) < base.Ui64(v1857) {
		v1916 = v1847
		v1926 = v1857
		v1933 = int32(16)
		goto L9
	} else {
		goto L123
	}
L11:
	;
	v68 = v29 << (uint(int64(2)) % 64)
	if v34 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v53 = int64(-1)
	v56 = base.I64_extend_i32_u(int32(1075) - v34)
	if v29&(v53<<(uint(v56)%64)^v53) != int64(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v1847 = v3
	v1857 = int64(base.Ui64(v29|int64(4503599627370496)) >> (uint(v56) % 64))
	goto L10
L14:
	;
	v71 = v68 | int64(18014398509481984)
	goto L16
L15:
	;
	v71 = v68
	goto L16
L16:
	;
	v76 = base.B2i32(base.Ui32(v34) < base.Ui32(int32(2))) | base.B2i32(v29 != int64(0))
	if v34 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v1834 = v1814 + v1816
	v1838 = v1828 + base.I64_extend_i32_u(v1833)&int64(1)
	if base.Ui64(v1838) <= base.Ui64(int64(9999999999999999)) {
		v1847 = v1834
		v1857 = v1838
		goto L10
	} else {
		goto L122
	}
L18:
	;
	v1382 = v25 + int32(160)
	v1384 = int64(base.Ui64(v1379) >> (uint(int64(2)) % 64))
	v1385 = int64(0)
	v1392 = int64(32)
	v1393 = int64(687194767)
	v1395 = int64(base.Ui64(v1384) >> (uint(v1392) % 64))
	v1398 = int64(4294967295)
	v1399 = int64(1546188227)
	v1401 = v1384 & v1398
	v1402 = v1399 * v1401
	v1406 = int64(base.Ui64(v1402)>>(uint(v1392)%64)) + v1399*v1395
	v1413 = v1401*v1393 + v1406&v1398
	*(*int64)(unsafe.Add(mBase, uint32(v1382)+8)) = v1384*v1385 + v1385 + v1393*v1395 + int64(base.Ui64(v1406)>>(uint(v1392)%64)) + int64(base.Ui64(v1413)>>(uint(v1392)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1382))) = v1402&v1398 | v1413<<(uint(v1392)%64)
	goto L104
L19:
	;
	v1063 = v25 + int32(240)
	v1064 = int64(0)
	v1071 = int64(32)
	v1072 = int64(3435973836)
	v1074 = int64(base.Ui64(v1060) >> (uint(v1071) % 64))
	v1077 = int64(4294967295)
	v1078 = int64(3435973837)
	v1080 = v1060 & v1077
	v1081 = v1078 * v1080
	v1085 = int64(base.Ui64(v1081)>>(uint(v1071)%64)) + v1078*v1074
	v1092 = v1080*v1072 + v1085&v1077
	*(*int64)(unsafe.Add(mBase, uint32(v1063)+8)) = v1060*v1064 + v1064 + v1072*v1074 + int64(base.Ui64(v1085)>>(uint(v1071)%64)) + int64(base.Ui64(v1092)>>(uint(v1071)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1063))) = v1081&v1077 | v1092<<(uint(v1071)%64)
	goto L92
L20:
	;
	v80 = v34 - int32(1077)
	goto L22
L21:
	;
	v80 = int32(-1076)
	goto L22
L22:
	;
	if int32(0) <= v80 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v84 = v25 + int32(432)
	v91 = int32(base.Ui32(v80*int32(_a_F_double_to_shortest_decimal_buf_0))>>(uint(int32(18))%32)) - base.B2i32(base.Ui32(int32(3)) < base.Ui32(v80))
	v93 = v91 << (uint(int32(4)) % 32)
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_double_to_shortest_decimal_buf[2])))
	v95 = int64(0)
	v97 = v71 | int64(2)
	v103 = int64(32)
	v104 = int64(base.Ui64(v97) >> (uint(v103) % 64))
	v106 = int64(base.Ui64(v94) >> (uint(v103) % 64))
	v109 = int64(4294967295)
	v110 = v97 & v109
	v112 = v94 & v109
	v113 = v110 * v112
	v117 = int64(base.Ui64(v113)>>(uint(v103)%64)) + v110*v106
	v124 = v112*v104 + v117&v109
	*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v94*v95 + v95*v97 + v104*v106 + int64(base.Ui64(v117)>>(uint(v103)%64)) + int64(base.Ui64(v124)>>(uint(v103)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v84))) = v113&v109 | v124<<(uint(v103)%64)
	goto L26
L24:
	;
	goto L25
L25:
	;
	v653 = v25 + int32(576)
	v655 = int32(0) - v80
	v662 = int32(base.Ui32(v80*int32(-732923))>>(uint(int32(20))%32)) - base.B2i32(base.Ui32(int32(1)) < base.Ui32(v655))
	v663 = v655 - v662
	v665 = v663 << (uint(int32(4)) % 32)
	v668 = *(*int64)(unsafe.Add(mBase, uint32(v665)+uint32(_c_F_double_to_shortest_decimal_buf[3])))
	v669 = int64(0)
	v671 = v71 | int64(2)
	v677 = int64(32)
	v678 = int64(base.Ui64(v671) >> (uint(v677) % 64))
	v680 = int64(base.Ui64(v668) >> (uint(v677) % 64))
	v683 = int64(4294967295)
	v684 = v671 & v683
	v686 = v668 & v683
	v687 = v684 * v686
	v691 = int64(base.Ui64(v687)>>(uint(v677)%64)) + v684*v680
	v698 = v686*v678 + v691&v683
	*(*int64)(unsafe.Add(mBase, uint32(v653)+8)) = v668*v669 + v669*v671 + v678*v680 + int64(base.Ui64(v691)>>(uint(v677)%64)) + int64(base.Ui64(v698)>>(uint(v677)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v653))) = v687&v683 | v698<<(uint(v677)%64)
	goto L64
L26:
	;
	v136 = v25 + int32(416)
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v93)+uint32(_c_F_double_to_shortest_decimal_buf[4])))
	v138 = int64(0)
	v144 = int64(32)
	v145 = int64(base.Ui64(v97) >> (uint(v144) % 64))
	v147 = int64(base.Ui64(v137) >> (uint(v144) % 64))
	v150 = int64(4294967295)
	v151 = v97 & v150
	v153 = v137 & v150
	v154 = v151 * v153
	v158 = int64(base.Ui64(v154)>>(uint(v144)%64)) + v151*v147
	v165 = v153*v145 + v158&v150
	*(*int64)(unsafe.Add(mBase, uint32(v136)+8)) = v137*v138 + v138*v97 + v145*v147 + int64(base.Ui64(v158)>>(uint(v144)%64)) + int64(base.Ui64(v165)>>(uint(v144)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v136))) = v154&v150 | v165<<(uint(v144)%64)
	goto L27
L27:
	;
	v177 = v25 + int32(400)
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v25)+440))
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v25)+416))
	v180 = v178 + v179
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v25)+424))
	v184 = v181 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v180) < base.Ui64(v178)))
	v190 = v91 - v80 + int32(base.Ui32(v91*int32(_a_F_double_to_shortest_decimal_buf_1))>>(uint(int32(19))%32))
	v192 = v190 + int32(58)
	if v192&int32(64) != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v217 = v25 + int32(336)
	v218 = int64(0)
	v222 = v71 + base.I64_extend_i32_s(v76^int32(-1))
	v228 = int64(32)
	v229 = int64(base.Ui64(v222) >> (uint(v228) % 64))
	v231 = int64(base.Ui64(v94) >> (uint(v228) % 64))
	v234 = int64(4294967295)
	v235 = v222 & v234
	v237 = v94 & v234
	v238 = v235 * v237
	v242 = int64(base.Ui64(v238)>>(uint(v228)%64)) + v235*v231
	v249 = v237*v229 + v242&v234
	*(*int64)(unsafe.Add(mBase, uint32(v217)+8)) = v94*v218 + v218*v222 + v229*v231 + int64(base.Ui64(v242)>>(uint(v228)%64)) + int64(base.Ui64(v249)>>(uint(v228)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v217))) = v238&v234 | v249<<(uint(v228)%64)
	goto L34
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v177))) = v211
	*(*int64)(unsafe.Add(mBase, uint32(v177)+8)) = v212
	goto L28
L30:
	;
	v211 = int64(base.Ui64(v184) >> (uint(base.I64_extend_i32_u(v190+int32(-6))) % 64))
	v212 = int64(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	if v192 == int32(0) {
		v211 = v180
		v212 = v184
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v207 = base.I64_extend_i32_u(v192)
	v211 = v184<<(uint(base.I64_extend_i32_u(int32(64)-v192))%64) | int64(base.Ui64(v180)>>(uint(v207)%64))
	v212 = int64(base.Ui64(v184) >> (uint(v207) % 64))
	goto L29
L34:
	;
	v261 = v25 + int32(320)
	v262 = int64(0)
	v268 = int64(32)
	v269 = int64(base.Ui64(v222) >> (uint(v268) % 64))
	v271 = int64(base.Ui64(v137) >> (uint(v268) % 64))
	v274 = int64(4294967295)
	v275 = v222 & v274
	v277 = v137 & v274
	v278 = v275 * v277
	v282 = int64(base.Ui64(v278)>>(uint(v268)%64)) + v275*v271
	v289 = v277*v269 + v282&v274
	*(*int64)(unsafe.Add(mBase, uint32(v261)+8)) = v137*v262 + v262*v222 + v269*v271 + int64(base.Ui64(v282)>>(uint(v268)%64)) + int64(base.Ui64(v289)>>(uint(v268)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v261))) = v278&v274 | v289<<(uint(v268)%64)
	goto L35
L35:
	;
	v301 = v25 + int32(304)
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v25)+344))
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v25)+320))
	v304 = v302 + v303
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v25)+328))
	v308 = v305 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v304) < base.Ui64(v302)))
	if v192&int32(64) != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v333 = v25 + int32(384)
	v334 = int64(0)
	v340 = int64(32)
	v341 = int64(base.Ui64(v71) >> (uint(v340) % 64))
	v343 = int64(base.Ui64(v94) >> (uint(v340) % 64))
	v346 = int64(4294967295)
	v347 = v71 & v346
	v349 = v94 & v346
	v350 = v347 * v349
	v354 = int64(base.Ui64(v350)>>(uint(v340)%64)) + v347*v343
	v361 = v349*v341 + v354&v346
	*(*int64)(unsafe.Add(mBase, uint32(v333)+8)) = v94*v334 + v334*v71 + v341*v343 + int64(base.Ui64(v354)>>(uint(v340)%64)) + int64(base.Ui64(v361)>>(uint(v340)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v333))) = v350&v346 | v361<<(uint(v340)%64)
	goto L42
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v301))) = v327
	*(*int64)(unsafe.Add(mBase, uint32(v301)+8)) = v328
	goto L36
L38:
	;
	v327 = int64(base.Ui64(v308) >> (uint(base.I64_extend_i32_u(v190+int32(-6))) % 64))
	v328 = int64(0)
	goto L37
L39:
	;
	goto L40
L40:
	;
	if v192 == int32(0) {
		v327 = v304
		v328 = v308
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v323 = base.I64_extend_i32_u(v192)
	v327 = v308<<(uint(base.I64_extend_i32_u(int32(64)-v192))%64) | int64(base.Ui64(v304)>>(uint(v323)%64))
	v328 = int64(base.Ui64(v308) >> (uint(v323) % 64))
	goto L37
L42:
	;
	v373 = v25 + int32(368)
	v374 = int64(0)
	v380 = int64(32)
	v381 = int64(base.Ui64(v71) >> (uint(v380) % 64))
	v383 = int64(base.Ui64(v137) >> (uint(v380) % 64))
	v386 = int64(4294967295)
	v387 = v71 & v386
	v389 = v137 & v386
	v390 = v387 * v389
	v394 = int64(base.Ui64(v390)>>(uint(v380)%64)) + v387*v383
	v401 = v389*v381 + v394&v386
	*(*int64)(unsafe.Add(mBase, uint32(v373)+8)) = v137*v374 + v374*v71 + v381*v383 + int64(base.Ui64(v394)>>(uint(v380)%64)) + int64(base.Ui64(v401)>>(uint(v380)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v373))) = v390&v386 | v401<<(uint(v380)%64)
	goto L43
L43:
	;
	v413 = v25 + int32(352)
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v25)+392))
	v415 = *(*int64)(unsafe.Add(mBase, uint32(v25)+368))
	v416 = v414 + v415
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v25)+376))
	v420 = v417 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v416) < base.Ui64(v414)))
	if v192&int32(64) != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v444 = *(*int64)(unsafe.Add(mBase, uint32(v25)+352))
	v445 = *(*int64)(unsafe.Add(mBase, uint32(v25)+304))
	v446 = *(*int64)(unsafe.Add(mBase, uint32(v25)+400))
	if base.Ui32(int32(21)) < base.Ui32(v91) {
		v1364 = v91
		v1376 = v444
		v1377 = v445
		v1379 = v446
		goto L18
	} else {
		goto L50
	}
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v413))) = v439
	*(*int64)(unsafe.Add(mBase, uint32(v413)+8)) = v440
	goto L44
L46:
	;
	v439 = int64(base.Ui64(v420) >> (uint(base.I64_extend_i32_u(v190+int32(-6))) % 64))
	v440 = int64(0)
	goto L45
L47:
	;
	goto L48
L48:
	;
	if v192 == int32(0) {
		v439 = v416
		v440 = v420
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v435 = base.I64_extend_i32_u(v192)
	v439 = v420<<(uint(base.I64_extend_i32_u(int32(64)-v192))%64) | int64(base.Ui64(v416)>>(uint(v435)%64))
	v440 = int64(base.Ui64(v420) >> (uint(v435) % 64))
	goto L45
L50:
	;
	v450 = v25 + int32(288)
	v451 = int64(0)
	v458 = int64(32)
	v459 = int64(3435973836)
	v461 = int64(base.Ui64(v71) >> (uint(v458) % 64))
	v464 = int64(4294967295)
	v465 = int64(3435973837)
	v467 = v71 & v464
	v468 = v465 * v467
	v472 = int64(base.Ui64(v468)>>(uint(v458)%64)) + v465*v461
	v479 = v467*v459 + v472&v464
	*(*int64)(unsafe.Add(mBase, uint32(v450)+8)) = v71*v451 + v451 + v459*v461 + int64(base.Ui64(v472)>>(uint(v458)%64)) + int64(base.Ui64(v479)>>(uint(v458)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v450))) = v468&v464 | v479<<(uint(v458)%64)
	goto L51
L51:
	;
	v490 = int32(0)
	v491 = *(*int64)(unsafe.Add(mBase, uint32(v25)+296))
	if base.I32_wrap_i64(int64(base.Ui64(v491)>>(uint(int64(2))%64))*int64(4294967291)+v71) == v490 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v503 = v490
	v515 = v71
	goto L55
L53:
	;
	goto L54
L54:
	;
	v578 = v490
	v591 = v97
	goto L60
L55:
	;
	v523 = v25 + int32(256)
	v524 = int64(0)
	v531 = int64(32)
	v532 = int64(3435973836)
	v534 = int64(base.Ui64(v515) >> (uint(v531) % 64))
	v537 = int64(4294967295)
	v538 = int64(3435973837)
	v540 = v515 & v537
	v541 = v538 * v540
	v545 = int64(base.Ui64(v541)>>(uint(v531)%64)) + v538*v534
	v552 = v540*v532 + v545&v537
	*(*int64)(unsafe.Add(mBase, uint32(v523)+8)) = v515*v524 + v524 + v532*v534 + int64(base.Ui64(v545)>>(uint(v531)%64)) + int64(base.Ui64(v552)>>(uint(v531)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v523))) = v541&v537 | v552<<(uint(v531)%64)
	goto L57
L56:
	;
	if base.Ui32(v503) < base.Ui32(v91) {
		v1364 = v91
		v1376 = v444
		v1377 = v445
		v1379 = v446
		goto L18
	} else {
		goto L59
	}
L57:
	;
	v565 = *(*int64)(unsafe.Add(mBase, uint32(v25)+264))
	v567 = int64(base.Ui64(v565) >> (uint(int64(2)) % 64))
	if base.I32_wrap_i64(v515+v567*int64(4294967291)) == int32(0) {
		v503 = v503 + int32(1)
		v515 = v567
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v1045 = v91
	v1057 = v444
	v1058 = v445
	v1060 = v446
	goto L19
L60:
	;
	v598 = v25 + int32(272)
	v599 = int64(0)
	v606 = int64(32)
	v607 = int64(3435973836)
	v609 = int64(base.Ui64(v591) >> (uint(v606) % 64))
	v612 = int64(4294967295)
	v613 = int64(3435973837)
	v615 = v591 & v612
	v616 = v613 * v615
	v620 = int64(base.Ui64(v616)>>(uint(v606)%64)) + v613*v609
	v627 = v615*v607 + v620&v612
	*(*int64)(unsafe.Add(mBase, uint32(v598)+8)) = v591*v599 + v599 + v607*v609 + int64(base.Ui64(v620)>>(uint(v606)%64)) + int64(base.Ui64(v627)>>(uint(v606)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v598))) = v616&v612 | v627<<(uint(v606)%64)
	goto L62
L61:
	;
	v1364 = v91
	v1376 = v444
	v1377 = v445
	v1379 = v446 - base.I64_extend_i32_u(base.B2i32(base.Ui32(v91) <= base.Ui32(v578)))
	goto L18
L62:
	;
	v640 = *(*int64)(unsafe.Add(mBase, uint32(v25)+280))
	v642 = int64(base.Ui64(v640) >> (uint(int64(2)) % 64))
	if base.I32_wrap_i64(v591+v642*int64(4294967291)) == int32(0) {
		v578 = v578 + int32(1)
		v591 = v642
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v710 = v25 + int32(560)
	v713 = *(*int64)(unsafe.Add(mBase, uint32(v665)+uint32(_c_F_double_to_shortest_decimal_buf[5])))
	v714 = int64(0)
	v720 = int64(32)
	v721 = int64(base.Ui64(v671) >> (uint(v720) % 64))
	v723 = int64(base.Ui64(v713) >> (uint(v720) % 64))
	v726 = int64(4294967295)
	v727 = v671 & v726
	v729 = v713 & v726
	v730 = v727 * v729
	v734 = int64(base.Ui64(v730)>>(uint(v720)%64)) + v727*v723
	v741 = v729*v721 + v734&v726
	*(*int64)(unsafe.Add(mBase, uint32(v710)+8)) = v713*v714 + v714*v671 + v721*v723 + int64(base.Ui64(v734)>>(uint(v720)%64)) + int64(base.Ui64(v741)>>(uint(v720)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v710))) = v730&v726 | v741<<(uint(v720)%64)
	goto L65
L65:
	;
	v753 = v25 + int32(544)
	v754 = *(*int64)(unsafe.Add(mBase, uint32(v25)+584))
	v755 = *(*int64)(unsafe.Add(mBase, uint32(v25)+560))
	v756 = v754 + v755
	v757 = *(*int64)(unsafe.Add(mBase, uint32(v25)+568))
	v760 = v757 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v756) < base.Ui64(v754)))
	v765 = v662 - int32(base.Ui32(v663*int32(_a_F_double_to_shortest_decimal_buf_1))>>(uint(int32(19))%32))
	v767 = v765 + int32(56)
	if v767&int32(64) != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v792 = v25 + int32(480)
	v793 = int64(0)
	v797 = v71 + base.I64_extend_i32_s(v76^int32(-1))
	v803 = int64(32)
	v804 = int64(base.Ui64(v797) >> (uint(v803) % 64))
	v806 = int64(base.Ui64(v668) >> (uint(v803) % 64))
	v809 = int64(4294967295)
	v810 = v797 & v809
	v812 = v668 & v809
	v813 = v810 * v812
	v817 = int64(base.Ui64(v813)>>(uint(v803)%64)) + v810*v806
	v824 = v812*v804 + v817&v809
	*(*int64)(unsafe.Add(mBase, uint32(v792)+8)) = v668*v793 + v793*v797 + v804*v806 + int64(base.Ui64(v817)>>(uint(v803)%64)) + int64(base.Ui64(v824)>>(uint(v803)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v792))) = v813&v809 | v824<<(uint(v803)%64)
	goto L72
L67:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v753))) = v786
	*(*int64)(unsafe.Add(mBase, uint32(v753)+8)) = v787
	goto L66
L68:
	;
	v786 = int64(base.Ui64(v760) >> (uint(base.I64_extend_i32_u(v765+int32(-8))) % 64))
	v787 = int64(0)
	goto L67
L69:
	;
	goto L70
L70:
	;
	if v767 == int32(0) {
		v786 = v756
		v787 = v760
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v782 = base.I64_extend_i32_u(v767)
	v786 = v760<<(uint(base.I64_extend_i32_u(int32(64)-v767))%64) | int64(base.Ui64(v756)>>(uint(v782)%64))
	v787 = int64(base.Ui64(v760) >> (uint(v782) % 64))
	goto L67
L72:
	;
	v836 = v25 + int32(464)
	v837 = int64(0)
	v843 = int64(32)
	v844 = int64(base.Ui64(v797) >> (uint(v843) % 64))
	v846 = int64(base.Ui64(v713) >> (uint(v843) % 64))
	v849 = int64(4294967295)
	v850 = v797 & v849
	v852 = v713 & v849
	v853 = v850 * v852
	v857 = int64(base.Ui64(v853)>>(uint(v843)%64)) + v850*v846
	v864 = v852*v844 + v857&v849
	*(*int64)(unsafe.Add(mBase, uint32(v836)+8)) = v713*v837 + v837*v797 + v844*v846 + int64(base.Ui64(v857)>>(uint(v843)%64)) + int64(base.Ui64(v864)>>(uint(v843)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v836))) = v853&v849 | v864<<(uint(v843)%64)
	goto L73
L73:
	;
	v876 = v25 + int32(448)
	v877 = *(*int64)(unsafe.Add(mBase, uint32(v25)+488))
	v878 = *(*int64)(unsafe.Add(mBase, uint32(v25)+464))
	v879 = v877 + v878
	v880 = *(*int64)(unsafe.Add(mBase, uint32(v25)+472))
	v883 = v880 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v879) < base.Ui64(v877)))
	if v767&int32(64) != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v908 = v25 + int32(528)
	v909 = int64(0)
	v915 = int64(32)
	v916 = int64(base.Ui64(v71) >> (uint(v915) % 64))
	v918 = int64(base.Ui64(v668) >> (uint(v915) % 64))
	v921 = int64(4294967295)
	v922 = v71 & v921
	v924 = v668 & v921
	v925 = v922 * v924
	v929 = int64(base.Ui64(v925)>>(uint(v915)%64)) + v922*v918
	v936 = v924*v916 + v929&v921
	*(*int64)(unsafe.Add(mBase, uint32(v908)+8)) = v668*v909 + v909*v71 + v916*v918 + int64(base.Ui64(v929)>>(uint(v915)%64)) + int64(base.Ui64(v936)>>(uint(v915)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v908))) = v925&v921 | v936<<(uint(v915)%64)
	goto L80
L75:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v876))) = v902
	*(*int64)(unsafe.Add(mBase, uint32(v876)+8)) = v903
	goto L74
L76:
	;
	v902 = int64(base.Ui64(v883) >> (uint(base.I64_extend_i32_u(v765+int32(-8))) % 64))
	v903 = int64(0)
	goto L75
L77:
	;
	goto L78
L78:
	;
	if v767 == int32(0) {
		v902 = v879
		v903 = v883
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v898 = base.I64_extend_i32_u(v767)
	v902 = v883<<(uint(base.I64_extend_i32_u(int32(64)-v767))%64) | int64(base.Ui64(v879)>>(uint(v898)%64))
	v903 = int64(base.Ui64(v883) >> (uint(v898) % 64))
	goto L75
L80:
	;
	v948 = v25 + int32(512)
	v949 = int64(0)
	v955 = int64(32)
	v956 = int64(base.Ui64(v71) >> (uint(v955) % 64))
	v958 = int64(base.Ui64(v713) >> (uint(v955) % 64))
	v961 = int64(4294967295)
	v962 = v71 & v961
	v964 = v713 & v961
	v965 = v962 * v964
	v969 = int64(base.Ui64(v965)>>(uint(v955)%64)) + v962*v958
	v976 = v964*v956 + v969&v961
	*(*int64)(unsafe.Add(mBase, uint32(v948)+8)) = v713*v949 + v949*v71 + v956*v958 + int64(base.Ui64(v969)>>(uint(v955)%64)) + int64(base.Ui64(v976)>>(uint(v955)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v948))) = v965&v961 | v976<<(uint(v955)%64)
	goto L81
L81:
	;
	v988 = v25 + int32(496)
	v989 = *(*int64)(unsafe.Add(mBase, uint32(v25)+536))
	v990 = *(*int64)(unsafe.Add(mBase, uint32(v25)+512))
	v991 = v989 + v990
	v992 = *(*int64)(unsafe.Add(mBase, uint32(v25)+520))
	v995 = v992 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v991) < base.Ui64(v989)))
	if v767&int32(64) != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v1019 = v80 + v662
	v1020 = *(*int64)(unsafe.Add(mBase, uint32(v25)+496))
	v1021 = *(*int64)(unsafe.Add(mBase, uint32(v25)+448))
	v1022 = *(*int64)(unsafe.Add(mBase, uint32(v25)+544))
	if base.Ui32(v662) <= base.Ui32(int32(1)) {
		goto L88
	} else {
		goto L89
	}
L83:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v988))) = v1014
	*(*int64)(unsafe.Add(mBase, uint32(v988)+8)) = v1015
	goto L82
L84:
	;
	v1014 = int64(base.Ui64(v995) >> (uint(base.I64_extend_i32_u(v765+int32(-8))) % 64))
	v1015 = int64(0)
	goto L83
L85:
	;
	goto L86
L86:
	;
	if v767 == int32(0) {
		v1014 = v991
		v1015 = v995
		goto L83
	} else {
		goto L87
	}
L87:
	;
	v1010 = base.I64_extend_i32_u(v767)
	v1014 = v995<<(uint(base.I64_extend_i32_u(int32(64)-v767))%64) | int64(base.Ui64(v991)>>(uint(v1010)%64))
	v1015 = int64(base.Ui64(v995) >> (uint(v1010) % 64))
	goto L83
L88:
	;
	v1045 = v1019
	v1057 = v1020
	v1058 = v1021
	v1060 = v1022 - int64(1)
	goto L19
L89:
	;
	goto L90
L90:
	;
	v1029 = int64(-1)
	if base.B2i32(base.Ui32(int32(62)) < base.Ui32(v662))|base.B2i32(v71&(v1029<<(uint(base.I64_extend_i32_u(v662-int32(1)))%64)^v1029) != int64(0)) != 0 {
		v1364 = v1019
		v1376 = v1020
		v1377 = v1021
		v1379 = v1022
		goto L18
	} else {
		goto L91
	}
L91:
	;
	v1045 = v1019
	v1057 = v1020
	v1058 = v1021
	v1060 = v1022
	goto L19
L92:
	;
	v1104 = v25 + int32(224)
	v1105 = int64(0)
	v1112 = int64(32)
	v1113 = int64(3435973836)
	v1115 = int64(base.Ui64(v1058) >> (uint(v1112) % 64))
	v1118 = int64(4294967295)
	v1119 = int64(3435973837)
	v1121 = v1058 & v1118
	v1122 = v1119 * v1121
	v1126 = int64(base.Ui64(v1122)>>(uint(v1112)%64)) + v1119*v1115
	v1133 = v1121*v1113 + v1126&v1118
	*(*int64)(unsafe.Add(mBase, uint32(v1104)+8)) = v1058*v1105 + v1105 + v1113*v1115 + int64(base.Ui64(v1126)>>(uint(v1112)%64)) + int64(base.Ui64(v1133)>>(uint(v1112)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1104))) = v1122&v1118 | v1133<<(uint(v1112)%64)
	goto L93
L93:
	;
	v1144 = int32(0)
	v1145 = *(*int64)(unsafe.Add(mBase, uint32(v25)+248))
	v1146 = int64(3)
	v1147 = int64(base.Ui64(v1145) >> (uint(v1146) % 64))
	v1148 = *(*int64)(unsafe.Add(mBase, uint32(v25)+232))
	v1150 = int64(base.Ui64(v1148) >> (uint(v1146) % 64))
	if base.Ui64(v1147) <= base.Ui64(v1150) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v1349 = v1328 & int32(255)
	v1814 = v1327
	v1816 = v1045
	v1828 = v1341
	v1833 = (base.I32_wrap_i64(v1341)|(v1346|base.B2i32(v1349 != int32(5))))&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v1349)) | base.B2i32(v1341 == v1342)
	goto L17
L95:
	;
	v1152 = int32(0)
	v1327 = v1144
	v1328 = v1152
	v1341 = v1057
	v1342 = v1058
	v1346 = v1152
	goto L94
L96:
	;
	goto L97
L97:
	;
	v1157 = int32(1)
	v1159 = v1144
	v1160 = int32(0)
	v1171 = v1150
	v1172 = v1147
	v1173 = v1057
	goto L98
L98:
	;
	v1179 = v25 + int32(208)
	v1180 = int64(0)
	v1187 = int64(32)
	v1188 = int64(3435973836)
	v1190 = int64(base.Ui64(v1173) >> (uint(v1187) % 64))
	v1193 = int64(4294967295)
	v1194 = int64(3435973837)
	v1196 = v1173 & v1193
	v1197 = v1194 * v1196
	v1201 = int64(base.Ui64(v1197)>>(uint(v1187)%64)) + v1194*v1190
	v1208 = v1196*v1188 + v1201&v1193
	*(*int64)(unsafe.Add(mBase, uint32(v1179)+8)) = v1173*v1180 + v1180 + v1188*v1190 + int64(base.Ui64(v1201)>>(uint(v1187)%64)) + int64(base.Ui64(v1208)>>(uint(v1187)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1179))) = v1197&v1193 | v1208<<(uint(v1187)%64)
	goto L100
L99:
	;
	v1327 = v1302
	v1328 = v1314
	v1341 = v1310
	v1342 = v1171
	v1346 = v1307 ^ int32(1)
	goto L94
L100:
	;
	v1220 = v25 + int32(192)
	v1221 = int64(0)
	v1228 = int64(32)
	v1229 = int64(3435973836)
	v1231 = int64(base.Ui64(v1172) >> (uint(v1228) % 64))
	v1234 = int64(4294967295)
	v1235 = int64(3435973837)
	v1237 = v1172 & v1234
	v1238 = v1235 * v1237
	v1242 = int64(base.Ui64(v1238)>>(uint(v1228)%64)) + v1235*v1231
	v1249 = v1237*v1229 + v1242&v1234
	*(*int64)(unsafe.Add(mBase, uint32(v1220)+8)) = v1172*v1221 + v1221 + v1229*v1231 + int64(base.Ui64(v1242)>>(uint(v1228)%64)) + int64(base.Ui64(v1249)>>(uint(v1228)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1220))) = v1238&v1234 | v1249<<(uint(v1228)%64)
	goto L101
L101:
	;
	v1261 = v25 + int32(176)
	v1262 = int64(0)
	v1269 = int64(32)
	v1270 = int64(3435973836)
	v1272 = int64(base.Ui64(v1171) >> (uint(v1269) % 64))
	v1275 = int64(4294967295)
	v1276 = int64(3435973837)
	v1278 = v1171 & v1275
	v1279 = v1276 * v1278
	v1283 = int64(base.Ui64(v1279)>>(uint(v1269)%64)) + v1276*v1272
	v1290 = v1278*v1270 + v1283&v1275
	*(*int64)(unsafe.Add(mBase, uint32(v1261)+8)) = v1171*v1262 + v1262 + v1270*v1272 + int64(base.Ui64(v1283)>>(uint(v1269)%64)) + int64(base.Ui64(v1290)>>(uint(v1269)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1261))) = v1279&v1275 | v1290<<(uint(v1269)%64)
	goto L102
L102:
	;
	v1302 = v1159 + int32(1)
	v1307 = v1157 & base.B2i32(v1160&int32(255) == int32(0))
	v1308 = *(*int64)(unsafe.Add(mBase, uint32(v25)+216))
	v1309 = int64(3)
	v1310 = int64(base.Ui64(v1308) >> (uint(v1309) % 64))
	v1314 = base.I32_wrap_i64(v1173 + v1310*int64(246))
	v1315 = *(*int64)(unsafe.Add(mBase, uint32(v25)+200))
	v1317 = int64(base.Ui64(v1315) >> (uint(v1309) % 64))
	v1318 = *(*int64)(unsafe.Add(mBase, uint32(v25)+184))
	v1320 = int64(base.Ui64(v1318) >> (uint(v1309) % 64))
	if base.Ui64(v1320) < base.Ui64(v1317) {
		v1157 = v1307
		v1159 = v1302
		v1160 = v1314
		v1171 = v1320
		v1172 = v1317
		v1173 = v1310
		goto L98
	} else {
		goto L103
	}
L103:
	;
	goto L99
L104:
	;
	v1425 = v25 + int32(144)
	v1427 = int64(base.Ui64(v1377) >> (uint(int64(2)) % 64))
	v1428 = int64(0)
	v1435 = int64(32)
	v1436 = int64(687194767)
	v1438 = int64(base.Ui64(v1427) >> (uint(v1435) % 64))
	v1441 = int64(4294967295)
	v1442 = int64(1546188227)
	v1444 = v1427 & v1441
	v1445 = v1442 * v1444
	v1449 = int64(base.Ui64(v1445)>>(uint(v1435)%64)) + v1442*v1438
	v1456 = v1444*v1436 + v1449&v1441
	*(*int64)(unsafe.Add(mBase, uint32(v1425)+8)) = v1427*v1428 + v1428 + v1436*v1438 + int64(base.Ui64(v1449)>>(uint(v1435)%64)) + int64(base.Ui64(v1456)>>(uint(v1435)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1425))) = v1445&v1441 | v1456<<(uint(v1435)%64)
	goto L105
L105:
	;
	v1468 = *(*int64)(unsafe.Add(mBase, uint32(v25)+168))
	v1469 = int64(2)
	v1470 = int64(base.Ui64(v1468) >> (uint(v1469) % 64))
	v1471 = *(*int64)(unsafe.Add(mBase, uint32(v25)+152))
	v1473 = int64(base.Ui64(v1471) >> (uint(v1469) % 64))
	if base.Ui64(v1470) <= base.Ui64(v1473) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v1535 = v25 + int32(112)
	v1536 = int64(0)
	v1543 = int64(32)
	v1544 = int64(3435973836)
	v1546 = int64(base.Ui64(v1530) >> (uint(v1543) % 64))
	v1549 = int64(4294967295)
	v1550 = int64(3435973837)
	v1552 = v1530 & v1549
	v1553 = v1550 * v1552
	v1557 = int64(base.Ui64(v1553)>>(uint(v1543)%64)) + v1550*v1546
	v1564 = v1552*v1544 + v1557&v1549
	*(*int64)(unsafe.Add(mBase, uint32(v1535)+8)) = v1530*v1536 + v1536 + v1544*v1546 + int64(base.Ui64(v1557)>>(uint(v1543)%64)) + int64(base.Ui64(v1564)>>(uint(v1543)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1535))) = v1553&v1549 | v1564<<(uint(v1543)%64)
	goto L111
L107:
	;
	v1529 = int32(0)
	v1530 = v1379
	v1531 = v1376
	v1532 = v1377
	v1533 = int32(0)
	goto L106
L108:
	;
	goto L109
L109:
	;
	v1477 = v25 + int32(128)
	v1479 = int64(base.Ui64(v1376) >> (uint(int64(2)) % 64))
	v1480 = int64(0)
	v1487 = int64(32)
	v1488 = int64(687194767)
	v1490 = int64(base.Ui64(v1479) >> (uint(v1487) % 64))
	v1493 = int64(4294967295)
	v1494 = int64(1546188227)
	v1496 = v1479 & v1493
	v1497 = v1494 * v1496
	v1501 = int64(base.Ui64(v1497)>>(uint(v1487)%64)) + v1494*v1490
	v1508 = v1496*v1488 + v1501&v1493
	*(*int64)(unsafe.Add(mBase, uint32(v1477)+8)) = v1479*v1480 + v1480 + v1488*v1490 + int64(base.Ui64(v1501)>>(uint(v1487)%64)) + int64(base.Ui64(v1508)>>(uint(v1487)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1477))) = v1497&v1493 | v1508<<(uint(v1487)%64)
	goto L110
L110:
	;
	v1519 = *(*int64)(unsafe.Add(mBase, uint32(v25)+136))
	v1521 = int64(base.Ui64(v1519) >> (uint(int64(2)) % 64))
	v1529 = base.B2i32(base.Ui32(int32(49)) < base.Ui32(base.I32_wrap_i64(v1521*int64(4294967196)+v1376)))
	v1530 = v1470
	v1531 = v1521
	v1532 = v1473
	v1533 = int32(2)
	goto L106
L111:
	;
	v1576 = v25 + int32(96)
	v1577 = int64(0)
	v1584 = int64(32)
	v1585 = int64(3435973836)
	v1587 = int64(base.Ui64(v1532) >> (uint(v1584) % 64))
	v1590 = int64(4294967295)
	v1591 = int64(3435973837)
	v1593 = v1532 & v1590
	v1594 = v1591 * v1593
	v1598 = int64(base.Ui64(v1594)>>(uint(v1584)%64)) + v1591*v1587
	v1605 = v1593*v1585 + v1598&v1590
	*(*int64)(unsafe.Add(mBase, uint32(v1576)+8)) = v1532*v1577 + v1577 + v1585*v1587 + int64(base.Ui64(v1598)>>(uint(v1584)%64)) + int64(base.Ui64(v1605)>>(uint(v1584)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1576))) = v1594&v1590 | v1605<<(uint(v1584)%64)
	goto L112
L112:
	;
	v1616 = *(*int64)(unsafe.Add(mBase, uint32(v25)+120))
	v1617 = int64(3)
	v1618 = int64(base.Ui64(v1616) >> (uint(v1617) % 64))
	v1619 = *(*int64)(unsafe.Add(mBase, uint32(v25)+104))
	v1621 = int64(base.Ui64(v1619) >> (uint(v1617) % 64))
	if base.Ui64(v1621) < base.Ui64(v1618) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v1626 = v1533
	v1638 = v1618
	v1639 = v1531
	v1640 = v1621
	goto L116
L114:
	;
	v1789 = v1533
	v1802 = v1531
	v1805 = v1532
	v1808 = v1529
	goto L115
L115:
	;
	v1814 = v1789
	v1816 = v1364
	v1828 = v1802
	v1833 = v1808 | base.B2i32(v1802 == v1805)
	goto L17
L116:
	;
	v1646 = v25 + int32(80)
	v1647 = int64(0)
	v1654 = int64(32)
	v1655 = int64(3435973836)
	v1657 = int64(base.Ui64(v1639) >> (uint(v1654) % 64))
	v1660 = int64(4294967295)
	v1661 = int64(3435973837)
	v1663 = v1639 & v1660
	v1664 = v1661 * v1663
	v1668 = int64(base.Ui64(v1664)>>(uint(v1654)%64)) + v1661*v1657
	v1675 = v1663*v1655 + v1668&v1660
	*(*int64)(unsafe.Add(mBase, uint32(v1646)+8)) = v1639*v1647 + v1647 + v1655*v1657 + int64(base.Ui64(v1668)>>(uint(v1654)%64)) + int64(base.Ui64(v1675)>>(uint(v1654)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1646))) = v1664&v1660 | v1675<<(uint(v1654)%64)
	goto L118
L117:
	;
	v1789 = v1769
	v1802 = v1772
	v1805 = v1640
	v1808 = base.B2i32(base.Ui32(int32(4)) < base.Ui32(base.I32_wrap_i64(v1772*int64(4294967286)+v1639)))
	goto L115
L118:
	;
	v1687 = v25 - int32(-64)
	v1688 = int64(0)
	v1695 = int64(32)
	v1696 = int64(3435973836)
	v1698 = int64(base.Ui64(v1638) >> (uint(v1695) % 64))
	v1701 = int64(4294967295)
	v1702 = int64(3435973837)
	v1704 = v1638 & v1701
	v1705 = v1702 * v1704
	v1709 = int64(base.Ui64(v1705)>>(uint(v1695)%64)) + v1702*v1698
	v1716 = v1704*v1696 + v1709&v1701
	*(*int64)(unsafe.Add(mBase, uint32(v1687)+8)) = v1638*v1688 + v1688 + v1696*v1698 + int64(base.Ui64(v1709)>>(uint(v1695)%64)) + int64(base.Ui64(v1716)>>(uint(v1695)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1687))) = v1705&v1701 | v1716<<(uint(v1695)%64)
	goto L119
L119:
	;
	v1728 = v25 + int32(48)
	v1729 = int64(0)
	v1736 = int64(32)
	v1737 = int64(3435973836)
	v1739 = int64(base.Ui64(v1640) >> (uint(v1736) % 64))
	v1742 = int64(4294967295)
	v1743 = int64(3435973837)
	v1745 = v1640 & v1742
	v1746 = v1743 * v1745
	v1750 = int64(base.Ui64(v1746)>>(uint(v1736)%64)) + v1743*v1739
	v1757 = v1745*v1737 + v1750&v1742
	*(*int64)(unsafe.Add(mBase, uint32(v1728)+8)) = v1640*v1729 + v1729 + v1737*v1739 + int64(base.Ui64(v1750)>>(uint(v1736)%64)) + int64(base.Ui64(v1757)>>(uint(v1736)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1728))) = v1746&v1742 | v1757<<(uint(v1736)%64)
	goto L120
L120:
	;
	v1769 = v1626 + int32(1)
	v1770 = *(*int64)(unsafe.Add(mBase, uint32(v25)+88))
	v1771 = int64(3)
	v1772 = int64(base.Ui64(v1770) >> (uint(v1771) % 64))
	v1773 = *(*int64)(unsafe.Add(mBase, uint32(v25)+72))
	v1775 = int64(base.Ui64(v1773) >> (uint(v1771) % 64))
	v1776 = *(*int64)(unsafe.Add(mBase, uint32(v25)+56))
	v1778 = int64(base.Ui64(v1776) >> (uint(v1771) % 64))
	if base.Ui64(v1778) < base.Ui64(v1775) {
		v1626 = v1769
		v1638 = v1775
		v1639 = v1772
		v1640 = v1778
		goto L116
	} else {
		goto L121
	}
L121:
	;
	goto L117
L122:
	;
	v1916 = v1834
	v1926 = v1838
	v1933 = int32(17)
	goto L9
L123:
	;
	if base.Ui64(int64(99999999999999)) < base.Ui64(v1857) {
		v1916 = v1847
		v1926 = v1857
		v1933 = int32(15)
		goto L9
	} else {
		goto L124
	}
L124:
	;
	if base.Ui64(int64(9999999999999)) < base.Ui64(v1857) {
		v1916 = v1847
		v1926 = v1857
		v1933 = int32(14)
		goto L9
	} else {
		goto L125
	}
L125:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v1857) {
		v1916 = v1847
		v1926 = v1857
		v1933 = int32(13)
		goto L9
	} else {
		goto L126
	}
L126:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v1857) {
		v1916 = v1847
		v1926 = v1857
		v1933 = int32(12)
		goto L9
	} else {
		goto L127
	}
L127:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v1857) {
		v1916 = v1847
		v1926 = v1857
		v1933 = int32(11)
		goto L9
	} else {
		goto L128
	}
L128:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v1857) {
		v1916 = v1847
		v1926 = v1857
		v1933 = int32(10)
		goto L9
	} else {
		goto L129
	}
L129:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v1857) {
		v1916 = v1847
		v1926 = v1857
		v1933 = int32(9)
		goto L9
	} else {
		goto L130
	}
L130:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v1857) {
		v1916 = v1847
		v1926 = v1857
		v1933 = int32(8)
		goto L9
	} else {
		goto L131
	}
L131:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v1857) {
		v1916 = v1847
		v1926 = v1857
		v1933 = int32(7)
		goto L9
	} else {
		goto L132
	}
L132:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v1857) {
		v1916 = v1847
		v1926 = v1857
		v1933 = int32(6)
		goto L9
	} else {
		goto L133
	}
L133:
	;
	if base.Ui64(int64(9999)) < base.Ui64(v1857) {
		v1916 = v1847
		v1926 = v1857
		v1933 = int32(5)
		goto L9
	} else {
		goto L134
	}
L134:
	;
	if base.Ui64(int64(999)) < base.Ui64(v1857) {
		v1916 = v1847
		v1926 = v1857
		v1933 = int32(4)
		goto L9
	} else {
		goto L135
	}
L135:
	;
	if base.Ui64(int64(99)) < base.Ui64(v1857) {
		v1916 = v1847
		v1926 = v1857
		v1933 = int32(3)
		goto L9
	} else {
		goto L136
	}
L136:
	;
	if base.Ui64(int64(9)) < base.Ui64(v1857) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v1910 = int32(2)
	goto L139
L138:
	;
	v1910 = int32(1)
	goto L139
L139:
	;
	v1916 = v1847
	v1926 = v1857
	v1933 = v1910
	goto L9
L140:
	;
	v1936 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v1936)
	v1939 = int32(1)
	goto L142
L141:
	;
	v1939 = v3
	goto L142
L142:
	;
	v1940 = v1916 + v1933
	if base.Ui32(v1940+int32(3)) <= base.Ui32(int32(18)) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v1945 = v1939 + l1
	if v1940 <= int32(0) {
		goto L147
	} else {
		goto L148
	}
L144:
	;
	goto L145
L145:
	;
	if v1916 != 0 {
		goto L190
	} else {
		goto L191
	}
L146:
	;
	if base.Ui64(v1926) < base.Ui64(int64(4294967296)) {
		goto L152
	} else {
		goto L153
	}
L147:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1945))) = int64(3472328296227679792)
	v1960 = int32(2) - v1940
	goto L146
L148:
	;
	goto L149
L149:
	;
	if v1916 < int32(0) {
		v1960 = int32(1)
		goto L146
	} else {
		goto L150
	}
L150:
	;
	v1955 = int64(3472328296227680304)
	*(*int64)(unsafe.Add(mBase, uint32(v1945)+8)) = v1955
	*(*int64)(unsafe.Add(mBase, uint32(v1945))) = v1955
	v1960 = int32(0)
	goto L146
L151:
	;
	v2066 = base.I32_wrap_i64(v2065)
	if base.Ui32(int32(_a_F_double_to_shortest_decimal_buf_2)) <= base.Ui32(v2066) {
		goto L156
	} else {
		goto L157
	}
L152:
	;
	v2061 = int32(0)
	v2065 = v1926
	goto L151
L153:
	;
	goto L154
L154:
	;
	v1964 = int64(0)
	v1971 = int64(32)
	v1972 = int64(2882303761)
	v1974 = int64(base.Ui64(v1926) >> (uint(v1971) % 64))
	v1977 = int64(4294967295)
	v1978 = int64(2221002493)
	v1980 = v1926 & v1977
	v1981 = v1978 * v1980
	v1985 = int64(base.Ui64(v1981)>>(uint(v1971)%64)) + v1978*v1974
	v1992 = v1980*v1972 + v1985&v1977
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v1926*v1964 + v1964 + v1972*v1974 + int64(base.Ui64(v1985)>>(uint(v1971)%64)) + int64(base.Ui64(v1992)>>(uint(v1971)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v1981&v1977 | v1992<<(uint(v1971)%64)
	goto L155
L155:
	;
	v2004 = v1945 + v1960 + v1933
	v2005 = int32(8)
	v2007 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
	v2009 = int64(base.Ui64(v2007) >> (uint(int64(26)) % 64))
	v2013 = base.I32_wrap_i64(v2009*int64(4194967296) + v1926)
	v2014 = int32(_a_F_double_to_shortest_decimal_buf_2)
	v2015 = base.I32_div_u_s(v2013, v2014)
	v2017 = base.I32_rem_u_s(v2015, v2014)
	v2018 = int32(100)
	v2019 = base.I32_div_u_s(v2017, v2018)
	v2020 = int32(1)
	v2022 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2019<<(uint(v2020)%32))+uint32(_c_F_double_to_shortest_decimal_buf[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2004-v2005))) = uint16(v2022)
	v2028 = v2013 - v2015*v2014
	v2029 = int32(_a_F_double_to_shortest_decimal_buf_3)
	v2032 = base.I32_div_u_s(v2028&v2029, v2018)
	v2035 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2032<<(uint(v2020)%32))+uint32(_c_F_double_to_shortest_decimal_buf[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2004-int32(4)))) = uint16(v2035)
	v2046 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2017-v2019*v2018)&v2029<<(uint(v2020)%32))+uint32(_c_F_double_to_shortest_decimal_buf[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2004-int32(6)))) = uint16(v2046)
	v2057 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2028-v2032*v2018)&v2029<<(uint(v2020)%32))+uint32(_c_F_double_to_shortest_decimal_buf[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2004-int32(2)))) = uint16(v2057)
	v2061 = v2005
	v2065 = v2009
	goto L151
L156:
	;
	v2072 = v2066
	v2074 = v2061
	goto L159
L157:
	;
	v2121 = v2066
	v2123 = v2061
	goto L158
L158:
	;
	if base.Ui32(v2121) < base.Ui32(int32(100)) {
		goto L163
	} else {
		goto L164
	}
L159:
	;
	v2093 = v1945 + v1960 + v1933 - v2074
	v2094 = int32(4)
	v2097 = base.I32_div_u_s(v2072, int32(_a_F_double_to_shortest_decimal_buf_2))
	v2100 = v2072 + v2097*int32(-10000)
	v2101 = int32(100)
	v2102 = base.I32_div_u_s(v2100, v2101)
	v2103 = int32(1)
	v2105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2102<<(uint(v2103)%32))+uint32(_c_F_double_to_shortest_decimal_buf[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2093-v2094))) = uint16(v2105)
	v2114 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2100-v2102*v2101)<<(uint(v2103)%32))+uint32(_c_F_double_to_shortest_decimal_buf[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2093-int32(2)))) = uint16(v2114)
	v2117 = v2074 + v2094
	if base.Ui32(int32(99999999)) < base.Ui32(v2072) {
		v2072 = v2097
		v2074 = v2117
		goto L159
	} else {
		goto L161
	}
L160:
	;
	v2121 = v2097
	v2123 = v2117
	goto L158
L161:
	;
	goto L160
L162:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v2165) {
		goto L167
	} else {
		goto L168
	}
L163:
	;
	v2164 = v2123
	v2165 = v2121
	goto L162
L164:
	;
	goto L165
L165:
	;
	v2147 = int32(2)
	v2149 = int32(_a_F_double_to_shortest_decimal_buf_3)
	v2151 = int32(100)
	v2152 = base.I32_div_u_s(v2121&v2149, v2151)
	v2160 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2121-v2152*v2151)&v2149<<(uint(int32(1))%32))+uint32(_c_F_double_to_shortest_decimal_buf[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1945+v1960+v1933-v2123-v2147))) = uint16(v2160)
	v2164 = v2123 | v2147
	v2165 = v2152
	goto L162
L166:
	;
	v2181 = int32(1)
	if v1960 == v2181 {
		goto L171
	} else {
		goto L172
	}
L167:
	;
	v2175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2165<<(uint(int32(1))%32))+uint32(_c_F_double_to_shortest_decimal_buf[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1945+v1960+v1933-v2164-int32(2)))) = uint16(v2175)
	goto L166
L168:
	;
	goto L169
L169:
	;
	v2179 = v2165 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1945+v1960))) = uint8(v2179)
	goto L166
L170:
	;
	v2657 = v2232 + base.I32_wrap_i64(int64(base.Ui64(v27)>>(uint(int64(63))%64)))
	goto L1
L171:
	;
	if v1940&int32(8) != 0 {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	goto L173
L173:
	;
	if v1916 < int32(0) {
		goto L186
	} else {
		goto L187
	}
L174:
	;
	v2186 = *(*int64)(unsafe.Add(mBase, uint32(v1945)+1))
	*(*int64)(unsafe.Add(mBase, uint32(v1945))) = v2186
	v2189 = int32(9)
	goto L176
L175:
	;
	v2189 = v2181
	goto L176
L176:
	;
	if v1940&int32(4) != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v2192 = v2189 + v1945
	v2195 = *(*int32)(unsafe.Add(mBase, uint32(v2192)))
	*(*int32)(unsafe.Add(mBase, uint32(v2192-int32(1)))) = v2195
	v2200 = v2189 | int32(4)
	goto L179
L178:
	;
	v2200 = v2189
	goto L179
L179:
	;
	if v1940&int32(2) != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v2203 = v2200 + v1945
	v2206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2203))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2203-int32(1)))) = uint16(v2206)
	v2211 = v2200 + int32(2)
	goto L182
L181:
	;
	v2211 = v2200
	goto L182
L182:
	;
	if v1940&int32(1) != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v2214 = v2211 + v1945
	v2217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2214))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2214-int32(1)))) = uint8(v2217)
	goto L185
L184:
	;
	goto L185
L185:
	;
	v2221 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1945+v1940))) = uint8(v2221)
	v2232 = v1933 + int32(1)
	goto L170
L186:
	;
	v2229 = int32(2) - v1916
	goto L188
L187:
	;
	v2229 = v1940
	goto L188
L188:
	;
	v2232 = v2229
	goto L170
L189:
	;
	if base.Ui64(v2331) < base.Ui64(int64(4294967296)) {
		goto L199
	} else {
		goto L200
	}
L190:
	;
	v2320 = v1933
	v2331 = v1926
	goto L189
L191:
	;
	goto L192
L192:
	;
	v2244 = v1933
	v2252 = v1926
	goto L193
L193:
	;
	if base.I32_wrap_i64(v2252)&int32(1) != 0 {
		v2320 = v2244
		v2331 = v2252
		goto L189
	} else {
		goto L195
	}
L194:
	;
	v2320 = v2244
	v2331 = v2252
	goto L189
L195:
	;
	v2263 = v25 + int32(32)
	v2264 = int64(0)
	v2271 = int64(32)
	v2272 = int64(3435973836)
	v2274 = int64(base.Ui64(v2252) >> (uint(v2271) % 64))
	v2277 = int64(4294967295)
	v2278 = int64(3435973837)
	v2280 = v2252 & v2277
	v2281 = v2278 * v2280
	v2285 = int64(base.Ui64(v2281)>>(uint(v2271)%64)) + v2278*v2274
	v2292 = v2280*v2272 + v2285&v2277
	*(*int64)(unsafe.Add(mBase, uint32(v2263)+8)) = v2252*v2264 + v2264 + v2272*v2274 + int64(base.Ui64(v2285)>>(uint(v2271)%64)) + int64(base.Ui64(v2292)>>(uint(v2271)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v2263))) = v2281&v2277 | v2292<<(uint(v2271)%64)
	goto L196
L196:
	;
	v2305 = *(*int64)(unsafe.Add(mBase, uint32(v25)+40))
	v2307 = int64(base.Ui64(v2305) >> (uint(int64(3)) % 64))
	if (v2252+v2307*int64(4294967286))&int64(4294967294) == int64(0) {
		v2244 = v2244 - int32(1)
		v2252 = v2307
		goto L193
	} else {
		goto L197
	}
L197:
	;
	goto L194
L198:
	;
	v2444 = base.I32_wrap_i64(v2443)
	if base.Ui32(int32(_a_F_double_to_shortest_decimal_buf_2)) <= base.Ui32(v2444) {
		goto L203
	} else {
		goto L204
	}
L199:
	;
	v2439 = int32(0)
	v2443 = v2331
	goto L198
L200:
	;
	goto L201
L201:
	;
	v2341 = v25 + int32(16)
	v2342 = int64(0)
	v2349 = int64(32)
	v2350 = int64(2882303761)
	v2352 = int64(base.Ui64(v2331) >> (uint(v2349) % 64))
	v2355 = int64(4294967295)
	v2356 = int64(2221002493)
	v2358 = v2331 & v2355
	v2359 = v2356 * v2358
	v2363 = int64(base.Ui64(v2359)>>(uint(v2349)%64)) + v2356*v2352
	v2370 = v2358*v2350 + v2363&v2355
	*(*int64)(unsafe.Add(mBase, uint32(v2341)+8)) = v2331*v2342 + v2342 + v2350*v2352 + int64(base.Ui64(v2363)>>(uint(v2349)%64)) + int64(base.Ui64(v2370)>>(uint(v2349)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v2341))) = v2359&v2355 | v2370<<(uint(v2349)%64)
	goto L202
L202:
	;
	v2382 = v1939 + l1 + v2320
	v2385 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
	v2387 = int64(base.Ui64(v2385) >> (uint(int64(26)) % 64))
	v2391 = base.I32_wrap_i64(v2387*int64(4194967296) + v2331)
	v2392 = int32(_a_F_double_to_shortest_decimal_buf_2)
	v2393 = base.I32_div_u_s(v2391, v2392)
	v2395 = base.I32_rem_u_s(v2393, v2392)
	v2396 = int32(100)
	v2397 = base.I32_div_u_s(v2395, v2396)
	v2398 = int32(1)
	v2400 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2397<<(uint(v2398)%32))+uint32(_c_F_double_to_shortest_decimal_buf[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2382-int32(7)))) = uint16(v2400)
	v2406 = v2391 - v2393*v2392
	v2407 = int32(_a_F_double_to_shortest_decimal_buf_3)
	v2410 = base.I32_div_u_s(v2406&v2407, v2396)
	v2413 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2410<<(uint(v2398)%32))+uint32(_c_F_double_to_shortest_decimal_buf[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2382-int32(3)))) = uint16(v2413)
	v2424 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2395-v2397*v2396)&v2407<<(uint(v2398)%32))+uint32(_c_F_double_to_shortest_decimal_buf[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2382-int32(5)))) = uint16(v2424)
	v2435 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2406-v2410*v2396)&v2407<<(uint(v2398)%32))+uint32(_c_F_double_to_shortest_decimal_buf[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2382-v2398))) = uint16(v2435)
	v2439 = int32(8)
	v2443 = v2387
	goto L198
L203:
	;
	v2450 = v2444
	v2452 = v2439
	goto L206
L204:
	;
	v2499 = v2444
	v2501 = v2439
	goto L205
L205:
	;
	if base.Ui32(v2499) < base.Ui32(int32(100)) {
		goto L210
	} else {
		goto L211
	}
L206:
	;
	v2471 = v1939 + l1 + v2320 - v2452
	v2475 = base.I32_div_u_s(v2450, int32(_a_F_double_to_shortest_decimal_buf_2))
	v2478 = v2450 + v2475*int32(-10000)
	v2479 = int32(100)
	v2480 = base.I32_div_u_s(v2478, v2479)
	v2481 = int32(1)
	v2483 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2480<<(uint(v2481)%32))+uint32(_c_F_double_to_shortest_decimal_buf[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2471-int32(3)))) = uint16(v2483)
	v2492 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2478-v2480*v2479)<<(uint(v2481)%32))+uint32(_c_F_double_to_shortest_decimal_buf[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2471-v2481))) = uint16(v2492)
	v2495 = v2452 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v2450) {
		v2450 = v2475
		v2452 = v2495
		goto L206
	} else {
		goto L208
	}
L207:
	;
	v2499 = v2475
	v2501 = v2495
	goto L205
L208:
	;
	goto L207
L209:
	;
	v2545 = v1940 - int32(1)
	v2546 = v1939 + l1
	if base.Ui32(int32(10)) <= base.Ui32(v2543) {
		goto L214
	} else {
		goto L215
	}
L210:
	;
	v2542 = v2501
	v2543 = v2499
	goto L209
L211:
	;
	goto L212
L212:
	;
	v2527 = int32(_a_F_double_to_shortest_decimal_buf_3)
	v2529 = int32(100)
	v2530 = base.I32_div_u_s(v2499&v2527, v2529)
	v2538 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2499-v2530*v2529)&v2527<<(uint(int32(1))%32))+uint32(_c_F_double_to_shortest_decimal_buf[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1939+l1+v2320+(v2501^int32(-1))))) = uint16(v2538)
	v2542 = v2501 | int32(2)
	v2543 = v2530
	goto L209
L213:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2546))) = uint8(v2560)
	if base.Ui32(int32(2)) <= base.Ui32(v2320) {
		goto L217
	} else {
		goto L218
	}
L214:
	;
	v2553 = v2543 << (uint(int32(1)) % 32)
	v2554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2553)+uint32(_c_F_double_to_shortest_decimal_buf[7]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1+(v2320+v1939-v2542)))) = uint8(v2554)
	v2556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2553)+uint32(_c_F_double_to_shortest_decimal_buf[6]))))
	v2560 = v2556
	goto L213
L215:
	;
	goto L216
L216:
	;
	v2560 = v2543 | int32(48)
	goto L213
L217:
	;
	v2564 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v2546)+1)) = uint8(v2564)
	v2569 = v2320 + int32(1)
	goto L219
L218:
	;
	v2569 = int32(1)
	goto L219
L219:
	;
	v2570 = v2569 + v1939
	v2571 = l1 + v2570
	v2572 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v2571))) = uint8(v2572)
	v2577 = base.B2i32(v2545 < int32(0))
	if v2545 < int32(0) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v2578 = int32(45)
	goto L222
L221:
	;
	v2578 = int32(43)
	goto L222
L222:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2571)+1)) = uint8(v2578)
	v2581 = v2570 + int32(2)
	if v2545 < int32(0) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v2584 = int32(1) - v1940
	goto L225
L224:
	;
	v2584 = v2545
	goto L225
L225:
	;
	if int32(100) <= v2584 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v2588 = int32(10)
	v2589 = base.I32_div_u_s(v2584, v2588)
	v2592 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2589<<(uint(int32(1))%32))+uint32(_c_F_double_to_shortest_decimal_buf[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2581+l1))) = uint16(v2592)
	v2598 = v2584 - v2589*v2588 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2571)+4)) = uint8(v2598)
	v2657 = v2570 + int32(5)
	goto L1
L227:
	;
	goto L228
L228:
	;
	v2607 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2584<<(uint(int32(1))%32))+uint32(_c_F_double_to_shortest_decimal_buf[6]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2581+l1))) = uint16(v2607)
	v2657 = v2570 + int32(4)
	goto L1
L229:
	;
	v2613 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v2613)
	goto L231
L230:
	;
	goto L231
L231:
	;
	v2618 = l1 + base.I32_wrap_i64(int64(base.Ui64(v27)>>(uint(int64(63))%64)))
	if v34 == int32(2047) {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2618))) = int64(8751735898823355977)
	if v27 < int64(0) {
		goto L235
	} else {
		goto L236
	}
L233:
	;
	goto L234
L234:
	;
	v2628 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2618))) = uint8(v2628)
	if v27 < int64(0) {
		goto L238
	} else {
		goto L239
	}
L235:
	;
	v2627 = int32(9)
	goto L237
L236:
	;
	v2627 = int32(8)
	goto L237
L237:
	;
	v2657 = v2627
	goto L1
L238:
	;
	v2634 = int32(2)
	goto L240
L239:
	;
	v2634 = int32(1)
	goto L240
L240:
	;
	v2657 = v2634
	goto L1
}
func F_dround(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)))
	v5 = F_Float8GetDatum(m, base.F64_nearest(v3))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_dsin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 float64
	_ = v4
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v35 float64
	_ = v35
	var v39 int32
	_ = v39
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v47 float64
	_ = v47
	var v48 float64
	_ = v48
	var v50 float64
	_ = v50
	var v52 float64
	_ = v52
	var v54 float64
	_ = v54
	var v61 float64
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(v3)))
	if base.Ui64(base.I64_reinterpret_f64(v4)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		*(*int32)(unsafe.Add(mBase, _c_F_dsin[0])) = int32(0)
		if base.F64_eq(base.F64_abs(v4), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_dsin_0), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_dsin_1), int32(1953), int32(_a_F_dsin_2))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v19 = m.G0
			v21 = v19 - int32(16)
			m.G0 = v21
			v28 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v4))>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(v28) <= base.Ui32(int32(1072243195)) {
				if base.Ui32(v28) < base.Ui32(int32(1045430272)) {
					v54 = v4
				} else {
					v35 = F___sin(m, v4, float64(0), int32(0))
					mBase = m.M
					v54 = v35
				}
			} else {
				if base.Ui32(int32(2146435072)) <= base.Ui32(v28) {
					v54 = base.F64_sub(v4, v4)
				} else {
					v39 = F___rem_pio2(m, v4, v21)
					mBase = m.M
					v40 = *(*float64)(unsafe.Add(mBase, uint32(v21)+8))
					v41 = *(*float64)(unsafe.Add(mBase, uint32(v21)))
					switch v39&int32(3) - int32(1) {
					case 0:
						v48 = F___cos(m, v41, v40)
						mBase = m.M
						v54 = v48
					case 1:
						v50 = F___sin(m, v41, v40, int32(1))
						mBase = m.M
						v54 = base.F64_neg(v50)
					case 2:
						v52 = F___cos(m, v41, v40)
						mBase = m.M
						v54 = base.F64_neg(v52)
					default:
						v47 = F___sin(m, v41, v40, int32(1))
						mBase = m.M
						v54 = v47
					}
				}
			}
			m.G0 = v21 + int32(16)
			v61 = v54
			v62 = F_Float8GetDatum(m, v61)
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				return v62
			}
		}
	} else {
		v61 = math.Float64frombits(uint64(0x7ff8000000000000))
		v62 = F_Float8GetDatum(m, v61)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			return v62
		}
	}
}
func F_dsinh(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v14 int64
	_ = v14
	var v23 int64
	_ = v23
	var v28 int32
	_ = v28
	var v64 int32
	_ = v64
	var v65 float64
	_ = v65
	var v71 int32
	_ = v71
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v81 float64
	_ = v81
	var v82 int32
	_ = v82
	var v83 float64
	_ = v83
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v103 float64
	_ = v103
	var v106 float64
	_ = v106
	var v112 float64
	_ = v112
	var v121 float64
	_ = v121
	var v136 float64
	_ = v136
	var v145 float64
	_ = v145
	var v150 float64
	_ = v150
	var v157 float64
	_ = v157
	var v160 float64
	_ = v160
	var v166 float64
	_ = v166
	var v176 float64
	_ = v176
	var v178 float64
	_ = v178
	var v190 float64
	_ = v190
	var v208 float64
	_ = v208
	var v212 float64
	_ = v212
	var v217 float64
	_ = v217
	var v220 float64
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	*(*int32)(unsafe.Add(mBase, _c_F_dsinh[0])) = int32(0)
	v12 = base.F64_copysign(float64(0.5), v7)
	v13 = base.F64_abs(v7)
	v14 = base.I64_reinterpret_f64(v13)
	if base.Ui64(v14) <= base.Ui64(int64(4649454526309335039)) {
		v23 = base.I64_reinterpret_f64(v13)
		v28 = base.I32_wrap_i64(int64(base.Ui64(v23)>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(int32(1078159482)) <= base.Ui32(v28) {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v13)&int64(9223372036854775807)) {
				v178 = v13
				v190 = v178
			} else {
				if v23 < int64(0) {
					v190 = float64(-1)
				} else {
					if base.F64_gt(v13, float64(709.782712893384)) == int32(0) {
						v64 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v13, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v13)))
						v65 = base.F64_convert_i32_s(v64)
						v71 = v64
						v72 = base.F64_mul(v65, float64(1.9082149292705877e-10))
						v74 = base.F64_add(v13, base.F64_mul(v65, float64(-0.6931471803691238)))
						v75 = base.F64_sub(v74, v72)
						v81 = v75
						v82 = v71
						v83 = base.F64_sub(base.F64_sub(v74, v75), v72)
						v86 = base.F64_mul(v81, float64(0.5))
						v87 = base.F64_mul(v81, v86)
						v103 = base.F64_add(base.F64_mul(v87, base.F64_add(base.F64_mul(v87, base.F64_add(base.F64_mul(v87, base.F64_add(base.F64_mul(v87, base.F64_add(base.F64_mul(v87, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
						v106 = base.F64_sub(float64(3), base.F64_mul(v103, v86))
						v112 = base.F64_mul(v87, base.F64_div(base.F64_sub(v103, v106), base.F64_sub(float64(6), base.F64_mul(v81, v106))))
						if v82 == int32(0) {
							v190 = base.F64_sub(v81, base.F64_sub(base.F64_mul(v81, v112), v87))
						} else {
							v121 = base.F64_sub(base.F64_sub(base.F64_mul(v81, base.F64_sub(v112, v83)), v83), v87)
							switch v82 + int32(1) {
							case 0:
								v190 = base.F64_add(base.F64_mul(base.F64_sub(v81, v121), float64(0.5)), float64(-0.5))
							default:
								v145 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v82+int32(1023)) << (uint(int64(52)) % 64))
								if base.Ui32(int32(57)) <= base.Ui32(v82) {
									v150 = base.F64_add(base.F64_sub(v81, v121), float64(1))
									if v82 == int32(1024) {
										v157 = base.F64_mul(base.F64_add(v150, v150), float64(8.98846567431158e+307))
									} else {
										v157 = base.F64_mul(v150, v145)
									}
									v190 = base.F64_add(v157, float64(-1))
								} else {
									v160 = float64(1)
									v166 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v82) << (uint(int64(52)) % 64))
									if base.Ui32(v82) <= base.Ui32(int32(19)) {
										v176 = base.F64_add(base.F64_sub(v160, v166), base.F64_sub(v81, v121))
									} else {
										v176 = base.F64_add(base.F64_sub(v81, base.F64_add(v121, v166)), v160)
									}
									v178 = base.F64_mul(v176, v145)
									v190 = v178
								}
							case 2:
								if base.F64_lt(v81, float64(-0.25)) != 0 {
									v190 = base.F64_mul(base.F64_sub(v121, base.F64_add(v81, float64(0.5))), float64(-2))
								} else {
									v136 = base.F64_sub(v81, v121)
									v190 = base.F64_add(base.F64_add(v136, v136), float64(1))
								}
							}
						}
					} else {
						v190 = base.F64_mul(v13, float64(8.98846567431158e+307))
					}
				}
			}
		} else {
			if base.Ui32(v28) < base.Ui32(int32(1071001155)) {
				if base.Ui32(v28) < base.Ui32(int32(1016070144)) {
					v178 = v13
					v190 = v178
				} else {
					v81 = v13
					v82 = int32(0)
					v83 = float64(0)
					v86 = base.F64_mul(v81, float64(0.5))
					v87 = base.F64_mul(v81, v86)
					v103 = base.F64_add(base.F64_mul(v87, base.F64_add(base.F64_mul(v87, base.F64_add(base.F64_mul(v87, base.F64_add(base.F64_mul(v87, base.F64_add(base.F64_mul(v87, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
					v106 = base.F64_sub(float64(3), base.F64_mul(v103, v86))
					v112 = base.F64_mul(v87, base.F64_div(base.F64_sub(v103, v106), base.F64_sub(float64(6), base.F64_mul(v81, v106))))
					if v82 == int32(0) {
						v190 = base.F64_sub(v81, base.F64_sub(base.F64_mul(v81, v112), v87))
					} else {
						v121 = base.F64_sub(base.F64_sub(base.F64_mul(v81, base.F64_sub(v112, v83)), v83), v87)
						switch v82 + int32(1) {
						case 0:
							v190 = base.F64_add(base.F64_mul(base.F64_sub(v81, v121), float64(0.5)), float64(-0.5))
						default:
							v145 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v82+int32(1023)) << (uint(int64(52)) % 64))
							if base.Ui32(int32(57)) <= base.Ui32(v82) {
								v150 = base.F64_add(base.F64_sub(v81, v121), float64(1))
								if v82 == int32(1024) {
									v157 = base.F64_mul(base.F64_add(v150, v150), float64(8.98846567431158e+307))
								} else {
									v157 = base.F64_mul(v150, v145)
								}
								v190 = base.F64_add(v157, float64(-1))
							} else {
								v160 = float64(1)
								v166 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v82) << (uint(int64(52)) % 64))
								if base.Ui32(v82) <= base.Ui32(int32(19)) {
									v176 = base.F64_add(base.F64_sub(v160, v166), base.F64_sub(v81, v121))
								} else {
									v176 = base.F64_add(base.F64_sub(v81, base.F64_add(v121, v166)), v160)
								}
								v178 = base.F64_mul(v176, v145)
								v190 = v178
							}
						case 2:
							if base.F64_lt(v81, float64(-0.25)) != 0 {
								v190 = base.F64_mul(base.F64_sub(v121, base.F64_add(v81, float64(0.5))), float64(-2))
							} else {
								v136 = base.F64_sub(v81, v121)
								v190 = base.F64_add(base.F64_add(v136, v136), float64(1))
							}
						}
					}
				}
			} else {
				if base.Ui32(int32(1072734897)) < base.Ui32(v28) {
					v64 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v13, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v13)))
					v65 = base.F64_convert_i32_s(v64)
					v71 = v64
					v72 = base.F64_mul(v65, float64(1.9082149292705877e-10))
					v74 = base.F64_add(v13, base.F64_mul(v65, float64(-0.6931471803691238)))
				} else {
					if int64(0) <= v23 {
						v71 = int32(1)
						v72 = float64(1.9082149292705877e-10)
						v74 = base.F64_add(v13, float64(-0.6931471803691238))
					} else {
						v71 = int32(-1)
						v72 = float64(-1.9082149292705877e-10)
						v74 = base.F64_add(v13, float64(0.6931471803691238))
					}
				}
				v75 = base.F64_sub(v74, v72)
				v81 = v75
				v82 = v71
				v83 = base.F64_sub(base.F64_sub(v74, v75), v72)
				v86 = base.F64_mul(v81, float64(0.5))
				v87 = base.F64_mul(v81, v86)
				v103 = base.F64_add(base.F64_mul(v87, base.F64_add(base.F64_mul(v87, base.F64_add(base.F64_mul(v87, base.F64_add(base.F64_mul(v87, base.F64_add(base.F64_mul(v87, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
				v106 = base.F64_sub(float64(3), base.F64_mul(v103, v86))
				v112 = base.F64_mul(v87, base.F64_div(base.F64_sub(v103, v106), base.F64_sub(float64(6), base.F64_mul(v81, v106))))
				if v82 == int32(0) {
					v190 = base.F64_sub(v81, base.F64_sub(base.F64_mul(v81, v112), v87))
				} else {
					v121 = base.F64_sub(base.F64_sub(base.F64_mul(v81, base.F64_sub(v112, v83)), v83), v87)
					switch v82 + int32(1) {
					case 0:
						v190 = base.F64_add(base.F64_mul(base.F64_sub(v81, v121), float64(0.5)), float64(-0.5))
					default:
						v145 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v82+int32(1023)) << (uint(int64(52)) % 64))
						if base.Ui32(int32(57)) <= base.Ui32(v82) {
							v150 = base.F64_add(base.F64_sub(v81, v121), float64(1))
							if v82 == int32(1024) {
								v157 = base.F64_mul(base.F64_add(v150, v150), float64(8.98846567431158e+307))
							} else {
								v157 = base.F64_mul(v150, v145)
							}
							v190 = base.F64_add(v157, float64(-1))
						} else {
							v160 = float64(1)
							v166 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v82) << (uint(int64(52)) % 64))
							if base.Ui32(v82) <= base.Ui32(int32(19)) {
								v176 = base.F64_add(base.F64_sub(v160, v166), base.F64_sub(v81, v121))
							} else {
								v176 = base.F64_add(base.F64_sub(v81, base.F64_add(v121, v166)), v160)
							}
							v178 = base.F64_mul(v176, v145)
							v190 = v178
						}
					case 2:
						if base.F64_lt(v81, float64(-0.25)) != 0 {
							v190 = base.F64_mul(base.F64_sub(v121, base.F64_add(v81, float64(0.5))), float64(-2))
						} else {
							v136 = base.F64_sub(v81, v121)
							v190 = base.F64_add(base.F64_add(v136, v136), float64(1))
						}
					}
				}
			}
		}
		if base.Ui64(v14) <= base.Ui64(int64(4607182418800017407)) {
			if base.Ui64(v14) < base.Ui64(int64(4490088828488384512)) {
				v217 = v7
				v220 = v217
			} else {
				v220 = base.F64_mul(v12, base.F64_sub(base.F64_add(v190, v190), base.F64_div(base.F64_mul(v190, v190), base.F64_add(v190, float64(1)))))
			}
		} else {
			v220 = base.F64_mul(v12, base.F64_add(v190, base.F64_div(v190, base.F64_add(v190, float64(1)))))
		}
	} else {
		v208 = float64(2.247116418577895e+307)
		v212 = F_exp(m, base.F64_add(v13, float64(-1416.0996898839683)))
		mBase = m.M
		v217 = base.F64_mul(base.F64_mul(base.F64_mul(base.F64_add(v12, v12), v208), v212), v208)
		v220 = v217
	}
	v221 = F_Float8GetDatum(m, v220)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		return int32(0)
	} else {
		return v221
	}
}
func F_dtoi2(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13868(m, l0, int32(_a_F_dtoi2_0), int32(1254), int32(_a_F_dtoi2_1), float64(32768), float64(-32768))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_dtoi4(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13868(m, l0, int32(_a_F_dtoi4_0), int32(1229), int32(_a_F_dtoi4_1), float64(2.147483648e+09), float64(-2.147483648e+09))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_dtoi8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 float64
	_ = v4
	var v5 float64
	_ = v5
	var v13 int32
	_ = v13
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(v3)))
	v5 = base.F64_nearest(v4)
	v13 = int32(0)
	if base.B2i32(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)))|base.B2i32(base.F64_ge(v5, float64(-9.223372036854776e+18)) == v13) == v13)&base.F64_lt(v5, float64(9.223372036854776e+18)) == v13 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_dtoi8_0), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_dtoi8_1), int32(1312), int32(_a_F_dtoi8_2))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v42 = F_Int64GetDatum(m, base.I64_trunc_sat_f64_s(v5))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			return v42
		}
	}
}
func F_durable_rename(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(1088)
	m.G0 = v9
	v14 = F_fsync_fname_ext(m, l0, v4, v4, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(1088)
	return v341
L2:
	;
	return int32(0)
L3:
	;
	if v14 != 0 {
		v341 = int32(-1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_durable_rename[0]))
	v21 = F_OpenTransientFilePerm(m, l1, int32(2), v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v127 = int32(-1)
	v128 = F_rename(m, l0, l1)
	mBase = m.M
	if v128 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L6:
	;
	if v21 < int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_durable_rename[1]))
	if v26 == int32(44) {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_durable_rename[2])))
	if v49 != int32(1) {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	v29 = int32(-1)
	v31 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	if v31 == int32(0) {
		v341 = v29
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
	F_errmsg(m, int32(_a_F_durable_rename_0), v9+int32(16))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_durable_rename_1), int32(808), int32(_a_F_durable_rename_2))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v341 = v29
	goto L1
L16:
	;
	v98 = F_CloseTransientFile(m, v21)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L32
	}
L17:
	;
	goto L18
L18:
	;
	v58 = F_fsync(m, v21)
	mBase = m.M
	if v58 != int32(-1) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_durable_rename[1]))
	v69 = F_CloseTransientFile(m, v21)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	if v58 == int32(0) {
		goto L16
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_durable_rename[1]))
	if v64 == int32(27) {
		goto L18
	} else {
		goto L25
	}
L24:
	;
	goto L20
L25:
	;
	goto L20
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_durable_rename[1])) = v68
	v73 = int32(-1)
	v75 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	if v75 == int32(0) {
		v341 = v73
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
	F_errmsg(m, int32(_a_F_durable_rename_3), v9+int32(32))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_durable_rename_1), int32(825), int32(_a_F_durable_rename_2))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v341 = v73
	goto L1
L32:
	;
	if v98 == int32(0) {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v102 = int32(-1)
	v104 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	if v104 == int32(0) {
		v341 = v102
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l1
	F_errmsg(m, int32(_a_F_durable_rename_4), v9+int32(48))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_durable_rename_1), int32(833), int32(_a_F_durable_rename_2))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v341 = v102
	goto L1
L39:
	;
	v132 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L2
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v148 = int32(0)
	v150 = F_fsync_fname_ext(m, l1, v148, v148, l2)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L47
	}
L42:
	;
	if v132 == int32(0) {
		v341 = v127
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	F_errmsg(m, int32(_a_F_durable_rename_5), v9)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_durable_rename_1), int32(844), int32(_a_F_durable_rename_2))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v341 = v127
	goto L1
L47:
	;
	if v150 != 0 {
		v341 = v127
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v153 = v9 - int32(-64)
	goto L52
L49:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v276 != 0 {
		goto L81
	} else {
		goto L82
	}
L50:
	;
	v270 = F_strlen(m, v259)
	mBase = m.M
	goto L49
L52:
	;
	goto L53
L53:
	;
	v160 = int32(1023)
	if (v153^l1)&int32(3) != 0 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v263 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v260))) = uint8(v263)
	goto L50
L55:
	;
	v244 = v239
	v245 = v240
	v246 = v241
	goto L76
L56:
	;
	if v234 == int32(0) {
		v259 = v232
		v260 = v233
		goto L54
	} else {
		goto L75
	}
L57:
	;
	v232 = l1
	v233 = v153
	v234 = v160
	goto L56
L58:
	;
	goto L59
L59:
	;
	v164 = int32(0)
	if base.B2i32(l1&int32(3) == v164)|int32(0) == v164 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v200 == int32(0) {
		v259 = v197
		v260 = v198
		goto L54
	} else {
		goto L69
	}
L61:
	;
	v176 = l1
	v177 = v153
	v178 = v160
	goto L64
L62:
	;
	goto L63
L63:
	;
	v197 = l1
	v198 = v153
	v199 = v160
	v200 = int32(1)
	goto L60
L64:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	*(*uint8)(unsafe.Add(mBase, uint32(v177))) = uint8(v180)
	if v180 == int32(0) {
		v239 = v176
		v240 = v177
		v241 = v178
		goto L55
	} else {
		goto L66
	}
L65:
	;
	v197 = v191
	v198 = v185
	v199 = v187
	v200 = v189
	goto L60
L66:
	;
	v184 = int32(1)
	v185 = v177 + v184
	v187 = v178 - v184
	v188 = int32(0)
	v189 = base.B2i32(v187 != v188)
	v191 = v176 + v184
	if v191&int32(3) == v188 {
		v197 = v191
		v198 = v185
		v199 = v187
		v200 = v189
		goto L60
	} else {
		goto L67
	}
L67:
	;
	if v187 != 0 {
		v176 = v191
		v177 = v185
		v178 = v187
		goto L64
	} else {
		goto L68
	}
L68:
	;
	goto L65
L69:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	if base.B2i32(v203 == int32(0))|base.B2i32(base.Ui32(v199) < base.Ui32(int32(4))) != 0 {
		v232 = v197
		v233 = v198
		v234 = v199
		goto L56
	} else {
		goto L70
	}
L70:
	;
	v210 = v197
	v211 = v198
	v212 = v199
	goto L71
L71:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v218 = int32(-2139062144)
	if (int32(16843008)-v215|v215)&v218 != v218 {
		v239 = v210
		v240 = v211
		v241 = v212
		goto L55
	} else {
		goto L73
	}
L72:
	;
	v232 = v226
	v233 = v224
	v234 = v228
	goto L56
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211))) = v215
	v223 = int32(4)
	v224 = v211 + v223
	v226 = v210 + v223
	v228 = v212 - v223
	if base.Ui32(int32(3)) < base.Ui32(v228) {
		v210 = v226
		v211 = v224
		v212 = v228
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v239 = v232
	v240 = v233
	v241 = v234
	goto L55
L76:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	*(*uint8)(unsafe.Add(mBase, uint32(v245))) = uint8(v248)
	if v248 == int32(0) {
		v259 = v244
		v260 = v245
		goto L54
	} else {
		goto L78
	}
L77:
	;
	v259 = v255
	v260 = v253
	goto L54
L78:
	;
	v252 = int32(1)
	v253 = v245 + v252
	v255 = v244 + v252
	v257 = v246 - v252
	if v257 != 0 {
		v244 = v255
		v245 = v253
		v246 = v257
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+64)))
	if v323 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L81:
	;
	v277 = F_strlen(m, v153)
	mBase = m.M
	v280 = v277 + v153
	goto L84
L82:
	;
	goto L83
L83:
	;
	goto L80
L84:
	;
	v284 = v280 - int32(1)
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	if base.B2i32(v285 == int32(47))&base.B2i32(base.Ui32(v153) < base.Ui32(v284)) != 0 {
		v280 = v284
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v291 = v284
	goto L87
L86:
	;
	goto L85
L87:
	;
	if base.Ui32(v153) < base.Ui32(v291) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v303 = v291
	goto L93
L89:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	if v297 != int32(47) {
		v291 = v291 - int32(1)
		goto L87
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	goto L88
L92:
	;
	goto L91
L93:
	;
	if base.Ui32(v153) < base.Ui32(v303) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	if v153 == v303 {
		goto L99
	} else {
		goto L100
	}
L95:
	;
	v307 = v303 - int32(1)
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
	if v308 == int32(47) {
		v303 = v307
		goto L93
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	goto L94
L98:
	;
	goto L97
L99:
	;
	v316 = v153 + base.B2i32(v276 == int32(47))
	goto L101
L100:
	;
	v316 = v303
	goto L101
L101:
	;
	v317 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v316))) = uint8(v317)
	goto L83
L102:
	;
	v326 = int32(46)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+64)) = uint16(v326)
	goto L104
L103:
	;
	goto L104
L104:
	;
	v329 = int32(0)
	v334 = F_fsync_fname_ext(m, v9-int32(-64), int32(1), v329, l2)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	if v334 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v336 = int32(-1)
	goto L108
L107:
	;
	v336 = v329
	goto L108
L108:
	;
	v341 = v336
	goto L1
}
func F_dutch_ISO_8859_1_create_env(m *base.Module) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_SN_create_env(m, int32(0), int32(4))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
