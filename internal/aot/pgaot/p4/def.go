package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_buildDefItem(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v9 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return v133
L2:
	;
	v123 = F_pstrdup(m, l0)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L8
	} else {
		goto L46
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
	v18 = F_strtol(m, l1, v7+int32(12), int32(10))
	mBase = m.M
	goto L4
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
	v37 = F_strtod(m, l1, v7+int32(12))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L12
	}
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v22 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v23 = F_pstrdup(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v27 = F_makeInteger(m, v18)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v30 = F_makeDefElem(m, v23, v27, int32(-1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v133 = v30
	goto L1
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v40 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v52 = int32(358953)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, _consts[471])))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v56 == int32(0) {
		v75 = v55
		v76 = v56
		goto L21
	} else {
		goto L22
	}
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v42 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v43 = F_pstrdup(m, l0)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v45 = F_pstrdup(m, l1)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v47 = F_makeFloat(m, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v50 = F_makeDefElem(m, v43, v47, int32(-1))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v133 = v50
	goto L1
L20:
	;
	if v76-v75 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	goto L20
L22:
	;
	if v55 != v56 {
		v75 = v55
		v76 = v56
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v60 = l1
	v61 = v52
	goto L24
L24:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	if v65 == int32(0) {
		v75 = v64
		v76 = v65
		goto L21
	} else {
		goto L26
	}
L25:
	;
	v75 = v64
	v76 = v65
	goto L21
L26:
	;
	v68 = int32(1)
	if v64 == v65 {
		v60 = v60 + v68
		v61 = v61 + v68
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v80 = F_pstrdup(m, l0)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v88 = int32(376448)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, _consts[472])))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v92 == int32(0) {
		v111 = v91
		v112 = v92
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v83 = F_makeBoolean(m, int32(1))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	v86 = F_makeDefElem(m, v80, v83, int32(-1))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	v133 = v86
	goto L1
L34:
	;
	if v112-v111 != 0 {
		goto L2
	} else {
		goto L42
	}
L35:
	;
	goto L34
L36:
	;
	if v91 != v92 {
		v111 = v91
		v112 = v92
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v96 = l1
	v97 = v88
	goto L38
L38:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
	if v101 == int32(0) {
		v111 = v100
		v112 = v101
		goto L35
	} else {
		goto L40
	}
L39:
	;
	v111 = v100
	v112 = v101
	goto L35
L40:
	;
	v104 = int32(1)
	if v100 == v101 {
		v96 = v96 + v104
		v97 = v97 + v104
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v114 = F_pstrdup(m, l0)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v117 = F_makeBoolean(m, int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v120 = F_makeDefElem(m, v114, v117, int32(-1))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	v133 = v120
	goto L1
L46:
	;
	v125 = F_pstrdup(m, l1)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	v127 = F_makeString(m, v125)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	v130 = F_makeDefElem(m, v123, v127, int32(-1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	v133 = v130
	goto L1
}
func F_defGetNumeric(m *base.Module, l0 int32) float64 {
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
	var v14 float64
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 float64
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		switch v10 - int32(465) {
		case 0:
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v39 = base.F64_convert_i32_s(v37)
			m.G0 = v7 + int32(32)
			return v39
		case 1:
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v14 = F_atof(m, v13)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return float64(0)
			} else {
				v39 = v14
				m.G0 = v7 + int32(32)
				return v39
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return float64(0)
			} else {
				F_errcode(m, int32(16801924))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return float64(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v25
					F_errmsg(m, int32(361269), v7+int32(16))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return float64(0)
					} else {
						F_errfinish(m, int32(519302), int32(85), int32(509601))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return float64(0)
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
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return float64(0)
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return float64(0)
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v51
				F_errmsg(m, int32(361269), v7)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return float64(0)
				} else {
					F_errfinish(m, int32(519302), int32(74), int32(509601))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return float64(0)
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
func F_defGetQualifiedName(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		if v10 == int32(1) {
			v46 = v9
			m.G0 = v7 + int32(32)
			return v46
		} else {
			if v10 != int32(468) {
				if v10 == int32(68) {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					v46 = v17
					m.G0 = v7 + int32(32)
					return v46
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(16801924))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v27
							F_errmsg(m, int32(397393), v7+int32(16))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(519302), int32(259), int32(397683))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
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
				*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v9
				*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v9
				v44 = F_list_make1_impl(m, int32(1), v7+int32(24))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v46 = v44
					m.G0 = v7 + int32(32)
					return v46
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v58
				F_errmsg(m, int32(226177), v7)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(519302), int32(245), int32(397683))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
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
func F_defGetString(m *base.Module, l0 int32) int32 {
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		switch v10 - int32(465) {
		case 0:
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v47
			v52 = F_psprintf(m, int32(449224), v7+int32(32))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				v73 = v52
				m.G0 = v7 + int32(48)
				return v73
			}
		case 1:
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v73 = v15
			m.G0 = v7 + int32(48)
			return v73
		case 2:
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)))
			if v18 != 0 {
				v19 = int32(358953)
			} else {
				v19 = int32(376448)
			}
			v73 = v19
			m.G0 = v7 + int32(48)
			return v73
		case 3:
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v73 = v20
			m.G0 = v7 + int32(48)
			return v73
		default:
			switch v10 - int32(68) {
			case 0:
				v21 = F_TypeNameToString(m, v9)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v73 = v21
					m.G0 = v7 + int32(48)
					return v73
				}
			case 1, 2, 3, 4, 5, 6, 7, 8:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v35
					F_errmsg_internal(m, int32(504634), v7+int32(16))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(519302), int32(59), int32(344700))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			case 9:
				v26 = F_pstrdup(m, int32(694900))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v73 = v26
					m.G0 = v7 + int32(48)
					return v73
				}
			default:
				if v10 == int32(1) {
					v71 = F_NameListToString(m, v9)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						v73 = v71
						m.G0 = v7 + int32(48)
						return v73
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v35
						F_errmsg_internal(m, int32(504634), v7+int32(16))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(519302), int32(59), int32(344700))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
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
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v61
				F_errmsg(m, int32(226177), v7)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(519302), int32(41), int32(344700))
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
	}
}
