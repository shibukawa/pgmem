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
		v3 = *(*int32)(unsafe.Add(mBase, _consts[173]))
		F_ResourceOwnerForget(m, v3, l0, int32(1744708))
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
		F_pairingheap_remove(m, int32(4141224), l0+int32(52))
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
						v18 = *(*int32)(unsafe.Add(mBase, _consts[254]))
						if v18 != 0 {
						} else {
							v20 = *(*int32)(unsafe.Add(mBase, _consts[255]))
							if v20 != 0 {
								v22 = *(*int32)(unsafe.Add(mBase, _consts[120]))
								v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
								v25 = *(*int32)(unsafe.Add(mBase, _consts[255]))
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
									*(*int32)(unsafe.Add(mBase, _consts[83])) = v44
									v48 = *(*int32)(unsafe.Add(mBase, _consts[120]))
									*(*int32)(unsafe.Add(mBase, uint32(v48)+40)) = v44
								}
							} else {
								v44 = int32(0)
								*(*int32)(unsafe.Add(mBase, _consts[83])) = v44
								v48 = *(*int32)(unsafe.Add(mBase, _consts[120]))
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
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
			v26 = F_GetSysCacheOid(m, int32(75), int32(95555), v20, v24, v24)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 != 0 {
					v54 = int32(0)
					v56 = v26
					v59 = l0 + v54<<(uint(int32(3))%32)
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
					v61 = F_pg_detoast_datum_packed(m, v60)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v63 = F_lookup_ts_dictionary_cache(m, v56)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v66 = v59 + int32(20)
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v63)+44))
							v71 = int32(1)
							v72 = v61 + v71
							v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
							v77 = v75 & v71
							if v77 != 0 {
								v78 = v72
							} else {
								v78 = v61 + int32(4)
							}
							if v75 == int32(1) {
								v81 = int32(4)
								v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
								if v83&int32(254) == int32(2) {
									v92 = v81
								} else {
									v92 = base.B2i32(v83 == int32(18)) << (uint(v81) % 32)
								}
								if v83 == int32(1) {
									v95 = v81
								} else {
									v95 = v92
								}
								v106 = v95
							} else {
								v96 = int32(1)
								if v77 != 0 {
									v106 = int32(base.Ui32(v75)>>(uint(v96)%32)) - v96
								} else {
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
									v106 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							v108 = F_FunctionCall4Coll(m, v63+int32(12), int32(0), v70, v78, v106, int32(0))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int32(0)
							} else {
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
								if v110 != v61 {
									F_pfree(m, v61)
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										if v108 == int32(0) {
											v116 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
											v117 = F_pg_detoast_datum_copy(m, v116)
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return int32(0)
											} else {
												v134 = v117
												m.G0 = v9 + int32(16)
												return v134
											}
										} else {
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
											if v119 == int32(0) {
												F_pfree(m, v108)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return int32(0)
												} else {
													v124 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
													v125 = F_pg_detoast_datum_copy(m, v124)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return int32(0)
													} else {
														v134 = v125
														m.G0 = v9 + int32(16)
														return v134
													}
												}
											} else {
												v127 = F_cstring_to_text(m, v119)
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return int32(0)
												} else {
													v129 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
													F_pfree(m, v129)
													mBase = m.M
													v131 = m.ExcPending
													if v131 != 0 {
														return int32(0)
													} else {
														F_pfree(m, v108)
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return int32(0)
														} else {
															v134 = v127
															m.G0 = v9 + int32(16)
															return v134
														}
													}
												}
											}
										}
									}
								} else {
									if v108 == int32(0) {
										v116 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
										v117 = F_pg_detoast_datum_copy(m, v116)
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											v134 = v117
											m.G0 = v9 + int32(16)
											return v134
										}
									} else {
										v119 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
										if v119 == int32(0) {
											F_pfree(m, v108)
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return int32(0)
											} else {
												v124 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
												v125 = F_pg_detoast_datum_copy(m, v124)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return int32(0)
												} else {
													v134 = v125
													m.G0 = v9 + int32(16)
													return v134
												}
											}
										} else {
											v127 = F_cstring_to_text(m, v119)
											mBase = m.M
											v128 = m.ExcPending
											if v128 != 0 {
												return int32(0)
											} else {
												v129 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
												F_pfree(m, v129)
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v108)
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int32(0)
													} else {
														v134 = v127
														m.G0 = v9 + int32(16)
														return v134
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
							v36 = F_get_namespace_name(m, v20)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(95555)
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v36
								F_errmsg(m, int32(69507), v9)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(485561), int32(464), int32(108169))
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
						}
					}
				}
			}
		}
	} else {
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v54 = v11
		v56 = v53
		v59 = l0 + v54<<(uint(int32(3))%32)
		v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
		v61 = F_pg_detoast_datum_packed(m, v60)
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return int32(0)
		} else {
			v63 = F_lookup_ts_dictionary_cache(m, v56)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				v66 = v59 + int32(20)
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v63)+44))
				v71 = int32(1)
				v72 = v61 + v71
				v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
				v77 = v75 & v71
				if v77 != 0 {
					v78 = v72
				} else {
					v78 = v61 + int32(4)
				}
				if v75 == int32(1) {
					v81 = int32(4)
					v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
					if v83&int32(254) == int32(2) {
						v92 = v81
					} else {
						v92 = base.B2i32(v83 == int32(18)) << (uint(v81) % 32)
					}
					if v83 == int32(1) {
						v95 = v81
					} else {
						v95 = v92
					}
					v106 = v95
				} else {
					v96 = int32(1)
					if v77 != 0 {
						v106 = int32(base.Ui32(v75)>>(uint(v96)%32)) - v96
					} else {
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
						v106 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v108 = F_FunctionCall4Coll(m, v63+int32(12), int32(0), v70, v78, v106, int32(0))
				mBase = m.M
				v109 = m.ExcPending
				if v109 != 0 {
					return int32(0)
				} else {
					v110 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
					if v110 != v61 {
						F_pfree(m, v61)
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int32(0)
						} else {
							if v108 == int32(0) {
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
								v117 = F_pg_detoast_datum_copy(m, v116)
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									v134 = v117
									m.G0 = v9 + int32(16)
									return v134
								}
							} else {
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
								if v119 == int32(0) {
									F_pfree(m, v108)
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return int32(0)
									} else {
										v124 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
										v125 = F_pg_detoast_datum_copy(m, v124)
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return int32(0)
										} else {
											v134 = v125
											m.G0 = v9 + int32(16)
											return v134
										}
									}
								} else {
									v127 = F_cstring_to_text(m, v119)
									mBase = m.M
									v128 = m.ExcPending
									if v128 != 0 {
										return int32(0)
									} else {
										v129 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
										F_pfree(m, v129)
										mBase = m.M
										v131 = m.ExcPending
										if v131 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v108)
											mBase = m.M
											v133 = m.ExcPending
											if v133 != 0 {
												return int32(0)
											} else {
												v134 = v127
												m.G0 = v9 + int32(16)
												return v134
											}
										}
									}
								}
							}
						}
					} else {
						if v108 == int32(0) {
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
							v117 = F_pg_detoast_datum_copy(m, v116)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								v134 = v117
								m.G0 = v9 + int32(16)
								return v134
							}
						} else {
							v119 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
							if v119 == int32(0) {
								F_pfree(m, v108)
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return int32(0)
								} else {
									v124 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
									v125 = F_pg_detoast_datum_copy(m, v124)
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return int32(0)
									} else {
										v134 = v125
										m.G0 = v9 + int32(16)
										return v134
									}
								}
							} else {
								v127 = F_cstring_to_text(m, v119)
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return int32(0)
								} else {
									v129 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
									F_pfree(m, v129)
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v108)
										mBase = m.M
										v133 = m.ExcPending
										if v133 != 0 {
											return int32(0)
										} else {
											v134 = v127
											m.G0 = v9 + int32(16)
											return v134
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v10 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v110
L2:
	;
	v11 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v14 = base.B2i32(v11 < v12)
	if v14 == v11 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L24
	} else {
		goto L37
	}
L5:
	;
	if v11 < v12 {
		goto L1
	} else {
		goto L36
	}
L6:
	;
	v110 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v17 = v11
	goto L10
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L24
	} else {
		goto L32
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22+v17<<(uint(int32(2))%32))))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v28 = int32(161560)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1559])))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v32 == int32(0) {
		v51 = v31
		v52 = v32
		goto L13
	} else {
		goto L14
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L24
	} else {
		goto L28
	}
L12:
	;
	if v52-v51 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	goto L12
L14:
	;
	if v31 != v32 {
		v51 = v31
		v52 = v32
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v36 = v27
	v37 = v28
	goto L16
L16:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v41 == int32(0) {
		v51 = v40
		v52 = v41
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v51 = v40
	v52 = v41
	goto L13
L18:
	;
	v44 = int32(1)
	if v40 == v41 {
		v36 = v36 + v44
		v37 = v37 + v44
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	if v17 == int32(1) {
		goto L9
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	goto L11
L23:
	;
	v58 = F_defGetString(m, v26)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	v62 = F_initTrie(m, v58)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v65 = v17 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v65 < v66 {
		v17 = v65
		goto L10
	} else {
		goto L27
	}
L27:
	;
	v110 = v62
	goto L5
L28:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v75
	F_errmsg(m, int32(701531), v8)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(485561), int32(363), int32(98496))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	F_errmsg(m, int32(130320), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L24
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(485561), int32(354), int32(98496))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L24
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	goto L4
L37:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L24
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(213772), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L24
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(485561), int32(371), int32(98496))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L24
	} else {
		goto L40
	}
L40:
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
			v13 = *(*int32)(unsafe.Add(mBase, _consts[140]))
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
							F_errmsg(m, int32(292168), v7)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								F_errfinish(m, int32(492032), int32(3849), int32(372012))
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
		v12 = *(*int32)(unsafe.Add(mBase, _consts[140]))
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
						F_errmsg(m, int32(295023), v6)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							F_errfinish(m, int32(491528), int32(6966), int32(380328))
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
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
	var v45 int32
	_ = v45
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v68
L2:
	;
	v68 = v2
	goto L1
L3:
	;
	goto L4
L4:
	;
	v13 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	F_deconstruct_array_builtin(m, v13, int32(25), v9+int32(12), int32(0), v9+int32(8))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v25 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v68 = v2
	goto L1
L9:
	;
	goto L10
L10:
	;
	v30 = v2
	v31 = v2
	goto L11
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v31<<(uint(int32(2))%32))))
	v40 = F_text_to_cstring(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v68 = v60
	goto L1
L13:
	;
	v42 = int32(61)
	v43 = F___strchrnul(m, v40, v42)
	mBase = m.M
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v45 == v42 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v49 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v49 = v43
	goto L17
L16:
	;
	v49 = int32(0)
	goto L17
L17:
	;
	goto L14
L18:
	;
	v50 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v50)
	v54 = F_makeString(m, v49+int32(1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	v56 = int32(0)
	goto L20
L20:
	;
	v58 = F_makeDefElem(m, v40, v56, int32(-1))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L22
	}
L21:
	;
	v56 = v54
	goto L20
L22:
	;
	v60 = F_lappend(m, v30, v58)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v63 = v31 + int32(1)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v63 < v64 {
		v30 = v60
		v31 = v63
		goto L11
	} else {
		goto L24
	}
L24:
	;
	goto L12
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
	var v68 int64
	_ = v68
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
	var v180 int32
	_ = v180
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
	var v236 int32
	_ = v236
	var v259 int32
	_ = v259
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
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
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
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
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
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
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
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
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
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
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v468 int32
	_ = v468
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
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
	if v468 != 0 {
		goto L73
	} else {
		goto L74
	}
L8:
	;
	v64 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+55)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	v68 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = v68
	*(*int64)(unsafe.Add(mBase, uint32(v26)+23)) = v68
	*(*int64)(unsafe.Add(mBase, uint32(v26)+40)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = v68
	*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = v68
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
	v468 = v43
	goto L10
L10:
	;
	v483 = v50 + int32(1)
	if v483 != l2 {
		v43 = v468
		v50 = v483
		goto L6
	} else {
		goto L72
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
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v60)+144))
	if int32(0) < v351 {
		goto L33
	} else {
		goto L34
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26-int32(-64)+v139<<(uint(v151)%32)))) = v343
	v345 = int32(1)
	v348 = v140 + v345
	if v348 != int32(5) {
		v139 = v139 + v345
		v140 = v348
		goto L11
	} else {
		goto L31
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
	v317 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26+int32(32)+v139))) = uint8(v317)
	v343 = int32(0)
	goto L13
L17:
	;
	v168 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v157) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v180 = v168
	v193 = int32(0)
	goto L21
L19:
	;
	v236 = v168
	goto L20
L20:
	;
	if v161 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v198 = v180 << (uint(int32(2)) % 32)
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
	v226 = v180 + v204
	v228 = v193 + v204
	if v228 != v157&int32(2147483644) {
		v180 = v226
		v193 = v228
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v236 = v226
	goto L20
L23:
	;
	goto L22
L24:
	;
	v259 = v236
	v271 = int32(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	v312 = F_construct_array_builtin(m, v166, v157, int32(700))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L4
	} else {
		goto L30
	}
L27:
	;
	v277 = v259 << (uint(int32(2)) % 32)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v279+v277)))
	*(*int32)(unsafe.Add(mBase, uint32(v166+v277))) = v281
	v283 = int32(1)
	v286 = v271 + v283
	if v286 != v161 {
		v259 = v259 + v283
		v271 = v286
		goto L27
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	goto L28
L30:
	;
	v343 = v312
	goto L13
L31:
	;
	goto L12
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+168)) = v364
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v60)+148))
	if v366 <= int32(0) {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v60)+164))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v60)+184))
	v356 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+204)))
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+214)))
	v358 = int32(*(*int8)(unsafe.Add(mBase, uint32(v60)+219)))
	v359 = F_construct_array(m, v354, v351, v355, v356, v357, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v361 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+58)) = uint8(v361)
	v364 = int32(0)
	goto L32
L36:
	;
	v364 = v359
	goto L32
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+172)) = v379
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v60)+152))
	if v381 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L38:
	;
	v369 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+59)) = uint8(v369)
	v379 = int32(0)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v60)+168))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v60)+188))
	v374 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+206)))
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+215)))
	v376 = int32(*(*int8)(unsafe.Add(mBase, uint32(v60)+220)))
	v377 = F_construct_array(m, v372, v366, v373, v374, v375, v376)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v379 = v377
	goto L37
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+176)) = v394
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v60)+156))
	if v396 <= int32(0) {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	v384 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+60)) = uint8(v384)
	v394 = int32(0)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v60)+172))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v60)+192))
	v389 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+208)))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+216)))
	v391 = int32(*(*int8)(unsafe.Add(mBase, uint32(v60)+221)))
	v392 = F_construct_array(m, v387, v381, v388, v389, v390, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v394 = v392
	goto L42
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+180)) = v409
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v60)+160))
	if v411 <= int32(0) {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v399 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+61)) = uint8(v399)
	v409 = int32(0)
	goto L47
L49:
	;
	goto L50
L50:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v60)+176))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v60)+196))
	v404 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+210)))
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+217)))
	v406 = int32(*(*int8)(unsafe.Add(mBase, uint32(v60)+222)))
	v407 = F_construct_array(m, v402, v396, v403, v404, v405, v406)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v409 = v407
	goto L47
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+184)) = v424
	v427 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+224)))
	v428 = F_SearchSysCache3(m, int32(65), l0, v427, l1)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L4
	} else {
		goto L57
	}
L53:
	;
	v414 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+62)) = uint8(v414)
	v424 = int32(0)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v60)+180))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v60)+200))
	v419 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+212)))
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+218)))
	v421 = int32(*(*int8)(unsafe.Add(mBase, uint32(v60)+223)))
	v422 = F_construct_array(m, v417, v411, v418, v419, v420, v421)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v424 = v422
	goto L52
L57:
	;
	if v43 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v432 = F_CatalogOpenIndexes(m, v32)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L4
	} else {
		goto L61
	}
L59:
	;
	v434 = v43
	goto L60
L60:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	if v428 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v434 = v432
	goto L60
L62:
	;
	F_pfree(m, v456)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L4
	} else {
		goto L71
	}
L63:
	;
	v440 = F_heap_modify_tuple(m, v428, v435, v26-int32(-64), v26+int32(32), v26)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L4
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v452 = F_heap_form_tuple(m, v435, v26-int32(-64), v26+int32(32))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L4
	} else {
		goto L69
	}
L66:
	;
	F_ReleaseCatCache(m, v428)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	F_CatalogTupleUpdateWithInfo(m, v32, v440+int32(4), v440, v434)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v456 = v440
	goto L62
L69:
	;
	F_CatalogTupleInsertWithInfo(m, v32, v452, v434)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v456 = v452
	goto L62
L71:
	;
	v468 = v434
	goto L10
L72:
	;
	goto L7
L73:
	;
	F_CatalogCloseIndexes(m, v468)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L4
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_sequence_close(m, v32, int32(3))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L4
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
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
	v11 = int32(4463656)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[385]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, _consts[385])) = v9 + int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(240032)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(992)
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
			*(*int32)(unsafe.Add(mBase, _consts[385])) = v41
			m.G0 = v9 + int32(32)
			return
		}
	} else {
		v41 = v12
		*(*int32)(unsafe.Add(mBase, _consts[385])) = v41
		m.G0 = v9 + int32(32)
		return
	}
}
