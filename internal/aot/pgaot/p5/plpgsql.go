package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_plpgsql_build_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	switch v12 {
	case 0:
		v14 = F_palloc0(m, int32(52))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(0)
			v20 = F_pstrdup(m, l0)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(v14)+44)) = uint16(v22)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v20
				v32 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
				v35 = *(*int32)(unsafe.Add(mBase, _consts[1351]))
				v38 = *(*int32)(unsafe.Add(mBase, _consts[1352]))
				if v35 == v38 {
					*(*int32)(unsafe.Add(mBase, _consts[1352])) = v35 << (uint(int32(1)) % 32)
					v50 = F_repalloc(m, v32, v35<<(uint(int32(3))%32))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[1350])) = v50
						v55 = *(*int32)(unsafe.Add(mBase, _consts[1351]))
						v56 = v55
						v57 = v50
						*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v56
						*(*int32)(unsafe.Add(mBase, _consts[1351])) = v56 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v57+v56<<(uint(int32(2))%32)))) = v14
						if l3 == int32(0) {
							v133 = v14
							m.G0 = v10 + int32(32)
							return v133
						} else {
							F_plpgsql_ns_additem(m, int32(1), v56, l0)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v133 = v14
								m.G0 = v10 + int32(32)
								return v133
							}
						}
					}
				} else {
					v56 = v35
					v57 = v32
					*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v56
					*(*int32)(unsafe.Add(mBase, _consts[1351])) = v56 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v57+v56<<(uint(int32(2))%32)))) = v14
					if l3 == int32(0) {
						v133 = v14
						m.G0 = v10 + int32(32)
						return v133
					} else {
						F_plpgsql_ns_additem(m, int32(1), v56, l0)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							v133 = v14
							m.G0 = v10 + int32(32)
							return v133
						}
					}
				}
			}
		}
	case 1:
		v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v74 = F_palloc0(m, int32(40))
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(2)
			v78 = F_pstrdup(m, l0)
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v74)+32)) = int64(4294967295)
				*(*int32)(unsafe.Add(mBase, uint32(v74)+28)) = v72
				*(*int32)(unsafe.Add(mBase, uint32(v74)+24)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = v78
				v89 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
				v92 = *(*int32)(unsafe.Add(mBase, _consts[1351]))
				v95 = *(*int32)(unsafe.Add(mBase, _consts[1352]))
				if v92 == v95 {
					*(*int32)(unsafe.Add(mBase, _consts[1352])) = v92 << (uint(int32(1)) % 32)
					v107 = F_repalloc(m, v89, v92<<(uint(int32(3))%32))
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[1350])) = v107
						v112 = *(*int32)(unsafe.Add(mBase, _consts[1351]))
						v113 = v107
						v114 = v112
						*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v114
						*(*int32)(unsafe.Add(mBase, _consts[1351])) = v114 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v113+v114<<(uint(int32(2))%32)))) = v74
						if l3 == int32(0) {
							v133 = v74
							m.G0 = v10 + int32(32)
							return v133
						} else {
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
							F_plpgsql_ns_additem(m, int32(2), v114, v127)
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int32(0)
							} else {
								v133 = v74
								m.G0 = v10 + int32(32)
								return v133
							}
						}
					}
				} else {
					v113 = v89
					v114 = v92
					*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v114
					*(*int32)(unsafe.Add(mBase, _consts[1351])) = v114 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v113+v114<<(uint(int32(2))%32)))) = v74
					if l3 == int32(0) {
						v133 = v74
						m.G0 = v10 + int32(32)
						return v133
					} else {
						v127 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
						F_plpgsql_ns_additem(m, int32(2), v114, v127)
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int32(0)
						} else {
							v133 = v74
							m.G0 = v10 + int32(32)
							return v133
						}
					}
				}
			}
		}
	case 2:
		F_errstart_cold(m, int32(21), int32(569208))
		mBase = m.M
		v144 = m.ExcPending
		if v144 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v147 = m.ExcPending
			if v147 != 0 {
				return int32(0)
			} else {
				v148 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				v149 = F_format_type_be(m, v148)
				mBase = m.M
				v150 = m.ExcPending
				if v150 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v149
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
					F_errmsg(m, int32(192378), v10+int32(16))
					mBase = m.M
					v158 = m.ExcPending
					if v158 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(507433), int32(1795), int32(404898))
						mBase = m.M
						v165 = m.ExcPending
						if v165 != 0 {
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
	default:
		F_errstart_cold(m, int32(21), int32(569208))
		mBase = m.M
		v171 = m.ExcPending
		if v171 != 0 {
			return int32(0)
		} else {
			v172 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v172
			F_errmsg_internal(m, int32(494259), v10)
			mBase = m.M
			v177 = m.ExcPending
			if v177 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(507433), int32(1799), int32(404898))
				mBase = m.M
				v184 = m.ExcPending
				if v184 != 0 {
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
func F_plpgsql_location_to_lineno(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	v3 = int32(0)
	if l0 < v3 {
		v58 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v58
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
	if v9 == int32(0) {
		v58 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = l0 + v9
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+188))
	if base.Ui32(v13) <= base.Ui32(v12) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+196))
	if v28 == int32(0) {
		v58 = v29
		goto L1
	} else {
		goto L12
	}
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+192))
	v28 = v15
	goto L4
L6:
	;
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+196)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v9
	v19 = int32(10)
	v20 = F___strchrnul(m, v9, v19)
	mBase = m.M
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v22 == v19 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+192)) = v26
	v28 = v26
	goto L4
L9:
	;
	v26 = v20
	goto L11
L10:
	;
	v26 = int32(0)
	goto L11
L11:
	;
	goto L8
L12:
	;
	if base.Ui32(v12) <= base.Ui32(v28) {
		v58 = v29
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v33 = v28
	v35 = v29
	goto L14
L14:
	;
	v38 = int32(1)
	v39 = v35 + v38
	*(*int32)(unsafe.Add(mBase, uint32(v8)+196)) = v39
	v42 = v33 + v38
	*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v42
	v44 = int32(10)
	v45 = F___strchrnul(m, v42, v44)
	mBase = m.M
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v47 == v44 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v58 = v39
	goto L1
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+192)) = v51
	if v51 == int32(0) {
		v58 = v39
		goto L1
	} else {
		goto L20
	}
L17:
	;
	v51 = v45
	goto L19
L18:
	;
	v51 = int32(0)
	goto L19
L19:
	;
	goto L16
L20:
	;
	if base.Ui32(v51) < base.Ui32(v12) {
		v33 = v51
		v35 = v39
		goto L14
	} else {
		goto L21
	}
L21:
	;
	goto L15
}
func F_plpgsql_ns_find_nearest_loop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v2 = l0
	goto L4
L2:
	;
	goto L3
L3:
	;
	return int32(0)
L4:
	;
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	if v3 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v2)+8))
	if v8 != 0 {
		v2 = v8
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	if v4 != int32(1) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	return v2
L9:
	;
	goto L5
}
func F_plpgsql_ns_lookup_label(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v3 = l0
	goto L4
L2:
	;
	goto L3
L3:
	;
	return int32(0)
L4:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	if v5 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	if v34 != 0 {
		v3 = v34
		goto L4
	} else {
		goto L17
	}
L7:
	;
	v7 = v3 + int32(12)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v11 == int32(0) {
		v30 = v10
		v31 = v11
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v31-v30 != 0 {
		goto L6
	} else {
		goto L16
	}
L9:
	;
	goto L8
L10:
	;
	if v10 != v11 {
		v30 = v10
		v31 = v11
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v15 = v7
	v16 = l1
	goto L12
L12:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v20 == int32(0) {
		v30 = v19
		v31 = v20
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v30 = v19
	v31 = v20
	goto L9
L14:
	;
	v23 = int32(1)
	if v19 == v20 {
		v15 = v15 + v23
		v16 = v16 + v23
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	return v3
L17:
	;
	goto L5
}
func F_plpgsql_param_eval_recfield(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
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
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v15 = int32(2)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14<<(uint(v15)%32)-int32(4))))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v13+v21<<(uint(v15)%32))))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if v26 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_instantiate_empty_record_variable(m, v12, v25)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v32 = v26
	goto L3
L3:
	;
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v20)+24))
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v32)+48))
	if v33 != v34 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	return
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	v32 = v31
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(569208))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L24
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(569208))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L20
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v39 = F_expanded_record_lookup_field(m, v32, v36, v20+int32(32))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	if v46 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	if v39 == int32(0) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v32)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = v43
	goto L10
L13:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v73 != v74 {
		goto L6
	} else {
		goto L19
	}
L14:
	;
	v67 = F_expanded_record_fetch_field(m, v32, v46, v45)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L18
	}
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	if v49&int32(4) == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v32)+64))
	if v54 < v46 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v57 = v46 - int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+v58))))
	*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v60)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v32)+56))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62+v57<<(uint(int32(2))%32))))
	v70 = v66
	goto L13
L18:
	;
	v70 = v67
	goto L13
L19:
	;
	m.G0 = v9 + int32(32)
	return
L20:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v88
	F_errmsg(m, int32(737066), v9+int32(16))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(512305), int32(6787), int32(441097))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	v116 = F_format_type_be(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v119 = F_format_type_be(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v114
	F_errmsg(m, int32(688925), v9)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(512305), int32(6803), int32(441097))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_plpgsql_parse_cwordrowtype(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(4536272)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v14 = *(*int32)(unsafe.Add(mBase, _consts[1355]))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v14
	v16 = F_makeRangeVarFromNameList(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(0)
		v24 = F_RangeVarGetRelidExtended(m, v16, v20, v20, v20, v20)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = F_get_rel_type_id(m, v24)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 != 0 {
					*(*int32)(unsafe.Add(mBase, _consts[10])) = v11
					v30 = F_makeTypeNameFromNameList(m, l0)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v33 = F_SearchSysCache1(m, int32(82), v26)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							if v33 == int32(0) {
								F_errstart_cold(m, int32(21), int32(569208))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v26
									F_errmsg_internal(m, int32(51355), v8+int32(16))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(507433), int32(1960), int32(374469))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v39 = F_build_datatype(m, v33, int32(-1), int32(0), v30)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									F_ReleaseCatCache(m, v33)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										m.G0 = v8 + int32(32)
										return v39
									}
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(569208))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v54
							F_errmsg(m, int32(378289), v8)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(507433), int32(1729), int32(373837))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
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
}
func F_plpgsql_parse_cwordtype(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
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
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(4536272)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v14 = *(*int32)(unsafe.Add(mBase, _consts[1355]))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v14
	if l0 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(569208))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L52
	} else {
		goto L69
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(569208))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L52
	} else {
		goto L65
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v11
	m.G0 = v8 + int32(32)
	return v262
L4:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	v227 = int32(0)
	v231 = F_RangeVarGetRelidExtended(m, v222, v227, v227, v227, v227)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L52
	} else {
		goto L57
	}
L5:
	;
	v209 = F_list_copy(m, l0)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L52
	} else {
		goto L54
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v18 != int32(2) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[1354]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v31 = v8 + int32(28)
	if v22 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	v202 = F_makeRangeVar(m, int32(0), v200, int32(-1))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L52
	} else {
		goto L53
	}
L9:
	;
	if v169 == int32(0) {
		goto L8
	} else {
		goto L46
	}
L10:
	;
	v169 = v159
	goto L9
L11:
	;
	v159 = v145
	goto L10
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v137
	v145 = v135
	goto L11
L13:
	;
	v127 = int32(0)
	if v31 == v127 {
		v159 = v127
		goto L10
	} else {
		goto L45
	}
L14:
	;
	v39 = v22
	goto L15
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v48 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L13
L17:
	;
	v53 = v39
	v55 = v48
	goto L20
L18:
	;
	v75 = v39
	goto L19
L19:
	;
	if v28 == int32(0) {
		v112 = v75
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v60 = F_strcmp(m, v53+int32(12), v26)
	mBase = m.M
	if v60 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v75 = v69
	goto L19
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v70 != 0 {
		v53 = v69
		v55 = v70
		goto L20
	} else {
		goto L26
	}
L23:
	;
	if base.B2i32(v55 == int32(1))&base.B2i32(v28 != int32(0)) != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if v31 == int32(0) {
		v145 = v53
		goto L11
	} else {
		goto L25
	}
L25:
	;
	v135 = v53
	v137 = int32(1)
	goto L12
L26:
	;
	goto L21
L27:
	;
	goto L43
L28:
	;
	v84 = F_strcmp(m, v75+int32(12), v26)
	mBase = m.M
	if v84 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v85 = v75
	goto L31
L30:
	;
	v85 = v39
	goto L31
L31:
	;
	if v84 != 0 {
		v112 = v85
		goto L27
	} else {
		goto L32
	}
L32:
	;
	if v48 == int32(0) {
		v112 = v85
		goto L27
	} else {
		goto L33
	}
L33:
	;
	v88 = v39
	v95 = v48
	goto L34
L34:
	;
	v99 = F_strcmp(m, v88+int32(12), v28)
	mBase = m.M
	if v99 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v112 = v106
	goto L27
L36:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	if v107 != 0 {
		v88 = v106
		v95 = v107
		goto L34
	} else {
		goto L42
	}
L37:
	;
	if int32(0)&base.B2i32(v95 == int32(1)) != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	if v31 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v169 = v88
	goto L9
L40:
	;
	goto L41
L41:
	;
	v135 = v88
	v137 = int32(2)
	goto L12
L42:
	;
	goto L35
L43:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	if v117 != 0 {
		v39 = v117
		goto L15
	} else {
		goto L44
	}
L44:
	;
	goto L16
L45:
	;
	v135 = v127
	v137 = v127
	goto L12
L46:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	if v172 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v176+v177<<(uint(int32(2))%32))))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+24))
	v262 = v182
	goto L3
L48:
	;
	goto L49
L49:
	;
	if v172 != int32(2) {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	if v185 != int32(2) {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v189+v190<<(uint(int32(2))%32))))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+24))
	v262 = v195
	goto L3
L52:
	;
	return int32(0)
L53:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v222 = v202
	v224 = v206 + int32(4)
	goto L4
L54:
	;
	v211 = F_list_delete_last(m, v209)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v213 = F_makeRangeVarFromNameList(m, v211)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v222 = v213
	v224 = v215 + v216<<(uint(int32(2))%32) - int32(4)
	goto L4
L57:
	;
	v233 = F_SearchSysCacheAttName(m, v231, v226)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L52
	} else {
		goto L58
	}
L58:
	;
	if v233 == int32(0) {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v233)+16))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+22)))
	v240 = v238 + v239
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+68))
	v242 = F_SearchSysCache1(m, int32(82), v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L52
	} else {
		goto L60
	}
L60:
	;
	if v242 == int32(0) {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v11
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v240)+76))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v240)+96))
	v251 = F_build_datatype(m, v242, v248, v249, int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L52
	} else {
		goto L62
	}
L62:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _consts[1355]))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v255
	F_ReleaseCatCache(m, v233)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L52
	} else {
		goto L63
	}
L63:
	;
	F_ReleaseCatCache(m, v242)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L52
	} else {
		goto L64
	}
L64:
	;
	v262 = v251
	goto L3
L65:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L52
	} else {
		goto L66
	}
L66:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v226
	F_errmsg(m, int32(73052), v8)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L52
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(507433), int32(1631), int32(374348))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L52
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v240)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v292
	F_errmsg_internal(m, int32(51355), v8+int32(16))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L52
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(507433), int32(1637), int32(374348))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L52
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_plpgsql_parse_tripword(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[1353]))
	if v13 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(80)
	return v255
L2:
	;
	v228 = F_makeString(m, l0)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L47
	} else {
		goto L57
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[1354]))
	v20 = v10 + int32(76)
	if v17 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	if v158 == int32(0) {
		goto L2
	} else {
		goto L41
	}
L5:
	;
	v158 = v148
	goto L4
L6:
	;
	v148 = v134
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v126
	v134 = v124
	goto L6
L8:
	;
	v116 = int32(0)
	if v20 == v116 {
		v148 = v116
		goto L5
	} else {
		goto L40
	}
L9:
	;
	v28 = v17
	goto L10
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L8
L12:
	;
	v42 = v28
	v44 = v37
	goto L15
L13:
	;
	v64 = v28
	goto L14
L14:
	;
	if l1 == int32(0) {
		v101 = v64
		goto L22
	} else {
		goto L23
	}
L15:
	;
	v49 = F_strcmp(m, v42+int32(12), l0)
	mBase = m.M
	if v49 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v64 = v58
	goto L14
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v59 != 0 {
		v42 = v58
		v44 = v59
		goto L15
	} else {
		goto L21
	}
L18:
	;
	if base.B2i32(v44 == int32(1))&base.B2i32(l1 != int32(0)) != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	if v20 == int32(0) {
		v134 = v42
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v124 = v42
	v126 = int32(1)
	goto L7
L21:
	;
	goto L16
L22:
	;
	goto L38
L23:
	;
	v73 = F_strcmp(m, v64+int32(12), l0)
	mBase = m.M
	if v73 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v74 = v64
	goto L26
L25:
	;
	v74 = v28
	goto L26
L26:
	;
	if v73 != 0 {
		v101 = v74
		goto L22
	} else {
		goto L27
	}
L27:
	;
	if v37 == int32(0) {
		v101 = v74
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v77 = v28
	v84 = v37
	goto L29
L29:
	;
	v88 = F_strcmp(m, v77+int32(12), l1)
	mBase = m.M
	if v88 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v101 = v95
	goto L22
L31:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v96 != 0 {
		v77 = v95
		v84 = v96
		goto L29
	} else {
		goto L37
	}
L32:
	;
	if base.B2i32(l2 != int32(0))&base.B2i32(v84 == int32(1)) != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	if v20 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v158 = v77
	goto L4
L35:
	;
	goto L36
L36:
	;
	v124 = v77
	v126 = int32(2)
	goto L7
L37:
	;
	goto L30
L38:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
	if v106 != 0 {
		v28 = v106
		goto L10
	} else {
		goto L39
	}
L39:
	;
	goto L11
L40:
	;
	v124 = v116
	v126 = v116
	goto L7
L41:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v161 != int32(2) {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _consts[1350]))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v165+v166<<(uint(int32(2))%32))))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
	if v171 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v219
	v221 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v221)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v221
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v218
	v255 = int32(1)
	goto L1
L44:
	;
	v174 = F_plpgsql_build_recfield(m, v170, l1)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v193 = F_plpgsql_build_recfield(m, v170, l2)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L47
	} else {
		goto L52
	}
L47:
	;
	return int32(0)
L48:
	;
	v178 = F_makeString(m, l0)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v178
	v181 = F_makeString(m, l1)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v181
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v10)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v185
	v191 = F_list_make2_impl(m, v10+int32(28), v10+int32(24))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L47
	} else {
		goto L51
	}
L51:
	;
	v218 = v174
	v219 = v191
	goto L43
L52:
	;
	v195 = F_makeString(m, l0)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L47
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v195
	v198 = F_makeString(m, l1)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L47
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v198
	v201 = F_makeString(m, l2)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L47
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v201
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v205
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v207
	v215 = F_list_make3_impl(m, v10+int32(40), v10+int32(36), v10+int32(32))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L47
	} else {
		goto L56
	}
L56:
	;
	v218 = v193
	v219 = v215
	goto L43
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v228
	v231 = F_makeString(m, l1)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L47
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v231
	v234 = F_makeString(m, l2)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L47
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v234
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v240
	v248 = F_list_make3_impl(m, v10+int32(20), v10+int32(16), v10+int32(12))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L47
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v248
	v255 = int32(0)
	goto L1
}
func F_plpgsql_parser_setup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(7197)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(7198)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(7199)
	return
}
func F_plpgsql_peek2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v15 = F_internal_yylex(m, v11+int32(24), l4)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v17 = F_internal_yylex(m, v11, l4)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v15
			if l2 != 0 {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v20
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
			if l3 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v23
			} else {
			}
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+72))
			if v26 < int32(4) {
				*(*int32)(unsafe.Add(mBase, uint32(v25+v26<<(uint(int32(2))%32))+76)) = v17
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+72))
				v36 = v25 + v33*int32(24)
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v36)+108)) = v37
				v39 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v36)+100)) = v39
				v41 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
				*(*int64)(unsafe.Add(mBase, uint32(v36)+92)) = v41
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+72))
				v46 = v44 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v43)+72)) = v46
				if int32(3) <= v44 {
					F_errstart_cold(m, int32(21), int32(569208))
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(326119), int32(0))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return
						} else {
							F_errfinish(m, int32(506894), int32(388), int32(287934))
							mBase = m.M
							v107 = m.ExcPending
							if v107 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v43+v46<<(uint(int32(2))%32))+76)) = v15
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v43)+72))
					v57 = v43 + v54*int32(24)
					v58 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
					*(*int64)(unsafe.Add(mBase, uint32(v57)+108)) = v58
					v60 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
					*(*int64)(unsafe.Add(mBase, uint32(v57)+100)) = v60
					v62 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v57)+92)) = v62
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
					*(*int32)(unsafe.Add(mBase, uint32(v64)+72)) = v65 + int32(1)
					m.G0 = v11 + int32(48)
					return
				}
			} else {
				F_errstart_cold(m, int32(21), int32(569208))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(326119), int32(0))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return
					} else {
						F_errfinish(m, int32(506894), int32(388), int32(287934))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
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
}
func F_plpgsql_scanner_errposition(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	v3 = int32(0)
	if l0 < v3 {
		v23 = v3
		return v23
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+60))
		if v8 == int32(0) {
			v23 = v3
			return v23
		} else {
			v11 = F_pg_mbstrlen_with_len(m, v8, l0)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_internalerrposition(m, v11+int32(1))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
					v21 = F_internalerrquery(m, v20)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v23 = v21
						return v23
					}
				}
			}
		}
	}
}
func F_plpgsql_yylex(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
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
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v144 int64
	_ = v144
	var v146 int64
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v186 int64
	_ = v186
	var v188 int64
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v208 int64
	_ = v208
	var v210 int64
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
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
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int64
	_ = v264
	var v266 int64
	_ = v266
	var v268 int64
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int64
	_ = v336
	var v338 int64
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	v10 = m.G0
	v12 = v10 - int32(128)
	m.G0 = v12
	v16 = F_internal_yylex(m, v12+int32(104), l2)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		switch v16 - int32(258) {
		case 0, 9:
			v24 = F_internal_yylex(m, v12+int32(80), l2)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v24 == int32(46) {
					v30 = F_internal_yylex(m, v12+int32(56), l2)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 == int32(258) {
							v36 = F_internal_yylex(m, v12+int32(32), l2)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								if v36 == int32(46) {
									v42 = F_internal_yylex(m, v12+int32(8), l2)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return int32(0)
									} else {
										if v42 == int32(258) {
											v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
											v51 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
											v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
											v54 = v12 + int32(104)
											v57 = F_plpgsql_parse_tripword(m, v50, v51, v52, v54, v54)
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return int32(0)
											} else {
												v117 = v12 + int32(28)
												v118 = v12 + int32(24)
												v120 = v57
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
												v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
												*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = v121 + (v122 - v123)
												if v120 != 0 {
													v129 = int32(277)
												} else {
													v129 = int32(276)
												}
												v331 = v129
												v336 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
												*(*int64)(unsafe.Add(mBase, uint32(l0))) = v336
												v338 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v338
												v340 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
												*(*int32)(unsafe.Add(mBase, uint32(l1))) = v340
												v342 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
												v343 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
												*(*int32)(unsafe.Add(mBase, uint32(v343)+68)) = v331
												*(*int32)(unsafe.Add(mBase, uint32(v343)+64)) = v342
												m.G0 = v12 + int32(128)
												return v331
											}
										} else {
											v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+72))
											if int32(4) <= v60 {
												F_errstart_cold(m, int32(21), int32(569208))
												mBase = m.M
												v355 = m.ExcPending
												if v355 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(326119), int32(0))
													mBase = m.M
													v360 = m.ExcPending
													if v360 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(506894), int32(388), int32(287934))
														mBase = m.M
														v367 = m.ExcPending
														if v367 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v59+v60<<(uint(int32(2))%32))+76)) = v42
												v67 = *(*int32)(unsafe.Add(mBase, uint32(v59)+72))
												v70 = v59 + v67*int32(24)
												v71 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
												*(*int64)(unsafe.Add(mBase, uint32(v70)+108)) = v71
												v73 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
												*(*int64)(unsafe.Add(mBase, uint32(v70)+100)) = v73
												v75 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
												*(*int64)(unsafe.Add(mBase, uint32(v70)+92)) = v75
												v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
												v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+72))
												v80 = v78 + int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v77)+72)) = v80
												if int32(3) <= v78 {
													F_errstart_cold(m, int32(21), int32(569208))
													mBase = m.M
													v373 = m.ExcPending
													if v373 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(326119), int32(0))
														mBase = m.M
														v378 = m.ExcPending
														if v378 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(506894), int32(388), int32(287934))
															mBase = m.M
															v385 = m.ExcPending
															if v385 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v77+v80<<(uint(int32(2))%32))+76)) = int32(46)
													v93 = *(*int32)(unsafe.Add(mBase, uint32(v77)+72))
													v96 = v77 + v93*int32(24)
													v97 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
													*(*int64)(unsafe.Add(mBase, uint32(v96)+108)) = v97
													v99 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
													*(*int64)(unsafe.Add(mBase, uint32(v96)+100)) = v99
													v101 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
													*(*int64)(unsafe.Add(mBase, uint32(v96)+92)) = v101
													v103 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
													v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+72))
													*(*int32)(unsafe.Add(mBase, uint32(v103)+72)) = v104 + int32(1)
													v108 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
													v109 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
													v111 = v12 + int32(104)
													v114 = F_plpgsql_parse_dblword(m, v108, v109, v111, v111)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														v117 = v12 + int32(76)
														v118 = v12 + int32(72)
														v120 = v114
														v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
														v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
														v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
														*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = v121 + (v122 - v123)
														if v120 != 0 {
															v129 = int32(277)
														} else {
															v129 = int32(276)
														}
														v331 = v129
														v336 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
														*(*int64)(unsafe.Add(mBase, uint32(l0))) = v336
														v338 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
														*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v338
														v340 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
														*(*int32)(unsafe.Add(mBase, uint32(l1))) = v340
														v342 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
														v343 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
														*(*int32)(unsafe.Add(mBase, uint32(v343)+68)) = v331
														*(*int32)(unsafe.Add(mBase, uint32(v343)+64)) = v342
														m.G0 = v12 + int32(128)
														return v331
													}
												}
											}
										}
									}
								} else {
									v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+72))
									if int32(4) <= v131 {
										F_errstart_cold(m, int32(21), int32(569208))
										mBase = m.M
										v391 = m.ExcPending
										if v391 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(326119), int32(0))
											mBase = m.M
											v396 = m.ExcPending
											if v396 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(506894), int32(388), int32(287934))
												mBase = m.M
												v403 = m.ExcPending
												if v403 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v130+v131<<(uint(int32(2))%32))+76)) = v36
										v138 = *(*int32)(unsafe.Add(mBase, uint32(v130)+72))
										v141 = v130 + v138*int32(24)
										v142 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v141)+108)) = v142
										v144 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v141)+100)) = v144
										v146 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v141)+92)) = v146
										v148 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+72))
										*(*int32)(unsafe.Add(mBase, uint32(v148)+72)) = v149 + int32(1)
										v153 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
										v154 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
										v156 = v12 + int32(104)
										v159 = F_plpgsql_parse_dblword(m, v153, v154, v156, v156)
										mBase = m.M
										v160 = m.ExcPending
										if v160 != 0 {
											return int32(0)
										} else {
											v161 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
											v162 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
											v163 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
											*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = v161 + (v162 - v163)
											if v159 != 0 {
												v169 = int32(277)
											} else {
												v169 = int32(276)
											}
											v331 = v169
											v336 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
											*(*int64)(unsafe.Add(mBase, uint32(l0))) = v336
											v338 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
											*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v338
											v340 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
											*(*int32)(unsafe.Add(mBase, uint32(l1))) = v340
											v342 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
											v343 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											*(*int32)(unsafe.Add(mBase, uint32(v343)+68)) = v331
											*(*int32)(unsafe.Add(mBase, uint32(v343)+64)) = v342
											m.G0 = v12 + int32(128)
											return v331
										}
									}
								}
							}
						} else {
							v170 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+72))
							if int32(4) <= v171 {
								F_errstart_cold(m, int32(21), int32(569208))
								mBase = m.M
								v409 = m.ExcPending
								if v409 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(326119), int32(0))
									mBase = m.M
									v414 = m.ExcPending
									if v414 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(506894), int32(388), int32(287934))
										mBase = m.M
										v421 = m.ExcPending
										if v421 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v170+v171<<(uint(int32(2))%32))+76)) = v30
								v178 = *(*int32)(unsafe.Add(mBase, uint32(v170)+72))
								v181 = v170 + v178*int32(24)
								v182 = *(*int64)(unsafe.Add(mBase, uint32(v12)+72))
								*(*int64)(unsafe.Add(mBase, uint32(v181)+108)) = v182
								v186 = *(*int64)(unsafe.Add(mBase, uint32(v12-int32(-64))))
								*(*int64)(unsafe.Add(mBase, uint32(v181)+100)) = v186
								v188 = *(*int64)(unsafe.Add(mBase, uint32(v12)+56))
								*(*int64)(unsafe.Add(mBase, uint32(v181)+92)) = v188
								v190 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+72))
								v193 = v191 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v190)+72)) = v193
								if int32(3) <= v191 {
									F_errstart_cold(m, int32(21), int32(569208))
									mBase = m.M
									v427 = m.ExcPending
									if v427 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(326119), int32(0))
										mBase = m.M
										v432 = m.ExcPending
										if v432 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(506894), int32(388), int32(287934))
											mBase = m.M
											v439 = m.ExcPending
											if v439 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v190+v193<<(uint(int32(2))%32))+76)) = int32(46)
									v202 = *(*int32)(unsafe.Add(mBase, uint32(v190)+72))
									v205 = v190 + v202*int32(24)
									v206 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
									*(*int64)(unsafe.Add(mBase, uint32(v205)+108)) = v206
									v208 = *(*int64)(unsafe.Add(mBase, uint32(v12)+88))
									*(*int64)(unsafe.Add(mBase, uint32(v205)+100)) = v208
									v210 = *(*int64)(unsafe.Add(mBase, uint32(v12)+80))
									*(*int64)(unsafe.Add(mBase, uint32(v205)+92)) = v210
									v212 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+72))
									v214 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v212)+72)) = v213 + v214
									v218 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
									v219 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
									v220 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
									v224 = v12 + int32(104)
									v227 = F_plpgsql_parse_word(m, v218, v219+v220, v214, v224, v224)
									mBase = m.M
									v228 = m.ExcPending
									if v228 != 0 {
										return int32(0)
									} else {
										if v227 != 0 {
											v331 = int32(277)
											v336 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
											*(*int64)(unsafe.Add(mBase, uint32(l0))) = v336
											v338 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
											*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v338
											v340 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
											*(*int32)(unsafe.Add(mBase, uint32(l1))) = v340
											v342 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
											v343 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											*(*int32)(unsafe.Add(mBase, uint32(v343)+68)) = v331
											*(*int32)(unsafe.Add(mBase, uint32(v343)+64)) = v342
											m.G0 = v12 + int32(128)
											return v331
										} else {
											v229 = int32(275)
											v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+108)))
											if v230 != 0 {
												v331 = v229
												v336 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
												*(*int64)(unsafe.Add(mBase, uint32(l0))) = v336
												v338 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v338
												v340 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
												*(*int32)(unsafe.Add(mBase, uint32(l1))) = v340
												v342 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
												v343 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
												*(*int32)(unsafe.Add(mBase, uint32(v343)+68)) = v331
												*(*int32)(unsafe.Add(mBase, uint32(v343)+64)) = v342
												m.G0 = v12 + int32(128)
												return v331
											} else {
												v231 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
												v233 = F_ScanKeywordLookup(m, v231, int32(4196960))
												mBase = m.M
												v234 = m.ExcPending
												if v234 != 0 {
													return int32(0)
												} else {
													if v233 < int32(0) {
														v331 = v229
													} else {
														v241 = v233 << (uint(int32(1)) % 32)
														v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241)+uint32(_consts[1434]))))
														*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = int32(2191008) + v245
														v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241)+uint32(_consts[1377]))))
														v331 = v251
													}
													v336 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
													*(*int64)(unsafe.Add(mBase, uint32(l0))) = v336
													v338 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v338
													v340 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
													*(*int32)(unsafe.Add(mBase, uint32(l1))) = v340
													v342 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
													v343 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
													*(*int32)(unsafe.Add(mBase, uint32(v343)+68)) = v331
													*(*int32)(unsafe.Add(mBase, uint32(v343)+64)) = v342
													m.G0 = v12 + int32(128)
													return v331
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v252 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+72))
					if int32(4) <= v253 {
						F_errstart_cold(m, int32(21), int32(569208))
						mBase = m.M
						v445 = m.ExcPending
						if v445 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(326119), int32(0))
							mBase = m.M
							v450 = m.ExcPending
							if v450 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(506894), int32(388), int32(287934))
								mBase = m.M
								v457 = m.ExcPending
								if v457 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v252+v253<<(uint(int32(2))%32))+76)) = v24
						v260 = *(*int32)(unsafe.Add(mBase, uint32(v252)+72))
						v263 = v252 + v260*int32(24)
						v264 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
						*(*int64)(unsafe.Add(mBase, uint32(v263)+108)) = v264
						v266 = *(*int64)(unsafe.Add(mBase, uint32(v12)+88))
						*(*int64)(unsafe.Add(mBase, uint32(v263)+100)) = v266
						v268 = *(*int64)(unsafe.Add(mBase, uint32(v12)+80))
						*(*int64)(unsafe.Add(mBase, uint32(v263)+92)) = v268
						v270 = int32(1)
						v271 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+72))
						*(*int32)(unsafe.Add(mBase, uint32(v271)+72)) = v272 + v270
						v276 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
						v277 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
						v279 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
						v280 = *(*int32)(unsafe.Add(mBase, uint32(v271)+68))
						if v280 <= int32(310) {
							if v280 == int32(59) {
								if v24 == int32(61) {
									v300 = v270
								} else {
									if v24 == int32(91) {
										v300 = v270
									} else {
										if v24 == int32(270) {
											v300 = v270
										} else {
											v300 = int32(0)
										}
									}
								}
							} else {
								if v280 == int32(287) {
									if v24 == int32(61) {
										v300 = v270
									} else {
										if v24 == int32(91) {
											v300 = v270
										} else {
											if v24 == int32(270) {
												v300 = v270
											} else {
												v300 = int32(0)
											}
										}
									}
								} else {
									v300 = v270
								}
							}
						} else {
							if v280 == int32(311) {
								if v24 == int32(61) {
									v300 = v270
								} else {
									if v24 == int32(91) {
										v300 = v270
									} else {
										if v24 == int32(270) {
											v300 = v270
										} else {
											v300 = int32(0)
										}
									}
								}
							} else {
								if v280 == int32(376) {
									if v24 == int32(61) {
										v300 = v270
									} else {
										if v24 == int32(91) {
											v300 = v270
										} else {
											if v24 == int32(270) {
												v300 = v270
											} else {
												v300 = int32(0)
											}
										}
									}
								} else {
									if v280 != int32(336) {
										v300 = v270
									} else {
										if v24 == int32(61) {
											v300 = v270
										} else {
											if v24 == int32(91) {
												v300 = v270
											} else {
												if v24 == int32(270) {
													v300 = v270
												} else {
													v300 = int32(0)
												}
											}
										}
									}
								}
							}
						}
						v303 = v12 + int32(104)
						v306 = F_plpgsql_parse_word(m, v279, v276+v277, v300, v303, v303)
						mBase = m.M
						v307 = m.ExcPending
						if v307 != 0 {
							return int32(0)
						} else {
							if v306 != 0 {
								v331 = int32(277)
								v336 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
								*(*int64)(unsafe.Add(mBase, uint32(l0))) = v336
								v338 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v338
								v340 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v340
								v342 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
								v343 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								*(*int32)(unsafe.Add(mBase, uint32(v343)+68)) = v331
								*(*int32)(unsafe.Add(mBase, uint32(v343)+64)) = v342
								m.G0 = v12 + int32(128)
								return v331
							} else {
								v308 = int32(275)
								v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+108)))
								if v309 != 0 {
									v331 = v308
									v336 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v336
									v338 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v338
									v340 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = v340
									v342 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
									v343 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									*(*int32)(unsafe.Add(mBase, uint32(v343)+68)) = v331
									*(*int32)(unsafe.Add(mBase, uint32(v343)+64)) = v342
									m.G0 = v12 + int32(128)
									return v331
								} else {
									v310 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
									v312 = F_ScanKeywordLookup(m, v310, int32(4196960))
									mBase = m.M
									v313 = m.ExcPending
									if v313 != 0 {
										return int32(0)
									} else {
										if v312 < int32(0) {
											v331 = v308
										} else {
											v320 = v312 << (uint(int32(1)) % 32)
											v324 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v320)+uint32(_consts[1434]))))
											*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = int32(2191008) + v324
											v330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v320)+uint32(_consts[1377]))))
											v331 = v330
										}
										v336 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
										*(*int64)(unsafe.Add(mBase, uint32(l0))) = v336
										v338 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v338
										v340 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
										*(*int32)(unsafe.Add(mBase, uint32(l1))) = v340
										v342 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
										v343 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										*(*int32)(unsafe.Add(mBase, uint32(v343)+68)) = v331
										*(*int32)(unsafe.Add(mBase, uint32(v343)+64)) = v342
										m.G0 = v12 + int32(128)
										return v331
									}
								}
							}
						}
					}
				}
			}
		default:
			v331 = v16
			v336 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v336
			v338 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v338
			v340 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v340
			v342 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
			v343 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			*(*int32)(unsafe.Add(mBase, uint32(v343)+68)) = v331
			*(*int32)(unsafe.Add(mBase, uint32(v343)+64)) = v342
			m.G0 = v12 + int32(128)
			return v331
		}
	}
}
