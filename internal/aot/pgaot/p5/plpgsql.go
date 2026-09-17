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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
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
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_variable[0]))
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_variable[1]))
				v34 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_variable[2]))
				if v32 == v34 {
					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_variable[2])) = v32 << (uint(int32(1)) % 32)
					v43 = F_repalloc(m, v30, v32<<(uint(int32(3))%32))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_variable[0])) = v43
						v47 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_variable[1]))
						v48 = v43
						v49 = v47
						*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v49
						*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_variable[1])) = v49 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v48+v49<<(uint(int32(2))%32)))) = v14
						if l3 == int32(0) {
							v117 = v14
							m.G0 = v10 + int32(32)
							return v117
						} else {
							F_plpgsql_ns_additem(m, int32(1), v49, l0)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v117 = v14
								m.G0 = v10 + int32(32)
								return v117
							}
						}
					}
				} else {
					v48 = v30
					v49 = v32
					*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v49
					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_variable[1])) = v49 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v48+v49<<(uint(int32(2))%32)))) = v14
					if l3 == int32(0) {
						v117 = v14
						m.G0 = v10 + int32(32)
						return v117
					} else {
						F_plpgsql_ns_additem(m, int32(1), v49, l0)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v117 = v14
							m.G0 = v10 + int32(32)
							return v117
						}
					}
				}
			}
		}
	case 1:
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v66 = F_palloc0(m, int32(40))
		mBase = m.M
		v67 = m.ExcPending
		if v67 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v66))) = int32(2)
			v70 = F_pstrdup(m, l0)
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v66)+32)) = int64(4294967295)
				*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = v64
				*(*int32)(unsafe.Add(mBase, uint32(v66)+24)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v66)+12)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = v70
				v79 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_variable[0]))
				v81 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_variable[1]))
				v83 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_variable[2]))
				if v81 == v83 {
					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_variable[2])) = v81 << (uint(int32(1)) % 32)
					v92 = F_repalloc(m, v79, v81<<(uint(int32(3))%32))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_variable[0])) = v92
						v96 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_variable[1]))
						v97 = v92
						v98 = v96
						*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v98
						*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_variable[1])) = v98 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v97+v98<<(uint(int32(2))%32)))) = v66
						if l3 == int32(0) {
							v117 = v66
							m.G0 = v10 + int32(32)
							return v117
						} else {
							v111 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
							F_plpgsql_ns_additem(m, int32(2), v98, v111)
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return int32(0)
							} else {
								v117 = v66
								m.G0 = v10 + int32(32)
								return v117
							}
						}
					}
				} else {
					v97 = v79
					v98 = v81
					*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v98
					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_build_variable[1])) = v98 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v97+v98<<(uint(int32(2))%32)))) = v66
					if l3 == int32(0) {
						v117 = v66
						m.G0 = v10 + int32(32)
						return v117
					} else {
						v111 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
						F_plpgsql_ns_additem(m, int32(2), v98, v111)
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int32(0)
						} else {
							v117 = v66
							m.G0 = v10 + int32(32)
							return v117
						}
					}
				}
			}
		}
	case 2:
		F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_build_variable_0))
		mBase = m.M
		v126 = m.ExcPending
		if v126 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v129 = m.ExcPending
			if v129 != 0 {
				return int32(0)
			} else {
				v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				v131 = F_format_type_be(m, v130)
				mBase = m.M
				v132 = m.ExcPending
				if v132 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v131
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
					F_errmsg(m, int32(_a_F_plpgsql_build_variable_1), v10+int32(16))
					mBase = m.M
					v139 = m.ExcPending
					if v139 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_plpgsql_build_variable_2), int32(1795), int32(_a_F_plpgsql_build_variable_3))
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
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
		F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_build_variable_0))
		mBase = m.M
		v148 = m.ExcPending
		if v148 != 0 {
			return int32(0)
		} else {
			v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v149
			F_errmsg_internal(m, int32(_a_F_plpgsql_build_variable_4), v10)
			mBase = m.M
			v153 = m.ExcPending
			if v153 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_plpgsql_build_variable_2), int32(1799), int32(_a_F_plpgsql_build_variable_3))
				mBase = m.M
				v158 = m.ExcPending
				if v158 != 0 {
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	v3 = int32(0)
	if l0 < v3 {
		v60 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v60
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
	if v9 == int32(0) {
		v60 = v3
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
	if base.B2i32(v28 == int32(0))|base.B2i32(base.Ui32(v12) <= base.Ui32(v28)) != 0 {
		v60 = v29
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
	v34 = v28
	v37 = v29
	goto L13
L13:
	;
	v39 = int32(1)
	v40 = v37 + v39
	*(*int32)(unsafe.Add(mBase, uint32(v8)+196)) = v40
	v43 = v34 + v39
	*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v43
	v45 = int32(10)
	v46 = F___strchrnul(m, v43, v45)
	mBase = m.M
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v48 == v45 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v60 = v40
	goto L1
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+192)) = v52
	if v52 == int32(0) {
		v60 = v40
		goto L1
	} else {
		goto L19
	}
L16:
	;
	v52 = v46
	goto L18
L17:
	;
	v52 = int32(0)
	goto L18
L18:
	;
	goto L15
L19:
	;
	if base.Ui32(v52) < base.Ui32(v12) {
		v34 = v52
		v37 = v40
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L14
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	if v35 != 0 {
		v3 = v35
		goto L4
	} else {
		goto L16
	}
L7:
	;
	v7 = v3 + int32(12)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v10 == int32(0))|base.B2i32(v10 != v13) != 0 {
		v31 = v10
		v32 = v13
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v31-v32 != 0 {
		goto L6
	} else {
		goto L15
	}
L9:
	;
	goto L8
L10:
	;
	v16 = v7
	v17 = l1
	goto L11
L11:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v21 == int32(0) {
		v31 = v21
		v32 = v20
		goto L9
	} else {
		goto L13
	}
L12:
	;
	v31 = v21
	v32 = v20
	goto L9
L13:
	;
	v24 = int32(1)
	if v21 == v20 {
		v16 = v16 + v24
		v17 = v17 + v24
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	return v3
L16:
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
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
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
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_param_eval_recfield_0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L24
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_param_eval_recfield_0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
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
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v86
	F_errmsg(m, int32(_a_F_plpgsql_param_eval_recfield_1), v9+int32(16))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_plpgsql_param_eval_recfield_2), int32(_a_F_plpgsql_param_eval_recfield_3), int32(_a_F_plpgsql_param_eval_recfield_4))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
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
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	v109 = F_format_type_be(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v112 = F_format_type_be(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v107
	F_errmsg(m, int32(_a_F_plpgsql_param_eval_recfield_5), v9)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_plpgsql_param_eval_recfield_2), int32(_a_F_plpgsql_param_eval_recfield_6), int32(_a_F_plpgsql_param_eval_recfield_4))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
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
	v10 = int32(_a_F_plpgsql_parse_cwordrowtype_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_cwordrowtype[0]))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_cwordrowtype[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_cwordrowtype[0])) = v14
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
					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_cwordrowtype[0])) = v11
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
								F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_parse_cwordrowtype_1))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v26
									F_errmsg_internal(m, int32(_a_F_plpgsql_parse_cwordrowtype_2), v8+int32(16))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_plpgsql_parse_cwordrowtype_3), int32(1960), int32(_a_F_plpgsql_parse_cwordrowtype_4))
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
					F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_parse_cwordrowtype_1))
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
							F_errmsg(m, int32(_a_F_plpgsql_parse_cwordrowtype_5), v8)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_plpgsql_parse_cwordrowtype_3), int32(1729), int32(_a_F_plpgsql_parse_cwordrowtype_6))
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
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v166 int32
	_ = v166
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
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
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(_a_F_plpgsql_parse_cwordtype_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_cwordtype[0]))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_cwordtype[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_cwordtype[0])) = v14
	if l0 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_parse_cwordtype_1))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L51
	} else {
		goto L68
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_parse_cwordtype_1))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L51
	} else {
		goto L64
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_cwordtype[0])) = v11
	m.G0 = v8 + int32(32)
	return v269
L4:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	v234 = int32(0)
	v238 = F_RangeVarGetRelidExtended(m, v229, v234, v234, v234, v234)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L51
	} else {
		goto L56
	}
L5:
	;
	v216 = F_list_copy(m, l0)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L51
	} else {
		goto L53
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
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_cwordtype[2]))
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
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	v209 = F_makeRangeVar(m, int32(0), v207, int32(-1))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L51
	} else {
		goto L52
	}
L9:
	;
	if v176 == int32(0) {
		goto L8
	} else {
		goto L45
	}
L10:
	;
	v176 = v166
	goto L9
L11:
	;
	v166 = v152
	goto L10
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v144
	v152 = v142
	goto L11
L13:
	;
	v134 = int32(0)
	if v31 == v134 {
		v166 = v134
		goto L10
	} else {
		goto L44
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
	v78 = v39
	goto L19
L19:
	;
	if v28 == int32(0) {
		v119 = v78
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v60 = F_strcmp(m, v53+int32(12), v26)
	mBase = m.M
	v63 = int32(0)
	if v60|base.B2i32(v55 == int32(1))&base.B2i32(v28 != v63) == v63 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v78 = v72
	goto L19
L22:
	;
	if v31 == int32(0) {
		v152 = v53
		goto L11
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v73 != 0 {
		v53 = v72
		v55 = v73
		goto L20
	} else {
		goto L26
	}
L25:
	;
	v142 = v53
	v144 = int32(1)
	goto L12
L26:
	;
	goto L21
L27:
	;
	goto L42
L28:
	;
	v87 = F_strcmp(m, v78+int32(12), v26)
	mBase = m.M
	if v87 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v88 = v78
	goto L31
L30:
	;
	v88 = v39
	goto L31
L31:
	;
	if v87|base.B2i32(v48 == int32(0)) != 0 {
		v119 = v88
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v92 = v39
	v99 = v48
	goto L33
L33:
	;
	v103 = F_strcmp(m, v92+int32(12), v28)
	mBase = m.M
	if v103|int32(0)&base.B2i32(v99 == int32(1)) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v119 = v113
	goto L27
L35:
	;
	if v31 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	if v114 != 0 {
		v92 = v113
		v99 = v114
		goto L33
	} else {
		goto L41
	}
L38:
	;
	v176 = v92
	goto L9
L39:
	;
	goto L40
L40:
	;
	v142 = v92
	v144 = int32(2)
	goto L12
L41:
	;
	goto L34
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
	if v124 != 0 {
		v39 = v124
		goto L15
	} else {
		goto L43
	}
L43:
	;
	goto L16
L44:
	;
	v142 = v134
	v144 = v134
	goto L12
L45:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	if v179 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_cwordtype[3]))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v183+v184<<(uint(int32(2))%32))))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+24))
	v269 = v189
	goto L3
L47:
	;
	goto L48
L48:
	;
	if v179 != int32(2) {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	if v192 != int32(2) {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_cwordtype[3]))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v196+v197<<(uint(int32(2))%32))))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+24))
	v269 = v202
	goto L3
L51:
	;
	return int32(0)
L52:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v229 = v209
	v231 = v213 + int32(4)
	goto L4
L53:
	;
	v218 = F_list_delete_last(m, v216)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L51
	} else {
		goto L54
	}
L54:
	;
	v220 = F_makeRangeVarFromNameList(m, v218)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v229 = v220
	v231 = v222 + v223<<(uint(int32(2))%32) - int32(4)
	goto L4
L56:
	;
	v240 = F_SearchSysCacheAttName(m, v238, v233)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L51
	} else {
		goto L57
	}
L57:
	;
	if v240 == int32(0) {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v240)+16))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+22)))
	v247 = v245 + v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+68))
	v249 = F_SearchSysCache1(m, int32(82), v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L51
	} else {
		goto L59
	}
L59:
	;
	if v249 == int32(0) {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_cwordtype[0])) = v11
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v247)+76))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v247)+96))
	v258 = F_build_datatype(m, v249, v255, v256, int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L51
	} else {
		goto L61
	}
L61:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_cwordtype[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_cwordtype[0])) = v262
	F_ReleaseCatCache(m, v240)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L51
	} else {
		goto L62
	}
L62:
	;
	F_ReleaseCatCache(m, v249)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L51
	} else {
		goto L63
	}
L63:
	;
	v269 = v258
	goto L3
L64:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L51
	} else {
		goto L65
	}
L65:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v229)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v233
	F_errmsg(m, int32(_a_F_plpgsql_parse_cwordtype_2), v8)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L51
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_plpgsql_parse_cwordtype_3), int32(1631), int32(_a_F_plpgsql_parse_cwordtype_4))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L51
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v247)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v299
	F_errmsg_internal(m, int32(_a_F_plpgsql_parse_cwordtype_5), v8+int32(16))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L51
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_plpgsql_parse_cwordtype_3), int32(1637), int32(_a_F_plpgsql_parse_cwordtype_4))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L51
	} else {
		goto L70
	}
L70:
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
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_tripword[0]))
	if v13 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(80)
	return v262
L2:
	;
	v235 = F_makeString(m, l0)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L46
	} else {
		goto L56
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_tripword[1]))
	v20 = v10 + int32(76)
	if v17 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	if v165 == int32(0) {
		goto L2
	} else {
		goto L40
	}
L5:
	;
	v165 = v155
	goto L4
L6:
	;
	v155 = v141
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v133
	v141 = v131
	goto L6
L8:
	;
	v123 = int32(0)
	if v20 == v123 {
		v155 = v123
		goto L5
	} else {
		goto L39
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
	v67 = v28
	goto L14
L14:
	;
	if l1 == int32(0) {
		v108 = v67
		goto L22
	} else {
		goto L23
	}
L15:
	;
	v49 = F_strcmp(m, v42+int32(12), l0)
	mBase = m.M
	v52 = int32(0)
	if v49|base.B2i32(v44 == int32(1))&base.B2i32(l1 != v52) == v52 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v67 = v61
	goto L14
L17:
	;
	if v20 == int32(0) {
		v141 = v42
		goto L6
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v62 != 0 {
		v42 = v61
		v44 = v62
		goto L15
	} else {
		goto L21
	}
L20:
	;
	v131 = v42
	v133 = int32(1)
	goto L7
L21:
	;
	goto L16
L22:
	;
	goto L37
L23:
	;
	v76 = F_strcmp(m, v67+int32(12), l0)
	mBase = m.M
	if v76 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v77 = v67
	goto L26
L25:
	;
	v77 = v28
	goto L26
L26:
	;
	if v76|base.B2i32(v37 == int32(0)) != 0 {
		v108 = v77
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v81 = v28
	v88 = v37
	goto L28
L28:
	;
	v92 = F_strcmp(m, v81+int32(12), l1)
	mBase = m.M
	if v92|base.B2i32(l2 != int32(0))&base.B2i32(v88 == int32(1)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v108 = v102
	goto L22
L30:
	;
	if v20 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	if v103 != 0 {
		v81 = v102
		v88 = v103
		goto L28
	} else {
		goto L36
	}
L33:
	;
	v165 = v81
	goto L4
L34:
	;
	goto L35
L35:
	;
	v131 = v81
	v133 = int32(2)
	goto L7
L36:
	;
	goto L29
L37:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	if v113 != 0 {
		v28 = v113
		goto L10
	} else {
		goto L38
	}
L38:
	;
	goto L11
L39:
	;
	v131 = v123
	v133 = v123
	goto L7
L40:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if v168 != int32(2) {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_parse_tripword[2]))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v172+v173<<(uint(int32(2))%32))))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
	if v178 == int32(1) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v226
	v228 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v228)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v225
	v262 = int32(1)
	goto L1
L43:
	;
	v181 = F_plpgsql_build_recfield(m, v177, l1)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	v200 = F_plpgsql_build_recfield(m, v177, l2)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L46
	} else {
		goto L51
	}
L46:
	;
	return int32(0)
L47:
	;
	v185 = F_makeString(m, l0)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v185
	v188 = F_makeString(m, l1)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L46
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v188
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v10)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v192
	v198 = F_list_make2_impl(m, v10+int32(28), v10+int32(24))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v225 = v181
	v226 = v198
	goto L42
L51:
	;
	v202 = F_makeString(m, l0)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L46
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v202
	v205 = F_makeString(m, l1)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L46
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v205
	v208 = F_makeString(m, l2)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L46
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v208
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v212
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v214
	v222 = F_list_make3_impl(m, v10+int32(40), v10+int32(36), v10+int32(32))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L46
	} else {
		goto L55
	}
L55:
	;
	v225 = v200
	v226 = v222
	goto L42
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v235
	v238 = F_makeString(m, l1)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L46
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v238
	v241 = F_makeString(m, l2)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L46
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v241
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v245
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v247
	v255 = F_list_make3_impl(m, v10+int32(20), v10+int32(16), v10+int32(12))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L46
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v255
	v262 = int32(0)
	goto L1
}
func F_plpgsql_parser_setup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(_a_F_plpgsql_parser_setup_0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(_a_F_plpgsql_parser_setup_1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(_a_F_plpgsql_parser_setup_2)
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
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
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
					F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_peek2_0))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_plpgsql_peek2_1), int32(0))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_plpgsql_peek2_2), int32(388), int32(_a_F_plpgsql_peek2_3))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
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
				F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_peek2_0))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_plpgsql_peek2_1), int32(0))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_plpgsql_peek2_2), int32(388), int32(_a_F_plpgsql_peek2_3))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
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
	var v24 int32
	_ = v24
	v3 = int32(0)
	if l0 < v3 {
		v24 = v3
		return v24
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+60))
		if v8 == int32(0) {
			v24 = v3
			return v24
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
						v24 = v21
						return v24
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int64
	_ = v176
	var v178 int64
	_ = v178
	var v180 int64
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int64
	_ = v198
	var v200 int64
	_ = v200
	var v202 int64
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int64
	_ = v250
	var v252 int64
	_ = v252
	var v254 int64
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int64
	_ = v319
	var v321 int64
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
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
											v55 = F_plpgsql_parse_tripword(m, v50, v51, v52, v54, v54)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return int32(0)
											} else {
												v113 = v12 + int32(24)
												v114 = v12 + int32(28)
												v116 = v55
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
												*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = v117 + (v118 - v119)
												if v116 != 0 {
													v125 = int32(277)
												} else {
													v125 = int32(276)
												}
												v315 = v125
												v319 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v319
												v321 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
												*(*int64)(unsafe.Add(mBase, uint32(l0))) = v321
												v323 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
												*(*int32)(unsafe.Add(mBase, uint32(l1))) = v323
												v325 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
												v326 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
												*(*int32)(unsafe.Add(mBase, uint32(v326)+68)) = v315
												*(*int32)(unsafe.Add(mBase, uint32(v326)+64)) = v325
												m.G0 = v12 + int32(128)
												return v315
											}
										} else {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+72))
											if int32(4) <= v58 {
												F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yylex_0))
												mBase = m.M
												v340 = m.ExcPending
												if v340 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(_a_F_plpgsql_yylex_1), int32(0))
													mBase = m.M
													v344 = m.ExcPending
													if v344 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_plpgsql_yylex_2), int32(388), int32(_a_F_plpgsql_yylex_3))
														mBase = m.M
														v349 = m.ExcPending
														if v349 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v57+v58<<(uint(int32(2))%32))+76)) = v42
												v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)+72))
												v68 = v57 + v65*int32(24)
												v69 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
												*(*int64)(unsafe.Add(mBase, uint32(v68)+108)) = v69
												v71 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
												*(*int64)(unsafe.Add(mBase, uint32(v68)+100)) = v71
												v73 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
												*(*int64)(unsafe.Add(mBase, uint32(v68)+92)) = v73
												v75 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
												v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+72))
												v78 = v76 + int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v75)+72)) = v78
												if int32(3) <= v76 {
													F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yylex_0))
													mBase = m.M
													v340 = m.ExcPending
													if v340 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_plpgsql_yylex_1), int32(0))
														mBase = m.M
														v344 = m.ExcPending
														if v344 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_plpgsql_yylex_2), int32(388), int32(_a_F_plpgsql_yylex_3))
															mBase = m.M
															v349 = m.ExcPending
															if v349 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v75+v78<<(uint(int32(2))%32))+76)) = int32(46)
													v91 = *(*int32)(unsafe.Add(mBase, uint32(v75)+72))
													v94 = v75 + v91*int32(24)
													v95 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
													*(*int64)(unsafe.Add(mBase, uint32(v94)+108)) = v95
													v97 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
													*(*int64)(unsafe.Add(mBase, uint32(v94)+100)) = v97
													v99 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
													*(*int64)(unsafe.Add(mBase, uint32(v94)+92)) = v99
													v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
													v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+72))
													*(*int32)(unsafe.Add(mBase, uint32(v101)+72)) = v102 + int32(1)
													v106 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
													v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
													v109 = v12 + int32(104)
													v110 = F_plpgsql_parse_dblword(m, v106, v107, v109, v109)
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return int32(0)
													} else {
														v113 = v12 + int32(72)
														v114 = v12 + int32(76)
														v116 = v110
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
														v119 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
														*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = v117 + (v118 - v119)
														if v116 != 0 {
															v125 = int32(277)
														} else {
															v125 = int32(276)
														}
														v315 = v125
														v319 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
														*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v319
														v321 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
														*(*int64)(unsafe.Add(mBase, uint32(l0))) = v321
														v323 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
														*(*int32)(unsafe.Add(mBase, uint32(l1))) = v323
														v325 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
														v326 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
														*(*int32)(unsafe.Add(mBase, uint32(v326)+68)) = v315
														*(*int32)(unsafe.Add(mBase, uint32(v326)+64)) = v325
														m.G0 = v12 + int32(128)
														return v315
													}
												}
											}
										}
									}
								} else {
									v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+72))
									if int32(4) <= v127 {
										F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yylex_0))
										mBase = m.M
										v340 = m.ExcPending
										if v340 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_plpgsql_yylex_1), int32(0))
											mBase = m.M
											v344 = m.ExcPending
											if v344 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_plpgsql_yylex_2), int32(388), int32(_a_F_plpgsql_yylex_3))
												mBase = m.M
												v349 = m.ExcPending
												if v349 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v126+v127<<(uint(int32(2))%32))+76)) = v36
										v134 = *(*int32)(unsafe.Add(mBase, uint32(v126)+72))
										v137 = v126 + v134*int32(24)
										v138 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v137)+108)) = v138
										v140 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v137)+100)) = v140
										v142 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v137)+92)) = v142
										v144 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+72))
										*(*int32)(unsafe.Add(mBase, uint32(v144)+72)) = v145 + int32(1)
										v149 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
										v150 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
										v152 = v12 + int32(104)
										v153 = F_plpgsql_parse_dblword(m, v149, v150, v152, v152)
										mBase = m.M
										v154 = m.ExcPending
										if v154 != 0 {
											return int32(0)
										} else {
											v155 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
											v156 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
											v157 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
											*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = v155 + (v156 - v157)
											if v153 != 0 {
												v163 = int32(277)
											} else {
												v163 = int32(276)
											}
											v315 = v163
											v319 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
											*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v319
											v321 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
											*(*int64)(unsafe.Add(mBase, uint32(l0))) = v321
											v323 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
											*(*int32)(unsafe.Add(mBase, uint32(l1))) = v323
											v325 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
											v326 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											*(*int32)(unsafe.Add(mBase, uint32(v326)+68)) = v315
											*(*int32)(unsafe.Add(mBase, uint32(v326)+64)) = v325
											m.G0 = v12 + int32(128)
											return v315
										}
									}
								}
							}
						} else {
							v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+72))
							if int32(4) <= v165 {
								F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yylex_0))
								mBase = m.M
								v340 = m.ExcPending
								if v340 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_plpgsql_yylex_1), int32(0))
									mBase = m.M
									v344 = m.ExcPending
									if v344 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_plpgsql_yylex_2), int32(388), int32(_a_F_plpgsql_yylex_3))
										mBase = m.M
										v349 = m.ExcPending
										if v349 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v164+v165<<(uint(int32(2))%32))+76)) = v30
								v172 = *(*int32)(unsafe.Add(mBase, uint32(v164)+72))
								v175 = v164 + v172*int32(24)
								v176 = *(*int64)(unsafe.Add(mBase, uint32(v12)+72))
								*(*int64)(unsafe.Add(mBase, uint32(v175)+108)) = v176
								v178 = *(*int64)(unsafe.Add(mBase, uint32(v12)+64))
								*(*int64)(unsafe.Add(mBase, uint32(v175)+100)) = v178
								v180 = *(*int64)(unsafe.Add(mBase, uint32(v12)+56))
								*(*int64)(unsafe.Add(mBase, uint32(v175)+92)) = v180
								v182 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+72))
								v185 = v183 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v182)+72)) = v185
								if int32(3) <= v183 {
									F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yylex_0))
									mBase = m.M
									v340 = m.ExcPending
									if v340 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(_a_F_plpgsql_yylex_1), int32(0))
										mBase = m.M
										v344 = m.ExcPending
										if v344 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_plpgsql_yylex_2), int32(388), int32(_a_F_plpgsql_yylex_3))
											mBase = m.M
											v349 = m.ExcPending
											if v349 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v182+v185<<(uint(int32(2))%32))+76)) = int32(46)
									v194 = *(*int32)(unsafe.Add(mBase, uint32(v182)+72))
									v197 = v182 + v194*int32(24)
									v198 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
									*(*int64)(unsafe.Add(mBase, uint32(v197)+108)) = v198
									v200 = *(*int64)(unsafe.Add(mBase, uint32(v12)+88))
									*(*int64)(unsafe.Add(mBase, uint32(v197)+100)) = v200
									v202 = *(*int64)(unsafe.Add(mBase, uint32(v12)+80))
									*(*int64)(unsafe.Add(mBase, uint32(v197)+92)) = v202
									v204 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+72))
									v206 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v204)+72)) = v205 + v206
									v210 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
									v211 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
									v212 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
									v216 = v12 + int32(104)
									v217 = F_plpgsql_parse_word(m, v210, v211+v212, v206, v216, v216)
									mBase = m.M
									v218 = m.ExcPending
									if v218 != 0 {
										return int32(0)
									} else {
										if v217 != 0 {
											v315 = int32(277)
											v319 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
											*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v319
											v321 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
											*(*int64)(unsafe.Add(mBase, uint32(l0))) = v321
											v323 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
											*(*int32)(unsafe.Add(mBase, uint32(l1))) = v323
											v325 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
											v326 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											*(*int32)(unsafe.Add(mBase, uint32(v326)+68)) = v315
											*(*int32)(unsafe.Add(mBase, uint32(v326)+64)) = v325
											m.G0 = v12 + int32(128)
											return v315
										} else {
											v219 = int32(275)
											v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+108)))
											if v220 != 0 {
												v315 = v219
												v319 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v319
												v321 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
												*(*int64)(unsafe.Add(mBase, uint32(l0))) = v321
												v323 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
												*(*int32)(unsafe.Add(mBase, uint32(l1))) = v323
												v325 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
												v326 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
												*(*int32)(unsafe.Add(mBase, uint32(v326)+68)) = v315
												*(*int32)(unsafe.Add(mBase, uint32(v326)+64)) = v325
												m.G0 = v12 + int32(128)
												return v315
											} else {
												v221 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
												v223 = F_ScanKeywordLookup(m, v221, int32(_a_F_plpgsql_yylex_4))
												mBase = m.M
												v224 = m.ExcPending
												if v224 != 0 {
													return int32(0)
												} else {
													if v223 < int32(0) {
														v315 = v219
													} else {
														v228 = v223 << (uint(int32(1)) % 32)
														v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228)+uint32(_c_F_plpgsql_yylex[0]))))
														*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = v231 + int32(_a_F_plpgsql_yylex_5)
														v237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228)+uint32(_c_F_plpgsql_yylex[1]))))
														v315 = v237
													}
													v319 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v319
													v321 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
													*(*int64)(unsafe.Add(mBase, uint32(l0))) = v321
													v323 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
													*(*int32)(unsafe.Add(mBase, uint32(l1))) = v323
													v325 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
													v326 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
													*(*int32)(unsafe.Add(mBase, uint32(v326)+68)) = v315
													*(*int32)(unsafe.Add(mBase, uint32(v326)+64)) = v325
													m.G0 = v12 + int32(128)
													return v315
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v238 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+72))
					if int32(4) <= v239 {
						F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_yylex_0))
						mBase = m.M
						v340 = m.ExcPending
						if v340 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_plpgsql_yylex_1), int32(0))
							mBase = m.M
							v344 = m.ExcPending
							if v344 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_plpgsql_yylex_2), int32(388), int32(_a_F_plpgsql_yylex_3))
								mBase = m.M
								v349 = m.ExcPending
								if v349 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v238+v239<<(uint(int32(2))%32))+76)) = v24
						v246 = *(*int32)(unsafe.Add(mBase, uint32(v238)+72))
						v249 = v238 + v246*int32(24)
						v250 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
						*(*int64)(unsafe.Add(mBase, uint32(v249)+108)) = v250
						v252 = *(*int64)(unsafe.Add(mBase, uint32(v12)+88))
						*(*int64)(unsafe.Add(mBase, uint32(v249)+100)) = v252
						v254 = *(*int64)(unsafe.Add(mBase, uint32(v12)+80))
						*(*int64)(unsafe.Add(mBase, uint32(v249)+92)) = v254
						v256 = int32(1)
						v257 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+72))
						*(*int32)(unsafe.Add(mBase, uint32(v257)+72)) = v258 + v256
						v262 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
						v263 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
						v265 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
						v266 = *(*int32)(unsafe.Add(mBase, uint32(v257)+68))
						if v266 <= int32(310) {
							if base.B2i32(v266 == int32(59))|base.B2i32(v266 == int32(287)) != 0 {
								if base.B2i32(v24 == int32(61))|base.B2i32(v24 == int32(91))|base.B2i32(v24 == int32(270)) != 0 {
									v290 = v256
								} else {
									v290 = int32(0)
								}
							} else {
								v290 = v256
							}
						} else {
							if base.B2i32(v266 == int32(311))|base.B2i32(v266 == int32(376)) != 0 {
								if base.B2i32(v24 == int32(61))|base.B2i32(v24 == int32(91))|base.B2i32(v24 == int32(270)) != 0 {
									v290 = v256
								} else {
									v290 = int32(0)
								}
							} else {
								if v266 != int32(336) {
									v290 = v256
								} else {
									if base.B2i32(v24 == int32(61))|base.B2i32(v24 == int32(91))|base.B2i32(v24 == int32(270)) != 0 {
										v290 = v256
									} else {
										v290 = int32(0)
									}
								}
							}
						}
						v293 = v12 + int32(104)
						v294 = F_plpgsql_parse_word(m, v265, v262+v263, v290, v293, v293)
						mBase = m.M
						v295 = m.ExcPending
						if v295 != 0 {
							return int32(0)
						} else {
							if v294 != 0 {
								v315 = int32(277)
								v319 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v319
								v321 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
								*(*int64)(unsafe.Add(mBase, uint32(l0))) = v321
								v323 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v323
								v325 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
								v326 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								*(*int32)(unsafe.Add(mBase, uint32(v326)+68)) = v315
								*(*int32)(unsafe.Add(mBase, uint32(v326)+64)) = v325
								m.G0 = v12 + int32(128)
								return v315
							} else {
								v296 = int32(275)
								v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+108)))
								if v297 != 0 {
									v315 = v296
									v319 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v319
									v321 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v321
									v323 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = v323
									v325 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
									v326 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									*(*int32)(unsafe.Add(mBase, uint32(v326)+68)) = v315
									*(*int32)(unsafe.Add(mBase, uint32(v326)+64)) = v325
									m.G0 = v12 + int32(128)
									return v315
								} else {
									v298 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
									v300 = F_ScanKeywordLookup(m, v298, int32(_a_F_plpgsql_yylex_4))
									mBase = m.M
									v301 = m.ExcPending
									if v301 != 0 {
										return int32(0)
									} else {
										if v300 < int32(0) {
											v315 = v296
										} else {
											v305 = v300 << (uint(int32(1)) % 32)
											v308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v305)+uint32(_c_F_plpgsql_yylex[0]))))
											*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = v308 + int32(_a_F_plpgsql_yylex_5)
											v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v305)+uint32(_c_F_plpgsql_yylex[1]))))
											v315 = v314
										}
										v319 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v319
										v321 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
										*(*int64)(unsafe.Add(mBase, uint32(l0))) = v321
										v323 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
										*(*int32)(unsafe.Add(mBase, uint32(l1))) = v323
										v325 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
										v326 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										*(*int32)(unsafe.Add(mBase, uint32(v326)+68)) = v315
										*(*int32)(unsafe.Add(mBase, uint32(v326)+64)) = v325
										m.G0 = v12 + int32(128)
										return v315
									}
								}
							}
						}
					}
				}
			}
		default:
			v315 = v16
			v319 = *(*int64)(unsafe.Add(mBase, uint32(v12)+112))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v319
			v321 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v321
			v323 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v323
			v325 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
			v326 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			*(*int32)(unsafe.Add(mBase, uint32(v326)+68)) = v315
			*(*int32)(unsafe.Add(mBase, uint32(v326)+64)) = v325
			m.G0 = v12 + int32(128)
			return v315
		}
	}
}
