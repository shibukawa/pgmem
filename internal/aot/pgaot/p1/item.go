package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_executeItemOptUnwrapResult(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
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
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int64
	_ = v149
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(112)
	m.G0 = v12
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if base.B2i32(l3 == v6)|base.B2i32(v16&int32(1) == v6) == v6 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v12 + int32(112)
	return v201
L2:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	if v180 == int32(0) {
		v201 = v31
		goto L1
	} else {
		goto L42
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = int64(8589934592)
	v29 = v12 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v29
	v31 = int32(2)
	v33 = F_executeItemOptUnwrapTarget(m, l0, l1, l2, v29, int32(1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v178 = F_executeItemOptUnwrapTarget(m, l0, l1, l2, l4, (l3^int32(1))&v16)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L41
	}
L6:
	;
	return int32(0)
L7:
	;
	if v33 == int32(2) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v43 = int32(0)
	v46 = v29
	goto L9
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v49 <= v43 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v159 = int32(0)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	if v160 == v159 {
		v201 = v159
		goto L1
	} else {
		goto L36
	}
L11:
	;
	goto L10
L12:
	;
	v51 = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	if v52 == v51 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	v55 = v43
	v56 = v46
	goto L14
L14:
	;
	v60 = v55 + int32(1)
	v62 = int32(16)
	v63 = v55<<(uint(int32(5))%32) + v56 + v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	switch v64 - v62 {
	case 0:
		goto L18
	default:
		goto L16
	case 2:
		goto L19
	}
L15:
	;
	v55 = v51
	v56 = v52
	goto L14
L16:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v114 < v115 {
		goto L29
	} else {
		goto L30
	}
L17:
	;
	v103 = int32(0)
	v104 = int32(1)
	v109 = F_executeAnyItem(m, l0, v103, v67, l4, v104, v104, v104, v103, v103)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L6
	} else {
		goto L28
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L25
	}
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v68&int32(536870912) != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	if v68&int32(1073741824) != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v77
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapResult_0), v12)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapResult_1), int32(3947), int32(_a_F_executeItemOptUnwrapResult_2))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
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
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v91
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapResult_3), v12+int32(16))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapResult_1), int32(1707), int32(_a_F_executeItemOptUnwrapResult_4))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	v43 = v60
	v46 = v56
	goto L9
L29:
	;
	v119 = v113 + v114<<(uint(int32(5))%32)
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v63)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v119)+40)) = v120
	v122 = *(*int64)(unsafe.Add(mBase, uint32(v63)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v119)+32)) = v122
	v124 = *(*int64)(unsafe.Add(mBase, uint32(v63)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v119)+24)) = v124
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
	*(*int64)(unsafe.Add(mBase, uint32(v119)+16)) = v126
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v128 + int32(1)
	v43 = v60
	v46 = v56
	goto L9
L30:
	;
	goto L31
L31:
	;
	v132 = int32(16)
	v134 = v115 << (uint(int32(1)) % 32)
	if v134 <= v132 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v137 = v132
	goto L34
L33:
	;
	v137 = v134
	goto L34
L34:
	;
	v142 = F_palloc(m, v137<<(uint(int32(5))%32)|int32(16))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = int32(1)
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
	*(*int64)(unsafe.Add(mBase, uint32(v142)+16)) = v149
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v63)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v142)+24)) = v151
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v63)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v142)+32)) = v153
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v63)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v142)+40)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v113)+8)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v142
	v43 = v60
	v46 = v56
	goto L9
L36:
	;
	v166 = v160
	goto L37
L37:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v166)+8))
	F_pfree(m, v166)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L6
	} else {
		goto L39
	}
L38:
	;
	v201 = v159
	goto L1
L39:
	;
	if v172 != 0 {
		v166 = v172
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v201 = v178
	goto L1
L42:
	;
	v186 = v180
	goto L43
L43:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	F_pfree(m, v186)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L6
	} else {
		goto L45
	}
L44:
	;
	v201 = v31
	goto L1
L45:
	;
	if v192 != 0 {
		v186 = v192
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
}
func F_write_item(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
	v14 = F_fwrite(m, v7+int32(12), int32(1), int32(4), l2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		if v14 == int32(4) {
			if l1 != 0 {
				v19 = F_fwrite(m, l0, int32(1), l1, l2)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					if v19 != l1 {
						F_errstart_cold(m, int32(22), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_write_item_0), int32(0))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_write_item_1), int32(_a_F_write_item_2), int32(_a_F_write_item_3))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
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
						m.G0 = v7 + int32(16)
						return
					}
				}
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		} else {
			F_errstart_cold(m, int32(22), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				F_errcode_for_file_access(m)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_write_item_0), int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_write_item_1), int32(_a_F_write_item_4), int32(_a_F_write_item_3))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
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
