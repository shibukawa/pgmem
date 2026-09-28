package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_UnregisterSnapshot(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	if l0 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, _c_F_UnregisterSnapshot[0]))
		F_ResourceOwnerForget(m, v3, base.I64_extend_i32_u(l0), int32(_a_F_UnregisterSnapshot_0))
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_UnregisterSnapshotNoOwner(m, l0)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		return
	}
}
func F_UnregisterSnapshotNoOwner(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v6 = v4 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v6
	if v6 != 0 {
		return
	} else {
		F_pairingheap_remove(m, int32(_a_F_UnregisterSnapshotNoOwner_0), l0+int32(52))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if v13 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				if v14 != 0 {
					return
				} else {
					F_pfree(m, l0)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, _c_F_UnregisterSnapshotNoOwner[0]))
						if v18 != 0 {
							return
						} else {
							v20 = *(*int32)(unsafe.Add(mBase, _c_F_UnregisterSnapshotNoOwner[1]))
							if v20 == int32(0) {
								v24 = int32(0)
								*(*int32)(unsafe.Add(mBase, _c_F_UnregisterSnapshotNoOwner[2])) = v24
								v27 = *(*int32)(unsafe.Add(mBase, _c_F_UnregisterSnapshotNoOwner[3]))
								*(*int32)(unsafe.Add(mBase, uint32(v27)+52)) = v24
								return
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, _c_F_UnregisterSnapshotNoOwner[3]))
								v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+52))
								v33 = int32(3)
								v36 = *(*int32)(unsafe.Add(mBase, _c_F_UnregisterSnapshotNoOwner[1]))
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v36-int32(48))))
								if base.B2i32(base.Ui32(v32) < base.Ui32(v33))|base.B2i32(base.Ui32(v39) < base.Ui32(v33)) == int32(0) {
									if v32-v39 < int32(0) {
										*(*int32)(unsafe.Add(mBase, _c_F_UnregisterSnapshotNoOwner[2])) = v39
										*(*int32)(unsafe.Add(mBase, uint32(v31)+52)) = v39
									} else {
									}
								} else {
									if base.Ui32(v39) <= base.Ui32(v32) {
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_UnregisterSnapshotNoOwner[2])) = v39
										*(*int32)(unsafe.Add(mBase, uint32(v31)+52)) = v39
									}
								}
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_uint32_hash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = int32(711645284)
	v10 = v4 - int32(1636608428) ^ v7 - int32(1455628627)
	v15 = v10 ^ int32(-1636608428) - base.I32_rotl(v10, int32(25))
	v20 = v15 ^ v7 - base.I32_rotl(v15, int32(16))
	v24 = v20 ^ v10 - base.I32_rotl(v20, int32(4))
	v28 = v24 ^ v15 - base.I32_rotl(v24, int32(14))
	return v28 ^ v20 - base.I32_rotl(v28, int32(24))
}
func F_unaccent_dict(m *base.Module, l0 int32) int64 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(1)
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v11 == v10 {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
		v19 = F_get_func_namespace(m, v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			v24 = int64(0)
			v26 = F_GetSysCacheOid(m, int32(75), int64(107119), base.I64_extend_i32_u(v19), v24, v24)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int64(0)
			} else {
				if v26 != 0 {
					v49 = int32(0)
					v50 = v26
					v54 = l0 + v49<<(uint(int32(4))%32)
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+24))
					v56 = F_pg_detoast_datum_packed(m, v55)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int64(0)
					} else {
						v58 = F_lookup_ts_dictionary_cache(m, v50)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int64(0)
						} else {
							v61 = v54 + int32(24)
							v65 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58)+44)))
							v66 = int32(1)
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
							v70 = v68 & v66
							if v70 != 0 {
								v71 = v66
							} else {
								v71 = int32(4)
							}
							if v68 == int32(1) {
								v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
								if v79 == int32(18) {
									v82 = int32(16)
								} else {
									v82 = int32(0)
								}
								if base.Ui32((v79-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v89 = int32(4)
								} else {
									v89 = v82
								}
								v100 = v89
							} else {
								v90 = int32(1)
								if v70 != 0 {
									v100 = int32(base.Ui32(v68)>>(uint(v90)%32)) - v90
								} else {
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
									v100 = int32(base.Ui32(v94)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							v103 = F_FunctionCall4Coll(m, v58+int32(12), int32(0), v65, base.I64_extend_i32_u(v56+v71), base.I64_extend_i32_s(v100), int64(0))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int64(0)
							} else {
								v105 = base.I32_wrap_i64(v103)
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
								if v106 != v56 {
									F_pfree(m, v56)
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int64(0)
									} else {
										if v105 == int32(0) {
											v112 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
											v113 = F_pg_detoast_datum_copy(m, v112)
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return int64(0)
											} else {
												v130 = v113
												m.G0 = v8 + int32(16)
												return base.I64_extend_i32_u(v130)
											}
										} else {
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
											if v115 == int32(0) {
												F_pfree(m, v105)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return int64(0)
												} else {
													v120 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
													v121 = F_pg_detoast_datum_copy(m, v120)
													mBase = m.M
													v122 = m.ExcPending
													if v122 != 0 {
														return int64(0)
													} else {
														v130 = v121
														m.G0 = v8 + int32(16)
														return base.I64_extend_i32_u(v130)
													}
												}
											} else {
												v123 = F_cstring_to_text(m, v115)
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return int64(0)
												} else {
													v125 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
													F_pfree(m, v125)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int64(0)
													} else {
														F_pfree(m, v105)
														mBase = m.M
														v129 = m.ExcPending
														if v129 != 0 {
															return int64(0)
														} else {
															v130 = v123
															m.G0 = v8 + int32(16)
															return base.I64_extend_i32_u(v130)
														}
													}
												}
											}
										}
									}
								} else {
									if v105 == int32(0) {
										v112 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
										v113 = F_pg_detoast_datum_copy(m, v112)
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int64(0)
										} else {
											v130 = v113
											m.G0 = v8 + int32(16)
											return base.I64_extend_i32_u(v130)
										}
									} else {
										v115 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
										if v115 == int32(0) {
											F_pfree(m, v105)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int64(0)
											} else {
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
												v121 = F_pg_detoast_datum_copy(m, v120)
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return int64(0)
												} else {
													v130 = v121
													m.G0 = v8 + int32(16)
													return base.I64_extend_i32_u(v130)
												}
											}
										} else {
											v123 = F_cstring_to_text(m, v115)
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
												return int64(0)
											} else {
												v125 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
												F_pfree(m, v125)
												mBase = m.M
												v127 = m.ExcPending
												if v127 != 0 {
													return int64(0)
												} else {
													F_pfree(m, v105)
													mBase = m.M
													v129 = m.ExcPending
													if v129 != 0 {
														return int64(0)
													} else {
														v130 = v123
														m.G0 = v8 + int32(16)
														return base.I64_extend_i32_u(v130)
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
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(67137668))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							v35 = F_get_namespace_name(m, v19)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(_a_F_unaccent_dict_0)
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v35
								F_errmsg(m, int32(_a_F_unaccent_dict_1), v8)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_unaccent_dict_2), int32(464), int32(_a_F_unaccent_dict_3))
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
					}
				}
			}
		}
	} else {
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v49 = v10
		v50 = v48
		v54 = l0 + v49<<(uint(int32(4))%32)
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+24))
		v56 = F_pg_detoast_datum_packed(m, v55)
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int64(0)
		} else {
			v58 = F_lookup_ts_dictionary_cache(m, v50)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int64(0)
			} else {
				v61 = v54 + int32(24)
				v65 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v58)+44)))
				v66 = int32(1)
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
				v70 = v68 & v66
				if v70 != 0 {
					v71 = v66
				} else {
					v71 = int32(4)
				}
				if v68 == int32(1) {
					v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
					if v79 == int32(18) {
						v82 = int32(16)
					} else {
						v82 = int32(0)
					}
					if base.Ui32((v79-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v89 = int32(4)
					} else {
						v89 = v82
					}
					v100 = v89
				} else {
					v90 = int32(1)
					if v70 != 0 {
						v100 = int32(base.Ui32(v68)>>(uint(v90)%32)) - v90
					} else {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
						v100 = int32(base.Ui32(v94)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v103 = F_FunctionCall4Coll(m, v58+int32(12), int32(0), v65, base.I64_extend_i32_u(v56+v71), base.I64_extend_i32_s(v100), int64(0))
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return int64(0)
				} else {
					v105 = base.I32_wrap_i64(v103)
					v106 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
					if v106 != v56 {
						F_pfree(m, v56)
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return int64(0)
						} else {
							if v105 == int32(0) {
								v112 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
								v113 = F_pg_detoast_datum_copy(m, v112)
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int64(0)
								} else {
									v130 = v113
									m.G0 = v8 + int32(16)
									return base.I64_extend_i32_u(v130)
								}
							} else {
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
								if v115 == int32(0) {
									F_pfree(m, v105)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int64(0)
									} else {
										v120 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
										v121 = F_pg_detoast_datum_copy(m, v120)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int64(0)
										} else {
											v130 = v121
											m.G0 = v8 + int32(16)
											return base.I64_extend_i32_u(v130)
										}
									}
								} else {
									v123 = F_cstring_to_text(m, v115)
									mBase = m.M
									v124 = m.ExcPending
									if v124 != 0 {
										return int64(0)
									} else {
										v125 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
										F_pfree(m, v125)
										mBase = m.M
										v127 = m.ExcPending
										if v127 != 0 {
											return int64(0)
										} else {
											F_pfree(m, v105)
											mBase = m.M
											v129 = m.ExcPending
											if v129 != 0 {
												return int64(0)
											} else {
												v130 = v123
												m.G0 = v8 + int32(16)
												return base.I64_extend_i32_u(v130)
											}
										}
									}
								}
							}
						}
					} else {
						if v105 == int32(0) {
							v112 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
							v113 = F_pg_detoast_datum_copy(m, v112)
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return int64(0)
							} else {
								v130 = v113
								m.G0 = v8 + int32(16)
								return base.I64_extend_i32_u(v130)
							}
						} else {
							v115 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
							if v115 == int32(0) {
								F_pfree(m, v105)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int64(0)
								} else {
									v120 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
									v121 = F_pg_detoast_datum_copy(m, v120)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int64(0)
									} else {
										v130 = v121
										m.G0 = v8 + int32(16)
										return base.I64_extend_i32_u(v130)
									}
								}
							} else {
								v123 = F_cstring_to_text(m, v115)
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
									return int64(0)
								} else {
									v125 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
									F_pfree(m, v125)
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
										return int64(0)
									} else {
										F_pfree(m, v105)
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int64(0)
										} else {
											v130 = v123
											m.G0 = v8 + int32(16)
											return base.I64_extend_i32_u(v130)
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
func F_unaccent_init(m *base.Module, l0 int32) int64 {
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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v9 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L17
	} else {
		goto L29
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L17
	} else {
		goto L25
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v12 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v16 = int32(0)
	goto L6
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L17
	} else {
		goto L21
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v16<<(uint(int32(2))%32))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v26 = int32(_a_F_unaccent_init_0)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_unaccent_init[0])))
	if base.B2i32(v29 == int32(0))|base.B2i32(v29 != v32) != 0 {
		v50 = v29
		v51 = v32
		goto L9
	} else {
		goto L10
	}
L7:
	;
	m.G0 = v7 + int32(16)
	return base.I64_extend_i32_u(v59)
L8:
	;
	if v50-v51 != 0 {
		goto L5
	} else {
		goto L15
	}
L9:
	;
	goto L8
L10:
	;
	v35 = v25
	v36 = v26
	goto L11
L11:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	if v40 == int32(0) {
		v50 = v40
		v51 = v39
		goto L9
	} else {
		goto L13
	}
L12:
	;
	v50 = v40
	v51 = v39
	goto L9
L13:
	;
	v43 = int32(1)
	if v40 == v39 {
		v35 = v35 + v43
		v36 = v36 + v43
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	if v16&int32(1) != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v55 = F_defGetString(m, v24)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int64(0)
L18:
	;
	v59 = F_initTrie(m, v55)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v61 = int32(1)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v61 < v62 {
		v16 = v61
		goto L6
	} else {
		goto L20
	}
L20:
	;
	goto L7
L21:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v77
	F_errmsg(m, int32(_a_F_unaccent_init_1), v7)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_unaccent_init_2), int32(363), int32(_a_F_unaccent_init_3))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(_a_F_unaccent_init_4), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L17
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_unaccent_init_2), int32(371), int32(_a_F_unaccent_init_3))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L17
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L17
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(_a_F_unaccent_init_5), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L17
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_unaccent_init_2), int32(354), int32(_a_F_unaccent_init_3))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L17
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_unlink_if_exists_fname(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
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
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 != 0 {
		v9 = F_rmdir(m, l0)
		mBase = m.M
		if v9 == int32(0) {
			m.G0 = v7 + int32(16)
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_unlink_if_exists_fname[0]))
			if v13 == int32(44) {
				m.G0 = v7 + int32(16)
				return
			} else {
				v17 = F_errstart(m, l2, int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					if v17 == int32(0) {
						m.G0 = v7 + int32(16)
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
							F_errmsg(m, int32(_a_F_unlink_if_exists_fname_0), v7)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_unlink_if_exists_fname_1), int32(3834), int32(_a_F_unlink_if_exists_fname_2))
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v33 = F_PathNameDeleteTemporaryFile(m, l0, int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_unlink_initfile(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_unlink(m, l0)
	mBase = m.M
	if int32(0) <= v8 {
		m.G0 = v6 + int32(16)
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_unlink_initfile[0]))
		if v12 == int32(44) {
			m.G0 = v6 + int32(16)
			return
		} else {
			v16 = F_errstart(m, l1, int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				if v16 == int32(0) {
					m.G0 = v6 + int32(16)
					return
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
						F_errmsg(m, int32(_a_F_unlink_initfile_0), v6)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_unlink_initfile_1), int32(_a_F_unlink_initfile_2), int32(_a_F_unlink_initfile_3))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								m.G0 = v6 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_untransformRelOptions(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = base.I32_wrap_i64(l0)
	if v12 == v2 {
		v75 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v75
L2:
	;
	v15 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	F_deconstruct_array_builtin(m, v15, int32(25), v10+int32(12), int32(0), v10+int32(8))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v27 <= int32(0) {
		v75 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v33 = int32(0)
	v36 = v2
	goto L7
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v33<<(uint(int32(3))%32))))
	v44 = F_text_to_cstring(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v75 = v64
	goto L1
L9:
	;
	v46 = int32(61)
	v47 = F___strchrnul(m, v44, v46)
	mBase = m.M
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v49 == v46 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v53 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v53 = v47
	goto L13
L12:
	;
	v53 = int32(0)
	goto L13
L13:
	;
	goto L10
L14:
	;
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v54)
	v58 = F_makeString(m, v53+int32(1))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L3
	} else {
		goto L17
	}
L15:
	;
	v60 = int32(0)
	goto L16
L16:
	;
	v62 = F_makeDefElem(m, v44, v60, int32(-1))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L3
	} else {
		goto L18
	}
L17:
	;
	v60 = v58
	goto L16
L18:
	;
	v64 = F_lappend(m, v36, v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v67 = v33 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v67 < v68 {
		v33 = v67
		v36 = v64
		goto L7
	} else {
		goto L20
	}
L20:
	;
	goto L8
}
func F_updateInitAclDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	var v11 int32
	_ = v11
	F_updateAclDependenciesWorker(m, l0, l1, l2, int32(0), int32(105), l3, l4, l5, l6)
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		return
	}
}
func F_update_attstats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v74 int64
	_ = v74
	var v83 int64
	_ = v83
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v120 int64
	_ = v120
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int64
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int64
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int64
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v240 int64
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v273 int32
	_ = v273
	var v286 int32
	_ = v286
	var v298 int32
	_ = v298
	var v302 int64
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v365 int64
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v440 int64
	_ = v440
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v449 int64
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v490 int32
	_ = v490
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	v5 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(320)
	m.G0 = v26
	if v5 < l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = base.I64_extend_i32_u(l1)
	v31 = base.I64_extend_i32_u(l0)
	v34 = F_table_open(m, int32(2619), int32(3))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v26 + int32(320)
	return
L4:
	;
	return
L5:
	;
	v45 = v5
	v53 = v5
	goto L6
L6:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l3+v53<<(uint(int32(2))%32))))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+36)))
	if v63 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v490 != 0 {
		goto L56
	} else {
		goto L57
	}
L8:
	;
	v66 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+55)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v26)+40)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v66
	v74 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = v74
	*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = v74
	*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = v74
	*(*int64)(unsafe.Add(mBase, uint32(v26)+23)) = v74
	*(*int64)(unsafe.Add(mBase, uint32(v26)+64)) = v31
	v83 = int64(*(*int16)(unsafe.Add(mBase, uint32(v62)+224)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+80)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v26)+72)) = v83
	v86 = int64(*(*int32)(unsafe.Add(mBase, uint32(v62)+40)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+88)) = v86
	v88 = int64(*(*int32)(unsafe.Add(mBase, uint32(v62)+44)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+96)) = v88
	v90 = int64(*(*int32)(unsafe.Add(mBase, uint32(v62)+48)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+104)) = v90
	v92 = int64(*(*int16)(unsafe.Add(mBase, uint32(v62)+52)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+112)) = v92
	v94 = int64(*(*int16)(unsafe.Add(mBase, uint32(v62)+54)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+120)) = v94
	v96 = int64(*(*int16)(unsafe.Add(mBase, uint32(v62)+56)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+128)) = v96
	v98 = int64(*(*int16)(unsafe.Add(mBase, uint32(v62)+58)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+136)) = v98
	v100 = int64(*(*int16)(unsafe.Add(mBase, uint32(v62)+60)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+144)) = v100
	v102 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v62)+64)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+152)) = v102
	v104 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v62)+68)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+160)) = v104
	v106 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v62)+72)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+168)) = v106
	v108 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v62)+76)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+176)) = v108
	v110 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v62)+80)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+184)) = v110
	v112 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v62)+84)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+192)) = v112
	v114 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v62)+88)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+200)) = v114
	v116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v62)+92)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+208)) = v116
	v118 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v62)+96)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+216)) = v118
	v120 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v62)+100)))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+224)) = v120
	v136 = int32(21)
	v139 = int32(0)
	goto L11
L9:
	;
	v490 = v45
	goto L10
L10:
	;
	v505 = v53 + int32(1)
	if v505 != l2 {
		v45 = v490
		v53 = v505
		goto L6
	} else {
		goto L55
	}
L11:
	;
	v157 = v139 << (uint(int32(2)) % 32)
	v158 = v62 + int32(124) + v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v159 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v387 = int32(26)
	v388 = int32(0)
	goto L32
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26-int32(-64)+v136<<(uint(int32(3))%32)))) = v365
	v367 = int32(1)
	v370 = v139 + v367
	if v370 != int32(5) {
		v136 = v136 + v367
		v139 = v370
		goto L11
	} else {
		goto L31
	}
L14:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157+(v62+int32(104)))))
	v164 = F_palloc(m, v161<<(uint(int32(3))%32))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v339 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26+int32(32)+v136))) = uint8(v339)
	v365 = int64(0)
	goto L13
L17:
	;
	if v161 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v333 = F_construct_array_builtin(m, v164, v161, int32(700))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L4
	} else {
		goto L30
	}
L19:
	;
	v169 = v161 & int32(3)
	v170 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v161) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v178 = v170
	v192 = int32(0)
	goto L23
L21:
	;
	v250 = v170
	goto L22
L22:
	;
	v273 = v250
	v286 = v170
	goto L27
L23:
	;
	v200 = int32(3)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v204 = int32(2)
	v207 = int64(*(*int32)(unsafe.Add(mBase, uint32(v203+v178<<(uint(v204)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v164+v178<<(uint(v200)%32)))) = v207
	v210 = v178 | int32(1)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v218 = int64(*(*int32)(unsafe.Add(mBase, uint32(v214+v210<<(uint(v204)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v164+v210<<(uint(v200)%32)))) = v218
	v221 = v178 | v204
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v229 = int64(*(*int32)(unsafe.Add(mBase, uint32(v225+v221<<(uint(v204)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v164+v221<<(uint(v200)%32)))) = v229
	v232 = v178 | v200
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v240 = int64(*(*int32)(unsafe.Add(mBase, uint32(v236+v232<<(uint(v204)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v164+v232<<(uint(v200)%32)))) = v240
	v242 = int32(4)
	v243 = v178 + v242
	v245 = v192 + v242
	if v245 != v161&int32(2147483644) {
		v178 = v243
		v192 = v245
		goto L23
	} else {
		goto L25
	}
L24:
	;
	if v169 == int32(0) {
		goto L18
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	v250 = v243
	goto L22
L27:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v302 = int64(*(*int32)(unsafe.Add(mBase, uint32(v298+v273<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v164+v273<<(uint(int32(3))%32)))) = v302
	v304 = int32(1)
	v307 = v286 + v304
	if v307 != v169 {
		v273 = v273 + v304
		v286 = v307
		goto L27
	} else {
		goto L29
	}
L28:
	;
	goto L18
L29:
	;
	goto L28
L30:
	;
	v365 = base.I64_extend_i32_u(v333)
	goto L13
L31:
	;
	goto L12
L32:
	;
	v416 = v388 << (uint(int32(2)) % 32)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v62+int32(164)+v416)))
	if v418 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v449 = int64(*(*int16)(unsafe.Add(mBase, uint32(v62)+224)))
	v450 = F_SearchSysCache3(m, int32(65), v31, v449, v30)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L4
	} else {
		goto L40
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26-int32(-64)+v387<<(uint(int32(3))%32)))) = v440
	v442 = int32(1)
	v445 = v388 + v442
	if v445 != int32(5) {
		v387 = v387 + v442
		v388 = v445
		goto L32
	} else {
		goto L39
	}
L35:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v416+(v62+int32(144)))))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v416+(v62+int32(184)))))
	v426 = int32(*(*int16)(unsafe.Add(mBase, uint32(v62+int32(204)+v388<<(uint(int32(1))%32)))))
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388+(v62+int32(214))))))
	v430 = int32(*(*int8)(unsafe.Add(mBase, uint32(v388+(v62+int32(219))))))
	v431 = F_construct_array(m, v418, v420, v422, v426, v428, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v437 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26+int32(32)+v387))) = uint8(v437)
	v440 = int64(0)
	goto L34
L38:
	;
	v440 = base.I64_extend_i32_u(v431)
	goto L34
L39:
	;
	goto L33
L40:
	;
	if v45 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v454 = F_CatalogOpenIndexes(m, v34)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	v456 = v45
	goto L43
L43:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	if v450 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v456 = v454
	goto L43
L45:
	;
	F_pfree(m, v478)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L4
	} else {
		goto L54
	}
L46:
	;
	v462 = F_heap_modify_tuple(m, v450, v457, v26-int32(-64), v26+int32(32), v26)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v474 = F_heap_form_tuple(m, v457, v26-int32(-64), v26+int32(32))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L4
	} else {
		goto L52
	}
L49:
	;
	F_ReleaseCatCache(m, v450)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	F_CatalogTupleUpdateWithInfo(m, v34, v462+int32(4), v462, v456)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v478 = v462
	goto L45
L52:
	;
	F_CatalogTupleInsertWithInfo(m, v34, v474, v456)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v478 = v474
	goto L45
L54:
	;
	v490 = v456
	goto L10
L55:
	;
	goto L7
L56:
	;
	F_CatalogCloseIndexes(m, v490)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L4
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	F_relation_close(m, v34, int32(3))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L4
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	goto L3
}
func F_update_progress_txn_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(_a_F_update_progress_txn_cb_wrapper_0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_update_progress_txn_cb_wrapper[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, _c_F_update_progress_txn_cb_wrapper[0])) = v9 + int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_update_progress_txn_cb_wrapper_1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(1058)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v9 + int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+147)) = uint8(v4)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+164)) = uint8(v4)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+152)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+160)) = v30
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v13)+128))
	if v36 != 0 {
		m.T0[v36].(func(*base.Module, int32, int64, int32, int32))(m, v13, l2, v30, int32(0))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v41 = v40
			*(*int32)(unsafe.Add(mBase, _c_F_update_progress_txn_cb_wrapper[0])) = v41
			m.G0 = v9 + int32(32)
			return
		}
	} else {
		v41 = v12
		*(*int32)(unsafe.Add(mBase, _c_F_update_progress_txn_cb_wrapper[0])) = v41
		m.G0 = v9 + int32(32)
		return
	}
}
