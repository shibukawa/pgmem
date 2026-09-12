package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_multirange_before_range(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
					v34 = int32(0)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
					if v41&int32(1) != 0 {
						v72 = v34
						m.G0 = v9 + int32(48)
						return v72
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
						if v44 == int32(0) {
							v72 = v34
							m.G0 = v9 + int32(48)
							return v72
						} else {
							F_range_deserialize(m, v33, v17, v9+int32(40), v9+int32(32), v9+int32(15))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								F_multirange_get_bounds(m, v33, v12, v55-int32(1), v9+int32(24), v9+int32(16))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									v68 = F_range_cmp_bounds(m, v33, v9+int32(40), v9+int32(16))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										v72 = base.B2i32(int32(0) < v68)
										m.G0 = v9 + int32(48)
										return v72
									}
								}
							}
						}
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(65536))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(362047), v9)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(483051), int32(558), int32(390036))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v34 = int32(0)
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
							v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
							if v41&int32(1) != 0 {
								v72 = v34
								m.G0 = v9 + int32(48)
								return v72
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								if v44 == int32(0) {
									v72 = v34
									m.G0 = v9 + int32(48)
									return v72
								} else {
									F_range_deserialize(m, v33, v17, v9+int32(40), v9+int32(32), v9+int32(15))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										F_multirange_get_bounds(m, v33, v12, v55-int32(1), v9+int32(24), v9+int32(16))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v68 = F_range_cmp_bounds(m, v33, v9+int32(40), v9+int32(16))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												v72 = base.B2i32(int32(0) < v68)
												m.G0 = v9 + int32(48)
												return v72
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(65536))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(362047), v9)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(483051), int32(558), int32(390036))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						v34 = int32(0)
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
						v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
						if v41&int32(1) != 0 {
							v72 = v34
							m.G0 = v9 + int32(48)
							return v72
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							if v44 == int32(0) {
								v72 = v34
								m.G0 = v9 + int32(48)
								return v72
							} else {
								F_range_deserialize(m, v33, v17, v9+int32(40), v9+int32(32), v9+int32(15))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									F_multirange_get_bounds(m, v33, v12, v55-int32(1), v9+int32(24), v9+int32(16))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										v68 = F_range_cmp_bounds(m, v33, v9+int32(40), v9+int32(16))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											v72 = base.B2i32(int32(0) < v68)
											m.G0 = v9 + int32(48)
											return v72
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
func F_multirange_constructor1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = F_get_fn_expr_rettype(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
		if v15 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			if v16 == v10 {
				v26 = v15
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
				if v27 == int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(132626), int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(483051), int32(1042), int32(541061))
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
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+296))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v32 = F_pg_detoast_datum(m, v31)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v32
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
						if v35 != v36 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v35
								F_errmsg_internal(m, int32(359971), v7+int32(16))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(483051), int32(1049), int32(541061))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v41 = F_make_multirange(m, v10, v30, int32(1), v7+int32(28))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(32)
								return v41
							}
						}
					}
				}
			} else {
				v19 = F_lookup_type_cache(m, v10, int32(65536))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+296))
					if v21 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10
							F_errmsg_internal(m, int32(362047), v7)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(483051), int32(558), int32(390036))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v19
						v26 = v19
						v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
						if v27 == int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(132626), int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(483051), int32(1042), int32(541061))
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
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+296))
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v32 = F_pg_detoast_datum(m, v31)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v32
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
								if v35 != v36 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v35
										F_errmsg_internal(m, int32(359971), v7+int32(16))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(483051), int32(1049), int32(541061))
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v41 = F_make_multirange(m, v10, v30, int32(1), v7+int32(28))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										m.G0 = v7 + int32(32)
										return v41
									}
								}
							}
						}
					}
				}
			}
		} else {
			v19 = F_lookup_type_cache(m, v10, int32(65536))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+296))
				if v21 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10
						F_errmsg_internal(m, int32(362047), v7)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(483051), int32(558), int32(390036))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v19
					v26 = v19
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
					if v27 == int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(132626), int32(0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(483051), int32(1042), int32(541061))
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
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+296))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v32 = F_pg_detoast_datum(m, v31)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v32
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
							if v35 != v36 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v35
									F_errmsg_internal(m, int32(359971), v7+int32(16))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(483051), int32(1049), int32(541061))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v41 = F_make_multirange(m, v10, v30, int32(1), v7+int32(28))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(32)
									return v41
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_multirange_contains_elem(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
		if v19 != 0 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v20 == v17 {
				v30 = v19
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+296))
				v32 = F_multirange_contains_elem_internal(m, v31, v12, v16)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 + int32(16)
					return v32
				}
			} else {
				v23 = F_lookup_type_cache(m, v17, int32(65536))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+296))
					if v25 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
							F_errmsg_internal(m, int32(362047), v9)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(483051), int32(558), int32(390036))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
						v30 = v23
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+296))
						v32 = F_multirange_contains_elem_internal(m, v31, v12, v16)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v32
						}
					}
				}
			}
		} else {
			v23 = F_lookup_type_cache(m, v17, int32(65536))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+296))
				if v25 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
						F_errmsg_internal(m, int32(362047), v9)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(483051), int32(558), int32(390036))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
					v30 = v23
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+296))
					v32 = F_multirange_contains_elem_internal(m, v31, v12, v16)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v32
					}
				}
			}
		}
	}
}
func F_multirange_intersect(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L36
	}
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+296))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v38 != 0 {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v26 == v23 {
		v36 = v25
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v29 = F_lookup_type_cache(m, v23, int32(65536))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+296))
	if v31 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v29
	v36 = v29
	goto L5
L12:
	;
	m.G0 = v13 + int32(16)
	return v130
L13:
	;
	if int32(0) < v38 {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v39 != 0 {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v41 = int32(0)
	v43 = F_make_multirange(m, v23, v37, v41, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v130 = v43
	goto L12
L19:
	;
	v50 = F_palloc(m, v38<<(uint(int32(2))%32))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v75 = v39
	v81 = v2
	goto L21
L21:
	;
	if int32(0) < v75 {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v52 = int32(0)
	goto L23
L23:
	;
	v65 = F_multirange_get_range(m, v37, v16, v52)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v75 = v71
	v81 = v50
	goto L21
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50+v52<<(uint(int32(2))%32)))) = v65
	v69 = v52 + int32(1)
	if v69 != v38 {
		v52 = v69
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v87 = F_palloc(m, v75<<(uint(int32(2))%32))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	v116 = v2
	goto L29
L29:
	;
	v118 = F_multirange_intersect_internal(m, v23, v37, v38, v81, v75, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L35
	}
L30:
	;
	v89 = int32(0)
	goto L31
L31:
	;
	v102 = F_multirange_get_range(m, v37, v21, v89)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	v116 = v87
	goto L29
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87+v89<<(uint(int32(2))%32)))) = v102
	v106 = v89 + int32(1)
	if v106 != v75 {
		v89 = v106
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v130 = v118
	goto L12
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v23
	F_errmsg_internal(m, int32(362047), v13)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(483051), int32(558), int32(390036))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_multirange_lower(m *base.Module, l0 int32) int32 {
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
		if v15 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v49 = int32(0)
			m.G0 = v8 + int32(32)
			return v49
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
			if v22 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				if v23 == v20 {
					v33 = v22
					v34 = int32(0)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
					F_multirange_get_bounds(m, v35, v11, v34, v8+int32(24), v8+int32(16))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
						if v43 == int32(0) {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
							v49 = v46
						} else {
							v47 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v47)
							v49 = v34
						}
						m.G0 = v8 + int32(32)
						return v49
					}
				} else {
					v26 = F_lookup_type_cache(m, v20, int32(65536))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+296))
						if v28 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v20
								F_errmsg_internal(m, int32(362047), v8)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(483051), int32(558), int32(390036))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
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
							v34 = int32(0)
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
							F_multirange_get_bounds(m, v35, v11, v34, v8+int32(24), v8+int32(16))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
								if v43 == int32(0) {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
									v49 = v46
								} else {
									v47 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v47)
									v49 = v34
								}
								m.G0 = v8 + int32(32)
								return v49
							}
						}
					}
				}
			} else {
				v26 = F_lookup_type_cache(m, v20, int32(65536))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+296))
					if v28 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v20
							F_errmsg_internal(m, int32(362047), v8)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(483051), int32(558), int32(390036))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
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
						v34 = int32(0)
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
						F_multirange_get_bounds(m, v35, v11, v34, v8+int32(24), v8+int32(16))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
							if v43 == int32(0) {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
								v49 = v46
							} else {
								v47 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v47)
								v49 = v34
							}
							m.G0 = v8 + int32(32)
							return v49
						}
					}
				}
			}
		}
	}
}
func F_multirange_lower_inf(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
		if v15 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
			if v18 != 0 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				if v19 == v16 {
					v29 = v18
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+296))
					F_multirange_get_bounds(m, v30, v11, int32(0), v8+int32(24), v8+int32(16))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
						v42 = v38
						m.G0 = v8 + int32(32)
						return v42
					}
				} else {
					v22 = F_lookup_type_cache(m, v16, int32(65536))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+296))
						if v24 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
								F_errmsg_internal(m, int32(362047), v8)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(483051), int32(558), int32(390036))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v22
							v29 = v22
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+296))
							F_multirange_get_bounds(m, v30, v11, int32(0), v8+int32(24), v8+int32(16))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
								v42 = v38
								m.G0 = v8 + int32(32)
								return v42
							}
						}
					}
				}
			} else {
				v22 = F_lookup_type_cache(m, v16, int32(65536))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+296))
					if v24 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
							F_errmsg_internal(m, int32(362047), v8)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(483051), int32(558), int32(390036))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v22
						v29 = v22
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+296))
						F_multirange_get_bounds(m, v30, v11, int32(0), v8+int32(24), v8+int32(16))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
							v42 = v38
							m.G0 = v8 + int32(32)
							return v42
						}
					}
				}
			}
		} else {
			v42 = int32(0)
			m.G0 = v8 + int32(32)
			return v42
		}
	}
}
