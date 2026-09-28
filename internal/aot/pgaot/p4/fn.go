package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_set_fn_opclass_options(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v4 = int32(-1)
	v5 = int32(0)
	v11 = F_makeConst(m, int32(17), v4, v5, v4, base.I64_extend_i32_u(l1), base.B2i32(l1 == v5), v5)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v11
		return
	}
}
func Fn14208(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if int32(0) < l0 {
		if base.Ui32(l7) <= base.Ui32(l0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = l4
					F_errmsg(m, l3, v12)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						F_errfinish(m, l2, l1, int32(_a_Fn14208_0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
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
			m.G0 = v12 + int32(16)
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_errcode(m, int32(130))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_errmsg(m, l6, int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_errfinish(m, l2, l5, int32(_a_Fn14208_0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
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
func Fn14217(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
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
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int64
	_ = v79
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	if l1 == int32(0) {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
		if v18 != 0 {
			v79 = int64(0)
			m.G0 = v14 + int32(96)
			return v79
		} else {
			v19 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v14)+28)) = v19
			*(*int64)(unsafe.Add(mBase, uint32(v14)+33)) = v19
			v23 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v14)+88)) = uint8(v23)
			*(*uint8)(unsafe.Add(mBase, uint32(v14)+72)) = uint8(v23)
			*(*uint8)(unsafe.Add(mBase, uint32(v14)+56)) = uint8(v23)
			v29 = int32(3)
			*(*uint16)(unsafe.Add(mBase, uint32(v14)+42)) = uint16(v29)
			*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = base.I64_extend_i32_s(l3)
			*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = base.I64_extend_i32_u(l2)
			*(*int64)(unsafe.Add(mBase, uint32(v14)+48)) = base.I64_extend_i32_u(l1)
			*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = l0
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v41 = m.T0[v40].(func(*base.Module, int32) int64)(m, v14+int32(24))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+40)))
				if l1 == int32(0) {
					if v45&int32(1) != 0 {
						v79 = v41
						m.G0 = v14 + int32(96)
						return v79
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int64(0)
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v14))) = v54
							F_errmsg_internal(m, l8, v14)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_Fn14217_0), l7, l4)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					if v45&int32(1) == int32(0) {
						v79 = v41
						m.G0 = v14 + int32(96)
						return v79
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int64(0)
						} else {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v69
							F_errmsg_internal(m, l6, v14+int32(16))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_Fn14217_0), l5, l4)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int64(0)
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
	} else {
		v19 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v14)+28)) = v19
		*(*int64)(unsafe.Add(mBase, uint32(v14)+33)) = v19
		v23 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+88)) = uint8(v23)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+72)) = uint8(v23)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+56)) = uint8(v23)
		v29 = int32(3)
		*(*uint16)(unsafe.Add(mBase, uint32(v14)+42)) = uint16(v29)
		*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = base.I64_extend_i32_s(l3)
		*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = base.I64_extend_i32_u(l2)
		*(*int64)(unsafe.Add(mBase, uint32(v14)+48)) = base.I64_extend_i32_u(l1)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = l0
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v41 = m.T0[v40].(func(*base.Module, int32) int64)(m, v14+int32(24))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int64(0)
		} else {
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+40)))
			if l1 == int32(0) {
				if v45&int32(1) != 0 {
					v79 = v41
					m.G0 = v14 + int32(96)
					return v79
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int64(0)
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v54
						F_errmsg_internal(m, l8, v14)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_Fn14217_0), l7, l4)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				if v45&int32(1) == int32(0) {
					v79 = v41
					m.G0 = v14 + int32(96)
					return v79
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int64(0)
					} else {
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v69
						F_errmsg_internal(m, l6, v14+int32(16))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_Fn14217_0), l5, l4)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int64(0)
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
func Fn14219(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v93 int64
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	v8 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v21 = F_SearchSysCache1(m, l6, base.I64_extend_i32_u(l0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v18 + int32(16)
	return v177
L2:
	;
	return int32(0)
L3:
	;
	if v21 == int32(0) {
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
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v27 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v27)
	v177 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
	F_errmsg_internal(m, l5, v18)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_Fn14219_0), l4, l3)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
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
	v44 = v40 + v41
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	if v45 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v21)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L44
	}
L15:
	;
	v48 = int32(0)
	v50 = *(*int32)(unsafe.Add(mBase, _c_Fn14219[0]))
	if v50 == v48 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v93 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v44)+4)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L32
	}
L18:
	;
	if v89 == int32(0) {
		v160 = v48
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v89 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v57 <= int32(0) {
		v83 = v48
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v89 = v83
	goto L18
L23:
	;
	v60 = int32(0)
	if v60 < v57 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v63 = v57
	goto L26
L25:
	;
	v63 = v60
	goto L26
L26:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v66 = int32(0)
	goto L27
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v64+v66<<(uint(int32(2))%32))))
	v75 = base.B2i32(v74 == v45)
	if v74 == v45 {
		v83 = v75
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v83 = v75
	goto L22
L29:
	;
	v77 = v66 + int32(1)
	if v77 != v63 {
		v66 = v77
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
	v97 = *(*int32)(unsafe.Add(mBase, _c_Fn14219[0]))
	if v97 == int32(0) {
		v151 = v8
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v160 = base.B2i32(l0 == v151)
	goto L14
L34:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if v100 <= int32(0) {
		v151 = v8
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_Fn14219[1]))
	v110 = int32(0)
	v117 = v107
	v121 = v100
	goto L36
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124+v110<<(uint(int32(2))%32))))
	if v117 != v128 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v151 = int32(0)
	goto L33
L38:
	;
	v132 = F_GetSysCacheOid(m, l2, v93, base.I64_extend_i32_u(v44+int32(8)), base.I64_extend_i32_u(v128), int64(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L2
	} else {
		goto L41
	}
L39:
	;
	v137 = v117
	v138 = v121
	goto L40
L40:
	;
	v140 = v110 + int32(1)
	if v140 < v138 {
		v110 = v140
		v117 = v137
		v121 = v138
		goto L36
	} else {
		goto L43
	}
L41:
	;
	if v132 != 0 {
		v151 = v132
		goto L33
	} else {
		goto L42
	}
L42:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v136 = *(*int32)(unsafe.Add(mBase, _c_Fn14219[1]))
	v137 = v136
	v138 = v134
	goto L40
L43:
	;
	goto L37
L44:
	;
	v177 = v160
	goto L1
}
func Fn14220(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	goto L2
L1:
	;
	return
L2:
	;
	v13 = int32(0)
	v14 = m.Env.Pgmem_sem(m, l4, l0, v13)
	mBase = m.M
	if v13 <= v14 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	if int32(0) <= v22 {
		goto L1
	} else {
		goto L8
	}
L5:
	;
	v22 = v14
	goto L4
L6:
	;
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn14220[0])) = int32(0) - v14
	v22 = int32(-1)
	goto L4
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_Fn14220[0]))
	if v26 == int32(27) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	goto L3
L10:
	;
	return
L11:
	;
	F_errmsg_internal(m, l3, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_Fn14220_0), l2, l1)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func Fn14231(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_check_encoding_conversion_args(m, v13, v14, v15, l1, int32(6))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v23 = F_LocalToUtf(m, v11, v15, v10, l5, l4, l3, l2, l1, base.B2i32(v12 != int64(0)))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_s(v23)
		}
	}
}
func Fn14237(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	v4 = int32(_a_Fn14237_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_Fn14237[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	*(*int32)(unsafe.Add(mBase, _c_Fn14237[0])) = v8
	F_varstr_sortsupport(m, v7, l1, int32(950))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_Fn14237[0])) = v5
		return int64(0)
	}
}
func Fn14244(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v36 int64
	_ = v36
	var v43 int64
	_ = v43
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = int64(63)
	v15 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	v22 = int64(32)
	v23 = int64(base.Ui64(v15) >> (uint(v22) % 64))
	v25 = int64(base.Ui64(v12) >> (uint(v22) % 64))
	v28 = int64(4294967295)
	v29 = v15 & v28
	v31 = v12 & v28
	v32 = v29 * v31
	v36 = int64(base.Ui64(v32)>>(uint(v22)%64)) + v29*v25
	v43 = v31*v23 + v36&v28
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v12*(v15>>(uint(v13)%64)) + v12>>(uint(v13)%64)*v15 + v23*v25 + int64(base.Ui64(v36)>>(uint(v22)%64)) + int64(base.Ui64(v43)>>(uint(v22)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v32&v28 | v43<<(uint(v22)%64)
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	if v54 != v55>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, l4, int32(0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, l3, l2, l1)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.G0 = v10 + int32(16)
		return v55
	}
}
func Fn14253(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v16 - int32(1) {
	case 0:
		if l1 == int32(0) {
			v69 = int32(0)
			m.G0 = v14 + int32(48)
			return v69
		} else {
			v48 = int32(23)
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v50 = F_errsave_start(m, v49)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				if v50 == int32(0) {
					v69 = v48
					m.G0 = v14 + int32(48)
					return v69
				} else {
					F_errcode(m, int32(33685634))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v57
						F_errmsg(m, l7, v14+int32(32))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v65 = F_errdetail(m, int32(_a_Fn14253_0), int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, v49, l4, l6, l2)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v69 = v48
									m.G0 = v14 + int32(48)
									return v69
								}
							}
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v78 = m.ExcPending
		if v78 != 0 {
			return int32(0)
		} else {
			v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v79
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = l5
			F_errmsg_internal(m, int32(_a_Fn14253_1), v14)
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, l4, l3, l2)
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 3:
		if l1 == int32(0) {
			v69 = int32(0)
			m.G0 = v14 + int32(48)
			return v69
		} else {
			v22 = int32(23)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v24 = F_errsave_start(m, v23)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v24 == int32(0) {
					v69 = v22
					m.G0 = v14 + int32(48)
					return v69
				} else {
					F_errcode(m, int32(33685634))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v33
						F_errmsg(m, l7, v14+int32(16))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v41 = F_errdetail(m, int32(_a_Fn14253_2), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, v23, l4, l8, l2)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									v69 = v22
									m.G0 = v14 + int32(48)
									return v69
								}
							}
						}
					}
				}
			}
		}
	}
}
func Fn14259(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = F_get_fn_expr_argtype(m, v12, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		if v14 != 0 {
			v18 = F_enum_endpoint(m, v14, l4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				if v18 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int64(0)
						} else {
							v48 = F_format_type_be(m, v14)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v48
								F_errmsg(m, int32(_a_Fn14259_0), v10)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_Fn14259_1), l2, l1)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int64(0)
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
					m.G0 = v10 + int32(16)
					return base.I64_extend_i32_u(v18)
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_Fn14259_2), int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_Fn14259_1), l3, l1)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int64(0)
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
func Fn14264(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = l9
	v18 = int32(1)
	v22 = F_LookupFuncName(m, l0, v18, v15+int32(44), v18)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		if v22 != 0 {
			v26 = F_get_func_rettype(m, v22)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 != l8 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(117833860))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							v83 = F_NameListToString(m, l0)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v83
								F_errmsg(m, l3, v15+int32(32))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_Fn14264_0), l2, l1)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
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
					v29 = F_func_volatile(m, v22)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						if v29 != int32(118) {
							m.G0 = v15 + int32(48)
							return v22
						} else {
							v35 = F_errstart(m, int32(19), int32(0))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								if v35 == int32(0) {
									m.G0 = v15 + int32(48)
									return v22
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return int32(0)
									} else {
										v42 = F_NameListToString(m, l0)
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v42
											F_errmsg(m, l7, v15+int32(16))
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_Fn14264_0), l6, l1)
												mBase = m.M
												v51 = m.ExcPending
												if v51 != 0 {
													return int32(0)
												} else {
													m.G0 = v15 + int32(48)
													return v22
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(52461700))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					v67 = F_func_signature_string(m, l0, int32(1), int32(0), v15+int32(44))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v67
						F_errmsg(m, int32(_a_Fn14264_1), v15)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn14264_0), l5, l1)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
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
func Fn14268(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = F_gbt_var_picksplit(m, v4, base.I32_wrap_i64(v5), v7, l1, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		return v5 & int64(4294967295)
	}
}
func Fn14273(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
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
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = F_gbt_num_same(m, v6, v7, l1, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v4)))) = uint8(v9)
		return v4 & int64(4294967295)
	}
}
func Fn14275(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
		v22 = v11 + int32(8)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v25 = int32(4)
		v26 = v23 + v25
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = v26
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		v29 = int32(2)
		v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		if base.Ui32(v30+v25) < base.Ui32(int32(base.Ui32(v38)>>(uint(v29)%32))) {
			v42 = v26 + (v30+int32(3))&int32(2147483644)
		} else {
			v42 = v26
		}
		*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v42
		v44 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v44)
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
		v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+16)))
		v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47+v48)+12)))
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v54 = F_gbt_var_consistent(m, v22, v15, v19, v46, v50&int32(1), l1, v53)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int64(0)
		} else {
			m.G0 = v11 + int32(16)
			return base.I64_extend_i32_u(v54)
		}
	}
}
func Fn14279(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 float64
	_ = v30
	var v33 int32
	_ = v33
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v12)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v15 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v14 + v15
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v14
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+16)))
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+v24)+12)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = F_gbt_num_distance(m, v9+v15, v9+int32(12), v26&int32(1), l1, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int64(0)
	} else {
		m.G0 = v9 + int32(16)
		return base.I64_reinterpret_f64(v30)
	}
}
func Fn14280(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+14)) = uint16(v14)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v16 + l2
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v16
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+16)))
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+v28)+12)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = F_gbt_num_consistent(m, v10+int32(4), v12, v10+int32(14), v30&int32(1), l1, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return int64(0)
	} else {
		m.G0 = v10 + int32(16)
		return base.I64_extend_i32_u(v34)
	}
}
func Fn14291(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_SearchSysCache1(m, l1, base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14)+96))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v16
			}
		}
	}
}
func Fn14307(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int64(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v24 = F_pg_detoast_datum_packed(m, v23)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int64(0)
		} else {
			v26 = F_get_role_oid_or_public(m, v17)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int64(0)
			} else {
				v29 = F_text_to_cstring(m, v19)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					v32 = F_DirectFunctionCall1Coll(m, l7, int32(0), base.I64_extend_i32_u(v29))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						v34 = base.I32_wrap_i64(v32)
						if v34 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int64(0)
							} else {
								F_errcode(m, l6)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15))) = v29
									F_errmsg(m, l5, v15)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_Fn14307_0), l4, l3)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v49 = F_convert_any_priv_string(m, v24, l1)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int64(0)
							} else {
								v51 = F_object_aclcheck(m, l2, v34, v26, v49)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int64(0)
								} else {
									m.G0 = v15 + int32(16)
									return base.I64_extend_i32_u(base.B2i32(v51 == int32(0)))
								}
							}
						}
					}
				}
			}
		}
	}
}
func Fn14309(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v13 = F_table_open(m, l3, int32(1))
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = v8 + int32(-56)
		F_ScanKeyInit(m, v18, l2, int32(3), int32(184), base.I64_extend_i32_u(l0))
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = int32(1)
			v27 = F_systable_beginscan(m, v13, l1, v24, int32(0), v24, v18)
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = F_systable_getnext(m, v27)
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					F_systable_endscan(m, v27)
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_relation_close(m, v13, int32(1))
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							m.G0 = v10 - int32(-64)
							return base.B2i32(v29 != int32(0))
						}
					}
				}
			}
		}
	}
}
func Fn14312(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int64
	_ = v119
	var v120 int32
	_ = v120
	var v123 int64
	_ = v123
	var v124 int32
	_ = v124
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v131 int64
	_ = v131
	var v132 int32
	_ = v132
	var v135 int64
	_ = v135
	var v139 int64
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v155 int64
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v165 int64
	_ = v165
	var v166 int32
	_ = v166
	var v169 int64
	_ = v169
	var v170 int64
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v257 int64
	_ = v257
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v272 int64
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int64
	_ = v279
	var v280 int32
	_ = v280
	var v281 int64
	_ = v281
	var v282 int32
	_ = v282
	var v283 int64
	_ = v283
	var v284 int32
	_ = v284
	var v285 int64
	_ = v285
	var v289 int64
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v297 int64
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int64
	_ = v304
	var v308 int32
	_ = v308
	var v309 int64
	_ = v309
	var v310 int64
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int64
	_ = v318
	var v328 int64
	_ = v328
	var v329 int64
	_ = v329
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v344 int64
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int64
	_ = v351
	var v352 int32
	_ = v352
	var v353 int64
	_ = v353
	var v354 int32
	_ = v354
	var v355 int64
	_ = v355
	var v356 int32
	_ = v356
	var v357 int64
	_ = v357
	var v361 int64
	_ = v361
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v369 int64
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int64
	_ = v376
	var v380 int32
	_ = v380
	var v381 int64
	_ = v381
	var v382 int64
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v390 int64
	_ = v390
	var v400 int64
	_ = v400
	var v409 int64
	_ = v409
	var v421 int64
	_ = v421
	v9 = int64(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v11 = v10 & l3
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v12&l3 != 0 {
		if v11 != 0 {
			return int32(0)
		} else {
			v17 = l1 + int32(8)
			if int32(7) < l2 {
				v257 = int64(0)
				v258 = int32(0)
				if l2 == v258 {
					v328 = int64(0)
				} else {
					v265 = l2 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(l2) {
						v270 = v17
						v272 = v257
						v275 = v258
						for {
							v276 = int32(4)
							v277 = v270 + v276
							v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+3)))
							v279 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v278)+uint32(_c_Fn14312[0]))))
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+2)))
							v281 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v280)+uint32(_c_Fn14312[0]))))
							v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+1)))
							v283 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v282)+uint32(_c_Fn14312[0]))))
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
							v285 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v284)+uint32(_c_Fn14312[0]))))
							v289 = v279 + (v281 + (v283 + (v272 + v285)))
							v291 = v275 + v276
							if v291 != l2&int32(-4) {
								v270 = v277
								v272 = v289
								v275 = v291
								continue
							} else {
								break
							}
							break
						}
						if v265 == int32(0) {
							v318 = v289
						} else {
							v295 = v277
							v297 = v289
							v302 = v295
							v303 = int32(0)
							v304 = v297
							for {
								v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
								v309 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v308)+uint32(_c_Fn14312[0]))))
								v310 = v304 + v309
								v311 = int32(1)
								v314 = v303 + v311
								if v314 != v265 {
									v302 = v302 + v311
									v303 = v314
									v304 = v310
									continue
								} else {
									break
								}
								break
							}
							v318 = v310
						}
					} else {
						v295 = v17
						v297 = v257
						v302 = v295
						v303 = int32(0)
						v304 = v297
						for {
							v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
							v309 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v308)+uint32(_c_Fn14312[0]))))
							v310 = v304 + v309
							v311 = int32(1)
							v314 = v303 + v311
							if v314 != v265 {
								v302 = v302 + v311
								v303 = v314
								v304 = v310
								continue
							} else {
								break
							}
							break
						}
						v318 = v310
					}
					v328 = v318
				}
				v421 = v328
			} else {
				if l2 == int32(0) {
					v421 = v9
				} else {
					v25 = l2 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(l2) {
						v31 = int32(0)
						v32 = v17
						v39 = v9
						for {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+3)))
							v43 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v40)+uint32(_c_Fn14312[0]))))
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+2)))
							v47 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v44)+uint32(_c_Fn14312[0]))))
							v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
							v51 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v48)+uint32(_c_Fn14312[0]))))
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
							v55 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v52)+uint32(_c_Fn14312[0]))))
							v59 = v43 + (v47 + (v51 + (v39 + v55)))
							v60 = int32(4)
							v61 = v32 + v60
							v63 = v31 + v60
							if v63 != l2&int32(-4) {
								v31 = v63
								v32 = v61
								v39 = v59
								continue
							} else {
								break
							}
							break
						}
						if v25 == int32(0) {
							v421 = v59
						} else {
							v68 = v61
							v75 = v59
							v77 = int32(0)
							v78 = v68
							v85 = v75
							for {
								v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
								v89 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v86)+uint32(_c_Fn14312[0]))))
								v90 = v85 + v89
								v91 = int32(1)
								v94 = v77 + v91
								if v94 != v25 {
									v77 = v94
									v78 = v78 + v91
									v85 = v90
									continue
								} else {
									break
								}
								break
							}
							v421 = v90
						}
					} else {
						v68 = v17
						v75 = v9
						v77 = int32(0)
						v78 = v68
						v85 = v75
						for {
							v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
							v89 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v86)+uint32(_c_Fn14312[0]))))
							v90 = v85 + v89
							v91 = int32(1)
							v94 = v77 + v91
							if v94 != v25 {
								v77 = v94
								v78 = v78 + v91
								v85 = v90
								continue
							} else {
								break
							}
							break
						}
						v421 = v90
					}
				}
			}
			return l2<<(uint(int32(3))%32) - base.I32_wrap_i64(v421)
		}
	} else {
		if v11 != 0 {
			v97 = l0 + int32(8)
			if int32(7) < l2 {
				v329 = int64(0)
				v330 = int32(0)
				if l2 == v330 {
					v400 = int64(0)
				} else {
					v337 = l2 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(l2) {
						v342 = v97
						v344 = v329
						v347 = v330
						for {
							v348 = int32(4)
							v349 = v342 + v348
							v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+3)))
							v351 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v350)+uint32(_c_Fn14312[0]))))
							v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+2)))
							v353 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v352)+uint32(_c_Fn14312[0]))))
							v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+1)))
							v355 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v354)+uint32(_c_Fn14312[0]))))
							v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342))))
							v357 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v356)+uint32(_c_Fn14312[0]))))
							v361 = v351 + (v353 + (v355 + (v344 + v357)))
							v363 = v347 + v348
							if v363 != l2&int32(-4) {
								v342 = v349
								v344 = v361
								v347 = v363
								continue
							} else {
								break
							}
							break
						}
						if v337 == int32(0) {
							v390 = v361
						} else {
							v367 = v349
							v369 = v361
							v374 = v367
							v375 = int32(0)
							v376 = v369
							for {
								v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374))))
								v381 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v380)+uint32(_c_Fn14312[0]))))
								v382 = v376 + v381
								v383 = int32(1)
								v386 = v375 + v383
								if v386 != v337 {
									v374 = v374 + v383
									v375 = v386
									v376 = v382
									continue
								} else {
									break
								}
								break
							}
							v390 = v382
						}
					} else {
						v367 = v97
						v369 = v329
						v374 = v367
						v375 = int32(0)
						v376 = v369
						for {
							v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374))))
							v381 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v380)+uint32(_c_Fn14312[0]))))
							v382 = v376 + v381
							v383 = int32(1)
							v386 = v375 + v383
							if v386 != v337 {
								v374 = v374 + v383
								v375 = v386
								v376 = v382
								continue
							} else {
								break
							}
							break
						}
						v390 = v382
					}
					v400 = v390
				}
				v409 = v400
			} else {
				if l2 == int32(0) {
					v409 = v9
				} else {
					v105 = l2 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(l2) {
						v111 = int32(0)
						v112 = v97
						v119 = v9
						for {
							v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+3)))
							v123 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v120)+uint32(_c_Fn14312[0]))))
							v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+2)))
							v127 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v124)+uint32(_c_Fn14312[0]))))
							v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+1)))
							v131 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v128)+uint32(_c_Fn14312[0]))))
							v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
							v135 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v132)+uint32(_c_Fn14312[0]))))
							v139 = v123 + (v127 + (v131 + (v119 + v135)))
							v140 = int32(4)
							v141 = v112 + v140
							v143 = v111 + v140
							if v143 != l2&int32(-4) {
								v111 = v143
								v112 = v141
								v119 = v139
								continue
							} else {
								break
							}
							break
						}
						if v105 == int32(0) {
							v409 = v139
						} else {
							v148 = v141
							v155 = v139
							v157 = int32(0)
							v158 = v148
							v165 = v155
							for {
								v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
								v169 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v166)+uint32(_c_Fn14312[0]))))
								v170 = v165 + v169
								v171 = int32(1)
								v174 = v157 + v171
								if v174 != v105 {
									v157 = v174
									v158 = v158 + v171
									v165 = v170
									continue
								} else {
									break
								}
								break
							}
							v409 = v170
						}
					} else {
						v148 = v97
						v155 = v9
						v157 = int32(0)
						v158 = v148
						v165 = v155
						for {
							v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v169 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v166)+uint32(_c_Fn14312[0]))))
							v170 = v165 + v169
							v171 = int32(1)
							v174 = v157 + v171
							if v174 != v105 {
								v157 = v174
								v158 = v158 + v171
								v165 = v170
								continue
							} else {
								break
							}
							break
						}
						v409 = v170
					}
				}
			}
			return l2<<(uint(int32(3))%32) - base.I32_wrap_i64(v409)
		} else {
			if l2 <= int32(0) {
				return int32(0)
			} else {
				v180 = int32(8)
				v181 = l1 + v180
				v183 = l0 + v180
				v184 = int32(0)
				if l2 != int32(1) {
					v193 = v184
					v194 = v184
					v195 = int32(0)
					for {
						v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+v183))))
						v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+v181))))
						v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203^v205)+uint32(_c_Fn14312[0]))))
						v212 = v194 | int32(1)
						v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181+v212))))
						v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212+v183))))
						v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214^v216)+uint32(_c_Fn14312[0]))))
						v221 = v193 + v209 + v220
						v222 = int32(2)
						v223 = v194 + v222
						v225 = v195 + v222
						if v225 != l2&int32(2147483646) {
							v193 = v221
							v194 = v223
							v195 = v225
							continue
						} else {
							break
						}
						break
					}
					if l2&int32(1) == int32(0) {
						v247 = v221
					} else {
						v229 = v221
						v230 = v223
						v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v181))))
						v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v183))))
						v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239^v241)+uint32(_c_Fn14312[0]))))
						v247 = v229 + v245
					}
				} else {
					v229 = v184
					v230 = v184
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v181))))
					v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v183))))
					v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239^v241)+uint32(_c_Fn14312[0]))))
					v247 = v229 + v245
				}
				return v247
			}
		}
	}
}
func Fn14318(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v15 int64
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		*(*uint32)(unsafe.Add(mBase, uint32(v8)+12)) = uint32(v15)
		v21 = F_get_worker(m, v11, int32(0), v8+int32(12), int32(1), l1)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			if v21 != 0 {
				v27 = base.I64_extend_i32_u(v21)
			} else {
				v24 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
				v27 = int64(0)
			}
			m.G0 = v8 + int32(16)
			return v27
		}
	}
}
func Fn14321(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v4
	v11 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v11
	v16 = v7 + int32(8)
	F_pushJsonbValue(m, v16, l2, v4)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		F_pushJsonbValue(m, v16, l1, int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			v26 = F_JsonbValueToJsonb(m, v25)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int64(0)
			} else {
				m.G0 = v7 + int32(32)
				return base.I64_extend_i32_u(v26)
			}
		}
	}
}
func Fn14327(m *base.Module, l0 int32, l1 int64) int64 {
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int64
	_ = v52
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
	v8 = int32(8)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v11 = int32(16)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)))
	v15 = v7<<(uint(v8)%32) | v10<<(uint(v11)%32) | v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+2)))
	v25 = v17<<(uint(v8)%32) | v20<<(uint(v11)%32) | v24
	if base.Ui32(v15) < base.Ui32(v25) {
		v52 = l1
	} else {
		if base.Ui32(v25) < base.Ui32(v15) {
			v52 = int64(1)
		} else {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)))
			v31 = int32(8)
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)))
			v34 = int32(16)
			v37 = v29 | (v30<<(uint(v31)%32) | v33<<(uint(v34)%32))
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+5)))
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+4)))
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+3)))
			v46 = v38 | (v39<<(uint(v31)%32) | v42<<(uint(v34)%32))
			if base.Ui32(v37) < base.Ui32(v46) {
				v52 = l1
			} else {
				v52 = base.I64_extend_i32_u(base.B2i32(base.Ui32(v46) < base.Ui32(v37)))
			}
		}
	}
	return v52
}
func Fn14329(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
		if v18 == int32(1) {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
			if v24 == int32(18) {
				v27 = int32(16)
			} else {
				v27 = int32(0)
			}
			if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v34 = int32(4)
			} else {
				v34 = v27
			}
			v47 = v34
		} else {
			v35 = int32(1)
			if v18&v35 != 0 {
				v47 = int32(base.Ui32(v18)>>(uint(v35)%32)) - v35
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v48 = int32(1)
		if v18&v48 != 0 {
			v52 = v48
		} else {
			v52 = int32(4)
		}
		v58 = F_pg_md5_hash(m, v12+v52, v47, v7+int32(-48), v7+int32(-52))
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return int64(0)
		} else {
			if v58 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(2600))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a_Fn14329_0)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v71
						F_errmsg(m, int32(_a_Fn14329_1), v9)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_Fn14329_2), l2, l1)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v81 = F_cstring_to_text(m, v7+int32(-48))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int64(0)
				} else {
					m.G0 = v9 - int32(-64)
					return base.I64_extend_i32_u(v81)
				}
			}
		}
	}
}
func Fn14330(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = F_strlen(m, v6)
		mBase = m.M
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v14 = F_RE_compile_and_cache(m, v8, l1, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v19 = F_palloc_mul(m, int32(4), v12+int32(1))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				v21 = F_pg_mb2wchar_with_len(m, v6, v19, v12)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					v23 = int32(0)
					v26 = F_RE_wchar_execute(m, v19, v21, v23, v23, v23)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						F_pfree(m, v19)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v26)
						}
					}
				}
			}
		}
	}
}
func Fn14336(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int64
	_ = v28
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v11 != 0 {
		v14 = int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v14 = v13
	}
	v17 = F_numeric_stddev_internal(m, v14, l2, l1, v9+int32(15))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
		if v21 == int32(1) {
			v24 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
			v28 = int64(0)
		} else {
			v28 = base.I64_extend_i32_u(v17)
		}
		m.G0 = v9 + int32(16)
		return v28
	}
}
func Fn14343(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(34209793)
	*(*uint32)(unsafe.Add(mBase, uint32(v7)+8)) = uint32(v9)
	v14 = int64(base.Ui64(v9) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v7)+4)) = uint32(v14)
	v17 = *(*int32)(unsafe.Add(mBase, _c_Fn14343[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v17
	v20 = F_LockRelease(m, v7, l1, int32(1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int64(0)
	} else {
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v20)
	}
}
func Fn14349(m *base.Module, l0 int64) int32 {
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
	v4 = int32(_a_Fn14349_0)
	v5 = int32(_a_Fn14349_1)
	v6 = *(*int64)(unsafe.Add(mBase, _c_Fn14349[0]))
	v8 = *(*int64)(unsafe.Add(mBase, _c_Fn14349[1]))
	v9 = v6 ^ v8
	*(*int64)(unsafe.Add(mBase, _c_Fn14349[1])) = base.I64_rotl(v9, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_Fn14349[0])) = v9<<(uint(int64(16))%64) ^ base.I64_rotl(v6, int64(24)) ^ v9
	return base.I32_wrap_i64(int64(base.Ui64(base.I64_rotl(v6*int64(5), int64(7))*int64(9)) >> (uint(l0) % 64)))
}
func Fn14356(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 float64
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 float64
	_ = v146
	var v150 int64
	_ = v150
	var v155 float64
	_ = v155
	var v156 float64
	_ = v156
	var v157 float64
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v163 int32
	_ = v163
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	v3 = int32(0)
	v13 = float64(0)
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = F_palloc0(m, v20<<(uint(int32(2))%32))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		if base.B2i32(v25 == int64(0)) == int32(0) {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v36 = v3
			v40 = v3
			v42 = v3
			for {
				v49 = v31 + v36<<(uint(int32(3))%32)
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)))
				if v50 == int32(1) {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
					v55 = int32(16)
					v59 = (int32(base.Ui32(v54)>>(uint(v55)%32)) ^ v54) * int32(-2048144789)
					v64 = (int32(base.Ui32(v59)>>(uint(int32(13))%32)) ^ v59) * int32(-1028477387)
					v68 = v53 & (int32(base.Ui32(v64)>>(uint(v55)%32)) ^ v64)
					v71 = v23 + v68<<(uint(int32(2))%32)
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
					*(*int32)(unsafe.Add(mBase, uint32(v71))) = v72 + int32(1)
					if base.Ui32(v36) < base.Ui32(v68) {
						v79 = base.I32_wrap_i64(v25)
					} else {
						v79 = int32(0)
					}
					v80 = v36 - v68 + v79
					if base.Ui32(v42) < base.Ui32(v80) {
						v82 = v80
					} else {
						v82 = v42
					}
					v85 = v80 + v40
					v87 = v82
				} else {
					v85 = v40
					v87 = v42
				}
				v89 = v36 + int32(1)
				if base.Ui64(base.I64_extend_i32_u(v89)) < base.Ui64(v25) {
					v36 = v89
					v40 = v85
					v42 = v87
					continue
				} else {
					break
				}
				break
			}
			v92 = int32(0)
			v97 = v92
			v99 = v92
			v101 = v92
			for {
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v23+v101<<(uint(int32(2))%32))))
				v115 = v113 - int32(1)
				if base.Ui32(v99) < base.Ui32(v115) {
					v117 = v115
				} else {
					v117 = v99
				}
				if v113 != 0 {
					v118 = v117
				} else {
					v118 = v99
				}
				v122 = v113 - base.B2i32(v113 != int32(0)) + v97
				v124 = v101 + int32(1)
				if base.Ui64(base.I64_extend_i32_u(v124)) < base.Ui64(v25) {
					v97 = v122
					v99 = v118
					v101 = v124
					continue
				} else {
					break
				}
				break
			}
			v129 = v122
			v131 = v118
			v135 = v85
			v137 = v87
		} else {
			v129 = v3
			v131 = v3
			v135 = v3
			v137 = v3
		}
		F_pfree(m, v23)
		mBase = m.M
		v143 = m.ExcPending
		if v143 != 0 {
			return
		} else {
			v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v144 != 0 {
				v146 = base.F64_convert_i32_u(v144)
				v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				v155 = base.F64_div(base.F64_convert_i32_u(v129), v146)
				v156 = base.F64_div(base.F64_convert_i32_u(v135), v146)
				v157 = base.F64_div(v146, base.F64_convert_i64_u(v150))
			} else {
				v155 = v13
				v156 = v13
				v157 = float64(0)
			}
			v160 = F_errstart(m, int32(15), int32(0))
			mBase = m.M
			v161 = m.ExcPending
			if v161 != 0 {
				return
			} else {
				if v160 != 0 {
					v162 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
					v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*float64)(unsafe.Add(mBase, uint32(v18)+48)) = v155
					*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v131
					*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v129
					*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v156
					*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v137
					*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v135
					*(*float64)(unsafe.Add(mBase, uint32(v18)+16)) = v157
					*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v163
					*(*int64)(unsafe.Add(mBase, uint32(v18))) = v162
					F_errmsg_internal(m, int32(_a_Fn14356_0), v18)
					mBase = m.M
					v175 = m.ExcPending
					if v175 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_Fn14356_1), int32(1190), l1)
						mBase = m.M
						v179 = m.ExcPending
						if v179 != 0 {
							return
						} else {
							m.G0 = v18 - int32(-64)
							return
						}
					}
				} else {
					m.G0 = v18 - int32(-64)
					return
				}
			}
		}
	}
}
func Fn14358(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
	if v7 <= l3 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
		switch v9 - int32(3) {
		case 0, 2:
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v15 = v14
		default:
			v15 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v15
	} else {
	}
	return int32(0)
}
func Fn14363(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	v5 = int32(0)
	v7 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v7 == v5 {
		v29 = v5
		return v29
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v11-int32(2) <= v10 {
			v29 = v5
			return v29
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v11-int32(1)))))
			if v19 != l3 {
				v29 = v5
				return v29
			} else {
				v22 = F_find_among_b(m, l0, l2, l1, int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v29 = base.B2i32(v22 != int32(0))
					return v29
				}
			}
		}
	}
}
func Fn14365(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v4 = int32(0)
	v6 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v6 == v4 {
		v29 = v4
		return v29
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v11 = v9 - int32(1)
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v11 <= v12 {
			v29 = v4
			return v29
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v11))))
			if v16 != l2 {
				v29 = v4
				return v29
			} else {
				v20 = F_find_among_b(m, l0, l1, int32(4), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					if v20 == int32(0) {
						v29 = v4
					} else {
						v27 = Fn14364(m, l0, int32(121))
						mBase = m.M
						v29 = v27
					}
					return v29
				}
			}
		}
	}
}
func Fn14369(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
			if v22 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				if v23 == v20 {
					v33 = v22
					v34 = F_range_union_internal(m, v33, v13, v18, l1)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						m.G0 = v10 + int32(16)
						return base.I64_extend_i32_u(v34)
					}
				} else {
					v26 = F_lookup_type_cache(m, v20, int32(2048))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+200))
						if v28 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
								F_errmsg_internal(m, int32(_a_Fn14369_0), v10)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_Fn14369_1), int32(1946), int32(_a_Fn14369_2))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v26
							v33 = v26
							v34 = F_range_union_internal(m, v33, v13, v18, l1)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
							} else {
								m.G0 = v10 + int32(16)
								return base.I64_extend_i32_u(v34)
							}
						}
					}
				}
			} else {
				v26 = F_lookup_type_cache(m, v20, int32(2048))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+200))
					if v28 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
							F_errmsg_internal(m, int32(_a_Fn14369_0), v10)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_Fn14369_1), int32(1946), int32(_a_Fn14369_2))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v26
						v33 = v26
						v34 = F_range_union_internal(m, v33, v13, v18, l1)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							m.G0 = v10 + int32(16)
							return base.I64_extend_i32_u(v34)
						}
					}
				}
			}
		}
	}
}
func Fn14376(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v2 = l1
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	if v9 <= v6+int32(1) {
		F_appendStringInfoChar(m, v5, v2)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		*(*uint8)(unsafe.Add(mBase, uint32(v17+v6))) = uint8(v2)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
		v23 = v21 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v23
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		v27 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v25+v23))) = uint8(v27)
		return v27
	}
}
func Fn14378(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v13 = F_pg_detoast_datum_packed(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v15 == int32(1) {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
				if v21 == int32(18) {
					v24 = int32(16)
				} else {
					v24 = int32(0)
				}
				if base.Ui32((v21-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v31 = int32(4)
				} else {
					v31 = v24
				}
				v44 = v31
			} else {
				v32 = int32(1)
				if v15&v32 != 0 {
					v44 = int32(base.Ui32(v15)>>(uint(v32)%32)) - v32
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v46 = F_RE_compile_and_cache(m, v13, l1, v45)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int64(0)
			} else {
				v51 = F_palloc_mul(m, int32(4), v44+int32(1))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int64(0)
				} else {
					v53 = int32(1)
					if v15&v53 != 0 {
						v57 = v53
					} else {
						v57 = int32(4)
					}
					v59 = F_pg_mb2wchar_with_len(m, v8+v57, v51, v44)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int64(0)
					} else {
						v61 = int32(0)
						v64 = F_RE_wchar_execute(m, v51, v59, v61, v61, v61)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int64(0)
						} else {
							F_pfree(m, v51)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v64)
							}
						}
					}
				}
			}
		}
	}
}
func Fn14381(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int64 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v25 int64
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v87 int64
	_ = v87
	var v94 int64
	_ = v94
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v118 int64
	_ = v118
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v128 int64
	_ = v128
	var v139 int64
	_ = v139
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui64(v20-int64(9223372036854775807)) < base.Ui64(int64(2)) {
		v139 = v20
		m.G0 = v18 + int32(16)
		return v139
	} else {
		v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		if base.Ui64(v25-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v151 = m.ExcPending
			if v151 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v154 = m.ExcPending
				if v154 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_Fn14381_0), int32(0))
					mBase = m.M
					v158 = m.ExcPending
					if v158 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_Fn14381_1), l7, l1)
						mBase = m.M
						v161 = m.ExcPending
						if v161 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
			if v31 != 0 {
				if v31 != int32(2147483647) {
					if v31 != int32(-2147483648) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_Fn14381_2), int32(0))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_Fn14381_1), l8, l1)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
						if v36 != int32(-2147483648) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_Fn14381_2), int32(0))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_Fn14381_1), l8, l1)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v39 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
							if v39 != int64(-9223372036854775807-1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_Fn14381_2), int32(0))
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_Fn14381_1), l8, l1)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v223 = m.ExcPending
								if v223 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v226 = m.ExcPending
									if v226 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_Fn14381_3), int32(0))
										mBase = m.M
										v230 = m.ExcPending
										if v230 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_Fn14381_1), l2, l1)
											mBase = m.M
											v233 = m.ExcPending
											if v233 != 0 {
												return int64(0)
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
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
					if v42 != int32(2147483647) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_Fn14381_2), int32(0))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_Fn14381_1), l8, l1)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v45 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
						if v45 == int64(9223372036854775807) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v223 = m.ExcPending
							if v223 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v226 = m.ExcPending
								if v226 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_Fn14381_3), int32(0))
									mBase = m.M
									v230 = m.ExcPending
									if v230 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_Fn14381_1), l2, l1)
										mBase = m.M
										v233 = m.ExcPending
										if v233 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_Fn14381_2), int32(0))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_Fn14381_1), l8, l1)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int64(0)
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
			} else {
				v64 = int64(*(*int32)(unsafe.Add(mBase, uint32(v30)+8)))
				v73 = int64(32)
				v74 = int64(20)
				v76 = int64(base.Ui64(v64) >> (uint(v73) % 64))
				v79 = int64(4294967295)
				v80 = int64(500654080)
				v82 = v64 & v79
				v83 = v80 * v82
				v87 = int64(base.Ui64(v83)>>(uint(v73)%64)) + v80*v76
				v94 = v82*v74 + v87&v79
				*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v64*int64(0) + v64>>(uint(int64(63))%64)*int64(86400000000) + v74*v76 + int64(base.Ui64(v87)>>(uint(v73)%64)) + int64(base.Ui64(v94)>>(uint(v73)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v18))) = v83&v79 | v94<<(uint(v73)%64)
				v105 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
				v106 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
				if v105 != v106>>(uint(int64(63))%64) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v167 = m.ExcPending
					if v167 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_Fn14381_4), int32(0))
							mBase = m.M
							v174 = m.ExcPending
							if v174 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_Fn14381_1), l6, l1)
								mBase = m.M
								v177 = m.ExcPending
								if v177 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v110 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
					v113 = v106 + v110
					if base.B2i32(v110 < int64(0)) != base.B2i32(v113 < v106) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v167 = m.ExcPending
						if v167 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v170 = m.ExcPending
							if v170 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_Fn14381_4), int32(0))
								mBase = m.M
								v174 = m.ExcPending
								if v174 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_Fn14381_1), l6, l1)
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						if v113 <= int64(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v181 = m.ExcPending
							if v181 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_Fn14381_5), int32(0))
									mBase = m.M
									v188 = m.ExcPending
									if v188 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_Fn14381_1), l5, l1)
										mBase = m.M
										v191 = m.ExcPending
										if v191 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v118 = v20 - v25
							if base.B2i32(v118 < v20) != base.B2i32(int64(0) < v25) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v195 = m.ExcPending
								if v195 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v198 = m.ExcPending
									if v198 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_Fn14381_4), int32(0))
										mBase = m.M
										v202 = m.ExcPending
										if v202 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_Fn14381_1), l4, l1)
											mBase = m.M
											v205 = m.ExcPending
											if v205 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v123 = base.I64_rem_s(v118, v113)
								v125 = v118 - v123 + v25
								if int64(0) <= v123 {
									v139 = v125
									m.G0 = v18 + int32(16)
									return v139
								} else {
									v128 = v125 - v113
									if base.B2i32(v128 < v125)^base.B2i32(int64(0) < v113)|base.B2i32(base.Ui64(v128-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615))) != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v209 = m.ExcPending
										if v209 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_Fn14381_6), int32(0))
												mBase = m.M
												v216 = m.ExcPending
												if v216 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_Fn14381_1), l3, l1)
													mBase = m.M
													v219 = m.ExcPending
													if v219 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v139 = v128
										m.G0 = v18 + int32(16)
										return v139
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
func Fn14390(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l2
		*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v16)
		v19 = F_text_to_cstring(m, v12)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v25 = F_parse_tsquery(m, v19, int32(1273), v9+int32(8), l1, int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int64(0)
			} else {
				m.G0 = v9 + int32(16)
				return base.I64_extend_i32_u(v25)
			}
		}
	}
}
func Fn14396(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v6 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v15)+79)))
			F_ReleaseCatCache(m, v6)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v17 == l1)
			}
		}
	}
}
func Fn14400(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		if v12 == int32(1) {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
			if v18 == int32(18) {
				v21 = int32(16)
			} else {
				v21 = int32(0)
			}
			if base.Ui32((v18-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v28 = int32(4)
			} else {
				v28 = v21
			}
			v41 = v28
		} else {
			v29 = int32(1)
			if v12&v29 != 0 {
				v41 = int32(base.Ui32(v12)>>(uint(v29)%32)) - v29
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v41 = int32(base.Ui32(v35)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v43 = int32(1)
		if v12&v43 != 0 {
			v47 = v43
		} else {
			v47 = int32(4)
		}
		v49 = F_uuid_generate_internal(m, l1, base.I32_wrap_i64(v6), v8+v47, v41)
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int64(0)
		} else {
			return v49
		}
	}
}
func Fn14411(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	F_errstart_cold(m, int32(21), int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		F_errcode(m, int32(1088))
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			F_errmsg(m, int32(_a_Fn14411_0), int32(0))
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				v19 = F_errdetail(m, int32(_a_Fn14411_1), int32(0))
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_Fn14411_2), l2, l1)
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
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
