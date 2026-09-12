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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[65]))
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
				F_errmsg(m, int32(270861), int32(0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					F_errfinish(m, int32(512898), int32(4387), int32(94247))
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
					F_errmsg(m, int32(270861), int32(0))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						F_errfinish(m, int32(512898), int32(4387), int32(94247))
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
			v14 = *(*int32)(unsafe.Add(mBase, _consts[50]))
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
						F_errmsg(m, int32(270861), int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							F_errfinish(m, int32(512898), int32(4387), int32(94247))
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
					if v21&int32(4104) == int32(0) {
						if v21&int32(1044455) != 0 {
							F_errstart_cold(m, int32(22), int32(0))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
								if base.Ui32(v80) <= base.Ui32(int32(19)) {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v80<<(uint(int32(2))%32))+uint32(_consts[164])))
									v90 = v89
								} else {
									v90 = int32(565099)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v90
								F_errmsg_internal(m, int32(196045), v7+int32(16))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return
								} else {
									F_errfinish(m, int32(512898), int32(4446), int32(94247))
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
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
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(539303)
									F_errmsg(m, int32(162212), v7)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return
									} else {
										F_errfinish(m, int32(512898), int32(4424), int32(94247))
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
								v50 = *(*int32)(unsafe.Add(mBase, _consts[65]))
								v52 = *(*int32)(unsafe.Add(mBase, _consts[72]))
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
		F_appendStringInfo(m, l0, int32(54049), v9+int32(16))
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
		F_appendStringInfo(m, l0, int32(53992), v9+int32(32))
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
		F_appendStringInfo(m, l0, int32(54076), v9+int32(48))
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
		F_appendStringInfo(m, l0, int32(53954), v9-int32(-64))
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
		F_appendStringInfo(m, l0, int32(54032), v9+int32(80))
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
		F_appendStringInfo(m, l0, int32(48990), v9+int32(96))
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
		F_appendStringInfo(m, l0, int32(41491), v9+int32(112))
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
		F_appendStringInfo(m, l0, int32(48896), v9+int32(128))
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
		F_appendStringInfo(m, l0, int32(53861), v9+int32(144))
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
		F_appendStringInfo(m, l0, int32(528840), v9+int32(160))
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
		F_appendStringInfo(m, l0, int32(528861), v9+int32(176))
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
		F_appendStringInfo(m, l0, int32(53898), v9+int32(192))
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
		F_appendStringInfo(m, l0, int32(494596), v9)
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
	var v122 int32
	_ = v122
	var v124 int64
	_ = v124
	var v137 int64
	_ = v137
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v169 int32
	_ = v169
	var v181 int32
	_ = v181
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v15 <= int32(-4713) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v181
L2:
	;
	v169 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v169
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
	v181 = v169
	goto L1
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v35 = int32(60)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v46 = base.B2i32(int32(2) < v31)
	if int32(2) < v31 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	if v15 != int32(-4713) {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if v15 <= int32(5874897) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(10) < v20 {
		v31 = v20
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L2
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v31 = v25
	goto L3
L10:
	;
	goto L11
L11:
	;
	if v15 != int32(5874898) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(5) < v28 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v31 = v28
	goto L3
L14:
	;
	v47 = int32(4800)
	goto L16
L15:
	;
	v47 = int32(4799)
	goto L16
L16:
	;
	v48 = v47 + v15
	v56 = base.I32_div_u_s(v48, int32(100))
	v59 = base.I32_div_u_s(v48, int32(400))
	if int32(2) < v31 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v63 = int32(1)
	goto L19
L18:
	;
	v63 = int32(13)
	goto L19
L19:
	;
	v68 = base.I32_div_s((v63+v31)*int32(7834), int32(256))
	v71 = v42 + v48*int32(365) + int32(base.Ui32(v48)>>(uint(int32(2))%32)) - v56 + v59 + v68 - int32(2472755)
	v75 = base.I64_extend_i32_s(v32+(v33+v34*v35)*v35) + base.I64_extend_i32_s(v71)*int64(86400)
	if base.B2i32(int32(0) < v71)&base.B2i32(v75 < int64(0)) != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v75 - int64(86400)
	v94 = F_pg_next_dst_boundary(m, v13+int32(24), v13+int32(12), v13+int32(4), v13+int32(16), v13+int32(8), v13, l1)
	mBase = m.M
	if v94 < int32(0) {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	if v94 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v75 - base.I64_extend_i32_s(v101)
	v181 = int32(0) - v101
	goto L1
L23:
	;
	goto L24
L24:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v109 = v75 - base.I64_extend_i32_s(v107)
	if int64(0) <= v75 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if v75 <= int64(0) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	if v107 <= int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	if int64(0) < v109 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v124 = v75 - base.I64_extend_i32_s(v122)
	if int64(0) <= v75 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	if int32(0) < v107 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	if v109 < int64(0) {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	if v75 <= int64(0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	if v122 <= int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	if int64(0) < v124 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
	if v137 <= v109 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	if int32(0) < v122 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	if v124 < int64(0) {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	if v109 <= v137 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	if v137 <= v124 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v109
	v181 = int32(0) - v107
	goto L1
L44:
	;
	if v107 < v122 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	if v124 < v137 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v147
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v124
	v181 = int32(0) - v122
	goto L1
L47:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v153
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v109
	v181 = int32(0) - v107
	goto L1
L48:
	;
	goto L49
L49:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v158
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v124
	v181 = int32(0) - v122
	goto L1
}
func F_DoLockModesConflict(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[74])))
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
	var v18 int64
	_ = v18
	var v23 int32
	_ = v23
	var v36 float64
	_ = v36
	var v47 float64
	_ = v47
	var v59 float64
	_ = v59
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v66 float64
	_ = v66
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v78 float64
	_ = v78
	var v84 float64
	_ = v84
	var v89 float64
	_ = v89
	var v93 float64
	_ = v93
	var v97 float64
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v8 = base.F64_abs(v7)
	if base.Ui64(base.I64_reinterpret_f64(v8)) <= base.Ui64(int64(9218868437227405312)) {
		if base.F64_gt(v8, float64(1)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v109 = m.ExcPending
				if v109 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(418208), int32(0))
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(512946), int32(1772), int32(144171))
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
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
			v18 = base.I64_reinterpret_f64(v7)
			v23 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1072693248)) <= base.Ui32(v23) {
				if base.I32_wrap_i64(v18)|(v23-int32(1072693248)) == int32(0) {
					if int64(0) <= v18 {
						v36 = float64(0)
					} else {
						v36 = float64(3.141592653589793)
					}
					v93 = v36
				} else {
					v93 = base.F64_div(float64(0), base.F64_sub(v7, v7))
				}
			} else {
				if base.Ui32(v23) <= base.Ui32(int32(1071644671)) {
					if base.Ui32(v23) < base.Ui32(int32(1012924417)) {
						v89 = float64(1.5707963267948966)
						v93 = v89
					} else {
						v47 = F_R(m, base.F64_mul(v7, v7))
						mBase = m.M
						v93 = base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(v7, v47)), v7), float64(1.5707963267948966))
					}
				} else {
					if v18 < int64(0) {
						v59 = base.F64_mul(base.F64_add(v7, float64(1)), float64(0.5))
						v60 = base.F64_sqrt(v59)
						v61 = F_R(m, v59)
						mBase = m.M
						v66 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v60, base.F64_add(base.F64_mul(v60, v61), float64(-6.123233995736766e-17))))
						v93 = base.F64_add(v66, v66)
					} else {
						v71 = base.F64_mul(base.F64_sub(float64(1), v7), float64(0.5))
						v72 = base.F64_sqrt(v71)
						v73 = F_R(m, v71)
						mBase = m.M
						v78 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v72) & int64(-4294967296))
						v84 = base.F64_add(base.F64_add(base.F64_mul(v72, v73), base.F64_div(base.F64_sub(v71, base.F64_mul(v78, v78)), base.F64_add(v72, v78))), v78)
						v89 = base.F64_add(v84, v84)
						v93 = v89
					}
				}
			}
			if base.F64_eq(base.F64_abs(v93), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v97 = v93
				v98 = F_Float8GetDatum(m, v97)
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return int32(0)
				} else {
					return v98
				}
			}
		}
	} else {
		v97 = math.Float64frombits(uint64(0x7ff8000000000000))
		v98 = F_Float8GetDatum(m, v97)
		mBase = m.M
		v101 = m.ExcPending
		if v101 != 0 {
			return int32(0)
		} else {
			return v98
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
	var v129 int64
	_ = v129
	var v134 int32
	_ = v134
	var v147 float64
	_ = v147
	var v158 float64
	_ = v158
	var v170 float64
	_ = v170
	var v171 float64
	_ = v171
	var v172 float64
	_ = v172
	var v177 float64
	_ = v177
	var v182 float64
	_ = v182
	var v183 float64
	_ = v183
	var v184 float64
	_ = v184
	var v189 float64
	_ = v189
	var v195 float64
	_ = v195
	var v200 float64
	_ = v200
	var v204 float64
	_ = v204
	var v208 float64
	_ = v208
	var v212 float64
	_ = v212
	var v220 int64
	_ = v220
	var v225 int32
	_ = v225
	var v248 float64
	_ = v248
	var v255 float64
	_ = v255
	var v256 float64
	_ = v256
	var v257 float64
	_ = v257
	var v262 float64
	_ = v262
	var v267 float64
	_ = v267
	var v271 float64
	_ = v271
	var v280 float64
	_ = v280
	var v289 float64
	_ = v289
	var v293 float64
	_ = v293
	var v294 float64
	_ = v294
	var v302 float64
	_ = v302
	var v306 float64
	_ = v306
	var v314 int64
	_ = v314
	var v319 int32
	_ = v319
	var v332 float64
	_ = v332
	var v343 float64
	_ = v343
	var v355 float64
	_ = v355
	var v356 float64
	_ = v356
	var v357 float64
	_ = v357
	var v362 float64
	_ = v362
	var v367 float64
	_ = v367
	var v368 float64
	_ = v368
	var v369 float64
	_ = v369
	var v374 float64
	_ = v374
	var v380 float64
	_ = v380
	var v385 float64
	_ = v385
	var v389 float64
	_ = v389
	var v393 float64
	_ = v393
	var v399 float64
	_ = v399
	var v403 float64
	_ = v403
	var v407 float64
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	v13 = base.F64_abs(v12)
	if base.Ui64(base.I64_reinterpret_f64(v13)) <= base.Ui64(int64(9218868437227405312)) {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1081])))
		if v18 == int32(0) {
			F_init_degree_constants(m)
			mBase = m.M
		} else {
		}
		if base.F64_gt(v13, float64(1)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v419 = m.ExcPending
			if v419 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v422 = m.ExcPending
				if v422 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(418208), int32(0))
					mBase = m.M
					v426 = m.ExcPending
					if v426 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(512946), int32(2127), int32(436818))
						mBase = m.M
						v431 = m.ExcPending
						if v431 != 0 {
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
					v119 = *(*float64)(unsafe.Add(mBase, _consts[1082]))
					v403 = base.F64_add(base.F64_mul(base.F64_div(v115, v119), float64(-30)), float64(90))
				} else {
					v129 = base.I64_reinterpret_f64(v12)
					v134 = base.I32_wrap_i64(int64(base.Ui64(v129)>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072693248)) <= base.Ui32(v134) {
						if base.I32_wrap_i64(v129)|(v134-int32(1072693248)) == int32(0) {
							if int64(0) <= v129 {
								v147 = float64(0)
							} else {
								v147 = float64(3.141592653589793)
							}
							v204 = v147
						} else {
							v204 = base.F64_div(float64(0), base.F64_sub(v12, v12))
						}
					} else {
						if base.Ui32(v134) <= base.Ui32(int32(1071644671)) {
							if base.Ui32(v134) < base.Ui32(int32(1012924417)) {
								v200 = float64(1.5707963267948966)
								v204 = v200
							} else {
								v158 = F_R(m, base.F64_mul(v12, v12))
								mBase = m.M
								v204 = base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(v12, v158)), v12), float64(1.5707963267948966))
							}
						} else {
							if v129 < int64(0) {
								v170 = base.F64_mul(base.F64_add(v12, float64(1)), float64(0.5))
								v171 = base.F64_sqrt(v170)
								v172 = F_R(m, v170)
								mBase = m.M
								v177 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v171, base.F64_add(base.F64_mul(v171, v172), float64(-6.123233995736766e-17))))
								v204 = base.F64_add(v177, v177)
							} else {
								v182 = base.F64_mul(base.F64_sub(float64(1), v12), float64(0.5))
								v183 = base.F64_sqrt(v182)
								v184 = F_R(m, v182)
								mBase = m.M
								v189 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v183) & int64(-4294967296))
								v195 = base.F64_add(base.F64_add(base.F64_mul(v183, v184), base.F64_div(base.F64_sub(v182, base.F64_mul(v189, v189)), base.F64_add(v183, v189))), v189)
								v200 = base.F64_add(v195, v195)
								v204 = v200
							}
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v204
					v208 = *(*float64)(unsafe.Add(mBase, _consts[1083]))
					v403 = base.F64_mul(base.F64_div(v204, v208), float64(60))
				}
			} else {
				v212 = base.F64_neg(v12)
				if base.F64_ge(v12, float64(-0.5)) != 0 {
					v220 = base.I64_reinterpret_f64(v212)
					v225 = base.I32_wrap_i64(int64(base.Ui64(v220)>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072693248)) <= base.Ui32(v225) {
						if base.I32_wrap_i64(v220)|(v225-int32(1072693248)) == int32(0) {
							v302 = base.F64_add(base.F64_mul(v212, float64(1.5707963267948966)), float64(7.52316384526264e-37))
						} else {
							v302 = base.F64_div(float64(0), base.F64_sub(v212, v212))
						}
					} else {
						if base.Ui32(v225) <= base.Ui32(int32(1071644671)) {
							if base.Ui32(v225+int32(-1048576)) < base.Ui32(int32(1044381696)) {
								v294 = v212
								v302 = v294
							} else {
								v248 = F_R(m, base.F64_mul(v212, v212))
								mBase = m.M
								v302 = base.F64_add(base.F64_mul(v212, v248), v212)
							}
						} else {
							v255 = base.F64_mul(base.F64_sub(float64(1), base.F64_abs(v212)), float64(0.5))
							v256 = base.F64_sqrt(v255)
							v257 = F_R(m, v255)
							mBase = m.M
							if base.Ui32(int32(1072640819)) <= base.Ui32(v225) {
								v262 = base.F64_add(base.F64_mul(v256, v257), v256)
								v289 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v262, v262), float64(-6.123233995736766e-17)))
							} else {
								v267 = float64(0.7853981633974483)
								v271 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v256) & int64(-4294967296))
								v280 = base.F64_div(base.F64_sub(v255, base.F64_mul(v271, v271)), base.F64_add(v256, v271))
								v289 = base.F64_add(base.F64_sub(base.F64_sub(v267, base.F64_add(v271, v271)), base.F64_sub(base.F64_mul(base.F64_add(v256, v256), v257), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v280, v280)))), v267)
							}
							if v220 < int64(0) {
								v293 = base.F64_neg(v289)
							} else {
								v293 = v289
							}
							v294 = v293
							v302 = v294
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v302
					v306 = *(*float64)(unsafe.Add(mBase, _consts[1082]))
					v399 = base.F64_mul(base.F64_div(v302, v306), float64(30))
				} else {
					v314 = base.I64_reinterpret_f64(v212)
					v319 = base.I32_wrap_i64(int64(base.Ui64(v314)>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072693248)) <= base.Ui32(v319) {
						if base.I32_wrap_i64(v314)|(v319-int32(1072693248)) == int32(0) {
							if int64(0) <= v314 {
								v332 = float64(0)
							} else {
								v332 = float64(3.141592653589793)
							}
							v389 = v332
						} else {
							v389 = base.F64_div(float64(0), base.F64_sub(v212, v212))
						}
					} else {
						if base.Ui32(v319) <= base.Ui32(int32(1071644671)) {
							if base.Ui32(v319) < base.Ui32(int32(1012924417)) {
								v385 = float64(1.5707963267948966)
								v389 = v385
							} else {
								v343 = F_R(m, base.F64_mul(v212, v212))
								mBase = m.M
								v389 = base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(v212, v343)), v212), float64(1.5707963267948966))
							}
						} else {
							if v314 < int64(0) {
								v355 = base.F64_mul(base.F64_add(v212, float64(1)), float64(0.5))
								v356 = base.F64_sqrt(v355)
								v357 = F_R(m, v355)
								mBase = m.M
								v362 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v356, base.F64_add(base.F64_mul(v356, v357), float64(-6.123233995736766e-17))))
								v389 = base.F64_add(v362, v362)
							} else {
								v367 = base.F64_mul(base.F64_sub(float64(1), v212), float64(0.5))
								v368 = base.F64_sqrt(v367)
								v369 = F_R(m, v367)
								mBase = m.M
								v374 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v368) & int64(-4294967296))
								v380 = base.F64_add(base.F64_add(base.F64_mul(v368, v369), base.F64_div(base.F64_sub(v367, base.F64_mul(v374, v374)), base.F64_add(v368, v374))), v374)
								v385 = base.F64_add(v380, v380)
								v389 = v385
							}
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v389
					v393 = *(*float64)(unsafe.Add(mBase, _consts[1083]))
					v399 = base.F64_add(base.F64_mul(base.F64_div(v389, v393), float64(-60)), float64(90))
				}
				v403 = base.F64_add(v399, float64(90))
			}
			if base.F64_eq(base.F64_abs(v403), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v433 = m.ExcPending
				if v433 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v407 = v403
				v408 = F_Float8GetDatum(m, v407)
				mBase = m.M
				v411 = m.ExcPending
				if v411 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(16)
					return v408
				}
			}
		}
	} else {
		v407 = math.Float64frombits(uint64(0x7ff8000000000000))
		v408 = F_Float8GetDatum(m, v407)
		mBase = m.M
		v411 = m.ExcPending
		if v411 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v408
		}
	}
}
func F_danish_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v354 int32
	_ = v354
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = v10 + int32(3)
	if v8 < v12 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	if v133 < v136 {
		goto L38
	} else {
		goto L39
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v12
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v24 < v10 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v64 < int32(0) {
		goto L1
	} else {
		goto L18
	}
L4:
	;
	v26 = v10
	goto L6
L5:
	;
	v26 = v24
	goto L6
L6:
	;
	v33 = v10
	goto L8
L7:
	;
	v64 = v44
	goto L3
L8:
	;
	if v33 == v26 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v64 = int32(-1)
	goto L3
L11:
	;
	goto L12
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v33))))
	if int32(248) < v39 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v56 = v33 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v56
	v33 = v56
	goto L8
L14:
	;
	v41 = v39 - int32(97)
	if v41 < int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v44 = int32(1)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v41)>>(uint(int32(3))%32)))+uint32(_consts[1317]))))
	if int32(base.Ui32(v48)>>(uint(v41&int32(7))%32))&v44 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	goto L13
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v76 < v75 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v119 < int32(0) {
		goto L1
	} else {
		goto L33
	}
L20:
	;
	v78 = v75
	goto L22
L21:
	;
	v78 = v76
	goto L22
L22:
	;
	v85 = v75
	goto L24
L23:
	;
	v119 = int32(1)
	goto L19
L24:
	;
	if v85 == v78 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v119 = int32(-1)
	goto L19
L27:
	;
	goto L28
L28:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91+v85))))
	if int32(248) < v93 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v95 = v93 - int32(97)
	if v95 < int32(0) {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v95)>>(uint(int32(3))%32)))+uint32(_consts[1317]))))
	if int32(base.Ui32(v101)>>(uint(v95&int32(7))%32))&int32(1) == int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v110 = v85 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v110
	v85 = v110
	goto L24
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v123 = v122 + v119
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	if v126 < v123 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v128 = v123
	goto L36
L35:
	;
	v128 = v126
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125)+4)) = v128
	goto L1
L37:
	;
	return v442
L38:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v233
	v235 = F_r_consonant_pair_1(m, l0)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L44
	} else {
		goto L65
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v136
	if v133 <= v136 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	goto L38
L41:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141+v133-int32(1)))))
	if v145&int32(224) != int32(96) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	if int32(1)<<(uint(v145)%32)&int32(1851440) == int32(0) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v158 = F_find_among_b(m, l0, int32(4230352), int32(32))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	return int32(0)
L45:
	;
	if v158 == int32(0) {
		goto L40
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v165
	switch v158 - int32(1) {
	case 0:
		goto L48
	case 1:
		goto L47
	default:
		goto L38
	}
L47:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L53
L48:
	;
	v169 = F_slice_del(m, l0)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L44
	} else {
		goto L49
	}
L49:
	;
	if int32(0) <= v169 {
		goto L38
	} else {
		goto L50
	}
L50:
	;
	v442 = v169
	goto L37
L51:
	;
	if v225 != 0 {
		goto L38
	} else {
		goto L62
	}
L52:
	;
	v225 = v221
	goto L51
L53:
	;
	if v181 <= v182 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v221 = int32(0)
	goto L52
L55:
	;
	v225 = int32(-1)
	goto L51
L56:
	;
	goto L57
L57:
	;
	v194 = int32(1)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195+v181-v194))))
	if int32(229) < v199 {
		v221 = v194
		goto L52
	} else {
		goto L58
	}
L58:
	;
	v201 = v199 - int32(97)
	if v201 < int32(0) {
		v221 = v194
		goto L52
	} else {
		goto L59
	}
L59:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v201)>>(uint(int32(3))%32)))+uint32(_consts[1318]))))
	if int32(base.Ui32(v207)>>(uint(v201&int32(7))%32))&int32(1) == int32(0) {
		v221 = v194
		goto L52
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v181 - int32(1)
	goto L61
L61:
	;
	goto L54
L62:
	;
	v226 = F_slice_del(m, l0)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L44
	} else {
		goto L63
	}
L63:
	;
	if int32(0) <= v226 {
		goto L38
	} else {
		goto L64
	}
L64:
	;
	v442 = v226
	goto L37
L65:
	;
	if v235 < int32(0) {
		v442 = v235
		goto L37
	} else {
		goto L66
	}
L66:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v239
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v239
	v242 = int32(2)
	v244 = int32(0)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v239-v247 < v242 {
		v257 = v244
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v285
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	if v285 < v288 {
		goto L80
	} else {
		goto L81
	}
L68:
	;
	if v257 == int32(0) {
		goto L67
	} else {
		goto L72
	}
L69:
	;
	goto L68
L70:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v253 = F_memcmp(m, v250+v239-v242, int32(2209220), v242)
	mBase = m.M
	if v253 != 0 {
		v257 = v244
		goto L69
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v239 - v242
	v257 = int32(1)
	goto L69
L72:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v260
	v262 = int32(2)
	v264 = int32(0)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v260-v267 < v262 {
		v277 = v264
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v277 == int32(0) {
		goto L67
	} else {
		goto L77
	}
L74:
	;
	goto L73
L75:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v273 = F_memcmp(m, v270+v260-v262, int32(2209222), v262)
	mBase = m.M
	if v273 != 0 {
		v277 = v264
		goto L74
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v260 - v262
	v277 = int32(1)
	goto L74
L77:
	;
	v280 = F_slice_del(m, l0)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L44
	} else {
		goto L78
	}
L78:
	;
	if v280 < int32(0) {
		v442 = v280
		goto L37
	} else {
		goto L79
	}
L79:
	;
	goto L67
L80:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v338
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)+4))
	if v338 < v341 {
		goto L96
	} else {
		goto L97
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v285
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v288
	v294 = v285 - int32(1)
	if v294 <= v288 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v291
	goto L80
L83:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296+v294))))
	if v298&int32(224) != int32(96) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	if int32(1)<<(uint(v298)%32)&int32(1572992) == int32(0) {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v311 = F_find_among_b(m, l0, int32(4231072), int32(5))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L44
	} else {
		goto L86
	}
L86:
	;
	if v311 == int32(0) {
		goto L82
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v291
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v316
	switch v311 - int32(1) {
	case 0:
		goto L89
	case 1:
		goto L88
	default:
		goto L80
	}
L88:
	;
	v330 = F_slice_from_s(m, l0, int32(3), int32(2209224))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L44
	} else {
		goto L94
	}
L89:
	;
	v320 = F_slice_del(m, l0)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L44
	} else {
		goto L90
	}
L90:
	;
	if v320 < int32(0) {
		v442 = v320
		goto L37
	} else {
		goto L91
	}
L91:
	;
	v324 = F_r_consonant_pair_1(m, l0)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L44
	} else {
		goto L92
	}
L92:
	;
	if int32(0) <= v324 {
		goto L80
	} else {
		goto L93
	}
L93:
	;
	v442 = v324
	goto L37
L94:
	;
	if int32(0) <= v330 {
		goto L80
	} else {
		goto L95
	}
L95:
	;
	v442 = v330
	goto L37
L96:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v439
	v442 = int32(1)
	goto L37
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v338
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v341
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L100
L98:
	;
	if v398 != 0 {
		goto L109
	} else {
		goto L110
	}
L99:
	;
	v398 = v394
	goto L98
L100:
	;
	if v354 <= v341 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v394 = int32(0)
	goto L99
L102:
	;
	v398 = int32(-1)
	goto L98
L103:
	;
	goto L104
L104:
	;
	v367 = int32(1)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368+v354-v367))))
	if int32(122) < v372 {
		v394 = v367
		goto L99
	} else {
		goto L105
	}
L105:
	;
	v374 = v372 - int32(98)
	if v374 < int32(0) {
		v394 = v367
		goto L99
	} else {
		goto L106
	}
L106:
	;
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v374)>>(uint(int32(3))%32)))+uint32(_consts[1319]))))
	if int32(base.Ui32(v380)>>(uint(v374&int32(7))%32))&int32(1) == int32(0) {
		v394 = v367
		goto L99
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v354 - int32(1)
	goto L108
L108:
	;
	goto L101
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
	goto L96
L110:
	;
	goto L111
L111:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v400
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	v404 = F_slice_to(m, l0, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L44
	} else {
		goto L112
	}
L112:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v406))) = v404
	if v404 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	return int32(-1)
L114:
	;
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
	v413 = int32(0)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v404-int32(4))))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v419-v344 < v418 {
		v430 = v413
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v430 == int32(0) {
		goto L96
	} else {
		goto L120
	}
L117:
	;
	goto L116
L118:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v426 = F_memcmp(m, v423+v419-v418, v404, v418)
	mBase = m.M
	if v426 != 0 {
		v430 = v413
		goto L117
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v419 - v418
	v430 = int32(1)
	goto L117
L120:
	;
	v433 = F_slice_del(m, l0)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L44
	} else {
		goto L121
	}
L121:
	;
	if v433 < int32(0) {
		v442 = v433
		goto L37
	} else {
		goto L122
	}
L122:
	;
	goto L96
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
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 float64
	_ = v47
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v85 float64
	_ = v85
	var v102 float64
	_ = v102
	var v103 float64
	_ = v103
	var v119 float64
	_ = v119
	var v120 float64
	_ = v120
	var v129 float64
	_ = v129
	var v136 int32
	_ = v136
	var v137 float64
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	v5 = math.Float64frombits(uint64(0x7ff8000000000000))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v7)&int64(9223372036854775807)) {
		v137 = v5
		v139 = F_Float8GetDatum(m, v137)
		mBase = m.M
		v140 = m.ExcPending
		if v140 != 0 {
			return int32(0)
		} else {
			return v139
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v13)))
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v14)&int64(9223372036854775807)) {
			v137 = v5
			v139 = F_Float8GetDatum(m, v137)
			mBase = m.M
			v140 = m.ExcPending
			if v140 != 0 {
				return int32(0)
			} else {
				return v139
			}
		} else {
			if base.Ui64(base.I64_reinterpret_f64(v14)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
				if base.Ui64(base.I64_reinterpret_f64(v7)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
					v37 = base.I64_reinterpret_f64(v14)
					v40 = base.I32_wrap_i64(int64(base.Ui64(v37) >> (uint(int64(32)) % 64)))
					v43 = base.I32_wrap_i64(v37)
					if v40-int32(1072693248)|v43 == int32(0) {
						v47 = F_atan(m, v7)
						mBase = m.M
						v129 = v47
					} else {
						v51 = int32(base.Ui32(v40)>>(uint(int32(30))%32)) & int32(2)
						v52 = base.I64_reinterpret_f64(v7)
						v56 = v51 | base.I32_wrap_i64(int64(base.Ui64(v52)>>(uint(int64(63))%64)))
						v61 = base.I32_wrap_i64(int64(base.Ui64(v52)>>(uint(int64(32))%64))) & int32(2147483647)
						if v61|base.I32_wrap_i64(v52) == int32(0) {
							switch v56 - int32(2) {
							case 0:
								v129 = float64(3.141592653589793)
							case 1:
								v129 = float64(-3.141592653589793)
							default:
								v120 = v7
								v129 = v120
							}
						} else {
							v71 = v40 & int32(2147483647)
							if v71|v43 == int32(0) {
								v129 = base.F64_copysign(float64(1.5707963267948966), v7)
							} else {
								if v71 == int32(2146435072) {
									if v61 != int32(2146435072) {
										v119 = *(*float64)(unsafe.Add(mBase, uint32(v56<<(uint(int32(3))%32))+uint32(_consts[1079])))
										v120 = v119
										v129 = v120
									} else {
										v85 = *(*float64)(unsafe.Add(mBase, uint32(v56<<(uint(int32(3))%32))+uint32(_consts[1080])))
										v129 = v85
									}
								} else {
									if base.B2i32(v61 != int32(2146435072))&base.B2i32(base.Ui32(v61) <= base.Ui32(v71+int32(67108864))) == int32(0) {
										v129 = base.F64_copysign(float64(1.5707963267948966), v7)
									} else {
										if v51 != 0 {
											if base.Ui32(v61+int32(67108864)) < base.Ui32(v71) {
												v103 = float64(0)
											} else {
												v102 = F_atan(m, base.F64_abs(base.F64_div(v7, v14)))
												mBase = m.M
												v103 = v102
											}
										} else {
											v102 = F_atan(m, base.F64_abs(base.F64_div(v7, v14)))
											mBase = m.M
											v103 = v102
										}
										switch v56 - int32(1) {
										case 0:
											v129 = base.F64_neg(v103)
										case 1:
											v129 = base.F64_sub(float64(3.141592653589793), base.F64_add(v103, float64(-1.2246467991473532e-16)))
										case 2:
											v129 = base.F64_add(base.F64_add(v103, float64(-1.2246467991473532e-16)), float64(-3.141592653589793))
										default:
											v120 = v103
											v129 = v120
										}
									}
								}
							}
						}
					}
				} else {
					v129 = base.F64_add(v7, v14)
				}
			} else {
				v129 = base.F64_add(v7, v14)
			}
			if base.F64_ne(base.F64_abs(v129), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v137 = v129
				v139 = F_Float8GetDatum(m, v137)
				mBase = m.M
				v140 = m.ExcPending
				if v140 != 0 {
					return int32(0)
				} else {
					return v139
				}
			} else {
				F_float_overflow_error(m)
				mBase = m.M
				v136 = m.ExcPending
				if v136 != 0 {
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
	var v124 float64
	_ = v124
	var v129 float64
	_ = v129
	var v132 float64
	_ = v132
	var v136 float64
	_ = v136
	var v137 float64
	_ = v137
	var v149 float64
	_ = v149
	var v153 float64
	_ = v153
	var v156 float64
	_ = v156
	var v163 int32
	_ = v163
	var v164 float64
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v11)&int64(9223372036854775807)) {
		v164 = math.Float64frombits(uint64(0x7ff8000000000000))
		v165 = F_Float8GetDatum(m, v164)
		mBase = m.M
		v166 = m.ExcPending
		if v166 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v165
		}
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1081])))
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
			v149 = v43
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
						v149 = base.F64_sub(v80, base.F64_mul(v80, base.F64_add(v97, v114)))
					} else {
						v121 = v81 << (uint(int32(3)) % 32)
						v124 = *(*float64)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[1084])))
						v129 = *(*float64)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[1085])))
						v132 = base.F64_sub(v124, base.F64_sub(base.F64_sub(base.F64_mul(v80, base.F64_add(v97, v114)), v129), v80))
						if v28 < int64(0) {
							v136 = base.F64_neg(v132)
						} else {
							v136 = v132
						}
						v137 = v136
						v149 = v137
					}
				} else {
					v137 = v11
					v149 = v137
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
					v149 = base.F64_sub(v80, base.F64_mul(v80, base.F64_add(v97, v114)))
				} else {
					v121 = v81 << (uint(int32(3)) % 32)
					v124 = *(*float64)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[1084])))
					v129 = *(*float64)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[1085])))
					v132 = base.F64_sub(v124, base.F64_sub(base.F64_sub(base.F64_mul(v80, base.F64_add(v97, v114)), v129), v80))
					if v28 < int64(0) {
						v136 = base.F64_neg(v132)
					} else {
						v136 = v132
					}
					v137 = v136
					v149 = v137
				}
			}
		}
		*(*float64)(unsafe.Add(mBase, uint32(v7)+8)) = v149
		v153 = *(*float64)(unsafe.Add(mBase, _consts[1086]))
		v156 = base.F64_mul(base.F64_div(v149, v153), float64(45))
		if base.F64_ne(base.F64_abs(v156), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v164 = v156
			v165 = F_Float8GetDatum(m, v164)
			mBase = m.M
			v166 = m.ExcPending
			if v166 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return v165
			}
		} else {
			F_float_overflow_error(m)
			mBase = m.M
			v163 = m.ExcPending
			if v163 != 0 {
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
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
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
	v128 = m.ExcPending
	if v128 != 0 {
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
		v120 = v12
		goto L12
	} else {
		goto L13
	}
L12:
	;
	m.G0 = v9 + int32(32)
	return v120
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
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
	if v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if base.Ui32(v43-int32(2147483647)) < base.Ui32(int32(2)) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
	if v48 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	if base.Ui32(int32(2147483494)) <= base.Ui32(v43+int32(2451546)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v53 = int32(0)
	v54 = F_errsave_start(m, v16)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v70 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)) = uint8(v70)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v43 + v70
	goto L14
L21:
	;
	if v54 == int32(0) {
		v120 = v53
		goto L12
	} else {
		goto L22
	}
L22:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errmsg(m, int32(418579), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errsave_finish(m, v16, int32(514002), int32(1648), int32(326901))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v120 = v53
	goto L12
L26:
	;
	v117 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L38
	}
L27:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if base.Ui32(v77-int32(2147483647)) < base.Ui32(int32(2)) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
	if v82 != int32(1) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	if base.Ui32(int32(2147483494)) <= base.Ui32(v77+int32(2451546)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v89 = int32(0)
	v90 = F_errsave_start(m, v16)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v106 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v77 + int32(1)
	goto L26
L33:
	;
	if v90 == int32(0) {
		v120 = v89
		goto L12
	} else {
		goto L34
	}
L34:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errmsg(m, int32(418579), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errsave_finish(m, v16, int32(514002), int32(1663), int32(326901))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v120 = v89
	goto L12
L38:
	;
	v120 = v117
	goto L12
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
	F_errmsg_internal(m, int32(385707), v9)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(514002), int32(1776), int32(414934))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
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
								F_errmsg(m, int32(418494), int32(0))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(519090), int32(2044), int32(247043))
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
						F_errmsg(m, int32(247302), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(519090), int32(658), int32(32371))
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
							F_errmsg(m, int32(247302), int32(0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(519090), int32(2981), int32(6960))
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
						F_errmsg(m, int32(247302), int32(0))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(519090), int32(2971), int32(6960))
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
		return int32(557347)
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		v12 = v2
		return v12
	case 16:
		v12 = int32(547969)
		return v12
	default:
		if v4 != 0 {
			v12 = v2
			return v12
		} else {
			return int32(530055)
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
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v36 int64
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 float64
	_ = v45
	var v48 int64
	_ = v48
	var v55 float64
	_ = v55
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v76 int32
	_ = v76
	var v90 int32
	_ = v90
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v101 int64
	_ = v101
	var v105 int64
	_ = v105
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int64
	_ = v123
	var v127 int64
	_ = v127
	var v134 int64
	_ = v134
	var v138 int32
	_ = v138
	var v140 int64
	_ = v140
	var v144 int32
	_ = v144
	var v148 int64
	_ = v148
	var v150 int32
	_ = v150
	var v153 int64
	_ = v153
	var v167 int64
	_ = v167
	var v178 float64
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 float64
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 float64
	_ = v192
	var v195 int32
	_ = v195
	var v196 float64
	_ = v196
	var v201 float64
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v221 float64
	_ = v221
	var v225 int32
	_ = v225
	var v226 float64
	_ = v226
	var v227 float64
	_ = v227
	var v232 float64
	_ = v232
	var v234 float64
	_ = v234
	var v236 float64
	_ = v236
	var v239 float64
	_ = v239
	var v243 float64
	_ = v243
	var v247 float64
	_ = v247
	var v251 float64
	_ = v251
	var v260 float64
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v280 float64
	_ = v280
	var v284 int32
	_ = v284
	var v285 float64
	_ = v285
	var v286 float64
	_ = v286
	var v292 float64
	_ = v292
	var v293 float64
	_ = v293
	var v295 float64
	_ = v295
	var v297 float64
	_ = v297
	var v299 float64
	_ = v299
	var v308 float64
	_ = v308
	var v312 float64
	_ = v312
	var v313 int32
	_ = v313
	var v317 float64
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v337 float64
	_ = v337
	var v341 int32
	_ = v341
	var v342 float64
	_ = v342
	var v343 float64
	_ = v343
	var v349 float64
	_ = v349
	var v350 float64
	_ = v350
	var v352 float64
	_ = v352
	var v354 float64
	_ = v354
	var v356 float64
	_ = v356
	var v365 float64
	_ = v365
	var v373 float64
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v393 float64
	_ = v393
	var v397 int32
	_ = v397
	var v398 float64
	_ = v398
	var v399 float64
	_ = v399
	var v404 float64
	_ = v404
	var v406 float64
	_ = v406
	var v408 float64
	_ = v408
	var v411 float64
	_ = v411
	var v415 float64
	_ = v415
	var v419 float64
	_ = v419
	var v423 float64
	_ = v423
	var v429 float64
	_ = v429
	var v430 float64
	_ = v430
	var v432 float64
	_ = v432
	var v435 float64
	_ = v435
	var v438 float64
	_ = v438
	var v441 float64
	_ = v441
	var v447 float64
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	if base.Ui64(base.I64_reinterpret_f64(v13)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L127
	} else {
		goto L129
	}
L2:
	;
	if base.F64_eq(base.F64_abs(v13), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	v447 = math.Float64frombits(uint64(0x7ff8000000000000))
	goto L4
L4:
	;
	v448 = F_Float8GetDatum(m, v447)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L127
	} else {
		goto L128
	}
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1081])))
	if v23 == int32(0) {
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
	v27 = int32(0)
	v36 = base.I64_reinterpret_f64(v13)
	v40 = int32(2047)
	v41 = base.I32_wrap_i64(int64(base.Ui64(v36)>>(uint(int64(52))%64))) & v40
	if v41 != v40 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v180 = base.F64_lt(v178, float64(0))
	if v180 != 0 {
		goto L50
	} else {
		goto L51
	}
L10:
	;
	v48 = v36 << (uint(int64(1)) % 64)
	if base.Ui64(v48) <= base.Ui64(int64(-9156662467374350336)) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v45 = base.F64_mul(v13, float64(360))
	v178 = base.F64_div(v45, v45)
	goto L9
L12:
	;
	if v48 == int64(-9156662467374350336) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if v41 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v55 = base.F64_mul(v13, float64(0))
	goto L17
L16:
	;
	v55 = v13
	goto L17
L17:
	;
	v178 = v55
	goto L9
L18:
	;
	if int32(1031) < v90 {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	v58 = int32(0)
	v60 = v36 << (uint(int64(12)) % 64)
	if int64(0) <= v60 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v90 = v41
	v95 = v36&int64(4503599627370495) | int64(4503599627370496)
	goto L18
L22:
	;
	v64 = v58
	v66 = v60
	goto L25
L23:
	;
	v76 = v58
	goto L24
L24:
	;
	v90 = v76
	v95 = v36 << (uint(base.I64_extend_i32_u(int32(1)-v76)) % 64)
	goto L18
L25:
	;
	v70 = v64 - int32(1)
	v72 = v66 << (uint(int64(1)) % 64)
	if int64(0) <= v72 {
		v64 = v70
		v66 = v72
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v76 = v70
	goto L24
L27:
	;
	goto L26
L28:
	;
	v99 = v90
	v101 = v95
	goto L31
L29:
	;
	v121 = v90
	v123 = v95
	goto L30
L30:
	;
	v127 = v123 - int64(6333186975989760)
	if v127 < int64(0) {
		v134 = v123
		goto L37
	} else {
		goto L38
	}
L31:
	;
	v105 = v101 - int64(6333186975989760)
	if v105 < int64(0) {
		v112 = v101
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v121 = int32(1031)
	v123 = v114
	goto L30
L33:
	;
	v114 = v112 << (uint(int64(1)) % 64)
	v116 = v99 - int32(1)
	if int32(1031) < v116 {
		v99 = v116
		v101 = v114
		goto L31
	} else {
		goto L36
	}
L34:
	;
	if v105 != int64(0) {
		v112 = v105
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v178 = base.F64_mul(v13, float64(0))
	goto L9
L36:
	;
	goto L32
L37:
	;
	if base.Ui64(int64(4503599627370495)) < base.Ui64(v134) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	if v127 != int64(0) {
		v134 = v127
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v178 = base.F64_mul(v13, float64(0))
	goto L9
L40:
	;
	if int32(0) < v150 {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	v150 = v121
	v153 = v134
	goto L40
L42:
	;
	goto L43
L43:
	;
	v138 = v121
	v140 = v134
	goto L44
L44:
	;
	v144 = v138 - int32(1)
	v148 = v140 << (uint(int64(1)) % 64)
	if base.Ui64(v140) < base.Ui64(int64(2251799813685248)) {
		v138 = v144
		v140 = v148
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v150 = v144
	v153 = v148
	goto L40
L46:
	;
	goto L45
L47:
	;
	v167 = v153 - int64(4503599627370496) | base.I64_extend_i32_u(v150)<<(uint(int64(52))%64)
	goto L49
L48:
	;
	v167 = int64(base.Ui64(v153) >> (uint(base.I64_extend_i32_u(int32(1)-v150)) % 64))
	goto L49
L49:
	;
	v178 = base.F64_reinterpret_i64(v167 | v36&int64(-9223372036854775807-1))
	goto L9
L50:
	;
	v181 = int32(-1)
	goto L52
L51:
	;
	v181 = int32(1)
	goto L52
L52:
	;
	if v180 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v184 = base.F64_neg(v178)
	goto L55
L54:
	;
	v184 = v178
	goto L55
L55:
	;
	v186 = base.F64_gt(v184, float64(180))
	if v186 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v187 = v27 - v181
	goto L58
L57:
	;
	v187 = v181
	goto L58
L58:
	;
	if v186 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v195 != 0 {
		goto L93
	} else {
		goto L94
	}
L60:
	;
	v192 = base.F64_sub(float64(360), v184)
	goto L62
L61:
	;
	v192 = v184
	goto L62
L62:
	;
	v195 = base.F64_gt(v192, float64(90))
	if v195 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v196 = base.F64_sub(float64(180), v192)
	goto L65
L64:
	;
	v196 = v192
	goto L65
L65:
	;
	if base.F64_le(v196, float64(60)) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v201 = base.F64_mul(v196, float64(0.017453292519943295))
	v205 = m.G0
	v207 = v205 - int32(16)
	m.G0 = v207
	v214 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v201))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v214) <= base.Ui32(int32(1072243195)) {
		goto L71
	} else {
		goto L72
	}
L67:
	;
	goto L68
L68:
	;
	v260 = base.F64_mul(base.F64_sub(float64(90), v196), float64(0.017453292519943295))
	v264 = m.G0
	v266 = v264 - int32(16)
	m.G0 = v266
	v273 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v260))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v273) <= base.Ui32(int32(1072243195)) {
		goto L82
	} else {
		goto L83
	}
L69:
	;
	v247 = base.F64_sub(float64(1), v243)
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v247
	v251 = *(*float64)(unsafe.Add(mBase, _consts[1087]))
	v312 = base.F64_add(base.F64_mul(base.F64_div(v247, v251), float64(-0.5)), float64(1))
	goto L59
L70:
	;
	m.G0 = v207 + int32(16)
	goto L69
L71:
	;
	if base.Ui32(v214) < base.Ui32(int32(1044816030)) {
		v243 = float64(1)
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v214) {
		v243 = base.F64_sub(v201, v201)
		goto L70
	} else {
		goto L75
	}
L74:
	;
	v221 = F___cos(m, v201, float64(0))
	mBase = m.M
	v243 = v221
	goto L70
L75:
	;
	v225 = F___rem_pio2(m, v201, v207)
	mBase = m.M
	v226 = *(*float64)(unsafe.Add(mBase, uint32(v207)+8))
	v227 = *(*float64)(unsafe.Add(mBase, uint32(v207)))
	switch v225&int32(3) - int32(1) {
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
	v239 = F___sin(m, v227, v226, int32(1))
	mBase = m.M
	v243 = v239
	goto L70
L77:
	;
	v236 = F___cos(m, v227, v226)
	mBase = m.M
	v243 = base.F64_neg(v236)
	goto L70
L78:
	;
	v234 = F___sin(m, v227, v226, int32(1))
	mBase = m.M
	v243 = base.F64_neg(v234)
	goto L70
L79:
	;
	v232 = F___cos(m, v227, v226)
	mBase = m.M
	v243 = v232
	goto L70
L80:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v299
	v308 = *(*float64)(unsafe.Add(mBase, _consts[1088]))
	v312 = base.F64_mul(base.F64_div(v299, v308), float64(0.5))
	goto L59
L81:
	;
	m.G0 = v266 + int32(16)
	goto L80
L82:
	;
	if base.Ui32(v273) < base.Ui32(int32(1045430272)) {
		v299 = v260
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v273) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v280 = F___sin(m, v260, float64(0), int32(0))
	mBase = m.M
	v299 = v280
	goto L81
L86:
	;
	v299 = base.F64_sub(v260, v260)
	goto L81
L87:
	;
	goto L88
L88:
	;
	v284 = F___rem_pio2(m, v260, v266)
	mBase = m.M
	v285 = *(*float64)(unsafe.Add(mBase, uint32(v266)+8))
	v286 = *(*float64)(unsafe.Add(mBase, uint32(v266)))
	switch v284&int32(3) - int32(1) {
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
	v297 = F___cos(m, v286, v285)
	mBase = m.M
	v299 = base.F64_neg(v297)
	goto L81
L90:
	;
	v295 = F___sin(m, v286, v285, int32(1))
	mBase = m.M
	v299 = base.F64_neg(v295)
	goto L81
L91:
	;
	v293 = F___cos(m, v286, v285)
	mBase = m.M
	v299 = v293
	goto L81
L92:
	;
	v292 = F___sin(m, v286, v285, int32(1))
	mBase = m.M
	v299 = v292
	goto L81
L93:
	;
	v313 = v27 - v187
	goto L95
L94:
	;
	v313 = v187
	goto L95
L95:
	;
	if base.F64_le(v196, float64(30)) != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v430 = base.F64_div(v312, v429)
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v430
	v432 = float64(0)
	v435 = *(*float64)(unsafe.Add(mBase, _consts[1089]))
	v438 = base.F64_mul(base.F64_div(v430, v435), base.F64_convert_i32_s(v313))
	if base.F64_eq(v438, v432) != 0 {
		goto L124
	} else {
		goto L125
	}
L97:
	;
	v317 = base.F64_mul(v196, float64(0.017453292519943295))
	v321 = m.G0
	v323 = v321 - int32(16)
	m.G0 = v323
	v330 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v317))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v330) <= base.Ui32(int32(1072243195)) {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	goto L99
L99:
	;
	v373 = base.F64_mul(base.F64_sub(float64(90), v196), float64(0.017453292519943295))
	v377 = m.G0
	v379 = v377 - int32(16)
	m.G0 = v379
	v386 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v373))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v386) <= base.Ui32(int32(1072243195)) {
		goto L115
	} else {
		goto L116
	}
L100:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v356
	v365 = *(*float64)(unsafe.Add(mBase, _consts[1088]))
	v429 = base.F64_mul(base.F64_div(v356, v365), float64(0.5))
	goto L96
L101:
	;
	m.G0 = v323 + int32(16)
	goto L100
L102:
	;
	if base.Ui32(v330) < base.Ui32(int32(1045430272)) {
		v356 = v317
		goto L101
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v330) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v337 = F___sin(m, v317, float64(0), int32(0))
	mBase = m.M
	v356 = v337
	goto L101
L106:
	;
	v356 = base.F64_sub(v317, v317)
	goto L101
L107:
	;
	goto L108
L108:
	;
	v341 = F___rem_pio2(m, v317, v323)
	mBase = m.M
	v342 = *(*float64)(unsafe.Add(mBase, uint32(v323)+8))
	v343 = *(*float64)(unsafe.Add(mBase, uint32(v323)))
	switch v341&int32(3) - int32(1) {
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
	v354 = F___cos(m, v343, v342)
	mBase = m.M
	v356 = base.F64_neg(v354)
	goto L101
L110:
	;
	v352 = F___sin(m, v343, v342, int32(1))
	mBase = m.M
	v356 = base.F64_neg(v352)
	goto L101
L111:
	;
	v350 = F___cos(m, v343, v342)
	mBase = m.M
	v356 = v350
	goto L101
L112:
	;
	v349 = F___sin(m, v343, v342, int32(1))
	mBase = m.M
	v356 = v349
	goto L101
L113:
	;
	v419 = base.F64_sub(float64(1), v415)
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v419
	v423 = *(*float64)(unsafe.Add(mBase, _consts[1087]))
	v429 = base.F64_add(base.F64_mul(base.F64_div(v419, v423), float64(-0.5)), float64(1))
	goto L96
L114:
	;
	m.G0 = v379 + int32(16)
	goto L113
L115:
	;
	if base.Ui32(v386) < base.Ui32(int32(1044816030)) {
		v415 = float64(1)
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v386) {
		v415 = base.F64_sub(v373, v373)
		goto L114
	} else {
		goto L119
	}
L118:
	;
	v393 = F___cos(m, v373, float64(0))
	mBase = m.M
	v415 = v393
	goto L114
L119:
	;
	v397 = F___rem_pio2(m, v373, v379)
	mBase = m.M
	v398 = *(*float64)(unsafe.Add(mBase, uint32(v379)+8))
	v399 = *(*float64)(unsafe.Add(mBase, uint32(v379)))
	switch v397&int32(3) - int32(1) {
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
	v411 = F___sin(m, v399, v398, int32(1))
	mBase = m.M
	v415 = v411
	goto L114
L121:
	;
	v408 = F___cos(m, v399, v398)
	mBase = m.M
	v415 = base.F64_neg(v408)
	goto L114
L122:
	;
	v406 = F___sin(m, v399, v398, int32(1))
	mBase = m.M
	v415 = base.F64_neg(v406)
	goto L114
L123:
	;
	v404 = F___cos(m, v399, v398)
	mBase = m.M
	v415 = v404
	goto L114
L124:
	;
	v441 = v432
	goto L126
L125:
	;
	v441 = v438
	goto L126
L126:
	;
	v447 = v441
	goto L4
L127:
	;
	return int32(0)
L128:
	;
	m.G0 = v9 + int32(16)
	return v448
L129:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L127
	} else {
		goto L130
	}
L130:
	;
	F_errmsg(m, int32(418208), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L127
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(512946), int32(2390), int32(436806))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v36 int32
	_ = v36
	var v49 int32
	_ = v49
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
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v100 int32
	_ = v100
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v19 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(20)
	v36 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_pg_printf(m, int32(784019), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L9
	} else {
		goto L12
	}
L4:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v55 = l2 + v22 + v49<<(uint(int32(4))%32) + v36*int32(100)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+82)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+68))
	v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55)+72)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(28)))) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(24)))) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v17+v22))) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(16)))) = int32(785690)
	if v56 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v67 = int32(116)
	goto L8
L7:
	;
	v67 = int32(102)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(32)))) = v67
	v69 = int32(785690)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v55 + int32(4)
	v77 = v36 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v77
	F_pg_printf(m, int32(784722), v17)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	if v77 != v19 {
		v36 = v77
		goto L4
	} else {
		goto L11
	}
L11:
	;
	goto L5
L12:
	;
	m.G0 = v17 + int32(48)
	return
}
func F_debugtup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v42 int32
	_ = v42
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v153 int32
	_ = v153
	v18 = m.G0
	v20 = v18 - int32(48)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if int32(0) < v23 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = int32(20)
	v42 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_pg_printf(m, int32(784019), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L9
	} else {
		goto L30
	}
L4:
	;
	v59 = v42 + int32(1)
	v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v60 <= v42 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	F_slot_getsomeattrs_int(m, l0, v59)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v42))))
	if v68 == int32(0) {
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
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v42<<(uint(int32(2))%32))))
	v77 = v42 * int32(100)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v77+(v22+int32(88)+v78<<(uint(int32(4))%32)))))
	F_getTypeOutputInfo(m, v83, v20+int32(44), v20+int32(43))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v59 != v23 {
		v42 = v59
		goto L4
	} else {
		goto L29
	}
L14:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v91 = F_OidOutputFunctionCall(m, v90, v75)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v97 = v22 + v26 + v93<<(uint(int32(4))%32) + v77
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+82)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)+68))
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v97)+72)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(28)))) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(24)))) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v20+v26))) = v99
	if v91 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v107 = int32(759221)
	goto L18
L17:
	;
	v107 = int32(785690)
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(16)))) = v107
	if v98 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v111 = int32(116)
	goto L21
L20:
	;
	v111 = int32(102)
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(32)))) = v111
	if v91 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v114 = v91
	goto L24
L23:
	;
	v114 = int32(785690)
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v114
	if v91 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v118 = int32(759079)
	goto L27
L26:
	;
	v118 = int32(785690)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v97 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v59
	F_pg_printf(m, int32(784722), v20)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
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
	m.G0 = v20 + int32(48)
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
	v64 = int32(773381)
	goto L16
L15:
	;
	v64 = int32(785690)
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
	v66 = int32(785690)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v66
	F_appendStringInfo(m, l3, int32(184765), v10)
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		m.Env.Pgmem_zstream_free(m, v5)
		mBase = m.M
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		if v7 != 0 {
			F_ResourceOwnerForget(m, v7, v4, int32(4427916))
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
					v16 = F___memset(m, l0, int32(0), int32(8216))
					mBase = m.M
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
				v16 = F___memset(m, l0, int32(0), int32(8216))
				mBase = m.M
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
		v16 = F___memset(m, l0, int32(0), int32(8216))
		mBase = m.M
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
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
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v312 int32
	_ = v312
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v607 int32
	_ = v607
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v768 int32
	_ = v768
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
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
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v972 int32
	_ = v972
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v994 int32
	_ = v994
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	v20 = F_palloc0(m, int32(44))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l1
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v26 - int32(63) {
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
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v1001 = F_lappend(m, v1000, v20)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L209
	}
L4:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v978 == int32(0) {
		v994 = v972
		goto L3
	} else {
		goto L207
	}
L5:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v944 = F_bms_add_member(m, v942, v943)
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L1
	} else {
		goto L203
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L1
	} else {
		goto L200
	}
L7:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v106 {
	case 0:
		goto L37
	case 1, 5:
		goto L41
	case 2:
		goto L39
	default:
		goto L38
	case 4:
		goto L40
	}
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = l2
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v32 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v994 = v6
	goto L3
L10:
	;
	goto L11
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v35 <= int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v972 = v6
	goto L4
L13:
	;
	goto L14
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v45 = v6
	v47 = v6
	v48 = v38
	goto L15
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v45<<(uint(int32(2))%32))))
	v58 = F_deconstruct_recurse(m, l0, v57, l2, v20, l4)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v972 = v101
	goto L4
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v62+v63<<(uint(int32(2))%32)-int32(4))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v71 = F_bms_add_members(m, v60, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v71
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v74
	if v58 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v103 = v45 + int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v103 < v104 {
		v45 = v103
		v47 = v101
		v48 = v99
		goto L15
	} else {
		goto L34
	}
L20:
	;
	v80 = F_list_concat(m, v47, v58)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v83 = v48 - int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if int32(2) <= v84 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v99 = v48 - int32(1)
	v101 = v80
	goto L19
L24:
	;
	v97 = F_lappend(m, v47, v58)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L33
	}
L25:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _consts[495]))
	if v47 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v95 = F_list_concat(m, v47, v58)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L32
	}
L28:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v91 = v89
	goto L30
L29:
	;
	v91 = int32(0)
	goto L30
L30:
	;
	if v88 < v91+(v83+v84) {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	goto L27
L32:
	;
	v99 = v83
	v101 = v95
	goto L19
L33:
	;
	v99 = v83
	v101 = v97
	goto L19
L34:
	;
	goto L16
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v866
	v868 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v868 == int32(2) {
		goto L174
	} else {
		goto L175
	}
L36:
	;
	v861 = v849
	v862 = v850
	v866 = int32(0)
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = l2
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v813 = F_deconstruct_recurse(m, l0, v812, l2, v20, l4)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L1
	} else {
		goto L171
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L1
	} else {
		goto L168
	}
L39:
	;
	v379 = F_palloc0(m, int32(8))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L88
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = l2
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v341 = F_deconstruct_recurse(m, l0, v340, l2, v20, l4)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L84
	}
L41:
	;
	v108 = F_palloc0(m, int32(8))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v108))) = int64(272)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v113 = F_lappend(m, v112, v108)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v108
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v118 = F_deconstruct_recurse(m, l0, v117, l2, v20, l4)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v121+v122<<(uint(int32(2))%32)-int32(4))))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v130 = F_deconstruct_recurse(m, l0, v129, v108, v20, l4)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v133+v134<<(uint(int32(2))%32)-int32(4))))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v143 = F_bms_add_members(m, v141, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v143
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	v148 = F_bms_union(m, v146, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v148
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v151 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
	v331 = F_bms_union(m, v329, v330)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L83
	}
L49:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v155 = F_bms_add_member(m, v154, v151)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v155
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v160 = F_bms_add_member(m, v158, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v160
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v165 = F_bms_add_member(m, v163, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v165
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	if v169 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	if v226 <= int32(0) {
		goto L48
	} else {
		goto L64
	}
L54:
	;
	v226 = base.I32_ctz(v212) | v213<<(uint(int32(5))%32)
	goto L53
L55:
	;
	v226 = int32(-2)
	goto L53
L56:
	;
	v179 = base.I32_div_s(int32(0), int32(32))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	if v180 <= v179 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v183 = v169 + int32(8)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v183+v179<<(uint(int32(2))%32))))
	v190 = v187 & int32(-1)
	if v190 != 0 {
		v212 = v190
		v213 = v179
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v192 = v179 + int32(1)
	if v192 == v180 {
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v195 = v192
	goto L60
L60:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v183+v195<<(uint(int32(2))%32))))
	if v202 != 0 {
		v212 = v202
		v213 = v195
		goto L54
	} else {
		goto L62
	}
L61:
	;
	goto L55
L62:
	;
	v204 = v195 + int32(1)
	if v204 != v180 {
		v195 = v204
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v232 = v226
	goto L65
L65:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v232 == v243 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L48
L67:
	;
	if v169 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L68:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v245+v232<<(uint(int32(2))%32))))
	if v249 == int32(0) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v249)+96))
	v253 = F_bms_add_member(m, v252, v168)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v249)+96)) = v253
	goto L67
L71:
	;
	if int32(0) < v312 {
		v232 = v312
		goto L65
	} else {
		goto L82
	}
L72:
	;
	v312 = base.I32_ctz(v298) | v299<<(uint(int32(5))%32)
	goto L71
L73:
	;
	v312 = int32(-2)
	goto L71
L74:
	;
	v263 = v232 + int32(1)
	v265 = base.I32_div_s(v263, int32(32))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	if v266 <= v265 {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v269 = v169 + int32(8)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v269+v265<<(uint(int32(2))%32))))
	v276 = v273 & (int32(-1) << (uint(v263) % 32))
	if v276 != 0 {
		v298 = v276
		v299 = v265
		goto L72
	} else {
		goto L76
	}
L76:
	;
	v278 = v265 + int32(1)
	if v278 == v266 {
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v281 = v278
	goto L78
L78:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v269+v281<<(uint(int32(2))%32))))
	if v288 != 0 {
		v298 = v288
		v299 = v281
		goto L72
	} else {
		goto L80
	}
L79:
	;
	goto L73
L80:
	;
	v290 = v281 + int32(1)
	if v290 != v266 {
		v281 = v290
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	goto L66
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v331
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v334
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v336
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v861 = v130
	v862 = v118
	v866 = v338
	goto L35
L84:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+12))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v343)+4))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v344+v345<<(uint(int32(2))%32)-int32(4))))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v353 = F_deconstruct_recurse(m, l0, v352, l2, v20, l4)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+12))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v357+v358<<(uint(int32(2))%32)-int32(4))))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)+12))
	v366 = F_bms_union(m, v355, v365)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v366
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v351)+16))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v364)+16))
	v371 = F_bms_union(m, v369, v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v371
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v374
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v364)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v376
	v849 = v353
	v850 = v341
	goto L36
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v379))) = int32(272)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v384 = F_lappend(m, v383, v379)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v379
	v389 = F_palloc0(m, int32(8))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v389))) = int64(272)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v394 = F_lappend(m, v393, v389)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v394
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v398 = F_deconstruct_recurse(m, l0, v397, v389, v20, l4)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+12))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v400)+4))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v401+v402<<(uint(int32(2))%32)-int32(4))))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	v410 = F_bms_copy(m, v409)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v379)+4)) = v410
	v414 = F_palloc0(m, int32(8))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v414))) = int64(272)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v419 = F_lappend(m, v418, v414)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v419
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v423 = F_deconstruct_recurse(m, l0, v422, v414, v20, l4)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)+12))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v425)+4))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v426+v427<<(uint(int32(2))%32)-int32(4))))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v414)+4))
	v436 = F_bms_add_members(m, v434, v435)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v379)+4)) = v436
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v440 = F_bms_add_members(m, v439, v436)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v440
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v408)+12))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v433)+12))
	v445 = F_bms_union(m, v443, v444)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v445
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v450 = F_bms_add_member(m, v448, v449)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v450
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v455 = F_bms_add_member(m, v453, v454)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v455
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v460 = F_bms_add_member(m, v458, v459)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v460
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v408)+12))
	if v464 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	if int32(0) < v521 {
		goto L114
	} else {
		goto L115
	}
L104:
	;
	v521 = base.I32_ctz(v507) | v508<<(uint(int32(5))%32)
	goto L103
L105:
	;
	v521 = int32(-2)
	goto L103
L106:
	;
	v474 = base.I32_div_s(int32(0), int32(32))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v464)+4))
	if v475 <= v474 {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v478 = v464 + int32(8)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v478+v474<<(uint(int32(2))%32))))
	v485 = v482 & int32(-1)
	if v485 != 0 {
		v507 = v485
		v508 = v474
		goto L104
	} else {
		goto L108
	}
L108:
	;
	v487 = v474 + int32(1)
	if v487 == v475 {
		goto L105
	} else {
		goto L109
	}
L109:
	;
	v490 = v487
	goto L110
L110:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v478+v490<<(uint(int32(2))%32))))
	if v497 != 0 {
		v507 = v497
		v508 = v490
		goto L104
	} else {
		goto L112
	}
L111:
	;
	goto L105
L112:
	;
	v499 = v490 + int32(1)
	if v499 != v475 {
		v490 = v499
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v527 = v521
	goto L117
L115:
	;
	goto L116
L116:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v433)+12))
	if v625 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L117:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v527 == v538 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	goto L116
L119:
	;
	if v464 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L120:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v540+v527<<(uint(int32(2))%32))))
	if v544 == int32(0) {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v544)+96))
	v548 = F_bms_add_member(m, v547, v463)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v544)+96)) = v548
	goto L119
L123:
	;
	if int32(0) < v607 {
		v527 = v607
		goto L117
	} else {
		goto L134
	}
L124:
	;
	v607 = base.I32_ctz(v593) | v594<<(uint(int32(5))%32)
	goto L123
L125:
	;
	v607 = int32(-2)
	goto L123
L126:
	;
	v558 = v527 + int32(1)
	v560 = base.I32_div_s(v558, int32(32))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v464)+4))
	if v561 <= v560 {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v564 = v464 + int32(8)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v564+v560<<(uint(int32(2))%32))))
	v571 = v568 & (int32(-1) << (uint(v558) % 32))
	if v571 != 0 {
		v593 = v571
		v594 = v560
		goto L124
	} else {
		goto L128
	}
L128:
	;
	v573 = v560 + int32(1)
	if v573 == v561 {
		goto L125
	} else {
		goto L129
	}
L129:
	;
	v576 = v573
	goto L130
L130:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v564+v576<<(uint(int32(2))%32))))
	if v583 != 0 {
		v593 = v583
		v594 = v576
		goto L124
	} else {
		goto L132
	}
L131:
	;
	goto L125
L132:
	;
	v585 = v576 + int32(1)
	if v585 != v561 {
		v576 = v585
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	goto L118
L135:
	;
	if int32(0) < v682 {
		goto L146
	} else {
		goto L147
	}
L136:
	;
	v682 = base.I32_ctz(v668) | v669<<(uint(int32(5))%32)
	goto L135
L137:
	;
	v682 = int32(-2)
	goto L135
L138:
	;
	v635 = base.I32_div_s(int32(0), int32(32))
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v625)+4))
	if v636 <= v635 {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v639 = v625 + int32(8)
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v639+v635<<(uint(int32(2))%32))))
	v646 = v643 & int32(-1)
	if v646 != 0 {
		v668 = v646
		v669 = v635
		goto L136
	} else {
		goto L140
	}
L140:
	;
	v648 = v635 + int32(1)
	if v648 == v636 {
		goto L137
	} else {
		goto L141
	}
L141:
	;
	v651 = v648
	goto L142
L142:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v639+v651<<(uint(int32(2))%32))))
	if v658 != 0 {
		v668 = v658
		v669 = v651
		goto L136
	} else {
		goto L144
	}
L143:
	;
	goto L137
L144:
	;
	v660 = v651 + int32(1)
	if v660 != v636 {
		v651 = v660
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v688 = v682
	goto L149
L147:
	;
	goto L148
L148:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v408)+16))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v433)+16))
	v787 = F_bms_union(m, v785, v786)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L1
	} else {
		goto L167
	}
L149:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v688 == v699 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	goto L148
L151:
	;
	if v625 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L152:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v701+v688<<(uint(int32(2))%32))))
	if v705 == int32(0) {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v705)+96))
	v709 = F_bms_add_member(m, v708, v624)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v705)+96)) = v709
	goto L151
L155:
	;
	if int32(0) < v768 {
		v688 = v768
		goto L149
	} else {
		goto L166
	}
L156:
	;
	v768 = base.I32_ctz(v754) | v755<<(uint(int32(5))%32)
	goto L155
L157:
	;
	v768 = int32(-2)
	goto L155
L158:
	;
	v719 = v688 + int32(1)
	v721 = base.I32_div_s(v719, int32(32))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v625)+4))
	if v722 <= v721 {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v725 = v625 + int32(8)
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v725+v721<<(uint(int32(2))%32))))
	v732 = v729 & (int32(-1) << (uint(v719) % 32))
	if v732 != 0 {
		v754 = v732
		v755 = v721
		goto L156
	} else {
		goto L160
	}
L160:
	;
	v734 = v721 + int32(1)
	if v734 == v722 {
		goto L157
	} else {
		goto L161
	}
L161:
	;
	v737 = v734
	goto L162
L162:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v725+v737<<(uint(int32(2))%32))))
	if v744 != 0 {
		v754 = v744
		v755 = v737
		goto L156
	} else {
		goto L164
	}
L163:
	;
	goto L157
L164:
	;
	v746 = v737 + int32(1)
	if v746 != v722 {
		v737 = v746
		goto L162
	} else {
		goto L165
	}
L165:
	;
	goto L163
L166:
	;
	goto L150
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v787
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v408)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v790
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v433)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v792
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v861 = v423
	v862 = v398
	v866 = v794
	goto L35
L168:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v799
	F_errmsg_internal(m, int32(503829), v15+int32(-48))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(516724), int32(1405), int32(375370))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v815)+12))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v815)+4))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v816+v817<<(uint(int32(2))%32)-int32(4))))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v825 = F_deconstruct_recurse(m, l0, v824, l2, v20, l4)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v823)+12))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v828)+12))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v828)+4))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v829+v830<<(uint(int32(2))%32)-int32(4))))
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v836)+12))
	v838 = F_bms_union(m, v827, v837)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v838
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v823)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v842
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v836)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v844
	v849 = v825
	v850 = v813
	goto L36
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v861
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v862
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v862
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v861
	v879 = F_list_make2_impl(m, v15+int32(-36), v15+int32(-40))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L1
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v888 = int32(0)
	v892 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	if v862 != 0 {
		goto L179
	} else {
		goto L180
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = v879
	v886 = F_list_make1_impl(m, int32(1), v15+int32(-44))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	v994 = v886
	goto L3
L179:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v862)+4))
	v894 = v893
	goto L181
L180:
	;
	v894 = v888
	goto L181
L181:
	;
	if v861 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v861)+4))
	v896 = v895
	goto L184
L183:
	;
	v896 = v888
	goto L184
L184:
	;
	if v894+v896 <= v892 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v899 = F_list_concat(m, v862, v861)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L1
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	if v862 == int32(0) {
		v908 = v888
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v994 = v899
	goto L3
L189:
	;
	if v861 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L190:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v862)+4))
	if v903 != int32(1) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v908 = v862
	goto L189
L192:
	;
	goto L193
L193:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v862)+12))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v906)))
	v908 = v907
	goto L189
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v917
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v908
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v908
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v917
	v926 = F_list_make2_impl(m, v15+int32(-28), v15+int32(-32))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L1
	} else {
		goto L199
	}
L195:
	;
	v917 = int32(0)
	goto L194
L196:
	;
	goto L197
L197:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v861)+4))
	if v912 != int32(1) {
		v917 = v861
		goto L194
	} else {
		goto L198
	}
L198:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v861)+12))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v915)))
	v917 = v916
	goto L194
L199:
	;
	v994 = v926
	goto L3
L200:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v932
	F_errmsg_internal(m, int32(504634), v17)
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(516724), int32(1446), int32(375370))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = l2
	v948 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v949 = F_bms_add_member(m, v948, v943)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v949
	v952 = F_bms_make_singleton(m, v943)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v952
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = l1
	v962 = F_list_make1_impl(m, int32(1), v15+int32(-52))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	v994 = v962
	goto L3
L207:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v978)+4))
	if v981 <= int32(1) {
		v994 = v972
		goto L3
	} else {
		goto L208
	}
L208:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v984
	v994 = v972
	goto L3
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1001
	m.G0 = v17 - int32(-64)
	return v994
}
func F_deleteObjectsInList(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
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
	var v47 int32
	_ = v47
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v341 int32
	_ = v341
	var v353 int32
	_ = v353
	var v369 int32
	_ = v369
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v585 int32
	_ = v585
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
	if l2&int32(1) != 0 {
		v100 = v35
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
	if v100 <= int32(0) {
		goto L8
	} else {
		goto L29
	}
L10:
	;
	if v34 == int32(0) {
		v100 = v35
		goto L9
	} else {
		goto L11
	}
L11:
	;
	if v35 <= int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v47 = int32(0)
	goto L13
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v47<<(uint(int32(4))%32))))
	v65 = int32(0)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v69 = v66 + v47*int32(12)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v70 <= int32(3465) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v100 = v95
	goto L9
L15:
	;
	if v84 != 0 {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	v84 = int32(1)
	goto L15
L17:
	;
	if base.Ui32(v70-int32(1260)) < base.Ui32(int32(3)) {
		v84 = v65
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v70 == int32(3466) {
		v84 = v65
		goto L15
	} else {
		goto L22
	}
L20:
	;
	if v70 != int32(1213) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v84 = v65
	goto L15
L22:
	;
	if v70 == int32(6243) {
		v84 = v65
		goto L15
	} else {
		goto L23
	}
L23:
	;
	goto L16
L24:
	;
	F_EventTriggerSQLDropAddObject(m, v69, v64&int32(1), base.B2i32(v64&int32(66) != int32(0)))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L2
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v94 = v47 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v94 < v95 {
		v47 = v94
		goto L13
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	goto L14
L29:
	;
	v131 = v100
	v139 = int32(0)
	goto L30
L30:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l2&int32(8) != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L8
L32:
	;
	v585 = v139 + int32(1)
	if v585 < v570 {
		v131 = v570
		v139 = v585
		goto L30
	} else {
		goto L154
	}
L33:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146+v139<<(uint(int32(4))%32)))))
	if v150&int32(1) != 0 {
		v570 = v131
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v155 = v145 + v139*int32(12)
	v157 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	if v157 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
	F_RunObjectDropHook(m, v158, v159, v160, l2)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L2
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if l2&int32(2) != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L39
L41:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	F_ScanKeyInit(m, v20, int32(1), int32(3), int32(184), v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L49
	}
L42:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_sequence_close(m, v163, int32(3))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L2
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_doDeletion(m, v155, l2)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L48
	}
L45:
	;
	F_doDeletion(m, v155, l2)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v171 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v171
	goto L41
L48:
	;
	goto L41
L49:
	;
	v182 = int32(2)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	F_ScanKeyInit(m, v20+int32(48), v182, int32(3), int32(184), v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
	if v189 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v190 = int32(3)
	F_ScanKeyInit(m, v20+int32(96), v190, v190, int32(65), v189)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L2
	} else {
		goto L54
	}
L52:
	;
	v196 = v182
	goto L53
L53:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v201 = F_systable_beginscan(m, v197, int32(2673), int32(1), int32(0), v196, v20)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L2
	} else {
		goto L55
	}
L54:
	;
	v196 = v190
	goto L53
L55:
	;
	goto L56
L56:
	;
	v220 = F_systable_getnext(m, v201)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L2
	} else {
		goto L58
	}
L57:
	;
	F_systable_endscan(m, v201)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L2
	} else {
		goto L63
	}
L58:
	;
	if v220 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_CatalogTupleDelete(m, v222, v220+int32(4))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L2
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L57
L62:
	;
	goto L56
L63:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
	F_deleteSharedDependencyRecordsFor(m, v229, v230, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
	v237 = m.G0
	v239 = v237 - int32(144)
	m.G0 = v239
	F_ScanKeyInit(m, v239, int32(1), int32(3), int32(184), v234)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	v246 = int32(2)
	F_ScanKeyInit(m, v239+int32(48), v246, int32(3), int32(184), v235)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	if v236 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v256 = int32(3)
	F_ScanKeyInit(m, v239+int32(96), v256, v256, int32(65), v236)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L2
	} else {
		goto L70
	}
L68:
	;
	v262 = v246
	goto L69
L69:
	;
	v265 = F_table_open(m, int32(2609), int32(3))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L2
	} else {
		goto L71
	}
L70:
	;
	v262 = int32(3)
	goto L69
L71:
	;
	v270 = F_systable_beginscan(m, v265, int32(2675), int32(1), int32(0), v262, v239)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v272 = F_systable_getnext(m, v270)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	if v272 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v281 = v272
	goto L77
L75:
	;
	goto L76
L76:
	;
	F_systable_endscan(m, v270)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L2
	} else {
		goto L82
	}
L77:
	;
	F_CatalogTupleDelete(m, v265, v281+int32(4))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L2
	} else {
		goto L79
	}
L78:
	;
	goto L76
L79:
	;
	v295 = F_systable_getnext(m, v270)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	if v295 != 0 {
		v281 = v295
		goto L77
	} else {
		goto L81
	}
L81:
	;
	goto L78
L82:
	;
	F_sequence_close(m, v265, int32(3))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	v319 = int32(144)
	m.G0 = v239 + v319
	v322 = m.G0
	v324 = v322 - v319
	m.G0 = v324
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v329 = int32(1)
	if v326 <= int32(3591) {
		goto L88
	} else {
		goto L89
	}
L84:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v397 != 0 {
		goto L113
	} else {
		goto L114
	}
L85:
	;
	goto L84
L86:
	;
	v397 = int32(0)
	goto L85
L87:
	;
	if base.Ui32(v326-int32(2964)) < base.Ui32(int32(4)) {
		v397 = v329
		goto L85
	} else {
		goto L110
	}
L88:
	;
	if v326 <= int32(2670) {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	if v326 <= int32(5999) {
		goto L99
	} else {
		goto L100
	}
L91:
	;
	switch v326 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v397 = v329
		goto L85
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L86
	default:
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v341 = v326 - int32(2671)
	if base.Ui32(int32(27)) < base.Ui32(v341) {
		goto L87
	} else {
		goto L96
	}
L94:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v326-int32(2396)) {
		goto L86
	} else {
		goto L95
	}
L95:
	;
	v397 = v329
	goto L85
L96:
	;
	if int32(1)<<(uint(v341)%32)&int32(226492515) == int32(0) {
		goto L87
	} else {
		goto L97
	}
L97:
	;
	v397 = v329
	goto L85
L98:
	;
	if base.Ui32(v326-int32(3592)) < base.Ui32(int32(2)) {
		v397 = v329
		goto L85
	} else {
		goto L108
	}
L99:
	;
	v353 = v326 - int32(4177)
	if base.Ui32(int32(9)) < base.Ui32(v353) {
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	switch v326 - int32(6243) {
	case 0, 1, 2, 3, 4, 59, 60:
		v397 = v329
		goto L85
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L86
	default:
		goto L104
	}
L102:
	;
	if int32(1)<<(uint(v353)%32)&int32(963) == int32(0) {
		goto L98
	} else {
		goto L103
	}
L103:
	;
	v397 = v329
	goto L85
L104:
	;
	if base.Ui32(v326-int32(6000)) < base.Ui32(int32(3)) {
		v397 = v329
		goto L85
	} else {
		goto L105
	}
L105:
	;
	v369 = v326 - int32(6100)
	if base.Ui32(int32(15)) < base.Ui32(v369) {
		goto L86
	} else {
		goto L106
	}
L106:
	;
	if int32(1)<<(uint(v369)%32)&int32(49153) != 0 {
		v397 = v329
		goto L85
	} else {
		goto L107
	}
L107:
	;
	goto L86
L108:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v326-int32(4060)) {
		goto L86
	} else {
		goto L109
	}
L109:
	;
	v397 = v329
	goto L85
L110:
	;
	if base.Ui32(v326-int32(2846)) < base.Ui32(int32(2)) {
		v397 = v329
		goto L85
	} else {
		goto L111
	}
L111:
	;
	goto L86
L112:
	;
	m.G0 = v324 + int32(144)
	v504 = F_table_open(m, int32(3394), int32(3))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L2
	} else {
		goto L136
	}
L113:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	F_DeleteSharedSecurityLabel(m, v398, v399)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L2
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	F_ScanKeyInit(m, v324, int32(1), int32(3), int32(184), v398)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L2
	} else {
		goto L117
	}
L116:
	;
	goto L112
L117:
	;
	v407 = int32(2)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	F_ScanKeyInit(m, v324+int32(48), v407, int32(3), int32(184), v413)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
	if v416 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v419 = int32(3)
	F_ScanKeyInit(m, v324+int32(96), v419, v419, int32(65), v416)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L2
	} else {
		goto L122
	}
L120:
	;
	v425 = v407
	goto L121
L121:
	;
	v428 = F_table_open(m, int32(3596), int32(3))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L2
	} else {
		goto L123
	}
L122:
	;
	v425 = int32(3)
	goto L121
L123:
	;
	v433 = F_systable_beginscan(m, v428, int32(3597), int32(1), int32(0), v425, v324)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	v435 = F_systable_getnext(m, v433)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L2
	} else {
		goto L125
	}
L125:
	;
	if v435 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v443 = v435
	goto L129
L127:
	;
	goto L128
L128:
	;
	F_systable_endscan(m, v433)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L2
	} else {
		goto L134
	}
L129:
	;
	F_CatalogTupleDelete(m, v428, v443+int32(4))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L2
	} else {
		goto L131
	}
L130:
	;
	goto L128
L131:
	;
	v458 = F_systable_getnext(m, v433)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L2
	} else {
		goto L132
	}
L132:
	;
	if v458 != 0 {
		v443 = v458
		goto L129
	} else {
		goto L133
	}
L133:
	;
	goto L130
L134:
	;
	F_sequence_close(m, v428, int32(3))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L2
	} else {
		goto L135
	}
L135:
	;
	goto L112
L136:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	F_ScanKeyInit(m, v20+int32(144), int32(1), int32(3), int32(184), v511)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L2
	} else {
		goto L137
	}
L137:
	;
	v514 = int32(2)
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	F_ScanKeyInit(m, v20+int32(192), v514, int32(3), int32(184), v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L2
	} else {
		goto L138
	}
L138:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
	if v521 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v522 = int32(3)
	F_ScanKeyInit(m, v20+int32(240), v522, v522, int32(65), v521)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L2
	} else {
		goto L142
	}
L140:
	;
	v528 = v514
	goto L141
L141:
	;
	v534 = F_systable_beginscan(m, v504, int32(3395), int32(1), int32(0), v528, v20+int32(144))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L2
	} else {
		goto L143
	}
L142:
	;
	v528 = v522
	goto L141
L143:
	;
	goto L144
L144:
	;
	v553 = F_systable_getnext(m, v534)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L2
	} else {
		goto L146
	}
L145:
	;
	F_systable_endscan(m, v534)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L2
	} else {
		goto L151
	}
L146:
	;
	if v553 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	F_CatalogTupleDelete(m, v504, v553+int32(4))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L2
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	goto L145
L150:
	;
	goto L144
L151:
	;
	F_sequence_close(m, v504, int32(3))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L2
	} else {
		goto L152
	}
L152:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L2
	} else {
		goto L153
	}
L153:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v570 = v566
	goto L32
L154:
	;
	goto L31
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	v4 = l0
	goto L2
L1:
	;
	if v11&int32(254) != int32(2) {
		goto L45
	} else {
		goto L46
	}
L2:
	;
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v7 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v22&int32(1) != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	return v4
L5:
	;
	goto L6
L6:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+1)))
	if v11 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v11 != int32(18) {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v4)+2))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v22 == int32(1) {
		v4 = v21
		goto L2
	} else {
		goto L13
	}
L10:
	;
	v16 = F_toast_fetch_datum(m, v4)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	return v16
L13:
	;
	goto L3
L14:
	;
	v32 = int32(base.Ui32(v22) >> (uint(int32(1)) % 32))
	goto L16
L15:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v32 = int32(base.Ui32(v29) >> (uint(int32(2)) % 32))
	goto L16
L16:
	;
	v33 = F_palloc(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v35 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v38 = int32(6)
	v40 = int32(18)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v42 == v40 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	if v35&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L21:
	;
	v45 = v40
	goto L23
L22:
	;
	v45 = int32(2)
	goto L23
L23:
	;
	if v42&int32(254) == int32(2) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v50 = v38
	goto L26
L25:
	;
	v50 = v45
	goto L26
L26:
	;
	if v42 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v53 = v38
	goto L29
L28:
	;
	v53 = v50
	goto L29
L29:
	;
	if v53 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	return v55
L31:
	;
	v54 = F__emscripten_memcpy_bulkmem(m, v33, v21, v53)
	mBase = m.M
	v55 = v54
	goto L33
L32:
	;
	v55 = v33
	goto L33
L33:
	;
	goto L30
L34:
	;
	v60 = int32(base.Ui32(v35) >> (uint(int32(1)) % 32))
	if v60 != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	goto L36
L36:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v66 = int32(base.Ui32(v64) >> (uint(int32(2)) % 32))
	if v66 != 0 {
		goto L42
	} else {
		goto L43
	}
L37:
	;
	return v62
L38:
	;
	v61 = F__emscripten_memcpy_bulkmem(m, v33, v21, v60)
	mBase = m.M
	v62 = v61
	goto L40
L39:
	;
	v62 = v33
	goto L40
L40:
	;
	goto L37
L41:
	;
	return v68
L42:
	;
	v67 = F__emscripten_memcpy_bulkmem(m, v33, v21, v66)
	mBase = m.M
	v68 = v67
	goto L44
L43:
	;
	v68 = v33
	goto L44
L44:
	;
	goto L41
L45:
	;
	return v4
L46:
	;
	goto L47
L47:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v4)+2))
	v76 = F_EOH_get_flat_size(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L11
	} else {
		goto L48
	}
L48:
	;
	v78 = F_palloc(m, v76)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	F_EOH_flatten_into(m, v75, v78, v76)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L11
	} else {
		goto L50
	}
L50:
	;
	return v78
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
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
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
	v60 = v31 + (v53+int32(1))&int32(131070) + int32(8)
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
	v97 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)) = uint8(v97)
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
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+28)) = v65
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v83 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v80 = F__emscripten_memcpy_bulkmem(m, v75, l0+int32(32), v74)
	mBase = m.M
	goto L24
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8
	goto L27
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v8
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v67
	goto L18
}
func F_do_des(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var __phi147 int32
	_ = __phi147
	var v148 int32
	_ = v148
	var __phi148 int32
	_ = __phi148
	var v154 int32
	_ = v154
	var __phi154 int32
	_ = __phi154
	var v155 int32
	_ = v155
	var __phi155 int32
	_ = __phi155
	var v156 int32
	_ = v156
	var __phi156 int32
	_ = __phi156
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
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
	v27 = base.B2i32(int32(0) < l4)
	if int32(0) < l4 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = int32(4712912)
	goto L6
L5:
	;
	v28 = int32(4712976)
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
	v33 = int32(4712784)
	goto L9
L8:
	;
	v33 = int32(4712848)
	goto L9
L9:
	;
	v36 = int32(14)
	v38 = int32(1020)
	v39 = int32(base.Ui32(l0)>>(uint(v36)%32)) & v38
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1342])))
	v42 = int32(22)
	v45 = int32(base.Ui32(l0)>>(uint(v42)%32)) & v38
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+uint32(_consts[1343])))
	v49 = int32(6)
	v52 = int32(base.Ui32(l0)>>(uint(v49)%32)) & v38
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)+uint32(_consts[1344])))
	v56 = int32(2)
	v59 = l0 << (uint(v56) % 32) & v38
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)+uint32(_consts[1345])))
	v66 = int32(base.Ui32(l1)>>(uint(v42)%32)) & v38
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)+uint32(_consts[1346])))
	v73 = int32(base.Ui32(l1)>>(uint(v36)%32)) & v38
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+uint32(_consts[1347])))
	v80 = int32(base.Ui32(l1)>>(uint(v49)%32)) & v38
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+uint32(_consts[1348])))
	v87 = l1 << (uint(v56) % 32) & v38
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+uint32(_consts[1349])))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1350])))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v45)+uint32(_consts[1351])))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v52)+uint32(_consts[1352])))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v59)+uint32(_consts[1353])))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v66)+uint32(_consts[1354])))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v73)+uint32(_consts[1355])))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v80)+uint32(_consts[1356])))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v87)+uint32(_consts[1357])))
	v117 = l4 >> (uint(int32(31)) % 32)
	v125 = v41 | v47 | v54 | v61 | v68 | v75 | v82 | v89
	v128 = v94 | v96 | v99 | v102 | v105 | v108 | v111 | v114
	v133 = l4 ^ v117 - v117
	goto L10
L10:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v137 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v262 = int32(14)
	v264 = int32(1020)
	v265 = int32(base.Ui32(v252)>>(uint(v262)%32)) & v264
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v265)+uint32(_consts[1358])))
	v268 = int32(22)
	v271 = int32(base.Ui32(v252)>>(uint(v268)%32)) & v264
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v271)+uint32(_consts[1359])))
	v275 = int32(6)
	v278 = int32(base.Ui32(v252)>>(uint(v275)%32)) & v264
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v278)+uint32(_consts[1360])))
	v282 = int32(2)
	v285 = v252 << (uint(v282) % 32) & v264
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v285)+uint32(_consts[1361])))
	v292 = int32(base.Ui32(v154)>>(uint(v268)%32)) & v264
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v292)+uint32(_consts[1362])))
	v299 = int32(base.Ui32(v154)>>(uint(v262)%32)) & v264
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v299)+uint32(_consts[1363])))
	v306 = int32(base.Ui32(v154)>>(uint(v275)%32)) & v264
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v306)+uint32(_consts[1364])))
	v313 = v154 << (uint(v282) % 32) & v264
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v313)+uint32(_consts[1365])))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v267 | v273 | v280 | v287 | v294 | v301 | v308 | v315
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v265)+uint32(_consts[1366])))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v271)+uint32(_consts[1367])))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v278)+uint32(_consts[1368])))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v285)+uint32(_consts[1369])))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v292)+uint32(_consts[1370])))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v299)+uint32(_consts[1371])))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v306)+uint32(_consts[1372])))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v313)+uint32(_consts[1373])))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v321 | v323 | v326 | v329 | v332 | v335 | v338 | v341
	return int32(0)
L12:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v143 = v133 - int32(1)
	v145 = *(*int32)(unsafe.Add(mBase, _consts[1374]))
	__phi147 = v33
	__phi148 = v28
	__phi154 = v125
	__phi155 = v128
	__phi156 = int32(16)
	v147 = __phi147
	v148 = __phi148
	v154 = __phi154
	v155 = __phi155
	v156 = __phi156
	goto L17
L15:
	;
	return int32(0)
L16:
	;
	goto L14
L17:
	;
	v170 = int32(16515072)
	v174 = int32(258048)
	v179 = int32(4032)
	v186 = base.I32_rotl(v154, int32(23))&v170 | int32(base.Ui32(v154)>>(uint(int32(11))%32))&v174 | int32(base.Ui32(v154)>>(uint(int32(13))%32))&v179 | int32(base.Ui32(v154)>>(uint(int32(15))%32))&int32(63)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v189 = int32(1)
	v210 = v154<<(uint(v189)%32)&int32(62) | (v154<<(uint(int32(3))%32)&v179 | (v154<<(uint(int32(7))%32)&v170 | (v154<<(uint(int32(5))%32)&v174 | int32(base.Ui32(v154)>>(uint(int32(31))%32)))))
	v212 = v145 & (v186 ^ v210)
	v213 = v186 ^ v187 ^ v212
	v214 = int32(4095)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213&v214)+uint32(_consts[1375]))))
	v218 = int32(2)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217<<(uint(v218)%32))+uint32(_consts[1376])))
	v222 = int32(12)
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v213)>>(uint(v222)%32)))+uint32(_consts[1377]))))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225<<(uint(v218)%32))+uint32(_consts[1378])))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v233 = v210 ^ v231 ^ v212
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v233)>>(uint(v222)%32)))+uint32(_consts[1379]))))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v237<<(uint(v218)%32))+uint32(_consts[1380])))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233&v214)+uint32(_consts[1381]))))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v246<<(uint(v218)%32))+uint32(_consts[1382])))
	v252 = v221 | v229 | v241 | v250 ^ v155
	v253 = int32(4)
	v258 = v156 - v189
	if v258 != 0 {
		__phi147 = v147 + v253
		__phi148 = v148 + v253
		__phi154 = v252
		__phi155 = v154
		__phi156 = v258
		v147 = __phi147
		v148 = __phi148
		v154 = __phi154
		v155 = __phi155
		v156 = __phi156
		goto L17
	} else {
		goto L19
	}
L18:
	;
	if v143 != 0 {
		v125 = v154
		v128 = v252
		v133 = v143
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
	*(*int32)(unsafe.Add(mBase, _consts[158])) = int32(0) - v14
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
	v26 = *(*int32)(unsafe.Add(mBase, _consts[158]))
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
	v32 = *(*int32)(unsafe.Add(mBase, _consts[158]))
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
	*(*int32)(unsafe.Add(mBase, _consts[158])) = v32
	goto L11
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg(m, int32(311188), v7)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(520146), int32(356), int32(372480))
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v12 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v43 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v15 = int32(4)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v17&int32(254) == int32(2) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v30 = int32(1)
	if v12&v30 != 0 {
		v42 = int32(base.Ui32(v12)>>(uint(v30)%32)) - v30
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v26 = v15
	goto L7
L6:
	;
	v26 = base.B2i32(v17 == int32(18)) << (uint(v15) % 32)
	goto L7
L7:
	;
	if v17 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v29 = v15
	goto L10
L9:
	;
	v29 = v26
	goto L10
L10:
	;
	v42 = v29
	goto L1
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v42 = int32(base.Ui32(v36)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L12:
	;
	if v42 <= int32(0) {
		v220 = l0
		goto L23
	} else {
		goto L24
	}
L13:
	;
	v46 = int32(4)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v48&int32(254) == int32(2) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v61 = int32(1)
	if v43&v61 != 0 {
		v73 = int32(base.Ui32(v43)>>(uint(v61)%32)) - v61
		goto L12
	} else {
		goto L22
	}
L16:
	;
	v57 = v46
	goto L18
L17:
	;
	v57 = base.B2i32(v48 == int32(18)) << (uint(v46) % 32)
	goto L18
L18:
	;
	if v48 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v60 = v46
	goto L21
L20:
	;
	v60 = v57
	goto L21
L21:
	;
	v73 = v60
	goto L12
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v73 = int32(base.Ui32(v67)>>(uint(int32(2))%32)) - int32(4)
	goto L12
L23:
	;
	return v220
L24:
	;
	if v73 <= int32(0) {
		v220 = l0
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v78 = int32(1)
	if v43&v78 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v82 = v78
	goto L28
L27:
	;
	v82 = int32(4)
	goto L28
L28:
	;
	v83 = l1 + v82
	v85 = int32(1)
	v86 = v83 + v73 - v85
	if v12&v85 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v91 = v85
	goto L31
L30:
	;
	v91 = int32(4)
	goto L31
L31:
	;
	v92 = l0 + v91
	if l2 == int32(0) {
		v139 = v42
		v141 = v92
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v208 = v200 + int32(4)
	v209 = F_palloc(m, v208)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L59
	} else {
		goto L60
	}
L33:
	;
	if l3 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L34:
	;
	v103 = v42
	v105 = v92
	goto L35
L35:
	;
	if base.Ui32(v86) < base.Ui32(v83) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v198 = l0 + v91 + v42
	v200 = int32(0)
	goto L32
L37:
	;
	v139 = v42
	v141 = v92
	goto L33
L38:
	;
	goto L39
L39:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	v110 = v83
	goto L40
L40:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v121 != v109 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v126 = int32(1)
	if v126 < v103 {
		v103 = v103 - v126
		v105 = v105 + v126
		goto L35
	} else {
		goto L46
	}
L42:
	;
	v124 = v110 + int32(1)
	if base.Ui32(v124) <= base.Ui32(v86) {
		v110 = v124
		goto L40
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	goto L41
L45:
	;
	v139 = v103
	v141 = v105
	goto L33
L46:
	;
	goto L36
L47:
	;
	v198 = v141
	v200 = v139
	goto L32
L48:
	;
	if v139 <= int32(0) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v153 = v139
	v154 = v42 + v92
	goto L50
L50:
	;
	if base.Ui32(v86) < base.Ui32(v83) {
		goto L47
	} else {
		goto L52
	}
L51:
	;
	v198 = v141
	v200 = int32(0)
	goto L32
L52:
	;
	v162 = v154 - int32(1)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	v164 = v83
	goto L54
L53:
	;
	v180 = int32(1)
	if v180 < v153 {
		v153 = v153 - v180
		v154 = v162
		goto L50
	} else {
		goto L58
	}
L54:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v163 == v175 {
		goto L53
	} else {
		goto L56
	}
L55:
	;
	v198 = v141
	v200 = v153
	goto L32
L56:
	;
	v178 = v164 + int32(1)
	if base.Ui32(v178) <= base.Ui32(v86) {
		v164 = v178
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	goto L51
L59:
	;
	return int32(0)
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v208 << (uint(int32(2)) % 32)
	if v200 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v220 = v209
	goto L23
L62:
	;
	v218 = F__emscripten_memcpy_bulkmem(m, v209+int32(4), v198, v200)
	mBase = m.M
	goto L64
L63:
	;
	goto L64
L64:
	;
	goto L61
}
func F_double_to_shortest_decimal_buf(m *base.Module, l0 float64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v69 int64
	_ = v69
	var v72 int64
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v120 int64
	_ = v120
	var v127 int64
	_ = v127
	var v139 int32
	_ = v139
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v149 int64
	_ = v149
	var v150 int64
	_ = v150
	var v152 int64
	_ = v152
	var v155 int64
	_ = v155
	var v156 int64
	_ = v156
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v163 int64
	_ = v163
	var v170 int64
	_ = v170
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v186 int64
	_ = v186
	var v189 int64
	_ = v189
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v212 int64
	_ = v212
	var v216 int64
	_ = v216
	var v217 int64
	_ = v217
	var v222 int32
	_ = v222
	var v223 int64
	_ = v223
	var v227 int64
	_ = v227
	var v233 int64
	_ = v233
	var v234 int64
	_ = v234
	var v236 int64
	_ = v236
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v243 int64
	_ = v243
	var v247 int64
	_ = v247
	var v254 int64
	_ = v254
	var v266 int32
	_ = v266
	var v267 int64
	_ = v267
	var v273 int64
	_ = v273
	var v274 int64
	_ = v274
	var v276 int64
	_ = v276
	var v279 int64
	_ = v279
	var v280 int64
	_ = v280
	var v282 int64
	_ = v282
	var v283 int64
	_ = v283
	var v287 int64
	_ = v287
	var v294 int64
	_ = v294
	var v305 int64
	_ = v305
	var v306 int64
	_ = v306
	var v307 int64
	_ = v307
	var v308 int64
	_ = v308
	var v311 int64
	_ = v311
	var v326 int64
	_ = v326
	var v330 int64
	_ = v330
	var v331 int64
	_ = v331
	var v336 int32
	_ = v336
	var v337 int64
	_ = v337
	var v343 int64
	_ = v343
	var v344 int64
	_ = v344
	var v346 int64
	_ = v346
	var v349 int64
	_ = v349
	var v350 int64
	_ = v350
	var v352 int64
	_ = v352
	var v353 int64
	_ = v353
	var v357 int64
	_ = v357
	var v364 int64
	_ = v364
	var v376 int32
	_ = v376
	var v377 int64
	_ = v377
	var v383 int64
	_ = v383
	var v384 int64
	_ = v384
	var v386 int64
	_ = v386
	var v389 int64
	_ = v389
	var v390 int64
	_ = v390
	var v392 int64
	_ = v392
	var v393 int64
	_ = v393
	var v397 int64
	_ = v397
	var v404 int64
	_ = v404
	var v416 int32
	_ = v416
	var v417 int64
	_ = v417
	var v418 int64
	_ = v418
	var v419 int64
	_ = v419
	var v420 int64
	_ = v420
	var v423 int64
	_ = v423
	var v438 int64
	_ = v438
	var v442 int64
	_ = v442
	var v443 int64
	_ = v443
	var v447 int64
	_ = v447
	var v448 int64
	_ = v448
	var v449 int64
	_ = v449
	var v452 int32
	_ = v452
	var v453 int64
	_ = v453
	var v454 int64
	_ = v454
	var v455 int64
	_ = v455
	var v458 int64
	_ = v458
	var v463 int64
	_ = v463
	var v466 int64
	_ = v466
	var v487 int32
	_ = v487
	var v502 int64
	_ = v502
	var v509 int64
	_ = v509
	var v510 int64
	_ = v510
	var v511 int64
	_ = v511
	var v514 int64
	_ = v514
	var v519 int64
	_ = v519
	var v522 int64
	_ = v522
	var v534 int64
	_ = v534
	var v546 int32
	_ = v546
	var v559 int64
	_ = v559
	var v568 int64
	_ = v568
	var v569 int64
	_ = v569
	var v570 int64
	_ = v570
	var v573 int64
	_ = v573
	var v578 int64
	_ = v578
	var v581 int64
	_ = v581
	var v593 int64
	_ = v593
	var v606 int32
	_ = v606
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int64
	_ = v619
	var v620 int64
	_ = v620
	var v622 int64
	_ = v622
	var v628 int64
	_ = v628
	var v629 int64
	_ = v629
	var v631 int64
	_ = v631
	var v634 int64
	_ = v634
	var v635 int64
	_ = v635
	var v637 int64
	_ = v637
	var v638 int64
	_ = v638
	var v642 int64
	_ = v642
	var v649 int64
	_ = v649
	var v661 int32
	_ = v661
	var v664 int64
	_ = v664
	var v665 int64
	_ = v665
	var v671 int64
	_ = v671
	var v672 int64
	_ = v672
	var v674 int64
	_ = v674
	var v677 int64
	_ = v677
	var v678 int64
	_ = v678
	var v680 int64
	_ = v680
	var v681 int64
	_ = v681
	var v685 int64
	_ = v685
	var v692 int64
	_ = v692
	var v704 int32
	_ = v704
	var v705 int64
	_ = v705
	var v706 int64
	_ = v706
	var v707 int64
	_ = v707
	var v708 int64
	_ = v708
	var v711 int64
	_ = v711
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v733 int64
	_ = v733
	var v737 int64
	_ = v737
	var v738 int64
	_ = v738
	var v743 int32
	_ = v743
	var v744 int64
	_ = v744
	var v748 int64
	_ = v748
	var v754 int64
	_ = v754
	var v755 int64
	_ = v755
	var v757 int64
	_ = v757
	var v760 int64
	_ = v760
	var v761 int64
	_ = v761
	var v763 int64
	_ = v763
	var v764 int64
	_ = v764
	var v768 int64
	_ = v768
	var v775 int64
	_ = v775
	var v787 int32
	_ = v787
	var v788 int64
	_ = v788
	var v794 int64
	_ = v794
	var v795 int64
	_ = v795
	var v797 int64
	_ = v797
	var v800 int64
	_ = v800
	var v801 int64
	_ = v801
	var v803 int64
	_ = v803
	var v804 int64
	_ = v804
	var v808 int64
	_ = v808
	var v815 int64
	_ = v815
	var v827 int32
	_ = v827
	var v828 int64
	_ = v828
	var v829 int64
	_ = v829
	var v830 int64
	_ = v830
	var v831 int64
	_ = v831
	var v834 int64
	_ = v834
	var v849 int64
	_ = v849
	var v853 int64
	_ = v853
	var v854 int64
	_ = v854
	var v859 int32
	_ = v859
	var v860 int64
	_ = v860
	var v866 int64
	_ = v866
	var v867 int64
	_ = v867
	var v869 int64
	_ = v869
	var v872 int64
	_ = v872
	var v873 int64
	_ = v873
	var v875 int64
	_ = v875
	var v876 int64
	_ = v876
	var v880 int64
	_ = v880
	var v887 int64
	_ = v887
	var v899 int32
	_ = v899
	var v900 int64
	_ = v900
	var v906 int64
	_ = v906
	var v907 int64
	_ = v907
	var v909 int64
	_ = v909
	var v912 int64
	_ = v912
	var v913 int64
	_ = v913
	var v915 int64
	_ = v915
	var v916 int64
	_ = v916
	var v920 int64
	_ = v920
	var v927 int64
	_ = v927
	var v939 int32
	_ = v939
	var v940 int64
	_ = v940
	var v941 int64
	_ = v941
	var v942 int64
	_ = v942
	var v943 int64
	_ = v943
	var v946 int64
	_ = v946
	var v961 int64
	_ = v961
	var v965 int64
	_ = v965
	var v966 int64
	_ = v966
	var v970 int64
	_ = v970
	var v971 int64
	_ = v971
	var v972 int64
	_ = v972
	var v979 int64
	_ = v979
	var v996 int32
	_ = v996
	var v1005 int64
	_ = v1005
	var v1007 int64
	_ = v1007
	var v1008 int64
	_ = v1008
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var __phi1018 int32
	_ = __phi1018
	var v1019 int32
	_ = v1019
	var __phi1019 int32
	_ = __phi1019
	var v1021 int32
	_ = v1021
	var __phi1021 int32
	_ = __phi1021
	var v1032 int64
	_ = v1032
	var __phi1032 int64
	_ = __phi1032
	var v1034 int64
	_ = v1034
	var __phi1034 int64
	_ = __phi1034
	var v1035 int64
	_ = v1035
	var __phi1035 int64
	_ = __phi1035
	var v1040 int64
	_ = v1040
	var v1041 int64
	_ = v1041
	var v1042 int64
	_ = v1042
	var v1044 int64
	_ = v1044
	var v1047 int64
	_ = v1047
	var v1050 int64
	_ = v1050
	var v1053 int64
	_ = v1053
	var v1058 int64
	_ = v1058
	var v1064 int64
	_ = v1064
	var v1065 int64
	_ = v1065
	var v1067 int64
	_ = v1067
	var v1071 int64
	_ = v1071
	var v1076 int64
	_ = v1076
	var v1091 int64
	_ = v1091
	var v1100 int64
	_ = v1100
	var v1101 int64
	_ = v1101
	var v1102 int64
	_ = v1102
	var v1105 int64
	_ = v1105
	var v1110 int64
	_ = v1110
	var v1113 int64
	_ = v1113
	var v1125 int64
	_ = v1125
	var v1134 int32
	_ = v1134
	var v1151 int32
	_ = v1151
	var v1160 int64
	_ = v1160
	var v1162 int64
	_ = v1162
	var v1163 int64
	_ = v1163
	var v1169 int64
	_ = v1169
	var v1170 int64
	_ = v1170
	var v1171 int64
	_ = v1171
	var v1173 int64
	_ = v1173
	var v1175 int64
	_ = v1175
	var v1176 int64
	_ = v1176
	var v1179 int64
	_ = v1179
	var v1181 int64
	_ = v1181
	var v1184 int64
	_ = v1184
	var v1196 int64
	_ = v1196
	var v1198 int64
	_ = v1198
	var v1204 int64
	_ = v1204
	var v1209 int64
	_ = v1209
	var v1224 int64
	_ = v1224
	var v1228 int64
	_ = v1228
	var v1229 int64
	_ = v1229
	var v1231 int64
	_ = v1231
	var v1233 int64
	_ = v1233
	var v1234 int64
	_ = v1234
	var v1237 int64
	_ = v1237
	var v1239 int64
	_ = v1239
	var v1242 int64
	_ = v1242
	var v1254 int64
	_ = v1254
	var v1262 int32
	_ = v1262
	var v1263 int64
	_ = v1263
	var v1264 int64
	_ = v1264
	var v1265 int64
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var __phi1271 int32
	_ = __phi1271
	var v1272 int32
	_ = v1272
	var __phi1272 int32
	_ = __phi1272
	var v1284 int64
	_ = v1284
	var __phi1284 int64
	_ = __phi1284
	var v1285 int64
	_ = v1285
	var __phi1285 int64
	_ = __phi1285
	var v1286 int64
	_ = v1286
	var __phi1286 int64
	_ = __phi1286
	var v1293 int64
	_ = v1293
	var v1294 int64
	_ = v1294
	var v1295 int64
	_ = v1295
	var v1297 int64
	_ = v1297
	var v1300 int64
	_ = v1300
	var v1303 int64
	_ = v1303
	var v1306 int64
	_ = v1306
	var v1311 int64
	_ = v1311
	var v1317 int64
	_ = v1317
	var v1318 int64
	_ = v1318
	var v1320 int64
	_ = v1320
	var v1324 int64
	_ = v1324
	var v1329 int64
	_ = v1329
	var v1344 int64
	_ = v1344
	var v1346 int64
	_ = v1346
	var v1347 int64
	_ = v1347
	var v1348 int64
	_ = v1348
	var v1351 int64
	_ = v1351
	var v1356 int64
	_ = v1356
	var v1359 int64
	_ = v1359
	var v1371 int64
	_ = v1371
	var v1384 int32
	_ = v1384
	var v1389 int32
	_ = v1389
	var v1400 int64
	_ = v1400
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1411 int64
	_ = v1411
	var v1424 int32
	_ = v1424
	var v1430 int64
	_ = v1430
	var v1485 int32
	_ = v1485
	var v1495 int32
	_ = v1495
	var v1501 int64
	_ = v1501
	var v1510 int32
	_ = v1510
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1522 int32
	_ = v1522
	var v1532 int64
	_ = v1532
	var v1537 int32
	_ = v1537
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1545 int64
	_ = v1545
	var v1546 int64
	_ = v1546
	var v1547 int64
	_ = v1547
	var v1549 int64
	_ = v1549
	var v1550 int64
	_ = v1550
	var v1555 int64
	_ = v1555
	var v1558 int64
	_ = v1558
	var v1570 int64
	_ = v1570
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1585 int32
	_ = v1585
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1595 int32
	_ = v1595
	var v1600 int32
	_ = v1600
	var v1613 int32
	_ = v1613
	var v1626 int32
	_ = v1626
	var v1630 int32
	_ = v1630
	var v1634 int64
	_ = v1634
	var v1637 int32
	_ = v1637
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1670 int32
	_ = v1670
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1680 int32
	_ = v1680
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1741 int32
	_ = v1741
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1758 int32
	_ = v1758
	var v1762 int32
	_ = v1762
	var v1764 int32
	_ = v1764
	var v1769 int64
	_ = v1769
	var v1772 int32
	_ = v1772
	var v1775 int32
	_ = v1775
	var v1778 int32
	_ = v1778
	var v1782 int32
	_ = v1782
	var v1786 int32
	_ = v1786
	var v1789 int32
	_ = v1789
	var v1793 int32
	_ = v1793
	var v1797 int32
	_ = v1797
	var v1800 int32
	_ = v1800
	var v1804 int32
	_ = v1804
	var v1812 int32
	_ = v1812
	var v1815 int32
	_ = v1815
	var v1830 int32
	_ = v1830
	var v1835 int64
	_ = v1835
	var v1849 int64
	_ = v1849
	var v1850 int64
	_ = v1850
	var v1851 int64
	_ = v1851
	var v1853 int64
	_ = v1853
	var v1854 int64
	_ = v1854
	var v1859 int64
	_ = v1859
	var v1862 int64
	_ = v1862
	var v1874 int64
	_ = v1874
	var v1891 int32
	_ = v1891
	var v1898 int64
	_ = v1898
	var v1910 int32
	_ = v1910
	var v1913 int64
	_ = v1913
	var v1914 int64
	_ = v1914
	var v1915 int64
	_ = v1915
	var v1917 int64
	_ = v1917
	var v1918 int64
	_ = v1918
	var v1923 int64
	_ = v1923
	var v1926 int64
	_ = v1926
	var v1938 int64
	_ = v1938
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1949 int32
	_ = v1949
	var v1953 int32
	_ = v1953
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1963 int32
	_ = v1963
	var v1968 int32
	_ = v1968
	var v1981 int32
	_ = v1981
	var v1994 int32
	_ = v1994
	var v1998 int32
	_ = v1998
	var v2002 int64
	_ = v2002
	var v2005 int32
	_ = v2005
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2034 int32
	_ = v2034
	var v2038 int32
	_ = v2038
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2048 int32
	_ = v2048
	var v2059 int32
	_ = v2059
	var v2062 int32
	_ = v2062
	var v2068 int32
	_ = v2068
	var v2070 int32
	_ = v2070
	var v2096 int32
	_ = v2096
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2109 int32
	_ = v2109
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2121 int32
	_ = v2121
	var v2124 int32
	_ = v2124
	var v2128 int32
	_ = v2128
	var v2132 int32
	_ = v2132
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2139 int32
	_ = v2139
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2156 int32
	_ = v2156
	var v2159 int32
	_ = v2159
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2169 int32
	_ = v2169
	var v2175 int32
	_ = v2175
	var v2184 int32
	_ = v2184
	var v2190 int32
	_ = v2190
	var v2195 int32
	_ = v2195
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2211 int32
	_ = v2211
	var v2236 int32
	_ = v2236
	var v2241 int32
	_ = v2241
	v3 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(288)
	m.G0 = v27
	v29 = base.I64_reinterpret_f64(l0)
	v31 = v29 & int64(4503599627370495)
	v35 = int32(2047)
	v36 = base.I32_wrap_i64(int64(base.Ui64(v29)>>(uint(int64(52))%64))) & v35
	if v36 != v35 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v27 + int32(288)
	v2241 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1+v2236))) = uint8(v2241)
	return
L2:
	;
	if v29 < int64(0) {
		goto L211
	} else {
		goto L212
	}
L3:
	;
	if base.Ui32(int32(52)) < base.Ui32(v36-int32(1023)) {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	if v31 != int64(0) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if v31 == int64(0) {
		goto L2
	} else {
		goto L9
	}
L7:
	;
	if v36 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L2
L9:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1274])))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)) = uint8(v44)
	v47 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1275])))
	*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v47)
	v2236 = int32(3)
	goto L1
L10:
	;
	if v29 < int64(0) {
		goto L123
	} else {
		goto L124
	}
L11:
	;
	if base.Ui64(int64(999999999999999)) < base.Ui64(v1430) {
		v1495 = v1424
		v1501 = v1430
		v1510 = int32(16)
		goto L10
	} else {
		goto L106
	}
L12:
	;
	v69 = v31 << (uint(int64(2)) % 64)
	if v36 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v54 = int64(-1)
	v57 = base.I64_extend_i32_u(int32(1075) - v36)
	if v31&(v54<<(uint(v57)%64)^v54) != int64(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v1424 = v3
	v1430 = int64(base.Ui64(v31|int64(4503599627370496)) >> (uint(v57) % 64))
	goto L11
L15:
	;
	v72 = v69 | int64(18014398509481984)
	goto L17
L16:
	;
	v72 = v69
	goto L17
L17:
	;
	v77 = base.B2i32(base.Ui32(v36) < base.Ui32(int32(2))) | base.B2i32(v31 != int64(0))
	if v36 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v1407 = v1384 + v1389
	v1411 = v1400 + base.I64_extend_i32_u(v1406)&int64(1)
	if base.Ui64(v1411) <= base.Ui64(int64(9999999999999999)) {
		v1424 = v1407
		v1430 = v1411
		goto L11
	} else {
		goto L105
	}
L19:
	;
	v1169 = int64(34)
	v1170 = int64(base.Ui64(v1160) >> (uint(v1169) % 64))
	v1171 = int64(1546188227)
	v1173 = int64(2)
	v1175 = int64(4294967295)
	v1176 = int64(base.Ui64(v1160)>>(uint(v1173)%64)) & v1175
	v1179 = int64(32)
	v1181 = v1170*v1171 + int64(base.Ui64(v1176*v1171)>>(uint(v1179)%64))
	v1184 = int64(687194767)
	v1196 = int64(base.Ui64(int64(base.Ui64(v1181)>>(uint(v1179)%64))+v1170*v1184+int64(base.Ui64(v1176*v1184+v1181&v1175)>>(uint(v1179)%64))) >> (uint(v1173) % 64))
	v1198 = int64(base.Ui64(v1163) >> (uint(v1169) % 64))
	v1204 = int64(base.Ui64(v1163)>>(uint(v1173)%64)) & v1175
	v1209 = v1198*v1171 + int64(base.Ui64(v1204*v1171)>>(uint(v1179)%64))
	v1224 = int64(base.Ui64(int64(base.Ui64(v1209)>>(uint(v1179)%64))+v1198*v1184+int64(base.Ui64(v1204*v1184+v1209&v1175)>>(uint(v1179)%64))) >> (uint(v1173) % 64))
	if base.Ui64(v1196) <= base.Ui64(v1224) {
		goto L97
	} else {
		goto L98
	}
L20:
	;
	v1013 = int32(0)
	__phi1018 = v1013
	__phi1019 = v1013
	__phi1021 = int32(1)
	__phi1032 = v1005
	__phi1034 = v1007
	__phi1035 = v1008
	v1018 = __phi1018
	v1019 = __phi1019
	v1021 = __phi1021
	v1032 = __phi1032
	v1034 = __phi1034
	v1035 = __phi1035
	goto L91
L21:
	;
	v81 = v36 - int32(1077)
	goto L23
L22:
	;
	v81 = int32(-1076)
	goto L23
L23:
	;
	if int32(0) <= v81 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v85 = v27 + int32(128)
	v92 = int32(base.Ui32(v81*int32(78913))>>(uint(int32(18))%32)) - base.B2i32(base.Ui32(int32(3)) < base.Ui32(v81))
	v94 = v92 << (uint(int32(4)) % 32)
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v94)+uint32(_consts[1276])))
	v98 = int64(0)
	v100 = v72 | int64(2)
	v106 = int64(32)
	v107 = int64(base.Ui64(v100) >> (uint(v106) % 64))
	v109 = int64(base.Ui64(v97) >> (uint(v106) % 64))
	v112 = int64(4294967295)
	v113 = v100 & v112
	v115 = v97 & v112
	v116 = v113 * v115
	v120 = int64(base.Ui64(v116)>>(uint(v106)%64)) + v113*v109
	v127 = v115*v107 + v120&v112
	*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v97*v98 + v98*v100 + v107*v109 + int64(base.Ui64(v120)>>(uint(v106)%64)) + int64(base.Ui64(v127)>>(uint(v106)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v85))) = v116&v112 | v127<<(uint(v106)%64)
	goto L27
L25:
	;
	goto L26
L26:
	;
	v606 = v27 + int32(272)
	v614 = int32(base.Ui32(v81*int32(-732923))>>(uint(int32(20))%32)) - base.B2i32(v81 != int32(-1))
	v615 = v614 + v81
	v617 = v615 << (uint(int32(4)) % 32)
	v619 = *(*int64)(unsafe.Add(mBase, uint32(int32(1865744)-v617)))
	v620 = int64(0)
	v622 = v72 | int64(2)
	v628 = int64(32)
	v629 = int64(base.Ui64(v622) >> (uint(v628) % 64))
	v631 = int64(base.Ui64(v619) >> (uint(v628) % 64))
	v634 = int64(4294967295)
	v635 = v622 & v634
	v637 = v619 & v634
	v638 = v635 * v637
	v642 = int64(base.Ui64(v638)>>(uint(v628)%64)) + v635*v631
	v649 = v637*v629 + v642&v634
	*(*int64)(unsafe.Add(mBase, uint32(v606)+8)) = v619*v620 + v620*v622 + v629*v631 + int64(base.Ui64(v642)>>(uint(v628)%64)) + int64(base.Ui64(v649)>>(uint(v628)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v606))) = v638&v634 | v649<<(uint(v628)%64)
	goto L62
L27:
	;
	v139 = v27 + int32(112)
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v94)+uint32(_consts[1277])))
	v143 = int64(0)
	v149 = int64(32)
	v150 = int64(base.Ui64(v100) >> (uint(v149) % 64))
	v152 = int64(base.Ui64(v142) >> (uint(v149) % 64))
	v155 = int64(4294967295)
	v156 = v100 & v155
	v158 = v142 & v155
	v159 = v156 * v158
	v163 = int64(base.Ui64(v159)>>(uint(v149)%64)) + v156*v152
	v170 = v158*v150 + v163&v155
	*(*int64)(unsafe.Add(mBase, uint32(v139)+8)) = v142*v143 + v143*v100 + v150*v152 + int64(base.Ui64(v163)>>(uint(v149)%64)) + int64(base.Ui64(v170)>>(uint(v149)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v139))) = v159&v155 | v170<<(uint(v149)%64)
	goto L28
L28:
	;
	v182 = v27 + int32(96)
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v27)+136))
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v27)+112))
	v185 = v183 + v184
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v27)+120))
	v189 = v186 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v185) < base.Ui64(v183)))
	v195 = v92 - v81 + int32(base.Ui32(v92*int32(1217359))>>(uint(int32(19))%32))
	v197 = v195 + int32(58)
	if v197&int32(64) != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v222 = v27 + int32(32)
	v223 = int64(0)
	v227 = v72 + base.I64_extend_i32_s(v77^int32(-1))
	v233 = int64(32)
	v234 = int64(base.Ui64(v227) >> (uint(v233) % 64))
	v236 = int64(base.Ui64(v97) >> (uint(v233) % 64))
	v239 = int64(4294967295)
	v240 = v227 & v239
	v242 = v97 & v239
	v243 = v240 * v242
	v247 = int64(base.Ui64(v243)>>(uint(v233)%64)) + v240*v236
	v254 = v242*v234 + v247&v239
	*(*int64)(unsafe.Add(mBase, uint32(v222)+8)) = v97*v223 + v223*v227 + v234*v236 + int64(base.Ui64(v247)>>(uint(v233)%64)) + int64(base.Ui64(v254)>>(uint(v233)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v222))) = v243&v239 | v254<<(uint(v233)%64)
	goto L35
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v182))) = v216
	*(*int64)(unsafe.Add(mBase, uint32(v182)+8)) = v217
	goto L29
L31:
	;
	v216 = int64(base.Ui64(v189) >> (uint(base.I64_extend_i32_u(v195+int32(-6))) % 64))
	v217 = int64(0)
	goto L30
L32:
	;
	goto L33
L33:
	;
	if v197 == int32(0) {
		v216 = v185
		v217 = v189
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v212 = base.I64_extend_i32_u(v197)
	v216 = v189<<(uint(base.I64_extend_i32_u(int32(64)-v197))%64) | int64(base.Ui64(v185)>>(uint(v212)%64))
	v217 = int64(base.Ui64(v189) >> (uint(v212) % 64))
	goto L30
L35:
	;
	v266 = v27 + int32(16)
	v267 = int64(0)
	v273 = int64(32)
	v274 = int64(base.Ui64(v227) >> (uint(v273) % 64))
	v276 = int64(base.Ui64(v142) >> (uint(v273) % 64))
	v279 = int64(4294967295)
	v280 = v227 & v279
	v282 = v142 & v279
	v283 = v280 * v282
	v287 = int64(base.Ui64(v283)>>(uint(v273)%64)) + v280*v276
	v294 = v282*v274 + v287&v279
	*(*int64)(unsafe.Add(mBase, uint32(v266)+8)) = v142*v267 + v267*v227 + v274*v276 + int64(base.Ui64(v287)>>(uint(v273)%64)) + int64(base.Ui64(v294)>>(uint(v273)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v266))) = v283&v279 | v294<<(uint(v273)%64)
	goto L36
L36:
	;
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v27)+40))
	v306 = *(*int64)(unsafe.Add(mBase, uint32(v27)+16))
	v307 = v305 + v306
	v308 = *(*int64)(unsafe.Add(mBase, uint32(v27)+24))
	v311 = v308 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v307) < base.Ui64(v305)))
	if v197&int32(64) != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v336 = v27 + int32(80)
	v337 = int64(0)
	v343 = int64(32)
	v344 = int64(base.Ui64(v72) >> (uint(v343) % 64))
	v346 = int64(base.Ui64(v97) >> (uint(v343) % 64))
	v349 = int64(4294967295)
	v350 = v72 & v349
	v352 = v97 & v349
	v353 = v350 * v352
	v357 = int64(base.Ui64(v353)>>(uint(v343)%64)) + v350*v346
	v364 = v352*v344 + v357&v349
	*(*int64)(unsafe.Add(mBase, uint32(v336)+8)) = v97*v337 + v337*v72 + v344*v346 + int64(base.Ui64(v357)>>(uint(v343)%64)) + int64(base.Ui64(v364)>>(uint(v343)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v336))) = v353&v349 | v364<<(uint(v343)%64)
	goto L43
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27))) = v330
	*(*int64)(unsafe.Add(mBase, uint32(v27)+8)) = v331
	goto L37
L39:
	;
	v330 = int64(base.Ui64(v311) >> (uint(base.I64_extend_i32_u(v195+int32(-6))) % 64))
	v331 = int64(0)
	goto L38
L40:
	;
	goto L41
L41:
	;
	if v197 == int32(0) {
		v330 = v307
		v331 = v311
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v326 = base.I64_extend_i32_u(v197)
	v330 = v311<<(uint(base.I64_extend_i32_u(int32(64)-v197))%64) | int64(base.Ui64(v307)>>(uint(v326)%64))
	v331 = int64(base.Ui64(v311) >> (uint(v326) % 64))
	goto L38
L43:
	;
	v376 = v27 - int32(-64)
	v377 = int64(0)
	v383 = int64(32)
	v384 = int64(base.Ui64(v72) >> (uint(v383) % 64))
	v386 = int64(base.Ui64(v142) >> (uint(v383) % 64))
	v389 = int64(4294967295)
	v390 = v72 & v389
	v392 = v142 & v389
	v393 = v390 * v392
	v397 = int64(base.Ui64(v393)>>(uint(v383)%64)) + v390*v386
	v404 = v392*v384 + v397&v389
	*(*int64)(unsafe.Add(mBase, uint32(v376)+8)) = v142*v377 + v377*v72 + v384*v386 + int64(base.Ui64(v397)>>(uint(v383)%64)) + int64(base.Ui64(v404)>>(uint(v383)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v376))) = v393&v389 | v404<<(uint(v383)%64)
	goto L44
L44:
	;
	v416 = v27 + int32(48)
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v27)+88))
	v418 = *(*int64)(unsafe.Add(mBase, uint32(v27)+64))
	v419 = v417 + v418
	v420 = *(*int64)(unsafe.Add(mBase, uint32(v27)+72))
	v423 = v420 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v419) < base.Ui64(v417)))
	if v197&int32(64) != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v447 = *(*int64)(unsafe.Add(mBase, uint32(v27)+48))
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v27)+96))
	if base.Ui32(int32(21)) < base.Ui32(v92) {
		v1151 = v92
		v1160 = v449
		v1162 = v447
		v1163 = v448
		goto L19
	} else {
		goto L51
	}
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v416))) = v442
	*(*int64)(unsafe.Add(mBase, uint32(v416)+8)) = v443
	goto L45
L47:
	;
	v442 = int64(base.Ui64(v423) >> (uint(base.I64_extend_i32_u(v195+int32(-6))) % 64))
	v443 = int64(0)
	goto L46
L48:
	;
	goto L49
L49:
	;
	if v197 == int32(0) {
		v442 = v419
		v443 = v423
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v438 = base.I64_extend_i32_u(v197)
	v442 = v423<<(uint(base.I64_extend_i32_u(int32(64)-v197))%64) | int64(base.Ui64(v419)>>(uint(v438)%64))
	v443 = int64(base.Ui64(v423) >> (uint(v438) % 64))
	goto L46
L51:
	;
	v452 = int32(0)
	v453 = int64(32)
	v454 = int64(base.Ui64(v72) >> (uint(v453) % 64))
	v455 = int64(3435973837)
	v458 = v72 & int64(4294967292)
	v463 = v454*v455 + int64(base.Ui64(v458*v455)>>(uint(v453)%64))
	v466 = int64(3435973836)
	if base.I32_wrap_i64(int64(base.Ui64(int64(base.Ui64(v463)>>(uint(v453)%64))+v454*v466+int64(base.Ui64(v458*v466+v463&int64(4294967280))>>(uint(v453)%64)))>>(uint(int64(2))%64))*int64(4294967291)+v72) == v452 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v487 = v452
	v502 = v72
	goto L55
L53:
	;
	goto L54
L54:
	;
	v546 = v452
	v559 = v100
	goto L59
L55:
	;
	v509 = int64(32)
	v510 = int64(base.Ui64(v502) >> (uint(v509) % 64))
	v511 = int64(3435973837)
	v514 = v502 & int64(4294967295)
	v519 = v510*v511 + int64(base.Ui64(v514*v511)>>(uint(v509)%64))
	v522 = int64(3435973836)
	v534 = int64(base.Ui64(int64(base.Ui64(v519)>>(uint(v509)%64))+v510*v522+int64(base.Ui64(v514*v522+v519&int64(4294967292))>>(uint(v509)%64))) >> (uint(int64(2)) % 64))
	if base.I32_wrap_i64(v534*int64(4294967291)+v502) == int32(0) {
		v487 = v487 + int32(1)
		v502 = v534
		goto L55
	} else {
		goto L57
	}
L56:
	;
	if base.Ui32(v487) < base.Ui32(v92) {
		v1151 = v92
		v1160 = v449
		v1162 = v447
		v1163 = v448
		goto L19
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v996 = v92
	v1005 = v449
	v1007 = v447
	v1008 = v448
	goto L20
L59:
	;
	v568 = int64(32)
	v569 = int64(base.Ui64(v559) >> (uint(v568) % 64))
	v570 = int64(3435973837)
	v573 = v559 & int64(4294967295)
	v578 = v569*v570 + int64(base.Ui64(v573*v570)>>(uint(v568)%64))
	v581 = int64(3435973836)
	v593 = int64(base.Ui64(int64(base.Ui64(v578)>>(uint(v568)%64))+v569*v581+int64(base.Ui64(v573*v581+v578&int64(4294967292))>>(uint(v568)%64))) >> (uint(int64(2)) % 64))
	if base.I32_wrap_i64(v593*int64(4294967291)+v559) == int32(0) {
		v546 = v546 + int32(1)
		v559 = v593
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v1151 = v92
	v1160 = v449 - base.I64_extend_i32_u(base.B2i32(base.Ui32(v92) <= base.Ui32(v546)))
	v1162 = v447
	v1163 = v448
	goto L19
L61:
	;
	goto L60
L62:
	;
	v661 = v27 + int32(256)
	v664 = *(*int64)(unsafe.Add(mBase, uint32(int32(1865752)-v617)))
	v665 = int64(0)
	v671 = int64(32)
	v672 = int64(base.Ui64(v622) >> (uint(v671) % 64))
	v674 = int64(base.Ui64(v664) >> (uint(v671) % 64))
	v677 = int64(4294967295)
	v678 = v622 & v677
	v680 = v664 & v677
	v681 = v678 * v680
	v685 = int64(base.Ui64(v681)>>(uint(v671)%64)) + v678*v674
	v692 = v680*v672 + v685&v677
	*(*int64)(unsafe.Add(mBase, uint32(v661)+8)) = v664*v665 + v665*v622 + v672*v674 + int64(base.Ui64(v685)>>(uint(v671)%64)) + int64(base.Ui64(v692)>>(uint(v671)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v661))) = v681&v677 | v692<<(uint(v671)%64)
	goto L63
L63:
	;
	v704 = v27 + int32(240)
	v705 = *(*int64)(unsafe.Add(mBase, uint32(v27)+280))
	v706 = *(*int64)(unsafe.Add(mBase, uint32(v27)+256))
	v707 = v705 + v706
	v708 = *(*int64)(unsafe.Add(mBase, uint32(v27)+264))
	v711 = v708 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v707) < base.Ui64(v705)))
	v716 = v614 - int32(base.Ui32(v615*int32(-1217359))>>(uint(int32(19))%32))
	v718 = v716 + int32(56)
	if v718&int32(64) != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v743 = v27 + int32(176)
	v744 = int64(0)
	v748 = v72 + base.I64_extend_i32_s(v77^int32(-1))
	v754 = int64(32)
	v755 = int64(base.Ui64(v748) >> (uint(v754) % 64))
	v757 = int64(base.Ui64(v619) >> (uint(v754) % 64))
	v760 = int64(4294967295)
	v761 = v748 & v760
	v763 = v619 & v760
	v764 = v761 * v763
	v768 = int64(base.Ui64(v764)>>(uint(v754)%64)) + v761*v757
	v775 = v763*v755 + v768&v760
	*(*int64)(unsafe.Add(mBase, uint32(v743)+8)) = v619*v744 + v744*v748 + v755*v757 + int64(base.Ui64(v768)>>(uint(v754)%64)) + int64(base.Ui64(v775)>>(uint(v754)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v743))) = v764&v760 | v775<<(uint(v754)%64)
	goto L70
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v704))) = v737
	*(*int64)(unsafe.Add(mBase, uint32(v704)+8)) = v738
	goto L64
L66:
	;
	v737 = int64(base.Ui64(v711) >> (uint(base.I64_extend_i32_u(v716+int32(-8))) % 64))
	v738 = int64(0)
	goto L65
L67:
	;
	goto L68
L68:
	;
	if v718 == int32(0) {
		v737 = v707
		v738 = v711
		goto L65
	} else {
		goto L69
	}
L69:
	;
	v733 = base.I64_extend_i32_u(v718)
	v737 = v711<<(uint(base.I64_extend_i32_u(int32(64)-v718))%64) | int64(base.Ui64(v707)>>(uint(v733)%64))
	v738 = int64(base.Ui64(v711) >> (uint(v733) % 64))
	goto L65
L70:
	;
	v787 = v27 + int32(160)
	v788 = int64(0)
	v794 = int64(32)
	v795 = int64(base.Ui64(v748) >> (uint(v794) % 64))
	v797 = int64(base.Ui64(v664) >> (uint(v794) % 64))
	v800 = int64(4294967295)
	v801 = v748 & v800
	v803 = v664 & v800
	v804 = v801 * v803
	v808 = int64(base.Ui64(v804)>>(uint(v794)%64)) + v801*v797
	v815 = v803*v795 + v808&v800
	*(*int64)(unsafe.Add(mBase, uint32(v787)+8)) = v664*v788 + v788*v748 + v795*v797 + int64(base.Ui64(v808)>>(uint(v794)%64)) + int64(base.Ui64(v815)>>(uint(v794)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v787))) = v804&v800 | v815<<(uint(v794)%64)
	goto L71
L71:
	;
	v827 = v27 + int32(144)
	v828 = *(*int64)(unsafe.Add(mBase, uint32(v27)+184))
	v829 = *(*int64)(unsafe.Add(mBase, uint32(v27)+160))
	v830 = v828 + v829
	v831 = *(*int64)(unsafe.Add(mBase, uint32(v27)+168))
	v834 = v831 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v830) < base.Ui64(v828)))
	if v718&int32(64) != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v859 = v27 + int32(224)
	v860 = int64(0)
	v866 = int64(32)
	v867 = int64(base.Ui64(v72) >> (uint(v866) % 64))
	v869 = int64(base.Ui64(v619) >> (uint(v866) % 64))
	v872 = int64(4294967295)
	v873 = v72 & v872
	v875 = v619 & v872
	v876 = v873 * v875
	v880 = int64(base.Ui64(v876)>>(uint(v866)%64)) + v873*v869
	v887 = v875*v867 + v880&v872
	*(*int64)(unsafe.Add(mBase, uint32(v859)+8)) = v619*v860 + v860*v72 + v867*v869 + int64(base.Ui64(v880)>>(uint(v866)%64)) + int64(base.Ui64(v887)>>(uint(v866)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v859))) = v876&v872 | v887<<(uint(v866)%64)
	goto L78
L73:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v827))) = v853
	*(*int64)(unsafe.Add(mBase, uint32(v827)+8)) = v854
	goto L72
L74:
	;
	v853 = int64(base.Ui64(v834) >> (uint(base.I64_extend_i32_u(v716+int32(-8))) % 64))
	v854 = int64(0)
	goto L73
L75:
	;
	goto L76
L76:
	;
	if v718 == int32(0) {
		v853 = v830
		v854 = v834
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v849 = base.I64_extend_i32_u(v718)
	v853 = v834<<(uint(base.I64_extend_i32_u(int32(64)-v718))%64) | int64(base.Ui64(v830)>>(uint(v849)%64))
	v854 = int64(base.Ui64(v834) >> (uint(v849) % 64))
	goto L73
L78:
	;
	v899 = v27 + int32(208)
	v900 = int64(0)
	v906 = int64(32)
	v907 = int64(base.Ui64(v72) >> (uint(v906) % 64))
	v909 = int64(base.Ui64(v664) >> (uint(v906) % 64))
	v912 = int64(4294967295)
	v913 = v72 & v912
	v915 = v664 & v912
	v916 = v913 * v915
	v920 = int64(base.Ui64(v916)>>(uint(v906)%64)) + v913*v909
	v927 = v915*v907 + v920&v912
	*(*int64)(unsafe.Add(mBase, uint32(v899)+8)) = v664*v900 + v900*v72 + v907*v909 + int64(base.Ui64(v920)>>(uint(v906)%64)) + int64(base.Ui64(v927)>>(uint(v906)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v899))) = v916&v912 | v927<<(uint(v906)%64)
	goto L79
L79:
	;
	v939 = v27 + int32(192)
	v940 = *(*int64)(unsafe.Add(mBase, uint32(v27)+232))
	v941 = *(*int64)(unsafe.Add(mBase, uint32(v27)+208))
	v942 = v940 + v941
	v943 = *(*int64)(unsafe.Add(mBase, uint32(v27)+216))
	v946 = v943 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v942) < base.Ui64(v940)))
	if v718&int32(64) != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	v970 = *(*int64)(unsafe.Add(mBase, uint32(v27)+192))
	v971 = *(*int64)(unsafe.Add(mBase, uint32(v27)+144))
	v972 = *(*int64)(unsafe.Add(mBase, uint32(v27)+240))
	if base.Ui32(v614) <= base.Ui32(int32(1)) {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v939))) = v965
	*(*int64)(unsafe.Add(mBase, uint32(v939)+8)) = v966
	goto L80
L82:
	;
	v965 = int64(base.Ui64(v946) >> (uint(base.I64_extend_i32_u(v716+int32(-8))) % 64))
	v966 = int64(0)
	goto L81
L83:
	;
	goto L84
L84:
	;
	if v718 == int32(0) {
		v965 = v942
		v966 = v946
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v961 = base.I64_extend_i32_u(v718)
	v965 = v946<<(uint(base.I64_extend_i32_u(int32(64)-v718))%64) | int64(base.Ui64(v942)>>(uint(v961)%64))
	v966 = int64(base.Ui64(v946) >> (uint(v961) % 64))
	goto L81
L86:
	;
	v996 = v615
	v1005 = v972 - int64(1)
	v1007 = v970
	v1008 = v971
	goto L20
L87:
	;
	goto L88
L88:
	;
	if base.Ui32(int32(62)) < base.Ui32(v614) {
		v1151 = v615
		v1160 = v972
		v1162 = v970
		v1163 = v971
		goto L19
	} else {
		goto L89
	}
L89:
	;
	v979 = int64(-1)
	if v72&(v979<<(uint(base.I64_extend_i32_u(v614-int32(1)))%64)^v979) != int64(0) {
		v1151 = v615
		v1160 = v972
		v1162 = v970
		v1163 = v971
		goto L19
	} else {
		goto L90
	}
L90:
	;
	v996 = v615
	v1005 = v972
	v1007 = v970
	v1008 = v971
	goto L20
L91:
	;
	v1040 = int64(4294967295)
	v1041 = v1032 & v1040
	v1042 = int64(3435973837)
	v1044 = int64(32)
	v1047 = int64(base.Ui64(v1032) >> (uint(v1044) % 64))
	v1050 = int64(base.Ui64(v1041*v1042)>>(uint(v1044)%64)) + v1047*v1042
	v1053 = int64(3435973836)
	v1058 = int64(4294967292)
	v1064 = int64(3)
	v1065 = int64(base.Ui64(int64(base.Ui64(v1050)>>(uint(v1044)%64))+v1047*v1053+int64(base.Ui64(v1041*v1053+v1050&v1058)>>(uint(v1044)%64))) >> (uint(v1064) % 64))
	v1067 = int64(base.Ui64(v1035) >> (uint(v1044) % 64))
	v1071 = v1035 & v1040
	v1076 = v1067*v1042 + int64(base.Ui64(v1071*v1042)>>(uint(v1044)%64))
	v1091 = int64(base.Ui64(int64(base.Ui64(v1076)>>(uint(v1044)%64))+v1067*v1053+int64(base.Ui64(v1071*v1053+v1076&v1058)>>(uint(v1044)%64))) >> (uint(v1064) % 64))
	if base.Ui64(v1091) < base.Ui64(v1065) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v1134 = v1019 & int32(255)
	v1384 = v1018
	v1389 = v996
	v1400 = v1034
	v1406 = (base.I32_wrap_i64(v1034)|(v1021^int32(-1)|base.B2i32(v1134 != int32(5))))&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v1134)) | base.B2i32(v1034 == v1035)
	goto L18
L93:
	;
	v1100 = int64(32)
	v1101 = int64(base.Ui64(v1034) >> (uint(v1100) % 64))
	v1102 = int64(3435973837)
	v1105 = v1034 & int64(4294967295)
	v1110 = v1101*v1102 + int64(base.Ui64(v1105*v1102)>>(uint(v1100)%64))
	v1113 = int64(3435973836)
	v1125 = int64(base.Ui64(int64(base.Ui64(v1110)>>(uint(v1100)%64))+v1101*v1113+int64(base.Ui64(v1105*v1113+v1110&int64(4294967292))>>(uint(v1100)%64))) >> (uint(int64(3)) % 64))
	__phi1018 = v1018 + int32(1)
	__phi1019 = base.I32_wrap_i64(v1125*int64(246) + v1034)
	__phi1021 = base.B2i32(v1019&int32(255) == int32(0)) & v1021
	__phi1032 = v1065
	__phi1034 = v1125
	__phi1035 = v1091
	v1018 = __phi1018
	v1019 = __phi1019
	v1021 = __phi1021
	v1032 = __phi1032
	v1034 = __phi1034
	v1035 = __phi1035
	goto L91
L94:
	;
	goto L95
L95:
	;
	goto L92
L96:
	;
	__phi1271 = v1268
	__phi1272 = v1262
	__phi1284 = v1263
	__phi1285 = v1264
	__phi1286 = v1265
	v1271 = __phi1271
	v1272 = __phi1272
	v1284 = __phi1284
	v1285 = __phi1285
	v1286 = __phi1286
	goto L100
L97:
	;
	v1262 = int32(0)
	v1263 = v1160
	v1264 = v1162
	v1265 = v1163
	v1268 = int32(0)
	goto L96
L98:
	;
	goto L99
L99:
	;
	v1228 = int64(base.Ui64(v1162) >> (uint(int64(34)) % 64))
	v1229 = int64(1546188227)
	v1231 = int64(2)
	v1233 = int64(4294967295)
	v1234 = int64(base.Ui64(v1162)>>(uint(v1231)%64)) & v1233
	v1237 = int64(32)
	v1239 = v1228*v1229 + int64(base.Ui64(v1234*v1229)>>(uint(v1237)%64))
	v1242 = int64(687194767)
	v1254 = int64(base.Ui64(int64(base.Ui64(v1239)>>(uint(v1237)%64))+v1228*v1242+int64(base.Ui64(v1234*v1242+v1239&v1233)>>(uint(v1237)%64))) >> (uint(v1231) % 64))
	v1262 = base.B2i32(base.Ui32(int32(49)) < base.Ui32(base.I32_wrap_i64(v1254*int64(4294967196)+v1162)))
	v1263 = v1196
	v1264 = v1254
	v1265 = v1224
	v1268 = int32(2)
	goto L96
L100:
	;
	v1293 = int64(4294967295)
	v1294 = v1284 & v1293
	v1295 = int64(3435973837)
	v1297 = int64(32)
	v1300 = int64(base.Ui64(v1284) >> (uint(v1297) % 64))
	v1303 = int64(base.Ui64(v1294*v1295)>>(uint(v1297)%64)) + v1300*v1295
	v1306 = int64(3435973836)
	v1311 = int64(4294967292)
	v1317 = int64(3)
	v1318 = int64(base.Ui64(int64(base.Ui64(v1303)>>(uint(v1297)%64))+v1300*v1306+int64(base.Ui64(v1294*v1306+v1303&v1311)>>(uint(v1297)%64))) >> (uint(v1317) % 64))
	v1320 = int64(base.Ui64(v1286) >> (uint(v1297) % 64))
	v1324 = v1286 & v1293
	v1329 = v1320*v1295 + int64(base.Ui64(v1324*v1295)>>(uint(v1297)%64))
	v1344 = int64(base.Ui64(int64(base.Ui64(v1329)>>(uint(v1297)%64))+v1320*v1306+int64(base.Ui64(v1324*v1306+v1329&v1311)>>(uint(v1297)%64))) >> (uint(v1317) % 64))
	if base.Ui64(v1344) < base.Ui64(v1318) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v1384 = v1271
	v1389 = v1151
	v1400 = v1285
	v1406 = base.B2i32(v1285 == v1286) | v1272
	goto L18
L102:
	;
	v1346 = int64(32)
	v1347 = int64(base.Ui64(v1285) >> (uint(v1346) % 64))
	v1348 = int64(3435973837)
	v1351 = v1285 & int64(4294967295)
	v1356 = v1347*v1348 + int64(base.Ui64(v1351*v1348)>>(uint(v1346)%64))
	v1359 = int64(3435973836)
	v1371 = int64(base.Ui64(int64(base.Ui64(v1356)>>(uint(v1346)%64))+v1347*v1359+int64(base.Ui64(v1351*v1359+v1356&int64(4294967292))>>(uint(v1346)%64))) >> (uint(int64(3)) % 64))
	__phi1271 = v1271 + int32(1)
	__phi1272 = base.B2i32(base.Ui32(int32(4)) < base.Ui32(base.I32_wrap_i64(v1371*int64(4294967286)+v1285)))
	__phi1284 = v1318
	__phi1285 = v1371
	__phi1286 = v1344
	v1271 = __phi1271
	v1272 = __phi1272
	v1284 = __phi1284
	v1285 = __phi1285
	v1286 = __phi1286
	goto L100
L103:
	;
	goto L104
L104:
	;
	goto L101
L105:
	;
	v1495 = v1407
	v1501 = v1411
	v1510 = int32(17)
	goto L10
L106:
	;
	if base.Ui64(int64(99999999999999)) < base.Ui64(v1430) {
		v1495 = v1424
		v1501 = v1430
		v1510 = int32(15)
		goto L10
	} else {
		goto L107
	}
L107:
	;
	if base.Ui64(int64(9999999999999)) < base.Ui64(v1430) {
		v1495 = v1424
		v1501 = v1430
		v1510 = int32(14)
		goto L10
	} else {
		goto L108
	}
L108:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v1430) {
		v1495 = v1424
		v1501 = v1430
		v1510 = int32(13)
		goto L10
	} else {
		goto L109
	}
L109:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v1430) {
		v1495 = v1424
		v1501 = v1430
		v1510 = int32(12)
		goto L10
	} else {
		goto L110
	}
L110:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v1430) {
		v1495 = v1424
		v1501 = v1430
		v1510 = int32(11)
		goto L10
	} else {
		goto L111
	}
L111:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v1430) {
		v1495 = v1424
		v1501 = v1430
		v1510 = int32(10)
		goto L10
	} else {
		goto L112
	}
L112:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v1430) {
		v1495 = v1424
		v1501 = v1430
		v1510 = int32(9)
		goto L10
	} else {
		goto L113
	}
L113:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v1430) {
		v1495 = v1424
		v1501 = v1430
		v1510 = int32(8)
		goto L10
	} else {
		goto L114
	}
L114:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v1430) {
		v1495 = v1424
		v1501 = v1430
		v1510 = int32(7)
		goto L10
	} else {
		goto L115
	}
L115:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v1430) {
		v1495 = v1424
		v1501 = v1430
		v1510 = int32(6)
		goto L10
	} else {
		goto L116
	}
L116:
	;
	if base.Ui64(int64(9999)) < base.Ui64(v1430) {
		v1495 = v1424
		v1501 = v1430
		v1510 = int32(5)
		goto L10
	} else {
		goto L117
	}
L117:
	;
	if base.Ui64(int64(999)) < base.Ui64(v1430) {
		v1495 = v1424
		v1501 = v1430
		v1510 = int32(4)
		goto L10
	} else {
		goto L118
	}
L118:
	;
	if base.Ui64(int64(99)) < base.Ui64(v1430) {
		v1495 = v1424
		v1501 = v1430
		v1510 = int32(3)
		goto L10
	} else {
		goto L119
	}
L119:
	;
	if base.Ui64(int64(9)) < base.Ui64(v1430) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v1485 = int32(2)
	goto L122
L121:
	;
	v1485 = int32(1)
	goto L122
L122:
	;
	v1495 = v1424
	v1501 = v1430
	v1510 = v1485
	goto L10
L123:
	;
	v1513 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v1513)
	v1516 = int32(1)
	goto L125
L124:
	;
	v1516 = v3
	goto L125
L125:
	;
	v1517 = v1495 + v1510
	if base.Ui32(v1517+int32(3)) <= base.Ui32(int32(18)) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v1522 = v1516 + l1
	if v1517 <= int32(0) {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	goto L128
L128:
	;
	if v1495 != 0 {
		goto L173
	} else {
		goto L174
	}
L129:
	;
	if base.Ui64(v1501) < base.Ui64(int64(4294967296)) {
		goto L135
	} else {
		goto L136
	}
L130:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1522))) = int64(3472328296227679792)
	v1537 = int32(2) - v1517
	goto L129
L131:
	;
	goto L132
L132:
	;
	if v1495 < int32(0) {
		v1537 = int32(1)
		goto L129
	} else {
		goto L133
	}
L133:
	;
	v1532 = int64(3472328296227680304)
	*(*int64)(unsafe.Add(mBase, uint32(v1522))) = v1532
	*(*int64)(unsafe.Add(mBase, uint32(v1522)+8)) = v1532
	v1537 = int32(0)
	goto L129
L134:
	;
	v1637 = base.I32_wrap_i64(v1634)
	if base.Ui32(v1637) < base.Ui32(int32(10000)) {
		goto L139
	} else {
		goto L140
	}
L135:
	;
	v1630 = int32(0)
	v1634 = v1501
	goto L134
L136:
	;
	goto L137
L137:
	;
	v1542 = v1522 + v1537 + v1510
	v1543 = int32(8)
	v1545 = int64(32)
	v1546 = int64(base.Ui64(v1501) >> (uint(v1545) % 64))
	v1547 = int64(2221002493)
	v1549 = int64(4294967295)
	v1550 = v1501 & v1549
	v1555 = v1546*v1547 + int64(base.Ui64(v1550*v1547)>>(uint(v1545)%64))
	v1558 = int64(2882303761)
	v1570 = int64(base.Ui64(int64(base.Ui64(v1555)>>(uint(v1545)%64))+v1546*v1558+int64(base.Ui64(v1550*v1558+v1555&v1549)>>(uint(v1545)%64))) >> (uint(int64(26)) % 64))
	v1574 = base.I32_wrap_i64(v1570*int64(4194967296) + v1501)
	v1575 = int32(10000)
	v1576 = base.I32_div_u_s(v1574, v1575)
	v1578 = base.I32_rem_u_s(v1576, v1575)
	v1579 = int32(100)
	v1580 = base.I32_div_u_s(v1578, v1579)
	v1581 = int32(1)
	v1585 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1580<<(uint(v1581)%32))+uint32(_consts[1278]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1542-v1543))) = uint16(v1585)
	v1591 = v1574 - v1576*v1575
	v1592 = int32(65535)
	v1595 = base.I32_div_u_s(v1591&v1592, v1579)
	v1600 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1595<<(uint(v1581)%32))+uint32(_consts[1278]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1542-int32(4)))) = uint16(v1600)
	v1613 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1578-v1580*v1579)&v1592<<(uint(v1581)%32))+uint32(_consts[1278]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1542-int32(6)))) = uint16(v1613)
	v1626 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1591-v1595*v1579)&v1592<<(uint(v1581)%32))+uint32(_consts[1278]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1542-int32(2)))) = uint16(v1626)
	v1630 = v1543
	v1634 = v1570
	goto L134
L138:
	;
	if base.Ui32(v1702) < base.Ui32(int32(100)) {
		goto L146
	} else {
		goto L147
	}
L139:
	;
	v1700 = v1630
	v1702 = v1637
	goto L138
L140:
	;
	goto L141
L141:
	;
	v1644 = v1637
	v1645 = v1630
	goto L142
L142:
	;
	v1666 = v1522 + v1537 + v1510 - v1645
	v1667 = int32(4)
	v1670 = base.I32_div_u_s(v1644, int32(10000))
	v1673 = v1670*int32(-10000) + v1644
	v1674 = int32(100)
	v1675 = base.I32_div_u_s(v1673, v1674)
	v1676 = int32(1)
	v1680 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1675<<(uint(v1676)%32))+uint32(_consts[1278]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1666-v1667))) = uint16(v1680)
	v1691 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1673-v1675*v1674)<<(uint(v1676)%32))+uint32(_consts[1278]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1666-int32(2)))) = uint16(v1691)
	v1694 = v1645 + v1667
	if base.Ui32(int32(99999999)) < base.Ui32(v1644) {
		v1644 = v1670
		v1645 = v1694
		goto L142
	} else {
		goto L144
	}
L143:
	;
	v1700 = v1694
	v1702 = v1670
	goto L138
L144:
	;
	goto L143
L145:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1745) {
		goto L150
	} else {
		goto L151
	}
L146:
	;
	v1745 = v1702
	v1746 = v1700
	goto L145
L147:
	;
	goto L148
L148:
	;
	v1726 = int32(2)
	v1728 = int32(65535)
	v1730 = int32(100)
	v1731 = base.I32_div_u_s(v1702&v1728, v1730)
	v1741 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1702-v1731*v1730)&v1728<<(uint(int32(1))%32))+uint32(_consts[1278]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1522+v1537+v1510-v1700-v1726))) = uint16(v1741)
	v1745 = v1731
	v1746 = v1700 | v1726
	goto L145
L149:
	;
	v1764 = int32(1)
	if v1537 == v1764 {
		goto L154
	} else {
		goto L155
	}
L150:
	;
	v1758 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1745<<(uint(int32(1))%32))+uint32(_consts[1278]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1522+v1537+v1510-v1746-int32(2)))) = uint16(v1758)
	goto L149
L151:
	;
	goto L152
L152:
	;
	v1762 = v1745 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1522+v1537))) = uint8(v1762)
	goto L149
L153:
	;
	v2236 = v1815 + base.I32_wrap_i64(int64(base.Ui64(v29)>>(uint(int64(63))%64)))
	goto L1
L154:
	;
	if v1517&int32(8) != 0 {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	goto L156
L156:
	;
	if v1495 < int32(0) {
		goto L169
	} else {
		goto L170
	}
L157:
	;
	v1769 = *(*int64)(unsafe.Add(mBase, uint32(v1522)+1))
	*(*int64)(unsafe.Add(mBase, uint32(v1522))) = v1769
	v1772 = int32(9)
	goto L159
L158:
	;
	v1772 = v1764
	goto L159
L159:
	;
	if v1517&int32(4) != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v1775 = v1772 + v1522
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(v1775)))
	*(*int32)(unsafe.Add(mBase, uint32(v1775-int32(1)))) = v1778
	v1782 = v1772 | int32(4)
	goto L162
L161:
	;
	v1782 = v1772
	goto L162
L162:
	;
	if v1517&int32(2) != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v1786 = v1782 + v1522
	v1789 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1786))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1786-int32(1)))) = uint16(v1789)
	v1793 = v1782 + int32(2)
	goto L165
L164:
	;
	v1793 = v1782
	goto L165
L165:
	;
	if v1517&int32(1) != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v1797 = v1793 + v1522
	v1800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1797))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1797-int32(1)))) = uint8(v1800)
	goto L168
L167:
	;
	goto L168
L168:
	;
	v1804 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1522+v1517))) = uint8(v1804)
	v1815 = v1510 + int32(1)
	goto L153
L169:
	;
	v1812 = int32(2) - v1495
	goto L171
L170:
	;
	v1812 = v1517
	goto L171
L171:
	;
	v1815 = v1812
	goto L153
L172:
	;
	if base.Ui64(v1898) < base.Ui64(int64(4294967296)) {
		goto L181
	} else {
		goto L182
	}
L173:
	;
	v1891 = v1510
	v1898 = v1501
	goto L172
L174:
	;
	goto L175
L175:
	;
	v1830 = v1510
	v1835 = v1501
	goto L176
L176:
	;
	if base.I32_wrap_i64(v1835)&int32(1) != 0 {
		v1891 = v1830
		v1898 = v1835
		goto L172
	} else {
		goto L178
	}
L177:
	;
	v1891 = v1830
	v1898 = v1835
	goto L172
L178:
	;
	v1849 = int64(32)
	v1850 = int64(base.Ui64(v1835) >> (uint(v1849) % 64))
	v1851 = int64(3435973837)
	v1853 = int64(4294967294)
	v1854 = v1835 & v1853
	v1859 = v1850*v1851 + int64(base.Ui64(v1854*v1851)>>(uint(v1849)%64))
	v1862 = int64(3435973836)
	v1874 = int64(base.Ui64(int64(base.Ui64(v1859)>>(uint(v1849)%64))+v1850*v1862+int64(base.Ui64(v1854*v1862+v1859&int64(4294967288))>>(uint(v1849)%64))) >> (uint(int64(3)) % 64))
	if (v1835+v1874*int64(4294967286))&v1853 == int64(0) {
		v1830 = v1830 - int32(1)
		v1835 = v1874
		goto L176
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	v2005 = base.I32_wrap_i64(v2002)
	if base.Ui32(v2005) < base.Ui32(int32(10000)) {
		goto L185
	} else {
		goto L186
	}
L181:
	;
	v1998 = int32(0)
	v2002 = v1898
	goto L180
L182:
	;
	goto L183
L183:
	;
	v1910 = v1516 + l1 + v1891
	v1913 = int64(32)
	v1914 = int64(base.Ui64(v1898) >> (uint(v1913) % 64))
	v1915 = int64(2221002493)
	v1917 = int64(4294967295)
	v1918 = v1898 & v1917
	v1923 = v1914*v1915 + int64(base.Ui64(v1918*v1915)>>(uint(v1913)%64))
	v1926 = int64(2882303761)
	v1938 = int64(base.Ui64(int64(base.Ui64(v1923)>>(uint(v1913)%64))+v1914*v1926+int64(base.Ui64(v1918*v1926+v1923&v1917)>>(uint(v1913)%64))) >> (uint(int64(26)) % 64))
	v1942 = base.I32_wrap_i64(v1938*int64(4194967296) + v1898)
	v1943 = int32(10000)
	v1944 = base.I32_div_u_s(v1942, v1943)
	v1946 = base.I32_rem_u_s(v1944, v1943)
	v1947 = int32(100)
	v1948 = base.I32_div_u_s(v1946, v1947)
	v1949 = int32(1)
	v1953 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1948<<(uint(v1949)%32))+uint32(_consts[1278]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1910-int32(7)))) = uint16(v1953)
	v1959 = v1942 - v1944*v1943
	v1960 = int32(65535)
	v1963 = base.I32_div_u_s(v1959&v1960, v1947)
	v1968 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1963<<(uint(v1949)%32))+uint32(_consts[1278]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1910-int32(3)))) = uint16(v1968)
	v1981 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1946-v1948*v1947)&v1960<<(uint(v1949)%32))+uint32(_consts[1278]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1910-int32(5)))) = uint16(v1981)
	v1994 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1959-v1963*v1947)&v1960<<(uint(v1949)%32))+uint32(_consts[1278]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1910-v1949))) = uint16(v1994)
	v1998 = int32(8)
	v2002 = v1938
	goto L180
L184:
	;
	if base.Ui32(v2070) < base.Ui32(int32(100)) {
		goto L192
	} else {
		goto L193
	}
L185:
	;
	v2068 = v1998
	v2070 = v2005
	goto L184
L186:
	;
	goto L187
L187:
	;
	v2012 = v2005
	v2013 = v1998
	goto L188
L188:
	;
	v2034 = v1516 + l1 + v1891 - v2013
	v2038 = base.I32_div_u_s(v2012, int32(10000))
	v2041 = v2038*int32(-10000) + v2012
	v2042 = int32(100)
	v2043 = base.I32_div_u_s(v2041, v2042)
	v2044 = int32(1)
	v2048 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2043<<(uint(v2044)%32))+uint32(_consts[1278]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2034-int32(3)))) = uint16(v2048)
	v2059 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2041-v2043*v2042)<<(uint(v2044)%32))+uint32(_consts[1278]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v2034-v2044))) = uint16(v2059)
	v2062 = v2013 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v2012) {
		v2012 = v2038
		v2013 = v2062
		goto L188
	} else {
		goto L190
	}
L189:
	;
	v2068 = v2062
	v2070 = v2038
	goto L184
L190:
	;
	goto L189
L191:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v2113) {
		goto L196
	} else {
		goto L197
	}
L192:
	;
	v2113 = v2070
	v2114 = v2068
	goto L191
L193:
	;
	goto L194
L194:
	;
	v2096 = int32(65535)
	v2098 = int32(100)
	v2099 = base.I32_div_u_s(v2070&v2096, v2098)
	v2109 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v2070-v2099*v2098)&v2096<<(uint(int32(1))%32))+uint32(_consts[1278]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1516+l1+v1891+(v2068^int32(-1))))) = uint16(v2109)
	v2113 = v2099
	v2114 = v2068 | int32(2)
	goto L191
L195:
	;
	v2134 = v1517 - int32(1)
	v2135 = v1516 + l1
	*(*uint8)(unsafe.Add(mBase, uint32(v2135))) = uint8(v2132)
	if base.Ui32(int32(2)) <= base.Ui32(v1891) {
		goto L199
	} else {
		goto L200
	}
L196:
	;
	v2121 = v2113 << (uint(int32(1)) % 32)
	v2124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2121)+uint32(_consts[1279]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1+(v1516+v1891-v2114)))) = uint8(v2124)
	v2128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2121)+uint32(_consts[1278]))))
	v2132 = v2128
	goto L195
L197:
	;
	goto L198
L198:
	;
	v2132 = v2113 | int32(48)
	goto L195
L199:
	;
	v2139 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v2135)+1)) = uint8(v2139)
	v2144 = v1891 + int32(1)
	goto L201
L200:
	;
	v2144 = int32(1)
	goto L201
L201:
	;
	v2145 = v2144 + v1516
	v2146 = l1 + v2145
	v2147 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v2146))) = uint8(v2147)
	v2152 = base.B2i32(v2134 < int32(0))
	if v2134 < int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v2153 = int32(45)
	goto L204
L203:
	;
	v2153 = int32(43)
	goto L204
L204:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2146)+1)) = uint8(v2153)
	v2156 = v2145 + int32(2)
	if v2134 < int32(0) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v2159 = int32(1) - v1517
	goto L207
L206:
	;
	v2159 = v2134
	goto L207
L207:
	;
	if int32(100) <= v2159 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v2163 = int32(10)
	v2164 = base.I32_div_u_s(v2159, v2163)
	v2169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2164<<(uint(int32(1))%32))+uint32(_consts[1278]))))
	*(*uint16)(unsafe.Add(mBase, uint32(l1+v2156))) = uint16(v2169)
	v2175 = v2159 - v2164*v2163 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2146)+4)) = uint8(v2175)
	v2236 = v2145 + int32(5)
	goto L1
L209:
	;
	goto L210
L210:
	;
	v2184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2159<<(uint(int32(1))%32))+uint32(_consts[1278]))))
	*(*uint16)(unsafe.Add(mBase, uint32(l1+v2156))) = uint16(v2184)
	v2236 = v2145 + int32(4)
	goto L1
L211:
	;
	v2190 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v2190)
	goto L213
L212:
	;
	goto L213
L213:
	;
	v2195 = l1 + base.I32_wrap_i64(int64(base.Ui64(v29)>>(uint(int64(63))%64)))
	if v36 == int32(2047) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2195))) = int64(8751735898823355977)
	if v29 < int64(0) {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	goto L216
L216:
	;
	v2205 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2195))) = uint8(v2205)
	if v29 < int64(0) {
		goto L220
	} else {
		goto L221
	}
L217:
	;
	v2204 = int32(9)
	goto L219
L218:
	;
	v2204 = int32(8)
	goto L219
L219:
	;
	v2236 = v2204
	goto L1
L220:
	;
	v2211 = int32(2)
	goto L222
L221:
	;
	v2211 = int32(1)
	goto L222
L222:
	;
	v2236 = v2211
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
		*(*int32)(unsafe.Add(mBase, _consts[158])) = int32(0)
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
					F_errmsg(m, int32(418208), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(512946), int32(1953), int32(286385))
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
	var v63 float64
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v77 float64
	_ = v77
	var v78 int32
	_ = v78
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v89 int32
	_ = v89
	var v91 float64
	_ = v91
	var v92 float64
	_ = v92
	var v108 float64
	_ = v108
	var v111 float64
	_ = v111
	var v117 float64
	_ = v117
	var v126 float64
	_ = v126
	var v141 float64
	_ = v141
	var v150 float64
	_ = v150
	var v155 float64
	_ = v155
	var v162 float64
	_ = v162
	var v165 float64
	_ = v165
	var v171 float64
	_ = v171
	var v181 float64
	_ = v181
	var v183 float64
	_ = v183
	var v195 float64
	_ = v195
	var v213 float64
	_ = v213
	var v217 float64
	_ = v217
	var v222 float64
	_ = v222
	var v225 float64
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	*(*int32)(unsafe.Add(mBase, _consts[158])) = int32(0)
	v12 = base.F64_copysign(float64(0.5), v7)
	v13 = base.F64_abs(v7)
	v14 = base.I64_reinterpret_f64(v13)
	if base.Ui64(v14) <= base.Ui64(int64(4649454526309335039)) {
		v23 = base.I64_reinterpret_f64(v13)
		v28 = base.I32_wrap_i64(int64(base.Ui64(v23)>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(int32(1078159482)) <= base.Ui32(v28) {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v13)&int64(9223372036854775807)) {
				v183 = v13
				v195 = v183
			} else {
				if v23 < int64(0) {
					v195 = float64(-1)
				} else {
					if base.F64_gt(v13, float64(709.782712893384)) == int32(0) {
						v63 = base.F64_add(base.F64_mul(v13, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v13))
						if base.F64_lt(base.F64_abs(v63), float64(2.147483648e+09)) != 0 {
							v67 = base.I32_trunc_f64_s(v63)
							v69 = v67
						} else {
							v69 = int32(-2147483648)
						}
						v70 = base.F64_convert_i32_s(v69)
						v77 = base.F64_mul(v70, float64(1.9082149292705877e-10))
						v78 = v69
						v79 = base.F64_add(v13, base.F64_mul(v70, float64(-0.6931471803691238)))
						v80 = base.F64_sub(v79, v77)
						v86 = v80
						v88 = base.F64_sub(base.F64_sub(v79, v80), v77)
						v89 = v78
						v91 = base.F64_mul(v86, float64(0.5))
						v92 = base.F64_mul(v86, v91)
						v108 = base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
						v111 = base.F64_sub(float64(3), base.F64_mul(v108, v91))
						v117 = base.F64_mul(v92, base.F64_div(base.F64_sub(v108, v111), base.F64_sub(float64(6), base.F64_mul(v86, v111))))
						if v89 == int32(0) {
							v195 = base.F64_sub(v86, base.F64_sub(base.F64_mul(v86, v117), v92))
						} else {
							v126 = base.F64_sub(base.F64_sub(base.F64_mul(v86, base.F64_sub(v117, v88)), v88), v92)
							switch v89 + int32(1) {
							case 0:
								v195 = base.F64_add(base.F64_mul(base.F64_sub(v86, v126), float64(0.5)), float64(-0.5))
							default:
								v150 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v89+int32(1023)) << (uint(int64(52)) % 64))
								if base.Ui32(int32(57)) <= base.Ui32(v89) {
									v155 = base.F64_add(base.F64_sub(v86, v126), float64(1))
									if v89 == int32(1024) {
										v162 = base.F64_mul(base.F64_add(v155, v155), float64(8.98846567431158e+307))
									} else {
										v162 = base.F64_mul(v155, v150)
									}
									v195 = base.F64_add(v162, float64(-1))
								} else {
									v165 = float64(1)
									v171 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v89) << (uint(int64(52)) % 64))
									if base.Ui32(v89) <= base.Ui32(int32(19)) {
										v181 = base.F64_add(base.F64_sub(v165, v171), base.F64_sub(v86, v126))
									} else {
										v181 = base.F64_add(base.F64_sub(v86, base.F64_add(v126, v171)), v165)
									}
									v183 = base.F64_mul(v181, v150)
									v195 = v183
								}
							case 2:
								if base.F64_lt(v86, float64(-0.25)) != 0 {
									v195 = base.F64_mul(base.F64_sub(v126, base.F64_add(v86, float64(0.5))), float64(-2))
								} else {
									v141 = base.F64_sub(v86, v126)
									v195 = base.F64_add(base.F64_add(v141, v141), float64(1))
								}
							}
						}
					} else {
						v195 = base.F64_mul(v13, float64(8.98846567431158e+307))
					}
				}
			}
		} else {
			if base.Ui32(v28) < base.Ui32(int32(1071001155)) {
				if base.Ui32(v28) < base.Ui32(int32(1016070144)) {
					v183 = v13
					v195 = v183
				} else {
					v86 = v13
					v88 = float64(0)
					v89 = int32(0)
					v91 = base.F64_mul(v86, float64(0.5))
					v92 = base.F64_mul(v86, v91)
					v108 = base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
					v111 = base.F64_sub(float64(3), base.F64_mul(v108, v91))
					v117 = base.F64_mul(v92, base.F64_div(base.F64_sub(v108, v111), base.F64_sub(float64(6), base.F64_mul(v86, v111))))
					if v89 == int32(0) {
						v195 = base.F64_sub(v86, base.F64_sub(base.F64_mul(v86, v117), v92))
					} else {
						v126 = base.F64_sub(base.F64_sub(base.F64_mul(v86, base.F64_sub(v117, v88)), v88), v92)
						switch v89 + int32(1) {
						case 0:
							v195 = base.F64_add(base.F64_mul(base.F64_sub(v86, v126), float64(0.5)), float64(-0.5))
						default:
							v150 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v89+int32(1023)) << (uint(int64(52)) % 64))
							if base.Ui32(int32(57)) <= base.Ui32(v89) {
								v155 = base.F64_add(base.F64_sub(v86, v126), float64(1))
								if v89 == int32(1024) {
									v162 = base.F64_mul(base.F64_add(v155, v155), float64(8.98846567431158e+307))
								} else {
									v162 = base.F64_mul(v155, v150)
								}
								v195 = base.F64_add(v162, float64(-1))
							} else {
								v165 = float64(1)
								v171 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v89) << (uint(int64(52)) % 64))
								if base.Ui32(v89) <= base.Ui32(int32(19)) {
									v181 = base.F64_add(base.F64_sub(v165, v171), base.F64_sub(v86, v126))
								} else {
									v181 = base.F64_add(base.F64_sub(v86, base.F64_add(v126, v171)), v165)
								}
								v183 = base.F64_mul(v181, v150)
								v195 = v183
							}
						case 2:
							if base.F64_lt(v86, float64(-0.25)) != 0 {
								v195 = base.F64_mul(base.F64_sub(v126, base.F64_add(v86, float64(0.5))), float64(-2))
							} else {
								v141 = base.F64_sub(v86, v126)
								v195 = base.F64_add(base.F64_add(v141, v141), float64(1))
							}
						}
					}
				}
			} else {
				if base.Ui32(int32(1072734897)) < base.Ui32(v28) {
					v63 = base.F64_add(base.F64_mul(v13, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v13))
					if base.F64_lt(base.F64_abs(v63), float64(2.147483648e+09)) != 0 {
						v67 = base.I32_trunc_f64_s(v63)
						v69 = v67
					} else {
						v69 = int32(-2147483648)
					}
					v70 = base.F64_convert_i32_s(v69)
					v77 = base.F64_mul(v70, float64(1.9082149292705877e-10))
					v78 = v69
					v79 = base.F64_add(v13, base.F64_mul(v70, float64(-0.6931471803691238)))
				} else {
					if int64(0) <= v23 {
						v77 = float64(1.9082149292705877e-10)
						v78 = int32(1)
						v79 = base.F64_add(v13, float64(-0.6931471803691238))
					} else {
						v77 = float64(-1.9082149292705877e-10)
						v78 = int32(-1)
						v79 = base.F64_add(v13, float64(0.6931471803691238))
					}
				}
				v80 = base.F64_sub(v79, v77)
				v86 = v80
				v88 = base.F64_sub(base.F64_sub(v79, v80), v77)
				v89 = v78
				v91 = base.F64_mul(v86, float64(0.5))
				v92 = base.F64_mul(v86, v91)
				v108 = base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
				v111 = base.F64_sub(float64(3), base.F64_mul(v108, v91))
				v117 = base.F64_mul(v92, base.F64_div(base.F64_sub(v108, v111), base.F64_sub(float64(6), base.F64_mul(v86, v111))))
				if v89 == int32(0) {
					v195 = base.F64_sub(v86, base.F64_sub(base.F64_mul(v86, v117), v92))
				} else {
					v126 = base.F64_sub(base.F64_sub(base.F64_mul(v86, base.F64_sub(v117, v88)), v88), v92)
					switch v89 + int32(1) {
					case 0:
						v195 = base.F64_add(base.F64_mul(base.F64_sub(v86, v126), float64(0.5)), float64(-0.5))
					default:
						v150 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v89+int32(1023)) << (uint(int64(52)) % 64))
						if base.Ui32(int32(57)) <= base.Ui32(v89) {
							v155 = base.F64_add(base.F64_sub(v86, v126), float64(1))
							if v89 == int32(1024) {
								v162 = base.F64_mul(base.F64_add(v155, v155), float64(8.98846567431158e+307))
							} else {
								v162 = base.F64_mul(v155, v150)
							}
							v195 = base.F64_add(v162, float64(-1))
						} else {
							v165 = float64(1)
							v171 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v89) << (uint(int64(52)) % 64))
							if base.Ui32(v89) <= base.Ui32(int32(19)) {
								v181 = base.F64_add(base.F64_sub(v165, v171), base.F64_sub(v86, v126))
							} else {
								v181 = base.F64_add(base.F64_sub(v86, base.F64_add(v126, v171)), v165)
							}
							v183 = base.F64_mul(v181, v150)
							v195 = v183
						}
					case 2:
						if base.F64_lt(v86, float64(-0.25)) != 0 {
							v195 = base.F64_mul(base.F64_sub(v126, base.F64_add(v86, float64(0.5))), float64(-2))
						} else {
							v141 = base.F64_sub(v86, v126)
							v195 = base.F64_add(base.F64_add(v141, v141), float64(1))
						}
					}
				}
			}
		}
		if base.Ui64(v14) <= base.Ui64(int64(4607182418800017407)) {
			if base.Ui64(v14) < base.Ui64(int64(4490088828488384512)) {
				v222 = v7
				v225 = v222
			} else {
				v225 = base.F64_mul(v12, base.F64_sub(base.F64_add(v195, v195), base.F64_div(base.F64_mul(v195, v195), base.F64_add(v195, float64(1)))))
			}
		} else {
			v225 = base.F64_mul(v12, base.F64_add(v195, base.F64_div(v195, base.F64_add(v195, float64(1)))))
		}
	} else {
		v213 = float64(2.247116418577895e+307)
		v217 = F_exp(m, base.F64_add(v13, float64(-1416.0996898839683)))
		mBase = m.M
		v222 = base.F64_mul(base.F64_mul(base.F64_mul(base.F64_add(v12, v12), v213), v217), v213)
		v225 = v222
	}
	v226 = F_Float8GetDatum(m, v225)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		return int32(0)
	} else {
		return v226
	}
}
func F_dtoi2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 float64
	_ = v4
	var v5 float64
	_ = v5
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(v3)))
	v5 = base.F64_nearest(v4)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(418110), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(512946), int32(1254), int32(582762))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
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
		if base.F64_ge(v5, float64(-32768)) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(418110), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(512946), int32(1254), int32(582762))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
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
			if base.F64_lt(v5, float64(32768)) == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(418110), int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(512946), int32(1254), int32(582762))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
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
				if base.F64_lt(base.F64_abs(v5), float64(2.147483648e+09)) != 0 {
					v22 = base.I32_trunc_f64_s(v5)
					return v22
				} else {
					return int32(-2147483648)
				}
			}
		}
	}
}
func F_dtoi4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 float64
	_ = v4
	var v5 float64
	_ = v5
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(v3)))
	v5 = base.F64_nearest(v4)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(418422), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(512946), int32(1229), int32(580574))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
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
		if base.F64_ge(v5, float64(-2.147483648e+09)) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(418422), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(512946), int32(1229), int32(580574))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
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
			if base.F64_lt(v5, float64(2.147483648e+09)) == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(418422), int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(512946), int32(1229), int32(580574))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
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
				if base.F64_lt(base.F64_abs(v5), float64(2.147483648e+09)) != 0 {
					v22 = base.I32_trunc_f64_s(v5)
					return v22
				} else {
					return int32(-2147483648)
				}
			}
		}
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
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(v3)))
	v5 = base.F64_nearest(v4)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(418132), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(521324), int32(1312), int32(576888))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
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
		if base.F64_ge(v5, float64(-9.223372036854776e+18)) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(418132), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(521324), int32(1312), int32(576888))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
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
			if base.F64_lt(v5, float64(9.223372036854776e+18)) == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(418132), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(521324), int32(1312), int32(576888))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
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
				if base.F64_lt(base.F64_abs(v5), float64(9.223372036854776e+18)) != 0 {
					v22 = base.I64_trunc_f64_s(v5)
					v23 = F_Int64GetDatum(m, v22)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						return v23
					}
				} else {
					v29 = F_Int64GetDatum(m, int64(-9223372036854775807-1))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						return v29
					}
				}
			}
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
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
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
	return v338
L2:
	;
	return int32(0)
L3:
	;
	if v14 != 0 {
		v338 = int32(-1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[760]))
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
	v26 = *(*int32)(unsafe.Add(mBase, _consts[158]))
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
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _consts[194])))
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
		v338 = v29
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
	F_errmsg(m, int32(310478), v9+int32(16))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(520480), int32(808), int32(393362))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v338 = v29
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
	v68 = *(*int32)(unsafe.Add(mBase, _consts[158]))
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
	v64 = *(*int32)(unsafe.Add(mBase, _consts[158]))
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
	*(*int32)(unsafe.Add(mBase, _consts[158])) = v68
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
		v338 = v73
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
	F_errmsg(m, int32(311426), v9+int32(32))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(520480), int32(825), int32(393362))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v338 = v73
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
		v338 = v102
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
	F_errmsg(m, int32(311221), v9+int32(48))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(520480), int32(833), int32(393362))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v338 = v102
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
		v338 = v127
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
	F_errmsg(m, int32(309189), v9)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(520480), int32(844), int32(393362))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v338 = v127
	goto L1
L47:
	;
	if v150 != 0 {
		v338 = v127
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
	v270 = v9 - int32(-64)
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	if v274 != 0 {
		goto L82
	} else {
		goto L83
	}
L50:
	;
	v266 = F_strlen(m, v255)
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
	v259 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v256))) = uint8(v259)
	goto L50
L55:
	;
	v240 = v235
	v241 = v236
	v242 = v237
	goto L77
L56:
	;
	if v230 == int32(0) {
		v255 = v228
		v256 = v229
		goto L54
	} else {
		goto L76
	}
L57:
	;
	v228 = l1
	v229 = v153
	v230 = v160
	goto L56
L58:
	;
	goto L59
L59:
	;
	if l1&int32(3) == int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v197 == int32(0) {
		v255 = v194
		v256 = v195
		goto L54
	} else {
		goto L69
	}
L61:
	;
	v194 = l1
	v195 = v153
	v196 = v160
	v197 = int32(1)
	goto L60
L62:
	;
	goto L63
L63:
	;
	v173 = l1
	v174 = v153
	v175 = v160
	goto L64
L64:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	*(*uint8)(unsafe.Add(mBase, uint32(v174))) = uint8(v177)
	if v177 == int32(0) {
		v235 = v173
		v236 = v174
		v237 = v175
		goto L55
	} else {
		goto L66
	}
L65:
	;
	v194 = v188
	v195 = v182
	v196 = v184
	v197 = v186
	goto L60
L66:
	;
	v181 = int32(1)
	v182 = v174 + v181
	v184 = v175 - v181
	v185 = int32(0)
	v186 = base.B2i32(v184 != v185)
	v188 = v173 + v181
	if v188&int32(3) == v185 {
		v194 = v188
		v195 = v182
		v196 = v184
		v197 = v186
		goto L60
	} else {
		goto L67
	}
L67:
	;
	if v184 != 0 {
		v173 = v188
		v174 = v182
		v175 = v184
		goto L64
	} else {
		goto L68
	}
L68:
	;
	goto L65
L69:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v200 == int32(0) {
		v228 = v194
		v229 = v195
		v230 = v196
		goto L56
	} else {
		goto L70
	}
L70:
	;
	if base.Ui32(v196) < base.Ui32(int32(4)) {
		v228 = v194
		v229 = v195
		v230 = v196
		goto L56
	} else {
		goto L71
	}
L71:
	;
	v206 = v194
	v207 = v195
	v208 = v196
	goto L72
L72:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v214 = int32(-2139062144)
	if (int32(16843008)-v211|v211)&v214 != v214 {
		v235 = v206
		v236 = v207
		v237 = v208
		goto L55
	} else {
		goto L74
	}
L73:
	;
	v228 = v222
	v229 = v220
	v230 = v224
	goto L56
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207))) = v211
	v219 = int32(4)
	v220 = v207 + v219
	v222 = v206 + v219
	v224 = v208 - v219
	if base.Ui32(int32(3)) < base.Ui32(v224) {
		v206 = v222
		v207 = v220
		v208 = v224
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v235 = v228
	v236 = v229
	v237 = v230
	goto L55
L77:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	*(*uint8)(unsafe.Add(mBase, uint32(v241))) = uint8(v244)
	if v244 == int32(0) {
		v255 = v240
		v256 = v241
		goto L54
	} else {
		goto L79
	}
L78:
	;
	v255 = v251
	v256 = v249
	goto L54
L79:
	;
	v248 = int32(1)
	v249 = v241 + v248
	v251 = v240 + v248
	v253 = v242 - v248
	if v253 != 0 {
		v240 = v251
		v241 = v249
		v242 = v253
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+64)))
	if v320 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L82:
	;
	v275 = F_strlen(m, v270)
	mBase = m.M
	v278 = v275 + v270
	goto L85
L83:
	;
	goto L84
L84:
	;
	goto L81
L85:
	;
	v282 = v278 - int32(1)
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282))))
	if v283 == int32(47) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v288 = v282
	goto L91
L87:
	;
	if base.Ui32(v270) < base.Ui32(v282) {
		v278 = v282
		goto L85
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	goto L86
L90:
	;
	goto L89
L91:
	;
	if base.Ui32(v270) < base.Ui32(v288) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v300 = v288
	goto L97
L93:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	if v294 != int32(47) {
		v288 = v288 - int32(1)
		goto L91
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	goto L92
L96:
	;
	goto L95
L97:
	;
	if base.Ui32(v270) < base.Ui32(v300) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if v270 == v300 {
		goto L103
	} else {
		goto L104
	}
L99:
	;
	v304 = v300 - int32(1)
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304))))
	if v305 == int32(47) {
		v300 = v304
		goto L97
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	goto L98
L102:
	;
	goto L101
L103:
	;
	v313 = v270 + base.B2i32(v274 == int32(47))
	goto L105
L104:
	;
	v313 = v300
	goto L105
L105:
	;
	v314 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v313))) = uint8(v314)
	goto L84
L106:
	;
	v323 = int32(46)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+64)) = uint16(v323)
	goto L108
L107:
	;
	goto L108
L108:
	;
	v326 = int32(0)
	v331 = F_fsync_fname_ext(m, v9-int32(-64), int32(1), v326, l2)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L2
	} else {
		goto L109
	}
L109:
	;
	if v331 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v333 = int32(-1)
	goto L112
L111:
	;
	v333 = v326
	goto L112
L112:
	;
	v338 = v333
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
