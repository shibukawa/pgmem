package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_fn_expr_rettype(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	v2 = int32(0)
	if l0 == v2 {
		v13 = v2
		return v13
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v5 == int32(0) {
			v13 = v2
			return v13
		} else {
			v8 = F_exprType(m, v5)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				v13 = v8
				return v13
			}
		}
	}
}
func Fn13844(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if int32(0) < l0 {
		if base.Ui32(l7) <= base.Ui32(l0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = l4
					F_errmsg(m, l3, v12)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						F_errfinish(m, l2, l1, int32(_a_Fn13844_0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
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
			m.G0 = v12 + int32(16)
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_errcode(m, int32(130))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_errmsg(m, l6, int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_errfinish(m, l2, l5, int32(_a_Fn13844_0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
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
func Fn13855(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_SearchSysCacheExists(m, l6, l1, l0, l2, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		if v14 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				F_errcode(m, int32(_a_Fn13855_0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v23 = F_get_am_name(m, l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						v25 = F_get_namespace_name(m, l2)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v25
							*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v23
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
							F_errmsg(m, l5, v11)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_Fn13855_1), l4, l3)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
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
func Fn13862(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
			F_errmsg(m, int32(_a_Fn13862_0), int32(0))
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errdetail(m, int32(_a_Fn13862_1), int32(0))
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_Fn13862_2), l2, l1)
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
func Fn13864(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = F_pg_detoast_datum(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v23 = F_array_iterator(m, v14, l1, v19, v11+int32(12))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 == int32(0) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v27 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v31 != v19 {
								F_pfree(m, v19)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return int32(0)
								} else {
									v35 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v35)
									v56 = int32(0)
									m.G0 = v11 + int32(16)
									return v56
								}
							} else {
								v35 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v35)
								v56 = int32(0)
								m.G0 = v11 + int32(16)
								return v56
							}
						}
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v31 != v19 {
							F_pfree(m, v19)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v35 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v35)
								v56 = int32(0)
								m.G0 = v11 + int32(16)
								return v56
							}
						} else {
							v35 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v35)
							v56 = int32(0)
							m.G0 = v11 + int32(16)
							return v56
						}
					}
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
					v41 = F_palloc0(m, int32(base.Ui32(v38)>>(uint(int32(2))%32)))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
						v45 = int32(base.Ui32(v43) >> (uint(int32(2)) % 32))
						if v45 != 0 {
							base.MemoryCopy(m, v41, v37, v45)
						} else {
						}
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v47 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v19 == v51 {
									v56 = v41
									m.G0 = v11 + int32(16)
									return v56
								} else {
									F_pfree(m, v19)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v56 = v41
										m.G0 = v11 + int32(16)
										return v56
									}
								}
							}
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v19 == v51 {
								v56 = v41
								m.G0 = v11 + int32(16)
								return v56
							} else {
								F_pfree(m, v19)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v56 = v41
									m.G0 = v11 + int32(16)
									return v56
								}
							}
						}
					}
				}
			}
		}
	}
}
func Fn13868(m *base.Module, l0 int32, l1 int32) int32 {
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
func Fn13875(m *base.Module, l0 int32, l1 int32) int32 {
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
	v4 = int32(_a_Fn13875_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_Fn13875[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	*(*int32)(unsafe.Add(mBase, _c_Fn13875[0])) = v8
	F_varstr_sortsupport(m, v7, l1, int32(950))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_Fn13875[0])) = v5
		return int32(0)
	}
}
func Fn13879(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
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
	v8 = l1 << (uint(l3) % 32)
	v10 = v8 + int32(24)
	v11 = F_palloc0(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(0)
		if base.B2i32(v8 == v15)|(base.B2i32(l0 == v15)|base.B2i32(l1 <= v15)) == v15 {
			base.MemoryCopy(m, v11+int32(24), l0, v8)
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l2
		*(*int64)(unsafe.Add(mBase, uint32(v11)+4)) = int64(1)
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v10 << (uint(int32(2)) % 32)
		return v11
	}
}
func Fn13880(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v13 = v12 - v8
	v16 = int32(0)
	if base.B2i32(base.B2i32(int64(0) < v8)^base.B2i32(v13 < v12) == v16)&base.B2i32(v13 != int64(-9223372036854775807-1)) == v16 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, l3, int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l2, int32(107), l1)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
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
		v39 = v13 >> (uint(int64(63)) % 64)
		v42 = F_Int64GetDatum(m, v13^v39-v39)
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			return v42
		}
	}
}
func Fn13891(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = F_string2ean(m, v8, v9, v6+int32(8), l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v24 = int32(0)
			m.G0 = v6 + int32(16)
			return v24
		} else {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
			v22 = F_Int64GetDatum(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = v22
				m.G0 = v6 + int32(16)
				return v24
			}
		}
	}
}
func Fn13903(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = F_gbt_var_same(m, v5, v6, v7, l1, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v4))) = uint8(v9)
		return v4
	}
}
func Fn13905(m *base.Module, l0 int32, l1 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = F_gbt_num_same(m, v5, v6, l1, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v4))) = uint8(v8)
		return v4
	}
}
func Fn13909(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 float64
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v15 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v14 + v15
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v14
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+16)))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v22)+12)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = F_gbt_num_distance(m, v8, v8+v15, v24&int32(1), l1, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int32(0)
	} else {
		v32 = F_Float8GetDatum(m, v28)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v32
		}
	}
}
func Fn13914(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_palloc0(m, int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(16)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = F_gbt_num_union(m, v7, v5, l1, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
func Fn13918(m *base.Module, l0 int32, l1 int32) int32 {
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
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+92))
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
func Fn13921(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
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
func Fn13929(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
					v47 = v16 & int32(_a_Fn13929_0)
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
							F_errmsg_internal(m, int32(_a_Fn13929_1), v13)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn13929_2), int32(97), int32(_a_Fn13929_3))
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
func Fn13930(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l3 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v20 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = int32(0)
	goto L4
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37+v25<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = l0
	v43 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+24)) = uint16(v43)
	v46 = v25 + int32(1)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v46 < v47 {
		v25 = v46
		goto L4
	} else {
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	goto L5
L7:
	;
	m.G0 = v16 + int32(16)
	return
L8:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v64 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v69 = int32(0)
	goto L10
L10:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81+v69<<(uint(int32(2))%32))))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	if base.Ui32(l10) < base.Ui32(v86) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L7
L12:
	;
	v116 = v69 + int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v116 < v117 {
		v69 = v116
		goto L10
	} else {
		goto L23
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = l0
	v113 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v85)+24)) = uint16(v113)
	goto L12
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v89 = int32(1) << (uint(v86) % 32)
	if v89&l9 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v89&l8 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v94 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v85)+24)) = uint8(v94)
	goto L12
L18:
	;
	return
L19:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v104
	F_errmsg(m, int32(_a_Fn13930_0), v16)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, l6, l5, l4)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	goto L11
}
func Fn13941(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
func Fn13947(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	v7 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v11)+13)) = v15
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+36)) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v13
	v30 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+22)) = uint16(v30)
	v34 = F_array_recv(m, v11+int32(4))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return int32(0)
	} else {
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
		if v38 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50462850))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, l4, int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, l3, l2, l1)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
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
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
			if v41 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50462850))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, l4, int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, l3, l2, l1)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
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
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
				if v42 != l5 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50462850))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, l4, int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, l3, l2, l1)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
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
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
					if v44 == int32(0) {
						m.G0 = v11 + int32(48)
						return v34
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50462850))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, l4, int32(0))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, l3, l2, l1)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
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
}
func Fn13949(m *base.Module, l0 int32, l1 int32) int32 {
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
func Fn13950(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v23, v24, v25, int32(-1), int32(6))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int32(0)
	} else {
		v32 = v23 - l8
		if base.B2i32(base.Ui32(v32) <= base.Ui32(l7))&(int32(base.Ui32(l6)>>(uint(v32)%32))&int32(1)) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(2600))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v23
					F_errmsg(m, l5, v18)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, l4, l3, l2)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
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
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32)+l1)))
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
			v57 = int32(0)
			v62 = F_LocalToUtf(m, v22, v25, v21, v56, v57, v57, v57, v23, base.B2i32(v20 != v57))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int32(0)
			} else {
				m.G0 = v18 + int32(16)
				return v62
			}
		}
	}
}
func Fn13958(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v11, v12, v13, l3, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v20 = F_local2local(m, v9, v8, v13, l3, l2, l1, base.B2i32(v10 != int32(0)))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			return v20
		}
	}
}
func Fn13963(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
	v8 = int32(8)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v11 = int32(16)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)))
	v15 = v7<<(uint(v8)%32) | v10<<(uint(v11)%32) | v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+2)))
	v25 = v17<<(uint(v8)%32) | v20<<(uint(v11)%32) | v24
	if base.Ui32(v15) < base.Ui32(v25) {
		v51 = l1
	} else {
		if base.Ui32(v25) < base.Ui32(v15) {
			v51 = int32(1)
		} else {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)))
			v31 = int32(8)
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)))
			v34 = int32(16)
			v37 = v29 | (v30<<(uint(v31)%32) | v33<<(uint(v34)%32))
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+5)))
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+4)))
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+3)))
			v46 = v38 | (v39<<(uint(v31)%32) | v42<<(uint(v34)%32))
			if base.Ui32(v37) < base.Ui32(v46) {
				v51 = l1
			} else {
				v51 = base.B2i32(base.Ui32(v46) < base.Ui32(v37))
			}
		}
	}
	return v51
}
func Fn13969(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_strlen(m, v6)
		mBase = m.M
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v14 = F_RE_compile_and_cache(m, v8, l1, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v20 = F_palloc(m, v12<<(uint(int32(2))%32)+int32(4))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = F_pg_mb2wchar_with_len(m, v6, v20, v12)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = int32(0)
					v27 = F_RE_wchar_execute(m, v20, v22, v24, v24, v24)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v20)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							return v27
						}
					}
				}
			}
		}
	}
}
func Fn13974(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v17 int32
	_ = v17
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_palloc(m, l2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
		v16 = F_pg_snprintf(m, v11, l2, l1, v8)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v11
		}
	}
}
func Fn13983(m *base.Module, l0 int32, l1 int32) int32 {
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
func Fn13985(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
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
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
		if v20 == int32(_a_Fn13985_0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, l4, int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_Fn13985_1), l3, l2)
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
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = v14
			v38 = v11 + int32(16)
			v41 = F_pg_snprintf(m, v38, int32(32), int32(_a_Fn13985_2), v11)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				v44 = int32(0)
				v50 = F_DirectFunctionCall3Coll(m, int32(408), v44, v38, v44, int32(-1))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					v52 = F_DirectFunctionCall2Coll(m, l1, v44, v50, v16)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = F_DirectFunctionCall1Coll(m, int32(1465), v44, v52)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							m.G0 = v11 + int32(48)
							return v54
						}
					}
				}
			}
		}
	}
}
func Fn13994(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(176)
	m.G0 = v14
	v23 = v6
	v24 = int32(-1)
	v25 = v6
	v26 = v6
	goto L1
L1:
	;
	if v24 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn13994[0])) = v47
	*(*int32)(unsafe.Add(mBase, _c_Fn13994[1])) = v48
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v98 - int32(1)
	m.G0 = v14 + int32(176)
	return
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v31 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v30 + v31
	v35 = *(*int32)(unsafe.Add(mBase, _c_Fn13994[0]))
	v37 = *(*int32)(unsafe.Add(mBase, _c_Fn13994[1]))
	v39 = v14 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v14 + int32(12)
	goto L6
L4:
	;
	v46 = v23
	v47 = v25
	v48 = v26
	goto L5
L5:
	;
	goto L8
L6:
	;
	v46 = int32(0)
	v47 = v35
	v48 = v37
	goto L5
L7:
	;
	goto L2
L8:
	;
	if v46 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L7
L10:
	;
	v73 = int32(m.ExcTag)
	v74 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v73 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	F_standard_ExecutorRun(m, l0, l1, l2)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L10
	} else {
		goto L18
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn13994[0])) = v14 + int32(16)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v55 == int32(0) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn13994[1])) = v48
	*(*int32)(unsafe.Add(mBase, _c_Fn13994[0])) = v47
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v64 - int32(1)
	F_pg_re_throw(m)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L17
	}
L15:
	;
	m.T0[v55].(func(*base.Module, int32, int32, int64))(m, l0, l1, l2)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L7
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	goto L9
L19:
	;
	v78 = int32(v74)
	m.G0 = v14
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v14+int32(12) == v84 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	m.ExcPending = 1
	goto L28
L21:
	;
	if v88 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v88 = v86
	goto L24
L23:
	;
	v88 = int32(0)
	goto L24
L24:
	;
	goto L21
L25:
	;
	F___wasm_longjmp(m, v81, v80)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v23 = v80
	v24 = v88
	v25 = v47
	v26 = v48
	goto L1
L28:
	;
	return
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func Fn14001(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v6 = F_palloc0(m, int32(24))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l2
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(0)
		v14 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v14
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v14
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v6
		return v14
	}
}
func Fn14007(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v12 = int32(1)
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4+int32(base.Ui32(v8)>>(uint(int32(2))%32))-v12))))
		return int32(base.Ui32(v14)>>(uint(l1)%32)) & v12
	}
}
func Fn14010(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 float64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = F_patternsel_common(m, v3, v4, v5, v6, v7, v8, l1, v5)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_Float8GetDatum(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v14
		}
	}
}
func Fn14023(m *base.Module, l0 int32, l1 int32) int32 {
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
						F_errmsg(m, int32(_a_Fn14023_0), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn14023_1), int32(65), int32(_a_Fn14023_2))
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
				v37 = F_anytime_typmod_check(m, l1, v36)
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
func Fn14029(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v2 = l1
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = F_palloc0(m, int32(36))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(440)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l7
		*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = l7
		v27 = F_list_make1_impl(m, int32(472), v13+int32(8))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = l6
			*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l5
			*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l2
			*(*uint16)(unsafe.Add(mBase, uint32(v16)+8)) = uint16(v2)
			*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v27
			m.G0 = v13 + int32(16)
			return v16
		}
	}
}
func Fn14030(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	v2 = l1
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_copy(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v13 = F_pg_detoast_datum_copy(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			if v15 == int32(0) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v8 != v18 {
					v76 = v13
					v78 = v8
					F_pfree(m, v78)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						v81 = v76
						return v81
					}
				} else {
					v81 = v13
					return v81
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
				if v20 == int32(0) {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v23 != v13 {
						v76 = v8
						v78 = v13
						F_pfree(m, v78)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							v81 = v76
							return v81
						}
					} else {
						v81 = v8
						return v81
					}
				} else {
					v26 = F_palloc0(m, int32(24))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v28 | int32(1)
						v33 = F_palloc0(m, int32(12))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v26))) = v33
							v36 = int32(2)
							*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v36)
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
							*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)) = uint8(v2)
							v41 = F_palloc0(m, int32(8))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v41
								v45 = v13 + int32(8)
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								v50 = F_QT2QTN(m, v45, v45+v46*int32(12))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
									*(*int32)(unsafe.Add(mBase, uint32(v52))) = v50
									v55 = v8 + int32(8)
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
									v60 = F_QT2QTN(m, v55, v55+v56*int32(12))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v60
										*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = int32(2)
										v66 = F_QTN2QT(m, v26)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											F_QTNFree(m, v26)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												if v70 != v8 {
													F_pfree(m, v8)
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return int32(0)
													} else {
														v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														if v74 == v13 {
															v81 = v66
															return v81
														} else {
															v76 = v66
															v78 = v13
															F_pfree(m, v78)
															mBase = m.M
															v80 = m.ExcPending
															if v80 != 0 {
																return int32(0)
															} else {
																v81 = v76
																return v81
															}
														}
													}
												} else {
													v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													if v74 == v13 {
														v81 = v66
														return v81
													} else {
														v76 = v66
														v78 = v13
														F_pfree(m, v78)
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return int32(0)
														} else {
															v81 = v76
															return v81
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
			}
		}
	}
}
func Fn14036(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v23, v24, v25, int32(6), int32(-1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int32(0)
	} else {
		v32 = v24 - l8
		if base.B2i32(base.Ui32(v32) <= base.Ui32(l7))&(int32(base.Ui32(l6)>>(uint(v32)%32))&int32(1)) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(2600))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v24
					F_errmsg(m, l5, v18)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, l4, l3, l2)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
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
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32)+l1)))
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
			v57 = int32(0)
			v62 = F_UtfToLocal(m, v22, v25, v21, v56, v57, v57, v57, v24, base.B2i32(v20 != v57))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int32(0)
			} else {
				m.G0 = v18 + int32(16)
				return v62
			}
		}
	}
}
