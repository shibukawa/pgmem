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
								F_errfinish(m, int32(_a_F_get_multirange_io_data_2), int32(432), int32(_a_F_get_multirange_io_data_3))
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
													F_errfinish(m, int32(_a_F_get_multirange_io_data_2), int32(451), int32(_a_F_get_multirange_io_data_3))
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
													F_errfinish(m, int32(_a_F_get_multirange_io_data_2), int32(456), int32(_a_F_get_multirange_io_data_3))
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
							F_errfinish(m, int32(_a_F_get_multirange_io_data_2), int32(432), int32(_a_F_get_multirange_io_data_3))
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
												F_errfinish(m, int32(_a_F_get_multirange_io_data_2), int32(451), int32(_a_F_get_multirange_io_data_3))
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
												F_errfinish(m, int32(_a_F_get_multirange_io_data_2), int32(456), int32(_a_F_get_multirange_io_data_3))
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
func F_multirange_after_range(m *base.Module, l0 int32) int32 {
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
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
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
						v66 = v34
						m.G0 = v9 + int32(48)
						return v66
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
						if v44 == int32(0) {
							v66 = v34
							m.G0 = v9 + int32(48)
							return v66
						} else {
							v50 = v9 + int32(32)
							F_range_deserialize(m, v33, v17, v9+int32(40), v50, v9+int32(15))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v57 = v9 + int32(24)
								F_multirange_get_bounds(m, v33, v12, int32(0), v57, v9+int32(16))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v62 = F_range_cmp_bounds(m, v33, v50, v57)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										v66 = int32(base.Ui32(v62) >> (uint(int32(31)) % 32))
										m.G0 = v9 + int32(48)
										return v66
									}
								}
							}
						}
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_after_range_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_multirange_after_range_1), v9)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_after_range_2), int32(558), int32(_a_F_multirange_after_range_3))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
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
								v66 = v34
								m.G0 = v9 + int32(48)
								return v66
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								if v44 == int32(0) {
									v66 = v34
									m.G0 = v9 + int32(48)
									return v66
								} else {
									v50 = v9 + int32(32)
									F_range_deserialize(m, v33, v17, v9+int32(40), v50, v9+int32(15))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v57 = v9 + int32(24)
										F_multirange_get_bounds(m, v33, v12, int32(0), v57, v9+int32(16))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											v62 = F_range_cmp_bounds(m, v33, v50, v57)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												v66 = int32(base.Ui32(v62) >> (uint(int32(31)) % 32))
												m.G0 = v9 + int32(48)
												return v66
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_after_range_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_multirange_after_range_1), v9)
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_after_range_2), int32(558), int32(_a_F_multirange_after_range_3))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
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
							v66 = v34
							m.G0 = v9 + int32(48)
							return v66
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							if v44 == int32(0) {
								v66 = v34
								m.G0 = v9 + int32(48)
								return v66
							} else {
								v50 = v9 + int32(32)
								F_range_deserialize(m, v33, v17, v9+int32(40), v50, v9+int32(15))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v57 = v9 + int32(24)
									F_multirange_get_bounds(m, v33, v12, int32(0), v57, v9+int32(16))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v62 = F_range_cmp_bounds(m, v33, v50, v57)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v66 = int32(base.Ui32(v62) >> (uint(int32(31)) % 32))
											m.G0 = v9 + int32(48)
											return v66
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
func F_multirange_cmp(m *base.Module, l0 int32) int32 {
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
	var v24 int32
	_ = v24
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
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
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v112 int32
	_ = v112
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v24 = F_pg_detoast_datum(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v26 == v27 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L54
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	if v30 != 0 {
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
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L51
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v43 < v42 {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v31 == v26 {
		v41 = v30
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v34 = F_lookup_type_cache(m, v26, int32(_a_F_multirange_cmp_0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+296))
	if v36 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v34
	v41 = v34
	goto L8
L15:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v124 != v19 {
		goto L43
	} else {
		goto L44
	}
L16:
	;
	v45 = v42
	goto L18
L17:
	;
	v45 = v43
	goto L18
L18:
	;
	if int32(0) < v45 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v48 = int32(0)
	if v48 < v43 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v112 = int32(0)
	goto L15
L22:
	;
	v52 = v43
	goto L24
L23:
	;
	v52 = v48
	goto L24
L24:
	;
	v53 = int32(0)
	if v53 < v42 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v56 = v42
	goto L27
L26:
	;
	v56 = v53
	goto L27
L27:
	;
	v59 = v48
	goto L28
L28:
	;
	if v59 == v56 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L21
L30:
	;
	v112 = int32(-1)
	goto L15
L31:
	;
	goto L32
L32:
	;
	if v59 == v52 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v112 = int32(1)
	goto L15
L34:
	;
	goto L35
L35:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v41)+296))
	v76 = v16 + int32(40)
	v78 = v16 + int32(32)
	F_multirange_get_bounds(m, v74, v19, v59, v76, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v41)+296))
	v83 = v16 + int32(24)
	v85 = v16 + int32(16)
	F_multirange_get_bounds(m, v81, v24, v59, v83, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v41)+296))
	v89 = F_range_cmp_bounds(m, v88, v76, v83)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v89 != 0 {
		v112 = v89
		goto L15
	} else {
		goto L39
	}
L39:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v41)+296))
	v92 = F_range_cmp_bounds(m, v91, v78, v85)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v92 != 0 {
		v112 = v92
		goto L15
	} else {
		goto L41
	}
L41:
	;
	v95 = v59 + int32(1)
	if v95 != v45 {
		v59 = v95
		goto L28
	} else {
		goto L42
	}
L42:
	;
	goto L29
L43:
	;
	F_pfree(m, v19)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v128 != v24 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	F_pfree(m, v24)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	m.G0 = v16 + int32(48)
	return v112
L50:
	;
	goto L49
L51:
	;
	F_errmsg_internal(m, int32(_a_F_multirange_cmp_1), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_multirange_cmp_2), int32(2589), int32(_a_F_multirange_cmp_3))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v26
	F_errmsg_internal(m, int32(_a_F_multirange_cmp_4), v16)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_multirange_cmp_2), int32(558), int32(_a_F_multirange_cmp_5))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
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
	v12 = v10 - int32(32)
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
	m.G0 = v12 + int32(32)
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
	F_multirange_get_bounds(m, l0, l1, int32(0), v12+int32(24), v12+int32(16))
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
	v43 = v12 + int32(8)
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
	v48 = F_range_cmp_bounds(m, l0, v12+int32(16), v43)
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
	v91 = F_range_cmp_bounds(m, l0, v12+int32(24), v12+int32(8))
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
	v68 = v12 + int32(16)
	F_multirange_get_bounds(m, l0, l1, v62, v12+int32(24), v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v73 = F_range_cmp_bounds(m, l0, v68, v12+int32(8))
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
	v97 = F_range_cmp_bounds(m, l0, v12+int32(16), v12)
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	v4 = int32(0)
	v12 = l1 + int32(8)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
	if l2 <= v4 {
		v43 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v52 = v50 << (uint(int32(2)) % 32)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+(v12+v52)))))
	v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+8)))
	if v14 == int32(105) {
		v82 = (v50*int32(5) + int32(11)) & int32(-4)
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v20 = v4
	v21 = l2
	goto L3
L3:
	;
	v27 = int32(2)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v12+v21<<(uint(v27)%32))))
	v33 = v30&int32(2147483647) + v20
	if base.Ui32(v21) < base.Ui32(v27) {
		v43 = v33
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v43 = v33
	goto L1
L5:
	;
	if int32(0) <= v30 {
		v20 = v33
		v21 = v21 - int32(1)
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	v201 = v200 - v84
	v203 = v201 + int32(9)
	v204 = F_palloc0(m, v203)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L57
	} else {
		goto L58
	}
L8:
	;
	v194 = F_strlen(m, v157)
	mBase = m.M
	v200 = v194 + v157 + int32(1)
	goto L7
L9:
	;
	v167 = v164 & int32(255)
	if v167 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L10:
	;
	if v14 == int32(105) {
		goto L40
	} else {
		goto L41
	}
L11:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v135 != 0 {
		v164 = v135
		v165 = v134
		goto L9
	} else {
		goto L38
	}
L12:
	;
	v130 = v84 + v129
	if v55&int32(80) != 0 {
		v200 = v130
		goto L7
	} else {
		goto L37
	}
L13:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v129 = int32(base.Ui32(v125) >> (uint(int32(2)) % 32))
	goto L12
L14:
	;
	if v55&int32(81) != 0 {
		v200 = v119
		goto L7
	} else {
		goto L35
	}
L15:
	;
	v84 = v82 + l1 + v43
	if v55&int32(41) != 0 {
		v119 = v84
		goto L14
	} else {
		goto L20
	}
L16:
	;
	switch v14 - int32(99) {
	case 0:
		goto L19
	case 1:
		goto L18
	default:
		goto L17
	}
L17:
	;
	v82 = (v50*int32(5) + int32(9)) & int32(-2)
	goto L15
L18:
	;
	v82 = (v50*int32(5) + int32(15)) & int32(-8)
	goto L15
L19:
	;
	v82 = v50 + v52 + int32(8)
	goto L15
L20:
	;
	if int32(0) < v56 {
		v119 = v84 + v56
		goto L14
	} else {
		goto L21
	}
L21:
	;
	if v56 == int32(-1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v92 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v115 = F_strlen(m, v84)
	mBase = m.M
	v119 = v115 + v84 + int32(1)
	goto L14
L25:
	;
	v96 = int32(18)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	if v98 == v96 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if v92&int32(1) == int32(0) {
		goto L13
	} else {
		goto L34
	}
L28:
	;
	v101 = v96
	goto L30
L29:
	;
	v101 = int32(2)
	goto L30
L30:
	;
	if base.Ui32((v98-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v108 = int32(6)
	goto L33
L32:
	;
	v108 = v101
	goto L33
L33:
	;
	v129 = v108
	goto L12
L34:
	;
	v129 = int32(base.Ui32(v92) >> (uint(int32(1)) % 32))
	goto L12
L35:
	;
	if v56 == int32(-1) {
		v134 = v119
		goto L11
	} else {
		goto L36
	}
L36:
	;
	v138 = v56
	v139 = v119
	v140 = int32(0)
	goto L10
L37:
	;
	v134 = v130
	goto L11
L38:
	;
	v138 = int32(-1)
	v139 = v134
	v140 = int32(1)
	goto L10
L39:
	;
	if int32(0) < v56 {
		v200 = v138 + v157
		goto L7
	} else {
		goto L45
	}
L40:
	;
	v157 = (v139 + int32(3)) & int32(-4)
	goto L39
L41:
	;
	goto L42
L42:
	;
	switch v14 - int32(99) {
	case 0:
		v157 = v139
		goto L39
	case 1:
		goto L44
	default:
		goto L43
	}
L43:
	;
	v157 = (v139 + int32(1)) & int32(-2)
	goto L39
L44:
	;
	v157 = (v139 + int32(7)) & int32(-8)
	goto L39
L45:
	;
	if v140 == int32(0) {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	v164 = v163
	v165 = v157
	goto L9
L47:
	;
	v171 = int32(18)
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
	if v173 == v171 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v185 = int32(1)
	if v167&v185 != 0 {
		v200 = v165 + int32(base.Ui32(v167)>>(uint(v185)%32))
		goto L7
	} else {
		goto L56
	}
L50:
	;
	v176 = v171
	goto L52
L51:
	;
	v176 = int32(2)
	goto L52
L52:
	;
	if base.Ui32((v173-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v183 = int32(6)
	goto L55
L54:
	;
	v183 = v176
	goto L55
L55:
	;
	v200 = v165 + v183
	goto L7
L56:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v200 = v165 + int32(base.Ui32(v190)>>(uint(int32(2))%32))
	goto L7
L57:
	;
	return int32(0)
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v203 << (uint(int32(2)) % 32)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v204)+4)) = v211
	v214 = v204 + int32(8)
	if v201 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	base.MemoryCopy(m, v214, v84, v201)
	goto L61
L60:
	;
	goto L61
L61:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v214+v201))) = uint8(v55)
	return v204
}
func F_multirange_out(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v19 = F_get_multirange_io_data(m, l0, v17, int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_initStringInfo(m, v10)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_appendStringInfoChar(m, v10, int32(123))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v26 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_appendStringInfoChar(m, v10, int32(125))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L22
	}
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+296))
	v34 = F_palloc(m, v26<<(uint(int32(2))%32))
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
	v46 = F_multirange_get_range(m, v30, v13, v36)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v53 = v19 + int32(4)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v55 = F_OutputFunctionCall(m, v53, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34+v36<<(uint(int32(2))%32)))) = v46
	v50 = v36 + int32(1)
	if v50 != v26 {
		v36 = v50
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	F_appendStringInfoString(m, v10, v55)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v59 = int32(1)
	if v26 == v59 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v62 = v59
	goto L16
L16:
	;
	F_appendStringInfoChar(m, v10, int32(44))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L6
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v34+v62<<(uint(int32(2))%32))))
	v76 = F_OutputFunctionCall(m, v53, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_appendStringInfoString(m, v10, v76)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v81 = v62 + int32(1)
	if v81 != v26 {
		v62 = v81
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	m.G0 = v10 + int32(16)
	return v93
}
func F_multirange_overlaps_range(m *base.Module, l0 int32) int32 {
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
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
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
					v34 = F_range_overlaps_multirange_internal(m, v33, v17, v12)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v34
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_overlaps_range_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_multirange_overlaps_range_1), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_overlaps_range_2), int32(558), int32(_a_F_multirange_overlaps_range_3))
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
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v34 = F_range_overlaps_multirange_internal(m, v33, v17, v12)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v34
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_overlaps_range_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_multirange_overlaps_range_1), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_overlaps_range_2), int32(558), int32(_a_F_multirange_overlaps_range_3))
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
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						v34 = F_range_overlaps_multirange_internal(m, v33, v17, v12)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v34
						}
					}
				}
			}
		}
	}
}
func F_multirange_overright_range(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v2 = int32(0)
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
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			if v19 == int32(0) {
				v70 = v2
				m.G0 = v9 + int32(48)
				return v70
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				v28 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17+int32(base.Ui32(v22)>>(uint(int32(2))%32))-int32(1)))))
				if v28&int32(1) != 0 {
					v70 = v2
					m.G0 = v9 + int32(48)
					return v70
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
					if v33 != 0 {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						if v34 == v31 {
							v44 = v33
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
							v48 = v9 + int32(40)
							F_multirange_get_bounds(m, v45, v12, int32(0), v48, v9+int32(32))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
								v55 = v9 + int32(24)
								F_range_deserialize(m, v53, v17, v55, v9+int32(16), v9+int32(15))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
									v63 = F_range_cmp_bounds(m, v62, v48, v55)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										v70 = int32(base.Ui32(v63^int32(-1)) >> (uint(int32(31)) % 32))
										m.G0 = v9 + int32(48)
										return v70
									}
								}
							}
						} else {
							v37 = F_lookup_type_cache(m, v31, int32(_a_F_multirange_overright_range_0))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+296))
								if v39 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v31
										F_errmsg_internal(m, int32(_a_F_multirange_overright_range_1), v9)
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_multirange_overright_range_2), int32(558), int32(_a_F_multirange_overright_range_3))
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
									v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v37
									v44 = v37
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
									v48 = v9 + int32(40)
									F_multirange_get_bounds(m, v45, v12, int32(0), v48, v9+int32(32))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
										v55 = v9 + int32(24)
										F_range_deserialize(m, v53, v17, v55, v9+int32(16), v9+int32(15))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											v62 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
											v63 = F_range_cmp_bounds(m, v62, v48, v55)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												v70 = int32(base.Ui32(v63^int32(-1)) >> (uint(int32(31)) % 32))
												m.G0 = v9 + int32(48)
												return v70
											}
										}
									}
								}
							}
						}
					} else {
						v37 = F_lookup_type_cache(m, v31, int32(_a_F_multirange_overright_range_0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+296))
							if v39 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v31
									F_errmsg_internal(m, int32(_a_F_multirange_overright_range_1), v9)
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_multirange_overright_range_2), int32(558), int32(_a_F_multirange_overright_range_3))
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
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v37
								v44 = v37
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
								v48 = v9 + int32(40)
								F_multirange_get_bounds(m, v45, v12, int32(0), v48, v9+int32(32))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
									v55 = v9 + int32(24)
									F_range_deserialize(m, v53, v17, v55, v9+int32(16), v9+int32(15))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
										v63 = F_range_cmp_bounds(m, v62, v48, v55)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int32(0)
										} else {
											v70 = int32(base.Ui32(v63^int32(-1)) >> (uint(int32(31)) % 32))
											m.G0 = v9 + int32(48)
											return v70
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
func F_multirange_send(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v31 int32
	_ = v31
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v15 = F_makeStringInfo(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = F_get_multirange_io_data(m, l0, v14, int32(3))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_pq_begintypsend(m, v15)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	F_enlargeStringInfo(m, v15, int32(4))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v31 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v26+v27))) = base.I32_rotr(v22, int32(24))&v31 | base.I32_rotr(v22&v31, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v26 + int32(4)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if int32(0) < v42 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
	v50 = F_palloc(m, v42<<(uint(int32(2))%32))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v132 << (uint(int32(2)) % 32)
	goto L21
L10:
	;
	v52 = int32(0)
	goto L11
L11:
	;
	v63 = F_multirange_get_range(m, v46, v10, v52)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v72 = int32(0)
	goto L15
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50+v52<<(uint(int32(2))%32)))) = v63
	v67 = v52 + int32(1)
	if v67 != v42 {
		v52 = v67
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v50+v72<<(uint(int32(2))%32))))
	v84 = F_SendFunctionCall(m, v18+int32(4), v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L9
L17:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	F_enlargeStringInfo(m, v15, int32(4))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v93 = int32(2)
	v95 = int32(4)
	v96 = int32(base.Ui32(v86)>>(uint(v93)%32)) - v95
	v97 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v90+v91))) = base.I32_rotr(v96&v97, int32(8)) | base.I32_rotr(v96, int32(24))&v97
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v90 + v95
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	F_appendBinaryStringInfo(m, v15, v84+v95, int32(base.Ui32(v112)>>(uint(v93)%32))-v95)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v120 = v72 + int32(1)
	if v120 != v42 {
		v72 = v120
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	return v131
}
