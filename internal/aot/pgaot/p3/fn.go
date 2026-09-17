package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func Fn13824(m *base.Module, l0 float32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
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
						F_errfinish(m, l2, l1, int32(_a_Fn13824_0))
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
					F_errfinish(m, l2, l4, int32(_a_Fn13824_0))
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
func Fn13828(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = F_SearchSysCache1(m, l5, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
				F_errmsg_internal(m, l4, v11)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l3, l2, l1)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v28+v29)+4))
			F_ReleaseCatCache(m, v13)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v11 + int32(16)
				return v31
			}
		}
	}
}
func Fn13837(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v4 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v6 = F_SearchCatCache2(m, v5, l0, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v11)+91)))
			if v13 == int32(0) {
				v16 = F_heap_copytuple(m, v6)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v18 = v16
					F_ReleaseCatCache(m, v6)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v22 = v18
						return v22
					}
				}
			} else {
				v18 = v4
				F_ReleaseCatCache(m, v6)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v22 = v18
					return v22
				}
			}
		} else {
			v22 = v4
			return v22
		}
	}
}
func Fn13840(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	F_errstart_cold(m, int32(21), int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(_a_Fn13840_0), int32(0))
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errdetail(m, int32(_a_Fn13840_1), int32(0))
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn13840_2), l2, l1)
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
		}
	}
}
func Fn13846(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_DatumGetAnyArrayP(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_DatumGetAnyArrayP(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = F_array_contain_compare(m, v7, v12, v14, l1, v15+int32(16))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				if v20 == int32(-1) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v27 == int32(-1) {
						return v18
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v12 == v30 {
							return v18
						} else {
							F_pfree(m, v12)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								return v18
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v7 == v23 {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						if v27 == int32(-1) {
							return v18
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v12 == v30 {
								return v18
							} else {
								F_pfree(m, v12)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									return v18
								}
							}
						}
					} else {
						F_pfree(m, v7)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							if v27 == int32(-1) {
								return v18
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v12 == v30 {
									return v18
								} else {
									F_pfree(m, v12)
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return int32(0)
									} else {
										return v18
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
func Fn13853(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	v4 = int32(_a_Fn13853_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_Fn13853[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	*(*int32)(unsafe.Add(mBase, _c_Fn13853[0])) = v8
	F_varstr_sortsupport(m, v7, l1, int32(950))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_Fn13853[0])) = v5
		return int32(0)
	}
}
func Fn13859(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v14 = v13 - v9
	if base.B2i32(int64(0) < v9) != base.B2i32(v14 < v13) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, l4, int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l3, l2, l1)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
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
		v31 = F_Int64GetDatum(m, v14)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			return v31
		}
	}
}
func Fn13860(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v37 int64
	_ = v37
	var v44 int64
	_ = v44
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v14 = int64(63)
	v16 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v23 = int64(32)
	v24 = int64(base.Ui64(v16) >> (uint(v23) % 64))
	v26 = int64(base.Ui64(v13) >> (uint(v23) % 64))
	v29 = int64(4294967295)
	v30 = v16 & v29
	v32 = v13 & v29
	v33 = v30 * v32
	v37 = int64(base.Ui64(v33)>>(uint(v23)%64)) + v30*v26
	v44 = v32*v24 + v37&v29
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v13*(v16>>(uint(v14)%64)) + v13>>(uint(v14)%64)*v16 + v24*v26 + int64(base.Ui64(v37)>>(uint(v23)%64)) + int64(base.Ui64(v44)>>(uint(v23)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v33&v29 | v44<<(uint(v23)%64)
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	if v55 != v56>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, l4, int32(0))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l3, l2, l1)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
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
		v74 = F_Int64GetDatum(m, v56)
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			m.G0 = v10 + int32(16)
			return v74
		}
	}
}
func Fn13871(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = F_get_fn_expr_argtype(m, v12, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v18 = F_enum_endpoint(m, v14, l4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = F_format_type_be(m, v14)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v47
								F_errmsg(m, int32(_a_Fn13871_0), v10)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_Fn13871_1), l2, l1)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
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
					m.G0 = v10 + int32(16)
					return v18
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_Fn13871_2), int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_Fn13871_1), l3, l1)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
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
func Fn13877(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = l9
	v18 = int32(1)
	v22 = F_LookupFuncName(m, l0, v18, v15+int32(44), v18)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		if v22 != 0 {
			v26 = F_get_func_rettype(m, v22)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 != l8 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(117833860))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							v83 = F_NameListToString(m, l0)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v83
								F_errmsg(m, l3, v15+int32(32))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_Fn13877_0), l2, l1)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
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
					v29 = F_func_volatile(m, v22)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						if v29 != int32(118) {
							m.G0 = v15 + int32(48)
							return v22
						} else {
							v35 = F_errstart(m, int32(19), int32(0))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								if v35 == int32(0) {
									m.G0 = v15 + int32(48)
									return v22
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return int32(0)
									} else {
										v42 = F_NameListToString(m, l0)
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v42
											F_errmsg(m, l7, v15+int32(16))
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_Fn13877_0), l6, l1)
												mBase = m.M
												v51 = m.ExcPending
												if v51 != 0 {
													return int32(0)
												} else {
													m.G0 = v15 + int32(48)
													return v22
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(52461700))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					v67 = F_func_signature_string(m, l0, int32(1), int32(0), v15+int32(44))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v67
						F_errmsg(m, int32(_a_Fn13877_1), v15)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn13877_0), l5, l1)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
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
func Fn13882(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = F_pg_detoast_datum(m, l1)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v19 = v10 + int32(8)
			v21 = int32(4)
			v22 = v12 + v21
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = v22
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			v25 = int32(2)
			v26 = int32(base.Ui32(v24) >> (uint(v25) % 32))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			if base.Ui32(v26+v21) < base.Ui32(int32(base.Ui32(v34)>>(uint(v25)%32))) {
				v38 = v22 + (v26+int32(3))&int32(2147483644)
			} else {
				v38 = v22
			}
			*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v38
			v41 = int32(4)
			v42 = v16 + v41
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v42
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
			v45 = int32(2)
			v46 = int32(base.Ui32(v44) >> (uint(v45) % 32))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if base.Ui32(v46+v41) < base.Ui32(int32(base.Ui32(v54)>>(uint(v45)%32))) {
				v58 = v42 + (v46+int32(3))&int32(2147483644)
			} else {
				v58 = v42
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v58
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v63 = F_DirectFunctionCall2Coll(m, l3, int32(0), v61, v62)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				if l0 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						if l1 != v16 {
							F_pfree(m, v16)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 + int32(16)
								return v63
							}
						} else {
							m.G0 = v10 + int32(16)
							return v63
						}
					}
				} else {
					if l1 != v16 {
						F_pfree(m, v16)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							m.G0 = v10 + int32(16)
							return v63
						}
					} else {
						m.G0 = v10 + int32(16)
						return v63
					}
				}
			}
		}
	}
}
func Fn13884(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_palloc(m, l2)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = F_gbt_num_union(m, v7, v6, l1, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v14
		}
	}
}
func Fn13888(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v11 int32
	_ = v11
	var v13 int32
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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+10)) = uint16(v13)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+16)))
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+v28)+12)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = F_gbt_num_consistent(m, v8, v8+int32(12), v8+int32(10), v30&int32(1), l1, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return v34
	}
}
func Fn13897(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, l1, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+8))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func Fn13899(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v9 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v17 = F_GetSysCacheOid(m, l7, l0, v9, v9, v9)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if l1|v17 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_errcode(m, l6)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
					F_errmsg(m, l5, v12)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, l4, l3, l2)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
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
			return v17
		}
	}
}
func Fn13907(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = F_palloc(m, int32(4))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v26 = F_palloc(m, int32(16))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = v28
			v31 = F_palloc(m, v28)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v31
				v34 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = l2
				*(*uint8)(unsafe.Add(mBase, uint32(v26)+8)) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v17
				*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v16)
				v42 = F_palloc(m, int32(4))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v42
					*(*int32)(unsafe.Add(mBase, uint32(v42))) = v26
					v47 = v16 & int32(_a_Fn13907_0)
					switch v47 - int32(1) {
					case 0, 1:
						v66 = F_palloc0(m, l1)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v21))) = v66
							v69 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v69)
							m.G0 = v13 + int32(16)
							return v21
						}
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v17
						m.G0 = v13 + int32(16)
						return v21
					case 3, 4:
						v50 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v50)
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v17
						m.G0 = v13 + int32(16)
						return v21
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13))) = v47
							F_errmsg_internal(m, int32(_a_Fn13907_1), v13)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn13907_2), int32(97), int32(_a_Fn13907_3))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
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
func Fn13912(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v23 = F_ArrayGetIntegerTypmods(m, v17, v14+int32(12))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
			if v25 == int32(1) {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				if v28 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, l7, int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, l3, l6, l1)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
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
					if base.Ui32(l9) <= base.Ui32(v28) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = l5
								F_errmsg(m, l4, v14)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, l3, l2, l1)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
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
						m.G0 = v14 + int32(16)
						return v28
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_Fn13912_0), int32(0))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, l3, l8, l1)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
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
func Fn13927(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	F_ean2string(m, v9, v6, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_pstrdup(m, v6)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(32)
			return v14
		}
	}
}
func Fn13932(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v10 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v87
L2:
	;
	v13 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
	v87 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = F_palloc(m, int32(32))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v73 = int32(0)
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v73
	v77 = F_pushJsonbValue(m, v8, l1, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L17
	}
L8:
	;
	return int32(0)
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v28
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v30
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v34
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+28)) = uint8(v36)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+29)))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+29)) = uint8(v38)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v40 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v41 = v24
	v43 = v40
	goto L13
L11:
	;
	v63 = v24
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+24)) = int32(0)
	v73 = v24
	goto L7
L13:
	;
	v47 = F_palloc(m, int32(32))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v63 = v47
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = v47
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v50
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v47)+8)) = v52
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	*(*int64)(unsafe.Add(mBase, uint32(v47))) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = v56
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+28)) = uint8(v58)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+29)))
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+29)) = uint8(v60)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
	if v62 != 0 {
		v41 = v47
		v43 = v62
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v77
	v80 = F_JsonbValueToJsonb(m, v77)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v87 = v80
	goto L1
}
func Fn13934(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v21 = F_JsonbExtractScalar(m, v13+int32(4), v10+int32(12))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			if v21 != 0 {
				switch v23 {
				case 0:
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v24 != v13 {
						F_pfree(m, v13)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
							v41 = int32(0)
							m.G0 = v10 + int32(32)
							return v41
						}
					} else {
						v28 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
						v41 = int32(0)
						m.G0 = v10 + int32(32)
						return v41
					}
				default:
					F_cannotCastJsonbValue(m, v23, l1)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 2:
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
					v35 = F_DirectFunctionCall1Coll(m, l2, int32(0), v34)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v13 == v37 {
							v41 = v35
							m.G0 = v10 + int32(32)
							return v41
						} else {
							F_pfree(m, v13)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = v35
								m.G0 = v10 + int32(32)
								return v41
							}
						}
					}
				}
			} else {
				F_cannotCastJsonbValue(m, v23, l1)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
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
func Fn13943(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
func Fn13956(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(34209793)
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+8)) = uint32(v11)
	v16 = int64(base.Ui64(v11) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+4)) = uint32(v16)
	v19 = *(*int32)(unsafe.Add(mBase, _c_Fn13956[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v19
	v22 = F_LockAcquire(m, v8, l2, l1, int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return int32(0)
	}
}
func Fn13961(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_SearchSysCache1(m, int32(47), v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
			v32 = int32(0)
			m.G0 = v7 + int32(16)
			return v32
		} else {
			F_initStringInfo(m, v7)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v23 = F_print_function_arguments(m, v7, v11, int32(0), l1)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_ReleaseCatCache(m, v11)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
						v28 = F_cstring_to_text(m, v27)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v27)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								v32 = v28
								m.G0 = v7 + int32(16)
								return v32
							}
						}
					}
				}
			}
		}
	}
}
func Fn13967(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(34209793)
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+8)) = uint32(v11)
	v16 = int64(base.Ui64(v11) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+4)) = uint32(v16)
	v19 = *(*int32)(unsafe.Add(mBase, _c_Fn13967[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v19
	v22 = F_LockAcquire(m, v8, l2, l1, int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return base.B2i32(v22 != int32(0))
	}
}
func Fn13970(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v16 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v20 = F_pg_detoast_datum_packed(m, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = v20
					v23 = F_encrypt_internal(m, l2, l1, v9, v14, v22)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v25 != v9 {
							F_pfree(m, v9)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v29 != v14 {
									F_pfree(m, v14)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return int32(0)
									} else {
										v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
										if v33 < int32(3) {
											return v23
										} else {
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v22 == v36 {
												return v23
											} else {
												F_pfree(m, v22)
												mBase = m.M
												v39 = m.ExcPending
												if v39 != 0 {
													return int32(0)
												} else {
													return v23
												}
											}
										}
									}
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v22 == v36 {
											return v23
										} else {
											F_pfree(m, v22)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							}
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v29 != v14 {
								F_pfree(m, v14)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v22 == v36 {
											return v23
										} else {
											F_pfree(m, v22)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v22 == v36 {
										return v23
									} else {
										F_pfree(m, v22)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						}
					}
				}
			} else {
				v22 = int32(0)
				v23 = F_encrypt_internal(m, l2, l1, v9, v14, v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v25 != v9 {
						F_pfree(m, v9)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v29 != v14 {
								F_pfree(m, v14)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
									if v33 < int32(3) {
										return v23
									} else {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v22 == v36 {
											return v23
										} else {
											F_pfree(m, v22)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												return v23
											}
										}
									}
								}
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v22 == v36 {
										return v23
									} else {
										F_pfree(m, v22)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v29 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
								if v33 < int32(3) {
									return v23
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v22 == v36 {
										return v23
									} else {
										F_pfree(m, v22)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											return v23
										}
									}
								}
							}
						} else {
							v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
							if v33 < int32(3) {
								return v23
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v22 == v36 {
									return v23
								} else {
									F_pfree(m, v22)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										return v23
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
func Fn13976(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
	if v7 <= l3 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
		switch v9 - int32(3) {
		case 0, 2:
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v15 = v14
		default:
			v15 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v15
	} else {
	}
	return int32(0)
}
func Fn13987(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
			if v22 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				if v23 == v20 {
					v33 = v22
					v34 = F_range_union_internal(m, v33, v13, v18, l1)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						m.G0 = v10 + int32(16)
						return v34
					}
				} else {
					v26 = F_lookup_type_cache(m, v20, int32(2048))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+200))
						if v28 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
								F_errmsg_internal(m, int32(_a_Fn13987_0), v10)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_Fn13987_1), int32(1776), int32(_a_Fn13987_2))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v26
							v33 = v26
							v34 = F_range_union_internal(m, v33, v13, v18, l1)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 + int32(16)
								return v34
							}
						}
					}
				}
			} else {
				v26 = F_lookup_type_cache(m, v20, int32(2048))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+200))
					if v28 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
							F_errmsg_internal(m, int32(_a_Fn13987_0), v10)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn13987_1), int32(1776), int32(_a_Fn13987_2))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v26
						v33 = v26
						v34 = F_range_union_internal(m, v33, v13, v18, l1)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							m.G0 = v10 + int32(16)
							return v34
						}
					}
				}
			}
		}
	}
}
func Fn13989(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 float64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_get_negator(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_Fn13989_0), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn13989_1), int32(773), int32(_a_Fn13989_2))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v33 = F_patternsel_common(m, v10, v12, int32(0), v9, v8, v7, l1, int32(1))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = F_Float8GetDatum(m, v33)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					return v35
				}
			}
		}
	}
}
func Fn13992(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	v12 = F_pg_snprintf(m, l0, int32(12), int32(_a_Fn13992_0), v6)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return l0
	}
}
func Fn13998(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v89 int64
	_ = v89
	var v96 int64
	_ = v96
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v112 int64
	_ = v112
	var v115 int64
	_ = v115
	var v120 int64
	_ = v120
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v130 int64
	_ = v130
	var v141 int64
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	if base.Ui64(v21-int64(9223372036854775807)) < base.Ui64(int64(2)) {
		v141 = v21
		v146 = F_Int64GetDatum(m, v141)
		mBase = m.M
		v147 = m.ExcPending
		if v147 != 0 {
			return int32(0)
		} else {
			m.G0 = v18 + int32(16)
			return v146
		}
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v27 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
		if base.Ui64(v27-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v155 = m.ExcPending
			if v155 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v158 = m.ExcPending
				if v158 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_Fn13998_0), int32(0))
					mBase = m.M
					v162 = m.ExcPending
					if v162 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_Fn13998_1), l7, l1)
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
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
			if v33 != 0 {
				if v33 != int32(2147483647) {
					if v33 != int32(-2147483648) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_Fn13998_2), int32(0))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_Fn13998_1), l8, l1)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
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
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
						if v38 != int32(-2147483648) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_Fn13998_2), int32(0))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_Fn13998_1), l8, l1)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
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
							v41 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
							if v41 != int64(-9223372036854775807-1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_Fn13998_2), int32(0))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_Fn13998_1), l8, l1)
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
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
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v227 = m.ExcPending
								if v227 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v230 = m.ExcPending
									if v230 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_Fn13998_3), int32(0))
										mBase = m.M
										v234 = m.ExcPending
										if v234 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_Fn13998_1), l2, l1)
											mBase = m.M
											v237 = m.ExcPending
											if v237 != 0 {
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
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
					if v44 != int32(2147483647) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_Fn13998_2), int32(0))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_Fn13998_1), l8, l1)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
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
						v47 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
						if v47 == int64(9223372036854775807) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v227 = m.ExcPending
							if v227 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v230 = m.ExcPending
								if v230 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_Fn13998_3), int32(0))
									mBase = m.M
									v234 = m.ExcPending
									if v234 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_Fn13998_1), l2, l1)
										mBase = m.M
										v237 = m.ExcPending
										if v237 != 0 {
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
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_Fn13998_2), int32(0))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_Fn13998_1), l8, l1)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
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
			} else {
				v66 = int64(*(*int32)(unsafe.Add(mBase, uint32(v32)+8)))
				v75 = int64(32)
				v76 = int64(20)
				v78 = int64(base.Ui64(v66) >> (uint(v75) % 64))
				v81 = int64(4294967295)
				v82 = int64(500654080)
				v84 = v66 & v81
				v85 = v82 * v84
				v89 = int64(base.Ui64(v85)>>(uint(v75)%64)) + v82*v78
				v96 = v84*v76 + v89&v81
				*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v66*int64(0) + v66>>(uint(int64(63))%64)*int64(86400000000) + v76*v78 + int64(base.Ui64(v89)>>(uint(v75)%64)) + int64(base.Ui64(v96)>>(uint(v75)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v18))) = v85&v81 | v96<<(uint(v75)%64)
				v107 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
				v108 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
				if v107 != v108>>(uint(int64(63))%64) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v174 = m.ExcPending
						if v174 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_Fn13998_4), int32(0))
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn13998_1), l6, l1)
								mBase = m.M
								v181 = m.ExcPending
								if v181 != 0 {
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
					v112 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
					v115 = v108 + v112
					if base.B2i32(v112 < int64(0)) != base.B2i32(v115 < v108) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v174 = m.ExcPending
							if v174 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_Fn13998_4), int32(0))
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_Fn13998_1), l6, l1)
									mBase = m.M
									v181 = m.ExcPending
									if v181 != 0 {
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
						if v115 <= int64(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v185 = m.ExcPending
							if v185 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v188 = m.ExcPending
								if v188 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_Fn13998_5), int32(0))
									mBase = m.M
									v192 = m.ExcPending
									if v192 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_Fn13998_1), l5, l1)
										mBase = m.M
										v195 = m.ExcPending
										if v195 != 0 {
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
							v120 = v21 - v27
							if base.B2i32(v120 < v21) != base.B2i32(int64(0) < v27) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v199 = m.ExcPending
								if v199 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v202 = m.ExcPending
									if v202 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_Fn13998_4), int32(0))
										mBase = m.M
										v206 = m.ExcPending
										if v206 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_Fn13998_1), l4, l1)
											mBase = m.M
											v209 = m.ExcPending
											if v209 != 0 {
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
								v125 = base.I64_rem_s(v120, v115)
								v127 = v120 - v125 + v27
								if int64(0) <= v125 {
									v141 = v127
									v146 = F_Int64GetDatum(m, v141)
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
										return int32(0)
									} else {
										m.G0 = v18 + int32(16)
										return v146
									}
								} else {
									v130 = v127 - v115
									if base.B2i32(v130 < v127)^base.B2i32(int64(0) < v115)|base.B2i32(base.Ui64(v130-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615))) != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v213 = m.ExcPending
										if v213 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v216 = m.ExcPending
											if v216 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_Fn13998_6), int32(0))
												mBase = m.M
												v220 = m.ExcPending
												if v220 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_Fn13998_1), l3, l1)
													mBase = m.M
													v223 = m.ExcPending
													if v223 != 0 {
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
										v141 = v130
										v146 = F_Int64GetDatum(m, v141)
										mBase = m.M
										v147 = m.ExcPending
										if v147 != 0 {
											return int32(0)
										} else {
											m.G0 = v18 + int32(16)
											return v146
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
func Fn14003(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v10 = m.G0
	v11 = int32(-64)
	v12 = v10 + v11
	m.G0 = v12
	v14 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+20)))
	v16 = v12 - v11
	v17 = v16
	v21 = v14
	goto L1
L1:
	;
	v27 = v17 - int32(1)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v21)&l3)+uint32(_c_Fn14003[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v30)
	if base.Ui64(v21) < base.Ui64(l2) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v36 = v16 - v27
	v38 = v36 + int32(4)
	v39 = F_palloc(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L2
L4:
	;
	if base.Ui32(v12) < base.Ui32(v27) {
		v17 = v27
		v21 = int64(base.Ui64(v21) >> (uint(l1) % 64))
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v38 << (uint(int32(2)) % 32)
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	base.MemoryCopy(m, v39+int32(4), v27, v36)
	goto L10
L9:
	;
	goto L10
L10:
	;
	m.G0 = v12 - int32(-64)
	return v39
}
func Fn14016(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
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
	v24 = F_DirectFunctionCall1Coll(m, int32(3376), int32(0), v10)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		m.G0 = v10 + int32(48)
		return v24
	}
}
func Fn14021(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v86 int32
	_ = v86
	var v87 float32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v19 = int32(1)
			v20 = v15 + v19
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			v23 = v21 & v19
			if v23 != 0 {
				v24 = v20
			} else {
				v24 = v15 + int32(4)
			}
			if v21 == int32(1) {
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				if v30 == int32(18) {
					v33 = int32(16)
				} else {
					v33 = int32(0)
				}
				if base.Ui32((v30-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v40 = int32(4)
				} else {
					v40 = v33
				}
				v51 = v40
			} else {
				v41 = int32(1)
				if v23 != 0 {
					v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v52 = int32(1)
			v53 = v10 + v52
			v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v58 = v56 & v52
			if v58 != 0 {
				v59 = v53
			} else {
				v59 = v10 + int32(4)
			}
			if v56 == int32(1) {
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				if v65 == int32(18) {
					v68 = int32(16)
				} else {
					v68 = int32(0)
				}
				if base.Ui32((v65-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v75 = int32(4)
				} else {
					v75 = v68
				}
				v86 = v75
			} else {
				v76 = int32(1)
				if v58 != 0 {
					v86 = int32(base.Ui32(v56)>>(uint(v76)%32)) - v76
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v86 = int32(base.Ui32(v80)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v87 = F_calc_word_similarity(m, v24, v51, v59, v86, l1)
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int32(0)
			} else {
				v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v89 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v93 != v15 {
							F_pfree(m, v15)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int32(0)
							} else {
								return base.I32_reinterpret_f32(base.F32_sub(float32(1), v87))
							}
						} else {
							return base.I32_reinterpret_f32(base.F32_sub(float32(1), v87))
						}
					}
				} else {
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v93 != v15 {
						F_pfree(m, v15)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							return base.I32_reinterpret_f32(base.F32_sub(float32(1), v87))
						}
					} else {
						return base.I32_reinterpret_f32(base.F32_sub(float32(1), v87))
					}
				}
			}
		}
	}
}
