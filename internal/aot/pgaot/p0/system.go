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
	v2 = *(*int32)(unsafe.Add(mBase, _consts[175]))
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
			F_errmsg_internal(m, int32(493993), v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(520529), int32(239), int32(263423))
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
		v30 = *(*int32)(unsafe.Add(mBase, uint32((l0^int32(-1))<<(uint(int32(2))%32))+uint32(_consts[196])))
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
				F_errmsg(m, int32(594451), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					F_errfinish(m, int32(522098), int32(151), int32(298251))
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
					F_errmsg(m, int32(594451), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						F_errfinish(m, int32(522098), int32(151), int32(298251))
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
						F_errmsg(m, int32(594451), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							F_errfinish(m, int32(522098), int32(151), int32(298251))
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
	var v28 int32
	_ = v28
	var v35 float64
	_ = v35
	var v38 int32
	_ = v38
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 float64
	_ = v75
	var v78 float64
	_ = v78
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var __phi92 int32
	_ = __phi92
	var v94 int32
	_ = v94
	var __phi94 int32
	_ = __phi94
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v157 int32
	_ = v157
	var v159 int64
	_ = v159
	var v163 int32
	_ = v163
	var v175 int64
	_ = v175
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
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
	return v183
L2:
	;
	v183 = int32(-1)
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
	v138 = v19
	goto L5
L5:
	;
	v149 = v138 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v149
	v151 = int32(-1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if base.Ui32(v152) < base.Ui32(v149) {
		v183 = v151
		goto L1
	} else {
		goto L42
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
	v123 = v22
	goto L8
L8:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v133
	v138 = v123
	goto L5
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	F_pg_prng_seed(m, v16, base.I64_extend_i32_u(v28))
	mBase = m.M
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = l1
	goto L13
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v48
	if base.Ui32(int32(2)) <= base.Ui32(v38) {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v39 = base.F64_convert_i32_u(v38)
	v40 = base.F64_mul(v35, v39)
	if base.F64_lt(v40, float64(4.294967296e+09))&base.F64_ge(v40, float64(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v35 = F_pg_prng_double(m, v16)
	mBase = m.M
	if base.F64_eq(v35, float64(0)) != 0 {
		goto L13
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	goto L14
L16:
	;
	v46 = base.I32_trunc_f64_u(v40)
	v48 = v46
	goto L11
L17:
	;
	goto L18
L18:
	;
	v48 = int32(0)
	goto L11
L19:
	;
	goto L22
L20:
	;
	v112 = int32(1)
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v112
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v123 = v119
	goto L8
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v67 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v112 = v86
	goto L21
L24:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	goto L31
L27:
	;
	return int32(0)
L28:
	;
	goto L26
L29:
	;
	if v86 == int32(0) {
		goto L22
	} else {
		goto L37
	}
L30:
	;
	v78 = base.F64_mul(v75, v39)
	if base.F64_lt(v78, float64(4.294967296e+09))&base.F64_ge(v78, float64(0)) != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v75 = F_pg_prng_double(m, v16)
	mBase = m.M
	if base.F64_eq(v75, float64(0)) != 0 {
		goto L31
	} else {
		goto L33
	}
L32:
	;
	goto L30
L33:
	;
	goto L32
L34:
	;
	v84 = base.I32_trunc_f64_u(v78)
	v86 = v84
	goto L29
L35:
	;
	goto L36
L36:
	;
	v86 = int32(0)
	goto L29
L37:
	;
	__phi92 = v86
	__phi94 = v38
	v92 = __phi92
	v94 = __phi94
	goto L38
L38:
	;
	v102 = base.I32_rem_u_s(v94, v92)
	if v102 != 0 {
		__phi92 = v102
		__phi94 = v92
		v92 = __phi92
		v94 = __phi94
		goto L38
	} else {
		goto L40
	}
L39:
	;
	if base.Ui32(int32(1)) < base.Ui32(v92) {
		goto L22
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	goto L23
L42:
	;
	v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	if v155 <= v154 {
		v183 = v151
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v159 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v18)+36)))
	v163 = v157
	goto L44
L44:
	;
	v175 = base.I64_rem_u_s(v159+base.I64_extend_i32_u(v163), base.I64_extend_i32_u(v152))
	v176 = base.I32_wrap_i64(v175)
	if base.Ui32(l1) <= base.Ui32(v176) {
		v163 = v176
		goto L44
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v176
	v183 = v176
	goto L1
L46:
	;
	goto L45
}
func F_system_time_beginsamplescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	if base.F64_lt(v7, float64(0)) == int32(0) {
		if base.Ui64(base.I64_reinterpret_f64(v7)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
			v38 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v37)+28)) = v38
			*(*uint16)(unsafe.Add(mBase, uint32(v37)+24)) = uint16(v38)
			*(*float64)(unsafe.Add(mBase, uint32(v37)+8)) = v7
			*(*int32)(unsafe.Add(mBase, uint32(v37))) = l3
			return
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				F_errcode(m, int32(403177602))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_errmsg(m, int32(361063), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						F_errfinish(m, int32(523869), int32(201), int32(298300))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
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
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			F_errcode(m, int32(403177602))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				F_errmsg(m, int32(361063), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_errfinish(m, int32(523869), int32(201), int32(298300))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
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
func F_system_time_samplescangetsamplesize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 float64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v24 float64
	_ = v24
	var v35 float64
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v48 float64
	_ = v48
	var v50 float64
	_ = v50
	var v52 float64
	_ = v52
	var v60 float64
	_ = v60
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	var v70 int32
	_ = v70
	var v77 float64
	_ = v77
	var v79 float64
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 float64
	_ = v89
	var v97 float64
	_ = v97
	var v101 float64
	_ = v101
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = float64(1000)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = F_estimate_expression_value(m, l0, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		if v19 != int32(7) {
			v35 = v14
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+24)))
			if v22 != 0 {
				v35 = v14
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
				v24 = *(*float64)(unsafe.Add(mBase, uint32(v23)))
				if base.F64_lt(v24, float64(0)) == int32(0) {
					if base.Ui64(base.I64_reinterpret_f64(v24)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
						v35 = v24
					} else {
						v35 = float64(1000)
					}
				} else {
					v35 = float64(1000)
				}
			}
		}
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
		F_get_tablespace_page_costs(m, v36, v12+int32(8), int32(0))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return
		} else {
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
			v43 = base.F64_convert_i32_u(v42)
			v44 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
			if base.F64_gt(v44, float64(0)) != 0 {
				v48 = base.F64_div(v35, v44)
			} else {
				v48 = v35
			}
			if base.F64_gt(v48, v43) != 0 {
				v50 = v43
			} else {
				v50 = v48
			}
			v52 = float64(1e+100)
			if base.F64_gt(v50, v52) != 0 {
				v64 = v52
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v50)&int64(9223372036854775807)) {
					v64 = v52
				} else {
					v60 = float64(1)
					if base.F64_le(v50, v60) != 0 {
						v64 = v60
					} else {
						v64 = base.F64_nearest(v50)
					}
				}
			}
			v65 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
			if base.F64_gt(v65, float64(0)) == int32(0) {
				v77 = v64
			} else {
				v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
				if v70 == int32(0) {
					v77 = v64
				} else {
					v77 = base.F64_mul(v64, base.F64_div(v65, base.F64_convert_i32_u(v70)))
				}
			}
			if base.F64_gt(v77, v65) != 0 {
				v79 = v65
			} else {
				v79 = v77
			}
			if base.F64_lt(v64, float64(4.294967296e+09))&base.F64_ge(v64, float64(0)) != 0 {
				v85 = base.I32_trunc_f64_u(v64)
				v87 = v85
			} else {
				v87 = int32(0)
			}
			v89 = float64(1e+100)
			if base.F64_gt(v79, v89) != 0 {
				v101 = v89
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v79)&int64(9223372036854775807)) {
					v101 = v89
				} else {
					v97 = float64(1)
					if base.F64_le(v79, v97) != 0 {
						v101 = v97
					} else {
						v101 = base.F64_nearest(v79)
					}
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
			*(*float64)(unsafe.Add(mBase, uint32(l4))) = v101
			m.G0 = v12 + int32(16)
			return
		}
	}
}
