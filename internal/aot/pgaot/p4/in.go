package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_in_grouping_b_U(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = v14
	goto L2
L1:
	;
	return v126
L2:
	;
	if v25 <= v15 {
		v126 = int32(-1)
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v126 = int32(0)
	goto L1
L4:
	;
	v32 = int32(1)
	v33 = v25 - v32
	v35 = int32(*(*int8)(unsafe.Add(mBase, uint32(v16+v33))))
	v37 = v35 & int32(255)
	if base.B2i32(v33 == v15)|base.B2i32(int32(0) <= v35) != 0 {
		v95 = v37
		v99 = v32
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l3 < v95 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v44 = v37 & int32(63)
	v46 = v25 - int32(2)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v46))))
	v50 = v48 << (uint(int32(6)) % 32)
	if base.B2i32(v46 != v15)&base.B2i32(base.Ui32(v48) < base.Ui32(int32(192))) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v95 = v50&int32(1984) | v44
	v99 = int32(2)
	goto L5
L8:
	;
	goto L9
L9:
	;
	v63 = v50&int32(4032) | v44
	v65 = v25 - int32(3)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v65))))
	if base.B2i32(v65 != v15)&base.B2i32(base.Ui32(v67) < base.Ui32(int32(224))) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v95 = v67<<(uint(int32(12))%32)&int32(_a_F_in_grouping_b_U_0) | v63
	v99 = int32(3)
	goto L5
L11:
	;
	goto L12
L12:
	;
	v85 = int32(4)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v16-v85))))
	v95 = v67<<(uint(int32(12))%32)&int32(_a_F_in_grouping_b_U_1) | v87&int32(7)<<(uint(int32(18))%32) | v63
	v99 = v85
	goto L5
L13:
	;
	return v99
L14:
	;
	goto L15
L15:
	;
	v102 = v95 - l2
	if v102 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return v99
L17:
	;
	goto L18
L18:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v102)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v109)>>(uint(v102&int32(7))%32))&int32(1) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return v99
L20:
	;
	goto L21
L21:
	;
	v118 = v25 - v99
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v118
	if l4 != 0 {
		v25 = v118
		goto L2
	} else {
		goto L22
	}
L22:
	;
	goto L3
}
func F_in_range_timestamp_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v45 int64
	_ = v45
	var v52 int64
	_ = v52
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v68 int64
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v22 = base.I64_extend_i32_s(v16)*int64(30) + base.I64_extend_i32_s(v20)
	v31 = int64(32)
	v32 = int64(20)
	v34 = int64(base.Ui64(v22) >> (uint(v31) % 64))
	v37 = int64(4294967295)
	v38 = int64(500654080)
	v40 = v22 & v37
	v41 = v38 * v40
	v45 = int64(base.Ui64(v41)>>(uint(v31)%64)) + v38*v34
	v52 = v40*v32 + v45&v37
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v22*int64(0) + v22>>(uint(int64(63))%64)*int64(86400000000) + v32*v34 + int64(base.Ui64(v45)>>(uint(v31)%64)) + int64(base.Ui64(v52)>>(uint(v31)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v41&v37 | v52<<(uint(v31)%64)
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	if int64(0) <= v63+v64>>(uint(int64(63))%64)+base.I64_extend_i32_u(base.B2i32(base.Ui64(v68+v64) < base.Ui64(v68))) {
		v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
		v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
		v81 = int32(2147483647)
		if base.B2i32(v16 != v81)|base.B2i32(v20 != v81)|base.B2i32(v64 != int64(9223372036854775807)) == int32(0) {
			if v76 != 0 {
				if v78 != int64(9223372036854775807) {
					v102 = int32(1268)
					v104 = F_Int64GetDatum(m, v78)
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						v108 = F_DirectFunctionCall2Coll(m, v102, int32(0), v104, v15)
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return int32(0)
						} else {
							v110 = *(*int64)(unsafe.Add(mBase, uint32(v108)))
							if v75 != 0 {
								v114 = base.B2i32(v80 <= v110)
							} else {
								v114 = base.B2i32(v110 <= v80)
							}
							m.G0 = v13 + int32(16)
							return v114
						}
					}
				} else {
					v114 = int32(1)
					m.G0 = v13 + int32(16)
					return v114
				}
			} else {
				if v78 != int64(-9223372036854775807-1) {
					v102 = int32(1267)
					v104 = F_Int64GetDatum(m, v78)
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						v108 = F_DirectFunctionCall2Coll(m, v102, int32(0), v104, v15)
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return int32(0)
						} else {
							v110 = *(*int64)(unsafe.Add(mBase, uint32(v108)))
							if v75 != 0 {
								v114 = base.B2i32(v80 <= v110)
							} else {
								v114 = base.B2i32(v110 <= v80)
							}
							m.G0 = v13 + int32(16)
							return v114
						}
					}
				} else {
					v114 = int32(1)
					m.G0 = v13 + int32(16)
					return v114
				}
			}
		} else {
			if v76 == int32(0) {
				v102 = int32(1267)
			} else {
				v102 = int32(1268)
			}
			v104 = F_Int64GetDatum(m, v78)
			mBase = m.M
			v107 = m.ExcPending
			if v107 != 0 {
				return int32(0)
			} else {
				v108 = F_DirectFunctionCall2Coll(m, v102, int32(0), v104, v15)
				mBase = m.M
				v109 = m.ExcPending
				if v109 != 0 {
					return int32(0)
				} else {
					v110 = *(*int64)(unsafe.Add(mBase, uint32(v108)))
					if v75 != 0 {
						v114 = base.B2i32(v80 <= v110)
					} else {
						v114 = base.B2i32(v110 <= v80)
					}
					m.G0 = v13 + int32(16)
					return v114
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v122 = m.ExcPending
		if v122 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v125 = m.ExcPending
			if v125 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_in_range_timestamp_interval_0), int32(0))
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_in_range_timestamp_interval_1), int32(3906), int32(_a_F_in_range_timestamp_interval_2))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
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
func F_in_range_timestamptz_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v45 int64
	_ = v45
	var v52 int64
	_ = v52
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v68 int64
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v110 int32
	_ = v110
	var v112 int64
	_ = v112
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v22 = base.I64_extend_i32_s(v16)*int64(30) + base.I64_extend_i32_s(v20)
	v31 = int64(32)
	v32 = int64(20)
	v34 = int64(base.Ui64(v22) >> (uint(v31) % 64))
	v37 = int64(4294967295)
	v38 = int64(500654080)
	v40 = v22 & v37
	v41 = v38 * v40
	v45 = int64(base.Ui64(v41)>>(uint(v31)%64)) + v38*v34
	v52 = v40*v32 + v45&v37
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v22*int64(0) + v22>>(uint(int64(63))%64)*int64(86400000000) + v32*v34 + int64(base.Ui64(v45)>>(uint(v31)%64)) + int64(base.Ui64(v52)>>(uint(v31)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v41&v37 | v52<<(uint(v31)%64)
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	if int64(0) <= v63+v64>>(uint(int64(63))%64)+base.I64_extend_i32_u(base.B2i32(base.Ui64(v68+v64) < base.Ui64(v68))) {
		v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
		v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
		v81 = int32(2147483647)
		if base.B2i32(v16 != v81)|base.B2i32(v20 != v81)|base.B2i32(v64 != int64(9223372036854775807)) == int32(0) {
			if v76 != 0 {
				if v78 != int64(9223372036854775807) {
					v100 = v13 + int32(16)
					F_interval_um_internal(m, v15, v100)
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return int32(0)
					} else {
						v106 = F_timestamptz_pl_interval_internal(m, v78, v100, int32(0))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							v112 = v106
							if v75 != 0 {
								v117 = base.B2i32(v80 <= v112)
							} else {
								v117 = base.B2i32(v112 <= v80)
							}
							m.G0 = v13 + int32(32)
							return v117
						}
					}
				} else {
					v117 = int32(1)
					m.G0 = v13 + int32(32)
					return v117
				}
			} else {
				if v78 != int64(-9223372036854775807-1) {
					v109 = F_timestamptz_pl_interval_internal(m, v78, v15, int32(0))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						v112 = v109
						if v75 != 0 {
							v117 = base.B2i32(v80 <= v112)
						} else {
							v117 = base.B2i32(v112 <= v80)
						}
						m.G0 = v13 + int32(32)
						return v117
					}
				} else {
					v117 = int32(1)
					m.G0 = v13 + int32(32)
					return v117
				}
			}
		} else {
			if v76 == int32(0) {
				v109 = F_timestamptz_pl_interval_internal(m, v78, v15, int32(0))
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int32(0)
				} else {
					v112 = v109
					if v75 != 0 {
						v117 = base.B2i32(v80 <= v112)
					} else {
						v117 = base.B2i32(v112 <= v80)
					}
					m.G0 = v13 + int32(32)
					return v117
				}
			} else {
				v100 = v13 + int32(16)
				F_interval_um_internal(m, v15, v100)
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return int32(0)
				} else {
					v106 = F_timestamptz_pl_interval_internal(m, v78, v100, int32(0))
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						v112 = v106
						if v75 != 0 {
							v117 = base.B2i32(v80 <= v112)
						} else {
							v117 = base.B2i32(v112 <= v80)
						}
						m.G0 = v13 + int32(32)
						return v117
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v125 = m.ExcPending
		if v125 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v128 = m.ExcPending
			if v128 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_in_range_timestamptz_interval_0), int32(0))
				mBase = m.M
				v132 = m.ExcPending
				if v132 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_in_range_timestamptz_interval_1), int32(3869), int32(_a_F_in_range_timestamptz_interval_2))
					mBase = m.M
					v137 = m.ExcPending
					if v137 != 0 {
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
