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
	if v8 == int32(420) {
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
				F_errfinish(m, int32(_a_F_GetJsonTableExecContext_1), int32(_a_F_GetJsonTableExecContext_4), int32(_a_F_GetJsonTableExecContext_3))
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
func F_add_json(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	v6 = l5
	v8 = l7
	v11 = m.G0
	v13 = v11 - int32(176)
	m.G0 = v13
	F_jspInit(m, v13+int32(112), l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = l4 + int32(4)
	v25 = F_JsonbExtractScalar(m, v22, v13+int32(80))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v25 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+92)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = int32(18)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v32 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+144)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+140)) = l1
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+156)) = int64(0)
	v71 = int32(base.Ui32(v67) >> (uint(int32(31)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+173)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+172)) = uint8(v71)
	v75 = v13 + int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+152)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v13)+148)) = v75
	v78 = m.T0[l3].(func(*base.Module, int32) int32)(m, l1)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L18
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v61
	goto L6
L8:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
	if v38 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v49 = int32(1)
	if v32&v49 != 0 {
		v61 = int32(base.Ui32(v32)>>(uint(v49)%32)) - v49
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v41 = int32(16)
	goto L13
L12:
	;
	v41 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v38-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v48 = int32(4)
	goto L16
L15:
	;
	v48 = v41
	goto L16
L16:
	;
	v61 = v48
	goto L7
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v61 = int32(base.Ui32(v55)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+175)) = uint8(v8)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+174)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+168)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+164)) = v78 + int32(1)
	v87 = int32(0)
	if l6|base.B2i32(v67 < v87) == v87 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	m.G0 = v13 + int32(176)
	return v153
L20:
	;
	v92 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v92
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = int64(8589934592)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v13
	v102 = F_executeItemOptUnwrapTarget(m, v13+int32(140), v13+int32(112), v75, v13, v92)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v141 = F_executeItemOptUnwrapTarget(m, v13+int32(140), v13+int32(112), v13+int32(80), l6, v71)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L34
	}
L23:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v105 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v106 = v105
	goto L27
L25:
	;
	goto L26
L26:
	;
	v129 = int32(2)
	if v102 == v129 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	F_pfree(m, v106)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	if v116 != 0 {
		v106 = v116
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v134 = v129
	goto L33
L32:
	;
	v134 = base.B2i32(v104 == int32(0))
	goto L33
L33:
	;
	v153 = v134
	goto L19
L34:
	;
	v153 = v141
	goto L19
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
func F_json_agg_transfn_worker(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
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
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = v10 + int32(12)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 == v3 {
		v32 = int32(0)
		if v13 == v32 {
			v40 = v32
		} else {
			v35 = v32
			v36 = v3
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = v35
			v40 = v36
		}
		v43 = v40
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		switch v18 - int32(435) {
		case 0:
			if v13 == int32(0) {
				v43 = int32(1)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+168))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
				v35 = v25
				v36 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v35
				v40 = v36
				v43 = v40
			}
		case 1:
			if v13 == int32(0) {
				v43 = int32(2)
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+376))
				v35 = v30
				v36 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v35
				v40 = v36
				v43 = v40
			}
		default:
			v32 = int32(0)
			if v13 == v32 {
				v40 = v32
			} else {
				v35 = v32
				v36 = v3
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v35
				v40 = v36
			}
			v43 = v40
		}
	}
	if v43 != 0 {
		v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v44 == int32(1) {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v49 = F_get_fn_expr_argtype(m, v47, int32(1))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int64(0)
			} else {
				if v49 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v145 = m.ExcPending
					if v145 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v148 = m.ExcPending
						if v148 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_json_agg_transfn_worker_0), int32(0))
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_json_agg_transfn_worker_1), int32(769), int32(_a_F_json_agg_transfn_worker_2))
								mBase = m.M
								v157 = m.ExcPending
								if v157 != 0 {
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
					v55 = int32(_a_F_json_agg_transfn_worker_3)
					v56 = *(*int32)(unsafe.Add(mBase, _c_F_json_agg_transfn_worker[0]))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					*(*int32)(unsafe.Add(mBase, _c_F_json_agg_transfn_worker[0])) = v58
					v61 = F_palloc(m, int32(44))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int64(0)
					} else {
						v63 = F_makeStringInfo(m)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v61))) = v63
							*(*int32)(unsafe.Add(mBase, _c_F_json_agg_transfn_worker[0])) = v56
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
							F_appendStringInfoChar(m, v68, int32(91))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int64(0)
							} else {
								F_json_categorize_type(m, v49, int32(0), v61+int32(12), v61+int32(16))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int64(0)
								} else {
									v80 = v61
									if l1 != 0 {
										v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
										if v83 != 0 {
											m.G0 = v10 + int32(16)
											return base.I64_extend_i32_u(v80)
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
											v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
											if int32(2) <= v85 {
												F_appendStringInfoString(m, v84, int32(_a_F_json_agg_transfn_worker_4))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return int64(0)
												} else {
													v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
													if v91 == int32(1) {
														v94 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
														F_check_stack_depth(m)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return int64(0)
														} else {
															F_appendBinaryStringInfo(m, v94, int32(_a_F_json_agg_transfn_worker_5), int32(4))
															mBase = m.M
															v100 = m.ExcPending
															if v100 != 0 {
																return int64(0)
															} else {
																m.G0 = v10 + int32(16)
																return base.I64_extend_i32_u(v80)
															}
														}
													} else {
														v101 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
														v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
														v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
														if v103 != 0 {
															v116 = v101
															v117 = int32(0)
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
															v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
															F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
															mBase = m.M
															v122 = m.ExcPending
															if v122 != 0 {
																return int64(0)
															} else {
																m.G0 = v10 + int32(16)
																return base.I64_extend_i32_u(v80)
															}
														} else {
															v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
															if v104 < int32(2) {
																v116 = v101
																v117 = int32(0)
																v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
																v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
																F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
																mBase = m.M
																v122 = m.ExcPending
																if v122 != 0 {
																	return int64(0)
																} else {
																	m.G0 = v10 + int32(16)
																	return base.I64_extend_i32_u(v80)
																}
															} else {
																v107 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
																if v107&int32(-2) != int32(8) {
																	v116 = v101
																	v117 = int32(0)
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
																	v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
																	F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
																	mBase = m.M
																	v122 = m.ExcPending
																	if v122 != 0 {
																		return int64(0)
																	} else {
																		m.G0 = v10 + int32(16)
																		return base.I64_extend_i32_u(v80)
																	}
																} else {
																	F_appendStringInfoString(m, v101, int32(_a_F_json_agg_transfn_worker_6))
																	mBase = m.M
																	v114 = m.ExcPending
																	if v114 != 0 {
																		return int64(0)
																	} else {
																		v115 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
																		v116 = v115
																		v117 = int32(0)
																		v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
																		v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
																		F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
																		mBase = m.M
																		v122 = m.ExcPending
																		if v122 != 0 {
																			return int64(0)
																		} else {
																			m.G0 = v10 + int32(16)
																			return base.I64_extend_i32_u(v80)
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
												if v91 == int32(1) {
													v94 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
													F_check_stack_depth(m)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return int64(0)
													} else {
														F_appendBinaryStringInfo(m, v94, int32(_a_F_json_agg_transfn_worker_5), int32(4))
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return int64(0)
														} else {
															m.G0 = v10 + int32(16)
															return base.I64_extend_i32_u(v80)
														}
													}
												} else {
													v101 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
													v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
													v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
													if v103 != 0 {
														v116 = v101
														v117 = int32(0)
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
														v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
														F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return int64(0)
														} else {
															m.G0 = v10 + int32(16)
															return base.I64_extend_i32_u(v80)
														}
													} else {
														v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
														if v104 < int32(2) {
															v116 = v101
															v117 = int32(0)
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
															v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
															F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
															mBase = m.M
															v122 = m.ExcPending
															if v122 != 0 {
																return int64(0)
															} else {
																m.G0 = v10 + int32(16)
																return base.I64_extend_i32_u(v80)
															}
														} else {
															v107 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
															if v107&int32(-2) != int32(8) {
																v116 = v101
																v117 = int32(0)
																v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
																v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
																F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
																mBase = m.M
																v122 = m.ExcPending
																if v122 != 0 {
																	return int64(0)
																} else {
																	m.G0 = v10 + int32(16)
																	return base.I64_extend_i32_u(v80)
																}
															} else {
																F_appendStringInfoString(m, v101, int32(_a_F_json_agg_transfn_worker_6))
																mBase = m.M
																v114 = m.ExcPending
																if v114 != 0 {
																	return int64(0)
																} else {
																	v115 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
																	v116 = v115
																	v117 = int32(0)
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
																	v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
																	F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
																	mBase = m.M
																	v122 = m.ExcPending
																	if v122 != 0 {
																		return int64(0)
																	} else {
																		m.G0 = v10 + int32(16)
																		return base.I64_extend_i32_u(v80)
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
										v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
										if int32(2) <= v85 {
											F_appendStringInfoString(m, v84, int32(_a_F_json_agg_transfn_worker_4))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int64(0)
											} else {
												v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
												if v91 == int32(1) {
													v94 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
													F_check_stack_depth(m)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return int64(0)
													} else {
														F_appendBinaryStringInfo(m, v94, int32(_a_F_json_agg_transfn_worker_5), int32(4))
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return int64(0)
														} else {
															m.G0 = v10 + int32(16)
															return base.I64_extend_i32_u(v80)
														}
													}
												} else {
													v101 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
													v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
													v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
													if v103 != 0 {
														v116 = v101
														v117 = int32(0)
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
														v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
														F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return int64(0)
														} else {
															m.G0 = v10 + int32(16)
															return base.I64_extend_i32_u(v80)
														}
													} else {
														v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
														if v104 < int32(2) {
															v116 = v101
															v117 = int32(0)
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
															v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
															F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
															mBase = m.M
															v122 = m.ExcPending
															if v122 != 0 {
																return int64(0)
															} else {
																m.G0 = v10 + int32(16)
																return base.I64_extend_i32_u(v80)
															}
														} else {
															v107 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
															if v107&int32(-2) != int32(8) {
																v116 = v101
																v117 = int32(0)
																v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
																v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
																F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
																mBase = m.M
																v122 = m.ExcPending
																if v122 != 0 {
																	return int64(0)
																} else {
																	m.G0 = v10 + int32(16)
																	return base.I64_extend_i32_u(v80)
																}
															} else {
																F_appendStringInfoString(m, v101, int32(_a_F_json_agg_transfn_worker_6))
																mBase = m.M
																v114 = m.ExcPending
																if v114 != 0 {
																	return int64(0)
																} else {
																	v115 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
																	v116 = v115
																	v117 = int32(0)
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
																	v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
																	F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
																	mBase = m.M
																	v122 = m.ExcPending
																	if v122 != 0 {
																		return int64(0)
																	} else {
																		m.G0 = v10 + int32(16)
																		return base.I64_extend_i32_u(v80)
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
											if v91 == int32(1) {
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
												F_check_stack_depth(m)
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return int64(0)
												} else {
													F_appendBinaryStringInfo(m, v94, int32(_a_F_json_agg_transfn_worker_5), int32(4))
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return int64(0)
													} else {
														m.G0 = v10 + int32(16)
														return base.I64_extend_i32_u(v80)
													}
												}
											} else {
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
												v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
												v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
												if v103 != 0 {
													v116 = v101
													v117 = int32(0)
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
													F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
													mBase = m.M
													v122 = m.ExcPending
													if v122 != 0 {
														return int64(0)
													} else {
														m.G0 = v10 + int32(16)
														return base.I64_extend_i32_u(v80)
													}
												} else {
													v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
													if v104 < int32(2) {
														v116 = v101
														v117 = int32(0)
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
														v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
														F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return int64(0)
														} else {
															m.G0 = v10 + int32(16)
															return base.I64_extend_i32_u(v80)
														}
													} else {
														v107 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
														if v107&int32(-2) != int32(8) {
															v116 = v101
															v117 = int32(0)
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
															v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
															F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
															mBase = m.M
															v122 = m.ExcPending
															if v122 != 0 {
																return int64(0)
															} else {
																m.G0 = v10 + int32(16)
																return base.I64_extend_i32_u(v80)
															}
														} else {
															F_appendStringInfoString(m, v101, int32(_a_F_json_agg_transfn_worker_6))
															mBase = m.M
															v114 = m.ExcPending
															if v114 != 0 {
																return int64(0)
															} else {
																v115 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
																v116 = v115
																v117 = int32(0)
																v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
																v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
																F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
																mBase = m.M
																v122 = m.ExcPending
																if v122 != 0 {
																	return int64(0)
																} else {
																	m.G0 = v10 + int32(16)
																	return base.I64_extend_i32_u(v80)
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
			v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v80 = v79
			if l1 != 0 {
				v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
				if v83 != 0 {
					m.G0 = v10 + int32(16)
					return base.I64_extend_i32_u(v80)
				} else {
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
					if int32(2) <= v85 {
						F_appendStringInfoString(m, v84, int32(_a_F_json_agg_transfn_worker_4))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int64(0)
						} else {
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
							if v91 == int32(1) {
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
								F_check_stack_depth(m)
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int64(0)
								} else {
									F_appendBinaryStringInfo(m, v94, int32(_a_F_json_agg_transfn_worker_5), int32(4))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return int64(0)
									} else {
										m.G0 = v10 + int32(16)
										return base.I64_extend_i32_u(v80)
									}
								}
							} else {
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
								v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
								v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
								if v103 != 0 {
									v116 = v101
									v117 = int32(0)
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
									F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int64(0)
									} else {
										m.G0 = v10 + int32(16)
										return base.I64_extend_i32_u(v80)
									}
								} else {
									v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
									if v104 < int32(2) {
										v116 = v101
										v117 = int32(0)
										v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
										v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
										F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int64(0)
										} else {
											m.G0 = v10 + int32(16)
											return base.I64_extend_i32_u(v80)
										}
									} else {
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
										if v107&int32(-2) != int32(8) {
											v116 = v101
											v117 = int32(0)
											v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
											F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return int64(0)
											} else {
												m.G0 = v10 + int32(16)
												return base.I64_extend_i32_u(v80)
											}
										} else {
											F_appendStringInfoString(m, v101, int32(_a_F_json_agg_transfn_worker_6))
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return int64(0)
											} else {
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
												v116 = v115
												v117 = int32(0)
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
												F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return int64(0)
												} else {
													m.G0 = v10 + int32(16)
													return base.I64_extend_i32_u(v80)
												}
											}
										}
									}
								}
							}
						}
					} else {
						v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
						if v91 == int32(1) {
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
							F_check_stack_depth(m)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int64(0)
							} else {
								F_appendBinaryStringInfo(m, v94, int32(_a_F_json_agg_transfn_worker_5), int32(4))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int64(0)
								} else {
									m.G0 = v10 + int32(16)
									return base.I64_extend_i32_u(v80)
								}
							}
						} else {
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
							v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
							v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
							if v103 != 0 {
								v116 = v101
								v117 = int32(0)
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
								F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int64(0)
								} else {
									m.G0 = v10 + int32(16)
									return base.I64_extend_i32_u(v80)
								}
							} else {
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
								if v104 < int32(2) {
									v116 = v101
									v117 = int32(0)
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
									F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int64(0)
									} else {
										m.G0 = v10 + int32(16)
										return base.I64_extend_i32_u(v80)
									}
								} else {
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
									if v107&int32(-2) != int32(8) {
										v116 = v101
										v117 = int32(0)
										v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
										v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
										F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int64(0)
										} else {
											m.G0 = v10 + int32(16)
											return base.I64_extend_i32_u(v80)
										}
									} else {
										F_appendStringInfoString(m, v101, int32(_a_F_json_agg_transfn_worker_6))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int64(0)
										} else {
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
											v116 = v115
											v117 = int32(0)
											v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
											F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return int64(0)
											} else {
												m.G0 = v10 + int32(16)
												return base.I64_extend_i32_u(v80)
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
				if int32(2) <= v85 {
					F_appendStringInfoString(m, v84, int32(_a_F_json_agg_transfn_worker_4))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int64(0)
					} else {
						v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
						if v91 == int32(1) {
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
							F_check_stack_depth(m)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int64(0)
							} else {
								F_appendBinaryStringInfo(m, v94, int32(_a_F_json_agg_transfn_worker_5), int32(4))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int64(0)
								} else {
									m.G0 = v10 + int32(16)
									return base.I64_extend_i32_u(v80)
								}
							}
						} else {
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
							v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
							v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
							if v103 != 0 {
								v116 = v101
								v117 = int32(0)
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
								F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int64(0)
								} else {
									m.G0 = v10 + int32(16)
									return base.I64_extend_i32_u(v80)
								}
							} else {
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
								if v104 < int32(2) {
									v116 = v101
									v117 = int32(0)
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
									F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int64(0)
									} else {
										m.G0 = v10 + int32(16)
										return base.I64_extend_i32_u(v80)
									}
								} else {
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
									if v107&int32(-2) != int32(8) {
										v116 = v101
										v117 = int32(0)
										v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
										v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
										F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int64(0)
										} else {
											m.G0 = v10 + int32(16)
											return base.I64_extend_i32_u(v80)
										}
									} else {
										F_appendStringInfoString(m, v101, int32(_a_F_json_agg_transfn_worker_6))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int64(0)
										} else {
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
											v116 = v115
											v117 = int32(0)
											v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
											F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return int64(0)
											} else {
												m.G0 = v10 + int32(16)
												return base.I64_extend_i32_u(v80)
											}
										}
									}
								}
							}
						}
					}
				} else {
					v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
					if v91 == int32(1) {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
						F_check_stack_depth(m)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int64(0)
						} else {
							F_appendBinaryStringInfo(m, v94, int32(_a_F_json_agg_transfn_worker_5), int32(4))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int64(0)
							} else {
								m.G0 = v10 + int32(16)
								return base.I64_extend_i32_u(v80)
							}
						}
					} else {
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
						v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
						v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
						if v103 != 0 {
							v116 = v101
							v117 = int32(0)
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
							v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
							F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
							mBase = m.M
							v122 = m.ExcPending
							if v122 != 0 {
								return int64(0)
							} else {
								m.G0 = v10 + int32(16)
								return base.I64_extend_i32_u(v80)
							}
						} else {
							v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
							if v104 < int32(2) {
								v116 = v101
								v117 = int32(0)
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
								F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int64(0)
								} else {
									m.G0 = v10 + int32(16)
									return base.I64_extend_i32_u(v80)
								}
							} else {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
								if v107&int32(-2) != int32(8) {
									v116 = v101
									v117 = int32(0)
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
									F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int64(0)
									} else {
										m.G0 = v10 + int32(16)
										return base.I64_extend_i32_u(v80)
									}
								} else {
									F_appendStringInfoString(m, v101, int32(_a_F_json_agg_transfn_worker_6))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int64(0)
									} else {
										v115 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
										v116 = v115
										v117 = int32(0)
										v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
										v119 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
										F_datum_to_json_internal(m, v102, v117, v116, v118, v119, v117)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int64(0)
										} else {
											m.G0 = v10 + int32(16)
											return base.I64_extend_i32_u(v80)
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
		v132 = m.ExcPending
		if v132 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_json_agg_transfn_worker_7), int32(0))
			mBase = m.M
			v136 = m.ExcPending
			if v136 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_json_agg_transfn_worker_1), int32(759), int32(_a_F_json_agg_transfn_worker_2))
				mBase = m.M
				v141 = m.ExcPending
				if v141 != 0 {
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
func F_json_extract_path_text(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_get_path_all(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_json_lex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int64
	_ = v64
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
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
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v777 int32
	_ = v777
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v815 int32
	_ = v815
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v849 int32
	_ = v849
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v933 int32
	_ = v933
	var v943 int32
	_ = v943
	var v960 int64
	_ = v960
	var v962 int64
	_ = v962
	var v968 int64
	_ = v968
	var v969 int64
	_ = v969
	var v987 int64
	_ = v987
	var v1020 int64
	_ = v1020
	var v1023 int64
	_ = v1023
	var v1062 int64
	_ = v1062
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1115 int32
	_ = v1115
	var v1132 int32
	_ = v1132
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1180 int32
	_ = v1180
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1255 int32
	_ = v1255
	var v1261 int32
	_ = v1261
	var v1267 int32
	_ = v1267
	var v1284 int32
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1293 int32
	_ = v1293
	var v1299 int32
	_ = v1299
	var v1305 int32
	_ = v1305
	var v1322 int32
	_ = v1322
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1337 int32
	_ = v1337
	var v1343 int32
	_ = v1343
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1367 int32
	_ = v1367
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1411 int32
	_ = v1411
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1460 int32
	_ = v1460
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1490 int32
	_ = v1490
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1520 int32
	_ = v1520
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1563 int32
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1570 int32
	_ = v1570
	var v1572 int32
	_ = v1572
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1607 int32
	_ = v1607
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1644 int32
	_ = v1644
	var v1650 int32
	_ = v1650
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1657 int32
	_ = v1657
	var v1660 int32
	_ = v1660
	var v1666 int32
	_ = v1666
	var v1684 int32
	_ = v1684
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1718 int32
	_ = v1718
	var v1725 int32
	_ = v1725
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1747 int32
	_ = v1747
	var v1752 int32
	_ = v1752
	var v1756 int32
	_ = v1756
	var v1760 int32
	_ = v1760
	var v1764 int32
	_ = v1764
	var v1767 int32
	_ = v1767
	var v1773 int32
	_ = v1773
	var v1779 int32
	_ = v1779
	var v1819 int32
	_ = v1819
	v2 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(80)
	m.G0 = v22
	v24 = int32(16)
	if l0 == int32(_a_F_json_lex_0) {
		v1819 = v24
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v22 + int32(80)
	return v1819
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v27 == int32(_a_F_json_lex_1) {
		v1819 = v24
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v32 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v1819 = int32(0)
	goto L1
L5:
	;
	v727 = v31 + v30
	if base.Ui32(v723) < base.Ui32(v727) {
		goto L128
	} else {
		goto L129
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v723 = v35
	v725 = v2
	goto L5
L7:
	;
	goto L8
L8:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+2)))
	if v36 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v60 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v58 = v39
	goto L9
L11:
	;
	goto L12
L12:
	;
	v41 = v27 + int32(4)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v43 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v43)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = v43
	goto L13
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v52 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+2)) = uint8(v52)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v55 != int32(1) {
		v723 = v54
		v725 = v2
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v58 = v54
	goto L9
L15:
	;
	v723 = v58
	v725 = int32(1)
	goto L5
L16:
	;
	goto L17
L17:
	;
	v64 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+72)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v22)+64)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v64
	v79 = v59 + int32(4)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	switch v81 - int32(34) {
	case 0:
		goto L27
	default:
		goto L26
	case 11:
		goto L25
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v665 - v663
	v683 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v683 + v663
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v686
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v690
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v692
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v695 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+36)) = uint8(v695)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v694
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+68)) = uint8(v698)
	v700 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v700
	v704 = F_json_lex(m, v22+int32(12))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L41
	} else {
		goto L122
	}
L19:
	;
	v651 = int32(1)
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652)+1)))
	if v653 != v651 {
		v1819 = v651
		goto L1
	} else {
		goto L121
	}
L20:
	;
	if v606 != 0 {
		v663 = v588
		v665 = v590
		goto L18
	} else {
		goto L120
	}
L21:
	;
	if v622 != 0 {
		goto L116
	} else {
		goto L117
	}
L22:
	;
	if base.Ui32(v478) < base.Ui32(v480) {
		goto L103
	} else {
		goto L104
	}
L23:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v478 = v458
	v480 = v476
	goto L22
L24:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v663 = v184
	v665 = v456
	goto L18
L25:
	;
	v208 = int32(0)
	if v60 <= v208 {
		v298 = v208
		v307 = v2
		goto L49
	} else {
		goto L50
	}
L26:
	;
	if base.Ui32(int32(9)) < base.Ui32((v81-int32(48))&int32(255)) {
		v458 = int32(0)
		goto L23
	} else {
		goto L48
	}
L27:
	;
	v84 = int32(0)
	v86 = v60 - int32(1)
	if v86 <= v84 {
		v121 = v84
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v136 == int32(0) {
		goto L19
	} else {
		goto L34
	}
L29:
	;
	v90 = v86
	v93 = v84
	goto L30
L30:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v80))))
	if v109 != int32(92) {
		v121 = v93
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v121 = v86
	goto L28
L32:
	;
	v112 = int32(1)
	v115 = v93 + v112
	if v115 != v86 {
		v90 = v90 - v112
		v93 = v115
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v141 = int32(0)
	v144 = v121
	goto L35
L35:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v161 = int32(*(*int8)(unsafe.Add(mBase, uint32(v159+v141))))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v162 <= v163+int32(1) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L19
L37:
	;
	v183 = int32(1)
	v184 = v141 + v183
	if base.B2i32(v144&v183 == int32(0))&base.B2i32(v161 == int32(34)) != 0 {
		goto L24
	} else {
		goto L43
	}
L38:
	;
	F_appendStringInfoChar(m, v79, v161)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v171+v163))) = uint8(v161)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v176 = v174 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v176
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v180 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v178+v176))) = uint8(v180)
	goto L37
L41:
	;
	return int32(0)
L42:
	;
	goto L37
L43:
	;
	if v161 == int32(92) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v197 = v144 + int32(1)
	goto L46
L45:
	;
	v197 = int32(0)
	goto L46
L46:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v184) < base.Ui32(v198) {
		v141 = v184
		v144 = v197
		goto L35
	} else {
		goto L47
	}
L47:
	;
	goto L36
L48:
	;
	goto L25
L49:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v315 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L50:
	;
	v211 = int32(0)
	if v60 != int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v220 = v211
	v221 = v208
	v223 = int32(0)
	v230 = v2
	goto L54
L52:
	;
	v267 = v211
	v268 = v208
	v277 = v2
	goto L53
L53:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267+v80))))
	v293 = base.B2i32(v286 == int32(46))
	if v286 == int32(46) {
		goto L64
	} else {
		goto L65
	}
L54:
	;
	v238 = v220 + v80
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	v246 = base.B2i32(v239 == int32(46))
	if v239 == int32(46) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v60&int32(1) == int32(0) {
		v298 = v256
		v307 = v258
		goto L49
	} else {
		goto L63
	}
L56:
	;
	v247 = v221
	goto L58
L57:
	;
	v247 = v221 | base.B2i32(v239&int32(223) == int32(69))
	goto L58
L58:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+1)))
	v255 = base.B2i32(v248 == int32(46))
	if v248 == int32(46) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v256 = v247
	goto L61
L60:
	;
	v256 = v247 | base.B2i32(v248&int32(223) == int32(69))
	goto L61
L61:
	;
	v258 = v246 | v255 | v230
	v259 = int32(2)
	v260 = v220 + v259
	v262 = v223 + v259
	if v262 != v60&int32(2147483646) {
		v220 = v260
		v221 = v256
		v223 = v262
		v230 = v258
		goto L54
	} else {
		goto L62
	}
L62:
	;
	goto L55
L63:
	;
	v267 = v260
	v268 = v256
	v277 = v258
	goto L53
L64:
	;
	v294 = v268
	goto L66
L65:
	;
	v294 = v268 | base.B2i32(v286&int32(223) == int32(69))
	goto L66
L66:
	;
	v298 = v294
	v307 = v293 | v277
	goto L49
L67:
	;
	v611 = int32(0)
	v622 = v2
	goto L21
L68:
	;
	goto L69
L69:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v60-int32(1)))))
	v323 = int32(0)
	v326 = v323
	v327 = v298
	v329 = v323
	v330 = v322
	v336 = v307
	goto L70
L70:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344+v329))))
	v347 = base.I32_extend8_s(v346)
	if base.Ui32(int32(10)) <= base.Ui32(v346-int32(48)) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v478 = v451
	v480 = v454
	goto L22
L72:
	;
	v450 = int32(1)
	v451 = v326 + v450
	v453 = v329 + v450
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v453) < base.Ui32(v454) {
		v326 = v451
		v327 = v447
		v329 = v453
		v330 = v347
		v336 = v449
		goto L70
	} else {
		goto L101
	}
L73:
	;
	switch v346 - int32(43) {
	case 0, 2:
		goto L78
	case 1, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
		v458 = v326
		goto L23
	case 3:
		goto L77
	case 26:
		goto L79
	default:
		goto L80
	}
L74:
	;
	goto L75
L75:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v429 <= v430+int32(1) {
		goto L97
	} else {
		goto L98
	}
L76:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v409 <= v410+int32(1) {
		goto L93
	} else {
		goto L94
	}
L77:
	;
	if (v327|v336)&int32(1) != 0 {
		v458 = v326
		goto L23
	} else {
		goto L88
	}
L78:
	;
	if v330&int32(223) != int32(69) {
		v458 = v326
		goto L23
	} else {
		goto L83
	}
L79:
	;
	if v327&int32(1) == int32(0) {
		goto L76
	} else {
		goto L82
	}
L80:
	;
	if v346 != int32(101) {
		v458 = v326
		goto L23
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v458 = v326
	goto L23
L83:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v364 <= v365+int32(1) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	F_appendStringInfoChar(m, v79, v347)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L41
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v371+v365))) = uint8(v347)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v376 = v374 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v376
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v380 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v378+v376))) = uint8(v380)
	v447 = v327
	v449 = v336
	goto L72
L87:
	;
	v447 = v327
	v449 = v336
	goto L72
L88:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v385 <= v386+int32(1) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	F_appendStringInfoChar(m, v79, int32(46))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L41
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v397 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v395+v386))) = uint8(v397)
	v399 = int32(1)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v402 = v400 + v399
	*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v402
	v404 = int32(0)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v405+v402))) = uint8(v404)
	v447 = v404
	v449 = v399
	goto L72
L92:
	;
	v447 = int32(0)
	v449 = int32(1)
	goto L72
L93:
	;
	F_appendStringInfoChar(m, v79, v347)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L41
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v417+v410))) = uint8(v347)
	v420 = int32(1)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v423 = v421 + v420
	*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v423
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v427 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v425+v423))) = uint8(v427)
	v447 = v420
	v449 = v336
	goto L72
L96:
	;
	v447 = int32(1)
	v449 = v336
	goto L72
L97:
	;
	F_appendStringInfoChar(m, v79, v347)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L41
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v436+v430))) = uint8(v347)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v441 = v439 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v441
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v445 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v443+v441))) = uint8(v445)
	v447 = v327
	v449 = v336
	goto L72
L100:
	;
	v447 = v327
	v449 = v336
	goto L72
L101:
	;
	goto L71
L102:
	;
	if v588 != v590 {
		goto L20
	} else {
		goto L115
	}
L103:
	;
	v498 = v478
	v500 = v480
	goto L106
L104:
	;
	v568 = v478
	v570 = v480
	goto L105
L105:
	;
	v588 = v568
	v590 = v570
	v606 = int32(0)
	goto L102
L106:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517+v498))))
	v520 = base.I32_extend8_s(v519)
	v521 = int32(0)
	v527 = int32(255)
	if base.B2i32(v520 < v521)|base.B2i32(base.Ui32((v519&int32(223)-int32(65))&v527) < base.Ui32(int32(26)))|(base.B2i32(v520 == int32(95))|base.B2i32(base.Ui32(int32(246)) <= base.Ui32((v520-int32(58))&v527))) == v521 {
		v588 = v498
		v590 = v500
		v606 = int32(1)
		goto L102
	} else {
		goto L108
	}
L107:
	;
	v568 = v564
	v570 = v565
	goto L105
L108:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v544 <= v545+int32(1) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v564 = v498 + int32(1)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v564) < base.Ui32(v565) {
		v498 = v564
		v500 = v565
		goto L106
	} else {
		goto L114
	}
L110:
	;
	F_appendStringInfoChar(m, v79, v520)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L41
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v551+v545))) = uint8(v520)
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v556 = v554 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v556
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v560 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v558+v556))) = uint8(v560)
	goto L109
L113:
	;
	goto L109
L114:
	;
	goto L107
L115:
	;
	v611 = v590
	v622 = v606
	goto L21
L116:
	;
	v663 = v611
	v665 = v611
	goto L18
L117:
	;
	goto L118
L118:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627)+1)))
	if v628&int32(1) != 0 {
		v663 = v611
		v665 = v611
		goto L18
	} else {
		goto L119
	}
L119:
	;
	v1819 = int32(1)
	goto L1
L120:
	;
	goto L19
L121:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v656
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v656 + v658
	v1819 = int32(15)
	goto L1
L122:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v706
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v708
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v712
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v714
	if v704 != 0 {
		v1819 = v704
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v716 != v714-v712 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v1819 = int32(15)
	goto L1
L125:
	;
	goto L126
L126:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v721 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v720)+2)) = uint8(v721)
	goto L4
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v734
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734))))
	switch v807 - int32(34) {
	case 0:
		goto L146
	default:
		goto L151
	case 10:
		goto L148
	case 11:
		goto L145
	case 14, 15, 16, 17, 18, 19, 20, 21, 22, 23:
		goto L144
	case 24:
		goto L147
	case 57:
		goto L150
	case 59:
		goto L149
	case 89:
		v1773 = int32(3)
		goto L141
	case 91:
		goto L142
	}
L128:
	;
	v730 = v723 + (v727 - v723)
	v734 = v723
	goto L131
L129:
	;
	v777 = v723
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v777
	v797 = int32(1)
	v798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v798 != v797 {
		goto L4
	} else {
		goto L138
	}
L131:
	;
	v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734))))
	v752 = v750 - int32(9)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v752))|base.B2i32(int32(1)<<(uint(v752)%32)&int32(_a_F_json_lex_2) == int32(0)) != 0 {
		goto L127
	} else {
		goto L133
	}
L132:
	;
	v777 = v730
	goto L130
L133:
	;
	v763 = v734 + int32(1)
	if v750 == int32(10) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v763
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v767 + int32(1)
	goto L136
L135:
	;
	goto L136
L136:
	;
	if v763 != v730 {
		v734 = v763
		goto L131
	} else {
		goto L137
	}
L137:
	;
	goto L132
L138:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801)+1)))
	if v802 == int32(1) {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	v1819 = v797
	goto L1
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v1779
	goto L4
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v734 + int32(1)
	v1779 = v1773
	goto L140
L142:
	;
	v1773 = int32(4)
	goto L141
L143:
	;
	if v725 == int32(0) {
		goto L380
	} else {
		goto L381
	}
L144:
	;
	v1730 = int32(0)
	v1732 = F_json_lex_number(m, l0, v734, v1730, v1730)
	mBase = m.M
	v1733 = m.ExcPending
	if v1733 != 0 {
		goto L41
	} else {
		goto L378
	}
L145:
	;
	v1725 = int32(0)
	v1727 = F_json_lex_number(m, l0, v734+int32(1), v1725, v1725)
	mBase = m.M
	v1728 = m.ExcPending
	if v1728 != 0 {
		goto L41
	} else {
		goto L376
	}
L146:
	;
	v887 = m.G0
	v889 = v887 - int32(32)
	m.G0 = v889
	v891 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v893 = v891 + v892
	v894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v894 == int32(1) {
		goto L161
	} else {
		goto L162
	}
L147:
	;
	v1773 = int32(8)
	goto L141
L148:
	;
	v1773 = int32(7)
	goto L141
L149:
	;
	v1773 = int32(6)
	goto L141
L150:
	;
	v1773 = int32(5)
	goto L141
L151:
	;
	if base.Ui32(v734) < base.Ui32(v727) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v815 = v734
	goto L156
L153:
	;
	goto L154
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v734 + int32(1)
	v1819 = int32(15)
	goto L1
L155:
	;
	if v734 != v857 {
		goto L143
	} else {
		goto L160
	}
L156:
	;
	v830 = int32(*(*int8)(unsafe.Add(mBase, uint32(v815))))
	v833 = int32(255)
	v849 = int32(0)
	if base.B2i32(base.B2i32(base.Ui32((v830-int32(48))&v833) < base.Ui32(int32(10)))|base.B2i32(base.Ui32((v830&int32(-33)-int32(65))&v833) < base.Ui32(int32(26)))|base.B2i32(v830 == int32(95)) == v849)&base.B2i32(v849 <= v830) != 0 {
		v857 = v815
		goto L155
	} else {
		goto L158
	}
L157:
	;
	v857 = v730
	goto L155
L158:
	;
	v855 = v815 + int32(1)
	if base.Ui32(v855) < base.Ui32(v727) {
		v815 = v855
		goto L156
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	goto L154
L161:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v897)))
	v899 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v898))) = uint8(v899)
	*(*int32)(unsafe.Add(mBase, uint32(v897)+12)) = v899
	*(*int32)(unsafe.Add(mBase, uint32(v897)+4)) = v899
	goto L164
L162:
	;
	goto L163
L163:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v907 = v905 + int32(1)
	if base.Ui32(v893) <= base.Ui32(v907) {
		v1666 = v907
		goto L166
	} else {
		goto L167
	}
L164:
	;
	goto L163
L165:
	;
	m.G0 = v889 + int32(32)
	if v1718 != 0 {
		v1819 = v1718
		goto L1
	} else {
		goto L375
	}
L166:
	;
	v1684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v1684 != int32(1) {
		goto L371
	} else {
		goto L372
	}
L167:
	;
	v910 = v893 - int32(8)
	v915 = v907
	v916 = v905
	v918 = int32(-1)
	goto L168
L168:
	;
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v916)+1)))
	if v933 != int32(92) {
		goto L172
	} else {
		goto L173
	}
L169:
	;
	v1650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v1650 != int32(1) {
		goto L367
	} else {
		goto L368
	}
L170:
	;
	goto L169
L171:
	;
	v1644 = v1626 + int32(1)
	if base.Ui32(v1644) < base.Ui32(v893) {
		v915 = v1644
		v916 = v1626
		v918 = v1628
		goto L168
	} else {
		goto L366
	}
L172:
	;
	if v933 != int32(34) {
		goto L176
	} else {
		goto L177
	}
L173:
	;
	goto L174
L174:
	;
	v1196 = v916 + int32(2)
	if base.Ui32(v893) <= base.Ui32(v1196) {
		goto L214
	} else {
		goto L215
	}
L175:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1187 = v893 - v915
	v1188 = F_pg_encoding_mblen_or_incomplete(m, v1186, v915, v1187)
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L41
	} else {
		goto L210
	}
L176:
	;
	if v918 != int32(-1) {
		goto L175
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	if v918 != int32(-1) {
		goto L207
	} else {
		goto L208
	}
L179:
	;
	if base.Ui32(v910) <= base.Ui32(v915) {
		v1095 = v915
		goto L180
	} else {
		goto L181
	}
L180:
	;
	if base.Ui32(v893) <= base.Ui32(v1095) {
		v1147 = v1095
		goto L194
	} else {
		goto L195
	}
L181:
	;
	v943 = v915
	goto L182
L182:
	;
	v960 = *(*int64)(unsafe.Add(mBase, uint32(v943)))
	v962 = v960 ^ int64(6655295901103053916)
	if int64(0) <= v960 {
		goto L185
	} else {
		goto L186
	}
L183:
	;
	v1095 = v1091
	goto L180
L184:
	;
	v1091 = v943 + int32(8)
	if base.Ui32(v1091) < base.Ui32(v910) {
		v943 = v1091
		goto L182
	} else {
		goto L193
	}
L185:
	;
	v968 = v960&int64(36170086419038336) ^ int64(-9187201950435737472)
	v969 = int64(72340172838076673)
	if v968&(v962-v969)|v968&(v960^int64(2459565876494606882)-v969) != int64(0) {
		v1095 = v943
		goto L180
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v987 = int64(0)
	if base.B2i32(v962&int64(71776119061217280) == v987)|base.B2i32(v962&int64(280375465082880) == v987)|(base.B2i32(v962&int64(1095216660480) == v987)|base.B2i32(v962&int64(4278190080) == v987))|(base.B2i32(v962&int64(65280) == v987)|(base.B2i32(v962&int64(16711680) == v987)|base.B2i32(v962&int64(255) == v987))) != 0 {
		v1095 = v943
		goto L180
	} else {
		goto L190
	}
L188:
	;
	if v968&(v960-int64(2314885530818453536)) == int64(0) {
		goto L184
	} else {
		goto L189
	}
L189:
	;
	v1095 = v943
	goto L180
L190:
	;
	v1020 = v960 ^ int64(2459565876494606882)
	v1023 = int64(0)
	if base.B2i32(v1020&int64(71776119061217280) == v1023)|base.B2i32(v1020&int64(280375465082880) == v1023)|(base.B2i32(v1020&int64(1095216660480) == v1023)|base.B2i32(v1020&int64(4278190080) == v1023))|(base.B2i32(v1020&int64(16711680) == v1023)|base.B2i32(v1020&int64(255) == v1023)|(base.B2i32(v1020&int64(65280) == v1023)|base.B2i32(v960&int64(63050394783186944) == v1023))) != 0 {
		v1095 = v943
		goto L180
	} else {
		goto L191
	}
L191:
	;
	v1062 = int64(0)
	if base.B2i32(v960&int64(246290604621824) == v1062)|base.B2i32(v960&int64(962072674304) == v1062)|(base.B2i32(v960&int64(3758096384) == v1062)|base.B2i32(v960&int64(14680064) == v1062))|(base.B2i32(v960&int64(224) == v1062)|base.B2i32(v960&int64(57344) == v1062)) != 0 {
		v1095 = v943
		goto L180
	} else {
		goto L192
	}
L192:
	;
	goto L184
L193:
	;
	goto L183
L194:
	;
	v1164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v1164 == int32(1) {
		goto L203
	} else {
		goto L204
	}
L195:
	;
	v1115 = v1095
	goto L196
L196:
	;
	v1132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1115))))
	if base.B2i32(v1132 == int32(34))|base.B2i32(v1132 == int32(92)) != 0 {
		v1147 = v1115
		goto L194
	} else {
		goto L198
	}
L197:
	;
	v1147 = v893
	goto L194
L198:
	;
	if base.Ui32(v1132) <= base.Ui32(int32(31)) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1115
	v1718 = int32(5)
	goto L165
L200:
	;
	goto L201
L201:
	;
	v1143 = v1115 + int32(1)
	if v1143 != v893 {
		v1115 = v1143
		goto L196
	} else {
		goto L202
	}
L202:
	;
	goto L197
L203:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendBinaryStringInfo(m, v1167, v915, v1147-v915)
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L41
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v1626 = v1147 - int32(1)
	v1628 = int32(-1)
	goto L171
L206:
	;
	goto L205
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v916 + int32(2)
	v1718 = int32(22)
	goto L165
L208:
	;
	goto L209
L209:
	;
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1180
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v916 + int32(2)
	v1718 = int32(0)
	goto L165
L210:
	;
	if v1187 < v1188 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v1192 = v893
	goto L213
L212:
	;
	v1192 = v915 + v1188
	goto L213
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1192
	v1718 = int32(22)
	goto L165
L214:
	;
	v1198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v1198 != int32(1) {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	goto L216
L216:
	;
	v1213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1196))))
	if v1213 == int32(117) {
		goto L223
	} else {
		goto L224
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1196
	v1718 = int32(15)
	goto L165
L218:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1201)+1)))
	if v1202 != 0 {
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v1201+int32(4), v1205, v893-v1205)
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L41
	} else {
		goto L220
	}
L220:
	;
	v1718 = int32(1)
	goto L165
L221:
	;
	v1626 = v1196
	v1628 = int32(-1)
	goto L171
L222:
	;
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1615 = v893 - v1612
	v1616 = F_pg_encoding_mblen_or_incomplete(m, v1614, v1612, v1615)
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		goto L41
	} else {
		goto L362
	}
L223:
	;
	v1216 = v893 - int32(2) - v916
	if v1216 == int32(1) {
		goto L170
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	v1441 = base.I32_extend8_s(v1213)
	v1442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v1442 == int32(1) {
		goto L301
	} else {
		goto L302
	}
L226:
	;
	v1220 = v916 + int32(3)
	v1221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1220))))
	v1223 = v1221 - int32(48)
	if base.Ui32(v1223&int32(255)) < base.Ui32(int32(10)) {
		v1244 = v1223
		goto L227
	} else {
		goto L228
	}
L227:
	;
	if v1216 == int32(2) {
		goto L170
	} else {
		goto L233
	}
L228:
	;
	if base.Ui32((v1221-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1244 = v1221 - int32(87)
	goto L227
L230:
	;
	goto L231
L231:
	;
	if base.Ui32(int32(5)) < base.Ui32((v1221-int32(65))&int32(255)) {
		v1612 = v1220
		goto L222
	} else {
		goto L232
	}
L232:
	;
	v1244 = v1221 - int32(55)
	goto L227
L233:
	;
	v1247 = int32(255)
	v1248 = v1244 & v1247
	v1250 = v916 + int32(4)
	v1251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1250))))
	v1255 = (v1251 - int32(48)) & v1247
	if base.Ui32(int32(10)) <= base.Ui32(v1255) {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	if v1216 == int32(3) {
		goto L170
	} else {
		goto L242
	}
L235:
	;
	v1261 = (v1251 - int32(97)) & int32(255)
	if base.Ui32(int32(6)) <= base.Ui32(v1261) {
		goto L238
	} else {
		goto L239
	}
L236:
	;
	goto L237
L237:
	;
	v1284 = v1248<<(uint(int32(4))%32) | v1255
	goto L234
L238:
	;
	v1267 = (v1251 - int32(65)) & int32(255)
	if base.Ui32(int32(5)) < base.Ui32(v1267) {
		v1612 = v1250
		goto L222
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v1284 = v1261 + v1248<<(uint(int32(4))%32) + int32(10)
	goto L234
L241:
	;
	v1284 = v1267 + v1248<<(uint(int32(4))%32) + int32(10)
	goto L234
L242:
	;
	v1288 = v916 + int32(5)
	v1289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1288))))
	v1293 = (v1289 - int32(48)) & int32(255)
	if base.Ui32(int32(10)) <= base.Ui32(v1293) {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	if v1216 == int32(4) {
		goto L170
	} else {
		goto L251
	}
L244:
	;
	v1299 = (v1289 - int32(97)) & int32(255)
	if base.Ui32(int32(6)) <= base.Ui32(v1299) {
		goto L247
	} else {
		goto L248
	}
L245:
	;
	goto L246
L246:
	;
	v1322 = v1284<<(uint(int32(4))%32) | v1293
	goto L243
L247:
	;
	v1305 = (v1289 - int32(65)) & int32(255)
	if base.Ui32(int32(5)) < base.Ui32(v1305) {
		v1612 = v1288
		goto L222
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v1322 = v1299 + v1284<<(uint(int32(4))%32) + int32(10)
	goto L243
L250:
	;
	v1322 = v1305 + v1284<<(uint(int32(4))%32) + int32(10)
	goto L243
L251:
	;
	v1326 = v916 + int32(6)
	v1327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1326))))
	v1331 = (v1327 - int32(48)) & int32(255)
	if base.Ui32(int32(10)) <= base.Ui32(v1331) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	v1362 = v916 + int32(6)
	v1363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v1363 != int32(1) {
		v1626 = v1362
		v1628 = v918
		goto L171
	} else {
		goto L260
	}
L253:
	;
	v1337 = (v1327 - int32(97)) & int32(255)
	if base.Ui32(int32(6)) <= base.Ui32(v1337) {
		goto L256
	} else {
		goto L257
	}
L254:
	;
	goto L255
L255:
	;
	v1360 = v1322<<(uint(int32(4))%32) | v1331
	goto L252
L256:
	;
	v1343 = (v1327 - int32(65)) & int32(255)
	if base.Ui32(int32(5)) < base.Ui32(v1343) {
		v1612 = v1326
		goto L222
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	v1360 = v1337 + v1322<<(uint(int32(4))%32) + int32(10)
	goto L252
L259:
	;
	v1360 = v1343 + v1322<<(uint(int32(4))%32) + int32(10)
	goto L252
L260:
	;
	v1367 = v1360 & int32(67107840)
	if v1367 != int32(_a_F_json_lex_3) {
		goto L264
	} else {
		goto L265
	}
L261:
	;
	v1424 = F_pg_unicode_to_server_noerror(m, v1423, v889)
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L41
	} else {
		goto L292
	}
L262:
	;
	v1423 = v918<<(uint(int32(10))%32)&int32(_a_F_json_lex_4) | v1360&int32(1023) + int32(_a_F_json_lex_5)
	goto L261
L263:
	;
	if v918 != int32(-1) {
		goto L280
	} else {
		goto L281
	}
L264:
	;
	if v1367 != int32(_a_F_json_lex_6) {
		goto L263
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	if v918 != int32(-1) {
		goto L262
	} else {
		goto L275
	}
L267:
	;
	if v918 == int32(-1) {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v1626 = v1362
	v1628 = v1360
	goto L171
L269:
	;
	goto L270
L270:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1375 = v893 - v1362
	v1376 = F_pg_encoding_mblen_or_incomplete(m, v1374, v1362, v1375)
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L41
	} else {
		goto L271
	}
L271:
	;
	if v1375 < v1376 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v1380 = v893
	goto L274
L273:
	;
	v1380 = v1362 + v1376
	goto L274
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1380
	v1718 = int32(21)
	goto L165
L275:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1386 = v893 - v1362
	v1387 = F_pg_encoding_mblen_or_incomplete(m, v1385, v1362, v1386)
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L41
	} else {
		goto L276
	}
L276:
	;
	if v1386 < v1387 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1391 = v893
	goto L279
L278:
	;
	v1391 = v1362 + v1387
	goto L279
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1391
	v1718 = int32(22)
	goto L165
L280:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1397 = v893 - v1362
	v1398 = F_pg_encoding_mblen_or_incomplete(m, v1396, v1362, v1397)
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L41
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	if v1360 != 0 {
		v1423 = v1360
		goto L261
	} else {
		goto L287
	}
L283:
	;
	if v1397 < v1398 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1402 = v893
	goto L286
L285:
	;
	v1402 = v1362 + v1398
	goto L286
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1402
	v1718 = int32(22)
	goto L165
L287:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1406 = v893 - v1362
	v1407 = F_pg_encoding_mblen_or_incomplete(m, v1405, v1362, v1406)
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L41
	} else {
		goto L288
	}
L288:
	;
	if v1406 < v1407 {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1411 = v893
	goto L291
L290:
	;
	v1411 = v1362 + v1407
	goto L291
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1411
	v1718 = int32(17)
	goto L165
L292:
	;
	if v1424 == int32(0) {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1429 = v893 - v1362
	v1430 = F_pg_encoding_mblen_or_incomplete(m, v1428, v1362, v1429)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L41
	} else {
		goto L296
	}
L294:
	;
	goto L295
L295:
	;
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoString(m, v1437, v889)
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L41
	} else {
		goto L300
	}
L296:
	;
	if v1429 < v1430 {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1434 = v893
	goto L299
L298:
	;
	v1434 = v1362 + v1430
	goto L299
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1434
	v1718 = int32(20)
	goto L165
L300:
	;
	v1626 = v1362
	v1628 = int32(-1)
	goto L171
L301:
	;
	if v918 != int32(-1) {
		goto L304
	} else {
		goto L305
	}
L302:
	;
	goto L303
L303:
	;
	goto L334
L304:
	;
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1448 = v893 - v1196
	v1449 = F_pg_encoding_mblen_or_incomplete(m, v1447, v1196, v1448)
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L41
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	switch v1213 - int32(47) {
	case 0, 45:
		goto L317
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 46, 47, 48, 49, 50, 52, 53, 54, 56, 57, 58, 59, 60, 61, 62, 64, 65, 66, 68:
		goto L311
	case 51:
		goto L316
	case 55:
		goto L315
	case 63:
		goto L314
	case 67:
		goto L313
	case 69:
		goto L312
	default:
		goto L318
	}
L307:
	;
	if v1448 < v1449 {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1453 = v893
	goto L310
L309:
	;
	v1453 = v1196 + v1449
	goto L310
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1453
	v1718 = int32(22)
	goto L165
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1196
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1485 = v893 - v1196
	v1486 = F_pg_encoding_mblen_or_incomplete(m, v1484, v1196, v1485)
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L41
	} else {
		goto L326
	}
L312:
	;
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoChar(m, v1479, int32(9))
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L41
	} else {
		goto L325
	}
L313:
	;
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoChar(m, v1475, int32(13))
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		goto L41
	} else {
		goto L324
	}
L314:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoChar(m, v1471, int32(10))
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L41
	} else {
		goto L323
	}
L315:
	;
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoChar(m, v1467, int32(12))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L41
	} else {
		goto L322
	}
L316:
	;
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoChar(m, v1463, int32(8))
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L41
	} else {
		goto L321
	}
L317:
	;
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoChar(m, v1460, v1441)
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L41
	} else {
		goto L320
	}
L318:
	;
	if v1213 != int32(34) {
		goto L311
	} else {
		goto L319
	}
L319:
	;
	goto L317
L320:
	;
	goto L221
L321:
	;
	goto L221
L322:
	;
	goto L221
L323:
	;
	goto L221
L324:
	;
	goto L221
L325:
	;
	goto L221
L326:
	;
	if v1485 < v1486 {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v1490 = v893
	goto L329
L328:
	;
	v1490 = v1196 + v1486
	goto L329
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1490
	v1718 = int32(4)
	goto L165
L330:
	;
	if v1599 != 0 {
		goto L355
	} else {
		goto L356
	}
L331:
	;
	v1599 = int32(0)
	goto L330
L332:
	;
	v1577 = v1570
	v1579 = v1572
	goto L349
L333:
	;
	if base.B2i32(v1516 != v1517) == int32(0) {
		goto L331
	} else {
		goto L340
	}
L334:
	;
	v1508 = int32(_a_F_json_lex_7)
	v1510 = int32(9)
	goto L335
L335:
	;
	v1513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1508))))
	if v1513 == v1441&int32(255) {
		v1570 = v1508
		v1572 = v1510
		goto L332
	} else {
		goto L337
	}
L336:
	;
	goto L333
L337:
	;
	v1515 = int32(1)
	v1516 = v1510 - v1515
	v1517 = int32(0)
	v1520 = v1508 + v1515
	if v1520&int32(3) == v1517 {
		goto L333
	} else {
		goto L338
	}
L338:
	;
	if v1516 != 0 {
		v1508 = v1520
		v1510 = v1516
		goto L335
	} else {
		goto L339
	}
L339:
	;
	goto L336
L340:
	;
	v1533 = v1441 & int32(255)
	v1534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1520))))
	if base.B2i32(v1533 == v1534)|base.B2i32(base.Ui32(v1516) < base.Ui32(int32(4))) == int32(0) {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v1543 = v1520
	v1545 = v1516
	goto L344
L342:
	;
	v1563 = v1520
	v1565 = v1516
	goto L343
L343:
	;
	if v1565 == int32(0) {
		goto L331
	} else {
		goto L348
	}
L344:
	;
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v1543)))
	v1550 = v1549 ^ v1533*int32(16843009)
	v1553 = int32(-2139062144)
	if (int32(16843008)-v1550|v1550)&v1553 != v1553 {
		v1570 = v1543
		v1572 = v1545
		goto L332
	} else {
		goto L346
	}
L345:
	;
	v1563 = v1558
	v1565 = v1560
	goto L343
L346:
	;
	v1557 = int32(4)
	v1558 = v1543 + v1557
	v1560 = v1545 - v1557
	if base.Ui32(int32(3)) < base.Ui32(v1560) {
		v1543 = v1558
		v1545 = v1560
		goto L344
	} else {
		goto L347
	}
L347:
	;
	goto L345
L348:
	;
	v1570 = v1563
	v1572 = v1565
	goto L332
L349:
	;
	v1582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1577))))
	if v1441&int32(255) == v1582 {
		goto L351
	} else {
		goto L352
	}
L350:
	;
	goto L331
L351:
	;
	v1599 = v1577
	goto L330
L352:
	;
	goto L353
L353:
	;
	v1584 = int32(1)
	v1587 = v1579 - v1584
	if v1587 != 0 {
		v1577 = v1577 + v1584
		v1579 = v1587
		goto L349
	} else {
		goto L354
	}
L354:
	;
	goto L350
L355:
	;
	v1626 = v1196
	v1628 = v918
	goto L171
L356:
	;
	goto L357
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1196
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1602 = v893 - v1196
	v1603 = F_pg_encoding_mblen_or_incomplete(m, v1601, v1196, v1602)
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L41
	} else {
		goto L358
	}
L358:
	;
	if v1602 < v1603 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1607 = v893
	goto L361
L360:
	;
	v1607 = v1196 + v1603
	goto L361
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1607
	v1718 = int32(4)
	goto L165
L362:
	;
	if v1615 < v1616 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1620 = v893
	goto L365
L364:
	;
	v1620 = v1612 + v1616
	goto L365
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1620
	v1718 = int32(18)
	goto L165
L366:
	;
	v1666 = v1644
	goto L166
L367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v893
	v1718 = int32(15)
	goto L165
L368:
	;
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1653)+1)))
	if v1654 != 0 {
		goto L367
	} else {
		goto L369
	}
L369:
	;
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v1653+int32(4), v1657, v893-v1657)
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L41
	} else {
		goto L370
	}
L370:
	;
	v1718 = int32(1)
	goto L165
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1666
	v1718 = int32(15)
	goto L165
L372:
	;
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1687)+1)))
	if v1688 != 0 {
		goto L371
	} else {
		goto L373
	}
L373:
	;
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v1687+int32(4), v1691, v893-v1691)
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L41
	} else {
		goto L374
	}
L374:
	;
	v1718 = int32(1)
	goto L165
L375:
	;
	v1779 = int32(1)
	goto L140
L376:
	;
	if v1727 != 0 {
		v1819 = v1727
		goto L1
	} else {
		goto L377
	}
L377:
	;
	v1779 = int32(2)
	goto L140
L378:
	;
	if v1732 != 0 {
		v1819 = v1732
		goto L1
	} else {
		goto L379
	}
L379:
	;
	v1779 = int32(2)
	goto L140
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v857
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v723
	v1752 = int32(15)
	switch v857 - v734 - int32(4) {
	case 0:
		goto L386
	case 1:
		goto L385
	default:
		v1819 = v1752
		goto L1
	}
L381:
	;
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1737)+1)))
	if v1738 != 0 {
		goto L380
	} else {
		goto L382
	}
L382:
	;
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v857 != v1739+v1740 {
		goto L380
	} else {
		goto L383
	}
L383:
	;
	F_appendBinaryStringInfo(m, v1737+int32(4), v734, v727-v734)
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L41
	} else {
		goto L384
	}
L384:
	;
	v1819 = int32(1)
	goto L1
L385:
	;
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v734)))
	v1767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734)+4)))
	if v1764^int32(1936482662)|(v1767^int32(101)) != 0 {
		v1819 = v1752
		goto L1
	} else {
		goto L391
	}
L386:
	;
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v734)))
	if v1756 == int32(1702195828) {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v1779 = int32(9)
	goto L140
L388:
	;
	goto L389
L389:
	;
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(v734)))
	if v1760 != int32(1819047278) {
		v1819 = v1752
		goto L1
	} else {
		goto L390
	}
L390:
	;
	v1779 = int32(11)
	goto L140
L391:
	;
	v1779 = int32(10)
	goto L140
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
func F_json_object(m *base.Module, l0 int32) int64 {
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
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v283 = m.ExcPending
	if v283 != 0 {
		goto L8
	} else {
		goto L88
	}
L2:
	;
	m.G0 = v11 + int32(32)
	return base.I64_extend_i32_u(v267)
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
	return int64(0)
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
	F_errfinish(m, int32(_a_F_json_object_2), int32(1398), int32(_a_F_json_object_3))
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
	F_errfinish(m, int32(_a_F_json_object_2), int32(1405), int32(_a_F_json_object_3))
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
	F_errfinish(m, int32(_a_F_json_object_2), int32(1411), int32(_a_F_json_object_3))
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
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119+v108<<(uint(int32(3))%32))))
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
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184+v175<<(uint(int32(3))%32))))
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
	v286 = m.ExcPending
	if v286 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	F_errmsg(m, int32(_a_F_json_object_9), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_json_object_2), int32(1427), int32(_a_F_json_object_3))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
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
func F_json_to_tsvector_byid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v19 = F_pg_detoast_datum(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v21 = F_parse_jsonb_index_flags(m, v19)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v12)
				v24 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v24
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v24
				v29 = v10 + int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v29
				F_iterate_json_values(m, v14, v21, v10+int32(24))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int64(0)
				} else {
					v35 = F_make_tsvector(m, v29)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int64(0)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v37 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int64(0)
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
								if v41 != v19 {
									F_pfree(m, v19)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int64(0)
									} else {
										m.G0 = v10 + int32(32)
										return base.I64_extend_i32_u(v35)
									}
								} else {
									m.G0 = v10 + int32(32)
									return base.I64_extend_i32_u(v35)
								}
							}
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							if v41 != v19 {
								F_pfree(m, v19)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int64(0)
								} else {
									m.G0 = v10 + int32(32)
									return base.I64_extend_i32_u(v35)
								}
							} else {
								m.G0 = v10 + int32(32)
								return base.I64_extend_i32_u(v35)
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
