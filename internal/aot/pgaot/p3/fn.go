package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func Fn14205(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v13 = F_mul_size(m, v12, l3)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = F_add_size(m, l5, v13)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v15
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
			F_ShmemRequestStructWithOpts(m, v10)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				m.G0 = v10 + int32(16)
				return
			}
		}
	}
}
func Fn14209(m *base.Module, l0 float32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	if base.Ui32(base.I32_reinterpret_f32(l0)&int32(2147483647)) < base.Ui32(int32(2139095041)) {
		if base.F32_eq(base.F32_abs(l0), math.Float32frombits(uint32(0x7f800000))) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				F_errcode(m, int32(130))
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					F_errmsg(m, l3, int32(0))
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						F_errfinish(m, l2, l1, int32(_a_Fn14209_0))
						v40 = m.ExcPending
						if v40 != 0 {
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
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			F_errcode(m, int32(130))
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_errmsg(m, l5, int32(0))
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					F_errfinish(m, l2, l4, int32(_a_Fn14209_0))
					v27 = m.ExcPending
					if v27 != 0 {
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
func Fn14216(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v1 = l0
	v6 = F_mul_size(m, l1, v1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_add_size(m, int32(8), v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = F_palloc0(m, v10)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)) = uint16(v1)
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v10 << (uint(int32(2)) % 32)
				return v12
			}
		}
	}
}
func Fn14218(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v17 = F_SearchSysCacheExists(m, l6, base.I64_extend_i32_u(l1), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l2), int64(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		if v17 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_errcode(m, int32(_a_Fn14218_0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = F_get_am_name(m, l1)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v28 = F_get_namespace_name(m, l2)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v28
							*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v26
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
							F_errmsg(m, l5, v11)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_Fn14218_1), l4, l3)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
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
		} else {
			m.G0 = v11 + int32(16)
			return
		}
	}
}
func Fn14221(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v1 = l0
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v1)
	v10 = F_psprintf(m, l1, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v10
	}
}
func Fn14230(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v11 < int32(20) {
		v15 = v11 << (uint(int32(4)) % 32)
		*(*int64)(unsafe.Add(mBase, uint32(v15+l7))) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v15+l6))) = l0
		*(*int32)(unsafe.Add(mBase, uint32(l5))) = v11 + int32(1)
		v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_Fn14230[0])))
		if v24 == int32(0) {
			v29 = *(*int32)(unsafe.Add(mBase, _c_Fn14230[1]))
			if v29 <= int32(31) {
				*(*int32)(unsafe.Add(mBase, _c_Fn14230[1])) = v29 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v29<<(uint(int32(2))%32))+uint32(_c_Fn14230[2]))) = int32(1196)
			} else {
			}
			v46 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_Fn14230[0])) = uint8(v46)
		} else {
		}
		return
	} else {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				F_errmsg_internal(m, l4, int32(0))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_Fn14230_0), l3, l2)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
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
func Fn14236(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = l1
	return int64(0)
}
func Fn14245(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v36 int64
	_ = v36
	var v43 int64
	_ = v43
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = int64(63)
	v15 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+40)))
	v22 = int64(32)
	v23 = int64(base.Ui64(v15) >> (uint(v22) % 64))
	v25 = int64(base.Ui64(v12) >> (uint(v22) % 64))
	v28 = int64(4294967295)
	v29 = v15 & v28
	v31 = v12 & v28
	v32 = v29 * v31
	v36 = int64(base.Ui64(v32)>>(uint(v22)%64)) + v29*v25
	v43 = v31*v23 + v36&v28
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v12*(v15>>(uint(v13)%64)) + v12>>(uint(v13)%64)*v15 + v23*v25 + int64(base.Ui64(v36)>>(uint(v22)%64)) + int64(base.Ui64(v43)>>(uint(v22)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v32&v28 | v43<<(uint(v22)%64)
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	if v54 != v55>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, l4, int32(0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, l3, l2, l1)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.G0 = v10 + int32(16)
		return v55
	}
}
func Fn14252(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	v5 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+40)))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = base.I32_wrap_i64(v6)
	if v7 == int32(-2147483648) {
		v40 = int64(-9223372036854775807 - 1)
		v42 = F_DirectFunctionCall2Coll(m, l1, int32(0), v40, v5)
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int64(0)
		} else {
			return v42
		}
	} else {
		if v7 == int32(2147483647) {
			v40 = int64(9223372036854775807)
			v42 = F_DirectFunctionCall2Coll(m, l1, int32(0), v40, v5)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int64(0)
			} else {
				return v42
			}
		} else {
			if int32(106751983) <= v7 {
				v16 = int64(9223372036854775807)
				v18 = F_errsave_start(m, int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					if v18 == int32(0) {
						v40 = v16
						v42 = F_DirectFunctionCall2Coll(m, l1, int32(0), v40, v5)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int64(0)
						} else {
							return v42
						}
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_Fn14252_0), int32(0))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, int32(0), int32(_a_Fn14252_1), int32(646), int32(_a_Fn14252_2))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int64(0)
								} else {
									v40 = v16
									v42 = F_DirectFunctionCall2Coll(m, l1, int32(0), v40, v5)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return int64(0)
									} else {
										return v42
									}
								}
							}
						}
					}
				}
			} else {
				v40 = base.I64_extend32_s(v6) * int64(86400000000)
				v42 = F_DirectFunctionCall2Coll(m, l1, int32(0), v40, v5)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int64(0)
				} else {
					return v42
				}
			}
		}
	}
}
func Fn14258(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	if v11 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v23
				F_errmsg(m, l3, v8)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn14258_0), l2, l1)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
		m.G0 = v8 + int32(16)
		return int32(0)
	}
}
func Fn14265(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	v4 = l0
	goto L1
L1:
	;
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v7 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return v14
L3:
	;
	goto L2
L4:
	;
	v14 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	if v7 == l1 {
		v14 = v4
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v4 = v4 + int32(1)
	goto L1
}
func Fn14269(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = F_gbt_var_same(m, v6, v7, v8, l1, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v4)))) = uint8(v10)
		return v4 & int64(4294967295)
	}
}
func Fn14272(m *base.Module, l0 int32, l1 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = F_gbt_num_picksplit(m, v3, v4, l1, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v6)
	}
}
func Fn14274(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_palloc(m, l2)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = F_gbt_num_union(m, v7, v6, l1, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v14)
		}
	}
}
func Fn14278(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+12)) = uint32(v11)
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+10)) = uint16(v13)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+16)))
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+v28)+12)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = F_gbt_num_consistent(m, v8, v8+int32(12), v8+int32(10), v30&int32(1), l1, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return int64(0)
	} else {
		m.G0 = v8 + int32(16)
		return base.I64_extend_i32_u(v34)
	}
}
func Fn14281(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v11 = F_DirectFunctionCall2Coll(m, l4, int32(0), base.I64_extend_i32_u(v7), base.I64_extend_i32_u(v9))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = base.I32_wrap_i64(v11)
		if v15 != 0 {
			v24 = v15
			return v24
		} else {
			v21 = F_DirectFunctionCall2Coll(m, l4, int32(0), base.I64_extend_i32_u(v7+l3), base.I64_extend_i32_u(v9+l3))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v24 = base.I32_wrap_i64(v21)
				return v24
			}
		}
	}
}
func Fn14290(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = int64(0)
	v18 = F_GetSysCacheOid(m, l7, base.I64_extend_i32_u(l0), v15, v15, v15)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if l1|v18 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				F_errcode(m, l6)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
					F_errmsg(m, l5, v12)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, l4, l3, l2)
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
		} else {
			m.G0 = v12 + int32(16)
			return v18
		}
	}
}
func Fn14306(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v23 = F_pg_detoast_datum_packed(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, _c_Fn14306[0]))
			v28 = F_text_to_cstring(m, v18)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				v31 = F_DirectFunctionCall1Coll(m, l7, int32(0), base.I64_extend_i32_u(v28))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					v33 = base.I32_wrap_i64(v31)
					if v33 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							F_errcode(m, l6)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15))) = v28
								F_errmsg(m, l5, v15)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_Fn14306_0), l4, l3)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v48 = F_convert_any_priv_string(m, v23, l1)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int64(0)
						} else {
							v50 = F_object_aclcheck(m, l2, v33, v26, v48)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int64(0)
							} else {
								m.G0 = v15 + int32(16)
								return base.I64_extend_i32_u(base.B2i32(v50 == int32(0)))
							}
						}
					}
				}
			}
		}
	}
}
func Fn14308(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	v5 = int32(1)
	if l0 == l1 {
		v54 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v54
L2:
	;
	v7 = F_superuser_arg(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v7 != 0 {
		v54 = v5
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v11 = int32(0)
	v13 = F_roles_is_member_of(m, l0, l2, v11, v11)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v15 = int32(0)
	if v13 == v15 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v54 = v53
	goto L1
L8:
	;
	v53 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v21 <= int32(0) {
		v47 = v15
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = v47
	goto L7
L12:
	;
	v24 = int32(0)
	if v24 < v21 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v27 = v21
	goto L15
L14:
	;
	v27 = v24
	goto L15
L15:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v30 = int32(0)
	goto L16
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28+v30<<(uint(int32(2))%32))))
	v39 = base.B2i32(v38 == l1)
	if v38 == l1 {
		v47 = v39
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v47 = v39
	goto L11
L18:
	;
	v41 = v30 + int32(1)
	if v41 != v27 {
		v30 = v41
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
}
func Fn14313(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int64) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v31 int32
	_ = v31
	var v35 int64
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	v8 = int32(0)
	v9 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+12)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v12)+17)) = v9
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+72)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = int64(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+56)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = l6
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+40)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v14
	v31 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+26)) = uint16(v31)
	v35 = F_array_recv(m, v12+int32(8))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		return int64(0)
	} else {
		v39 = base.I32_wrap_i64(v35)
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
		if v40 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50462850))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, l4, int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, l3, l2, l1)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
			if v43 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50462850))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, l4, int32(0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, l3, l2, l1)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
				if v44 != l5 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50462850))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, l4, int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, l3, l2, l1)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
					if v46 == int32(0) {
						m.G0 = v12 + int32(80)
						return v35 & int64(4294967295)
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50462850))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, l4, int32(0))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, l3, l2, l1)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int64(0)
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
}
func Fn14319(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v18 = F_text_to_cstring(m, v16)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v18
				v25 = F_get_worker(m, v11, v8+int32(12), int32(0), int32(1), l1)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					if v25 != 0 {
						v31 = base.I64_extend_i32_u(v25)
					} else {
						v28 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
						v31 = int64(0)
					}
					m.G0 = v8 + int32(16)
					return v31
				}
			}
		}
	}
}
func Fn14320(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v9 == int32(1) {
		v12 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
		v34 = int64(0)
		m.G0 = v7 + int32(32)
		return v34
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v16
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v18
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v20
		F_pushJsonbValue(m, v7+int32(8), l1, int32(0))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int64(0)
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			v30 = F_JsonbValueToJsonb(m, v29)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int64(0)
			} else {
				v34 = base.I64_extend_i32_u(v30)
				m.G0 = v7 + int32(32)
				return v34
			}
		}
	}
}
func Fn14326(m *base.Module, l0 int32, l1 int64) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v56 int64
	_ = v56
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = int32(16711935)
	v10 = int32(8)
	v12 = int32(24)
	v16 = base.I32_rotr(v7&v8, v10) | base.I32_rotr(v7, v12)&v8
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v27 = base.I32_rotr(v18&v8, v10) | base.I32_rotr(v18, v12)&v8
	if base.Ui32(v16) < base.Ui32(v27) {
		v56 = l1
	} else {
		if base.Ui32(v27) < base.Ui32(v16) {
			v56 = int64(1)
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v32 = int32(16711935)
			v34 = int32(8)
			v36 = int32(24)
			v40 = base.I32_rotr(v31&v32, v34) | base.I32_rotr(v31, v36)&v32
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v50 = base.I32_rotr(v41&v32, v34) | base.I32_rotr(v41, v36)&v32
			if base.Ui32(v40) < base.Ui32(v50) {
				v56 = l1
			} else {
				v56 = base.I64_extend_i32_u(base.B2i32(base.Ui32(v50) < base.Ui32(v40)))
			}
		}
	}
	return v56
}
func Fn14328(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_palloc0(m, int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l3
		return v7
	}
}
func Fn14331(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = F_strlen(m, v6)
		mBase = m.M
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v14 = F_RE_compile_and_cache(m, v8, l1, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v19 = F_palloc_mul(m, int32(4), v12+int32(1))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				v21 = F_pg_mb2wchar_with_len(m, v6, v19, v12)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					v23 = int32(0)
					v26 = F_RE_wchar_execute(m, v19, v21, v23, v23, v23)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						F_pfree(m, v19)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v26 ^ int32(1))
						}
					}
				}
			}
		}
	}
}
func Fn14337(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_palloc(m, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		*(*uint32)(unsafe.Add(mBase, uint32(v7))) = uint32(v9)
		v16 = F_pg_snprintf(m, v10, l1, int32(_a_Fn14337_0), v7)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			m.G0 = v7 + int32(16)
			return base.I64_extend_i32_u(v10)
		}
	}
}
func Fn14342(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(34209794)
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+8)) = uint32(v11)
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+4)) = uint32(v10)
	v17 = *(*int32)(unsafe.Add(mBase, _c_Fn14342[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v17
	v20 = F_LockRelease(m, v8, l1, int32(1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int64(0)
	} else {
		m.G0 = v8 + int32(16)
		return base.I64_extend_i32_u(v20)
	}
}
func Fn14348(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
		if v19 == int32(_a_Fn14348_0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, l4, int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_Fn14348_1), l3, l2)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = v13
			v37 = v11 + int32(16)
			v40 = F_pg_snprintf(m, v37, int32(32), int32(_a_Fn14348_2), v11)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int64(0)
			} else {
				v43 = int32(0)
				v50 = F_DirectFunctionCall3Coll(m, int32(434), v43, base.I64_extend_i32_u(v37), int64(0), int64(-1))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int64(0)
				} else {
					v53 = F_DirectFunctionCall2Coll(m, l1, v43, v50, base.I64_extend_i32_u(v15))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int64(0)
					} else {
						v55 = F_DirectFunctionCall1Coll(m, int32(1670), v43, v53)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							m.G0 = v11 + int32(48)
							return v55
						}
					}
				}
			}
		}
	}
}
func Fn14357(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if l3 < v12 {
		m.G0 = v9 + int32(16)
		return int32(0)
	} else {
		v14 = F_strlen(m, l1)
		mBase = m.M
		if base.Ui32(int32(63)) < base.Ui32(v14) {
			m.G0 = v9 + int32(16)
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v21 = F_hash_search(m, v17, l1, int32(1), v9+int32(15))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v25
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v27 != 0 {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
					v30 = v29 - v27
					v33 = F_palloc(m, v30+int32(1))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						if v30 != 0 {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							base.MemoryCopy(m, v33, v35, v30)
						} else {
						}
						v38 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v33+v30))) = uint8(v38)
						v41 = v33
						*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v41
						m.G0 = v9 + int32(16)
						return int32(0)
					}
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v41 = v40
					*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v41
					m.G0 = v9 + int32(16)
					return int32(0)
				}
			}
		}
	}
}
func Fn14359(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v4
	v17 = F_query_or_expression_tree_walker_impl(m, l1, int32(947), v7+int32(4), v4)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		m.G0 = v7 + int32(16)
		return v21
	}
}
func Fn14362(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	v5 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6-v11 < l3 {
		v21 = v5
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v17 = F_memcmp(m, v14+v6-l3, l2, l3)
		mBase = m.M
		if v17 != 0 {
			v21 = v5
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 - l3
			v21 = int32(1)
		}
	}
	if v21 == int32(0) {
		v116 = v5
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v24
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v27 = int32(0)
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v26-int32(4))))
		if v34 == v27 {
			v108 = int32(0)
		} else {
			v39 = v34 & int32(3)
			if base.Ui32(v34) < base.Ui32(int32(4)) {
				v75 = v26
				v76 = int32(0)
				v81 = v75
				v82 = v76
				v86 = v27
				for {
					v87 = int32(*(*int8)(unsafe.Add(mBase, uint32(v81))))
					v90 = v82 + base.B2i32(int32(-65) < v87)
					v91 = int32(1)
					v94 = v86 + v91
					if v94 != v39 {
						v81 = v81 + v91
						v82 = v90
						v86 = v94
						continue
					} else {
						break
					}
					break
				}
				v97 = v90
			} else {
				v46 = v26
				v47 = int32(0)
				v50 = v27
				for {
					v52 = int32(*(*int8)(unsafe.Add(mBase, uint32(v46))))
					v53 = int32(-65)
					v56 = int32(*(*int8)(unsafe.Add(mBase, uint32(v46)+1)))
					v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(v46)+2)))
					v64 = int32(*(*int8)(unsafe.Add(mBase, uint32(v46)+3)))
					v67 = v47 + base.B2i32(v53 < v52) + base.B2i32(v53 < v56) + base.B2i32(v53 < v60) + base.B2i32(v53 < v64)
					v68 = int32(4)
					v69 = v46 + v68
					v71 = v50 + v68
					if v71 != v34&int32(-4) {
						v46 = v69
						v47 = v67
						v50 = v71
						continue
					} else {
						break
					}
					break
				}
				if v39 == int32(0) {
					v97 = v67
				} else {
					v75 = v69
					v76 = v67
					v81 = v75
					v82 = v76
					v86 = v27
					for {
						v87 = int32(*(*int8)(unsafe.Add(mBase, uint32(v81))))
						v90 = v82 + base.B2i32(int32(-65) < v87)
						v91 = int32(1)
						v94 = v86 + v91
						if v94 != v39 {
							v81 = v81 + v91
							v82 = v90
							v86 = v94
							continue
						} else {
							break
						}
						break
					}
					v97 = v90
				}
			}
			v108 = v97
		}
		if v108 < l1 {
			v116 = v5
		} else {
			v111 = F_slice_del(m, l0)
			mBase = m.M
			if int32(0) <= v111 {
				v114 = int32(1)
			} else {
				v114 = v111
			}
			v116 = v114
		}
	}
	return v116
}
func Fn14364(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v223 int32
	_ = v223
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 <= v12 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	return v378
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v367
	v378 = v372
	goto L1
L3:
	;
	v367 = v365
	v372 = int32(1)
	goto L2
L4:
	;
	v365 = v20 - v9 + v151
	goto L3
L5:
	;
	v159 = v11 - v9
	v160 = v156 + v159
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v160
	if v160 <= v157 {
		goto L34
	} else {
		goto L35
	}
L6:
	;
	v156 = v9
	v157 = v12
	v158 = v10
	goto L5
L7:
	;
	goto L8
L8:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v11-int32(1)))))
	if v17 != l1 {
		v156 = v9
		v157 = v12
		v158 = v10
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v20 = v11 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L12
L10:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v150 == int32(0) {
		goto L4
	} else {
		goto L33
	}
L11:
	;
	v150 = v143
	goto L10
L12:
	;
	if v20 <= v35 {
		v143 = int32(-1)
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v143 = int32(0)
	goto L11
L14:
	;
	v52 = int32(1)
	v53 = v20 - v52
	v55 = int32(*(*int8)(unsafe.Add(mBase, uint32(v36+v53))))
	v57 = v55 & int32(255)
	if base.B2i32(v53 == v35)|base.B2i32(int32(0) <= v55) != 0 {
		v115 = v57
		v119 = v52
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if int32(305) < v115 {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	v64 = v57 & int32(63)
	v66 = v20 - int32(2)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v66))))
	v70 = v68 << (uint(int32(6)) % 32)
	if base.B2i32(v66 != v35)&base.B2i32(base.Ui32(v68) < base.Ui32(int32(192))) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v115 = v70&int32(1984) | v64
	v119 = int32(2)
	goto L15
L18:
	;
	goto L19
L19:
	;
	v83 = v70&int32(4032) | v64
	v85 = v20 - int32(3)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v85))))
	if base.B2i32(v85 != v35)&base.B2i32(base.Ui32(v87) < base.Ui32(int32(224))) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v115 = v87<<(uint(int32(12))%32)&int32(_a_Fn14364_0) | v83
	v119 = int32(3)
	goto L15
L21:
	;
	goto L22
L22:
	;
	v105 = int32(4)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v36-v105))))
	v115 = v87<<(uint(int32(12))%32)&int32(_a_Fn14364_1) | v107&int32(7)<<(uint(int32(18))%32) | v83
	v119 = v105
	goto L15
L23:
	;
	v150 = v119
	goto L10
L24:
	;
	goto L25
L25:
	;
	v121 = v115 - int32(97)
	if v121 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v150 = v119
	goto L10
L27:
	;
	goto L28
L28:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v121)>>(uint(int32(3))%32)))+uint32(_c_Fn14364[0]))))
	if int32(base.Ui32(v127)>>(uint(v121&int32(7))%32))&int32(1) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v150 = v119
	goto L10
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20 - v119
	goto L32
L32:
	;
	goto L13
L33:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v156 = v151
	v157 = v155
	v158 = v154
	goto L5
L34:
	;
	v171 = int32(0)
	goto L39
L35:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160+v158-int32(1)))))
	if v166 != l1 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v367 = v160 - int32(1)
	v372 = int32(0)
	goto L2
L37:
	;
	if v223 < int32(0) {
		v378 = v171
		goto L1
	} else {
		goto L56
	}
L39:
	;
	goto L40
L40:
	;
	goto L41
L41:
	;
	v178 = v160
	v180 = int32(1)
	goto L44
L43:
	;
	v223 = v205
	goto L37
L44:
	;
	if v178 <= v157 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L43
L46:
	;
	v223 = int32(-1)
	goto L37
L47:
	;
	goto L48
L48:
	;
	v185 = v178 - int32(1)
	v187 = int32(*(*int8)(unsafe.Add(mBase, uint32(v158+v185))))
	if base.B2i32(int32(0) <= v187)|base.B2i32(v185 <= v157) != 0 {
		v205 = v185
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v209 = int32(1)
	if v209 < v180 {
		v178 = v205
		v180 = v180 - v209
		goto L44
	} else {
		goto L55
	}
L50:
	;
	v193 = v185
	goto L51
L51:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158+v193))))
	if base.Ui32(int32(191)) < base.Ui32(v198) {
		v205 = v193
		goto L49
	} else {
		goto L53
	}
L52:
	;
	v205 = v157
	goto L49
L53:
	;
	v202 = v193 - int32(1)
	if v157 < v202 {
		v193 = v202
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	goto L45
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v223
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L59
L57:
	;
	if v355 != 0 {
		v378 = v171
		goto L1
	} else {
		goto L80
	}
L58:
	;
	v355 = v348
	goto L57
L59:
	;
	if v223 <= v240 {
		v348 = int32(-1)
		goto L58
	} else {
		goto L61
	}
L60:
	;
	v348 = int32(0)
	goto L58
L61:
	;
	v257 = int32(1)
	v258 = v223 - v257
	v260 = int32(*(*int8)(unsafe.Add(mBase, uint32(v241+v258))))
	v262 = v260 & int32(255)
	if base.B2i32(v258 == v240)|base.B2i32(int32(0) <= v260) != 0 {
		v320 = v262
		v324 = v257
		goto L62
	} else {
		goto L63
	}
L62:
	;
	if int32(305) < v320 {
		goto L70
	} else {
		goto L71
	}
L63:
	;
	v269 = v262 & int32(63)
	v271 = v223 - int32(2)
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241+v271))))
	v275 = v273 << (uint(int32(6)) % 32)
	if base.B2i32(v271 != v240)&base.B2i32(base.Ui32(v273) < base.Ui32(int32(192))) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v320 = v275&int32(1984) | v269
	v324 = int32(2)
	goto L62
L65:
	;
	goto L66
L66:
	;
	v288 = v275&int32(4032) | v269
	v290 = v223 - int32(3)
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241+v290))))
	if base.B2i32(v290 != v240)&base.B2i32(base.Ui32(v292) < base.Ui32(int32(224))) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v320 = v292<<(uint(int32(12))%32)&int32(_a_Fn14364_0) | v288
	v324 = int32(3)
	goto L62
L68:
	;
	goto L69
L69:
	;
	v310 = int32(4)
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223+v241-v310))))
	v320 = v292<<(uint(int32(12))%32)&int32(_a_Fn14364_1) | v312&int32(7)<<(uint(int32(18))%32) | v288
	v324 = v310
	goto L62
L70:
	;
	v355 = v324
	goto L57
L71:
	;
	goto L72
L72:
	;
	v326 = v320 - int32(97)
	if v326 < int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v355 = v324
	goto L57
L74:
	;
	goto L75
L75:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v326)>>(uint(int32(3))%32)))+uint32(_c_Fn14364[0]))))
	if int32(base.Ui32(v332)>>(uint(v326&int32(7))%32))&int32(1) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v355 = v324
	goto L57
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v223 - v324
	goto L79
L79:
	;
	goto L60
L80:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v365 = v356 + v159
	goto L3
}
func Fn14368(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = l1
	return int64(0)
}
func Fn14377(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	v4 = l3
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum_copy(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		if v17 != 0 {
			v18 = F_array_contains_nulls(m, v13)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				if v18 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(67108994))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_Fn14377_0), int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_Fn14377_1), l2, l1)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v23 = F_ArrayGetNItemsSafe(m, v20, v13+int32(16))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int64(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v4)
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
						if v26 != 0 {
							v34 = v26
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v34 = (v27<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						}
						F_isort(m, v34+v13, v23, v10+int32(15))
						mBase = m.M
						m.G0 = v10 + int32(16)
						return base.I64_extend_i32_u(v13)
					}
				}
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v23 = F_ArrayGetNItemsSafe(m, v20, v13+int32(16))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v4)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
				if v26 != 0 {
					v34 = v26
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v34 = (v27<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				}
				F_isort(m, v34+v13, v23, v10+int32(15))
				mBase = m.M
				m.G0 = v10 + int32(16)
				return base.I64_extend_i32_u(v13)
			}
		}
	}
}
func Fn14379(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v13 = F_pg_detoast_datum_packed(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v15 == int32(1) {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
				if v21 == int32(18) {
					v24 = int32(16)
				} else {
					v24 = int32(0)
				}
				if base.Ui32((v21-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v31 = int32(4)
				} else {
					v31 = v24
				}
				v44 = v31
			} else {
				v32 = int32(1)
				if v15&v32 != 0 {
					v44 = int32(base.Ui32(v15)>>(uint(v32)%32)) - v32
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v46 = F_RE_compile_and_cache(m, v13, l1, v45)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int64(0)
			} else {
				v51 = F_palloc_mul(m, int32(4), v44+int32(1))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int64(0)
				} else {
					v53 = int32(1)
					if v15&v53 != 0 {
						v57 = v53
					} else {
						v57 = int32(4)
					}
					v59 = F_pg_mb2wchar_with_len(m, v8+v57, v51, v44)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int64(0)
					} else {
						v61 = int32(0)
						v64 = F_RE_wchar_execute(m, v51, v59, v61, v61, v61)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int64(0)
						} else {
							F_pfree(m, v51)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v64 ^ int32(1))
							}
						}
					}
				}
			}
		}
	}
}
func Fn14380(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 float64
	_ = v16
	var v19 float64
	_ = v19
	var v22 float64
	_ = v22
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v37 int64
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v59 int64
	_ = v59
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	v9 = F_MemoryContextAllocZero(m, l0, int32(32))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l0
		v16 = float64(4.294967296e+09)
		v19 = base.F64_div(base.F64_convert_i32_u(l1), float64(0.9))
		if base.F64_ge(v19, v16) != 0 {
			v22 = v16
		} else {
			v22 = v19
		}
		v23 = base.I64_trunc_sat_f64_u(v22)
		if base.Ui64(v23) <= base.Ui64(int64(2)) {
			v26 = int64(2)
		} else {
			v26 = v23
		}
		v27 = int64(1)
		if v26&(v26-v27) == int64(0) {
			v37 = v26
		} else {
			v37 = v27 << (uint(int64(64)-base.I64_clz(v26)) % 64)
		}
		if base.Ui64(v37<<(uint(int64(3))%64)) < base.Ui64(int64(2147483647)) {
			v46 = F_MemoryContextAllocExtended(m, l0, base.I32_wrap_i64(v37)<<(uint(int32(3))%32), int32(5))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v46
				v49 = int64(1)
				if v37&(v37-v49) == int64(0) {
					v59 = v37
				} else {
					v59 = v49 << (uint(int64(64)-base.I64_clz(v37)) % 64)
				}
				if base.Ui64(int64(2147483647)) <= base.Ui64(v59<<(uint(int64(3))%64)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_Fn14380_0), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn14380_1), int32(332), l3)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v9))) = v59
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = base.I32_wrap_i64(v59) - int32(1)
					if v59 == int64(4294967296) {
						v76 = int32(-85899346)
					} else {
						v76 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v59), float64(0.9)))
					}
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v76
					return v9
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_Fn14380_0), int32(0))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn14380_1), int32(332), l3)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
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
func Fn14391(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v9 = F_getTSCurrentConfig(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			v13 = F_DirectFunctionCall2Coll(m, l1, int32(0), base.I64_extend_i32_u(v9), base.I64_extend_i32_u(v4))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				return v13
			}
		}
	}
}
func Fn14397(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_SearchSysCache1(m, l5, base.I64_extend_i32_u(l0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
				F_errmsg_internal(m, l4, v11)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l3, l2, l1)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v29+v30)+84))
			F_ReleaseCatCache(m, v14)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				m.G0 = v11 + int32(16)
				return v32
			}
		}
	}
}
func Fn14401(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+29)) = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v16
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v20
	v25 = F_DirectFunctionCall1Coll(m, int32(3591), int32(0), base.I64_extend_i32_u(v10))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return int64(0)
	} else {
		m.G0 = v10 + int32(48)
		return v25
	}
}
func Fn14410(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 float32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 float64
	_ = v95
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = int32(1)
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v21 = v19 & v17
			if v21 != 0 {
				v22 = v17
			} else {
				v22 = int32(4)
			}
			if v19 == int32(1) {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
				if v29 == int32(18) {
					v32 = int32(16)
				} else {
					v32 = int32(0)
				}
				if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v39 = int32(4)
				} else {
					v39 = v32
				}
				v50 = v39
			} else {
				v40 = int32(1)
				if v21 != 0 {
					v50 = int32(base.Ui32(v19)>>(uint(v40)%32)) - v40
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v51 = int32(1)
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			v55 = v53 & v51
			if v55 != 0 {
				v56 = v51
			} else {
				v56 = int32(4)
			}
			if v53 == int32(1) {
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
				if v63 == int32(18) {
					v66 = int32(16)
				} else {
					v66 = int32(0)
				}
				if base.Ui32((v63-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v73 = int32(4)
				} else {
					v73 = v66
				}
				v84 = v73
			} else {
				v74 = int32(1)
				if v55 != 0 {
					v84 = int32(base.Ui32(v53)>>(uint(v74)%32)) - v74
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v84 = int32(base.Ui32(v78)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v85 = F_calc_word_similarity(m, v22+v10, v50, v15+v56, v84, l2)
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return int64(0)
			} else {
				v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v87 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int64(0)
					} else {
						v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v91 != v15 {
							F_pfree(m, v15)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int64(0)
							} else {
								v95 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								return base.I64_extend_i32_u(base.F64_le(v95, base.F64_promote_f32(v85)))
							}
						} else {
							v95 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							return base.I64_extend_i32_u(base.F64_le(v95, base.F64_promote_f32(v85)))
						}
					}
				} else {
					v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v91 != v15 {
						F_pfree(m, v15)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int64(0)
						} else {
							v95 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							return base.I64_extend_i32_u(base.F64_le(v95, base.F64_promote_f32(v85)))
						}
					} else {
						v95 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
						return base.I64_extend_i32_u(base.F64_le(v95, base.F64_promote_f32(v85)))
					}
				}
			}
		}
	}
}
