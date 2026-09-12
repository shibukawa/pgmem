package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetSystemIdentifier(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _consts[178]))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	return v3
}
func F_SystemAttributeDefinition(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	if base.Ui32(l0) <= base.Ui32(int32(-7)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
			F_errmsg_internal(m, int32(472213), v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(496316), int32(239), int32(250795))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, uint32((l0^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[199])))
		m.G0 = v5 + int32(16)
		return v30
	}
}
func F_system_beginsamplescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 float32
	_ = v7
	var v12 float64
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 float64
	_ = v28
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v7 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	if base.F32_lt(v7, float32(0)) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return
		} else {
			F_errcode(m, int32(403177602))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				F_errmsg(m, int32(567062), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					F_errfinish(m, int32(497432), int32(151), int32(284751))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
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
		if base.F32_gt(v7, float32(100)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				F_errcode(m, int32(403177602))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					F_errmsg(m, int32(567062), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						F_errfinish(m, int32(497432), int32(151), int32(284751))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
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
			v12 = base.F64_promote_f32(v7)
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					F_errcode(m, int32(403177602))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						F_errmsg(m, int32(567062), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							F_errfinish(m, int32(497432), int32(151), int32(284751))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
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
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v19 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v18)+16)) = uint16(v19)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v19
				*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l3
				v28 = base.F64_nearest(base.F64_div(base.F64_mul(v12, float64(4.294967296e+09)), float64(100)))
				if base.F64_lt(v28, float64(1.8446744073709552e+19))&base.F64_ge(v28, float64(0)) != 0 {
					v34 = base.I64_trunc_f64_u(v28)
					v36 = v34
				} else {
					v36 = int64(0)
				}
				*(*int64)(unsafe.Add(mBase, uint32(v18))) = v36
				v38 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v38)
				v41 = base.F32_ge(v7, float32(1))
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)) = uint8(v41)
				return
			}
		}
	}
}
func F_system_rows_nextsampleblock(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
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
	var v28 int64
	_ = v28
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v48 int64
	_ = v48
	var v53 int64
	_ = v53
	var v58 int64
	_ = v58
	var v74 float64
	_ = v74
	var v77 int32
	_ = v77
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 float64
	_ = v114
	var v117 float64
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var __phi131 int32
	_ = __phi131
	var v133 int32
	_ = v133
	var __phi133 int32
	_ = __phi133
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int64
	_ = v193
	var v194 int64
	_ = v194
	var v196 int32
	_ = v196
	var v198 int64
	_ = v198
	var v202 int32
	_ = v202
	var v214 int64
	_ = v214
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	if v19 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v16 + int32(16)
	return v222
L2:
	;
	v222 = int32(-1)
	goto L1
L3:
	;
	v22 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	if v23 == v22 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v177 = v19
	goto L5
L5:
	;
	v188 = v177 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v188
	v190 = int32(-1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if base.Ui32(v191) < base.Ui32(v188) {
		v222 = v190
		goto L1
	} else {
		goto L45
	}
L6:
	;
	if l1 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L7:
	;
	v162 = v22
	goto L8
L8:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v172
	v177 = v162
	goto L5
L9:
	;
	v28 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v18))))
	v32 = v28 + int64(4354685564936845354)
	v33 = int64(30)
	v36 = int64(-4658895280553007687)
	v37 = (int64(base.Ui64(v32)>>(uint(v33)%64)) ^ v32) * v36
	v38 = int64(27)
	v41 = int64(-7723592293110705685)
	v42 = (int64(base.Ui64(v37)>>(uint(v38)%64)) ^ v37) * v41
	v43 = int64(31)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(base.Ui64(v42)>>(uint(v43)%64)) ^ v42
	v48 = v28 - int64(7046029254386353131)
	v53 = (int64(base.Ui64(v48)>>(uint(v33)%64)) ^ v48) * v36
	v58 = (int64(base.Ui64(v53)>>(uint(v38)%64)) ^ v53) * v41
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(base.Ui64(v58)>>(uint(v43)%64)) ^ v58
	if v48|v32 == int64(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = l1
	goto L16
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(6364136223846793005)
	goto L13
L12:
	;
	goto L13
L13:
	;
	goto L10
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v87
	if base.Ui32(int32(2)) <= base.Ui32(v77) {
		goto L22
	} else {
		goto L23
	}
L15:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v78 = base.F64_convert_i32_u(v77)
	v79 = base.F64_mul(v74, v78)
	if base.F64_lt(v79, float64(4.294967296e+09))&base.F64_ge(v79, float64(0)) != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v74 = F_pg_prng_double(m, v16)
	mBase = m.M
	if base.F64_eq(v74, float64(0)) != 0 {
		goto L16
	} else {
		goto L18
	}
L17:
	;
	goto L15
L18:
	;
	goto L17
L19:
	;
	v85 = base.I32_trunc_f64_u(v79)
	v87 = v85
	goto L14
L20:
	;
	goto L21
L21:
	;
	v87 = int32(0)
	goto L14
L22:
	;
	goto L25
L23:
	;
	v151 = int32(1)
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v151
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v162 = v158
	goto L8
L25:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v106 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v151 = v125
	goto L24
L27:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	goto L34
L30:
	;
	return int32(0)
L31:
	;
	goto L29
L32:
	;
	if v125 == int32(0) {
		goto L25
	} else {
		goto L40
	}
L33:
	;
	v117 = base.F64_mul(v114, v78)
	if base.F64_lt(v117, float64(4.294967296e+09))&base.F64_ge(v117, float64(0)) != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v114 = F_pg_prng_double(m, v16)
	mBase = m.M
	if base.F64_eq(v114, float64(0)) != 0 {
		goto L34
	} else {
		goto L36
	}
L35:
	;
	goto L33
L36:
	;
	goto L35
L37:
	;
	v123 = base.I32_trunc_f64_u(v117)
	v125 = v123
	goto L32
L38:
	;
	goto L39
L39:
	;
	v125 = int32(0)
	goto L32
L40:
	;
	__phi131 = v125
	__phi133 = v77
	v131 = __phi131
	v133 = __phi133
	goto L41
L41:
	;
	v141 = base.I32_rem_u_s(v133, v131)
	if v141 != 0 {
		__phi131 = v141
		__phi133 = v131
		v131 = __phi131
		v133 = __phi133
		goto L41
	} else {
		goto L43
	}
L42:
	;
	if base.Ui32(int32(1)) < base.Ui32(v131) {
		goto L25
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	goto L26
L45:
	;
	v193 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	if v194 <= v193 {
		v222 = v190
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v198 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v18)+36)))
	v202 = v196
	goto L47
L47:
	;
	v214 = base.I64_rem_u_s(v198+base.I64_extend_i32_u(v202), base.I64_extend_i32_u(v191))
	v215 = base.I32_wrap_i64(v214)
	if base.Ui32(l1) <= base.Ui32(v215) {
		v202 = v215
		goto L47
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v215
	v222 = v215
	goto L1
L49:
	;
	goto L48
}
