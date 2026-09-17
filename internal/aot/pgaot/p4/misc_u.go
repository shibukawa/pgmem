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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	if l0 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, _c_F_UnregisterSnapshot[0]))
		F_ResourceOwnerForget(m, v3, l0, int32(_a_F_UnregisterSnapshot_0))
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			F_UnregisterSnapshotNoOwner(m, l0)
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
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
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v5 = v3 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v5
	if v5 != 0 {
		return
	} else {
		F_pairingheap_remove(m, int32(_a_F_UnregisterSnapshotNoOwner_0), l0+int32(52))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if v12 != 0 {
				return
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				if v13 != 0 {
					return
				} else {
					F_pfree(m, l0)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, _c_F_UnregisterSnapshotNoOwner[0]))
						if v18 != 0 {
						} else {
							v20 = *(*int32)(unsafe.Add(mBase, _c_F_UnregisterSnapshotNoOwner[1]))
							if v20 != 0 {
								v22 = *(*int32)(unsafe.Add(mBase, _c_F_UnregisterSnapshotNoOwner[2]))
								v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
								v25 = *(*int32)(unsafe.Add(mBase, _c_F_UnregisterSnapshotNoOwner[1]))
								v27 = v25 - int32(48)
								v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
								if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v28))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v23)) == int32(0) {
									v40 = base.B2i32(base.Ui32(v23) < base.Ui32(v28))
								} else {
									v40 = int32(base.Ui32(v23-v28) >> (uint(int32(31)) % 32))
								}
								if v40 == int32(0) {
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
									v44 = v43
									*(*int32)(unsafe.Add(mBase, _c_F_UnregisterSnapshotNoOwner[3])) = v44
									v48 = *(*int32)(unsafe.Add(mBase, _c_F_UnregisterSnapshotNoOwner[2]))
									*(*int32)(unsafe.Add(mBase, uint32(v48)+40)) = v44
								}
							} else {
								v44 = int32(0)
								*(*int32)(unsafe.Add(mBase, _c_F_UnregisterSnapshotNoOwner[3])) = v44
								v48 = *(*int32)(unsafe.Add(mBase, _c_F_UnregisterSnapshotNoOwner[2]))
								*(*int32)(unsafe.Add(mBase, uint32(v48)+40)) = v44
							}
						}
						return
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
func F_unaccent_dict(m *base.Module, l0 int32) int32 {
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
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
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
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
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(1)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v12 == v11 {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
		v20 = F_get_func_namespace(m, v19)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = int32(0)
			v26 = F_GetSysCacheOid(m, int32(75), int32(_a_F_unaccent_dict_0), v20, v24, v24)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 != 0 {
					v49 = int32(0)
					v50 = v26
					v54 = l0 + v49<<(uint(int32(3))%32)
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
					v56 = F_pg_detoast_datum_packed(m, v55)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						v58 = F_lookup_ts_dictionary_cache(m, v50)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							v61 = v54 + int32(20)
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v58)+44))
							v66 = int32(1)
							v67 = v56 + v66
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
							v72 = v70 & v66
							if v72 != 0 {
								v73 = v67
							} else {
								v73 = v56 + int32(4)
							}
							if v70 == int32(1) {
								v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
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
								if v72 != 0 {
									v100 = int32(base.Ui32(v70)>>(uint(v90)%32)) - v90
								} else {
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
									v100 = int32(base.Ui32(v94)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							v102 = F_FunctionCall4Coll(m, v58+int32(12), int32(0), v65, v73, v100, int32(0))
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return int32(0)
							} else {
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
								if v104 != v56 {
									F_pfree(m, v56)
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return int32(0)
									} else {
										if v102 == int32(0) {
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
											v111 = F_pg_detoast_datum_copy(m, v110)
											mBase = m.M
											v112 = m.ExcPending
											if v112 != 0 {
												return int32(0)
											} else {
												v128 = v111
												m.G0 = v9 + int32(16)
												return v128
											}
										} else {
											v113 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
											if v113 == int32(0) {
												F_pfree(m, v102)
												mBase = m.M
												v117 = m.ExcPending
												if v117 != 0 {
													return int32(0)
												} else {
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
													v119 = F_pg_detoast_datum_copy(m, v118)
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return int32(0)
													} else {
														v128 = v119
														m.G0 = v9 + int32(16)
														return v128
													}
												}
											} else {
												v121 = F_cstring_to_text(m, v113)
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return int32(0)
												} else {
													v123 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
													F_pfree(m, v123)
													mBase = m.M
													v125 = m.ExcPending
													if v125 != 0 {
														return int32(0)
													} else {
														F_pfree(m, v102)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return int32(0)
														} else {
															v128 = v121
															m.G0 = v9 + int32(16)
															return v128
														}
													}
												}
											}
										}
									}
								} else {
									if v102 == int32(0) {
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
										v111 = F_pg_detoast_datum_copy(m, v110)
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return int32(0)
										} else {
											v128 = v111
											m.G0 = v9 + int32(16)
											return v128
										}
									} else {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
										if v113 == int32(0) {
											F_pfree(m, v102)
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int32(0)
											} else {
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
												v119 = F_pg_detoast_datum_copy(m, v118)
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return int32(0)
												} else {
													v128 = v119
													m.G0 = v9 + int32(16)
													return v128
												}
											}
										} else {
											v121 = F_cstring_to_text(m, v113)
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return int32(0)
											} else {
												v123 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
												F_pfree(m, v123)
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v102)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int32(0)
													} else {
														v128 = v121
														m.G0 = v9 + int32(16)
														return v128
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
						return int32(0)
					} else {
						F_errcode(m, int32(67137668))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = F_get_namespace_name(m, v20)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(_a_F_unaccent_dict_0)
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v35
								F_errmsg(m, int32(_a_F_unaccent_dict_1), v9)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_unaccent_dict_2), int32(464), int32(_a_F_unaccent_dict_3))
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
				}
			}
		}
	} else {
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v49 = v11
		v50 = v48
		v54 = l0 + v49<<(uint(int32(3))%32)
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
		v56 = F_pg_detoast_datum_packed(m, v55)
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int32(0)
		} else {
			v58 = F_lookup_ts_dictionary_cache(m, v50)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v61 = v54 + int32(20)
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v58)+44))
				v66 = int32(1)
				v67 = v56 + v66
				v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
				v72 = v70 & v66
				if v72 != 0 {
					v73 = v67
				} else {
					v73 = v56 + int32(4)
				}
				if v70 == int32(1) {
					v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
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
					if v72 != 0 {
						v100 = int32(base.Ui32(v70)>>(uint(v90)%32)) - v90
					} else {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
						v100 = int32(base.Ui32(v94)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v102 = F_FunctionCall4Coll(m, v58+int32(12), int32(0), v65, v73, v100, int32(0))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
					if v104 != v56 {
						F_pfree(m, v56)
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							if v102 == int32(0) {
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
								v111 = F_pg_detoast_datum_copy(m, v110)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									v128 = v111
									m.G0 = v9 + int32(16)
									return v128
								}
							} else {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
								if v113 == int32(0) {
									F_pfree(m, v102)
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return int32(0)
									} else {
										v118 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
										v119 = F_pg_detoast_datum_copy(m, v118)
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return int32(0)
										} else {
											v128 = v119
											m.G0 = v9 + int32(16)
											return v128
										}
									}
								} else {
									v121 = F_cstring_to_text(m, v113)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int32(0)
									} else {
										v123 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
										F_pfree(m, v123)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v102)
											mBase = m.M
											v127 = m.ExcPending
											if v127 != 0 {
												return int32(0)
											} else {
												v128 = v121
												m.G0 = v9 + int32(16)
												return v128
											}
										}
									}
								}
							}
						}
					} else {
						if v102 == int32(0) {
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
							v111 = F_pg_detoast_datum_copy(m, v110)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								v128 = v111
								m.G0 = v9 + int32(16)
								return v128
							}
						} else {
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
							if v113 == int32(0) {
								F_pfree(m, v102)
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return int32(0)
								} else {
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
									v119 = F_pg_detoast_datum_copy(m, v118)
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return int32(0)
									} else {
										v128 = v119
										m.G0 = v9 + int32(16)
										return v128
									}
								}
							} else {
								v121 = F_cstring_to_text(m, v113)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int32(0)
								} else {
									v123 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
									F_pfree(m, v123)
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v102)
										mBase = m.M
										v127 = m.ExcPending
										if v127 != 0 {
											return int32(0)
										} else {
											v128 = v121
											m.G0 = v9 + int32(16)
											return v128
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
func F_unaccent_init(m *base.Module, l0 int32) int32 {
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
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v9 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L17
	} else {
		goto L29
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
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
	v72 = m.ExcPending
	if v72 != 0 {
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
	return v59
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
	return int32(0)
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
	v75 = m.ExcPending
	if v75 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v76
	F_errmsg(m, int32(_a_F_unaccent_init_1), v7)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_unaccent_init_2), int32(363), int32(_a_F_unaccent_init_3))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
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
	v92 = m.ExcPending
	if v92 != 0 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(_a_F_unaccent_init_4), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L17
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_unaccent_init_2), int32(371), int32(_a_F_unaccent_init_3))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
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
	v108 = m.ExcPending
	if v108 != 0 {
		goto L17
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(_a_F_unaccent_init_5), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L17
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_unaccent_init_2), int32(354), int32(_a_F_unaccent_init_3))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
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
								F_errfinish(m, int32(_a_F_unlink_if_exists_fname_1), int32(3849), int32(_a_F_unlink_if_exists_fname_2))
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
func F_untransformRelOptions(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
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
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 == v2 {
		v71 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v71
L2:
	;
	v13 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	F_deconstruct_array_builtin(m, v13, int32(25), v9+int32(12), int32(0), v9+int32(8))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v25 <= int32(0) {
		v71 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v29 = int32(0)
	v33 = v2
	goto L7
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v29<<(uint(int32(2))%32))))
	v41 = F_text_to_cstring(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v71 = v61
	goto L1
L9:
	;
	v43 = int32(61)
	v44 = F___strchrnul(m, v41, v43)
	mBase = m.M
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v46 == v43 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v50 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v50 = v44
	goto L13
L12:
	;
	v50 = int32(0)
	goto L13
L13:
	;
	goto L10
L14:
	;
	v51 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v50))) = uint8(v51)
	v55 = F_makeString(m, v50+int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L3
	} else {
		goto L17
	}
L15:
	;
	v57 = int32(0)
	goto L16
L16:
	;
	v59 = F_makeDefElem(m, v41, v57, int32(-1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L3
	} else {
		goto L18
	}
L17:
	;
	v57 = v55
	goto L16
L18:
	;
	v61 = F_lappend(m, v33, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v64 = v29 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v64 < v65 {
		v29 = v64
		v33 = v61
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	var v72 int64
	_ = v72
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v182 int32
	_ = v182
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v240 int32
	_ = v240
	var v263 int32
	_ = v263
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v470 int32
	_ = v470
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	v5 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(192)
	m.G0 = v26
	if v5 < l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v32 = F_table_open(m, int32(2619), int32(3))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v26 + int32(192)
	return
L4:
	;
	return
L5:
	;
	v43 = v5
	v50 = v5
	goto L6
L6:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l3+v50<<(uint(int32(2))%32))))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+36)))
	if v61 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v470 != 0 {
		goto L72
	} else {
		goto L73
	}
L8:
	;
	v64 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+55)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v26)+40)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v64
	v72 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v26)+23)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = l0
	v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+224)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+72)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v81
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v60)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v60)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v60)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+84)) = v88
	v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+52)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+88)) = v90
	v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+54)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+92)) = v92
	v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+56)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = v94
	v96 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+58)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+100)) = v96
	v98 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+60)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+104)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v60)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+108)) = v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v60)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+112)) = v102
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v60)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+116)) = v104
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v60)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+120)) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v60)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+124)) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v60)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+128)) = v110
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v60)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+132)) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v60)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+136)) = v114
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v60)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+140)) = v116
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v60)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+144)) = v118
	v139 = int32(21)
	v140 = int32(0)
	goto L11
L9:
	;
	v470 = v43
	goto L10
L10:
	;
	v485 = v50 + int32(1)
	if v485 != l2 {
		v43 = v470
		v50 = v485
		goto L6
	} else {
		goto L71
	}
L11:
	;
	v151 = int32(2)
	v155 = v140 << (uint(v151) % 32)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v60+int32(104)+v155)))
	if int32(0) < v157 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v60)+144))
	if int32(0) < v353 {
		goto L32
	} else {
		goto L33
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26-int32(-64)+v139<<(uint(v151)%32)))) = v345
	v347 = int32(1)
	v350 = v140 + v347
	if v350 != int32(5) {
		v139 = v139 + v347
		v140 = v350
		goto L11
	} else {
		goto L30
	}
L14:
	;
	v161 = v157 & int32(3)
	v162 = v155 + (v60 + int32(124))
	v166 = F_palloc(m, v157<<(uint(int32(2))%32))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v319 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26+int32(32)+v139))) = uint8(v319)
	v345 = int32(0)
	goto L13
L17:
	;
	v168 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v157) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v314 = F_construct_array_builtin(m, v166, v157, int32(700))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L4
	} else {
		goto L29
	}
L19:
	;
	v182 = v168
	v193 = int32(0)
	goto L22
L20:
	;
	v240 = v168
	goto L21
L21:
	;
	v263 = v240
	v273 = int32(0)
	goto L26
L22:
	;
	v198 = v182 << (uint(int32(2)) % 32)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v200+v198)))
	*(*int32)(unsafe.Add(mBase, uint32(v166+v198))) = v202
	v204 = int32(4)
	v205 = v198 | v204
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v207+v205)))
	*(*int32)(unsafe.Add(mBase, uint32(v166+v205))) = v209
	v212 = v198 | int32(8)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v214+v212)))
	*(*int32)(unsafe.Add(mBase, uint32(v166+v212))) = v216
	v219 = v198 | int32(12)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v221+v219)))
	*(*int32)(unsafe.Add(mBase, uint32(v166+v219))) = v223
	v226 = v182 + v204
	v228 = v193 + v204
	if v228 != v157&int32(2147483644) {
		v182 = v226
		v193 = v228
		goto L22
	} else {
		goto L24
	}
L23:
	;
	if v161 == int32(0) {
		goto L18
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v240 = v226
	goto L21
L26:
	;
	v279 = v263 << (uint(int32(2)) % 32)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v281+v279)))
	*(*int32)(unsafe.Add(mBase, uint32(v166+v279))) = v283
	v285 = int32(1)
	v288 = v273 + v285
	if v288 != v161 {
		v263 = v263 + v285
		v273 = v288
		goto L26
	} else {
		goto L28
	}
L27:
	;
	goto L18
L28:
	;
	goto L27
L29:
	;
	v345 = v314
	goto L13
L30:
	;
	goto L12
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+168)) = v366
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v60)+148))
	if v368 <= int32(0) {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v60)+164))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v60)+184))
	v358 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+204)))
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+214)))
	v360 = int32(*(*int8)(unsafe.Add(mBase, uint32(v60)+219)))
	v361 = F_construct_array(m, v356, v353, v357, v358, v359, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v363 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+58)) = uint8(v363)
	v366 = int32(0)
	goto L31
L35:
	;
	v366 = v361
	goto L31
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+172)) = v381
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v60)+152))
	if v383 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L37:
	;
	v371 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+59)) = uint8(v371)
	v381 = int32(0)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v60)+168))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v60)+188))
	v376 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+206)))
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+215)))
	v378 = int32(*(*int8)(unsafe.Add(mBase, uint32(v60)+220)))
	v379 = F_construct_array(m, v374, v368, v375, v376, v377, v378)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v381 = v379
	goto L36
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+176)) = v396
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v60)+156))
	if v398 <= int32(0) {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	v386 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+60)) = uint8(v386)
	v396 = int32(0)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v60)+172))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v60)+192))
	v391 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+208)))
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+216)))
	v393 = int32(*(*int8)(unsafe.Add(mBase, uint32(v60)+221)))
	v394 = F_construct_array(m, v389, v383, v390, v391, v392, v393)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v396 = v394
	goto L41
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+180)) = v411
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v60)+160))
	if v413 <= int32(0) {
		goto L52
	} else {
		goto L53
	}
L47:
	;
	v401 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+61)) = uint8(v401)
	v411 = int32(0)
	goto L46
L48:
	;
	goto L49
L49:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v60)+176))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v60)+196))
	v406 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+210)))
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+217)))
	v408 = int32(*(*int8)(unsafe.Add(mBase, uint32(v60)+222)))
	v409 = F_construct_array(m, v404, v398, v405, v406, v407, v408)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v411 = v409
	goto L46
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+184)) = v426
	v429 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+224)))
	v430 = F_SearchSysCache3(m, int32(65), l0, v429, l1)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L4
	} else {
		goto L56
	}
L52:
	;
	v416 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+62)) = uint8(v416)
	v426 = int32(0)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v60)+180))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v60)+200))
	v421 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+212)))
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+218)))
	v423 = int32(*(*int8)(unsafe.Add(mBase, uint32(v60)+223)))
	v424 = F_construct_array(m, v419, v413, v420, v421, v422, v423)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v426 = v424
	goto L51
L56:
	;
	if v43 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v434 = F_CatalogOpenIndexes(m, v32)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L4
	} else {
		goto L60
	}
L58:
	;
	v436 = v43
	goto L59
L59:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	if v430 != 0 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v436 = v434
	goto L59
L61:
	;
	F_pfree(m, v458)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L4
	} else {
		goto L70
	}
L62:
	;
	v442 = F_heap_modify_tuple(m, v430, v437, v26-int32(-64), v26+int32(32), v26)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L4
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v454 = F_heap_form_tuple(m, v437, v26-int32(-64), v26+int32(32))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L4
	} else {
		goto L68
	}
L65:
	;
	F_ReleaseCatCache(m, v430)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	F_CatalogTupleUpdateWithInfo(m, v32, v442+int32(4), v442, v436)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	v458 = v442
	goto L61
L68:
	;
	F_CatalogTupleInsertWithInfo(m, v32, v454, v436)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v458 = v454
	goto L61
L70:
	;
	v470 = v436
	goto L10
L71:
	;
	goto L7
L72:
	;
	F_CatalogCloseIndexes(m, v470)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L4
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	F_relation_close(m, v32, int32(3))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L4
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
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
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(993)
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
