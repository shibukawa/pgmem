package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetJsonTableExecContext(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 == int32(414) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		if v12 != int32(418352867) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
				F_errmsg_internal(m, int32(_a_F_GetJsonTableExecContext_0), v6)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_GetJsonTableExecContext_1), int32(_a_F_GetJsonTableExecContext_2), int32(_a_F_GetJsonTableExecContext_3))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v6 + int32(32)
			return v11
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l1
			F_errmsg_internal(m, int32(_a_F_GetJsonTableExecContext_0), v6+int32(16))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_GetJsonTableExecContext_1), int32(4094), int32(_a_F_GetJsonTableExecContext_3))
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
}
func F__equalJsonArrayQueryConstructor(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v27 = v3
			return v27
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v14 = F_equal(m, v12, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 == int32(0) {
					v27 = v3
					return v27
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v20 = F_equal(m, v18, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						if v20 == int32(0) {
							v27 = v3
						} else {
							v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
							v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
							v27 = base.B2i32(v24 == v25)
						}
						return v27
					}
				}
			}
		}
	}
}
func F_add_json(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l3 != 0 {
		if l1 != 0 {
			v12 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v12
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v12
			v26 = int32(0)
			v27 = v12
			F_datum_to_json_internal(m, l0, l1, l2, v26, v27, l4)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				m.G0 = v10 + int32(16)
				return
			}
		} else {
			F_json_categorize_type(m, l3, int32(0), v10+int32(12), v10+int32(8))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				v26 = v24
				v27 = v25
				F_datum_to_json_internal(m, l0, l1, l2, v26, v27, l4)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					m.G0 = v10 + int32(16)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_add_json_0), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_add_json_1), int32(611), int32(_a_F_add_json_2))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
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
func F_executeJsonPath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	v6 = l5
	v8 = l7
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	F_jspInit(m, v13+int32(32), l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v22 = l4 + int32(4)
		v25 = F_JsonbExtractScalar(m, v22, v13+int32(12))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			if v25 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v22
				*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(18)
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
				if v32 == int32(1) {
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
					if v38 == int32(18) {
						v41 = int32(16)
					} else {
						v41 = int32(0)
					}
					if base.Ui32((v38-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v48 = int32(4)
					} else {
						v48 = v41
					}
					v61 = v48
				} else {
					v49 = int32(1)
					if v32&v49 != 0 {
						v61 = int32(base.Ui32(v32)>>(uint(v49)%32)) - v49
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						v61 = int32(base.Ui32(v55)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v61
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = l1
			v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int64)(unsafe.Add(mBase, uint32(v13)+76)) = int64(0)
			v71 = int32(base.Ui32(v67) >> (uint(int32(31)) % 32))
			*(*uint8)(unsafe.Add(mBase, uint32(v13)+93)) = uint8(v71)
			*(*uint8)(unsafe.Add(mBase, uint32(v13)+92)) = uint8(v71)
			v75 = v13 + int32(12)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v75
			*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v75
			v78 = m.T0[l3].(func(*base.Module, int32) int32)(m, l1)
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v13)+95)) = uint8(v8)
				*(*uint8)(unsafe.Add(mBase, uint32(v13)+94)) = uint8(v6)
				*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v78 + int32(1)
				v87 = int32(0)
				if l6|base.B2i32(v67 < v87) == v87 {
					*(*int64)(unsafe.Add(mBase, uint32(v13))) = int64(0)
					v100 = F_executeItemOptUnwrapTarget(m, v13+int32(60), v13+int32(32), v75, v13, int32(0))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
					} else {
						if v100 == int32(2) {
							v117 = int32(2)
						} else {
							v104 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
							v105 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v117 = base.B2i32(v104|v105 == int32(0))
						}
						m.G0 = v13 + int32(96)
						return v117
					}
				} else {
					v115 = F_executeItemOptUnwrapTarget(m, v13+int32(60), v13+int32(32), v13+int32(12), l6, v71)
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						v117 = v115
						m.G0 = v13 + int32(96)
						return v117
					}
				}
			}
		}
	}
}
func F_freeJsonLexContext(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	v2 = int32(0)
	if base.B2i32(l0 == v2)|base.B2i32(l0 == int32(_a_F_freeJsonLexContext_0)) == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v12&int32(2) != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_free_attrmap(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v18 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	return
L8:
	;
	goto L6
L9:
	;
	F_free_attrmap(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v21 != int32(1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v89&int32(1) != 0 {
		goto L48
	} else {
		goto L49
	}
L14:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	F_pfree(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v28 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_pfree(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v32 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	F_pfree(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v35&int32(4) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	if v66 != 0 {
		goto L34
	} else {
		goto L35
	}
L25:
	;
	v40 = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v41 < v40 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v45 = v40
	v46 = v41
	goto L27
L27:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v45<<(uint(int32(2))%32))))
	if v53 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L24
L29:
	;
	F_pfree(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L7
	} else {
		goto L32
	}
L30:
	;
	v57 = v46
	goto L31
L31:
	;
	v59 = v45 + int32(1)
	if v59 <= v57 {
		v45 = v59
		v46 = v57
		goto L27
	} else {
		goto L33
	}
L32:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v57 = v56
	goto L31
L33:
	;
	goto L28
L34:
	;
	F_pfree(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L7
	} else {
		goto L37
	}
L35:
	;
	v70 = v65
	goto L36
L36:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	if v71 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v70 = v69
	goto L36
L38:
	;
	F_pfree(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L7
	} else {
		goto L41
	}
L39:
	;
	v75 = v70
	goto L40
L40:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+24))
	if v76 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v75 = v74
	goto L40
L42:
	;
	F_pfree(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L7
	} else {
		goto L45
	}
L43:
	;
	v82 = v75
	goto L44
L44:
	;
	F_pfree(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L7
	} else {
		goto L47
	}
L45:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v79 == int32(0) {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	v82 = v79
	goto L44
L47:
	;
	goto L13
L48:
	;
	F_pfree(m, l0)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L7
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	base.MemoryFill(m, l0, int32(0), int32(68))
	goto L3
L51:
	;
	return
}
func F_get_json_expr_options(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		if base.Ui32(v8) <= base.Ui32(int32(3)) {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v8<<(uint(int32(2))%32))+uint32(_c_F_get_json_expr_options[0])))
			F_appendStringInfoString(m, v11, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
				if v20 != 0 {
					v21 = int32(_a_F_get_json_expr_options_0)
				} else {
					v21 = int32(_a_F_get_json_expr_options_1)
				}
				F_appendStringInfoString(m, v17, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v25 == int32(0) {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v33 == int32(0) {
							return
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
							if v36 == l2 {
								return
							} else {
								F_get_json_behavior(m, v33, l1, int32(_a_F_get_json_expr_options_2))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
						if v28 == l2 {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v33 == int32(0) {
								return
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
								if v36 == l2 {
									return
								} else {
									F_get_json_behavior(m, v33, l1, int32(_a_F_get_json_expr_options_2))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							F_get_json_behavior(m, v25, l1, int32(_a_F_get_json_expr_options_3))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v33 == int32(0) {
									return
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
									if v36 == l2 {
										return
									} else {
										F_get_json_behavior(m, v33, l1, int32(_a_F_get_json_expr_options_2))
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
			if v20 != 0 {
				v21 = int32(_a_F_get_json_expr_options_0)
			} else {
				v21 = int32(_a_F_get_json_expr_options_1)
			}
			F_appendStringInfoString(m, v17, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v25 == int32(0) {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v33 == int32(0) {
						return
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
						if v36 == l2 {
							return
						} else {
							F_get_json_behavior(m, v33, l1, int32(_a_F_get_json_expr_options_2))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
					if v28 == l2 {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v33 == int32(0) {
							return
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
							if v36 == l2 {
								return
							} else {
								F_get_json_behavior(m, v33, l1, int32(_a_F_get_json_expr_options_2))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						F_get_json_behavior(m, v25, l1, int32(_a_F_get_json_expr_options_3))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v33 == int32(0) {
								return
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
								if v36 == l2 {
									return
								} else {
									F_get_json_behavior(m, v33, l1, int32(_a_F_get_json_expr_options_2))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v25 == int32(0) {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v33 == int32(0) {
				return
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
				if v36 == l2 {
					return
				} else {
					F_get_json_behavior(m, v33, l1, int32(_a_F_get_json_expr_options_2))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
			if v28 == l2 {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v33 == int32(0) {
					return
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
					if v36 == l2 {
						return
					} else {
						F_get_json_behavior(m, v33, l1, int32(_a_F_get_json_expr_options_2))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				F_get_json_behavior(m, v25, l1, int32(_a_F_get_json_expr_options_3))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v33 == int32(0) {
						return
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
						if v36 == l2 {
							return
						} else {
							F_get_json_behavior(m, v33, l1, int32(_a_F_get_json_expr_options_2))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_get_json_returning(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v30 int32
	_ = v30
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
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v11 == int32(0) {
		m.G0 = v9 + int32(32)
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v15 = F_format_type_with_typemod(m, v11, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v15
			F_appendStringInfo(m, l1, int32(_a_F_get_json_returning_0), v9+int32(16))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
				if l2 != 0 {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v27 == int32(3802) {
						v30 = int32(-3)
					} else {
						v30 = int32(-2)
					}
					if v24&v30 != 0 {
						if v24 == int32(2) {
							v38 = int32(_a_F_get_json_returning_1)
						} else {
							v38 = int32(_a_F_get_json_returning_2)
						}
						F_appendStringInfoString(m, l1, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
							switch v42 {
							case 0:
								m.G0 = v9 + int32(32)
								return
							default:
								v45 = int32(_a_F_get_json_returning_3)
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v45
								F_appendStringInfo(m, l1, int32(_a_F_get_json_returning_4), v9)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							case 2:
								v45 = int32(_a_F_get_json_returning_5)
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v45
								F_appendStringInfo(m, l1, int32(_a_F_get_json_returning_4), v9)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							case 3:
								v45 = int32(_a_F_get_json_returning_6)
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v45
								F_appendStringInfo(m, l1, int32(_a_F_get_json_returning_4), v9)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				} else {
					if v24 == int32(0) {
						m.G0 = v9 + int32(32)
						return
					} else {
						if v24 == int32(2) {
							v38 = int32(_a_F_get_json_returning_1)
						} else {
							v38 = int32(_a_F_get_json_returning_2)
						}
						F_appendStringInfoString(m, l1, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
							switch v42 {
							case 0:
								m.G0 = v9 + int32(32)
								return
							default:
								v45 = int32(_a_F_get_json_returning_3)
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v45
								F_appendStringInfo(m, l1, int32(_a_F_get_json_returning_4), v9)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							case 2:
								v45 = int32(_a_F_get_json_returning_5)
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v45
								F_appendStringInfo(m, l1, int32(_a_F_get_json_returning_4), v9)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							case 3:
								v45 = int32(_a_F_get_json_returning_6)
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v45
								F_appendStringInfo(m, l1, int32(_a_F_get_json_returning_4), v9)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
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
func F_json_agg_transfn_worker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = v9 + int32(12)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14 == v3 {
		v31 = int32(0)
		if v12 == v31 {
			v39 = v31
		} else {
			v34 = v31
			v35 = v3
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v34
			v39 = v35
		}
		v42 = v39
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		switch v17 - int32(429) {
		case 0:
			if v12 == int32(0) {
				v42 = int32(1)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)+168))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
				v34 = v24
				v35 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v34
				v39 = v35
				v42 = v39
			}
		case 1:
			if v12 == int32(0) {
				v42 = int32(2)
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+368))
				v34 = v29
				v35 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v34
				v39 = v35
				v42 = v39
			}
		default:
			v31 = int32(0)
			if v12 == v31 {
				v39 = v31
			} else {
				v34 = v31
				v35 = v3
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v34
				v39 = v35
			}
			v42 = v39
		}
	}
	if v42 != 0 {
		v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v43 == int32(1) {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v48 = F_get_fn_expr_argtype(m, v46, int32(1))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				if v48 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v143 = m.ExcPending
					if v143 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v146 = m.ExcPending
						if v146 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_json_agg_transfn_worker_0), int32(0))
							mBase = m.M
							v150 = m.ExcPending
							if v150 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_json_agg_transfn_worker_1), int32(799), int32(_a_F_json_agg_transfn_worker_2))
								mBase = m.M
								v155 = m.ExcPending
								if v155 != 0 {
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
					v54 = int32(_a_F_json_agg_transfn_worker_3)
					v55 = *(*int32)(unsafe.Add(mBase, _c_F_json_agg_transfn_worker[0]))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					*(*int32)(unsafe.Add(mBase, _c_F_json_agg_transfn_worker[0])) = v57
					v60 = F_palloc(m, int32(44))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v62 = F_makeStringInfo(m)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v60))) = v62
							*(*int32)(unsafe.Add(mBase, _c_F_json_agg_transfn_worker[0])) = v55
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
							F_appendStringInfoChar(m, v67, int32(91))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								F_json_categorize_type(m, v48, int32(0), v60+int32(12), v60+int32(16))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									v79 = v60
									if l1 != 0 {
										v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
										if v82 != 0 {
											m.G0 = v9 + int32(16)
											return v79
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
											v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
											if int32(2) <= v84 {
												F_appendStringInfoString(m, v83, int32(_a_F_json_agg_transfn_worker_4))
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
													if v90 == int32(1) {
														v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
														F_check_stack_depth(m)
														mBase = m.M
														v95 = m.ExcPending
														if v95 != 0 {
															return int32(0)
														} else {
															F_appendBinaryStringInfo(m, v93, int32(_a_F_json_agg_transfn_worker_5), int32(4))
															mBase = m.M
															v99 = m.ExcPending
															if v99 != 0 {
																return int32(0)
															} else {
																m.G0 = v9 + int32(16)
																return v79
															}
														}
													} else {
														v100 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
														v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														if v102 != 0 {
															v115 = v100
															v116 = int32(0)
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
															F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
															mBase = m.M
															v121 = m.ExcPending
															if v121 != 0 {
																return int32(0)
															} else {
																m.G0 = v9 + int32(16)
																return v79
															}
														} else {
															v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
															if v103 < int32(2) {
																v115 = v100
																v116 = int32(0)
																v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
																F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v9 + int32(16)
																	return v79
																}
															} else {
																v106 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																if v106&int32(-2) != int32(8) {
																	v115 = v100
																	v116 = int32(0)
																	v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
																	F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
																	mBase = m.M
																	v121 = m.ExcPending
																	if v121 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v9 + int32(16)
																		return v79
																	}
																} else {
																	F_appendStringInfoString(m, v100, int32(_a_F_json_agg_transfn_worker_6))
																	mBase = m.M
																	v113 = m.ExcPending
																	if v113 != 0 {
																		return int32(0)
																	} else {
																		v114 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
																		v115 = v114
																		v116 = int32(0)
																		v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																		v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
																		F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
																		mBase = m.M
																		v121 = m.ExcPending
																		if v121 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v9 + int32(16)
																			return v79
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
												if v90 == int32(1) {
													v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
													F_check_stack_depth(m)
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return int32(0)
													} else {
														F_appendBinaryStringInfo(m, v93, int32(_a_F_json_agg_transfn_worker_5), int32(4))
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return int32(0)
														} else {
															m.G0 = v9 + int32(16)
															return v79
														}
													}
												} else {
													v100 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
													v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													if v102 != 0 {
														v115 = v100
														v116 = int32(0)
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
														F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return int32(0)
														} else {
															m.G0 = v9 + int32(16)
															return v79
														}
													} else {
														v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
														if v103 < int32(2) {
															v115 = v100
															v116 = int32(0)
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
															F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
															mBase = m.M
															v121 = m.ExcPending
															if v121 != 0 {
																return int32(0)
															} else {
																m.G0 = v9 + int32(16)
																return v79
															}
														} else {
															v106 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
															if v106&int32(-2) != int32(8) {
																v115 = v100
																v116 = int32(0)
																v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
																F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v9 + int32(16)
																	return v79
																}
															} else {
																F_appendStringInfoString(m, v100, int32(_a_F_json_agg_transfn_worker_6))
																mBase = m.M
																v113 = m.ExcPending
																if v113 != 0 {
																	return int32(0)
																} else {
																	v114 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
																	v115 = v114
																	v116 = int32(0)
																	v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
																	F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
																	mBase = m.M
																	v121 = m.ExcPending
																	if v121 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v9 + int32(16)
																		return v79
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
										if int32(2) <= v84 {
											F_appendStringInfoString(m, v83, int32(_a_F_json_agg_transfn_worker_4))
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
												if v90 == int32(1) {
													v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
													F_check_stack_depth(m)
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return int32(0)
													} else {
														F_appendBinaryStringInfo(m, v93, int32(_a_F_json_agg_transfn_worker_5), int32(4))
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return int32(0)
														} else {
															m.G0 = v9 + int32(16)
															return v79
														}
													}
												} else {
													v100 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
													v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													if v102 != 0 {
														v115 = v100
														v116 = int32(0)
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
														F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return int32(0)
														} else {
															m.G0 = v9 + int32(16)
															return v79
														}
													} else {
														v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
														if v103 < int32(2) {
															v115 = v100
															v116 = int32(0)
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
															F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
															mBase = m.M
															v121 = m.ExcPending
															if v121 != 0 {
																return int32(0)
															} else {
																m.G0 = v9 + int32(16)
																return v79
															}
														} else {
															v106 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
															if v106&int32(-2) != int32(8) {
																v115 = v100
																v116 = int32(0)
																v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
																F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v9 + int32(16)
																	return v79
																}
															} else {
																F_appendStringInfoString(m, v100, int32(_a_F_json_agg_transfn_worker_6))
																mBase = m.M
																v113 = m.ExcPending
																if v113 != 0 {
																	return int32(0)
																} else {
																	v114 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
																	v115 = v114
																	v116 = int32(0)
																	v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
																	F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
																	mBase = m.M
																	v121 = m.ExcPending
																	if v121 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v9 + int32(16)
																		return v79
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
											if v90 == int32(1) {
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
												F_check_stack_depth(m)
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return int32(0)
												} else {
													F_appendBinaryStringInfo(m, v93, int32(_a_F_json_agg_transfn_worker_5), int32(4))
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return int32(0)
													} else {
														m.G0 = v9 + int32(16)
														return v79
													}
												}
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
												v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												if v102 != 0 {
													v115 = v100
													v116 = int32(0)
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
													F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return int32(0)
													} else {
														m.G0 = v9 + int32(16)
														return v79
													}
												} else {
													v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
													if v103 < int32(2) {
														v115 = v100
														v116 = int32(0)
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
														F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return int32(0)
														} else {
															m.G0 = v9 + int32(16)
															return v79
														}
													} else {
														v106 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
														if v106&int32(-2) != int32(8) {
															v115 = v100
															v116 = int32(0)
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
															F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
															mBase = m.M
															v121 = m.ExcPending
															if v121 != 0 {
																return int32(0)
															} else {
																m.G0 = v9 + int32(16)
																return v79
															}
														} else {
															F_appendStringInfoString(m, v100, int32(_a_F_json_agg_transfn_worker_6))
															mBase = m.M
															v113 = m.ExcPending
															if v113 != 0 {
																return int32(0)
															} else {
																v114 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
																v115 = v114
																v116 = int32(0)
																v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
																F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v9 + int32(16)
																	return v79
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
					}
				}
			}
		} else {
			v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v79 = v78
			if l1 != 0 {
				v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v82 != 0 {
					m.G0 = v9 + int32(16)
					return v79
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
					if int32(2) <= v84 {
						F_appendStringInfoString(m, v83, int32(_a_F_json_agg_transfn_worker_4))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
							if v90 == int32(1) {
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
								F_check_stack_depth(m)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									F_appendBinaryStringInfo(m, v93, int32(_a_F_json_agg_transfn_worker_5), int32(4))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(16)
										return v79
									}
								}
							} else {
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
								v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
								if v102 != 0 {
									v115 = v100
									v116 = int32(0)
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
									F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(16)
										return v79
									}
								} else {
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
									if v103 < int32(2) {
										v115 = v100
										v116 = int32(0)
										v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
										v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
										F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											m.G0 = v9 + int32(16)
											return v79
										}
									} else {
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
										if v106&int32(-2) != int32(8) {
											v115 = v100
											v116 = int32(0)
											v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
											v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
											F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												m.G0 = v9 + int32(16)
												return v79
											}
										} else {
											F_appendStringInfoString(m, v100, int32(_a_F_json_agg_transfn_worker_6))
											mBase = m.M
											v113 = m.ExcPending
											if v113 != 0 {
												return int32(0)
											} else {
												v114 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
												v115 = v114
												v116 = int32(0)
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
												F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return int32(0)
												} else {
													m.G0 = v9 + int32(16)
													return v79
												}
											}
										}
									}
								}
							}
						}
					} else {
						v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
						if v90 == int32(1) {
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
							F_check_stack_depth(m)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								F_appendBinaryStringInfo(m, v93, int32(_a_F_json_agg_transfn_worker_5), int32(4))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(16)
									return v79
								}
							}
						} else {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
							v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v102 != 0 {
								v115 = v100
								v116 = int32(0)
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
								F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(16)
									return v79
								}
							} else {
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
								if v103 < int32(2) {
									v115 = v100
									v116 = int32(0)
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
									F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(16)
										return v79
									}
								} else {
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
									if v106&int32(-2) != int32(8) {
										v115 = v100
										v116 = int32(0)
										v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
										v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
										F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											m.G0 = v9 + int32(16)
											return v79
										}
									} else {
										F_appendStringInfoString(m, v100, int32(_a_F_json_agg_transfn_worker_6))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
											v115 = v114
											v116 = int32(0)
											v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
											v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
											F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												m.G0 = v9 + int32(16)
												return v79
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
				if int32(2) <= v84 {
					F_appendStringInfoString(m, v83, int32(_a_F_json_agg_transfn_worker_4))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return int32(0)
					} else {
						v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
						if v90 == int32(1) {
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
							F_check_stack_depth(m)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								F_appendBinaryStringInfo(m, v93, int32(_a_F_json_agg_transfn_worker_5), int32(4))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(16)
									return v79
								}
							}
						} else {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
							v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v102 != 0 {
								v115 = v100
								v116 = int32(0)
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
								F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(16)
									return v79
								}
							} else {
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
								if v103 < int32(2) {
									v115 = v100
									v116 = int32(0)
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
									F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(16)
										return v79
									}
								} else {
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
									if v106&int32(-2) != int32(8) {
										v115 = v100
										v116 = int32(0)
										v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
										v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
										F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											m.G0 = v9 + int32(16)
											return v79
										}
									} else {
										F_appendStringInfoString(m, v100, int32(_a_F_json_agg_transfn_worker_6))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
											v115 = v114
											v116 = int32(0)
											v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
											v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
											F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												m.G0 = v9 + int32(16)
												return v79
											}
										}
									}
								}
							}
						}
					}
				} else {
					v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v90 == int32(1) {
						v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
						F_check_stack_depth(m)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							F_appendBinaryStringInfo(m, v93, int32(_a_F_json_agg_transfn_worker_5), int32(4))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v79
							}
						}
					} else {
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
						v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
						if v102 != 0 {
							v115 = v100
							v116 = int32(0)
							v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
							F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v79
							}
						} else {
							v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
							if v103 < int32(2) {
								v115 = v100
								v116 = int32(0)
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
								F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(16)
									return v79
								}
							} else {
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
								if v106&int32(-2) != int32(8) {
									v115 = v100
									v116 = int32(0)
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
									F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(16)
										return v79
									}
								} else {
									F_appendStringInfoString(m, v100, int32(_a_F_json_agg_transfn_worker_6))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										v114 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
										v115 = v114
										v116 = int32(0)
										v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
										v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
										F_datum_to_json_internal(m, v101, v116, v115, v117, v118, v116)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											m.G0 = v9 + int32(16)
											return v79
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
		v130 = m.ExcPending
		if v130 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_json_agg_transfn_worker_7), int32(0))
			mBase = m.M
			v134 = m.ExcPending
			if v134 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_json_agg_transfn_worker_1), int32(789), int32(_a_F_json_agg_transfn_worker_2))
				mBase = m.M
				v139 = m.ExcPending
				if v139 != 0 {
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
func F_json_extract_path_text(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_get_path_all(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_json_lex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int64
	_ = v61
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v547 int32
	_ = v547
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v578 int32
	_ = v578
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v687 int32
	_ = v687
	var v696 int32
	_ = v696
	var v711 int64
	_ = v711
	var v713 int64
	_ = v713
	var v719 int64
	_ = v719
	var v720 int64
	_ = v720
	var v738 int64
	_ = v738
	var v771 int64
	_ = v771
	var v774 int64
	_ = v774
	var v813 int64
	_ = v813
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v863 int32
	_ = v863
	var v878 int32
	_ = v878
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v923 int32
	_ = v923
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
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v998 int32
	_ = v998
	var v1004 int32
	_ = v1004
	var v1010 int32
	_ = v1010
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1043 int32
	_ = v1043
	var v1049 int32
	_ = v1049
	var v1067 int32
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1082 int32
	_ = v1082
	var v1088 int32
	_ = v1088
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1372 int32
	_ = v1372
	var v1375 int32
	_ = v1375
	var v1388 int32
	_ = v1388
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1411 int32
	_ = v1411
	var v1425 int32
	_ = v1425
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1456 int32
	_ = v1456
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1485 int32
	_ = v1485
	var v1490 int32
	_ = v1490
	var v1494 int32
	_ = v1494
	var v1498 int32
	_ = v1498
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1511 int32
	_ = v1511
	var v1518 int32
	_ = v1518
	var v1552 int32
	_ = v1552
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(80)
	m.G0 = v19
	v21 = int32(16)
	if l0 == int32(_a_F_json_lex_0) {
		v1552 = v21
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v19 + int32(80)
	return v1552
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v24 == int32(_a_F_json_lex_1) {
		v1552 = v21
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v29 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v1552 = int32(0)
	goto L1
L5:
	;
	v496 = v27 + v28
	if base.Ui32(v493) < base.Ui32(v496) {
		goto L93
	} else {
		goto L94
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v493 = v32
	v495 = v2
	goto L5
L7:
	;
	goto L8
L8:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)))
	if v33 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v57 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v55 = v36
	goto L9
L11:
	;
	goto L12
L12:
	;
	v38 = v24 + int32(4)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v40 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v40)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v40
	goto L13
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v49 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+2)) = uint8(v49)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v52 != int32(1) {
		v493 = v51
		v495 = v2
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v55 = v51
	goto L9
L15:
	;
	v493 = v55
	v495 = int32(1)
	goto L5
L16:
	;
	goto L17
L17:
	;
	v61 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v19)+64)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v61
	v76 = v56 + int32(4)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	switch v78 - int32(34) {
	case 0:
		goto L26
	default:
		goto L25
	case 11:
		goto L27
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v436 - v437
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v453 + v437
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v456
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v456
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v456
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v460
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v462
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v465 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+36)) = uint8(v465)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v464
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+68)) = uint8(v468)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v470
	v474 = F_json_lex(m, v19+int32(12))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L41
	} else {
		goto L87
	}
L19:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v436 = v434
	v437 = v173
	goto L18
L20:
	;
	v423 = int32(1)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424)+1)))
	if v425 != v423 {
		v1552 = v423
		goto L1
	} else {
		goto L86
	}
L21:
	;
	if v384 != 0 {
		v436 = v369
		v437 = v370
		goto L18
	} else {
		goto L85
	}
L22:
	;
	if v390 != 0 {
		goto L81
	} else {
		goto L82
	}
L23:
	;
	if base.Ui32(v269) < base.Ui32(v268) {
		goto L68
	} else {
		goto L69
	}
L24:
	;
	if v197 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L25:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(int32(9)) < base.Ui32((v78-int32(48))&int32(255)) {
		v268 = v189
		v269 = int32(0)
		goto L23
	} else {
		goto L48
	}
L26:
	;
	v82 = int32(0)
	v84 = v57 - int32(1)
	if v84 <= v82 {
		v116 = v82
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v197 = v81
	goto L24
L28:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v128 == int32(0) {
		goto L20
	} else {
		goto L34
	}
L29:
	;
	v89 = v84
	v91 = v82
	goto L30
L30:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89+v77))))
	if v104 != int32(92) {
		v116 = v91
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v116 = v84
	goto L28
L32:
	;
	v107 = int32(1)
	v110 = v91 + v107
	if v110 != v84 {
		v89 = v89 - v107
		v91 = v110
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v134 = int32(0)
	v136 = v116
	goto L35
L35:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v150 = int32(*(*int8)(unsafe.Add(mBase, uint32(v148+v134))))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v151 <= v152+int32(1) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L20
L37:
	;
	v172 = int32(1)
	v173 = v134 + v172
	if base.B2i32(v136&v172 == int32(0))&base.B2i32(v150 == int32(34)) != 0 {
		goto L19
	} else {
		goto L43
	}
L38:
	;
	F_appendStringInfoChar(m, v76, v150)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v160+v152))) = uint8(v150)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v165 = v163 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v165
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v169 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v167+v165))) = uint8(v169)
	goto L37
L41:
	;
	return int32(0)
L42:
	;
	goto L37
L43:
	;
	if v150 == int32(92) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v186 = v136 + int32(1)
	goto L46
L45:
	;
	v186 = int32(0)
	goto L46
L46:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v173) < base.Ui32(v187) {
		v134 = v173
		v136 = v186
		goto L35
	} else {
		goto L47
	}
L47:
	;
	goto L36
L48:
	;
	v197 = v189
	goto L24
L49:
	;
	v201 = int32(0)
	v387 = v201
	v390 = v201
	goto L22
L50:
	;
	goto L51
L51:
	;
	v203 = int32(0)
	v206 = v197
	v207 = v203
	v209 = v203
	goto L52
L52:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221+v209))))
	if base.Ui32(v223-int32(48)) < base.Ui32(int32(10)) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v268 = v265
	v269 = v262
	goto L23
L54:
	;
	v241 = base.I32_extend8_s(v223)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v242 <= v243+int32(1) {
		goto L62
	} else {
		goto L63
	}
L55:
	;
	v229 = v223 - int32(43)
	if int32(1)<<(uint(v229)%32)&int32(67108869) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v237 = base.B2i32(base.Ui32(v229) <= base.Ui32(int32(26)))
	goto L58
L57:
	;
	v237 = int32(0)
	goto L58
L58:
	;
	if v237 != 0 {
		goto L54
	} else {
		goto L59
	}
L59:
	;
	if v223 != int32(101) {
		v268 = v206
		v269 = v207
		goto L23
	} else {
		goto L60
	}
L60:
	;
	goto L54
L61:
	;
	v261 = int32(1)
	v262 = v207 + v261
	v264 = v209 + v261
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v264) < base.Ui32(v265) {
		v206 = v265
		v207 = v262
		v209 = v264
		goto L52
	} else {
		goto L66
	}
L62:
	;
	F_appendStringInfoChar(m, v76, v241)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L41
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v249+v243))) = uint8(v241)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v254 = v252 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v258 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v256+v254))) = uint8(v258)
	goto L61
L65:
	;
	goto L61
L66:
	;
	goto L53
L67:
	;
	if v369 != v370 {
		goto L21
	} else {
		goto L80
	}
L68:
	;
	v285 = v268
	v286 = v269
	goto L71
L69:
	;
	v352 = v268
	v353 = v269
	goto L70
L70:
	;
	v369 = v352
	v370 = v353
	v384 = int32(0)
	goto L67
L71:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301+v286))))
	v304 = base.I32_extend8_s(v303)
	v305 = int32(0)
	v311 = int32(255)
	if base.B2i32(v304 < v305)|base.B2i32(base.Ui32((v303&int32(223)-int32(65))&v311) < base.Ui32(int32(26)))|(base.B2i32(v304 == int32(95))|base.B2i32(base.Ui32(int32(246)) <= base.Ui32((v304-int32(58))&v311))) == v305 {
		v369 = v285
		v370 = v286
		v384 = int32(1)
		goto L67
	} else {
		goto L73
	}
L72:
	;
	v352 = v349
	v353 = v348
	goto L70
L73:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v328 <= v329+int32(1) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v348 = v286 + int32(1)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v348) < base.Ui32(v349) {
		v285 = v349
		v286 = v348
		goto L71
	} else {
		goto L79
	}
L75:
	;
	F_appendStringInfoChar(m, v76, v304)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L41
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v335+v329))) = uint8(v304)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v340 = v338 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v340
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v344 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v342+v340))) = uint8(v344)
	goto L74
L78:
	;
	goto L74
L79:
	;
	goto L72
L80:
	;
	v387 = v369
	v390 = v384
	goto L22
L81:
	;
	v436 = v387
	v437 = v387
	goto L18
L82:
	;
	goto L83
L83:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402)+1)))
	if v403&int32(1) != 0 {
		v436 = v387
		v437 = v387
		goto L18
	} else {
		goto L84
	}
L84:
	;
	v1552 = int32(1)
	goto L1
L85:
	;
	goto L20
L86:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v428
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v428 + v430
	v1552 = int32(15)
	goto L1
L87:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v476
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v478
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v482
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v484
	if v474 != 0 {
		v1552 = v474
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v486 != v484-v482 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v1552 = int32(15)
	goto L1
L90:
	;
	goto L91
L91:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v491 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v490)+2)) = uint8(v491)
	goto L4
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v503
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503))))
	switch v570 - int32(34) {
	case 0:
		goto L111
	default:
		goto L116
	case 10:
		goto L113
	case 11:
		goto L110
	case 14, 15, 16, 17, 18, 19, 20, 21, 22, 23:
		goto L109
	case 24:
		goto L112
	case 57:
		goto L115
	case 59:
		goto L114
	case 89:
		v1511 = int32(3)
		goto L106
	case 91:
		goto L107
	}
L93:
	;
	v499 = v493 + (v496 - v493)
	v503 = v493
	goto L96
L94:
	;
	v547 = v493
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v547
	v560 = int32(1)
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v561 != v560 {
		goto L4
	} else {
		goto L103
	}
L96:
	;
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503))))
	v518 = v516 - int32(9)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v518))|base.B2i32(int32(1)<<(uint(v518)%32)&int32(_a_F_json_lex_2) == int32(0)) != 0 {
		goto L92
	} else {
		goto L98
	}
L97:
	;
	v547 = v499
	goto L95
L98:
	;
	v529 = v503 + int32(1)
	if v516 == int32(10) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v529
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v533 + int32(1)
	goto L101
L100:
	;
	goto L101
L101:
	;
	if v529 != v499 {
		v503 = v529
		goto L96
	} else {
		goto L102
	}
L102:
	;
	goto L97
L103:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564)+1)))
	if v565 == int32(1) {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	v1552 = v560
	goto L1
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v1518
	goto L4
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v503 + int32(1)
	v1518 = v1511
	goto L105
L107:
	;
	v1511 = int32(4)
	goto L106
L108:
	;
	if v495 == int32(0) {
		goto L345
	} else {
		goto L346
	}
L109:
	;
	v1468 = int32(0)
	v1470 = F_json_lex_number(m, l0, v503, v1468, v1468)
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L41
	} else {
		goto L343
	}
L110:
	;
	v1463 = int32(0)
	v1465 = F_json_lex_number(m, l0, v503+int32(1), v1463, v1463)
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L41
	} else {
		goto L341
	}
L111:
	;
	v644 = m.G0
	v646 = v644 - int32(32)
	m.G0 = v646
	v648 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v650 = v648 + v649
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v651 == int32(1) {
		goto L126
	} else {
		goto L127
	}
L112:
	;
	v1511 = int32(8)
	goto L106
L113:
	;
	v1511 = int32(7)
	goto L106
L114:
	;
	v1511 = int32(6)
	goto L106
L115:
	;
	v1511 = int32(5)
	goto L106
L116:
	;
	if base.Ui32(v503) < base.Ui32(v496) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v578 = v503
	goto L121
L118:
	;
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v503 + int32(1)
	v1552 = int32(15)
	goto L1
L120:
	;
	if v503 != v617 {
		goto L108
	} else {
		goto L125
	}
L121:
	;
	v590 = int32(*(*int8)(unsafe.Add(mBase, uint32(v578))))
	v593 = int32(255)
	v609 = int32(0)
	if base.B2i32(base.B2i32(base.Ui32((v590-int32(48))&v593) < base.Ui32(int32(10)))|base.B2i32(base.Ui32((v590&int32(-33)-int32(65))&v593) < base.Ui32(int32(26)))|base.B2i32(v590 == int32(95)) == v609)&base.B2i32(v609 <= v590) != 0 {
		v617 = v578
		goto L120
	} else {
		goto L123
	}
L122:
	;
	v617 = v499
	goto L120
L123:
	;
	v615 = v578 + int32(1)
	if base.Ui32(v615) < base.Ui32(v496) {
		v578 = v615
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	goto L119
L126:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v654)))
	v656 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v655))) = uint8(v656)
	*(*int32)(unsafe.Add(mBase, uint32(v654)+12)) = v656
	*(*int32)(unsafe.Add(mBase, uint32(v654)+4)) = v656
	goto L129
L127:
	;
	goto L128
L128:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v664 = v662 + int32(1)
	if base.Ui32(v650) <= base.Ui32(v664) {
		v1411 = v664
		goto L131
	} else {
		goto L132
	}
L129:
	;
	goto L128
L130:
	;
	m.G0 = v646 + int32(32)
	if v1456 != 0 {
		v1552 = v1456
		goto L1
	} else {
		goto L340
	}
L131:
	;
	v1425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v1425 != int32(1) {
		goto L336
	} else {
		goto L337
	}
L132:
	;
	v667 = v650 - int32(8)
	v672 = v662
	v673 = v664
	v675 = int32(-1)
	goto L133
L133:
	;
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672)+1)))
	if v687 != int32(92) {
		goto L137
	} else {
		goto L138
	}
L134:
	;
	v1394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v1394 != int32(1) {
		goto L332
	} else {
		goto L333
	}
L135:
	;
	goto L134
L136:
	;
	v1388 = v1372 + int32(1)
	if base.Ui32(v1388) < base.Ui32(v650) {
		v672 = v1372
		v673 = v1388
		v675 = v1375
		goto L133
	} else {
		goto L331
	}
L137:
	;
	if v687 != int32(34) {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	goto L139
L139:
	;
	v939 = v672 + int32(2)
	if base.Ui32(v650) <= base.Ui32(v939) {
		goto L179
	} else {
		goto L180
	}
L140:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v930 = v650 - v673
	v931 = F_pg_encoding_mblen_or_incomplete(m, v929, v673, v930)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L41
	} else {
		goto L175
	}
L141:
	;
	if v675 != int32(-1) {
		goto L140
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	if v675 != int32(-1) {
		goto L172
	} else {
		goto L173
	}
L144:
	;
	if base.Ui32(v667) <= base.Ui32(v673) {
		v846 = v673
		goto L145
	} else {
		goto L146
	}
L145:
	;
	if base.Ui32(v650) <= base.Ui32(v846) {
		v892 = v846
		goto L159
	} else {
		goto L160
	}
L146:
	;
	v696 = v673
	goto L147
L147:
	;
	v711 = *(*int64)(unsafe.Add(mBase, uint32(v696)))
	v713 = v711 ^ int64(6655295901103053916)
	if int64(0) <= v711 {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	v846 = v843
	goto L145
L149:
	;
	v843 = v696 + int32(8)
	if base.Ui32(v843) < base.Ui32(v667) {
		v696 = v843
		goto L147
	} else {
		goto L158
	}
L150:
	;
	v719 = v711&int64(36170086419038336) ^ int64(-9187201950435737472)
	v720 = int64(72340172838076673)
	if v719&(v713-v720)|v719&(v711^int64(2459565876494606882)-v720) != int64(0) {
		v846 = v696
		goto L145
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v738 = int64(0)
	if base.B2i32(v713&int64(71776119061217280) == v738)|base.B2i32(v713&int64(280375465082880) == v738)|(base.B2i32(v713&int64(1095216660480) == v738)|base.B2i32(v713&int64(4278190080) == v738))|(base.B2i32(v713&int64(65280) == v738)|(base.B2i32(v713&int64(16711680) == v738)|base.B2i32(v713&int64(255) == v738))) != 0 {
		v846 = v696
		goto L145
	} else {
		goto L155
	}
L153:
	;
	if v719&(v711-int64(2314885530818453536)) == int64(0) {
		goto L149
	} else {
		goto L154
	}
L154:
	;
	v846 = v696
	goto L145
L155:
	;
	v771 = v711 ^ int64(2459565876494606882)
	v774 = int64(0)
	if base.B2i32(v771&int64(71776119061217280) == v774)|base.B2i32(v771&int64(280375465082880) == v774)|(base.B2i32(v771&int64(1095216660480) == v774)|base.B2i32(v771&int64(4278190080) == v774))|(base.B2i32(v771&int64(16711680) == v774)|base.B2i32(v771&int64(255) == v774)|(base.B2i32(v771&int64(65280) == v774)|base.B2i32(v711&int64(63050394783186944) == v774))) != 0 {
		v846 = v696
		goto L145
	} else {
		goto L156
	}
L156:
	;
	v813 = int64(0)
	if base.B2i32(v711&int64(246290604621824) == v813)|base.B2i32(v711&int64(962072674304) == v813)|(base.B2i32(v711&int64(3758096384) == v813)|base.B2i32(v711&int64(14680064) == v813))|(base.B2i32(v711&int64(224) == v813)|base.B2i32(v711&int64(57344) == v813)) != 0 {
		v846 = v696
		goto L145
	} else {
		goto L157
	}
L157:
	;
	goto L149
L158:
	;
	goto L148
L159:
	;
	v907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v907 == int32(1) {
		goto L168
	} else {
		goto L169
	}
L160:
	;
	v863 = v846
	goto L161
L161:
	;
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v863))))
	if base.B2i32(v878 == int32(34))|base.B2i32(v878 == int32(92)) != 0 {
		v892 = v863
		goto L159
	} else {
		goto L163
	}
L162:
	;
	v892 = v650
	goto L159
L163:
	;
	if base.Ui32(v878) <= base.Ui32(int32(31)) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v863
	v1456 = int32(5)
	goto L130
L165:
	;
	goto L166
L166:
	;
	v889 = v863 + int32(1)
	if v889 != v650 {
		v863 = v889
		goto L161
	} else {
		goto L167
	}
L167:
	;
	goto L162
L168:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendBinaryStringInfo(m, v910, v673, v892-v673)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L41
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v1372 = v892 - int32(1)
	v1375 = int32(-1)
	goto L136
L171:
	;
	goto L170
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v672 + int32(2)
	v1456 = int32(22)
	goto L130
L173:
	;
	goto L174
L174:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v923
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v672 + int32(2)
	v1456 = int32(0)
	goto L130
L175:
	;
	if v930 < v931 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v935 = v650
	goto L178
L177:
	;
	v935 = v673 + v931
	goto L178
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v935
	v1456 = int32(22)
	goto L130
L179:
	;
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v941 != int32(1) {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	goto L181
L181:
	;
	v956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939))))
	if v956 == int32(117) {
		goto L188
	} else {
		goto L189
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v939
	v1456 = int32(15)
	goto L130
L183:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v944)+1)))
	if v945 != 0 {
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v944+int32(4), v948, v650-v948)
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L41
	} else {
		goto L185
	}
L185:
	;
	v1456 = int32(1)
	goto L130
L186:
	;
	v1372 = v939
	v1375 = int32(-1)
	goto L136
L187:
	;
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1362 = v650 - v1357
	v1363 = F_pg_encoding_mblen_or_incomplete(m, v1361, v1357, v1362)
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L41
	} else {
		goto L327
	}
L188:
	;
	v959 = v650 - int32(2) - v672
	if v959 == int32(1) {
		goto L135
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v1187 = base.I32_extend8_s(v956)
	v1188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v1188 == int32(1) {
		goto L266
	} else {
		goto L267
	}
L191:
	;
	v963 = v672 + int32(3)
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v963))))
	v966 = v964 - int32(48)
	if base.Ui32(v966&int32(255)) < base.Ui32(int32(10)) {
		v987 = v966
		goto L192
	} else {
		goto L193
	}
L192:
	;
	if v959 == int32(2) {
		goto L135
	} else {
		goto L198
	}
L193:
	;
	if base.Ui32((v964-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v987 = v964 - int32(87)
	goto L192
L195:
	;
	goto L196
L196:
	;
	if base.Ui32(int32(5)) < base.Ui32((v964-int32(65))&int32(255)) {
		v1357 = v963
		goto L187
	} else {
		goto L197
	}
L197:
	;
	v987 = v964 - int32(55)
	goto L192
L198:
	;
	v990 = int32(255)
	v991 = v987 & v990
	v993 = v672 + int32(4)
	v994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v993))))
	v998 = (v994 - int32(48)) & v990
	if base.Ui32(int32(10)) <= base.Ui32(v998) {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	if v959 == int32(3) {
		goto L135
	} else {
		goto L207
	}
L200:
	;
	v1004 = (v994 - int32(97)) & int32(255)
	if base.Ui32(int32(6)) <= base.Ui32(v1004) {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	goto L202
L202:
	;
	v1028 = v991<<(uint(int32(4))%32) | v998
	goto L199
L203:
	;
	v1010 = (v994 - int32(65)) & int32(255)
	if base.Ui32(int32(5)) < base.Ui32(v1010) {
		v1357 = v993
		goto L187
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v1028 = v1004 + v991<<(uint(int32(4))%32) + int32(10)
	goto L199
L206:
	;
	v1028 = v1010 + v991<<(uint(int32(4))%32) + int32(10)
	goto L199
L207:
	;
	v1032 = v672 + int32(5)
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1032))))
	v1037 = (v1033 - int32(48)) & int32(255)
	if base.Ui32(int32(10)) <= base.Ui32(v1037) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	if v959 == int32(4) {
		goto L135
	} else {
		goto L216
	}
L209:
	;
	v1043 = (v1033 - int32(97)) & int32(255)
	if base.Ui32(int32(6)) <= base.Ui32(v1043) {
		goto L212
	} else {
		goto L213
	}
L210:
	;
	goto L211
L211:
	;
	v1067 = v1028<<(uint(int32(4))%32) | v1037
	goto L208
L212:
	;
	v1049 = (v1033 - int32(65)) & int32(255)
	if base.Ui32(int32(5)) < base.Ui32(v1049) {
		v1357 = v1032
		goto L187
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	v1067 = v1043 + v1028<<(uint(int32(4))%32) + int32(10)
	goto L208
L215:
	;
	v1067 = v1049 + v1028<<(uint(int32(4))%32) + int32(10)
	goto L208
L216:
	;
	v1071 = v672 + int32(6)
	v1072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1071))))
	v1076 = (v1072 - int32(48)) & int32(255)
	if base.Ui32(int32(10)) <= base.Ui32(v1076) {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	v1108 = v672 + int32(6)
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v1109 != int32(1) {
		v1372 = v1108
		v1375 = v675
		goto L136
	} else {
		goto L225
	}
L218:
	;
	v1082 = (v1072 - int32(97)) & int32(255)
	if base.Ui32(int32(6)) <= base.Ui32(v1082) {
		goto L221
	} else {
		goto L222
	}
L219:
	;
	goto L220
L220:
	;
	v1106 = v1067<<(uint(int32(4))%32) | v1076
	goto L217
L221:
	;
	v1088 = (v1072 - int32(65)) & int32(255)
	if base.Ui32(int32(5)) < base.Ui32(v1088) {
		v1357 = v1071
		goto L187
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v1106 = v1082 + v1067<<(uint(int32(4))%32) + int32(10)
	goto L217
L224:
	;
	v1106 = v1088 + v1067<<(uint(int32(4))%32) + int32(10)
	goto L217
L225:
	;
	v1113 = v1106 & int32(67107840)
	if v1113 != int32(_a_F_json_lex_3) {
		goto L229
	} else {
		goto L230
	}
L226:
	;
	v1170 = F_pg_unicode_to_server_noerror(m, v1169, v646)
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L41
	} else {
		goto L257
	}
L227:
	;
	v1169 = v675<<(uint(int32(10))%32)&int32(_a_F_json_lex_4) | v1106&int32(1023) + int32(_a_F_json_lex_5)
	goto L226
L228:
	;
	if v675 != int32(-1) {
		goto L245
	} else {
		goto L246
	}
L229:
	;
	if v1113 != int32(_a_F_json_lex_6) {
		goto L228
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	if v675 != int32(-1) {
		goto L227
	} else {
		goto L240
	}
L232:
	;
	if v675 == int32(-1) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1372 = v1108
	v1375 = v1106
	goto L136
L234:
	;
	goto L235
L235:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1121 = v650 - v1108
	v1122 = F_pg_encoding_mblen_or_incomplete(m, v1120, v1108, v1121)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L41
	} else {
		goto L236
	}
L236:
	;
	if v1121 < v1122 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v1126 = v650
	goto L239
L238:
	;
	v1126 = v1108 + v1122
	goto L239
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1126
	v1456 = int32(21)
	goto L130
L240:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1132 = v650 - v1108
	v1133 = F_pg_encoding_mblen_or_incomplete(m, v1131, v1108, v1132)
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L41
	} else {
		goto L241
	}
L241:
	;
	if v1132 < v1133 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1137 = v650
	goto L244
L243:
	;
	v1137 = v1108 + v1133
	goto L244
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1137
	v1456 = int32(22)
	goto L130
L245:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1143 = v650 - v1108
	v1144 = F_pg_encoding_mblen_or_incomplete(m, v1142, v1108, v1143)
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L41
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	if v1106 != 0 {
		v1169 = v1106
		goto L226
	} else {
		goto L252
	}
L248:
	;
	if v1143 < v1144 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1148 = v650
	goto L251
L250:
	;
	v1148 = v1108 + v1144
	goto L251
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1148
	v1456 = int32(22)
	goto L130
L252:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1152 = v650 - v1108
	v1153 = F_pg_encoding_mblen_or_incomplete(m, v1151, v1108, v1152)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L41
	} else {
		goto L253
	}
L253:
	;
	if v1152 < v1153 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1157 = v650
	goto L256
L255:
	;
	v1157 = v1108 + v1153
	goto L256
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1157
	v1456 = int32(17)
	goto L130
L257:
	;
	if v1170 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1175 = v650 - v1108
	v1176 = F_pg_encoding_mblen_or_incomplete(m, v1174, v1108, v1175)
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L41
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoString(m, v1183, v646)
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L41
	} else {
		goto L265
	}
L261:
	;
	if v1175 < v1176 {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v1180 = v650
	goto L264
L263:
	;
	v1180 = v1108 + v1176
	goto L264
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1180
	v1456 = int32(20)
	goto L130
L265:
	;
	v1372 = v1108
	v1375 = int32(-1)
	goto L136
L266:
	;
	if v675 != int32(-1) {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	goto L268
L268:
	;
	goto L299
L269:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1194 = v650 - v939
	v1195 = F_pg_encoding_mblen_or_incomplete(m, v1193, v939, v1194)
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L41
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	switch v956 - int32(47) {
	case 0, 45:
		goto L282
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 46, 47, 48, 49, 50, 52, 53, 54, 56, 57, 58, 59, 60, 61, 62, 64, 65, 66, 68:
		goto L276
	case 51:
		goto L281
	case 55:
		goto L280
	case 63:
		goto L279
	case 67:
		goto L278
	case 69:
		goto L277
	default:
		goto L283
	}
L272:
	;
	if v1194 < v1195 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1199 = v650
	goto L275
L274:
	;
	v1199 = v939 + v1195
	goto L275
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1199
	v1456 = int32(22)
	goto L130
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v939
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1231 = v650 - v939
	v1232 = F_pg_encoding_mblen_or_incomplete(m, v1230, v939, v1231)
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L41
	} else {
		goto L291
	}
L277:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoChar(m, v1225, int32(9))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L41
	} else {
		goto L290
	}
L278:
	;
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoChar(m, v1221, int32(13))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L41
	} else {
		goto L289
	}
L279:
	;
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoChar(m, v1217, int32(10))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L41
	} else {
		goto L288
	}
L280:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoChar(m, v1213, int32(12))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L41
	} else {
		goto L287
	}
L281:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoChar(m, v1209, int32(8))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L41
	} else {
		goto L286
	}
L282:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoChar(m, v1206, v1187)
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L41
	} else {
		goto L285
	}
L283:
	;
	if v956 != int32(34) {
		goto L276
	} else {
		goto L284
	}
L284:
	;
	goto L282
L285:
	;
	goto L186
L286:
	;
	goto L186
L287:
	;
	goto L186
L288:
	;
	goto L186
L289:
	;
	goto L186
L290:
	;
	goto L186
L291:
	;
	if v1231 < v1232 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1236 = v650
	goto L294
L293:
	;
	v1236 = v939 + v1232
	goto L294
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1236
	v1456 = int32(4)
	goto L130
L295:
	;
	if v1345 != 0 {
		goto L320
	} else {
		goto L321
	}
L296:
	;
	v1345 = int32(0)
	goto L295
L297:
	;
	v1323 = v1316
	v1325 = v1318
	goto L314
L298:
	;
	if base.B2i32(v1262 != v1263) == int32(0) {
		goto L296
	} else {
		goto L305
	}
L299:
	;
	v1254 = int32(_a_F_json_lex_7)
	v1256 = int32(9)
	goto L300
L300:
	;
	v1259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1254))))
	if v1259 == v1187&int32(255) {
		v1316 = v1254
		v1318 = v1256
		goto L297
	} else {
		goto L302
	}
L301:
	;
	goto L298
L302:
	;
	v1261 = int32(1)
	v1262 = v1256 - v1261
	v1263 = int32(0)
	v1266 = v1254 + v1261
	if v1266&int32(3) == v1263 {
		goto L298
	} else {
		goto L303
	}
L303:
	;
	if v1262 != 0 {
		v1254 = v1266
		v1256 = v1262
		goto L300
	} else {
		goto L304
	}
L304:
	;
	goto L301
L305:
	;
	v1279 = v1187 & int32(255)
	v1280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1266))))
	if base.B2i32(v1279 == v1280)|base.B2i32(base.Ui32(v1262) < base.Ui32(int32(4))) == int32(0) {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1289 = v1266
	v1291 = v1262
	goto L309
L307:
	;
	v1309 = v1266
	v1311 = v1262
	goto L308
L308:
	;
	if v1311 == int32(0) {
		goto L296
	} else {
		goto L313
	}
L309:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1289)))
	v1296 = v1295 ^ v1279*int32(16843009)
	v1299 = int32(-2139062144)
	if (int32(16843008)-v1296|v1296)&v1299 != v1299 {
		v1316 = v1289
		v1318 = v1291
		goto L297
	} else {
		goto L311
	}
L310:
	;
	v1309 = v1304
	v1311 = v1306
	goto L308
L311:
	;
	v1303 = int32(4)
	v1304 = v1289 + v1303
	v1306 = v1291 - v1303
	if base.Ui32(int32(3)) < base.Ui32(v1306) {
		v1289 = v1304
		v1291 = v1306
		goto L309
	} else {
		goto L312
	}
L312:
	;
	goto L310
L313:
	;
	v1316 = v1309
	v1318 = v1311
	goto L297
L314:
	;
	v1328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1323))))
	if v1187&int32(255) == v1328 {
		goto L316
	} else {
		goto L317
	}
L315:
	;
	goto L296
L316:
	;
	v1345 = v1323
	goto L295
L317:
	;
	goto L318
L318:
	;
	v1330 = int32(1)
	v1333 = v1325 - v1330
	if v1333 != 0 {
		v1323 = v1323 + v1330
		v1325 = v1333
		goto L314
	} else {
		goto L319
	}
L319:
	;
	goto L315
L320:
	;
	v1372 = v939
	v1375 = v675
	goto L136
L321:
	;
	goto L322
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v939
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1348 = v650 - v939
	v1349 = F_pg_encoding_mblen_or_incomplete(m, v1347, v939, v1348)
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L41
	} else {
		goto L323
	}
L323:
	;
	if v1348 < v1349 {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1353 = v650
	goto L326
L325:
	;
	v1353 = v939 + v1349
	goto L326
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1353
	v1456 = int32(4)
	goto L130
L327:
	;
	if v1362 < v1363 {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	v1367 = v650
	goto L330
L329:
	;
	v1367 = v1357 + v1363
	goto L330
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1367
	v1456 = int32(18)
	goto L130
L331:
	;
	v1411 = v1388
	goto L131
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v650
	v1456 = int32(15)
	goto L130
L333:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1397)+1)))
	if v1398 != 0 {
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v1397+int32(4), v1401, v650-v1401)
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L41
	} else {
		goto L335
	}
L335:
	;
	v1456 = int32(1)
	goto L130
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1411
	v1456 = int32(15)
	goto L130
L337:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1428)+1)))
	if v1429 != 0 {
		goto L336
	} else {
		goto L338
	}
L338:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v1428+int32(4), v1432, v650-v1432)
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L41
	} else {
		goto L339
	}
L339:
	;
	v1456 = int32(1)
	goto L130
L340:
	;
	v1518 = int32(1)
	goto L105
L341:
	;
	if v1465 != 0 {
		v1552 = v1465
		goto L1
	} else {
		goto L342
	}
L342:
	;
	v1518 = int32(2)
	goto L105
L343:
	;
	if v1470 != 0 {
		v1552 = v1470
		goto L1
	} else {
		goto L344
	}
L344:
	;
	v1518 = int32(2)
	goto L105
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v617
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v493
	v1490 = int32(15)
	switch v617 - v503 - int32(4) {
	case 0:
		goto L351
	case 1:
		goto L350
	default:
		v1552 = v1490
		goto L1
	}
L346:
	;
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1475)+1)))
	if v1476 != 0 {
		goto L345
	} else {
		goto L347
	}
L347:
	;
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v617 != v1477+v1478 {
		goto L345
	} else {
		goto L348
	}
L348:
	;
	F_appendBinaryStringInfo(m, v1475+int32(4), v503, v496-v503)
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L41
	} else {
		goto L349
	}
L349:
	;
	v1552 = int32(1)
	goto L1
L350:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	v1505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503)+4)))
	if v1502^int32(1936482662)|(v1505^int32(101)) != 0 {
		v1552 = v1490
		goto L1
	} else {
		goto L356
	}
L351:
	;
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	if v1494 == int32(1702195828) {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v1518 = int32(9)
	goto L105
L353:
	;
	goto L354
L354:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	if v1498 != int32(1819047278) {
		v1552 = v1490
		goto L1
	} else {
		goto L355
	}
L355:
	;
	v1518 = int32(11)
	goto L105
L356:
	;
	v1518 = int32(10)
	goto L105
}
func F_json_manifest_object_end(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v172 int32
	_ = v172
	var v177 int64
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
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
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v410 int64
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int64
	_ = v428
	var v429 int64
	_ = v429
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int64
	_ = v440
	var v443 int64
	_ = v443
	var v444 int64
	_ = v444
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v485 int32
	_ = v485
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	v13 = m.G0
	v15 = v13 - int32(272)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v18 - int32(1) {
	case 0:
		v468 = int32(14)
		goto L13
	default:
		goto L1
	case 6:
		goto L15
	case 10:
		goto L14
	}
L1:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v590)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_json_manifest_object_end_0)
	m.T0[v591].(func(*base.Module, int32, int32, int32))(m, v590, int32(_a_F_json_manifest_object_end_1), v15)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L21
	} else {
		goto L166
	}
L2:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+192)) = int32(_a_F_json_manifest_object_end_2)
	m.T0[v582].(func(*base.Module, int32, int32, int32))(m, v581, int32(_a_F_json_manifest_object_end_1), v15+int32(192))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L21
	} else {
		goto L165
	}
L3:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+224)) = int32(_a_F_json_manifest_object_end_3)
	m.T0[v573].(func(*base.Module, int32, int32, int32))(m, v572, int32(_a_F_json_manifest_object_end_1), v15+int32(224))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L21
	} else {
		goto L164
	}
L4:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = int32(_a_F_json_manifest_object_end_4)
	m.T0[v564].(func(*base.Module, int32, int32, int32))(m, v563, int32(_a_F_json_manifest_object_end_1), v15+int32(256))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L21
	} else {
		goto L163
	}
L5:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v396)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = int32(_a_F_json_manifest_object_end_5)
	m.T0[v555].(func(*base.Module, int32, int32, int32))(m, v396, int32(_a_F_json_manifest_object_end_1), v15+int32(176))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L21
	} else {
		goto L162
	}
L6:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v396)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = int32(_a_F_json_manifest_object_end_6)
	m.T0[v547].(func(*base.Module, int32, int32, int32))(m, v396, int32(_a_F_json_manifest_object_end_1), v15+int32(160))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L21
	} else {
		goto L161
	}
L7:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v396)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = int32(_a_F_json_manifest_object_end_7)
	m.T0[v539].(func(*base.Module, int32, int32, int32))(m, v396, int32(_a_F_json_manifest_object_end_1), v15+int32(144))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L21
	} else {
		goto L160
	}
L8:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v530)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = int32(_a_F_json_manifest_object_end_8)
	m.T0[v531].(func(*base.Module, int32, int32, int32))(m, v530, int32(_a_F_json_manifest_object_end_1), v15+int32(80))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L21
	} else {
		goto L159
	}
L9:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v521)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = int32(_a_F_json_manifest_object_end_9)
	m.T0[v522].(func(*base.Module, int32, int32, int32))(m, v521, int32(_a_F_json_manifest_object_end_1), v15+int32(96))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L21
	} else {
		goto L158
	}
L10:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = int32(_a_F_json_manifest_object_end_10)
	m.T0[v501].(func(*base.Module, int32, int32, int32))(m, v21, int32(_a_F_json_manifest_object_end_1), v15+int32(112))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L21
	} else {
		goto L157
	}
L11:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = int32(_a_F_json_manifest_object_end_11)
	m.T0[v493].(func(*base.Module, int32, int32, int32))(m, v21, int32(_a_F_json_manifest_object_end_1), v15+int32(32))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L21
	} else {
		goto L156
	}
L12:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = int32(_a_F_json_manifest_object_end_12)
	m.T0[v485].(func(*base.Module, int32, int32, int32))(m, v21, int32(_a_F_json_manifest_object_end_1), v15+int32(128))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L21
	} else {
		goto L155
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v468
	m.G0 = v15 + int32(272)
	return int32(0)
L14:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v397 == int32(0) {
		goto L7
	} else {
		goto L133
	}
L15:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v23 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v38 == int32(0) {
		goto L11
	} else {
		goto L24
	}
L17:
	;
	if v22 != 0 {
		v37 = v22
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v22 != 0 {
		goto L12
	} else {
		goto L23
	}
L20:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(_a_F_json_manifest_object_end_13)
	m.T0[v26].(func(*base.Module, int32, int32, int32))(m, v21, int32(_a_F_json_manifest_object_end_1), v15+int32(16))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v37 = int32(0)
	goto L16
L24:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v41 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v44 != 0 {
		goto L10
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if v37 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L27
L29:
	;
	v45 = F_strlen(m, v37)
	mBase = m.M
	v47 = base.I32_div_s(v45, int32(2))
	v50 = F_palloc(m, v47+int32(1))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L21
	} else {
		goto L32
	}
L30:
	;
	v172 = v38
	goto L31
L31:
	;
	v177 = F_strtox_2(m, v172, v15+int32(268), int32(10), int64(-1))
	mBase = m.M
	goto L58
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v50
	if v45&int32(1) != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	if int32(2) <= v45 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v60 = int32(0)
	goto L37
L35:
	;
	v150 = v50
	goto L36
L36:
	;
	v152 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v150+v47))) = uint8(v152)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_pfree(m, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L21
	} else {
		goto L57
	}
L37:
	;
	v73 = v57 + v60<<(uint(int32(1))%32)
	v74 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73))))
	v76 = v74 - int32(48)
	if base.Ui32(v76&int32(255)) <= base.Ui32(int32(9)) {
		v99 = v76
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v150 = v137
	goto L36
L39:
	;
	v100 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73)+1)))
	v102 = v100 - int32(48)
	if base.Ui32(v102&int32(255)) <= base.Ui32(int32(9)) {
		v125 = v102
		goto L47
	} else {
		goto L48
	}
L40:
	;
	if base.Ui32((v74-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v99 = v74 - int32(87)
	goto L39
L42:
	;
	goto L43
L43:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v74-int32(65))&int32(255)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v98 = int32(-1)
	goto L46
L45:
	;
	v98 = v74 - int32(55)
	goto L46
L46:
	;
	v99 = v98
	goto L39
L47:
	;
	if v99|v125 < int32(0) {
		goto L9
	} else {
		goto L55
	}
L48:
	;
	if base.Ui32((v100-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v125 = v100 - int32(87)
	goto L47
L50:
	;
	goto L51
L51:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v100-int32(65))&int32(255)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v124 = int32(-1)
	goto L54
L53:
	;
	v124 = v100 - int32(55)
	goto L54
L54:
	;
	v125 = v124
	goto L47
L55:
	;
	v132 = v125 + v99<<(uint(int32(4))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v60+v50))) = uint8(v132)
	v135 = v60 + int32(1)
	if v135 != v47 {
		v60 = v135
		goto L37
	} else {
		goto L56
	}
L56:
	;
	goto L38
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v172 = v159
	goto L31
L58:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v15)+268))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if v179 != 0 {
		goto L8
	} else {
		goto L59
	}
L59:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v180 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v241 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = int32(0)
	goto L60
L62:
	;
	goto L63
L63:
	;
	v186 = v15 + int32(264)
	v188 = F_pg_strcasecmp(m, v180, int32(_a_F_json_manifest_object_end_14))
	mBase = m.M
	if v188 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v231 != 0 {
		goto L60
	} else {
		goto L83
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = int32(0)
	v231 = int32(1)
	goto L64
L66:
	;
	goto L67
L67:
	;
	v195 = F_pg_strcasecmp(m, v180, int32(_a_F_json_manifest_object_end_15))
	mBase = m.M
	if v195 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v198 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v198
	v231 = v198
	goto L64
L69:
	;
	goto L70
L70:
	;
	v202 = F_pg_strcasecmp(m, v180, int32(_a_F_json_manifest_object_end_16))
	mBase = m.M
	if v202 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = int32(2)
	v231 = int32(1)
	goto L64
L72:
	;
	goto L73
L73:
	;
	v209 = F_pg_strcasecmp(m, v180, int32(_a_F_json_manifest_object_end_17))
	mBase = m.M
	if v209 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = int32(3)
	v231 = int32(1)
	goto L64
L75:
	;
	goto L76
L76:
	;
	v216 = F_pg_strcasecmp(m, v180, int32(_a_F_json_manifest_object_end_18))
	mBase = m.M
	if v216 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = int32(4)
	v231 = int32(1)
	goto L64
L78:
	;
	goto L79
L79:
	;
	v225 = F_pg_strcasecmp(m, v180, int32(_a_F_json_manifest_object_end_19))
	mBase = m.M
	if v225 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v226 = int32(0)
	goto L82
L81:
	;
	v226 = int32(5)
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v226
	v231 = base.B2i32(v225 == int32(0))
	goto L64
L83:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v233
	m.T0[v232].(func(*base.Module, int32, int32, int32))(m, v21, int32(_a_F_json_manifest_object_end_20), v15-int32(-64))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L21
	} else {
		goto L84
	}
L84:
	;
	goto L60
L85:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v15)+264))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	m.T0[v377].(func(*base.Module, int32, int32, int64, int32, int32, int32))(m, v21, v375, v177, v376, v370, v369)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L21
	} else {
		goto L120
	}
L86:
	;
	v244 = int32(0)
	v369 = v244
	v370 = v244
	goto L85
L87:
	;
	goto L88
L88:
	;
	v246 = F_strlen(m, v241)
	mBase = m.M
	if v246 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v249 = int32(0)
	v369 = v249
	v370 = v249
	goto L85
L90:
	;
	goto L91
L91:
	;
	v252 = base.I32_div_s(v246, int32(2))
	v253 = F_palloc(m, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L21
	} else {
		goto L92
	}
L92:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v246&int32(1) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if v246 < int32(2) {
		v369 = v253
		v370 = v252
		goto L85
	} else {
		goto L96
	}
L94:
	;
	v351 = v255
	goto L95
L95:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v351
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v355
	m.T0[v354].(func(*base.Module, int32, int32, int32))(m, v21, int32(_a_F_json_manifest_object_end_21), v15+int32(48))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L21
	} else {
		goto L119
	}
L96:
	;
	v264 = int32(0)
	goto L97
L97:
	;
	v277 = v255 + v264<<(uint(int32(1))%32)
	v278 = int32(*(*int8)(unsafe.Add(mBase, uint32(v277))))
	v280 = v278 - int32(48)
	if base.Ui32(v280&int32(255)) <= base.Ui32(int32(9)) {
		v303 = v280
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v351 = v341
	goto L95
L99:
	;
	v304 = int32(*(*int8)(unsafe.Add(mBase, uint32(v277)+1)))
	v306 = v304 - int32(48)
	if base.Ui32(v306&int32(255)) <= base.Ui32(int32(9)) {
		v329 = v306
		goto L107
	} else {
		goto L108
	}
L100:
	;
	if base.Ui32((v278-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v303 = v278 - int32(87)
	goto L99
L102:
	;
	goto L103
L103:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v278-int32(65))&int32(255)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v302 = int32(-1)
	goto L106
L105:
	;
	v302 = v278 - int32(55)
	goto L106
L106:
	;
	v303 = v302
	goto L99
L107:
	;
	if int32(0) <= v303|v329 {
		goto L115
	} else {
		goto L116
	}
L108:
	;
	if base.Ui32((v304-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v329 = v304 - int32(87)
	goto L107
L110:
	;
	goto L111
L111:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v304-int32(65))&int32(255)) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v328 = int32(-1)
	goto L114
L113:
	;
	v328 = v304 - int32(55)
	goto L114
L114:
	;
	v329 = v328
	goto L107
L115:
	;
	v336 = v329 + v303<<(uint(int32(4))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v264+v253))) = uint8(v336)
	v339 = v264 + int32(1)
	if v339 != v252 {
		v264 = v339
		goto L97
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	goto L98
L118:
	;
	v369 = v253
	v370 = v252
	goto L85
L119:
	;
	v369 = v253
	v370 = v252
	goto L85
L120:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v380 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	F_pfree(m, v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L21
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v385 != 0 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	goto L123
L125:
	;
	F_pfree(m, v385)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L21
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v390 != 0 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
	goto L127
L129:
	;
	F_pfree(m, v390)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L21
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v468 = int32(6)
	goto L13
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
	goto L131
L133:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v400 == int32(0) {
		goto L6
	} else {
		goto L134
	}
L134:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v403 == int32(0) {
		goto L5
	} else {
		goto L135
	}
L135:
	;
	v410 = F_strtox_2(m, v397, v15+int32(260), int32(10), int64(4294967295))
	mBase = m.M
	goto L136
L136:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v15)+260))
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412))))
	if v413 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v416 = v15 + int32(268)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v416
	v419 = v15 + int32(264)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v419
	v424 = F_sscanf(m, v414, int32(_a_F_json_manifest_object_end_22), v15+int32(240))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L21
	} else {
		goto L138
	}
L138:
	;
	if v424 != int32(2) {
		goto L3
	} else {
		goto L139
	}
L139:
	;
	v428 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v15)+264)))
	v429 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v15)+268)))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+208)) = v416
	*(*int32)(unsafe.Add(mBase, uint32(v15)+212)) = v419
	v436 = F_sscanf(m, v430, int32(_a_F_json_manifest_object_end_22), v15+int32(208))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L21
	} else {
		goto L140
	}
L140:
	;
	if v436 != int32(2) {
		goto L2
	} else {
		goto L141
	}
L141:
	;
	v440 = int64(32)
	v443 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v15)+264)))
	v444 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v15)+268)))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v396)+16))
	m.T0[v448].(func(*base.Module, int32, int32, int64, int64))(m, v396, base.I32_wrap_i64(v410), v429<<(uint(v440)%64)|v428, v443|v444<<(uint(v440)%64))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L21
	} else {
		goto L142
	}
L142:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v451 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	F_pfree(m, v451)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L21
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v456 != 0 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
	goto L145
L147:
	;
	F_pfree(m, v456)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L21
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v461 != 0 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
	goto L149
L151:
	;
	F_pfree(m, v461)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L21
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v468 = int32(10)
	goto L13
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	goto L153
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_json_object(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L8
	} else {
		goto L88
	}
L2:
	;
	m.G0 = v11 + int32(32)
	return v267
L3:
	;
	F_deconstruct_array_builtin(m, v14, int32(25), v11+int32(12), v11+int32(8), v11+int32(4))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L8
	} else {
		goto L25
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L21
	}
L5:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v43 == int32(2) {
		goto L3
	} else {
		goto L16
	}
L6:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)))
	if v22&int32(1) == int32(0) {
		goto L3
	} else {
		goto L11
	}
L7:
	;
	v20 = F_cstring_to_text(m, int32(_a_F_json_object_0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L10
	}
L8:
	;
	return int32(0)
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	switch v18 {
	case 0:
		goto L7
	case 1:
		goto L6
	case 2:
		goto L5
	default:
		goto L4
	}
L10:
	;
	v267 = v20
	goto L2
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	F_errmsg(m, int32(_a_F_json_object_1), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_json_object_2), int32(1428), int32(_a_F_json_object_3))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	F_errmsg(m, int32(_a_F_json_object_4), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_json_object_2), int32(1435), int32(_a_F_json_object_3))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	F_errmsg(m, int32(_a_F_json_object_5), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_json_object_2), int32(1441), int32(_a_F_json_object_3))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v89 = v11 + int32(16)
	F_initStringInfo(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	F_appendStringInfoChar(m, v89, int32(123))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	if int32(2) <= v87 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v98 = base.I32_div_s(v87, int32(2))
	v104 = int32(0)
	goto L31
L29:
	;
	goto L30
L30:
	;
	F_appendStringInfoChar(m, v11+int32(16), int32(125))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L8
	} else {
		goto L83
	}
L31:
	;
	v107 = int32(1)
	v108 = v104 << (uint(v107) % 32)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+v109))))
	if v111 == v107 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	goto L30
L33:
	;
	if v104 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_appendStringInfoString(m, v11+int32(16), int32(_a_F_json_object_6))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L8
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119+v108<<(uint(int32(2))%32))))
	v124 = F_pg_detoast_datum_packed(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L8
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v158 = int32(1)
	if v126&v158 != 0 {
		goto L50
	} else {
		goto L51
	}
L39:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v126 == int32(1) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+1)))
	if v132 == int32(18) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v143 = int32(1)
	if v126&v143 != 0 {
		v155 = int32(base.Ui32(v126)>>(uint(v143)%32)) - v143
		goto L38
	} else {
		goto L49
	}
L43:
	;
	v135 = int32(16)
	goto L45
L44:
	;
	v135 = int32(0)
	goto L45
L45:
	;
	if base.Ui32((v132-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v142 = int32(4)
	goto L48
L47:
	;
	v142 = v135
	goto L48
L48:
	;
	v155 = v142
	goto L38
L49:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v155 = int32(base.Ui32(v149)>>(uint(int32(2))%32)) - int32(4)
	goto L38
L50:
	;
	v162 = v158
	goto L52
L51:
	;
	v162 = int32(4)
	goto L52
L52:
	;
	F_escape_json_with_len(m, v11+int32(16), v124+v162, v155)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	if v124 != v123 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_pfree(m, v124)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L8
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v170 = v11 + int32(16)
	F_appendStringInfoString(m, v170, int32(_a_F_json_object_7))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L8
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v174 = int32(1)
	v175 = v108 | v174
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+v176))))
	if v178 == v174 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v239 = v104 + int32(1)
	if v239 != v98 {
		v104 = v239
		goto L31
	} else {
		goto L82
	}
L60:
	;
	F_appendStringInfoString(m, v170, int32(_a_F_json_object_8))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L8
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184+v175<<(uint(int32(2))%32))))
	v189 = F_pg_detoast_datum_packed(m, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L8
	} else {
		goto L65
	}
L63:
	;
	goto L59
L64:
	;
	v223 = int32(1)
	if v191&v223 != 0 {
		goto L76
	} else {
		goto L77
	}
L65:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	if v191 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+1)))
	if v197 == int32(18) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	v208 = int32(1)
	if v191&v208 != 0 {
		v220 = int32(base.Ui32(v191)>>(uint(v208)%32)) - v208
		goto L64
	} else {
		goto L75
	}
L69:
	;
	v200 = int32(16)
	goto L71
L70:
	;
	v200 = int32(0)
	goto L71
L71:
	;
	if base.Ui32((v197-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v207 = int32(4)
	goto L74
L73:
	;
	v207 = v200
	goto L74
L74:
	;
	v220 = v207
	goto L64
L75:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v220 = int32(base.Ui32(v214)>>(uint(int32(2))%32)) - int32(4)
	goto L64
L76:
	;
	v227 = v223
	goto L78
L77:
	;
	v227 = int32(4)
	goto L78
L78:
	;
	F_escape_json_with_len(m, v11+int32(16), v189+v227, v220)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	if v189 == v188 {
		goto L59
	} else {
		goto L80
	}
L80:
	;
	F_pfree(m, v189)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L8
	} else {
		goto L81
	}
L81:
	;
	goto L59
L82:
	;
	goto L32
L83:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_pfree(m, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L8
	} else {
		goto L84
	}
L84:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	F_pfree(m, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v262 = F_cstring_to_text_with_len(m, v260, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	F_pfree(m, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L8
	} else {
		goto L87
	}
L87:
	;
	v267 = v262
	goto L2
L88:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	F_errmsg(m, int32(_a_F_json_object_9), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_json_object_2), int32(1457), int32(_a_F_json_object_3))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L8
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_json_to_tsvector_byid(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = F_parse_jsonb_index_flags(m, v18)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v11
				v23 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v23
				v28 = v9 + int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v28
				F_iterate_json_values(m, v13, v20, v9+int32(24))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = F_make_tsvector(m, v28)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v36 != v13 {
							F_pfree(m, v13)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v40 != v18 {
									F_pfree(m, v18)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(32)
										return v34
									}
								} else {
									m.G0 = v9 + int32(32)
									return v34
								}
							}
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if v40 != v18 {
								F_pfree(m, v18)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(32)
									return v34
								}
							} else {
								m.G0 = v9 + int32(32)
								return v34
							}
						}
					}
				}
			}
		}
	}
}
func F_makeJsonBehavior(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_palloc0(m, int32(20))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(47)
		return v6
	}
}
