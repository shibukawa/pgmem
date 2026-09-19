package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_set_fn_opclass_options(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v4 = int32(-1)
	v5 = int32(0)
	v10 = F_makeConst(m, int32(17), v4, v5, v4, l1, base.B2i32(l1 == v5), v5)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v10
		return
	}
}
func Fn13825(m *base.Module, l0 float32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	if base.Ui32(base.I32_reinterpret_f32(l0)&int32(2147483647)) < base.Ui32(int32(2139095041)) {
		if base.F32_eq(base.F32_abs(l0), math.Float32frombits(uint32(0x7f800000))) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				F_errcode(m, int32(130))
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					F_errmsg(m, l3, int32(0))
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						F_errfinish(m, l2, l1, int32(_a_Fn13825_0))
						v40 = m.ExcPending
						if v40 != 0 {
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
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			F_errcode(m, int32(130))
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_errmsg(m, l5, int32(0))
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					F_errfinish(m, l2, l4, int32(_a_Fn13825_0))
					v27 = m.ExcPending
					if v27 != 0 {
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
}
func Fn13829(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = F_SearchSysCache1(m, l5, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
				F_errmsg_internal(m, l4, v11)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l3, l2, l1)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v28+v29)+4))
			F_ReleaseCatCache(m, v13)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v11 + int32(16)
				return v31
			}
		}
	}
}
func Fn13836(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	v8 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = F_SearchSysCache1(m, l6, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v18 + int32(16)
	return v174
L2:
	;
	return int32(0)
L3:
	;
	if v20 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v26 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v26)
	v174 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
	F_errmsg_internal(m, l5, v18)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_Fn13836_0), l4, l3)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v43 = v39 + v40
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+72))
	if v44 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v20)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L2
	} else {
		goto L44
	}
L15:
	;
	v47 = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, _c_Fn13836[0]))
	if v49 == v47 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L32
	}
L18:
	;
	if v88 == int32(0) {
		v157 = v47
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v88 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v56 <= int32(0) {
		v82 = v47
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v88 = v82
	goto L18
L23:
	;
	v59 = int32(0)
	if v59 < v56 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v62 = v56
	goto L26
L25:
	;
	v62 = v59
	goto L26
L26:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v65 = int32(0)
	goto L27
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v63+v65<<(uint(int32(2))%32))))
	v74 = base.B2i32(v73 == v44)
	if v73 == v44 {
		v82 = v74
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v82 = v74
	goto L22
L29:
	;
	v76 = v65 + int32(1)
	if v76 != v62 {
		v65 = v76
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L17
L32:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_Fn13836[0]))
	if v96 == int32(0) {
		v148 = v8
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v157 = base.B2i32(l0 == v148)
	goto L14
L34:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v99 <= int32(0) {
		v148 = v8
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_Fn13836[1]))
	v108 = int32(0)
	v115 = v105
	v119 = v99
	goto L36
L36:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122+v108<<(uint(int32(2))%32))))
	if v115 != v126 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v148 = int32(0)
	goto L33
L38:
	;
	v129 = F_GetSysCacheOid(m, l2, v92, v43+int32(8), v126, int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L2
	} else {
		goto L41
	}
L39:
	;
	v134 = v115
	v135 = v119
	goto L40
L40:
	;
	v137 = v108 + int32(1)
	if v137 < v135 {
		v108 = v137
		v115 = v134
		v119 = v135
		goto L36
	} else {
		goto L43
	}
L41:
	;
	if v129 != 0 {
		v148 = v129
		goto L33
	} else {
		goto L42
	}
L42:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v133 = *(*int32)(unsafe.Add(mBase, _c_Fn13836[1]))
	v134 = v133
	v135 = v131
	goto L40
L43:
	;
	goto L37
L44:
	;
	v174 = v157
	goto L1
}
func Fn13841(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v11
	v15 = F_LockRelease(m, v7, l1, l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func Fn13847(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if int32(0) <= l1 {
		if base.Ui32(l1) < base.Ui32(int32(7)) {
			v44 = l1
			m.G0 = v13 + int32(32)
			return v44
		} else {
			v19 = int32(6)
			v22 = F_errstart(m, int32(19), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					v44 = v19
					m.G0 = v13 + int32(32)
					return v44
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = int32(6)
						if l0 != 0 {
							v35 = int32(_a_Fn13847_0)
						} else {
							v35 = int32(_a_Fn13847_1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
						F_errmsg(m, l7, v13+int32(16))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, l4, l6, l2)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v44 = v19
								m.G0 = v13 + int32(32)
								return v44
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				if l0 != 0 {
					v58 = int32(_a_Fn13847_0)
				} else {
					v58 = int32(_a_Fn13847_1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
				F_errmsg(m, l5, v13)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l4, l3, l2)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
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
func Fn13852(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+16))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v18 == int32(0) {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v22 = base.I32_extend16_s(l1)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+216))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+204))
		v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v24+v26*(v22-int32(1))<<(uint(int32(2))%32)+int32(44)-int32(4))))
		if v38 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(117833860))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_Fn13852_0), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(11)
						F_errdetail_internal(m, int32(_a_Fn13852_1), v11)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, l4, l3, l2)
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
			}
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v43 = F_index_getprocinfo(m, v41, v22, int32(11))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v48 = *(*int64)(unsafe.Add(mBase, uint32(v43)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v48
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v50
				v52 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v52
				v54 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
				*(*int64)(unsafe.Add(mBase, uint32(v17))) = v54
				*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v47
				*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = int32(0)
				m.G0 = v11 + int32(16)
				return v17
			}
		}
	} else {
		m.G0 = v11 + int32(16)
		return v17
	}
}
func Fn13858(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(1)
		v13 = v8 + v12
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v18 = v16 & v12
		if v18 != 0 {
			v19 = v13
		} else {
			v19 = v8 + int32(4)
		}
		if v16 == int32(1) {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v25 == int32(18) {
				v28 = int32(16)
			} else {
				v28 = int32(0)
			}
			if base.Ui32((v25-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v35 = int32(4)
			} else {
				v35 = v28
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v18 != 0 {
				v46 = int32(base.Ui32(v16)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v49 = F_dotrim(m, v19, v46, int32(_a_Fn13858_0), int32(1), l2, l1)
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			return v49
		}
	}
}
func Fn13861(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v14 = v13 - v9
	if base.B2i32(int64(0) < v9) != base.B2i32(v14 < v13) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, l4, int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l3, l2, l1)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
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
		v31 = F_Int64GetDatum(m, v14)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			return v31
		}
	}
}
func Fn13870(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float64, l5 float64) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v17 int32
	_ = v17
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v10 = base.F64_nearest(v9)
	v17 = int32(0)
	if base.B2i32(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v10)&int64(9223372036854775807)))|base.B2i32(base.F64_ge(v10, l5) == v17) == v17)&base.F64_lt(v10, l4) == v17 {
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
				F_errmsg(m, l3, int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn13870_0), l2, l1)
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
		return base.I32_trunc_sat_f64_s(v10)
	}
}
func Fn13876(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = int32(_a_Fn13876_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_Fn13876[0]))
	*(*int32)(unsafe.Add(mBase, _c_Fn13876[0])) = v16 + int32(1)
	v21 = *(*int32)(unsafe.Add(mBase, _c_Fn13876[1]))
	if int32(0) <= v21 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = int32(_a_Fn13876_1)
	v25 = *(*int32)(unsafe.Add(mBase, _c_Fn13876[2]))
	v28 = v21 * int32(100)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn13876[3])))
	*(*int32)(unsafe.Add(mBase, _c_Fn13876[2])) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn13876[4]))) = l0
	v35 = v12 + int32(16)
	F_initStringInfo(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn13876[1])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L21
	}
L4:
	;
	return
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn13876[5])))
	*(*int32)(unsafe.Add(mBase, _c_Fn13876[6])) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l1
	v42 = F_appendStringInfoVA(m, v35, l0, l1)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v42 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v50 = v42
	goto L10
L8:
	;
	goto L9
L9:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn13876[7])))
	if v72 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v54 = v12 + int32(16)
	F_enlargeStringInfo(m, v54, v50)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn13876[5])))
	*(*int32)(unsafe.Add(mBase, _c_Fn13876[6])) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l1
	v61 = F_appendStringInfoVA(m, v54, l0, l1)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	if v61 != 0 {
		v50 = v61
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	F_pfree(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v76 = F_pstrdup(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_Fn13876[7]))) = v76
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	F_pfree(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn13876[2])) = v25
	v84 = int32(_a_Fn13876_0)
	v86 = *(*int32)(unsafe.Add(mBase, _c_Fn13876[0]))
	*(*int32)(unsafe.Add(mBase, _c_Fn13876[0])) = v86 - int32(1)
	m.G0 = v12 + int32(32)
	return
L21:
	;
	F_errmsg_internal(m, int32(_a_Fn13876_2), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_Fn13876_3), l3, l2)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func Fn13883(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = F_gbt_var_same(m, v5, v6, v7, l1, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v4))) = uint8(v9)
		return v4
	}
}
func Fn13885(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = F_gbt_num_same(m, v5, v6, l1, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v4))) = uint8(v8)
		return v4
	}
}
func Fn13889(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 float64
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v15 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v14 + v15
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v14
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+16)))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v22)+12)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = F_gbt_num_distance(m, v8, v8+v15, v24&int32(1), l1, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int32(0)
	} else {
		v32 = F_Float8GetDatum(m, v28)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v32
		}
	}
}
func Fn13896(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = F_pg_detoast_datum(m, l1)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v19 = v10 + int32(8)
			v21 = int32(4)
			v22 = v12 + v21
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = v22
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			v25 = int32(2)
			v26 = int32(base.Ui32(v24) >> (uint(v25) % 32))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			if base.Ui32(v26+v21) < base.Ui32(int32(base.Ui32(v34)>>(uint(v25)%32))) {
				v38 = v22 + (v26+int32(3))&int32(2147483644)
			} else {
				v38 = v22
			}
			*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v38
			v41 = int32(4)
			v42 = v16 + v41
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v42
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
			v45 = int32(2)
			v46 = int32(base.Ui32(v44) >> (uint(v45) % 32))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if base.Ui32(v46+v41) < base.Ui32(int32(base.Ui32(v54)>>(uint(v45)%32))) {
				v58 = v42 + (v46+int32(3))&int32(2147483644)
			} else {
				v58 = v42
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v58
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v63 = F_DirectFunctionCall2Coll(m, l3, v60, v61, v62)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				if l0 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						if l1 != v16 {
							F_pfree(m, v16)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 + int32(16)
								return v63
							}
						} else {
							m.G0 = v10 + int32(16)
							return v63
						}
					}
				} else {
					if l1 != v16 {
						F_pfree(m, v16)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							m.G0 = v10 + int32(16)
							return v63
						}
					} else {
						m.G0 = v10 + int32(16)
						return v63
					}
				}
			}
		}
	}
}
func Fn13898(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, l1, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+92))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func Fn13906(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v86 int32
	_ = v86
	v5 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == v5 {
		v29 = v5
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v29&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if v16 == int32(0) {
		v29 = v5
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v19 != int32(7) {
		v29 = v5
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v22 != int32(17) {
		v29 = v5
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+24)))
	v29 = v25 ^ int32(1)
	goto L2
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = F_get_fn_opclass_options(m, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v38 = l3
	goto L9
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v40 = v39 & l2
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	if v41&l2 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	return int32(0)
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v38 = v37
	goto L9
L12:
	;
	return v9
L13:
	;
	v86 = int32(base.Ui32(v40) >> (uint(l1) % 32))
	goto L15
L14:
	;
	if v40 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v86)
	goto L12
L16:
	;
	v86 = int32(0)
	goto L15
L17:
	;
	v44 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v44)
	v46 = int32(0)
	if v38 <= v46 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v49 = int32(8)
	v53 = v46
	goto L19
L19:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+(v11+v49)))))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+(v10+v49)))))
	if v62 != v64 {
		goto L16
	} else {
		goto L21
	}
L20:
	;
	goto L12
L21:
	;
	v67 = v53 + int32(1)
	if v38 != v67 {
		v53 = v67
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
}
func Fn13913(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 float32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
			v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
			if v18 != v19 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+4)))
						v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v29
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v28
						F_errmsg(m, int32(_a_Fn13913_0), v8)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn13913_1), int32(80), int32(_a_Fn13913_2))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
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
				v41 = int32(8)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v46 = m.T0[v45].(func(*base.Module, int32, int32, int32) float32)(m, base.I32_extend16_s(v18), v11+v41, v16+v41)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v49 = F_Float8GetDatum(m, base.F64_promote_f32(v46))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return v49
					}
				}
			}
		}
	}
}
func Fn13926(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int64) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v67 int64
	_ = v67
	v17 = F_palloc0(m, int32(140))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v17)+4)) = l14
		*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(438)
		*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)) = uint16(v21)
		*(*int32)(unsafe.Add(mBase, uint32(v17)+19)) = v21
		v30 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v17)+18)) = uint8(v30)
		*(*uint16)(unsafe.Add(mBase, uint32(v17)+23)) = uint16(v21)
		*(*int32)(unsafe.Add(mBase, uint32(v17)+108)) = l13
		*(*int32)(unsafe.Add(mBase, uint32(v17)+104)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = l12
		*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = l11
		*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = l10
		*(*int32)(unsafe.Add(mBase, uint32(v17)+88)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v17)+84)) = l9
		*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = l8
		*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = l7
		*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = l6
		*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v21
		*(*uint8)(unsafe.Add(mBase, uint32(v17)+29)) = uint8(v30)
		*(*int32)(unsafe.Add(mBase, uint32(v17)+25)) = v30
		*(*int32)(unsafe.Add(mBase, uint32(v17)+136)) = v21
		v67 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v17)+128)) = v67
		*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = v67
		*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = v67
		return v17
	}
}
func Fn13933(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = F_text_to_cstring(m, v15)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v17
				v24 = F_get_worker(m, v10, v7+int32(12), int32(0), int32(1), l1)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					if v24 == int32(0) {
						v28 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
						v31 = int32(0)
					} else {
						v31 = v24
					}
					m.G0 = v7 + int32(16)
					return v31
				}
			}
		}
	}
}
func Fn13935(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v9
	v14 = F_pushJsonbValue(m, v7, l2, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v19 = F_pushJsonbValue(m, v7, l1, int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v19
			v22 = F_JsonbValueToJsonb(m, v19)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return v22
			}
		}
	}
}
func Fn13940(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	if l1 == int32(0) {
		v8 = F_palloc(m, int32(32))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(4)
			*(*int64)(unsafe.Add(mBase, uint32(v8))) = l2
			v16 = v8 + int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v16
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
			return v8
		}
	} else {
		F_new_head_cell(m, l1)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v22))) = l0
			return l1
		}
	}
}
func Fn13942(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 float64
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 float64
	_ = v56
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 float64
	_ = v61
	var v64 int32
	_ = v64
	var v65 float64
	_ = v65
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v103 float64
	_ = v103
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v107 float64
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v202 float64
	_ = v202
	var v204 float64
	_ = v204
	var v206 float64
	_ = v206
	var v219 float64
	_ = v219
	var v229 float64
	_ = v229
	var v240 float64
	_ = v240
	var v252 float64
	_ = v252
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int64
	_ = v267
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int64
	_ = v280
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int64
	_ = v298
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int64
	_ = v324
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	v11 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(96)
	m.G0 = v18
	if l0 < int32(_a_Fn13942_0) {
		v336 = v11
		m.G0 = v18 + int32(96)
		return v336
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
		if v23 < int64(10000) {
			v336 = v11
			m.G0 = v18 + int32(96)
			return v336
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
			if v26 != int32(1) {
				v336 = v11
				m.G0 = v18 + int32(96)
				return v336
			} else {
				v30 = v22 + int32(16)
				v31 = int32(0)
				v38 = float64(0)
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
				if v40 != 0 {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
					if v40 != int32(1) {
						v50 = v31
						v51 = v31
						v56 = v38
						for {
							v58 = float64(1)
							v59 = v50 + v41
							v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
							v61 = F_scalbn(m, v58, v60)
							mBase = m.M
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
							v65 = F_scalbn(m, v58, v64)
							mBase = m.M
							v70 = base.F64_add(base.F64_add(v56, base.F64_div(v58, v65)), base.F64_div(v58, v61))
							v71 = int32(2)
							v72 = v50 + v71
							v74 = v51 + v71
							if v74 != v40&int32(-2) {
								v50 = v72
								v51 = v74
								v56 = v70
								continue
							} else {
								break
							}
							break
						}
						if v40&int32(1) == int32(0) {
							v103 = v70
						} else {
							v80 = v72
							v86 = v70
							v88 = float64(1)
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v41))))
							v92 = F_scalbn(m, v88, v91)
							mBase = m.M
							v103 = base.F64_add(v86, base.F64_div(v88, v92))
						}
					} else {
						v80 = v31
						v86 = v38
						v88 = float64(1)
						v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v41))))
						v92 = F_scalbn(m, v88, v91)
						mBase = m.M
						v103 = base.F64_add(v86, base.F64_div(v88, v92))
					}
					v105 = *(*float64)(unsafe.Add(mBase, uint32(v30)+8))
					v106 = base.F64_div(v105, v103)
					v107 = base.F64_convert_i32_u(v40)
					if base.F64_le(v106, base.F64_mul(v107, float64(2.5))) == int32(0) {
						v219 = v106
						if base.F64_gt(v219, float64(1.4316557653333333e+08)) == int32(0) {
							v240 = v219
						} else {
							v229 = F_log(m, base.F64_add(base.F64_mul(v219, float64(-2.3283064365386963e-10)), float64(1)))
							mBase = m.M
							v240 = base.F64_mul(v229, float64(-4.294967296e+09))
						}
						v252 = v240
					} else {
						v114 = v40 & int32(3)
						v115 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
						v116 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v40) {
							v125 = int32(0)
							v126 = v116
							v127 = v116
							for {
								v134 = v126 + v115
								v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
								v136 = int32(0)
								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+1)))
								v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+2)))
								v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+3)))
								v150 = v127 + base.B2i32(v135 == v136) + base.B2i32(v139 == v136) + base.B2i32(v143 == v136) + base.B2i32(v147 == v136)
								v151 = int32(4)
								v152 = v126 + v151
								v154 = v125 + v151
								if v154 != v40&int32(-4) {
									v125 = v154
									v126 = v152
									v127 = v150
									continue
								} else {
									break
								}
								break
							}
							if v114 == int32(0) {
								v191 = v150
							} else {
								v160 = v152
								v161 = v150
								v170 = v160
								v171 = v161
								v174 = v116
								for {
									v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v115))))
									v182 = v171 + base.B2i32(v179 == int32(0))
									v183 = int32(1)
									v186 = v174 + v183
									if v186 != v114 {
										v170 = v170 + v183
										v171 = v182
										v174 = v186
										continue
									} else {
										break
									}
									break
								}
								v191 = v182
							}
						} else {
							v160 = v116
							v161 = v116
							v170 = v160
							v171 = v161
							v174 = v116
							for {
								v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v115))))
								v182 = v171 + base.B2i32(v179 == int32(0))
								v183 = int32(1)
								v186 = v174 + v183
								if v186 != v114 {
									v170 = v170 + v183
									v171 = v182
									v174 = v186
									continue
								} else {
									break
								}
								break
							}
							v191 = v182
						}
						if v191 == int32(0) {
							v240 = v106
							v252 = v240
						} else {
							v202 = F_log(m, base.F64_div(v107, base.F64_convert_i32_s(v191)))
							mBase = m.M
							v252 = base.F64_mul(v202, v107)
						}
					}
				} else {
					v204 = *(*float64)(unsafe.Add(mBase, uint32(v30)+8))
					v206 = base.F64_div(v204, float64(0))
					if base.F64_le(v206, base.F64_mul(base.F64_convert_i32_u(v40), float64(2.5))) != 0 {
						v240 = v206
					} else {
						v219 = v206
						if base.F64_gt(v219, float64(1.4316557653333333e+08)) == int32(0) {
							v240 = v219
						} else {
							v229 = F_log(m, base.F64_add(base.F64_mul(v219, float64(-2.3283064365386963e-10)), float64(1)))
							mBase = m.M
							v240 = base.F64_mul(v229, float64(-4.294967296e+09))
						}
					}
					v252 = v240
				}
				if base.F64_gt(v252, float64(100000)) != 0 {
					v256 = int32(*(*uint8)(unsafe.Add(mBase, _c_Fn13942[0])))
					if v256 != int32(1) {
						v276 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)) = uint8(v276)
						v336 = v11
						m.G0 = v18 + int32(96)
						return v336
					} else {
						v261 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v264 = m.ExcPending
						if v264 != 0 {
							return int32(0)
						} else {
							if v261 == int32(0) {
								v276 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)) = uint8(v276)
								v336 = v11
								m.G0 = v18 + int32(96)
								return v336
							} else {
								v267 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
								*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = l0
								*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v267
								*(*float64)(unsafe.Add(mBase, uint32(v18))) = v252
								F_errmsg_internal(m, l9, v18)
								mBase = m.M
								v272 = m.ExcPending
								if v272 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, l4, l8, l2)
									mBase = m.M
									v274 = m.ExcPending
									if v274 != 0 {
										return int32(0)
									} else {
										v276 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)) = uint8(v276)
										v336 = v11
										m.G0 = v18 + int32(96)
										return v336
									}
								}
							}
						}
					}
				} else {
					v279 = int32(*(*uint8)(unsafe.Add(mBase, _c_Fn13942[0])))
					v280 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
					if base.F64_gt(base.F64_add(base.F64_div(base.F64_convert_i64_s(v280), float64(2000)), float64(0.5)), v252) != 0 {
						v287 = int32(1)
						if v279&v287 == int32(0) {
							v336 = v287
							m.G0 = v18 + int32(96)
							return v336
						} else {
							v294 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v295 = m.ExcPending
							if v295 != 0 {
								return int32(0)
							} else {
								if v294 == int32(0) {
									v336 = v287
									m.G0 = v18 + int32(96)
									return v336
								} else {
									v298 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
									*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = l0
									*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v298
									*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v252
									*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = base.F64_add(base.F64_div(base.F64_convert_i64_s(v298), float64(2000)), float64(0.5))
									F_errmsg_internal(m, l7, v18+int32(32))
									mBase = m.M
									v311 = m.ExcPending
									if v311 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, l4, l6, l2)
										mBase = m.M
										v313 = m.ExcPending
										if v313 != 0 {
											return int32(0)
										} else {
											v336 = v287
											m.G0 = v18 + int32(96)
											return v336
										}
									}
								}
							}
						}
					} else {
						if v279&int32(1) == int32(0) {
							v336 = v11
							m.G0 = v18 + int32(96)
							return v336
						} else {
							v320 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v321 = m.ExcPending
							if v321 != 0 {
								return int32(0)
							} else {
								if v320 == int32(0) {
									v336 = v11
									m.G0 = v18 + int32(96)
									return v336
								} else {
									v324 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
									*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = l0
									*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v324
									*(*float64)(unsafe.Add(mBase, uint32(v18)+64)) = v252
									F_errmsg_internal(m, l5, v18-int32(-64))
									mBase = m.M
									v331 = m.ExcPending
									if v331 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, l4, l3, l2)
										mBase = m.M
										v333 = m.ExcPending
										if v333 != 0 {
											return int32(0)
										} else {
											v336 = v11
											m.G0 = v18 + int32(96)
											return v336
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
}
func Fn13957(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(34209794)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v10
	v17 = *(*int32)(unsafe.Add(mBase, _c_Fn13957[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v17
	v20 = F_LockAcquire(m, v8, l2, l1, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return int32(0)
	}
}
func Fn13960(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(34209793)
	*(*uint32)(unsafe.Add(mBase, uint32(v7)+8)) = uint32(v10)
	v15 = int64(base.Ui64(v10) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v7)+4)) = uint32(v15)
	v18 = *(*int32)(unsafe.Add(mBase, _c_Fn13960[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v18
	v21 = F_LockRelease(m, v7, l1, int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v21
	}
}
func Fn13966(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	v4 = int32(_a_Fn13966_0)
	v5 = int32(_a_Fn13966_1)
	v6 = *(*int64)(unsafe.Add(mBase, _c_Fn13966[0]))
	v8 = *(*int64)(unsafe.Add(mBase, _c_Fn13966[1]))
	v9 = v6 ^ v8
	*(*int64)(unsafe.Add(mBase, _c_Fn13966[1])) = base.I64_rotl(v9, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_Fn13966[0])) = v9<<(uint(int64(16))%64) ^ base.I64_rotl(v6, int64(24)) ^ v9
	return base.I32_wrap_i64(int64(base.Ui64(base.I64_rotl(v6*int64(5), int64(7))*int64(9)) >> (uint(l0) % 64)))
}
func Fn13971(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = l5
	v20 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_Fn13971[0]))))
	return int32(base.Ui32(v47&l2) >> (uint(l1) % 32))
L4:
	;
	v25 = base.I32_div_s(v19+v20, int32(2))
	v27 = v25 << (uint(int32(3)) % 32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27+l4)))
	if base.Ui32(v29) < base.Ui32(l0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	if v41 <= v40 {
		v19 = v40
		v20 = v41
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v40 = v19
	v41 = v25 + int32(1)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27+l3)))
	if base.Ui32(v34) <= base.Ui32(l0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(1)
L11:
	;
	goto L12
L12:
	;
	v40 = v25 - int32(1)
	v41 = v20
	goto L6
L13:
	;
	goto L5
}
func Fn13977(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v41 int32
	_ = v41
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if l3 < v12 {
		m.G0 = v9 + int32(16)
		return int32(0)
	} else {
		v14 = F_strlen(m, l1)
		mBase = m.M
		if base.Ui32(int32(63)) < base.Ui32(v14) {
			m.G0 = v9 + int32(16)
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v21 = F_hash_search(m, v17, l1, int32(1), v9+int32(15))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v25
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v27 != 0 {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
					v30 = v29 - v27
					v33 = F_palloc(m, v30+int32(1))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						if v30 != 0 {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							base.MemoryCopy(m, v33, v35, v30)
						} else {
						}
						v38 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v33+v30))) = uint8(v38)
						v41 = v33
						*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v41
						m.G0 = v9 + int32(16)
						return int32(0)
					}
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v41 = v40
					*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v41
					m.G0 = v9 + int32(16)
					return int32(0)
				}
			}
		}
	}
}
func Fn13986(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v5 <= v7 {
		v36 = v3
		return v36
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v5-int32(1)))))
		switch v13 - int32(105) {
		case 0, 5:
			v17 = F_find_among_b(m, l0, l1, int32(3))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v17 == int32(0) {
					v36 = v3
					return v36
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
					v25 = F_slice_del(m, l0)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if v25 < int32(0) {
							v36 = v25
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
							v31 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v30 - v31
							v36 = v31
						}
						return v36
					}
				}
			}
		default:
			v36 = v3
			return v36
		}
	}
}
func Fn13988(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = l1
	return v4
}
func Fn13993(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(255)
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0 & v8
	if base.Ui32((l0-int32(33))&v8) < base.Ui32(int32(94)) {
		v20 = int32(_a_Fn13993_0)
	} else {
		v20 = int32(_a_Fn13993_1)
	}
	v21 = F_pg_snprintf(m, l1, int32(5), v20, v6)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return
	} else {
		m.G0 = v6 + int32(16)
		return
	}
}
func Fn13999(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 float64
	_ = v16
	var v19 float64
	_ = v19
	var v22 float64
	_ = v22
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v37 int64
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v59 int64
	_ = v59
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	v9 = F_MemoryContextAllocZero(m, l0, int32(32))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l0
		v16 = float64(4.294967296e+09)
		v19 = base.F64_div(base.F64_convert_i32_u(l1), float64(0.9))
		if base.F64_ge(v19, v16) != 0 {
			v22 = v16
		} else {
			v22 = v19
		}
		v23 = base.I64_trunc_sat_f64_u(v22)
		if base.Ui64(v23) <= base.Ui64(int64(2)) {
			v26 = int64(2)
		} else {
			v26 = v23
		}
		v27 = int64(1)
		if v26&(v26-v27) == int64(0) {
			v37 = v26
		} else {
			v37 = v27 << (uint(int64(64)-base.I64_clz(v26)) % 64)
		}
		if base.Ui64(v37<<(uint(int64(3))%64)) < base.Ui64(int64(2147483647)) {
			v46 = F_MemoryContextAllocExtended(m, l0, base.I32_wrap_i64(v37)<<(uint(int32(3))%32), int32(5))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v46
				v49 = int64(1)
				if v37&(v37-v49) == int64(0) {
					v59 = v37
				} else {
					v59 = v49 << (uint(int64(64)-base.I64_clz(v37)) % 64)
				}
				if base.Ui64(int64(2147483647)) <= base.Ui64(v59<<(uint(int64(3))%64)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_Fn13999_0), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn13999_1), int32(327), l3)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v9))) = v59
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = base.I32_wrap_i64(v59) - int32(1)
					if v59 == int64(4294967296) {
						v76 = int32(-85899346)
					} else {
						v76 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v59), float64(0.9)))
					}
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v76
					return v9
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_Fn13999_0), int32(0))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn13999_1), int32(327), l3)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
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
func Fn14002(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v15 = F_ArrayGetIntegerTypmods(m, v9, v6+int32(12))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			if v17 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_Fn14002_0), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn14002_1), int32(118), int32(_a_Fn14002_2))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
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
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v37 = F_anytimestamp_typmod_check(m, l1, v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return v37
				}
			}
		}
	}
}
func Fn14017(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(1)
		v13 = v8 + v12
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v18 = v16 & v12
		if v18 != 0 {
			v19 = v13
		} else {
			v19 = v8 + int32(4)
		}
		if v16 == int32(1) {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v25 == int32(18) {
				v28 = int32(16)
			} else {
				v28 = int32(0)
			}
			if base.Ui32((v25-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v35 = int32(4)
			} else {
				v35 = v28
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v18 != 0 {
				v46 = int32(base.Ui32(v16)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v47 = F_uuid_generate_internal(m, l1, v6, v19, v46)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			return v47
		}
	}
}
func Fn14020(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = F_WinGetFuncArgInPartition(m, v9, l1, int32(1), v7+int32(15), v7+int32(14))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v19 == int32(1) {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
			v25 = int32(0)
		} else {
			v25 = v15
		}
		m.G0 = v7 + int32(16)
		return v25
	}
}
func Fn14026(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
		F_errmsg_internal(m, l4, v9)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			F_errfinish(m, l3, l2, l1)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
