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
				v23 = F_lookup_type_cache(m, l1, int32(65536))
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
							F_errmsg_internal(m, int32(352490), v9)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(471045), int32(432), int32(480494))
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
												F_errmsg(m, int32(179630), v9+int32(16))
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(471045), int32(451), int32(480494))
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
												F_errmsg(m, int32(179541), v9+int32(32))
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(471045), int32(456), int32(480494))
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
			v23 = F_lookup_type_cache(m, l1, int32(65536))
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
						F_errmsg_internal(m, int32(352490), v9)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(471045), int32(432), int32(480494))
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
											F_errmsg(m, int32(179630), v9+int32(16))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(471045), int32(451), int32(480494))
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
											F_errmsg(m, int32(179541), v9+int32(32))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(471045), int32(456), int32(480494))
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
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
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
						v70 = v34
						m.G0 = v9 + int32(48)
						return v70
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
						if v44 == int32(0) {
							v70 = v34
							m.G0 = v9 + int32(48)
							return v70
						} else {
							F_range_deserialize(m, v33, v17, v9+int32(40), v9+int32(32), v9+int32(15))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_multirange_get_bounds(m, v33, v12, int32(0), v9+int32(24), v9+int32(16))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v66 = F_range_cmp_bounds(m, v33, v9+int32(32), v9+int32(24))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										v70 = int32(base.Ui32(v66) >> (uint(int32(31)) % 32))
										m.G0 = v9 + int32(48)
										return v70
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
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(352490), v9)
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(471045), int32(558), int32(379733))
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
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v34 = int32(0)
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
							v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
							if v41&int32(1) != 0 {
								v70 = v34
								m.G0 = v9 + int32(48)
								return v70
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								if v44 == int32(0) {
									v70 = v34
									m.G0 = v9 + int32(48)
									return v70
								} else {
									F_range_deserialize(m, v33, v17, v9+int32(40), v9+int32(32), v9+int32(15))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										F_multirange_get_bounds(m, v33, v12, int32(0), v9+int32(24), v9+int32(16))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											v66 = F_range_cmp_bounds(m, v33, v9+int32(32), v9+int32(24))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int32(0)
											} else {
												v70 = int32(base.Ui32(v66) >> (uint(int32(31)) % 32))
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
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(352490), v9)
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(471045), int32(558), int32(379733))
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
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						v34 = int32(0)
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
						v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
						if v41&int32(1) != 0 {
							v70 = v34
							m.G0 = v9 + int32(48)
							return v70
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							if v44 == int32(0) {
								v70 = v34
								m.G0 = v9 + int32(48)
								return v70
							} else {
								F_range_deserialize(m, v33, v17, v9+int32(40), v9+int32(32), v9+int32(15))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									F_multirange_get_bounds(m, v33, v12, int32(0), v9+int32(24), v9+int32(16))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v66 = F_range_cmp_bounds(m, v33, v9+int32(32), v9+int32(24))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											v70 = int32(base.Ui32(v66) >> (uint(int32(31)) % 32))
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
func F_multirange_cmp(m *base.Module, l0 int32) int32 {
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	v11 = m.G0
	v13 = v11 - int32(48)
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
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v23 == v24 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L54
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if v27 != 0 {
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
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L51
	}
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v40 < v39 {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v28 == v23 {
		v38 = v27
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v31 = F_lookup_type_cache(m, v23, int32(65536))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+296))
	if v33 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v31
	v38 = v31
	goto L8
L15:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v120 != v16 {
		goto L43
	} else {
		goto L44
	}
L16:
	;
	v42 = v39
	goto L18
L17:
	;
	v42 = v40
	goto L18
L18:
	;
	if int32(0) < v42 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v45 = int32(0)
	if v45 < v40 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v113 = int32(0)
	goto L15
L22:
	;
	v49 = v40
	goto L24
L23:
	;
	v49 = v45
	goto L24
L24:
	;
	v50 = int32(0)
	if v50 < v39 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v53 = v39
	goto L27
L26:
	;
	v53 = v50
	goto L27
L27:
	;
	v56 = v45
	goto L28
L28:
	;
	if v56 == v53 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L21
L30:
	;
	v113 = int32(-1)
	goto L15
L31:
	;
	goto L32
L32:
	;
	if v56 == v49 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v113 = int32(1)
	goto L15
L34:
	;
	goto L35
L35:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
	F_multirange_get_bounds(m, v68, v16, v56, v13+int32(40), v13+int32(32))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
	F_multirange_get_bounds(m, v75, v21, v56, v13+int32(24), v13+int32(16))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
	v87 = F_range_cmp_bounds(m, v82, v13+int32(40), v13+int32(24))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v87 != 0 {
		v113 = v87
		goto L15
	} else {
		goto L39
	}
L39:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
	v94 = F_range_cmp_bounds(m, v89, v13+int32(32), v13+int32(16))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v94 != 0 {
		v113 = v94
		goto L15
	} else {
		goto L41
	}
L41:
	;
	v97 = v56 + int32(1)
	if v97 != v42 {
		v56 = v97
		goto L28
	} else {
		goto L42
	}
L42:
	;
	goto L29
L43:
	;
	F_pfree(m, v16)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v124 != v21 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	F_pfree(m, v21)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	m.G0 = v13 + int32(48)
	return v113
L50:
	;
	goto L49
L51:
	;
	F_errmsg_internal(m, int32(309174), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(471045), int32(2589), int32(223961))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v23
	F_errmsg_internal(m, int32(352490), v13)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(471045), int32(558), int32(379733))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = int32(1)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v15 == v4 {
		v113 = v14
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v113
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
	v113 = int32(0)
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
		v113 = v14
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v38 = v4
	v39 = v4
	goto L9
L9:
	;
	F_multirange_get_bounds(m, l0, l2, v38, v12+int32(8), v12)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v113 = v105
	goto L1
L11:
	;
	v50 = F_range_cmp_bounds(m, l0, v12+int32(16), v12+int32(8))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	if v50 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v60 = v39
	goto L16
L14:
	;
	v87 = v39
	goto L15
L15:
	;
	v90 = int32(0)
	v95 = F_range_cmp_bounds(m, l0, v12+int32(24), v12+int32(8))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L24
	}
L16:
	;
	v64 = v60 + int32(1)
	if v18 <= v64 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v87 = v64
	goto L15
L18:
	;
	v113 = int32(0)
	goto L1
L19:
	;
	goto L20
L20:
	;
	F_multirange_get_bounds(m, l0, l1, v64, v12+int32(24), v12+int32(16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v77 = F_range_cmp_bounds(m, l0, v12+int32(16), v12+int32(8))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	if v77 < int32(0) {
		v60 = v64
		goto L16
	} else {
		goto L23
	}
L23:
	;
	goto L17
L24:
	;
	if int32(0) < v95 {
		v113 = v90
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v101 = F_range_cmp_bounds(m, l0, v12+int32(16), v12)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	if v101 < int32(0) {
		v113 = v90
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v105 = int32(1)
	v107 = v38 + v105
	if v107 != v15 {
		v38 = v107
		v39 = v87
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	v4 = int32(0)
	v11 = l1 + int32(8)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
	if l2 <= v4 {
		v41 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v49 = v47 << (uint(int32(2)) % 32)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+(v11+v49)))))
	v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+8)))
	if v13 == int32(105) {
		v79 = (v47*int32(5) + int32(11)) & int32(-4)
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v19 = v4
	v22 = l2
	goto L3
L3:
	;
	v25 = int32(2)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11+v22<<(uint(v25)%32))))
	v31 = v28&int32(2147483647) + v19
	if base.Ui32(v22) < base.Ui32(v25) {
		v41 = v31
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v41 = v31
	goto L1
L5:
	;
	if int32(0) <= v28 {
		v19 = v31
		v22 = v22 - int32(1)
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	v316 = v315 - v81
	v318 = v316 + int32(9)
	v319 = F_palloc0(m, v318)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L97
	} else {
		goto L98
	}
L8:
	;
	if v211&int32(3) == int32(0) {
		v274 = v211
		goto L82
	} else {
		goto L83
	}
L9:
	;
	v222 = v219 & int32(255)
	if v222 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L10:
	;
	if v13 == int32(105) {
		goto L60
	} else {
		goto L61
	}
L11:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v189 != 0 {
		v219 = v189
		v220 = v188
		goto L9
	} else {
		goto L58
	}
L12:
	;
	v185 = v184 + v81
	if v52&int32(80) != 0 {
		v315 = v185
		goto L7
	} else {
		goto L57
	}
L13:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v184 = int32(base.Ui32(v180) >> (uint(int32(2)) % 32))
	goto L12
L14:
	;
	if v52&int32(81) != 0 {
		v315 = v174
		goto L7
	} else {
		goto L55
	}
L15:
	;
	v81 = v79 + l1 + v41
	if v52&int32(41) != 0 {
		v174 = v81
		goto L14
	} else {
		goto L20
	}
L16:
	;
	switch v13 - int32(99) {
	case 0:
		goto L19
	case 1:
		goto L18
	default:
		goto L17
	}
L17:
	;
	v79 = (v47*int32(5) + int32(9)) & int32(-2)
	goto L15
L18:
	;
	v79 = (v47*int32(5) + int32(15)) & int32(-8)
	goto L15
L19:
	;
	v79 = v49 + v47 + int32(8)
	goto L15
L20:
	;
	if int32(0) < v53 {
		v174 = v53 + v81
		goto L14
	} else {
		goto L21
	}
L21:
	;
	if v53 == int32(-1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v89 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	if v81&int32(3) == int32(0) {
		v137 = v81
		goto L40
	} else {
		goto L41
	}
L25:
	;
	v92 = int32(6)
	v94 = int32(18)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	if v96 == v94 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if v89&int32(1) == int32(0) {
		goto L13
	} else {
		goto L37
	}
L28:
	;
	v99 = v94
	goto L30
L29:
	;
	v99 = int32(2)
	goto L30
L30:
	;
	if v96&int32(254) == int32(2) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v104 = v92
	goto L33
L32:
	;
	v104 = v99
	goto L33
L33:
	;
	if v96 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v107 = v92
	goto L36
L35:
	;
	v107 = v104
	goto L36
L36:
	;
	v184 = v107
	goto L12
L37:
	;
	v184 = int32(base.Ui32(v89) >> (uint(int32(1)) % 32))
	goto L12
L38:
	;
	v174 = v170 + v81 + int32(1)
	goto L14
L39:
	;
	v170 = v162 - v81
	goto L38
L40:
	;
	v141 = v137
	goto L49
L41:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v121 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v170 = int32(0)
	goto L38
L43:
	;
	goto L44
L44:
	;
	v126 = v81
	goto L45
L45:
	;
	v130 = v126 + int32(1)
	if v130&int32(3) == int32(0) {
		v137 = v130
		goto L40
	} else {
		goto L47
	}
L46:
	;
	v162 = v130
	goto L39
L47:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v135 != 0 {
		v126 = v130
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v150 = int32(-2139062144)
	if (int32(16843008)-v147|v147)&v150 == v150 {
		v141 = v141 + int32(4)
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v156 = v141
	goto L52
L51:
	;
	goto L50
L52:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v160 != 0 {
		v156 = v156 + int32(1)
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v162 = v156
	goto L39
L54:
	;
	goto L53
L55:
	;
	if v53 == int32(-1) {
		v188 = v174
		goto L11
	} else {
		goto L56
	}
L56:
	;
	v192 = v53
	v193 = v174
	v194 = int32(0)
	goto L10
L57:
	;
	v188 = v185
	goto L11
L58:
	;
	v192 = int32(-1)
	v193 = v188
	v194 = int32(1)
	goto L10
L59:
	;
	if int32(0) < v53 {
		v315 = v192 + v211
		goto L7
	} else {
		goto L65
	}
L60:
	;
	v211 = (v193 + int32(3)) & int32(-4)
	goto L59
L61:
	;
	goto L62
L62:
	;
	switch v13 - int32(99) {
	case 0:
		v211 = v193
		goto L59
	case 1:
		goto L64
	default:
		goto L63
	}
L63:
	;
	v211 = (v193 + int32(1)) & int32(-2)
	goto L59
L64:
	;
	v211 = (v193 + int32(7)) & int32(-8)
	goto L59
L65:
	;
	if v194 == int32(0) {
		goto L8
	} else {
		goto L66
	}
L66:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	v219 = v217
	v220 = v211
	goto L9
L67:
	;
	v225 = int32(6)
	v227 = int32(18)
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+1)))
	if v229 == v227 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	v242 = int32(1)
	if v222&v242 != 0 {
		v315 = v220 + int32(base.Ui32(v222)>>(uint(v242)%32))
		goto L7
	} else {
		goto L79
	}
L70:
	;
	v232 = v227
	goto L72
L71:
	;
	v232 = int32(2)
	goto L72
L72:
	;
	if v229&int32(254) == int32(2) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v237 = v225
	goto L75
L74:
	;
	v237 = v232
	goto L75
L75:
	;
	if v229 == int32(1) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v240 = v225
	goto L78
L77:
	;
	v240 = v237
	goto L78
L78:
	;
	v315 = v220 + v240
	goto L7
L79:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v315 = v220 + int32(base.Ui32(v247)>>(uint(int32(2))%32))
	goto L7
L80:
	;
	v315 = v307 + v211 + int32(1)
	goto L7
L81:
	;
	v307 = v299 - v211
	goto L80
L82:
	;
	v278 = v274
	goto L91
L83:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	if v258 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v307 = int32(0)
	goto L80
L85:
	;
	goto L86
L86:
	;
	v263 = v211
	goto L87
L87:
	;
	v267 = v263 + int32(1)
	if v267&int32(3) == int32(0) {
		v274 = v267
		goto L82
	} else {
		goto L89
	}
L88:
	;
	v299 = v267
	goto L81
L89:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
	if v272 != 0 {
		v263 = v267
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	v287 = int32(-2139062144)
	if (int32(16843008)-v284|v284)&v287 == v287 {
		v278 = v278 + int32(4)
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v293 = v278
	goto L94
L93:
	;
	goto L92
L94:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	if v297 != 0 {
		v293 = v293 + int32(1)
		goto L94
	} else {
		goto L96
	}
L95:
	;
	v299 = v293
	goto L81
L96:
	;
	goto L95
L97:
	;
	return int32(0)
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v319))) = v318 << (uint(int32(2)) % 32)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v319)+4)) = v326
	v329 = v319 + int32(8)
	if v316 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v331+v316))) = uint8(v52)
	return v319
L100:
	;
	v330 = F__emscripten_memcpy_bulkmem(m, v329, v81, v316)
	mBase = m.M
	v331 = v330
	goto L102
L101:
	;
	v331 = v329
	goto L102
L102:
	;
	goto L99
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
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(352490), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(471045), int32(558), int32(379733))
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
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(352490), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(471045), int32(558), int32(379733))
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
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
				v73 = v2
				m.G0 = v9 + int32(48)
				return v73
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				v28 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17+int32(base.Ui32(v22)>>(uint(int32(2))%32))-int32(1)))))
				if v28&int32(1) != 0 {
					v73 = v2
					m.G0 = v9 + int32(48)
					return v73
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
					if v33 != 0 {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						if v34 == v31 {
							v44 = v33
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
							F_multirange_get_bounds(m, v45, v12, int32(0), v9+int32(40), v9+int32(32))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
								F_range_deserialize(m, v53, v17, v9+int32(24), v9+int32(16), v9+int32(15))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
									v67 = F_range_cmp_bounds(m, v62, v9+int32(40), v9+int32(24))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										v73 = int32(base.Ui32(v67^int32(-1)) >> (uint(int32(31)) % 32))
										m.G0 = v9 + int32(48)
										return v73
									}
								}
							}
						} else {
							v37 = F_lookup_type_cache(m, v31, int32(65536))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+296))
								if v39 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v31
										F_errmsg_internal(m, int32(352490), v9)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(471045), int32(558), int32(379733))
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
									v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v37
									v44 = v37
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
									F_multirange_get_bounds(m, v45, v12, int32(0), v9+int32(40), v9+int32(32))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
										F_range_deserialize(m, v53, v17, v9+int32(24), v9+int32(16), v9+int32(15))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											v62 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
											v67 = F_range_cmp_bounds(m, v62, v9+int32(40), v9+int32(24))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return int32(0)
											} else {
												v73 = int32(base.Ui32(v67^int32(-1)) >> (uint(int32(31)) % 32))
												m.G0 = v9 + int32(48)
												return v73
											}
										}
									}
								}
							}
						}
					} else {
						v37 = F_lookup_type_cache(m, v31, int32(65536))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+296))
							if v39 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v31
									F_errmsg_internal(m, int32(352490), v9)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(471045), int32(558), int32(379733))
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
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v37
								v44 = v37
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
								F_multirange_get_bounds(m, v45, v12, int32(0), v9+int32(40), v9+int32(32))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
									F_range_deserialize(m, v53, v17, v9+int32(24), v9+int32(16), v9+int32(15))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
										v67 = F_range_cmp_bounds(m, v62, v9+int32(40), v9+int32(24))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return int32(0)
										} else {
											v73 = int32(base.Ui32(v67^int32(-1)) >> (uint(int32(31)) % 32))
											m.G0 = v9 + int32(48)
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
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
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
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
	v29 = int32(24)
	v31 = int32(65280)
	v33 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v26+v27))) = v22<<(uint(v29)%32) | v22&v31<<(uint(v33)%32) | (int32(base.Ui32(v22)>>(uint(v33)%32))&v31 | int32(base.Ui32(v22)>>(uint(v29)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v26 + int32(4)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if int32(0) < v48 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+296))
	v56 = F_palloc(m, v48<<(uint(int32(2))%32))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v144 << (uint(int32(2)) % 32)
	goto L21
L10:
	;
	v58 = int32(0)
	goto L11
L11:
	;
	v69 = F_multirange_get_range(m, v52, v10, v58)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v81 = int32(0)
	goto L15
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56+v58<<(uint(int32(2))%32)))) = v69
	v73 = v58 + int32(1)
	if v73 != v48 {
		v58 = v73
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v56+v81<<(uint(int32(2))%32))))
	v90 = F_SendFunctionCall(m, v18+int32(4), v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L9
L17:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	F_enlargeStringInfo(m, v15, int32(4))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v99 = int32(2)
	v101 = int32(4)
	v102 = int32(base.Ui32(v92)>>(uint(v99)%32)) - v101
	v103 = int32(24)
	v105 = int32(65280)
	v107 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v96+v97))) = v102<<(uint(v103)%32) | v102&v105<<(uint(v107)%32) | (int32(base.Ui32(v102)>>(uint(v107)%32))&v105 | int32(base.Ui32(v102)>>(uint(v103)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v96 + v101
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	F_pq_sendbytes(m, v15, v90+v101, int32(base.Ui32(v124)>>(uint(v99)%32))-v101)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v132 = v81 + int32(1)
	if v132 != v48 {
		v81 = v132
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	return v143
}
