package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecSeqScanInstrumentEstimate(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+132)))
	if v4&int32(16) == int32(0) {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		if v9 == int32(0) {
			return
		} else {
			v14 = F_mul_size(m, v9, int32(56))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v16 = F_add_size(m, int32(8), v14)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					v23 = F_add_size(m, v18, (v16+int32(31))&int32(-32))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v23
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						v28 = F_add_size(m, v26, int32(1))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v28
							return
						}
					}
				}
			}
		}
	}
}
func F_ExecSeqScanWithQual(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 float64
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	F_MemoryContextReset(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L4
L3:
	;
	m.G0 = v13 + int32(16)
	return v126
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQual[0]))
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L34
	}
L6:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v39 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L8
L10:
	;
	goto L5
L11:
	;
	v42 = F_ScanRelIsReadOnly(m, l0)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v72 = v39
	goto L13
L13:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+40)) = v74
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+188))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
	v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, v72, v38, v36)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L23
	}
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQual[1]))
	if v45 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQual[2])))
	if v47&int32(1) == int32(0) {
		goto L10
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v54 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v37)+132))
	if v42 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v64 = int32(1473)
	goto L21
L20:
	;
	v64 = int32(449)
	goto L21
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v52)+188))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	v68 = m.T0[v67].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v52, v53, v54, v54, v54, v57<<(uint(int32(7))%32)&int32(2048)|v64)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v68
	v72 = v68
	goto L13
L23:
	;
	v82 = int32(0)
	if base.B2i32(v80 == v82)|base.B2i32(v36 == v82) != 0 {
		v126 = int32(0)
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
	if v87&int32(2) != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v126 = v36
	goto L3
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v36
	v91 = int32(_a_F_ExecSeqScanWithQual_0)
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQual[3]))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQual[3])) = v94
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v99 = m.T0[v98].(func(*base.Module, int32, int32, int32) int64)(m, v15, v16, v13+int32(15))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQual[3])) = v92
	if v99 != int64(0) {
		v126 = v36
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v105 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v106 = *(*float64)(unsafe.Add(mBase, uint32(v105)+424))
	*(*float64)(unsafe.Add(mBase, uint32(v105)+424)) = base.F64_add(v106, float64(1))
	goto L32
L31:
	;
	goto L32
L32:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	F_MemoryContextReset(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L4
L34:
	;
	F_errmsg_internal(m, int32(_a_F_ExecSeqScanWithQual_1), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_ExecSeqScanWithQual_2), int32(931), int32(_a_F_ExecSeqScanWithQual_3))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SeqNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v9 == int32(0) {
		v12 = F_ScanRelIsReadOnly(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_SeqNext[0]))
			if v17 != 0 {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SeqNext[1])))
				if v19&int32(1) == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_SeqNext_0), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_SeqNext_1), int32(931), int32(_a_F_SeqNext_2))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
					v26 = int32(0)
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+132))
					if v12 != 0 {
						v36 = int32(1473)
					} else {
						v36 = int32(449)
					}
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v24)+188))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
					v40 = m.T0[v39].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v24, v25, v26, v26, v26, v29<<(uint(int32(7))%32)&int32(2048)|v36)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v40
						v44 = v40
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v46
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+188))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
						v52 = m.T0[v51].(func(*base.Module, int32, int32, int32) int32)(m, v44, v8, v6)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							if v52 != 0 {
								v54 = v6
							} else {
								v54 = int32(0)
							}
							return v54
						}
					}
				}
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
				v26 = int32(0)
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+132))
				if v12 != 0 {
					v36 = int32(1473)
				} else {
					v36 = int32(449)
				}
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v24)+188))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
				v40 = m.T0[v39].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v24, v25, v26, v26, v26, v29<<(uint(int32(7))%32)&int32(2048)|v36)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v40
					v44 = v40
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v46
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+188))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
					v52 = m.T0[v51].(func(*base.Module, int32, int32, int32) int32)(m, v44, v8, v6)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						if v52 != 0 {
							v54 = v6
						} else {
							v54 = int32(0)
						}
						return v54
					}
				}
			}
		}
	} else {
		v44 = v9
		v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v46
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+188))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
		v52 = m.T0[v51].(func(*base.Module, int32, int32, int32) int32)(m, v44, v8, v6)
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int32(0)
		} else {
			if v52 != 0 {
				v54 = v6
			} else {
				v54 = int32(0)
			}
			return v54
		}
	}
}
func F_seq_identify(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	if base.Ui32(l0) < base.Ui32(int32(16)) {
		v6 = int32(_a_F_seq_identify_0)
	} else {
		v6 = int32(0)
	}
	return v6
}
