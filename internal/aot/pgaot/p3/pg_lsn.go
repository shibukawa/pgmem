package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_lsn_in(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int64
	_ = v224
	var v232 int64
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = int32(531878)
	v15 = m.G0
	v17 = v15 - int32(32)
	v18 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v18
	v26 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1015])))
	if v26 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v239
L2:
	;
	v224 = F_strtox_2(m, v10, int32(0), int32(16), int64(4294967295))
	mBase = m.M
	goto L56
L3:
	;
	v199 = int32(0)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v201 = F_errsave_start(m, v200)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L50
	} else {
		goto L51
	}
L4:
	;
	if base.Ui32(v94-int32(9)) < base.Ui32(int32(-8)) {
		goto L3
	} else {
		goto L25
	}
L5:
	;
	v94 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1016])))
	if v30 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v34 = v10
	goto L11
L9:
	;
	goto L10
L10:
	;
	v44 = v11
	v45 = v26
	goto L14
L11:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v40 == v26 {
		v34 = v34 + int32(1)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v94 = v34 - v10
	goto L4
L13:
	;
	goto L12
L14:
	;
	v52 = v17 + int32(base.Ui32(v45)>>(uint(int32(3))%32))&int32(28)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v54 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v53 | v54<<(uint(v45)%32)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	if v58 != 0 {
		v44 = v44 + v54
		v45 = v58
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v61 == int32(0) {
		v86 = v10
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v94 = v86 - v10
	goto L4
L18:
	;
	v65 = v10
	v66 = v61
	goto L19
L19:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(base.Ui32(v66)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v74)>>(uint(v66)%32))&int32(1) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v86 = v82
	goto L17
L21:
	;
	v86 = v65
	goto L17
L22:
	;
	goto L23
L23:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
	v82 = v65 + int32(1)
	if v80 != 0 {
		v65 = v82
		v66 = v80
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v99 = v94 + v10
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v100 != int32(47) {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v104 = v99 + int32(1)
	v105 = int32(531878)
	v109 = m.G0
	v111 = v109 - int32(32)
	v112 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v111)+24)) = v112
	*(*int64)(unsafe.Add(mBase, uint32(v111)+16)) = v112
	*(*int64)(unsafe.Add(mBase, uint32(v111)+8)) = v112
	*(*int64)(unsafe.Add(mBase, uint32(v111))) = v112
	v120 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1015])))
	if v120 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if base.Ui32(v188-int32(9)) < base.Ui32(int32(-8)) {
		goto L3
	} else {
		goto L48
	}
L28:
	;
	v188 = int32(0)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1016])))
	if v124 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v128 = v104
	goto L34
L32:
	;
	goto L33
L33:
	;
	v138 = v105
	v139 = v120
	goto L37
L34:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v134 == v120 {
		v128 = v128 + int32(1)
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v188 = v128 - v104
	goto L27
L36:
	;
	goto L35
L37:
	;
	v146 = v111 + int32(base.Ui32(v139)>>(uint(int32(3))%32))&int32(28)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v148 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v147 | v148<<(uint(v139)%32)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)))
	if v152 != 0 {
		v138 = v138 + v148
		v139 = v152
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v155 == int32(0) {
		v180 = v104
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	v188 = v180 - v104
	goto L27
L41:
	;
	v159 = v104
	v160 = v155
	goto L42
L42:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v111+int32(base.Ui32(v160)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v168)>>(uint(v160)%32))&int32(1) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v180 = v176
	goto L40
L44:
	;
	v180 = v159
	goto L40
L45:
	;
	goto L46
L46:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+1)))
	v176 = v159 + int32(1)
	if v174 != 0 {
		v159 = v176
		v160 = v174
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
L48:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104+v188))))
	if v194 == int32(0) {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	goto L3
L50:
	;
	return int32(0)
L51:
	;
	if v201 == int32(0) {
		v239 = v199
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(242529)
	F_errmsg(m, int32(704449), v8)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L50
	} else {
		goto L54
	}
L54:
	;
	F_errsave_finish(m, v200, int32(491207), int32(74), int32(276830))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L50
	} else {
		goto L55
	}
L55:
	;
	v239 = v199
	goto L1
L56:
	;
	v232 = F_strtox_2(m, v104, int32(0), int32(16), int64(4294967295))
	mBase = m.M
	goto L57
L57:
	;
	v236 = F_Int64GetDatum(m, base.I64_extend_i32_u(base.I32_wrap_i64(v224))<<(uint(int64(32))%64)|base.I64_extend_i32_u(base.I32_wrap_i64(v232)))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L50
	} else {
		goto L58
	}
L58:
	;
	v239 = v236
	goto L1
}
func F_pg_lsn_mii(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
		if v15 == int32(49152) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(242504), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(491207), int32(296), int32(317037))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
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
			*(*int64)(unsafe.Add(mBase, uint32(v6))) = v9
			v39 = F_pg_snprintf(m, v6+int32(16), int32(32), int32(38081), v6)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				v42 = int32(0)
				v51 = F_DirectFunctionCall3Coll(m, int32(408), v42, v6+int32(16), v42, int32(-1))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					v53 = F_DirectFunctionCall2Coll(m, int32(18), v42, v51, v11)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v55 = F_DirectFunctionCall1Coll(m, int32(1481), v42, v53)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							m.G0 = v6 + int32(48)
							return v55
						}
					}
				}
			}
		}
	}
}
func F_pg_lsn_pli(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
		if v15 == int32(49152) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(242479), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(491207), int32(262), int32(316984))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
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
			*(*int64)(unsafe.Add(mBase, uint32(v6))) = v9
			v39 = F_pg_snprintf(m, v6+int32(16), int32(32), int32(38081), v6)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				v42 = int32(0)
				v51 = F_DirectFunctionCall3Coll(m, int32(408), v42, v6+int32(16), v42, int32(-1))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					v53 = F_DirectFunctionCall2Coll(m, int32(1294), v42, v51, v11)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v55 = F_DirectFunctionCall1Coll(m, int32(1481), v42, v53)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							m.G0 = v6 + int32(48)
							return v55
						}
					}
				}
			}
		}
	}
}
