package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_get_fn_expr_argtype(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	v3 = int32(0)
	if l0 == v3 {
		v52 = v3
		return v52
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v7 == int32(0) {
			v52 = v3
			return v52
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v12 = v10 - int32(11)
			v19 = int32(0)
			if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v12))|base.B2i32(int32(base.Ui32(int32(977))>>(uint(v12)%32))&int32(1) == v19)|base.B2i32(l1 < v19) != 0 {
				v52 = v3
				return v52
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(2))%32))+uint32(_c_F_get_fn_expr_argtype[0])))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v7+v27)))
				if v29 == int32(0) {
					v52 = v3
					return v52
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
					if v32 <= l1 {
						v52 = v3
						return v52
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+l1<<(uint(int32(2))%32))))
						v39 = F_exprType(m, v38)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							if l1 != int32(1) {
								v52 = v39
								return v52
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
								if v45 != int32(20) {
									v52 = v39
									return v52
								} else {
									v48 = F_get_base_element_type(m, v39)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										v52 = v48
										return v52
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
func F_get_fn_opclass_options(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v2 = int32(0)
	if l0 == v2 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_get_fn_opclass_options_0), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_fn_opclass_options_1), int32(2109), int32(_a_F_get_fn_opclass_options_2))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
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
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v5 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_get_fn_opclass_options_0), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_get_fn_opclass_options_1), int32(2109), int32(_a_F_get_fn_opclass_options_2))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
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
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			if v8 != int32(7) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_get_fn_opclass_options_0), int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_get_fn_opclass_options_1), int32(2109), int32(_a_F_get_fn_opclass_options_2))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
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
				v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
				if v11 != int32(17) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_get_fn_opclass_options_0), int32(0))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_get_fn_opclass_options_1), int32(2109), int32(_a_F_get_fn_opclass_options_2))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
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
					v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+24)))
					if v14 != 0 {
						v25 = v2
						return v25
					} else {
						v15 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
						v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
						if v16&int32(3) == int32(0) {
							v25 = v15
							return v25
						} else {
							v21 = F_detoast_attr(m, v15)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return int32(0)
							} else {
								v25 = v21
								return v25
							}
						}
					}
				}
			}
		}
	}
}
func Fn13845(m *base.Module, l0 float32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
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
						F_errfinish(m, l2, l1, int32(_a_Fn13845_0))
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
					F_errfinish(m, l2, l4, int32(_a_Fn13845_0))
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
func Fn13854(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v10 = int32(0)
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	if l1 == v10 {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
		if v18 != 0 {
			v76 = v10
			m.G0 = v14 - int32(-64)
			return v76
		} else {
			v19 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v19
			*(*int64)(unsafe.Add(mBase, uint32(v14)+29)) = v19
			v23 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v14)+60)) = uint8(v23)
			*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = l3
			*(*uint8)(unsafe.Add(mBase, uint32(v14)+52)) = uint8(v23)
			*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l2
			*(*uint8)(unsafe.Add(mBase, uint32(v14)+44)) = uint8(v23)
			*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = l1
			v32 = int32(3)
			*(*uint16)(unsafe.Add(mBase, uint32(v14)+38)) = uint16(v32)
			*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l0
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v38 = m.T0[v37].(func(*base.Module, int32) int32)(m, v12+int32(-44))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+36)))
				if l1 == int32(0) {
					if v42&int32(1) != 0 {
						v76 = v38
						m.G0 = v14 - int32(-64)
						return v76
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v14))) = v51
							F_errmsg_internal(m, l8, v14)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn13854_0), l7, l4)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
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
					if v42&int32(1) == int32(0) {
						v76 = v38
						m.G0 = v14 - int32(-64)
						return v76
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v66
							F_errmsg_internal(m, l6, v12+int32(-48))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn13854_0), l5, l4)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
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
		v19 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v19
		*(*int64)(unsafe.Add(mBase, uint32(v14)+29)) = v19
		v23 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+60)) = uint8(v23)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = l3
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+52)) = uint8(v23)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l2
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+44)) = uint8(v23)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = l1
		v32 = int32(3)
		*(*uint16)(unsafe.Add(mBase, uint32(v14)+38)) = uint16(v32)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l0
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v38 = m.T0[v37].(func(*base.Module, int32) int32)(m, v12+int32(-44))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+36)))
			if l1 == int32(0) {
				if v42&int32(1) != 0 {
					v76 = v38
					m.G0 = v14 - int32(-64)
					return v76
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v51
						F_errmsg_internal(m, l8, v14)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn13854_0), l7, l4)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
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
				if v42&int32(1) == int32(0) {
					v76 = v38
					m.G0 = v14 - int32(-64)
					return v76
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v66
						F_errmsg_internal(m, l6, v12+int32(-48))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn13854_0), l5, l4)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
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
func Fn13863(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	F_errstart_cold(m, int32(21), int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		F_errcode(m, int32(1088))
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			F_errmsg(m, int32(_a_Fn13863_0), int32(0))
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_errdetail(m, int32(_a_Fn13863_1), int32(0))
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_Fn13863_2), l3, l2)
					v22 = m.ExcPending
					if v22 != 0 {
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
func Fn13865(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v15 = F_array_iterator(m, v7, l1, v12, int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v17 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v21 != v12 {
							F_pfree(m, v12)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return int32(0)
							} else {
								return v15
							}
						} else {
							return v15
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v21 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							return v15
						}
					} else {
						return v15
					}
				}
			}
		}
	}
}
func Fn13869(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
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
		v15 = v11 << (uint(int32(3)) % 32)
		*(*int32)(unsafe.Add(mBase, uint32(v15+l7))) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v15+l6))) = l0
		*(*int32)(unsafe.Add(mBase, uint32(l5))) = v11 + int32(1)
		v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_Fn13869[0])))
		if v24 == int32(0) {
			v29 = *(*int32)(unsafe.Add(mBase, _c_Fn13869[1]))
			if v29 <= int32(31) {
				*(*int32)(unsafe.Add(mBase, _c_Fn13869[1])) = v29 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v29<<(uint(int32(2))%32))+uint32(_c_Fn13869[2]))) = int32(1103)
			} else {
			}
			v46 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_Fn13869[0])) = uint8(v46)
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
					F_errfinish(m, int32(_a_Fn13869_0), l3, l2)
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
func Fn13874(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = l5
			F_errmsg(m, l4, v9)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, l3, l2, l1)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func Fn13878(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(1)
		v13 = v8 + v12
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v18 = v16 & v12
		if v18 != 0 {
			v19 = v13
		} else {
			v19 = v8 + int32(4)
		}
		if v16 == int32(1) {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v25 == int32(18) {
				v28 = int32(16)
			} else {
				v28 = int32(0)
			}
			if base.Ui32((v25-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v35 = int32(4)
			} else {
				v35 = v28
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v18 != 0 {
				v46 = int32(base.Ui32(v16)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v49 = F_dotrim(m, v19, v46, int32(_a_Fn13878_0), int32(1), l2, l1)
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			return v49
		}
	}
}
func Fn13881(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
func Fn13890(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float64, l5 float64) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v17 int32
	_ = v17
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v10 = base.F64_nearest(v9)
	v17 = int32(0)
	if base.B2i32(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v10)&int64(9223372036854775807)))|base.B2i32(base.F64_ge(v10, l5) == v17) == v17)&base.F64_lt(v10, l4) == v17 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, l3, int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn13890_0), l2, l1)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
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
		return base.I32_trunc_sat_f64_s(v10)
	}
}
func Fn13902(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = F_gbt_var_picksplit(m, v4, v5, v6, l1, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func Fn13904(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
func Fn13908(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+22)) = uint16(v14)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v16 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v16
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+16)))
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+v31)+12)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v37 = F_gbt_num_consistent(m, v8+int32(12), v8+int32(24), v8+int32(22), v33&int32(1), l1, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(32)
		return v37
	}
}
func Fn13915(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v22 = v11 + int32(8)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v25 = int32(4)
		v26 = v23 + v25
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = v26
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		v29 = int32(2)
		v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		if base.Ui32(v30+v25) < base.Ui32(int32(base.Ui32(v38)>>(uint(v29)%32))) {
			v42 = v26 + (v30+int32(3))&int32(2147483644)
		} else {
			v42 = v26
		}
		*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v42
		v44 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v44)
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v46 == v44 {
			v50 = *(*int32)(unsafe.Add(mBase, _c_Fn13915[0]))
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v51*int32(28))+uint32(_c_Fn13915[1])))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v56
		} else {
		}
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+16)))
		v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+v64)+12)))
		v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v70 = F_gbt_var_consistent(m, v11+int32(8), v15, v19&int32(_a_Fn13915_0), v62, v66&int32(1), l1, v69)
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return int32(0)
		} else {
			m.G0 = v11 + int32(16)
			return v70
		}
	}
}
func Fn13919(m *base.Module, l0 int32, l1 int32) int32 {
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
func Fn13920(m *base.Module, l0 int32, l1 int32) int32 {
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
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+68))
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
func Fn13928(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
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
					v47 = v16 & int32(_a_Fn13928_0)
					switch v47 - int32(1) {
					case 0, 1:
						v66 = F_Int64GetDatum(m, l1)
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
							F_errmsg_internal(m, int32(_a_Fn13928_1), v13)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn13928_2), int32(97), int32(_a_Fn13928_3))
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
func Fn13931(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	if l1 < int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, _c_Fn13931[0]))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v16+(l1^int32(-1))<<(uint(int32(2))%32))))
		v30 = v22
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, _c_Fn13931[1]))
		v30 = v24 + l1<<(uint(int32(13))%32) + int32(-8192)
	}
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+14)))
	if v31 != 0 {
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+19)))
		v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+16)))
		if (v32<<(uint(int32(8))%32)-v35)&int32(_a_Fn13931_0) != int32(16) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return
			} else {
				F_errcode(m, int32(33557032))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return
				} else {
					v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if l1 < int32(0) {
						v95 = *(*int32)(unsafe.Add(mBase, _c_Fn13931[2]))
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v95+(l1^int32(-1))<<(uint(int32(6))%32))+16))
						v110 = v101
					} else {
						v103 = *(*int32)(unsafe.Add(mBase, _c_Fn13931[3]))
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v103+l1<<(uint(int32(6))%32)+int32(-64))+16))
						v110 = v109
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v110
					*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v91 + int32(4)
					F_errmsg(m, int32(_a_Fn13931_1), v11+int32(16))
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return
					} else {
						F_errhint(m, int32(_a_Fn13931_2), int32(0))
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return
						} else {
							F_errfinish(m, l4, l3, l2)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
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
			m.G0 = v11 + int32(32)
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return
		} else {
			F_errcode(m, int32(33557032))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if l1 < int32(0) {
					v55 = *(*int32)(unsafe.Add(mBase, _c_Fn13931[2]))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v55+(l1^int32(-1))<<(uint(int32(6))%32))+16))
					v70 = v61
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, _c_Fn13931[3]))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v63+l1<<(uint(int32(6))%32)+int32(-64))+16))
					v70 = v69
				}
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v70
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v51 + int32(4)
				F_errmsg(m, int32(_a_Fn13931_3), v11)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return
				} else {
					F_errhint(m, int32(_a_Fn13931_2), int32(0))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						F_errfinish(m, l4, l5, l2)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
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
func Fn13946(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int64) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v67 int64
	_ = v67
	v17 = F_palloc0(m, int32(140))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v17)+4)) = l14
		*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(438)
		*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)) = uint16(v21)
		*(*int32)(unsafe.Add(mBase, uint32(v17)+19)) = v21
		v30 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v17)+18)) = uint8(v30)
		*(*uint16)(unsafe.Add(mBase, uint32(v17)+23)) = uint16(v21)
		*(*int32)(unsafe.Add(mBase, uint32(v17)+108)) = l13
		*(*int32)(unsafe.Add(mBase, uint32(v17)+104)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = l12
		*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = l11
		*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = l10
		*(*int32)(unsafe.Add(mBase, uint32(v17)+88)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v17)+84)) = l9
		*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = l8
		*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = l7
		*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = l6
		*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v21
		*(*uint8)(unsafe.Add(mBase, uint32(v17)+29)) = uint8(v30)
		*(*int32)(unsafe.Add(mBase, uint32(v17)+25)) = v30
		*(*int32)(unsafe.Add(mBase, uint32(v17)+136)) = v21
		v67 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v17)+128)) = v67
		*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = v67
		*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = v67
		return v17
	}
}
func Fn13948(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	F_ean2isn(m, v9, v6+int32(8), l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
		v17 = F_Int64GetDatum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(16)
			return v17
		}
	}
}
func Fn13951(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v11, v12, v13, l2, int32(7))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v21 = F_latin2mic_with_table(m, v9, v8, v13, l3, l2, l1, base.B2i32(v10 != int32(0)))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			return v21
		}
	}
}
func Fn13959(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	if l0 == int32(0) {
		v11 = F_palloc(m, int32(32))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(4)
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = l2
			v19 = v11 + int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v19
			v75 = v11
			v76 = v19
			v77 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v76+v77<<(uint(int32(2))%32)-int32(4)))) = l1
			return v75
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v23 <= v22 {
			v25 = int32(1)
			v27 = int32(16)
			v29 = v22 + v25
			if v29 <= v27 {
				v32 = v27
			} else {
				v32 = v29
			}
			if v32&(v32-int32(1)) != 0 {
				v39 = v25 << (uint(int32(32)-base.I32_clz(v32)) % 32)
			} else {
				v39 = v32
			}
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v42 = l0 + int32(16)
			if v40 == v42 {
				v44 = F_GetMemoryChunkContext(m, l0)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v48 = F_MemoryContextAlloc(m, v44, v39<<(uint(int32(2))%32))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v48
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v53 = v51 << (uint(int32(2)) % 32)
						if v53 == int32(0) {
						} else {
							base.MemoryCopy(m, v48, v42, v53)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v39
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v70 = v65
						v72 = v70 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v75 = l0
						v76 = v74
						v77 = v72
						*(*int32)(unsafe.Add(mBase, uint32(v76+v77<<(uint(int32(2))%32)-int32(4)))) = l1
						return v75
					}
				}
			} else {
				v59 = F_repalloc(m, v40, v39<<(uint(int32(2))%32))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v59
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v39
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v70 = v65
					v72 = v70 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v75 = l0
					v76 = v74
					v77 = v72
					*(*int32)(unsafe.Add(mBase, uint32(v76+v77<<(uint(int32(2))%32)-int32(4)))) = l1
					return v75
				}
			}
		} else {
			v70 = v22
			v72 = v70 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v75 = l0
			v76 = v74
			v77 = v72
			*(*int32)(unsafe.Add(mBase, uint32(v76+v77<<(uint(int32(2))%32)-int32(4)))) = l1
			return v75
		}
	}
}
func Fn13962(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 float64
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 float64
	_ = v56
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 float64
	_ = v61
	var v64 int32
	_ = v64
	var v65 float64
	_ = v65
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v103 float64
	_ = v103
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v107 float64
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v202 float64
	_ = v202
	var v204 float64
	_ = v204
	var v206 float64
	_ = v206
	var v219 float64
	_ = v219
	var v229 float64
	_ = v229
	var v240 float64
	_ = v240
	var v252 float64
	_ = v252
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int64
	_ = v267
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int64
	_ = v280
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int64
	_ = v298
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int64
	_ = v324
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	v11 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(96)
	m.G0 = v18
	if l0 < int32(_a_Fn13962_0) {
		v336 = v11
		m.G0 = v18 + int32(96)
		return v336
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
		if v23 < int64(10000) {
			v336 = v11
			m.G0 = v18 + int32(96)
			return v336
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
			if v26 != int32(1) {
				v336 = v11
				m.G0 = v18 + int32(96)
				return v336
			} else {
				v30 = v22 + int32(16)
				v31 = int32(0)
				v38 = float64(0)
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
				if v40 != 0 {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
					if v40 != int32(1) {
						v50 = v31
						v51 = v31
						v56 = v38
						for {
							v58 = float64(1)
							v59 = v50 + v41
							v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
							v61 = F_scalbn(m, v58, v60)
							mBase = m.M
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
							v65 = F_scalbn(m, v58, v64)
							mBase = m.M
							v70 = base.F64_add(base.F64_add(v56, base.F64_div(v58, v65)), base.F64_div(v58, v61))
							v71 = int32(2)
							v72 = v50 + v71
							v74 = v51 + v71
							if v74 != v40&int32(-2) {
								v50 = v72
								v51 = v74
								v56 = v70
								continue
							} else {
								break
							}
							break
						}
						if v40&int32(1) == int32(0) {
							v103 = v70
						} else {
							v80 = v72
							v86 = v70
							v88 = float64(1)
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v41))))
							v92 = F_scalbn(m, v88, v91)
							mBase = m.M
							v103 = base.F64_add(v86, base.F64_div(v88, v92))
						}
					} else {
						v80 = v31
						v86 = v38
						v88 = float64(1)
						v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v41))))
						v92 = F_scalbn(m, v88, v91)
						mBase = m.M
						v103 = base.F64_add(v86, base.F64_div(v88, v92))
					}
					v105 = *(*float64)(unsafe.Add(mBase, uint32(v30)+8))
					v106 = base.F64_div(v105, v103)
					v107 = base.F64_convert_i32_u(v40)
					if base.F64_le(v106, base.F64_mul(v107, float64(2.5))) == int32(0) {
						v219 = v106
						if base.F64_gt(v219, float64(1.4316557653333333e+08)) == int32(0) {
							v240 = v219
						} else {
							v229 = F_log(m, base.F64_add(base.F64_mul(v219, float64(-2.3283064365386963e-10)), float64(1)))
							mBase = m.M
							v240 = base.F64_mul(v229, float64(-4.294967296e+09))
						}
						v252 = v240
					} else {
						v114 = v40 & int32(3)
						v115 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
						v116 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v40) {
							v125 = int32(0)
							v126 = v116
							v127 = v116
							for {
								v134 = v126 + v115
								v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
								v136 = int32(0)
								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+1)))
								v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+2)))
								v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+3)))
								v150 = v127 + base.B2i32(v135 == v136) + base.B2i32(v139 == v136) + base.B2i32(v143 == v136) + base.B2i32(v147 == v136)
								v151 = int32(4)
								v152 = v126 + v151
								v154 = v125 + v151
								if v154 != v40&int32(-4) {
									v125 = v154
									v126 = v152
									v127 = v150
									continue
								} else {
									break
								}
								break
							}
							if v114 == int32(0) {
								v191 = v150
							} else {
								v160 = v152
								v161 = v150
								v170 = v160
								v171 = v161
								v174 = v116
								for {
									v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v115))))
									v182 = v171 + base.B2i32(v179 == int32(0))
									v183 = int32(1)
									v186 = v174 + v183
									if v186 != v114 {
										v170 = v170 + v183
										v171 = v182
										v174 = v186
										continue
									} else {
										break
									}
									break
								}
								v191 = v182
							}
						} else {
							v160 = v116
							v161 = v116
							v170 = v160
							v171 = v161
							v174 = v116
							for {
								v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v115))))
								v182 = v171 + base.B2i32(v179 == int32(0))
								v183 = int32(1)
								v186 = v174 + v183
								if v186 != v114 {
									v170 = v170 + v183
									v171 = v182
									v174 = v186
									continue
								} else {
									break
								}
								break
							}
							v191 = v182
						}
						if v191 == int32(0) {
							v240 = v106
							v252 = v240
						} else {
							v202 = F_log(m, base.F64_div(v107, base.F64_convert_i32_s(v191)))
							mBase = m.M
							v252 = base.F64_mul(v202, v107)
						}
					}
				} else {
					v204 = *(*float64)(unsafe.Add(mBase, uint32(v30)+8))
					v206 = base.F64_div(v204, float64(0))
					if base.F64_le(v206, base.F64_mul(base.F64_convert_i32_u(v40), float64(2.5))) != 0 {
						v240 = v206
					} else {
						v219 = v206
						if base.F64_gt(v219, float64(1.4316557653333333e+08)) == int32(0) {
							v240 = v219
						} else {
							v229 = F_log(m, base.F64_add(base.F64_mul(v219, float64(-2.3283064365386963e-10)), float64(1)))
							mBase = m.M
							v240 = base.F64_mul(v229, float64(-4.294967296e+09))
						}
					}
					v252 = v240
				}
				if base.F64_gt(v252, float64(100000)) != 0 {
					v256 = int32(*(*uint8)(unsafe.Add(mBase, _c_Fn13962[0])))
					if v256 != int32(1) {
						v276 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)) = uint8(v276)
						v336 = v11
						m.G0 = v18 + int32(96)
						return v336
					} else {
						v261 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v264 = m.ExcPending
						if v264 != 0 {
							return int32(0)
						} else {
							if v261 == int32(0) {
								v276 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)) = uint8(v276)
								v336 = v11
								m.G0 = v18 + int32(96)
								return v336
							} else {
								v267 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
								*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = l0
								*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v267
								*(*float64)(unsafe.Add(mBase, uint32(v18))) = v252
								F_errmsg_internal(m, l9, v18)
								mBase = m.M
								v272 = m.ExcPending
								if v272 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, l4, l8, l2)
									mBase = m.M
									v274 = m.ExcPending
									if v274 != 0 {
										return int32(0)
									} else {
										v276 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)) = uint8(v276)
										v336 = v11
										m.G0 = v18 + int32(96)
										return v336
									}
								}
							}
						}
					}
				} else {
					v279 = int32(*(*uint8)(unsafe.Add(mBase, _c_Fn13962[0])))
					v280 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
					if base.F64_gt(base.F64_add(base.F64_div(base.F64_convert_i64_s(v280), float64(2000)), float64(0.5)), v252) != 0 {
						v287 = int32(1)
						if v279&v287 == int32(0) {
							v336 = v287
							m.G0 = v18 + int32(96)
							return v336
						} else {
							v294 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v295 = m.ExcPending
							if v295 != 0 {
								return int32(0)
							} else {
								if v294 == int32(0) {
									v336 = v287
									m.G0 = v18 + int32(96)
									return v336
								} else {
									v298 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
									*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = l0
									*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v298
									*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v252
									*(*float64)(unsafe.Add(mBase, uint32(v18)+40)) = base.F64_add(base.F64_div(base.F64_convert_i64_s(v298), float64(2000)), float64(0.5))
									F_errmsg_internal(m, l7, v18+int32(32))
									mBase = m.M
									v311 = m.ExcPending
									if v311 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, l4, l6, l2)
										mBase = m.M
										v313 = m.ExcPending
										if v313 != 0 {
											return int32(0)
										} else {
											v336 = v287
											m.G0 = v18 + int32(96)
											return v336
										}
									}
								}
							}
						}
					} else {
						if v279&int32(1) == int32(0) {
							v336 = v11
							m.G0 = v18 + int32(96)
							return v336
						} else {
							v320 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v321 = m.ExcPending
							if v321 != 0 {
								return int32(0)
							} else {
								if v320 == int32(0) {
									v336 = v11
									m.G0 = v18 + int32(96)
									return v336
								} else {
									v324 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
									*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = l0
									*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v324
									*(*float64)(unsafe.Add(mBase, uint32(v18)+64)) = v252
									F_errmsg_internal(m, l5, v18-int32(-64))
									mBase = m.M
									v331 = m.ExcPending
									if v331 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, l4, l3, l2)
										mBase = m.M
										v333 = m.ExcPending
										if v333 != 0 {
											return int32(0)
										} else {
											v336 = v11
											m.G0 = v18 + int32(96)
											return v336
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
func Fn13968(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v10, v11, v12, int32(7), l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v20 = F_mic2latin(m, v8, v7, v12, l2, l1, base.B2i32(v9 != int32(0)))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			return v20
		}
	}
}
func Fn13975(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
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
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	v13 = int32(1)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v20 == v13 {
		if v16&int32(1) == int32(0) {
			v40 = v19
			v41 = v19
			v42 = v13
			if v15&int32(1) != 0 {
				if v14&int32(1) == int32(0) {
					v53 = v17
					v55 = F_DirectFunctionCall2Coll(m, l2, int32(0), v41, v53)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v105 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
							v115 = int32(0)
							return v115
						} else {
							v79 = v53
							v80 = int32(1)
							v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								if v83 != 0 {
									if v42 != 0 {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
										return v115
									} else {
										v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											if v80&base.B2i32(v86 == int32(0)) != 0 {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v115 = int32(0)
												return v115
											} else {
												return base.B2i32(v86 != int32(0))
											}
										}
									}
								} else {
									if v80|v42 == int32(0) {
										v115 = int32(1)
									} else {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
									}
									return v115
								}
							}
						}
					}
				} else {
					v105 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
					v115 = int32(0)
					return v115
				}
			} else {
				if v14&int32(1) == int32(0) {
					v58 = int32(0)
					v61 = F_DirectFunctionCall2Coll(m, l2, v58, v18, v17)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						if v61 != 0 {
							v63 = v17
						} else {
							v63 = v18
						}
						v64 = F_DirectFunctionCall2Coll(m, l2, v58, v41, v63)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							if v64 == int32(0) {
								v79 = v63
								v80 = v58
								v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									if v83 != 0 {
										if v42 != 0 {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
											return v115
										} else {
											v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												if v80&base.B2i32(v86 == int32(0)) != 0 {
													v105 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
													v115 = int32(0)
													return v115
												} else {
													return base.B2i32(v86 != int32(0))
												}
											}
										}
									} else {
										if v80|v42 == int32(0) {
											v115 = int32(1)
										} else {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
										}
										return v115
									}
								}
							} else {
								if v61 != 0 {
									v69 = v18
								} else {
									v69 = v17
								}
								v70 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v69)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									if v42&base.B2i32(v70 == int32(0)) != 0 {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
										return v115
									} else {
										return base.B2i32(v70 != int32(0))
									}
								}
							}
						}
					}
				} else {
					v53 = v18
					v55 = F_DirectFunctionCall2Coll(m, l2, int32(0), v41, v53)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v105 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
							v115 = int32(0)
							return v115
						} else {
							v79 = v53
							v80 = int32(1)
							v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								if v83 != 0 {
									if v42 != 0 {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
										return v115
									} else {
										v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											if v80&base.B2i32(v86 == int32(0)) != 0 {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v115 = int32(0)
												return v115
											} else {
												return base.B2i32(v86 != int32(0))
											}
										}
									}
								} else {
									if v80|v42 == int32(0) {
										v115 = int32(1)
									} else {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
									}
									return v115
								}
							}
						}
					}
				}
			}
		} else {
			v105 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
			v115 = int32(0)
			return v115
		}
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v16&int32(1) != 0 {
			v40 = v19
			v41 = v27
			v42 = v13
			if v15&int32(1) != 0 {
				if v14&int32(1) == int32(0) {
					v53 = v17
					v55 = F_DirectFunctionCall2Coll(m, l2, int32(0), v41, v53)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v105 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
							v115 = int32(0)
							return v115
						} else {
							v79 = v53
							v80 = int32(1)
							v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								if v83 != 0 {
									if v42 != 0 {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
										return v115
									} else {
										v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											if v80&base.B2i32(v86 == int32(0)) != 0 {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v115 = int32(0)
												return v115
											} else {
												return base.B2i32(v86 != int32(0))
											}
										}
									}
								} else {
									if v80|v42 == int32(0) {
										v115 = int32(1)
									} else {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
									}
									return v115
								}
							}
						}
					}
				} else {
					v105 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
					v115 = int32(0)
					return v115
				}
			} else {
				if v14&int32(1) == int32(0) {
					v58 = int32(0)
					v61 = F_DirectFunctionCall2Coll(m, l2, v58, v18, v17)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						if v61 != 0 {
							v63 = v17
						} else {
							v63 = v18
						}
						v64 = F_DirectFunctionCall2Coll(m, l2, v58, v41, v63)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							if v64 == int32(0) {
								v79 = v63
								v80 = v58
								v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									if v83 != 0 {
										if v42 != 0 {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
											return v115
										} else {
											v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												if v80&base.B2i32(v86 == int32(0)) != 0 {
													v105 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
													v115 = int32(0)
													return v115
												} else {
													return base.B2i32(v86 != int32(0))
												}
											}
										}
									} else {
										if v80|v42 == int32(0) {
											v115 = int32(1)
										} else {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
										}
										return v115
									}
								}
							} else {
								if v61 != 0 {
									v69 = v18
								} else {
									v69 = v17
								}
								v70 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v69)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									if v42&base.B2i32(v70 == int32(0)) != 0 {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
										return v115
									} else {
										return base.B2i32(v70 != int32(0))
									}
								}
							}
						}
					}
				} else {
					v53 = v18
					v55 = F_DirectFunctionCall2Coll(m, l2, int32(0), v41, v53)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v105 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
							v115 = int32(0)
							return v115
						} else {
							v79 = v53
							v80 = int32(1)
							v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								if v83 != 0 {
									if v42 != 0 {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
										return v115
									} else {
										v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											if v80&base.B2i32(v86 == int32(0)) != 0 {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v115 = int32(0)
												return v115
											} else {
												return base.B2i32(v86 != int32(0))
											}
										}
									}
								} else {
									if v80|v42 == int32(0) {
										v115 = int32(1)
									} else {
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v115 = int32(0)
									}
									return v115
								}
							}
						}
					}
				}
			}
		} else {
			v30 = int32(0)
			v32 = F_DirectFunctionCall2Coll(m, l2, v30, v27, v19)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				if v32 != 0 {
					v36 = v19
				} else {
					v36 = v27
				}
				if v32 != 0 {
					v37 = v27
				} else {
					v37 = v19
				}
				v40 = v37
				v41 = v36
				v42 = v30
				if v15&int32(1) != 0 {
					if v14&int32(1) == int32(0) {
						v53 = v17
						v55 = F_DirectFunctionCall2Coll(m, l2, int32(0), v41, v53)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							if v55 != 0 {
								v105 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
								v115 = int32(0)
								return v115
							} else {
								v79 = v53
								v80 = int32(1)
								v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									if v83 != 0 {
										if v42 != 0 {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
											return v115
										} else {
											v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												if v80&base.B2i32(v86 == int32(0)) != 0 {
													v105 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
													v115 = int32(0)
													return v115
												} else {
													return base.B2i32(v86 != int32(0))
												}
											}
										}
									} else {
										if v80|v42 == int32(0) {
											v115 = int32(1)
										} else {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
										}
										return v115
									}
								}
							}
						}
					} else {
						v105 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
						v115 = int32(0)
						return v115
					}
				} else {
					if v14&int32(1) == int32(0) {
						v58 = int32(0)
						v61 = F_DirectFunctionCall2Coll(m, l2, v58, v18, v17)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							if v61 != 0 {
								v63 = v17
							} else {
								v63 = v18
							}
							v64 = F_DirectFunctionCall2Coll(m, l2, v58, v41, v63)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								if v64 == int32(0) {
									v79 = v63
									v80 = v58
									v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										if v83 != 0 {
											if v42 != 0 {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v115 = int32(0)
												return v115
											} else {
												v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return int32(0)
												} else {
													if v80&base.B2i32(v86 == int32(0)) != 0 {
														v105 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
														v115 = int32(0)
														return v115
													} else {
														return base.B2i32(v86 != int32(0))
													}
												}
											}
										} else {
											if v80|v42 == int32(0) {
												v115 = int32(1)
											} else {
												v105 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
												v115 = int32(0)
											}
											return v115
										}
									}
								} else {
									if v61 != 0 {
										v69 = v18
									} else {
										v69 = v17
									}
									v70 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v69)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										if v42&base.B2i32(v70 == int32(0)) != 0 {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
											return v115
										} else {
											return base.B2i32(v70 != int32(0))
										}
									}
								}
							}
						}
					} else {
						v53 = v18
						v55 = F_DirectFunctionCall2Coll(m, l2, int32(0), v41, v53)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							if v55 != 0 {
								v105 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
								v115 = int32(0)
								return v115
							} else {
								v79 = v53
								v80 = int32(1)
								v83 = F_DirectFunctionCall2Coll(m, l1, int32(0), v41, v79)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									if v83 != 0 {
										if v42 != 0 {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
											return v115
										} else {
											v86 = F_DirectFunctionCall2Coll(m, l1, int32(0), v79, v40)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												if v80&base.B2i32(v86 == int32(0)) != 0 {
													v105 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
													v115 = int32(0)
													return v115
												} else {
													return base.B2i32(v86 != int32(0))
												}
											}
										}
									} else {
										if v80|v42 == int32(0) {
											v115 = int32(1)
										} else {
											v105 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
											v115 = int32(0)
										}
										return v115
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
func Fn13982(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v4 - int32(142) {
	case 0:
		v15 = l1
		return v15
	case 1:
		return int32(3)
	default:
		if int32(0) <= base.I32_extend8_s(v4) {
			v14 = int32(1)
		} else {
			v14 = int32(2)
		}
		v15 = v14
		return v15
	}
}
func Fn13984(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_get_statisticsobj_worker(m, v4, l1, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v12 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
			return int32(0)
		} else {
			v16 = F_cstring_to_text(m, v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v6)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					return v16
				}
			}
		}
	}
}
func Fn13995(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v52 int64
	_ = v52
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int64
	_ = v166
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	v4 = int32(0)
	v13 = int64(2)
	if base.Ui64(l1) <= base.Ui64(v13) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = v13
	goto L3
L2:
	;
	v16 = l1
	goto L3
L3:
	;
	v17 = int64(1)
	if v16&(v16-v17) == int64(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v27 = v16
	goto L6
L5:
	;
	v27 = v17 << (uint(int64(64)-base.I64_clz(v16)) % 64)
	goto L6
L6:
	;
	if base.Ui64(v27<<(uint(int64(3))%64)) < base.Ui64(int64(2147483647)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v39 = F_MemoryContextAllocExtended(m, v34, base.I32_wrap_i64(v27)<<(uint(int32(3))%32), int32(5))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L11
	} else {
		goto L42
	}
L10:
	;
	goto L9
L11:
	;
	return
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v39
	v42 = int64(1)
	if v27&(v27-v42) == int64(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v52 = v27
	goto L15
L14:
	;
	v52 = v42 << (uint(int64(64)-base.I64_clz(v27)) % 64)
	goto L15
L15:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v52<<(uint(int64(3))%64)) {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v52
	v60 = base.I32_wrap_i64(v52) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v60
	if v52 == int64(4294967296) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v69 = int32(-85899346)
	goto L19
L18:
	;
	v69 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v52), float64(0.9)))
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v69
	if v33 != int64(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v77 = v4
	goto L24
L21:
	;
	goto L22
L22:
	;
	F_pfree(m, v32)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L11
	} else {
		goto L41
	}
L23:
	;
	v118 = v113
	v124 = v4
	goto L29
L24:
	;
	v87 = v32 + v77<<(uint(int32(3))%32)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+4)))
	if v88 != int32(1) {
		v113 = v77
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v113 = int32(0)
	goto L23
L26:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v92 = int32(16)
	v96 = (int32(base.Ui32(v91)>>(uint(v92)%32)) ^ v91) * int32(-2048144789)
	v101 = (int32(base.Ui32(v96)>>(uint(int32(13))%32)) ^ v96) * int32(-1028477387)
	if (int32(base.Ui32(v101)>>(uint(v92)%32))^v101)&v60 == v77 {
		v113 = v77
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v108 = v77 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v108)) < base.Ui64(v33) {
		v77 = v108
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v128 = v32 + v118<<(uint(int32(3))%32)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+4)))
	if v129 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L22
L31:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v133 = int32(16)
	v137 = (int32(base.Ui32(v132)>>(uint(v133)%32)) ^ v132) * int32(-2048144789)
	v142 = (int32(base.Ui32(v137)>>(uint(int32(13))%32)) ^ v137) * int32(-1028477387)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v150 = int32(base.Ui32(v142)>>(uint(v133)%32)) ^ v142
	goto L34
L32:
	;
	goto L33
L33:
	;
	v181 = v118 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v181)) < base.Ui64(v33) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v159 = v150 & v146
	v164 = v39 + v159<<(uint(int32(3))%32)
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+4)))
	if v165 != 0 {
		v150 = v159 + int32(1)
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v128)))
	*(*int64)(unsafe.Add(mBase, uint32(v164))) = v166
	goto L33
L36:
	;
	goto L35
L37:
	;
	v185 = v181
	goto L39
L38:
	;
	v185 = int32(0)
	goto L39
L39:
	;
	v187 = v124 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v187)) < base.Ui64(v33) {
		v118 = v185
		v124 = v187
		goto L29
	} else {
		goto L40
	}
L40:
	;
	goto L30
L41:
	;
	return
L42:
	;
	F_errmsg_internal(m, int32(_a_Fn13995_0), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_Fn13995_1), int32(327), l2)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L11
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func Fn14000(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l3
	v15 = F_query_or_expression_tree_walker_impl(m, l0, l2, v8+int32(8), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		m.G0 = v8 + int32(16)
		return v19
	}
}
func Fn14006(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v5 <= v7 {
		v36 = v3
		return v36
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v5-int32(1)))))
		switch v13 - int32(105) {
		case 0, 5:
			v17 = F_find_among_b(m, l0, l1, int32(3))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v17 == int32(0) {
					v36 = v3
					return v36
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
					v25 = F_slice_del(m, l0)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if v25 < int32(0) {
							v36 = v25
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
							v31 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v30 - v31
							v36 = v31
						}
						return v36
					}
				}
			}
		default:
			v36 = v3
			return v36
		}
	}
}
func Fn14011(m *base.Module, l0 int32, l1 int32) int32 {
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
				F_errmsg_internal(m, int32(_a_Fn14011_0), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn14011_1), int32(773), int32(_a_Fn14011_2))
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
func Fn14022(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v15 = F_ArrayGetIntegerTypmods(m, v9, v6+int32(12))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			if v17 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_Fn14022_0), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn14022_1), int32(118), int32(_a_Fn14022_2))
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
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v37 = F_anytimestamp_typmod_check(m, l1, v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return v37
				}
			}
		}
	}
}
func Fn14028(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v15
		v18 = F_text_to_cstring(m, v11)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v24 = F_parse_tsquery(m, v18, int32(1158), v8+int32(8), l1, int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(16)
				return v24
			}
		}
	}
}
func Fn14031(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v7 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_Fn14031_0), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn14031_1), l4, l3)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
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
		v27 = F_heap_getsysattr(m, v7, l1, l2)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
	}
}
func Fn14037(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(1)
		v13 = v8 + v12
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v18 = v16 & v12
		if v18 != 0 {
			v19 = v13
		} else {
			v19 = v8 + int32(4)
		}
		if v16 == int32(1) {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v25 == int32(18) {
				v28 = int32(16)
			} else {
				v28 = int32(0)
			}
			if base.Ui32((v25-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v35 = int32(4)
			} else {
				v35 = v28
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v18 != 0 {
				v46 = int32(base.Ui32(v16)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v47 = F_uuid_generate_internal(m, l1, v6, v19, v46)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			return v47
		}
	}
}
func Fn14046(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
		F_errmsg_internal(m, l4, v9)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			F_errfinish(m, l3, l2, l1)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
