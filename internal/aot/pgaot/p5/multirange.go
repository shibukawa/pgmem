package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_multirange_io_data(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v12 != 0 {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		if v14 == l1 {
			v81 = v12
			m.G0 = v9 + int32(48)
			return v81
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
			v18 = F_MemoryContextAlloc(m, v16, int32(36))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v23 = F_lookup_type_cache(m, l1, int32(_a_F_get_multirange_io_data_0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v23
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+296))
					if v26 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
							F_errmsg_internal(m, int32(_a_F_get_multirange_io_data_1), v9)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_get_multirange_io_data_2), int32(435), int32(_a_F_get_multirange_io_data_3))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
						F_get_type_io_data(m, v29, l2, v9+int32(42), v9+int32(41), v9+int32(40), v9+int32(39), v18+int32(32), v9+int32(44))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
							if v44 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(52461700))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+296))
										v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
										v57 = F_format_type_be(m, v56)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
										} else {
											if l2 == int32(2) {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v57
												F_errmsg(m, int32(_a_F_get_multirange_io_data_4), v9+int32(16))
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_get_multirange_io_data_2), int32(454), int32(_a_F_get_multirange_io_data_3))
													mBase = m.M
													v110 = m.ExcPending
													if v110 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v57
												F_errmsg(m, int32(_a_F_get_multirange_io_data_5), v9+int32(32))
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_get_multirange_io_data_2), int32(459), int32(_a_F_get_multirange_io_data_3))
													mBase = m.M
													v71 = m.ExcPending
													if v71 != 0 {
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
							} else {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
								F_fmgr_info_cxt(m, v44, v18+int32(4), v75)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v18
									v81 = v18
									m.G0 = v9 + int32(48)
									return v81
								}
							}
						}
					}
				}
			}
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
		v18 = F_MemoryContextAlloc(m, v16, int32(36))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v23 = F_lookup_type_cache(m, l1, int32(_a_F_get_multirange_io_data_0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v23
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+296))
				if v26 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
						F_errmsg_internal(m, int32(_a_F_get_multirange_io_data_1), v9)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_get_multirange_io_data_2), int32(435), int32(_a_F_get_multirange_io_data_3))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
					F_get_type_io_data(m, v29, l2, v9+int32(42), v9+int32(41), v9+int32(40), v9+int32(39), v18+int32(32), v9+int32(44))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
						if v44 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(52461700))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+296))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
									v57 = F_format_type_be(m, v56)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										if l2 == int32(2) {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v57
											F_errmsg(m, int32(_a_F_get_multirange_io_data_4), v9+int32(16))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_get_multirange_io_data_2), int32(454), int32(_a_F_get_multirange_io_data_3))
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v57
											F_errmsg(m, int32(_a_F_get_multirange_io_data_5), v9+int32(32))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_get_multirange_io_data_2), int32(459), int32(_a_F_get_multirange_io_data_3))
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
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
						} else {
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
							F_fmgr_info_cxt(m, v44, v18+int32(4), v75)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v18
								v81 = v18
								m.G0 = v9 + int32(48)
								return v81
							}
						}
					}
				}
			}
		}
	}
}
func F_multirange_after_range(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int64
	_ = v69
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v7 = int64(0)
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
			if v22 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				if v23 == v20 {
					v33 = v22
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
					v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v18+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
					if v41&int32(1) != 0 {
						v69 = v7
						m.G0 = v10 + int32(80)
						return v69
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
						if v44 == int32(0) {
							v69 = v7
							m.G0 = v10 + int32(80)
							return v69
						} else {
							v50 = v10 + int32(48)
							F_range_deserialize(m, v34, v18, v10-int32(-64), v50, v10+int32(15))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int64(0)
							} else {
								v57 = v10 + int32(32)
								F_multirange_get_bounds(m, v34, v13, int32(0), v57, v10+int32(16))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int64(0)
								} else {
									v62 = F_range_cmp_bounds(m, v34, v50, v57)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int64(0)
									} else {
										v69 = base.I64_extend_i32_u(int32(base.Ui32(v62) >> (uint(int32(31)) % 32)))
										m.G0 = v10 + int32(80)
										return v69
									}
								}
							}
						}
					}
				} else {
					v26 = F_lookup_type_cache(m, v20, int32(_a_F_multirange_after_range_0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+296))
						if v28 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
								F_errmsg_internal(m, int32(_a_F_multirange_after_range_1), v10)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_after_range_2), int32(561), int32(_a_F_multirange_after_range_3))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int64(0)
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
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
							v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v18+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
							if v41&int32(1) != 0 {
								v69 = v7
								m.G0 = v10 + int32(80)
								return v69
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
								if v44 == int32(0) {
									v69 = v7
									m.G0 = v10 + int32(80)
									return v69
								} else {
									v50 = v10 + int32(48)
									F_range_deserialize(m, v34, v18, v10-int32(-64), v50, v10+int32(15))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int64(0)
									} else {
										v57 = v10 + int32(32)
										F_multirange_get_bounds(m, v34, v13, int32(0), v57, v10+int32(16))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int64(0)
										} else {
											v62 = F_range_cmp_bounds(m, v34, v50, v57)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int64(0)
											} else {
												v69 = base.I64_extend_i32_u(int32(base.Ui32(v62) >> (uint(int32(31)) % 32)))
												m.G0 = v10 + int32(80)
												return v69
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v26 = F_lookup_type_cache(m, v20, int32(_a_F_multirange_after_range_0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+296))
					if v28 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
							F_errmsg_internal(m, int32(_a_F_multirange_after_range_1), v10)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_after_range_2), int32(561), int32(_a_F_multirange_after_range_3))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int64(0)
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
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
						v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v18+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
						if v41&int32(1) != 0 {
							v69 = v7
							m.G0 = v10 + int32(80)
							return v69
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
							if v44 == int32(0) {
								v69 = v7
								m.G0 = v10 + int32(80)
								return v69
							} else {
								v50 = v10 + int32(48)
								F_range_deserialize(m, v34, v18, v10-int32(-64), v50, v10+int32(15))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int64(0)
								} else {
									v57 = v10 + int32(32)
									F_multirange_get_bounds(m, v34, v13, int32(0), v57, v10+int32(16))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int64(0)
									} else {
										v62 = F_range_cmp_bounds(m, v34, v50, v57)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int64(0)
										} else {
											v69 = base.I64_extend_i32_u(int32(base.Ui32(v62) >> (uint(int32(31)) % 32)))
											m.G0 = v10 + int32(80)
											return v69
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
func F_multirange_cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v118 int64
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	v14 = int64(0)
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v27 == v28 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L55
	}
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	if v31 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L52
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v44 < v43 {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v32 == v27 {
		v42 = v31
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v35 = F_lookup_type_cache(m, v27, int32(_a_F_multirange_cmp_0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+296))
	if v37 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v35
	v42 = v35
	goto L8
L15:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v119 != v20 {
		goto L44
	} else {
		goto L45
	}
L16:
	;
	v46 = v43
	goto L18
L17:
	;
	v46 = v44
	goto L18
L18:
	;
	if v46 <= int32(0) {
		v118 = v14
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v49 = int32(0)
	if v49 < v44 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v53 = v44
	goto L22
L21:
	;
	v53 = v49
	goto L22
L22:
	;
	v54 = int32(0)
	if v54 < v43 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v57 = v43
	goto L25
L24:
	;
	v57 = v54
	goto L25
L25:
	;
	v59 = v49
	goto L26
L26:
	;
	if v59 == v57 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v118 = v14
	goto L15
L28:
	;
	v118 = int64(-1)
	goto L15
L29:
	;
	goto L30
L30:
	;
	if v59 == v53 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v118 = int64(1)
	goto L15
L32:
	;
	goto L33
L33:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
	v78 = v17 - int32(-64)
	v80 = v17 + int32(48)
	F_multirange_get_bounds(m, v76, v20, v59, v78, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
	v85 = v17 + int32(32)
	v87 = v17 + int32(16)
	F_multirange_get_bounds(m, v83, v25, v59, v85, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
	v91 = F_range_cmp_bounds(m, v90, v78, v85)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	v103 = v59 + int32(1)
	if v103 != v46 {
		v59 = v103
		goto L26
	} else {
		goto L43
	}
L37:
	;
	if v91 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
	v96 = F_range_cmp_bounds(m, v95, v80, v87)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	v100 = v91
	goto L40
L40:
	;
	v118 = base.I64_extend_i32_s(v100)
	goto L15
L41:
	;
	if v96 == int32(0) {
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v100 = v96
	goto L40
L43:
	;
	goto L27
L44:
	;
	F_pfree(m, v20)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v123 != v25 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	F_pfree(m, v25)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	m.G0 = v17 + int32(80)
	return v118
L51:
	;
	goto L50
L52:
	;
	F_errmsg_internal(m, int32(_a_F_multirange_cmp_1), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_multirange_cmp_2), int32(2664), int32(_a_F_multirange_cmp_3))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v27
	F_errmsg_internal(m, int32(_a_F_multirange_cmp_4), v17)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_multirange_cmp_2), int32(561), int32(_a_F_multirange_cmp_5))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_multirange_contains_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = int32(1)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v15 == v4 {
		v109 = v14
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 - int32(-64)
	return v109
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v18 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v109 = int32(0)
	goto L1
L4:
	;
	goto L5
L5:
	;
	F_multirange_get_bounds(m, l0, l1, int32(0), v10+int32(-16), v10+int32(-32))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v15 <= int32(0) {
		v109 = v14
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v39 = v4
	v40 = v4
	goto L9
L9:
	;
	v43 = v10 + int32(-48)
	F_multirange_get_bounds(m, l0, l2, v39, v43, v12)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v109 = v101
	goto L1
L11:
	;
	v48 = F_range_cmp_bounds(m, l0, v10+int32(-32), v43)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	if v48 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v59 = v40
	goto L16
L14:
	;
	v84 = v40
	goto L15
L15:
	;
	v86 = int32(0)
	v91 = F_range_cmp_bounds(m, l0, v10+int32(-16), v10+int32(-48))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L24
	}
L16:
	;
	v62 = v59 + int32(1)
	if v18 <= v62 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v84 = v62
	goto L15
L18:
	;
	v109 = int32(0)
	goto L1
L19:
	;
	goto L20
L20:
	;
	v68 = v10 + int32(-32)
	F_multirange_get_bounds(m, l0, l1, v62, v10+int32(-16), v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v73 = F_range_cmp_bounds(m, l0, v68, v10+int32(-48))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	if v73 < int32(0) {
		v59 = v62
		goto L16
	} else {
		goto L23
	}
L23:
	;
	goto L17
L24:
	;
	if int32(0) < v91 {
		v109 = v86
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v97 = F_range_cmp_bounds(m, l0, v10+int32(-32), v12)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	if v97 < int32(0) {
		v109 = v86
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v101 = int32(1)
	v103 = v39 + v101
	if v103 != v15 {
		v39 = v103
		v40 = v84
		goto L9
	} else {
		goto L28
	}
L28:
	;
	goto L10
}
func F_multirange_get_range(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v17 = l1 + int32(8)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v19 = int32(*(*int8)(unsafe.Add(mBase, uint32(v18)+11)))
	v21 = v19 & int32(255)
	if l2 <= v4 {
		v53 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+8)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v62 = v60 << (uint(int32(2)) % 32)
	v63 = v60 + v62
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v62+l2))))
	switch v21 - int32(99) {
	case 0:
		v96 = int32(-1)
		v97 = v63 + int32(8)
		goto L7
	case 1:
		goto L8
	default:
		goto L11
	case 6:
		goto L9
	case 16:
		goto L10
	}
L2:
	;
	v27 = l2
	v29 = v4
	goto L3
L3:
	;
	v35 = int32(2)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v17+v27<<(uint(v35)%32))))
	v41 = v38&int32(2147483647) + v29
	if base.Ui32(v27) < base.Ui32(v35) {
		v53 = v41
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v53 = v41
	goto L1
L5:
	;
	if int32(0) <= v38 {
		v27 = v27 - int32(1)
		v29 = v41
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	v100 = l1 + v96&v97 + v53
	if v68&int32(41) != 0 {
		v135 = v100
		goto L23
	} else {
		goto L24
	}
L8:
	;
	v96 = int32(-8)
	v97 = v63 + int32(15)
	goto L7
L9:
	;
	v96 = int32(-4)
	v97 = v63 + int32(11)
	goto L7
L10:
	;
	v96 = int32(-2)
	v97 = v63 + int32(9)
	goto L7
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v19
	F_errmsg_internal(m, int32(_a_F_multirange_get_range_0), v14)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_multirange_get_range_1), int32(322), int32(_a_F_multirange_get_range_2))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	v233 = v232 - v100
	v235 = v233 + int32(9)
	v236 = F_palloc0(m, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L12
	} else {
		goto L63
	}
L17:
	;
	v225 = F_strlen(m, v187)
	mBase = m.M
	v232 = v225 + v187 + int32(1)
	goto L16
L18:
	;
	v198 = v196 & int32(255)
	if v198 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L19:
	;
	switch v21 - int32(99) {
	case 0:
		v185 = v155
		v186 = int32(-1)
		goto L43
	case 1:
		goto L44
	default:
		goto L47
	case 6:
		goto L45
	case 16:
		goto L46
	}
L20:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	if v151 != 0 {
		v195 = v150
		v196 = v151
		goto L18
	} else {
		goto L42
	}
L21:
	;
	v146 = v100 + v145
	if v68&int32(80) != 0 {
		v232 = v146
		goto L16
	} else {
		goto L41
	}
L22:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v145 = int32(base.Ui32(v141) >> (uint(int32(2)) % 32))
	goto L21
L23:
	;
	if v68&int32(81) != 0 {
		v232 = v135
		goto L16
	} else {
		goto L39
	}
L24:
	;
	if int32(0) < v59 {
		v135 = v100 + v59
		goto L23
	} else {
		goto L25
	}
L25:
	;
	if v59 == int32(-1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v108 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v131 = F_strlen(m, v100)
	mBase = m.M
	v135 = v131 + v100 + int32(1)
	goto L23
L29:
	;
	v112 = int32(18)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	if v114 == v112 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	if v108&int32(1) == int32(0) {
		goto L22
	} else {
		goto L38
	}
L32:
	;
	v117 = v112
	goto L34
L33:
	;
	v117 = int32(2)
	goto L34
L34:
	;
	if base.Ui32((v114-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v124 = int32(6)
	goto L37
L36:
	;
	v124 = v117
	goto L37
L37:
	;
	v145 = v124
	goto L21
L38:
	;
	v145 = int32(base.Ui32(v108) >> (uint(int32(1)) % 32))
	goto L21
L39:
	;
	if v59 == int32(-1) {
		v150 = v135
		goto L20
	} else {
		goto L40
	}
L40:
	;
	v154 = v59
	v155 = v135
	v157 = int32(0)
	goto L19
L41:
	;
	v150 = v146
	goto L20
L42:
	;
	v154 = int32(-1)
	v155 = v150
	v157 = int32(1)
	goto L19
L43:
	;
	v187 = v185 & v186
	if int32(0) < v59 {
		v232 = v187 + v154
		goto L16
	} else {
		goto L51
	}
L44:
	;
	v185 = v155 + int32(7)
	v186 = int32(-8)
	goto L43
L45:
	;
	v185 = v155 + int32(3)
	v186 = int32(-4)
	goto L43
L46:
	;
	v185 = v155 + int32(1)
	v186 = int32(-2)
	goto L43
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v21
	F_errmsg_internal(m, int32(_a_F_multirange_get_range_0), v14+int32(16))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_multirange_get_range_1), int32(322), int32(_a_F_multirange_get_range_2))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	if v157 == int32(0) {
		goto L17
	} else {
		goto L52
	}
L52:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v195 = v187
	v196 = v193
	goto L18
L53:
	;
	v202 = int32(18)
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+1)))
	if v204 == v202 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v216 = int32(1)
	if v198&v216 != 0 {
		v232 = v195 + int32(base.Ui32(v198)>>(uint(v216)%32))
		goto L16
	} else {
		goto L62
	}
L56:
	;
	v207 = v202
	goto L58
L57:
	;
	v207 = int32(2)
	goto L58
L58:
	;
	if base.Ui32((v204-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v214 = int32(6)
	goto L61
L60:
	;
	v214 = v207
	goto L61
L61:
	;
	v232 = v195 + v214
	goto L16
L62:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	v232 = v195 + int32(base.Ui32(v221)>>(uint(int32(2))%32))
	goto L16
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v236))) = v235 << (uint(int32(2)) % 32)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v236)+4)) = v241
	v244 = v236 + int32(8)
	if v233 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	base.MemoryCopy(m, v244, v100, v233)
	goto L66
L65:
	;
	goto L66
L66:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v244+v233))) = uint8(v68)
	m.G0 = v14 + int32(32)
	return v236
}
func F_multirange_minus_multi(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int64
	_ = v171
	var v175 int32
	_ = v175
	var v192 int64
	_ = v192
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L7
	} else {
		goto L39
	}
L2:
	;
	m.G0 = v16 + int32(16)
	return v192
L3:
	;
	goto L6
L4:
	;
	goto L5
L5:
	;
	v32 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L9
	}
L6:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int64(0)
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = int32(2)
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
	v192 = int64(0)
	goto L2
L9:
	;
	v34 = int32(_a_F_multirange_minus_multi_0)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_multirange_minus_multi[0]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_multirange_minus_multi[0])) = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v40 = F_pg_detoast_datum(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v43 = F_pg_detoast_datum(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v47 = F_lookup_type_cache(m, v45, int32(_a_F_multirange_minus_multi_1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+296))
	if v49 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	if v52 == int32(0) {
		v145 = v40
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_multirange_minus_multi[0])) = v35
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	goto L34
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	if v55 == int32(0) {
		v145 = v40
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if int32(0) < v52 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v62 = F_palloc_mul(m, int32(4), v52)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L7
	} else {
		goto L20
	}
L18:
	;
	v90 = v55
	v96 = v2
	goto L19
L19:
	;
	if int32(0) < v90 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	v65 = int32(0)
	goto L21
L21:
	;
	v80 = F_multirange_get_range(m, v49, v40, v65)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L23
	}
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v90 = v86
	v96 = v62
	goto L19
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62+v65<<(uint(int32(2))%32)))) = v80
	v84 = v65 + int32(1)
	if v84 != v52 {
		v65 = v84
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v104 = F_palloc_mul(m, int32(4), v90)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L7
	} else {
		goto L28
	}
L26:
	;
	v138 = v2
	goto L27
L27:
	;
	v141 = F_multirange_minus_internal(m, v45, v49, v52, v96, v90, v138)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L7
	} else {
		goto L33
	}
L28:
	;
	v107 = int32(0)
	goto L29
L29:
	;
	v122 = F_multirange_get_range(m, v49, v43, v107)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L7
	} else {
		goto L31
	}
L30:
	;
	v138 = v104
	goto L27
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104+v107<<(uint(int32(2))%32)))) = v122
	v126 = v107 + int32(1)
	if v126 != v90 {
		v107 = v126
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v145 = v141
	goto L14
L34:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
	if v160 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L7
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v159)))
	*(*int64)(unsafe.Add(mBase, uint32(v159))) = v171 + int64(1)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+20)) = int32(1)
	v192 = base.I64_extend_i32_u(v145)
	goto L2
L38:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+20)) = int32(2)
	v168 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v168)
	v192 = int64(0)
	goto L2
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v45
	F_errmsg_internal(m, int32(_a_F_multirange_minus_multi_2), v16)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_multirange_minus_multi_3), int32(1277), int32(_a_F_multirange_minus_multi_4))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_multirange_out(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
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
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v20 = F_get_multirange_io_data(m, l0, v18, int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_initStringInfo(m, v11)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_appendStringInfoChar(m, v11, int32(123))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v27 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_appendStringInfoChar(m, v11, int32(125))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L22
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+296))
	v34 = F_palloc_mul(m, int32(4), v27)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v36 = int32(0)
	goto L9
L9:
	;
	v47 = F_multirange_get_range(m, v31, v14, v36)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v54 = v20 + int32(4)
	v55 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v34))))
	v56 = F_OutputFunctionCall(m, v54, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34+v36<<(uint(int32(2))%32)))) = v47
	v51 = v36 + int32(1)
	if v51 != v27 {
		v36 = v51
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	F_appendStringInfoString(m, v11, v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v60 = int32(1)
	if v27 == v60 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v63 = v60
	goto L16
L16:
	;
	F_appendStringInfoChar(m, v11, int32(44))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L6
L18:
	;
	v77 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v34+v63<<(uint(int32(2))%32)))))
	v78 = F_OutputFunctionCall(m, v54, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_appendStringInfoString(m, v11, v78)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v83 = v63 + int32(1)
	if v83 != v27 {
		v63 = v83
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v96 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11))))
	m.G0 = v11 + int32(16)
	return v96
}
func F_multirange_overlaps_range(m *base.Module, l0 int32) int64 {
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
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
					v34 = F_range_overlaps_multirange_internal(m, v33, v17, v12)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v34)
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_overlaps_range_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_multirange_overlaps_range_1), v9)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_overlaps_range_2), int32(561), int32(_a_F_multirange_overlaps_range_3))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int64(0)
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
							v34 = F_range_overlaps_multirange_internal(m, v33, v17, v12)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
							} else {
								m.G0 = v9 + int32(16)
								return base.I64_extend_i32_u(v34)
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_overlaps_range_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_multirange_overlaps_range_1), v9)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_overlaps_range_2), int32(561), int32(_a_F_multirange_overlaps_range_3))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int64(0)
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
						v34 = F_range_overlaps_multirange_internal(m, v33, v17, v12)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							m.G0 = v9 + int32(16)
							return base.I64_extend_i32_u(v34)
						}
					}
				}
			}
		}
	}
}
func F_multirange_overright_range(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int64
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v7 = int64(0)
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			if v20 == int32(0) {
				v73 = v7
				m.G0 = v10 + int32(80)
				return v73
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v29 = int32(*(*int8)(unsafe.Add(mBase, uint32(v18+int32(base.Ui32(v23)>>(uint(int32(2))%32))-int32(1)))))
				if v29&int32(1) != 0 {
					v73 = v7
					m.G0 = v10 + int32(80)
					return v73
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
					if v34 != 0 {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
						if v35 == v32 {
							v45 = v34
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
							v49 = v10 - int32(-64)
							F_multirange_get_bounds(m, v46, v13, int32(0), v49, v10+int32(48))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int64(0)
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
								v56 = v10 + int32(32)
								F_range_deserialize(m, v54, v18, v56, v10+int32(16), v10+int32(15))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int64(0)
								} else {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
									v64 = F_range_cmp_bounds(m, v63, v49, v56)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int64(0)
									} else {
										v73 = base.I64_extend_i32_u(base.B2i32(int32(0) <= v64))
										m.G0 = v10 + int32(80)
										return v73
									}
								}
							}
						} else {
							v38 = F_lookup_type_cache(m, v32, int32(_a_F_multirange_overright_range_0))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int64(0)
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
								if v40 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v32
										F_errmsg_internal(m, int32(_a_F_multirange_overright_range_1), v10)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_multirange_overright_range_2), int32(561), int32(_a_F_multirange_overright_range_3))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v38
									v45 = v38
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
									v49 = v10 - int32(-64)
									F_multirange_get_bounds(m, v46, v13, int32(0), v49, v10+int32(48))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int64(0)
									} else {
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
										v56 = v10 + int32(32)
										F_range_deserialize(m, v54, v18, v56, v10+int32(16), v10+int32(15))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int64(0)
										} else {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
											v64 = F_range_cmp_bounds(m, v63, v49, v56)
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int64(0)
											} else {
												v73 = base.I64_extend_i32_u(base.B2i32(int32(0) <= v64))
												m.G0 = v10 + int32(80)
												return v73
											}
										}
									}
								}
							}
						}
					} else {
						v38 = F_lookup_type_cache(m, v32, int32(_a_F_multirange_overright_range_0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
							if v40 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v32
									F_errmsg_internal(m, int32(_a_F_multirange_overright_range_1), v10)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_multirange_overright_range_2), int32(561), int32(_a_F_multirange_overright_range_3))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v38
								v45 = v38
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
								v49 = v10 - int32(-64)
								F_multirange_get_bounds(m, v46, v13, int32(0), v49, v10+int32(48))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int64(0)
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
									v56 = v10 + int32(32)
									F_range_deserialize(m, v54, v18, v56, v10+int32(16), v10+int32(15))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int64(0)
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
										v64 = F_range_cmp_bounds(m, v63, v49, v56)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int64(0)
										} else {
											v73 = base.I64_extend_i32_u(base.B2i32(int32(0) <= v64))
											m.G0 = v10 + int32(80)
											return v73
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
func F_multirange_send(m *base.Module, l0 int32) int64 {
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v86 int64
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	F_initStringInfo(m, v11)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = F_get_multirange_io_data(m, l0, v18, int32(3))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_pq_begintypsend(m, v11)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	F_enlargeStringInfo(m, v11, int32(4))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v35 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v30+v31))) = base.I32_rotr(v26, int32(24))&v35 | base.I32_rotr(v26&v35, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v30 + int32(4)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if int32(0) < v46 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+296))
	v53 = F_palloc_mul(m, int32(4), v46)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v135 << (uint(int32(2)) % 32)
	goto L21
L10:
	;
	v55 = int32(0)
	goto L11
L11:
	;
	v66 = F_multirange_get_range(m, v50, v14, v55)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v75 = int32(0)
	goto L15
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53+v55<<(uint(int32(2))%32)))) = v66
	v70 = v55 + int32(1)
	if v70 != v46 {
		v55 = v70
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v86 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v53+v75<<(uint(int32(2))%32)))))
	v87 = F_SendFunctionCall(m, v22+int32(4), v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L9
L17:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	F_enlargeStringInfo(m, v11, int32(4))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v96 = int32(2)
	v98 = int32(4)
	v99 = int32(base.Ui32(v89)>>(uint(v96)%32)) - v98
	v100 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v93+v94))) = base.I32_rotr(v99&v100, int32(8)) | base.I32_rotr(v99, int32(24))&v100
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v93 + v98
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	F_appendBinaryStringInfo(m, v11, v87+v98, int32(base.Ui32(v115)>>(uint(v96)%32))-v98)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v123 = v75 + int32(1)
	if v123 != v46 {
		v75 = v123
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	m.G0 = v11 + int32(16)
	return base.I64_extend_i32_u(v134)
}
