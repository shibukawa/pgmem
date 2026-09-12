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
				F_errmsg_internal(m, int32(370110), v6)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(523173), int32(4097), int32(68127))
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
			F_errmsg_internal(m, int32(370110), v6+int32(16))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(523173), int32(4094), int32(68127))
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
				F_errmsg(m, int32(387605), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					F_errfinish(m, int32(518600), int32(611), int32(257073))
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
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
					v35 = int32(4)
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
					if v37&int32(254) == int32(2) {
						v46 = v35
					} else {
						v46 = base.B2i32(v37 == int32(18)) << (uint(v35) % 32)
					}
					if v37 == int32(1) {
						v49 = v35
					} else {
						v49 = v46
					}
					v62 = v49
				} else {
					v50 = int32(1)
					if v32&v50 != 0 {
						v62 = int32(base.Ui32(v32)>>(uint(v50)%32)) - v50
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						v62 = int32(base.Ui32(v56)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v62
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = l1
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int64)(unsafe.Add(mBase, uint32(v13)+76)) = int64(0)
			v72 = int32(base.Ui32(v68) >> (uint(int32(31)) % 32))
			*(*uint8)(unsafe.Add(mBase, uint32(v13)+93)) = uint8(v72)
			*(*uint8)(unsafe.Add(mBase, uint32(v13)+92)) = uint8(v72)
			v76 = v13 + int32(12)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v76
			*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v76
			v81 = m.T0[l3].(func(*base.Module, int32) int32)(m, l1)
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v13)+95)) = uint8(v8)
				*(*uint8)(unsafe.Add(mBase, uint32(v13)+94)) = uint8(v6)
				*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v81 + int32(1)
				if l6 != 0 {
					v115 = F_executeItemOptUnwrapTarget(m, v13+int32(60), v13+int32(32), v13+int32(12), l6, v72)
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						v117 = v115
						m.G0 = v13 + int32(96)
						return v117
					}
				} else {
					if v72 != 0 {
						v115 = F_executeItemOptUnwrapTarget(m, v13+int32(60), v13+int32(32), v13+int32(12), l6, v72)
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							v117 = v115
							m.G0 = v13 + int32(96)
							return v117
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v13))) = int64(0)
						v100 = F_executeItemOptUnwrapTarget(m, v13+int32(60), v13+int32(32), v13+int32(12), v13, int32(0))
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
					}
				}
			}
		}
	}
}
func F_freeJsonLexContext(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l0 == int32(4555352) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v9&int32(2) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_free_attrmap(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v15 != 0 {
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
	F_free_attrmap(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v18 != int32(1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v86&int32(1) != 0 {
		goto L48
	} else {
		goto L49
	}
L14:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	F_pfree(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v25 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_pfree(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v29 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	F_pfree(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v32&int32(4) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	if v63 != 0 {
		goto L34
	} else {
		goto L35
	}
L25:
	;
	v37 = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v38 < v37 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v42 = v37
	v43 = v38
	goto L27
L27:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v42<<(uint(int32(2))%32))))
	if v50 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L24
L29:
	;
	F_pfree(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L7
	} else {
		goto L32
	}
L30:
	;
	v54 = v43
	goto L31
L31:
	;
	v56 = v42 + int32(1)
	if v56 <= v54 {
		v42 = v56
		v43 = v54
		goto L27
	} else {
		goto L33
	}
L32:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v54 = v53
	goto L31
L33:
	;
	goto L28
L34:
	;
	F_pfree(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L7
	} else {
		goto L37
	}
L35:
	;
	v67 = v62
	goto L36
L36:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	if v68 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v67 = v66
	goto L36
L38:
	;
	F_pfree(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L7
	} else {
		goto L41
	}
L39:
	;
	v72 = v67
	goto L40
L40:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+24))
	if v73 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v72 = v71
	goto L40
L42:
	;
	F_pfree(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L45
	}
L43:
	;
	v79 = v72
	goto L44
L44:
	;
	F_pfree(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L47
	}
L45:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v76 == int32(0) {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	v79 = v76
	goto L44
L47:
	;
	goto L13
L48:
	;
	F_pfree(m, l0)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L7
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v94 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(0)), int32(68))
	mBase = m.M
	goto L52
L51:
	;
	return
L52:
	;
	goto L1
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		if base.Ui32(v8) <= base.Ui32(int32(3)) {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v8<<(uint(int32(2))%32))+uint32(_consts[1147])))
			F_appendStringInfoString(m, v11, v16)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
				if v22 != 0 {
					v23 = int32(547297)
				} else {
					v23 = int32(547310)
				}
				F_appendStringInfoString(m, v19, v23)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v27 == int32(0) {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v35 == int32(0) {
							return
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
							if v38 == l2 {
								return
							} else {
								F_get_json_behavior(m, v35, l1, int32(548915))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
						if v30 == l2 {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v35 == int32(0) {
								return
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
								if v38 == l2 {
									return
								} else {
									F_get_json_behavior(m, v35, l1, int32(548915))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							F_get_json_behavior(m, v27, l1, int32(532199))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v35 == int32(0) {
									return
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
									if v38 == l2 {
										return
									} else {
										F_get_json_behavior(m, v35, l1, int32(548915))
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
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
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
			if v22 != 0 {
				v23 = int32(547297)
			} else {
				v23 = int32(547310)
			}
			F_appendStringInfoString(m, v19, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v27 == int32(0) {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v35 == int32(0) {
						return
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
						if v38 == l2 {
							return
						} else {
							F_get_json_behavior(m, v35, l1, int32(548915))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					if v30 == l2 {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v35 == int32(0) {
							return
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
							if v38 == l2 {
								return
							} else {
								F_get_json_behavior(m, v35, l1, int32(548915))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						F_get_json_behavior(m, v27, l1, int32(532199))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v35 == int32(0) {
								return
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
								if v38 == l2 {
									return
								} else {
									F_get_json_behavior(m, v35, l1, int32(548915))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
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
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v27 == int32(0) {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v35 == int32(0) {
				return
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				if v38 == l2 {
					return
				} else {
					F_get_json_behavior(m, v35, l1, int32(548915))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
			if v30 == l2 {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v35 == int32(0) {
					return
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
					if v38 == l2 {
						return
					} else {
						F_get_json_behavior(m, v35, l1, int32(548915))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				F_get_json_behavior(m, v27, l1, int32(532199))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v35 == int32(0) {
						return
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
						if v38 == l2 {
							return
						} else {
							F_get_json_behavior(m, v35, l1, int32(548915))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
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
			F_appendStringInfo(m, l1, int32(208042), v9+int32(16))
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
							v38 = int32(569465)
						} else {
							v38 = int32(552374)
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
								if v42 == int32(3) {
									v47 = int32(586204)
								} else {
									v47 = int32(581025)
								}
								v48 = v47
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v48
								F_appendStringInfo(m, l1, int32(208056), v9)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							case 2:
								v48 = int32(581877)
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v48
								F_appendStringInfo(m, l1, int32(208056), v9)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
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
							v38 = int32(569465)
						} else {
							v38 = int32(552374)
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
								if v42 == int32(3) {
									v47 = int32(586204)
								} else {
									v47 = int32(581025)
								}
								v48 = v47
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v48
								F_appendStringInfo(m, l1, int32(208056), v9)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							case 2:
								v48 = int32(581877)
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v48
								F_appendStringInfo(m, l1, int32(208056), v9)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
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
					v145 = m.ExcPending
					if v145 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v148 = m.ExcPending
						if v148 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(387605), int32(0))
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(518600), int32(799), int32(230597))
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
				} else {
					v54 = int32(4553888)
					v55 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v57
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
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v55
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
												F_appendStringInfoString(m, v83, int32(778193))
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
															F_appendBinaryStringInfo(m, v93, int32(316837), int32(4))
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
														v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														v101 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
														v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														if v102 != 0 {
															v115 = v101
															v116 = int32(0)
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
															F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
															mBase = m.M
															v121 = m.ExcPending
															if v121 != 0 {
																return int32(0)
															} else {
																m.G0 = v9 + int32(16)
																return v79
															}
														} else {
															v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
															if v103 < int32(2) {
																v115 = v101
																v116 = int32(0)
																v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
																F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
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
																	v115 = v101
																	v116 = int32(0)
																	v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
																	F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
																	mBase = m.M
																	v121 = m.ExcPending
																	if v121 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v9 + int32(16)
																		return v79
																	}
																} else {
																	F_appendStringInfoString(m, v101, int32(778373))
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
																		F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
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
														F_appendBinaryStringInfo(m, v93, int32(316837), int32(4))
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
													v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													v101 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
													v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													if v102 != 0 {
														v115 = v101
														v116 = int32(0)
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
														F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return int32(0)
														} else {
															m.G0 = v9 + int32(16)
															return v79
														}
													} else {
														v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
														if v103 < int32(2) {
															v115 = v101
															v116 = int32(0)
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
															F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
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
																v115 = v101
																v116 = int32(0)
																v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
																F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v9 + int32(16)
																	return v79
																}
															} else {
																F_appendStringInfoString(m, v101, int32(778373))
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
																	F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
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
											F_appendStringInfoString(m, v83, int32(778193))
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
														F_appendBinaryStringInfo(m, v93, int32(316837), int32(4))
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
													v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													v101 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
													v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													if v102 != 0 {
														v115 = v101
														v116 = int32(0)
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
														F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return int32(0)
														} else {
															m.G0 = v9 + int32(16)
															return v79
														}
													} else {
														v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
														if v103 < int32(2) {
															v115 = v101
															v116 = int32(0)
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
															F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
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
																v115 = v101
																v116 = int32(0)
																v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
																F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v9 + int32(16)
																	return v79
																}
															} else {
																F_appendStringInfoString(m, v101, int32(778373))
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
																	F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
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
													F_appendBinaryStringInfo(m, v93, int32(316837), int32(4))
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
												v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
												v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												if v102 != 0 {
													v115 = v101
													v116 = int32(0)
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
													F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return int32(0)
													} else {
														m.G0 = v9 + int32(16)
														return v79
													}
												} else {
													v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
													if v103 < int32(2) {
														v115 = v101
														v116 = int32(0)
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
														F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
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
															v115 = v101
															v116 = int32(0)
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
															F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
															mBase = m.M
															v121 = m.ExcPending
															if v121 != 0 {
																return int32(0)
															} else {
																m.G0 = v9 + int32(16)
																return v79
															}
														} else {
															F_appendStringInfoString(m, v101, int32(778373))
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
																F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
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
						F_appendStringInfoString(m, v83, int32(778193))
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
									F_appendBinaryStringInfo(m, v93, int32(316837), int32(4))
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
								v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
								v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
								if v102 != 0 {
									v115 = v101
									v116 = int32(0)
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
									F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(16)
										return v79
									}
								} else {
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
									if v103 < int32(2) {
										v115 = v101
										v116 = int32(0)
										v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
										v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
										F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
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
											v115 = v101
											v116 = int32(0)
											v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
											v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
											F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												m.G0 = v9 + int32(16)
												return v79
											}
										} else {
											F_appendStringInfoString(m, v101, int32(778373))
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
												F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
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
								F_appendBinaryStringInfo(m, v93, int32(316837), int32(4))
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
							v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
							v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v102 != 0 {
								v115 = v101
								v116 = int32(0)
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
								F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(16)
									return v79
								}
							} else {
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
								if v103 < int32(2) {
									v115 = v101
									v116 = int32(0)
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
									F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
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
										v115 = v101
										v116 = int32(0)
										v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
										v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
										F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											m.G0 = v9 + int32(16)
											return v79
										}
									} else {
										F_appendStringInfoString(m, v101, int32(778373))
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
											F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
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
					F_appendStringInfoString(m, v83, int32(778193))
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
								F_appendBinaryStringInfo(m, v93, int32(316837), int32(4))
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
							v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
							v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v102 != 0 {
								v115 = v101
								v116 = int32(0)
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
								F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(16)
									return v79
								}
							} else {
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
								if v103 < int32(2) {
									v115 = v101
									v116 = int32(0)
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
									F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
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
										v115 = v101
										v116 = int32(0)
										v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
										v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
										F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											m.G0 = v9 + int32(16)
											return v79
										}
									} else {
										F_appendStringInfoString(m, v101, int32(778373))
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
											F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
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
							F_appendBinaryStringInfo(m, v93, int32(316837), int32(4))
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
						v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
						v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
						if v102 != 0 {
							v115 = v101
							v116 = int32(0)
							v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
							F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v79
							}
						} else {
							v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
							if v103 < int32(2) {
								v115 = v101
								v116 = int32(0)
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
								F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
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
									v115 = v101
									v116 = int32(0)
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
									F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(16)
										return v79
									}
								} else {
									F_appendStringInfoString(m, v101, int32(778373))
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
										F_datum_to_json_internal(m, v100, v116, v115, v117, v118, v116)
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
		v132 = m.ExcPending
		if v132 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(67023), int32(0))
			mBase = m.M
			v136 = m.ExcPending
			if v136 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(518600), int32(789), int32(230597))
				mBase = m.M
				v141 = m.ExcPending
				if v141 != 0 {
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int64
	_ = v137
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v549 int32
	_ = v549
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v641 int32
	_ = v641
	var v651 int32
	_ = v651
	var v665 int64
	_ = v665
	var v667 int64
	_ = v667
	var v673 int64
	_ = v673
	var v720 int64
	_ = v720
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v801 int32
	_ = v801
	var v815 int32
	_ = v815
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v859 int32
	_ = v859
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v940 int32
	_ = v940
	var v946 int32
	_ = v946
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v979 int32
	_ = v979
	var v985 int32
	_ = v985
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1018 int32
	_ = v1018
	var v1024 int32
	_ = v1024
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1049 int32
	_ = v1049
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1213 int32
	_ = v1213
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1322 int32
	_ = v1322
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1344 int32
	_ = v1344
	var v1359 int32
	_ = v1359
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1390 int32
	_ = v1390
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1419 int32
	_ = v1419
	var v1424 int32
	_ = v1424
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1486 int32
	_ = v1486
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1508 int32
	_ = v1508
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1572 int32
	_ = v1572
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1604 int32
	_ = v1604
	v17 = m.G0
	v19 = v17 - int32(80)
	m.G0 = v19
	v21 = int32(16)
	if l0 == int32(4555352) {
		v1604 = v21
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v19 + int32(80)
	return v1604
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v24 == int32(4555420) {
		v1604 = v21
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v29 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v1604 = int32(0)
	goto L1
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1527 - v1526
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1543 + v1526
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v1546
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v1546
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v1546
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v1550
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v1552
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1555 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+36)) = uint8(v1555)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v1554
	v1558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+68)) = uint8(v1558)
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v1560
	v1564 = F_json_lex(m, v19+int32(12))
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L54
	} else {
		goto L398
	}
L6:
	;
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1526 = v251
	v1527 = v1524
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v73
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	switch v529 - int32(34) {
	case 0:
		goto L109
	default:
		goto L114
	case 10:
		goto L111
	case 11:
		goto L108
	case 14, 15, 16, 17, 18, 19, 20, 21, 22, 23:
		goto L107
	case 24:
		goto L110
	case 57:
		goto L113
	case 59:
		goto L112
	case 89:
		v1502 = int32(3)
		goto L104
	case 91:
		goto L105
	}
L8:
	;
	v137 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v19-int32(-64)))) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v137
	v154 = v59 + int32(4)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	switch v156 - int32(34) {
	case 0:
		goto L39
	default:
		goto L38
	case 11:
		goto L40
	}
L9:
	;
	v66 = v27 + v28
	if base.Ui32(v63) < base.Ui32(v66) {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v60 != 0 {
		goto L8
	} else {
		goto L17
	}
L11:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v29 != 0 {
		v58 = v56
		goto L10
	} else {
		goto L16
	}
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)))
	if v32 != int32(1) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v36 = v24 + int32(4)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v38 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v38)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v38
	goto L14
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v47 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+2)) = uint8(v47)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v51&int32(1) == v47 {
		v63 = v49
		v65 = v47
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v58 = v49
	goto L10
L16:
	;
	v63 = v56
	v65 = int32(0)
	goto L9
L17:
	;
	v63 = v58
	v65 = int32(1)
	goto L9
L18:
	;
	v69 = v63 + (v66 - v63)
	v73 = v63
	goto L21
L19:
	;
	v112 = v63
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v112
	v129 = int32(1)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v130 != v129 {
		goto L4
	} else {
		goto L29
	}
L21:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	v88 = v86 - int32(9)
	if base.Ui32(int32(23)) < base.Ui32(v88) {
		goto L7
	} else {
		goto L23
	}
L22:
	;
	v112 = v69
	goto L20
L23:
	;
	if int32(1)<<(uint(v88)%32)&int32(8388627) == int32(0) {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v98 = v73 + int32(1)
	if v86 == int32(10) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v98
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v102 + int32(1)
	goto L27
L26:
	;
	goto L27
L27:
	;
	if v98 != v69 {
		v73 = v98
		goto L21
	} else {
		goto L28
	}
L28:
	;
	goto L22
L29:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+1)))
	if v134 == int32(1) {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v1604 = v129
	goto L1
L31:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v521
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v521 + v523
	v1604 = int32(15)
	goto L1
L32:
	;
	v498 = int32(1)
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+1)))
	v502 = v500 & v498
	if v502 != 0 {
		v1526 = v461
		v1527 = v461
		goto L5
	} else {
		goto L101
	}
L33:
	;
	v493 = int32(1)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+1)))
	if v495&v493 != 0 {
		goto L31
	} else {
		goto L100
	}
L34:
	;
	if v457 != 0 {
		v1526 = v442
		v1527 = v443
		goto L5
	} else {
		goto L99
	}
L35:
	;
	if v463 == int32(0) {
		goto L32
	} else {
		goto L98
	}
L36:
	;
	if base.Ui32(v346) < base.Ui32(v347) {
		goto L81
	} else {
		goto L82
	}
L37:
	;
	if v276 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L38:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(int32(9)) < base.Ui32((v156-int32(48))&int32(255)) {
		v346 = int32(0)
		v347 = v267
		goto L36
	} else {
		goto L61
	}
L39:
	;
	v160 = int32(0)
	v162 = v60 - int32(1)
	if v162 <= v160 {
		v194 = v160
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v276 = v159
	goto L37
L41:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v206 == int32(0) {
		goto L33
	} else {
		goto L47
	}
L42:
	;
	v166 = v162
	v169 = v160
	goto L43
L43:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166+v155))))
	if v182 != int32(92) {
		v194 = v169
		goto L41
	} else {
		goto L45
	}
L44:
	;
	v194 = v162
	goto L41
L45:
	;
	v185 = int32(1)
	v188 = v169 + v185
	if v188 != v162 {
		v166 = v166 - v185
		v169 = v188
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v211 = int32(0)
	v214 = v194
	goto L48
L48:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v228 = int32(*(*int8)(unsafe.Add(mBase, uint32(v226+v211))))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v229 <= v230+int32(1) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L33
L50:
	;
	v250 = int32(1)
	v251 = v211 + v250
	if base.B2i32(v214&v250 == int32(0))&base.B2i32(v228 == int32(34)) != 0 {
		goto L6
	} else {
		goto L56
	}
L51:
	;
	F_appendStringInfoChar(m, v154, v228)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v238+v230))) = uint8(v228)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v243 = v241 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v243
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v247 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v245+v243))) = uint8(v247)
	goto L50
L54:
	;
	return int32(0)
L55:
	;
	goto L50
L56:
	;
	if v228 == int32(92) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v264 = v214 + int32(1)
	goto L59
L58:
	;
	v264 = int32(0)
	goto L59
L59:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v251) < base.Ui32(v265) {
		v211 = v251
		v214 = v264
		goto L48
	} else {
		goto L60
	}
L60:
	;
	goto L49
L61:
	;
	v276 = v267
	goto L37
L62:
	;
	v279 = int32(0)
	v461 = v279
	v463 = v279
	goto L35
L63:
	;
	goto L64
L64:
	;
	v281 = int32(0)
	v284 = v281
	v285 = v276
	v287 = v281
	goto L65
L65:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299+v287))))
	if base.Ui32(v301-int32(48)) < base.Ui32(int32(10)) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v346 = v340
	v347 = v343
	goto L36
L67:
	;
	v319 = base.I32_extend8_s(v301)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v320 <= v321+int32(1) {
		goto L75
	} else {
		goto L76
	}
L68:
	;
	v307 = v301 - int32(43)
	if int32(1)<<(uint(v307)%32)&int32(67108869) != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v315 = base.B2i32(base.Ui32(v307) <= base.Ui32(int32(26)))
	goto L71
L70:
	;
	v315 = int32(0)
	goto L71
L71:
	;
	if v315 != 0 {
		goto L67
	} else {
		goto L72
	}
L72:
	;
	if v301 != int32(101) {
		v346 = v284
		v347 = v285
		goto L36
	} else {
		goto L73
	}
L73:
	;
	goto L67
L74:
	;
	v339 = int32(1)
	v340 = v284 + v339
	v342 = v287 + v339
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v342) < base.Ui32(v343) {
		v284 = v340
		v285 = v343
		v287 = v342
		goto L65
	} else {
		goto L79
	}
L75:
	;
	F_appendStringInfoChar(m, v154, v319)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L54
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v327+v321))) = uint8(v319)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v332 = v330 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v332
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v336 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v334+v332))) = uint8(v336)
	goto L74
L78:
	;
	goto L74
L79:
	;
	goto L66
L80:
	;
	if v442 != v443 {
		goto L34
	} else {
		goto L97
	}
L81:
	;
	v363 = v346
	v364 = v347
	goto L84
L82:
	;
	v425 = v346
	v426 = v347
	goto L83
L83:
	;
	v442 = v425
	v443 = v426
	v457 = int32(0)
	goto L80
L84:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378+v363))))
	v381 = base.I32_extend8_s(v380)
	if base.Ui32((v380&int32(223)-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v425 = v421
	v426 = v422
	goto L83
L86:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v401 <= v402+int32(1) {
		goto L92
	} else {
		goto L93
	}
L87:
	;
	if v381 < int32(0) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	if v381 == int32(95) {
		goto L86
	} else {
		goto L89
	}
L89:
	;
	if base.Ui32(int32(246)) <= base.Ui32((v381-int32(58))&int32(255)) {
		goto L86
	} else {
		goto L90
	}
L90:
	;
	v442 = v363
	v443 = v364
	v457 = int32(1)
	goto L80
L91:
	;
	v421 = v363 + int32(1)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v421) < base.Ui32(v422) {
		v363 = v421
		v364 = v422
		goto L84
	} else {
		goto L96
	}
L92:
	;
	F_appendStringInfoChar(m, v154, v381)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L54
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v408+v402))) = uint8(v381)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v413 = v411 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v413
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v417 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v415+v413))) = uint8(v417)
	goto L91
L95:
	;
	goto L91
L96:
	;
	goto L85
L97:
	;
	v461 = v443
	v463 = v457
	goto L35
L98:
	;
	v1526 = v461
	v1527 = v461
	goto L5
L99:
	;
	goto L33
L100:
	;
	v1604 = v493
	goto L1
L101:
	;
	if v502 == int32(0) {
		v1604 = v498
		goto L1
	} else {
		goto L102
	}
L102:
	;
	goto L31
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v1508
	goto L4
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v73 + int32(1)
	v1508 = v1502
	goto L103
L105:
	;
	v1502 = int32(4)
	goto L104
L106:
	;
	if v65 == int32(0) {
		goto L368
	} else {
		goto L369
	}
L107:
	;
	v1402 = int32(0)
	v1404 = F_json_lex_number(m, l0, v73, v1402, v1402)
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L54
	} else {
		goto L366
	}
L108:
	;
	v1397 = int32(0)
	v1399 = F_json_lex_number(m, l0, v73+int32(1), v1397, v1397)
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L54
	} else {
		goto L364
	}
L109:
	;
	v598 = m.G0
	v600 = v598 - int32(32)
	m.G0 = v600
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v604 = v602 + v603
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v605 == int32(1) {
		goto L128
	} else {
		goto L129
	}
L110:
	;
	v1502 = int32(8)
	goto L104
L111:
	;
	v1502 = int32(7)
	goto L104
L112:
	;
	v1502 = int32(6)
	goto L104
L113:
	;
	v1502 = int32(5)
	goto L104
L114:
	;
	if base.Ui32(v73) < base.Ui32(v66) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v534 = v73
	goto L119
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v73 + int32(1)
	v1604 = int32(15)
	goto L1
L118:
	;
	if v571 != v73 {
		goto L106
	} else {
		goto L127
	}
L119:
	;
	v549 = int32(*(*int8)(unsafe.Add(mBase, uint32(v534))))
	if base.Ui32((v549-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v571 = v69
	goto L118
L121:
	;
	v569 = v534 + int32(1)
	if base.Ui32(v569) < base.Ui32(v66) {
		v534 = v569
		goto L119
	} else {
		goto L126
	}
L122:
	;
	if base.Ui32((v549&int32(-33)-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	if v549 == int32(95) {
		goto L121
	} else {
		goto L124
	}
L124:
	;
	if int32(0) <= v549 {
		v571 = v534
		goto L118
	} else {
		goto L125
	}
L125:
	;
	goto L121
L126:
	;
	goto L120
L127:
	;
	goto L117
L128:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v608)))
	v610 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v609))) = uint8(v610)
	*(*int32)(unsafe.Add(mBase, uint32(v608)+12)) = v610
	*(*int32)(unsafe.Add(mBase, uint32(v608)+4)) = v610
	goto L131
L129:
	;
	goto L130
L130:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v618 = v616 + int32(1)
	if base.Ui32(v604) <= base.Ui32(v618) {
		v1344 = v618
		goto L133
	} else {
		goto L134
	}
L131:
	;
	goto L130
L132:
	;
	m.G0 = v600 + int32(32)
	if v1390 != 0 {
		v1604 = v1390
		goto L1
	} else {
		goto L363
	}
L133:
	;
	v1359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v1359 != int32(1) {
		goto L359
	} else {
		goto L360
	}
L134:
	;
	v621 = v604 - int32(8)
	v626 = v618
	v627 = v616
	v628 = int32(-1)
	goto L135
L135:
	;
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627)+1)))
	if v641 != int32(92) {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	v1328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v1328 != int32(1) {
		goto L355
	} else {
		goto L356
	}
L137:
	;
	goto L136
L138:
	;
	v1322 = v1307 + int32(1)
	if base.Ui32(v1322) < base.Ui32(v604) {
		v626 = v1322
		v627 = v1307
		v628 = v1308
		goto L135
	} else {
		goto L354
	}
L139:
	;
	if v641 != int32(34) {
		goto L143
	} else {
		goto L144
	}
L140:
	;
	goto L141
L141:
	;
	v875 = v627 + int32(2)
	if base.Ui32(v604) <= base.Ui32(v875) {
		goto L201
	} else {
		goto L202
	}
L142:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v866 = v604 - v626
	v867 = F_pg_encoding_mblen_or_incomplete(m, v865, v626, v866)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L54
	} else {
		goto L197
	}
L143:
	;
	if v628 != int32(-1) {
		goto L142
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	if v628 != int32(-1) {
		goto L194
	} else {
		goto L195
	}
L146:
	;
	if base.Ui32(v621) <= base.Ui32(v626) {
		v784 = v626
		goto L147
	} else {
		goto L148
	}
L147:
	;
	if base.Ui32(v604) <= base.Ui32(v784) {
		v829 = v784
		goto L180
	} else {
		goto L181
	}
L148:
	;
	v651 = v626
	goto L149
L149:
	;
	v665 = *(*int64)(unsafe.Add(mBase, uint32(v651)))
	v667 = v665 ^ int64(6655295901103053916)
	if int64(0) <= v665 {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	v784 = v780
	goto L147
L151:
	;
	v780 = v651 + int32(8)
	if base.Ui32(v780) < base.Ui32(v621) {
		v651 = v780
		goto L149
	} else {
		goto L179
	}
L152:
	;
	v673 = v665&int64(36170086419038336) ^ int64(-9187201950435737472)
	if v673&(v667-int64(72340172838076673)) != int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	if v667&int64(71776119061217280) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L158
	}
L155:
	;
	if v673&(v665^int64(2459565876494606882)-int64(72340172838076673)) != int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L156
	}
L156:
	;
	if v673&(v665-int64(2314885530818453536)) == int64(0) {
		goto L151
	} else {
		goto L157
	}
L157:
	;
	v784 = v651
	goto L147
L158:
	;
	if v667&int64(280375465082880) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L159
	}
L159:
	;
	if v667&int64(1095216660480) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L160
	}
L160:
	;
	if v667&int64(4278190080) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L161
	}
L161:
	;
	if v667&int64(16711680) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L162
	}
L162:
	;
	if v667&int64(255) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L163
	}
L163:
	;
	if v667&int64(65280) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L164
	}
L164:
	;
	v720 = v665 ^ int64(2459565876494606882)
	if v720&int64(71776119061217280) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L165
	}
L165:
	;
	if v720&int64(280375465082880) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L166
	}
L166:
	;
	if v720&int64(1095216660480) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L167
	}
L167:
	;
	if v720&int64(4278190080) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L168
	}
L168:
	;
	if v720&int64(16711680) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L169
	}
L169:
	;
	if v720&int64(255) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L170
	}
L170:
	;
	if v720&int64(65280) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L171
	}
L171:
	;
	if v665&int64(63050394783186944) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L172
	}
L172:
	;
	if v665&int64(246290604621824) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L173
	}
L173:
	;
	if v665&int64(962072674304) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L174
	}
L174:
	;
	if v665&int64(3758096384) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L175
	}
L175:
	;
	if v665&int64(14680064) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L176
	}
L176:
	;
	if v665&int64(224) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L177
	}
L177:
	;
	if v665&int64(57344) == int64(0) {
		v784 = v651
		goto L147
	} else {
		goto L178
	}
L178:
	;
	goto L151
L179:
	;
	goto L150
L180:
	;
	v843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v843 == int32(1) {
		goto L190
	} else {
		goto L191
	}
L181:
	;
	v801 = v784
	goto L182
L182:
	;
	v815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801))))
	if v815 == int32(34) {
		v829 = v801
		goto L180
	} else {
		goto L184
	}
L183:
	;
	v829 = v604
	goto L180
L184:
	;
	if v815 == int32(92) {
		v829 = v801
		goto L180
	} else {
		goto L185
	}
L185:
	;
	if base.Ui32(v815) <= base.Ui32(int32(31)) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v801
	v1390 = int32(5)
	goto L132
L187:
	;
	goto L188
L188:
	;
	v825 = v801 + int32(1)
	if v825 != v604 {
		v801 = v825
		goto L182
	} else {
		goto L189
	}
L189:
	;
	goto L183
L190:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendBinaryStringInfo(m, v846, v626, v829-v626)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L54
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v1307 = v829 - int32(1)
	v1308 = int32(-1)
	goto L138
L193:
	;
	goto L192
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v627 + int32(2)
	v1390 = int32(22)
	goto L132
L195:
	;
	goto L196
L196:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v859
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v627 + int32(2)
	v1390 = int32(0)
	goto L132
L197:
	;
	if v866 < v867 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v871 = v604
	goto L200
L199:
	;
	v871 = v626 + v867
	goto L200
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v871
	v1390 = int32(22)
	goto L132
L201:
	;
	v877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v877 != int32(1) {
		goto L204
	} else {
		goto L205
	}
L202:
	;
	goto L203
L203:
	;
	v892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v875))))
	if v892 == int32(117) {
		goto L210
	} else {
		goto L211
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v875
	v1390 = int32(15)
	goto L132
L205:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v880)+1)))
	if v881 != 0 {
		goto L204
	} else {
		goto L206
	}
L206:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v880+int32(4), v884, v604-v884)
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L54
	} else {
		goto L207
	}
L207:
	;
	v1390 = int32(1)
	goto L132
L208:
	;
	v1307 = v875
	v1308 = int32(-1)
	goto L138
L209:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1296 = v604 - v1293
	v1297 = F_pg_encoding_mblen_or_incomplete(m, v1295, v1293, v1296)
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L54
	} else {
		goto L350
	}
L210:
	;
	v895 = v604 - int32(2) - v627
	if v895 == int32(1) {
		goto L137
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v1123 = base.I32_extend8_s(v892)
	v1124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v1124 == int32(1) {
		goto L288
	} else {
		goto L289
	}
L213:
	;
	v899 = v627 + int32(3)
	v900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v899))))
	v902 = v900 - int32(48)
	if base.Ui32(v902&int32(255)) < base.Ui32(int32(10)) {
		v923 = v902
		goto L214
	} else {
		goto L215
	}
L214:
	;
	if v895 == int32(2) {
		goto L137
	} else {
		goto L220
	}
L215:
	;
	if base.Ui32((v900-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v923 = v900 - int32(87)
	goto L214
L217:
	;
	goto L218
L218:
	;
	if base.Ui32(int32(5)) < base.Ui32((v900-int32(65))&int32(255)) {
		v1293 = v899
		goto L209
	} else {
		goto L219
	}
L219:
	;
	v923 = v900 - int32(55)
	goto L214
L220:
	;
	v926 = int32(255)
	v927 = v923 & v926
	v929 = v627 + int32(4)
	v930 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v929))))
	v934 = (v930 - int32(48)) & v926
	if base.Ui32(int32(10)) <= base.Ui32(v934) {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	if v895 == int32(3) {
		goto L137
	} else {
		goto L229
	}
L222:
	;
	v940 = (v930 - int32(97)) & int32(255)
	if base.Ui32(int32(6)) <= base.Ui32(v940) {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	goto L224
L224:
	;
	v964 = v927<<(uint(int32(4))%32) | v934
	goto L221
L225:
	;
	v946 = (v930 - int32(65)) & int32(255)
	if base.Ui32(int32(5)) < base.Ui32(v946) {
		v1293 = v929
		goto L209
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v964 = v940 + v927<<(uint(int32(4))%32) + int32(10)
	goto L221
L228:
	;
	v964 = v946 + v927<<(uint(int32(4))%32) + int32(10)
	goto L221
L229:
	;
	v968 = v627 + int32(5)
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v968))))
	v973 = (v969 - int32(48)) & int32(255)
	if base.Ui32(int32(10)) <= base.Ui32(v973) {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	if v895 == int32(4) {
		goto L137
	} else {
		goto L238
	}
L231:
	;
	v979 = (v969 - int32(97)) & int32(255)
	if base.Ui32(int32(6)) <= base.Ui32(v979) {
		goto L234
	} else {
		goto L235
	}
L232:
	;
	goto L233
L233:
	;
	v1003 = v964<<(uint(int32(4))%32) | v973
	goto L230
L234:
	;
	v985 = (v969 - int32(65)) & int32(255)
	if base.Ui32(int32(5)) < base.Ui32(v985) {
		v1293 = v968
		goto L209
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v1003 = v979 + v964<<(uint(int32(4))%32) + int32(10)
	goto L230
L237:
	;
	v1003 = v985 + v964<<(uint(int32(4))%32) + int32(10)
	goto L230
L238:
	;
	v1007 = v627 + int32(6)
	v1008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1007))))
	v1012 = (v1008 - int32(48)) & int32(255)
	if base.Ui32(int32(10)) <= base.Ui32(v1012) {
		goto L240
	} else {
		goto L241
	}
L239:
	;
	v1044 = v627 + int32(6)
	v1045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v1045 != int32(1) {
		v1307 = v1044
		v1308 = v628
		goto L138
	} else {
		goto L247
	}
L240:
	;
	v1018 = (v1008 - int32(97)) & int32(255)
	if base.Ui32(int32(6)) <= base.Ui32(v1018) {
		goto L243
	} else {
		goto L244
	}
L241:
	;
	goto L242
L242:
	;
	v1042 = v1003<<(uint(int32(4))%32) | v1012
	goto L239
L243:
	;
	v1024 = (v1008 - int32(65)) & int32(255)
	if base.Ui32(int32(5)) < base.Ui32(v1024) {
		v1293 = v1007
		goto L209
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v1042 = v1018 + v1003<<(uint(int32(4))%32) + int32(10)
	goto L239
L246:
	;
	v1042 = v1024 + v1003<<(uint(int32(4))%32) + int32(10)
	goto L239
L247:
	;
	v1049 = v1042 & int32(67107840)
	if v1049 != int32(56320) {
		goto L251
	} else {
		goto L252
	}
L248:
	;
	v1106 = F_pg_unicode_to_server_noerror(m, v1105, v600)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L54
	} else {
		goto L279
	}
L249:
	;
	v1105 = v628<<(uint(int32(10))%32)&int32(1047552) | v1042&int32(1023) + int32(65536)
	goto L248
L250:
	;
	if v628 != int32(-1) {
		goto L267
	} else {
		goto L268
	}
L251:
	;
	if v1049 != int32(55296) {
		goto L250
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	if v628 != int32(-1) {
		goto L249
	} else {
		goto L262
	}
L254:
	;
	if v628 == int32(-1) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1307 = v1044
	v1308 = v1042
	goto L138
L256:
	;
	goto L257
L257:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1057 = v604 - v1044
	v1058 = F_pg_encoding_mblen_or_incomplete(m, v1056, v1044, v1057)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L54
	} else {
		goto L258
	}
L258:
	;
	if v1057 < v1058 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1062 = v604
	goto L261
L260:
	;
	v1062 = v1044 + v1058
	goto L261
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1062
	v1390 = int32(21)
	goto L132
L262:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1068 = v604 - v1044
	v1069 = F_pg_encoding_mblen_or_incomplete(m, v1067, v1044, v1068)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L54
	} else {
		goto L263
	}
L263:
	;
	if v1068 < v1069 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v1073 = v604
	goto L266
L265:
	;
	v1073 = v1044 + v1069
	goto L266
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1073
	v1390 = int32(22)
	goto L132
L267:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1079 = v604 - v1044
	v1080 = F_pg_encoding_mblen_or_incomplete(m, v1078, v1044, v1079)
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L54
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	if v1042 != 0 {
		v1105 = v1042
		goto L248
	} else {
		goto L274
	}
L270:
	;
	if v1079 < v1080 {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v1084 = v604
	goto L273
L272:
	;
	v1084 = v1044 + v1080
	goto L273
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1084
	v1390 = int32(22)
	goto L132
L274:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1088 = v604 - v1044
	v1089 = F_pg_encoding_mblen_or_incomplete(m, v1087, v1044, v1088)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L54
	} else {
		goto L275
	}
L275:
	;
	if v1088 < v1089 {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1093 = v604
	goto L278
L277:
	;
	v1093 = v1044 + v1089
	goto L278
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1093
	v1390 = int32(17)
	goto L132
L279:
	;
	if v1106 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1111 = v604 - v1044
	v1112 = F_pg_encoding_mblen_or_incomplete(m, v1110, v1044, v1111)
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L54
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoString(m, v1119, v600)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L54
	} else {
		goto L287
	}
L283:
	;
	if v1111 < v1112 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1116 = v604
	goto L286
L285:
	;
	v1116 = v1044 + v1112
	goto L286
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1116
	v1390 = int32(20)
	goto L132
L287:
	;
	v1307 = v1044
	v1308 = int32(-1)
	goto L138
L288:
	;
	if v628 != int32(-1) {
		goto L291
	} else {
		goto L292
	}
L289:
	;
	goto L290
L290:
	;
	goto L321
L291:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1130 = v604 - v875
	v1131 = F_pg_encoding_mblen_or_incomplete(m, v1129, v875, v1130)
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L54
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	switch v892 - int32(47) {
	case 0, 45:
		goto L304
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 46, 47, 48, 49, 50, 52, 53, 54, 56, 57, 58, 59, 60, 61, 62, 64, 65, 66, 68:
		goto L298
	case 51:
		goto L303
	case 55:
		goto L302
	case 63:
		goto L301
	case 67:
		goto L300
	case 69:
		goto L299
	default:
		goto L305
	}
L294:
	;
	if v1130 < v1131 {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1135 = v604
	goto L297
L296:
	;
	v1135 = v875 + v1131
	goto L297
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1135
	v1390 = int32(22)
	goto L132
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v875
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1167 = v604 - v875
	v1168 = F_pg_encoding_mblen_or_incomplete(m, v1166, v875, v1167)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L54
	} else {
		goto L313
	}
L299:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoChar(m, v1161, int32(9))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L54
	} else {
		goto L312
	}
L300:
	;
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoChar(m, v1157, int32(13))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L54
	} else {
		goto L311
	}
L301:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoChar(m, v1153, int32(10))
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L54
	} else {
		goto L310
	}
L302:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoChar(m, v1149, int32(12))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L54
	} else {
		goto L309
	}
L303:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoChar(m, v1145, int32(8))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L54
	} else {
		goto L308
	}
L304:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_appendStringInfoChar(m, v1142, v1123)
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L54
	} else {
		goto L307
	}
L305:
	;
	if v892 != int32(34) {
		goto L298
	} else {
		goto L306
	}
L306:
	;
	goto L304
L307:
	;
	goto L208
L308:
	;
	goto L208
L309:
	;
	goto L208
L310:
	;
	goto L208
L311:
	;
	goto L208
L312:
	;
	goto L208
L313:
	;
	if v1167 < v1168 {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1172 = v604
	goto L316
L315:
	;
	v1172 = v875 + v1168
	goto L316
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1172
	v1390 = int32(4)
	goto L132
L317:
	;
	if v1279 != 0 {
		goto L343
	} else {
		goto L344
	}
L318:
	;
	v1279 = int32(0)
	goto L317
L319:
	;
	v1257 = v1250
	v1259 = v1252
	goto L337
L320:
	;
	if base.B2i32(v1197 != v1198) == int32(0) {
		goto L318
	} else {
		goto L328
	}
L321:
	;
	goto L322
L322:
	;
	v1189 = int32(87310)
	v1191 = int32(9)
	goto L323
L323:
	;
	v1194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1189))))
	if v1194 == v1123&int32(255) {
		v1250 = v1189
		v1252 = v1191
		goto L319
	} else {
		goto L325
	}
L324:
	;
	goto L320
L325:
	;
	v1196 = int32(1)
	v1197 = v1191 - v1196
	v1198 = int32(0)
	v1201 = v1189 + v1196
	if v1201&int32(3) == v1198 {
		goto L320
	} else {
		goto L326
	}
L326:
	;
	if v1197 != 0 {
		v1189 = v1201
		v1191 = v1197
		goto L323
	} else {
		goto L327
	}
L327:
	;
	goto L324
L328:
	;
	v1213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1201))))
	if v1213 == v1123&int32(255) {
		v1243 = v1201
		v1245 = v1197
		goto L329
	} else {
		goto L330
	}
L329:
	;
	if v1245 == int32(0) {
		goto L318
	} else {
		goto L336
	}
L330:
	;
	if base.Ui32(v1197) < base.Ui32(int32(4)) {
		v1243 = v1201
		v1245 = v1197
		goto L329
	} else {
		goto L331
	}
L331:
	;
	v1223 = v1201
	v1225 = v1197
	goto L332
L332:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v1223)))
	v1230 = v1229 ^ v1123&int32(255)*int32(16843009)
	v1233 = int32(-2139062144)
	if (int32(16843008)-v1230|v1230)&v1233 != v1233 {
		v1250 = v1223
		v1252 = v1225
		goto L319
	} else {
		goto L334
	}
L333:
	;
	v1243 = v1238
	v1245 = v1240
	goto L329
L334:
	;
	v1237 = int32(4)
	v1238 = v1223 + v1237
	v1240 = v1225 - v1237
	if base.Ui32(int32(3)) < base.Ui32(v1240) {
		v1223 = v1238
		v1225 = v1240
		goto L332
	} else {
		goto L335
	}
L335:
	;
	goto L333
L336:
	;
	v1250 = v1243
	v1252 = v1245
	goto L319
L337:
	;
	v1262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1257))))
	if v1123&int32(255) == v1262 {
		goto L339
	} else {
		goto L340
	}
L338:
	;
	goto L318
L339:
	;
	v1279 = v1257
	goto L317
L340:
	;
	goto L341
L341:
	;
	v1264 = int32(1)
	v1267 = v1259 - v1264
	if v1267 != 0 {
		v1257 = v1257 + v1264
		v1259 = v1267
		goto L337
	} else {
		goto L342
	}
L342:
	;
	goto L338
L343:
	;
	v1307 = v875
	v1308 = v628
	goto L138
L344:
	;
	goto L345
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v875
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1282 = v604 - v875
	v1283 = F_pg_encoding_mblen_or_incomplete(m, v1281, v875, v1282)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L54
	} else {
		goto L346
	}
L346:
	;
	if v1282 < v1283 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v1287 = v604
	goto L349
L348:
	;
	v1287 = v875 + v1283
	goto L349
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1287
	v1390 = int32(4)
	goto L132
L350:
	;
	if v1296 < v1297 {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	v1301 = v604
	goto L353
L352:
	;
	v1301 = v1293 + v1297
	goto L353
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1301
	v1390 = int32(18)
	goto L132
L354:
	;
	v1344 = v1322
	goto L133
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v604
	v1390 = int32(15)
	goto L132
L356:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1331)+1)))
	if v1332 != 0 {
		goto L355
	} else {
		goto L357
	}
L357:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v1331+int32(4), v1335, v604-v1335)
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L54
	} else {
		goto L358
	}
L358:
	;
	v1390 = int32(1)
	goto L132
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1344
	v1390 = int32(15)
	goto L132
L360:
	;
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1362)+1)))
	if v1363 != 0 {
		goto L359
	} else {
		goto L361
	}
L361:
	;
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v1362+int32(4), v1366, v604-v1366)
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L54
	} else {
		goto L362
	}
L362:
	;
	v1390 = int32(1)
	goto L132
L363:
	;
	v1508 = int32(1)
	goto L103
L364:
	;
	if v1399 != 0 {
		v1604 = v1399
		goto L1
	} else {
		goto L365
	}
L365:
	;
	v1508 = int32(2)
	goto L103
L366:
	;
	if v1404 != 0 {
		v1604 = v1404
		goto L1
	} else {
		goto L367
	}
L367:
	;
	v1508 = int32(2)
	goto L103
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v571
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v63
	v1424 = int32(15)
	switch v571 - v73 - int32(4) {
	case 0:
		goto L374
	case 1:
		goto L373
	default:
		v1604 = v1424
		goto L1
	}
L369:
	;
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1409)+1)))
	if v1410 != 0 {
		goto L368
	} else {
		goto L370
	}
L370:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v571 != v1411+v1412 {
		goto L368
	} else {
		goto L371
	}
L371:
	;
	F_appendBinaryStringInfo(m, v1409+int32(4), v73, v66-v73)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L54
	} else {
		goto L372
	}
L372:
	;
	v1604 = int32(1)
	goto L1
L373:
	;
	v1436 = int32(377998)
	v1437 = int32(5)
	goto L382
L374:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v1428 == int32(1702195828) {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	v1508 = int32(9)
	goto L103
L376:
	;
	goto L377
L377:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v1432 != int32(1819047278) {
		v1604 = v1424
		goto L1
	} else {
		goto L378
	}
L378:
	;
	v1508 = int32(11)
	goto L103
L379:
	;
	if v1499 != 0 {
		v1604 = v1424
		goto L1
	} else {
		goto L397
	}
L380:
	;
	v1499 = int32(0)
	goto L379
L381:
	;
	v1473 = v1468
	v1474 = v1469
	v1475 = v1470
	goto L391
L382:
	;
	if (v73|v1436)&int32(3) != 0 {
		v1468 = v73
		v1469 = v1436
		v1470 = v1437
		goto L381
	} else {
		goto L385
	}
L384:
	;
	if v1458 == int32(0) {
		goto L380
	} else {
		goto L390
	}
L385:
	;
	v1445 = v73
	v1446 = v1436
	v1447 = v1437
	goto L386
L386:
	;
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v1445)))
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1446)))
	if v1450 != v1451 {
		v1468 = v1445
		v1469 = v1446
		v1470 = v1447
		goto L381
	} else {
		goto L388
	}
L387:
	;
	goto L384
L388:
	;
	v1453 = int32(4)
	v1454 = v1446 + v1453
	v1456 = v1445 + v1453
	v1458 = v1447 - v1453
	if base.Ui32(int32(3)) < base.Ui32(v1458) {
		v1445 = v1456
		v1446 = v1454
		v1447 = v1458
		goto L386
	} else {
		goto L389
	}
L389:
	;
	goto L387
L390:
	;
	v1468 = v1456
	v1469 = v1454
	v1470 = v1458
	goto L381
L391:
	;
	v1478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1473))))
	v1479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1474))))
	if v1478 == v1479 {
		goto L393
	} else {
		goto L394
	}
L392:
	;
	v1499 = v1478 - v1479
	goto L379
L393:
	;
	v1481 = int32(1)
	v1486 = v1475 - v1481
	if v1486 != 0 {
		v1473 = v1473 + v1481
		v1474 = v1474 + v1481
		v1475 = v1486
		goto L391
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	goto L392
L396:
	;
	goto L380
L397:
	;
	v1508 = int32(10)
	goto L103
L398:
	;
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v1566
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1568
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1572
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1574
	if v1564 != 0 {
		v1604 = v1564
		goto L1
	} else {
		goto L399
	}
L399:
	;
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v1577 != v1574-v1572 {
		v1604 = int32(15)
		goto L1
	} else {
		goto L400
	}
L400:
	;
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1581 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1580)+2)) = uint8(v1581)
	goto L4
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
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v173 int32
	_ = v173
	var v178 int64
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v261 int32
	_ = v261
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v375 int64
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int64
	_ = v393
	var v394 int64
	_ = v394
	var v395 int32
	_ = v395
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int64
	_ = v409
	var v412 int64
	_ = v412
	var v413 int64
	_ = v413
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	v13 = m.G0
	v15 = v13 - int32(272)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v18 - int32(1) {
	case 0:
		v585 = int32(14)
		goto L1
	default:
		goto L4
	case 6:
		goto L17
	case 10:
		goto L16
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v585
	m.G0 = v15 + int32(272)
	return int32(0)
L2:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v15)+264))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	m.T0[v565].(func(*base.Module, int32, int32, int64, int32, int32, int32))(m, v21, v563, v178, v564, v558, v559)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L23
	} else {
		goto L152
	}
L3:
	;
	v549 = int32(0)
	v558 = v549
	v559 = v549
	goto L2
L4:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(446390)
	m.T0[v542].(func(*base.Module, int32, int32, int32))(m, v541, int32(209304), v15)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L23
	} else {
		goto L151
	}
L5:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v532)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+192)) = int32(552213)
	m.T0[v533].(func(*base.Module, int32, int32, int32))(m, v532, int32(209304), v15+int32(192))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L23
	} else {
		goto L150
	}
L6:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+224)) = int32(552171)
	m.T0[v524].(func(*base.Module, int32, int32, int32))(m, v523, int32(209304), v15+int32(224))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L23
	} else {
		goto L149
	}
L7:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v514)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = int32(235286)
	m.T0[v515].(func(*base.Module, int32, int32, int32))(m, v514, int32(209304), v15+int32(256))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L23
	} else {
		goto L148
	}
L8:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v361)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = int32(552197)
	m.T0[v506].(func(*base.Module, int32, int32, int32))(m, v361, int32(209304), v15+int32(176))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L23
	} else {
		goto L147
	}
L9:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v361)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = int32(552153)
	m.T0[v498].(func(*base.Module, int32, int32, int32))(m, v361, int32(209304), v15+int32(160))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L23
	} else {
		goto L146
	}
L10:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v361)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = int32(390321)
	m.T0[v490].(func(*base.Module, int32, int32, int32))(m, v361, int32(209304), v15+int32(144))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L23
	} else {
		goto L145
	}
L11:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v481)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = int32(235258)
	m.T0[v482].(func(*base.Module, int32, int32, int32))(m, v481, int32(209304), v15+int32(80))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L23
	} else {
		goto L144
	}
L12:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = int32(398879)
	m.T0[v473].(func(*base.Module, int32, int32, int32))(m, v472, int32(209304), v15+int32(96))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L23
	} else {
		goto L143
	}
L13:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = int32(301598)
	m.T0[v452].(func(*base.Module, int32, int32, int32))(m, v21, int32(209304), v15+int32(112))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L23
	} else {
		goto L142
	}
L14:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = int32(357437)
	m.T0[v444].(func(*base.Module, int32, int32, int32))(m, v21, int32(209304), v15+int32(32))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L23
	} else {
		goto L141
	}
L15:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = int32(398604)
	m.T0[v436].(func(*base.Module, int32, int32, int32))(m, v21, int32(209304), v15+int32(128))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L23
	} else {
		goto L140
	}
L16:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v362 == int32(0) {
		goto L10
	} else {
		goto L118
	}
L17:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v23 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v38 == int32(0) {
		goto L14
	} else {
		goto L26
	}
L19:
	;
	if v22 != 0 {
		v37 = v22
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v22 != 0 {
		goto L15
	} else {
		goto L25
	}
L22:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(398586)
	m.T0[v26].(func(*base.Module, int32, int32, int32))(m, v21, int32(209304), v15+int32(16))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	v37 = int32(0)
	goto L18
L26:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v41 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v44 != 0 {
		goto L13
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v37 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v45 = F_strlen(m, v37)
	mBase = m.M
	v47 = base.I32_div_s(v45, int32(2))
	v50 = F_palloc(m, v47+int32(1))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L23
	} else {
		goto L34
	}
L32:
	;
	v173 = v38
	goto L33
L33:
	;
	v178 = F_strtox_2(m, v173, v15+int32(268), int32(10), int64(-1))
	mBase = m.M
	goto L61
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v50
	if v45&int32(1) != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	if int32(2) <= v45 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v60 = int32(0)
	goto L39
L37:
	;
	v151 = v50
	goto L38
L38:
	;
	v153 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v151+v47))) = uint8(v153)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_pfree(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L23
	} else {
		goto L60
	}
L39:
	;
	v73 = v57 + v60<<(uint(int32(1))%32)
	v74 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73))))
	v76 = v74 - int32(48)
	if base.Ui32(v76&int32(255)) <= base.Ui32(int32(9)) {
		v99 = v76
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v151 = v138
	goto L38
L41:
	;
	v100 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73)+1)))
	v102 = v100 - int32(48)
	if base.Ui32(v102&int32(255)) <= base.Ui32(int32(9)) {
		v125 = v102
		goto L49
	} else {
		goto L50
	}
L42:
	;
	if base.Ui32((v74-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v99 = v74 - int32(87)
	goto L41
L44:
	;
	goto L45
L45:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v74-int32(65))&int32(255)) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v98 = int32(-1)
	goto L48
L47:
	;
	v98 = v74 - int32(55)
	goto L48
L48:
	;
	v99 = v98
	goto L41
L49:
	;
	if v99 < int32(0) {
		goto L12
	} else {
		goto L57
	}
L50:
	;
	if base.Ui32((v100-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v125 = v100 - int32(87)
	goto L49
L52:
	;
	goto L53
L53:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v100-int32(65))&int32(255)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v124 = int32(-1)
	goto L56
L55:
	;
	v124 = v100 - int32(55)
	goto L56
L56:
	;
	v125 = v124
	goto L49
L57:
	;
	if v125 < int32(0) {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	v133 = v125 + v99<<(uint(int32(4))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v60+v50))) = uint8(v133)
	v136 = v60 + int32(1)
	if v136 != v47 {
		v60 = v136
		goto L39
	} else {
		goto L59
	}
L59:
	;
	goto L40
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v173 = v160
	goto L33
L61:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v15)+268))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	if v180 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v181 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v242 == int32(0) {
		goto L3
	} else {
		goto L88
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+264)) = int32(0)
	goto L63
L65:
	;
	goto L66
L66:
	;
	v187 = v15 + int32(264)
	v189 = F_pg_strcasecmp(m, v181, int32(389545))
	mBase = m.M
	if v189 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v232 != 0 {
		goto L63
	} else {
		goto L86
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = int32(0)
	v232 = int32(1)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v196 = F_pg_strcasecmp(m, v181, int32(513722))
	mBase = m.M
	if v196 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v199 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = v199
	v232 = v199
	goto L67
L72:
	;
	goto L73
L73:
	;
	v203 = F_pg_strcasecmp(m, v181, int32(584307))
	mBase = m.M
	if v203 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = int32(2)
	v232 = int32(1)
	goto L67
L75:
	;
	goto L76
L76:
	;
	v210 = F_pg_strcasecmp(m, v181, int32(581713))
	mBase = m.M
	if v210 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = int32(3)
	v232 = int32(1)
	goto L67
L78:
	;
	goto L79
L79:
	;
	v217 = F_pg_strcasecmp(m, v181, int32(583948))
	mBase = m.M
	if v217 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = int32(4)
	v232 = int32(1)
	goto L67
L81:
	;
	goto L82
L82:
	;
	v226 = F_pg_strcasecmp(m, v181, int32(586270))
	mBase = m.M
	if v226 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v227 = int32(0)
	goto L85
L84:
	;
	v227 = int32(5)
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = v227
	v232 = base.B2i32(v226 == int32(0))
	goto L67
L86:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v234
	m.T0[v233].(func(*base.Module, int32, int32, int32))(m, v21, int32(758321), v15-int32(-64))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L23
	} else {
		goto L87
	}
L87:
	;
	goto L63
L88:
	;
	v245 = F_strlen(m, v242)
	mBase = m.M
	if v245 == int32(0) {
		goto L3
	} else {
		goto L89
	}
L89:
	;
	v249 = base.I32_div_s(v245, int32(2))
	v250 = F_palloc(m, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L23
	} else {
		goto L90
	}
L90:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v245&int32(1) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if v245 < int32(2) {
		v558 = v249
		v559 = v250
		goto L2
	} else {
		goto L94
	}
L92:
	;
	v349 = v252
	goto L93
L93:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v353
	m.T0[v352].(func(*base.Module, int32, int32, int32))(m, v21, int32(760159), v15+int32(48))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L23
	} else {
		goto L117
	}
L94:
	;
	v261 = int32(0)
	goto L95
L95:
	;
	v274 = v252 + v261<<(uint(int32(1))%32)
	v275 = int32(*(*int8)(unsafe.Add(mBase, uint32(v274))))
	v277 = v275 - int32(48)
	if base.Ui32(v277&int32(255)) <= base.Ui32(int32(9)) {
		v300 = v277
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v349 = v339
	goto L93
L97:
	;
	v301 = int32(*(*int8)(unsafe.Add(mBase, uint32(v274)+1)))
	v303 = v301 - int32(48)
	if base.Ui32(v303&int32(255)) <= base.Ui32(int32(9)) {
		v326 = v303
		goto L105
	} else {
		goto L106
	}
L98:
	;
	if base.Ui32((v275-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v300 = v275 - int32(87)
	goto L97
L100:
	;
	goto L101
L101:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v275-int32(65))&int32(255)) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v299 = int32(-1)
	goto L104
L103:
	;
	v299 = v275 - int32(55)
	goto L104
L104:
	;
	v300 = v299
	goto L97
L105:
	;
	if v300 < int32(0) {
		goto L113
	} else {
		goto L114
	}
L106:
	;
	if base.Ui32((v301-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v326 = v301 - int32(87)
	goto L105
L108:
	;
	goto L109
L109:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v301-int32(65))&int32(255)) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v325 = int32(-1)
	goto L112
L111:
	;
	v325 = v301 - int32(55)
	goto L112
L112:
	;
	v326 = v325
	goto L105
L113:
	;
	goto L96
L114:
	;
	if v326 < int32(0) {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v334 = v326 + v300<<(uint(int32(4))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v261+v250))) = uint8(v334)
	v337 = v261 + int32(1)
	if v337 != v249 {
		v261 = v337
		goto L95
	} else {
		goto L116
	}
L116:
	;
	v558 = v249
	v559 = v250
	goto L2
L117:
	;
	v558 = v249
	v559 = v250
	goto L2
L118:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v365 == int32(0) {
		goto L9
	} else {
		goto L119
	}
L119:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v368 == int32(0) {
		goto L8
	} else {
		goto L120
	}
L120:
	;
	v375 = F_strtox_2(m, v362, v15+int32(260), int32(10), int64(4294967295))
	mBase = m.M
	goto L121
L121:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v15)+260))
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377))))
	if v378 != 0 {
		goto L7
	} else {
		goto L122
	}
L122:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v15 + int32(268)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v15 + int32(264)
	v389 = F_sscanf(m, v379, int32(539886), v15+int32(240))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L23
	} else {
		goto L123
	}
L123:
	;
	if v389 != int32(2) {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	v393 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v15)+264)))
	v394 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v15)+268)))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+208)) = v15 + int32(268)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+212)) = v15 + int32(264)
	v405 = F_sscanf(m, v395, int32(539886), v15+int32(208))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L23
	} else {
		goto L125
	}
L125:
	;
	if v405 != int32(2) {
		goto L5
	} else {
		goto L126
	}
L126:
	;
	v409 = int64(32)
	v412 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v15)+264)))
	v413 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v15)+268)))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v361)+16))
	m.T0[v417].(func(*base.Module, int32, int32, int64, int64))(m, v361, base.I32_wrap_i64(v375), v394<<(uint(v409)%64)|v393, v412|v413<<(uint(v409)%64))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L23
	} else {
		goto L127
	}
L127:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v420 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	F_pfree(m, v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L23
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v425 != 0 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
	goto L130
L132:
	;
	F_pfree(m, v425)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L23
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v430 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
	goto L134
L136:
	;
	F_pfree(m, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L23
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v585 = int32(10)
	goto L1
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	goto L138
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v568 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	F_pfree(m, v568)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L23
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v573 != 0 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	goto L155
L157:
	;
	F_pfree(m, v573)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L23
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v578 != 0 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
	goto L159
L161:
	;
	F_pfree(m, v578)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L23
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v585 = int32(6)
	goto L1
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
	goto L163
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
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
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
	v288 = m.ExcPending
	if v288 != 0 {
		goto L8
	} else {
		goto L88
	}
L2:
	;
	m.G0 = v11 + int32(32)
	return v273
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
	v20 = F_cstring_to_text(m, int32(4103))
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
	v273 = v20
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
	F_errmsg(m, int32(130606), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(518600), int32(1428), int32(117552))
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
	F_errmsg(m, int32(156163), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(518600), int32(1435), int32(117552))
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
	F_errmsg(m, int32(125755), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(518600), int32(1441), int32(117552))
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
	F_initStringInfo(m, v11+int32(16))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	F_appendStringInfoChar(m, v11+int32(16), int32(123))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
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
	v100 = base.I32_div_s(v87, int32(2))
	v106 = int32(0)
	goto L31
L29:
	;
	goto L30
L30:
	;
	F_appendStringInfoChar(m, v11+int32(16), int32(125))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L83
	}
L31:
	;
	v109 = int32(1)
	v110 = v106 << (uint(v109) % 32)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+v111))))
	if v113 == v109 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	goto L30
L33:
	;
	if v106 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_appendStringInfoString(m, v11+int32(16), int32(778193))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L8
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121+v110<<(uint(int32(2))%32))))
	v126 = F_pg_detoast_datum_packed(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L8
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v161 = int32(1)
	if v128&v161 != 0 {
		goto L50
	} else {
		goto L51
	}
L39:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v128 == int32(1) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v131 = int32(4)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
	if v133&int32(254) == int32(2) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v146 = int32(1)
	if v128&v146 != 0 {
		v158 = int32(base.Ui32(v128)>>(uint(v146)%32)) - v146
		goto L38
	} else {
		goto L49
	}
L43:
	;
	v142 = v131
	goto L45
L44:
	;
	v142 = base.B2i32(v133 == int32(18)) << (uint(v131) % 32)
	goto L45
L45:
	;
	if v133 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v145 = v131
	goto L48
L47:
	;
	v145 = v142
	goto L48
L48:
	;
	v158 = v145
	goto L38
L49:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v158 = int32(base.Ui32(v152)>>(uint(int32(2))%32)) - int32(4)
	goto L38
L50:
	;
	v165 = v161
	goto L52
L51:
	;
	v165 = int32(4)
	goto L52
L52:
	;
	F_escape_json_with_len(m, v11+int32(16), v126+v165, v158)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	if v126 != v125 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_pfree(m, v126)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L8
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	F_appendStringInfoString(m, v11+int32(16), int32(777944))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L8
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v177 = int32(1)
	v178 = v110 | v177
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178+v179))))
	if v181 == v177 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v245 = v106 + int32(1)
	if v245 != v100 {
		v106 = v245
		goto L31
	} else {
		goto L82
	}
L60:
	;
	F_appendStringInfoString(m, v11+int32(16), int32(316837))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L8
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189+v178<<(uint(int32(2))%32))))
	v194 = F_pg_detoast_datum_packed(m, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L8
	} else {
		goto L65
	}
L63:
	;
	goto L59
L64:
	;
	v229 = int32(1)
	if v196&v229 != 0 {
		goto L76
	} else {
		goto L77
	}
L65:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v196 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v199 = int32(4)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)))
	if v201&int32(254) == int32(2) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	v214 = int32(1)
	if v196&v214 != 0 {
		v226 = int32(base.Ui32(v196)>>(uint(v214)%32)) - v214
		goto L64
	} else {
		goto L75
	}
L69:
	;
	v210 = v199
	goto L71
L70:
	;
	v210 = base.B2i32(v201 == int32(18)) << (uint(v199) % 32)
	goto L71
L71:
	;
	if v201 == int32(1) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v213 = v199
	goto L74
L73:
	;
	v213 = v210
	goto L74
L74:
	;
	v226 = v213
	goto L64
L75:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v226 = int32(base.Ui32(v220)>>(uint(int32(2))%32)) - int32(4)
	goto L64
L76:
	;
	v233 = v229
	goto L78
L77:
	;
	v233 = int32(4)
	goto L78
L78:
	;
	F_escape_json_with_len(m, v11+int32(16), v194+v233, v226)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	if v194 == v193 {
		goto L59
	} else {
		goto L80
	}
L80:
	;
	F_pfree(m, v194)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
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
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_pfree(m, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L8
	} else {
		goto L84
	}
L84:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	F_pfree(m, v263)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v268 = F_cstring_to_text_with_len(m, v266, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	F_pfree(m, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L8
	} else {
		goto L87
	}
L87:
	;
	v273 = v268
	goto L2
L88:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	F_errmsg(m, int32(21782), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(518600), int32(1457), int32(117552))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = l0 + int32(28)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v21 = l0 + int32(36)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		v23 = F_pg_detoast_datum(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = F_parse_jsonb_index_flags(m, v23)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v12
				v28 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v28
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v28
				*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v10 + int32(8)
				F_iterate_json_values(m, v16, v25, v10+int32(24))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v41 = F_make_tsvector(m, v10+int32(8))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
						if v43 != v16 {
							F_pfree(m, v16)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
								if v47 != v23 {
									F_pfree(m, v23)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(32)
										return v41
									}
								} else {
									m.G0 = v10 + int32(32)
									return v41
								}
							}
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							if v47 != v23 {
								F_pfree(m, v23)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 + int32(32)
									return v41
								}
							} else {
								m.G0 = v10 + int32(32)
								return v41
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
