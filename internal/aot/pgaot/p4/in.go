package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_in_grouping_b_U(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = v18
	goto L2
L1:
	;
	return v126
L2:
	;
	if v28 <= v19 {
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
	v36 = int32(1)
	v37 = v28 - v36
	v39 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15+v37))))
	v41 = v39 & int32(255)
	if v37 == v19 {
		v96 = v41
		v97 = v36
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l3 < v96 {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	if int32(0) <= v39 {
		v96 = v41
		v97 = v36
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v47 = v41 & int32(63)
	v49 = v28 - int32(2)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v49))))
	v53 = v51 << (uint(int32(6)) % 32)
	if base.B2i32(v49 != v19)&base.B2i32(base.Ui32(v51) < base.Ui32(int32(192))) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v96 = v53&int32(1984) | v47
	v97 = int32(2)
	goto L5
L9:
	;
	goto L10
L10:
	;
	v66 = v53&int32(4032) | v47
	v68 = v28 - int32(3)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v68))))
	if base.B2i32(v68 != v19)&base.B2i32(base.Ui32(v70) < base.Ui32(int32(224))) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v96 = v70<<(uint(int32(12))%32)&int32(61440) | v66
	v97 = int32(3)
	goto L5
L12:
	;
	goto L13
L13:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+(v15-int32(4))))))
	v96 = v70<<(uint(int32(12))%32)&int32(258048) | v88&int32(7)<<(uint(int32(18))%32) | v66
	v97 = int32(4)
	goto L5
L14:
	;
	return v97
L15:
	;
	goto L16
L16:
	;
	v102 = v96 - l2
	if v102 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return v97
L18:
	;
	goto L19
L19:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v102)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v109)>>(uint(v102&int32(7))%32))&int32(1) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return v97
L21:
	;
	goto L22
L22:
	;
	v118 = v28 - v97
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v118
	if l4 != 0 {
		v28 = v118
		goto L2
	} else {
		goto L23
	}
L23:
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
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
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
		if v16 != int32(2147483647) {
			if v76 == int32(0) {
				v98 = int32(1282)
			} else {
				v98 = int32(1283)
			}
			v100 = F_Int64GetDatum(m, v78)
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return int32(0)
			} else {
				v104 = F_DirectFunctionCall2Coll(m, v98, int32(0), v100, v15)
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return int32(0)
				} else {
					v106 = *(*int64)(unsafe.Add(mBase, uint32(v104)))
					if v75 != 0 {
						v110 = base.B2i32(v80 <= v106)
					} else {
						v110 = base.B2i32(v106 <= v80)
					}
					m.G0 = v13 + int32(16)
					return v110
				}
			}
		} else {
			if v20 != int32(2147483647) {
				if v76 == int32(0) {
					v98 = int32(1282)
				} else {
					v98 = int32(1283)
				}
				v100 = F_Int64GetDatum(m, v78)
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					v104 = F_DirectFunctionCall2Coll(m, v98, int32(0), v100, v15)
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return int32(0)
					} else {
						v106 = *(*int64)(unsafe.Add(mBase, uint32(v104)))
						if v75 != 0 {
							v110 = base.B2i32(v80 <= v106)
						} else {
							v110 = base.B2i32(v106 <= v80)
						}
						m.G0 = v13 + int32(16)
						return v110
					}
				}
			} else {
				if v64 != int64(9223372036854775807) {
					if v76 == int32(0) {
						v98 = int32(1282)
					} else {
						v98 = int32(1283)
					}
					v100 = F_Int64GetDatum(m, v78)
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return int32(0)
					} else {
						v104 = F_DirectFunctionCall2Coll(m, v98, int32(0), v100, v15)
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return int32(0)
						} else {
							v106 = *(*int64)(unsafe.Add(mBase, uint32(v104)))
							if v75 != 0 {
								v110 = base.B2i32(v80 <= v106)
							} else {
								v110 = base.B2i32(v106 <= v80)
							}
							m.G0 = v13 + int32(16)
							return v110
						}
					}
				} else {
					if v76 != 0 {
						if v78 != int64(9223372036854775807) {
							v98 = int32(1283)
							v100 = F_Int64GetDatum(m, v78)
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return int32(0)
							} else {
								v104 = F_DirectFunctionCall2Coll(m, v98, int32(0), v100, v15)
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									v106 = *(*int64)(unsafe.Add(mBase, uint32(v104)))
									if v75 != 0 {
										v110 = base.B2i32(v80 <= v106)
									} else {
										v110 = base.B2i32(v106 <= v80)
									}
									m.G0 = v13 + int32(16)
									return v110
								}
							}
						} else {
							v110 = int32(1)
							m.G0 = v13 + int32(16)
							return v110
						}
					} else {
						if v78 != int64(-9223372036854775807-1) {
							v98 = int32(1282)
							v100 = F_Int64GetDatum(m, v78)
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return int32(0)
							} else {
								v104 = F_DirectFunctionCall2Coll(m, v98, int32(0), v100, v15)
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									v106 = *(*int64)(unsafe.Add(mBase, uint32(v104)))
									if v75 != 0 {
										v110 = base.B2i32(v80 <= v106)
									} else {
										v110 = base.B2i32(v106 <= v80)
									}
									m.G0 = v13 + int32(16)
									return v110
								}
							}
						} else {
							v110 = int32(1)
							m.G0 = v13 + int32(16)
							return v110
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v118 = m.ExcPending
		if v118 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v121 = m.ExcPending
			if v121 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(239981), int32(0))
				mBase = m.M
				v125 = m.ExcPending
				if v125 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(474326), int32(3906), int32(295081))
					mBase = m.M
					v130 = m.ExcPending
					if v130 != 0 {
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
	var v100 int32
	_ = v100
	var v104 int64
	_ = v104
	var v105 int32
	_ = v105
	var v107 int64
	_ = v107
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
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
		if v16 != int32(2147483647) {
			if v76 == int32(0) {
				v107 = F_timestamptz_pl_interval_internal(m, v78, v15, int32(0))
				mBase = m.M
				v108 = m.ExcPending
				if v108 != 0 {
					return int32(0)
				} else {
					v109 = v107
					if v75 != 0 {
						v113 = base.B2i32(v80 <= v109)
					} else {
						v113 = base.B2i32(v109 <= v80)
					}
					m.G0 = v13 + int32(32)
					return v113
				}
			} else {
				F_interval_um_internal(m, v15, v13+int32(16))
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return int32(0)
				} else {
					v104 = F_timestamptz_pl_interval_internal(m, v78, v13+int32(16), int32(0))
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return int32(0)
					} else {
						v109 = v104
						if v75 != 0 {
							v113 = base.B2i32(v80 <= v109)
						} else {
							v113 = base.B2i32(v109 <= v80)
						}
						m.G0 = v13 + int32(32)
						return v113
					}
				}
			}
		} else {
			if v20 != int32(2147483647) {
				if v76 == int32(0) {
					v107 = F_timestamptz_pl_interval_internal(m, v78, v15, int32(0))
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int32(0)
					} else {
						v109 = v107
						if v75 != 0 {
							v113 = base.B2i32(v80 <= v109)
						} else {
							v113 = base.B2i32(v109 <= v80)
						}
						m.G0 = v13 + int32(32)
						return v113
					}
				} else {
					F_interval_um_internal(m, v15, v13+int32(16))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						v104 = F_timestamptz_pl_interval_internal(m, v78, v13+int32(16), int32(0))
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return int32(0)
						} else {
							v109 = v104
							if v75 != 0 {
								v113 = base.B2i32(v80 <= v109)
							} else {
								v113 = base.B2i32(v109 <= v80)
							}
							m.G0 = v13 + int32(32)
							return v113
						}
					}
				}
			} else {
				if v64 != int64(9223372036854775807) {
					if v76 == int32(0) {
						v107 = F_timestamptz_pl_interval_internal(m, v78, v15, int32(0))
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int32(0)
						} else {
							v109 = v107
							if v75 != 0 {
								v113 = base.B2i32(v80 <= v109)
							} else {
								v113 = base.B2i32(v109 <= v80)
							}
							m.G0 = v13 + int32(32)
							return v113
						}
					} else {
						F_interval_um_internal(m, v15, v13+int32(16))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							v104 = F_timestamptz_pl_interval_internal(m, v78, v13+int32(16), int32(0))
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return int32(0)
							} else {
								v109 = v104
								if v75 != 0 {
									v113 = base.B2i32(v80 <= v109)
								} else {
									v113 = base.B2i32(v109 <= v80)
								}
								m.G0 = v13 + int32(32)
								return v113
							}
						}
					}
				} else {
					if v76 != 0 {
						if v78 != int64(9223372036854775807) {
							F_interval_um_internal(m, v15, v13+int32(16))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								v104 = F_timestamptz_pl_interval_internal(m, v78, v13+int32(16), int32(0))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									v109 = v104
									if v75 != 0 {
										v113 = base.B2i32(v80 <= v109)
									} else {
										v113 = base.B2i32(v109 <= v80)
									}
									m.G0 = v13 + int32(32)
									return v113
								}
							}
						} else {
							v113 = int32(1)
							m.G0 = v13 + int32(32)
							return v113
						}
					} else {
						if v78 != int64(-9223372036854775807-1) {
							v107 = F_timestamptz_pl_interval_internal(m, v78, v15, int32(0))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								v109 = v107
								if v75 != 0 {
									v113 = base.B2i32(v80 <= v109)
								} else {
									v113 = base.B2i32(v109 <= v80)
								}
								m.G0 = v13 + int32(32)
								return v113
							}
						} else {
							v113 = int32(1)
							m.G0 = v13 + int32(32)
							return v113
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v121 = m.ExcPending
		if v121 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v124 = m.ExcPending
			if v124 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(239981), int32(0))
				mBase = m.M
				v128 = m.ExcPending
				if v128 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(474326), int32(3869), int32(294872))
					mBase = m.M
					v133 = m.ExcPending
					if v133 != 0 {
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
