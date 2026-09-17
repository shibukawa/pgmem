package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_stat_get_analyze_count(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+160))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_archiver(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int64
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int64
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int64
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int64
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	v2 = int32(0)
	v3 = int64(0)
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v3
	*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v3
	*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v6)+11)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v2
	v21 = F_CreateTemplateTupleDesc(m, int32(7))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		F_TupleDescInitEntry(m, v21, int32(1), int32(_a_F_pg_stat_get_archiver_0), int32(20), int32(-1), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			F_TupleDescInitEntry(m, v21, int32(2), int32(_a_F_pg_stat_get_archiver_1), int32(25), int32(-1), int32(0))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				F_TupleDescInitEntry(m, v21, int32(3), int32(_a_F_pg_stat_get_archiver_2), int32(1184), int32(-1), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_TupleDescInitEntry(m, v21, int32(4), int32(_a_F_pg_stat_get_archiver_3), int32(20), int32(-1), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						F_TupleDescInitEntry(m, v21, int32(5), int32(_a_F_pg_stat_get_archiver_4), int32(25), int32(-1), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_TupleDescInitEntry(m, v21, int32(6), int32(_a_F_pg_stat_get_archiver_5), int32(1184), int32(-1), int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_TupleDescInitEntry(m, v21, int32(7), int32(_a_F_pg_stat_get_archiver_6), int32(1184), int32(-1), int32(0))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v74 = F_BlessTupleDesc(m, v21)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										F_pgstat_snapshot_fixed(m, int32(7))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											v80 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[0]))
											v81 = F_Int64GetDatum(m, v80)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v81
												v85 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[1])))
												if v85 == int32(0) {
													v88 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v6)+9)) = uint8(v88)
													v95 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[2]))
													if v95 == int64(0) {
														v98 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v6)+10)) = uint8(v98)
														v104 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[3]))
														v105 = F_Int64GetDatum(m, v104)
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v105
															v109 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[4])))
															if v109 == int32(0) {
																v112 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)) = uint8(v112)
																v119 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[5]))
																if v119 == int64(0) {
																	v122 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v122)
																	v128 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[6]))
																	if v128 == int64(0) {
																		v131 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v131)
																		v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																		mBase = m.M
																		v141 = m.ExcPending
																		if v141 != 0 {
																			return int32(0)
																		} else {
																			v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																			v143 = F_HeapTupleHeaderGetDatum(m, v142)
																			mBase = m.M
																			v144 = m.ExcPending
																			if v144 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v6 + int32(48)
																				return v143
																			}
																		}
																	} else {
																		v133 = F_Int64GetDatum(m, v128)
																		mBase = m.M
																		v134 = m.ExcPending
																		if v134 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v133
																			v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																			mBase = m.M
																			v141 = m.ExcPending
																			if v141 != 0 {
																				return int32(0)
																			} else {
																				v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																				v143 = F_HeapTupleHeaderGetDatum(m, v142)
																				mBase = m.M
																				v144 = m.ExcPending
																				if v144 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v6 + int32(48)
																					return v143
																				}
																			}
																		}
																	}
																} else {
																	v124 = F_Int64GetDatum(m, v119)
																	mBase = m.M
																	v125 = m.ExcPending
																	if v125 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v124
																		v128 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[6]))
																		if v128 == int64(0) {
																			v131 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v131)
																			v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																			mBase = m.M
																			v141 = m.ExcPending
																			if v141 != 0 {
																				return int32(0)
																			} else {
																				v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																				v143 = F_HeapTupleHeaderGetDatum(m, v142)
																				mBase = m.M
																				v144 = m.ExcPending
																				if v144 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v6 + int32(48)
																					return v143
																				}
																			}
																		} else {
																			v133 = F_Int64GetDatum(m, v128)
																			mBase = m.M
																			v134 = m.ExcPending
																			if v134 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v133
																				v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return int32(0)
																				} else {
																					v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																					v143 = F_HeapTupleHeaderGetDatum(m, v142)
																					mBase = m.M
																					v144 = m.ExcPending
																					if v144 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v6 + int32(48)
																						return v143
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																v115 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_archiver_7))
																mBase = m.M
																v116 = m.ExcPending
																if v116 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v115
																	v119 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[5]))
																	if v119 == int64(0) {
																		v122 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v122)
																		v128 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[6]))
																		if v128 == int64(0) {
																			v131 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v131)
																			v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																			mBase = m.M
																			v141 = m.ExcPending
																			if v141 != 0 {
																				return int32(0)
																			} else {
																				v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																				v143 = F_HeapTupleHeaderGetDatum(m, v142)
																				mBase = m.M
																				v144 = m.ExcPending
																				if v144 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v6 + int32(48)
																					return v143
																				}
																			}
																		} else {
																			v133 = F_Int64GetDatum(m, v128)
																			mBase = m.M
																			v134 = m.ExcPending
																			if v134 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v133
																				v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return int32(0)
																				} else {
																					v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																					v143 = F_HeapTupleHeaderGetDatum(m, v142)
																					mBase = m.M
																					v144 = m.ExcPending
																					if v144 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v6 + int32(48)
																						return v143
																					}
																				}
																			}
																		}
																	} else {
																		v124 = F_Int64GetDatum(m, v119)
																		mBase = m.M
																		v125 = m.ExcPending
																		if v125 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v124
																			v128 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[6]))
																			if v128 == int64(0) {
																				v131 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v131)
																				v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return int32(0)
																				} else {
																					v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																					v143 = F_HeapTupleHeaderGetDatum(m, v142)
																					mBase = m.M
																					v144 = m.ExcPending
																					if v144 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v6 + int32(48)
																						return v143
																					}
																				}
																			} else {
																				v133 = F_Int64GetDatum(m, v128)
																				mBase = m.M
																				v134 = m.ExcPending
																				if v134 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v133
																					v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																						v143 = F_HeapTupleHeaderGetDatum(m, v142)
																						mBase = m.M
																						v144 = m.ExcPending
																						if v144 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v6 + int32(48)
																							return v143
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
														v100 = F_Int64GetDatum(m, v95)
														mBase = m.M
														v101 = m.ExcPending
														if v101 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v100
															v104 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[3]))
															v105 = F_Int64GetDatum(m, v104)
															mBase = m.M
															v106 = m.ExcPending
															if v106 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v105
																v109 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[4])))
																if v109 == int32(0) {
																	v112 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)) = uint8(v112)
																	v119 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[5]))
																	if v119 == int64(0) {
																		v122 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v122)
																		v128 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[6]))
																		if v128 == int64(0) {
																			v131 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v131)
																			v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																			mBase = m.M
																			v141 = m.ExcPending
																			if v141 != 0 {
																				return int32(0)
																			} else {
																				v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																				v143 = F_HeapTupleHeaderGetDatum(m, v142)
																				mBase = m.M
																				v144 = m.ExcPending
																				if v144 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v6 + int32(48)
																					return v143
																				}
																			}
																		} else {
																			v133 = F_Int64GetDatum(m, v128)
																			mBase = m.M
																			v134 = m.ExcPending
																			if v134 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v133
																				v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return int32(0)
																				} else {
																					v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																					v143 = F_HeapTupleHeaderGetDatum(m, v142)
																					mBase = m.M
																					v144 = m.ExcPending
																					if v144 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v6 + int32(48)
																						return v143
																					}
																				}
																			}
																		}
																	} else {
																		v124 = F_Int64GetDatum(m, v119)
																		mBase = m.M
																		v125 = m.ExcPending
																		if v125 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v124
																			v128 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[6]))
																			if v128 == int64(0) {
																				v131 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v131)
																				v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return int32(0)
																				} else {
																					v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																					v143 = F_HeapTupleHeaderGetDatum(m, v142)
																					mBase = m.M
																					v144 = m.ExcPending
																					if v144 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v6 + int32(48)
																						return v143
																					}
																				}
																			} else {
																				v133 = F_Int64GetDatum(m, v128)
																				mBase = m.M
																				v134 = m.ExcPending
																				if v134 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v133
																					v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																						v143 = F_HeapTupleHeaderGetDatum(m, v142)
																						mBase = m.M
																						v144 = m.ExcPending
																						if v144 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v6 + int32(48)
																							return v143
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	v115 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_archiver_7))
																	mBase = m.M
																	v116 = m.ExcPending
																	if v116 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v115
																		v119 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[5]))
																		if v119 == int64(0) {
																			v122 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v122)
																			v128 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[6]))
																			if v128 == int64(0) {
																				v131 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v131)
																				v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return int32(0)
																				} else {
																					v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																					v143 = F_HeapTupleHeaderGetDatum(m, v142)
																					mBase = m.M
																					v144 = m.ExcPending
																					if v144 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v6 + int32(48)
																						return v143
																					}
																				}
																			} else {
																				v133 = F_Int64GetDatum(m, v128)
																				mBase = m.M
																				v134 = m.ExcPending
																				if v134 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v133
																					v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																						v143 = F_HeapTupleHeaderGetDatum(m, v142)
																						mBase = m.M
																						v144 = m.ExcPending
																						if v144 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v6 + int32(48)
																							return v143
																						}
																					}
																				}
																			}
																		} else {
																			v124 = F_Int64GetDatum(m, v119)
																			mBase = m.M
																			v125 = m.ExcPending
																			if v125 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v124
																				v128 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[6]))
																				if v128 == int64(0) {
																					v131 = int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v131)
																					v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																						v143 = F_HeapTupleHeaderGetDatum(m, v142)
																						mBase = m.M
																						v144 = m.ExcPending
																						if v144 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v6 + int32(48)
																							return v143
																						}
																					}
																				} else {
																					v133 = F_Int64GetDatum(m, v128)
																					mBase = m.M
																					v134 = m.ExcPending
																					if v134 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v133
																						v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return int32(0)
																						} else {
																							v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																							v143 = F_HeapTupleHeaderGetDatum(m, v142)
																							mBase = m.M
																							v144 = m.ExcPending
																							if v144 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v6 + int32(48)
																								return v143
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
												} else {
													v91 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_archiver_8))
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v91
														v95 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[2]))
														if v95 == int64(0) {
															v98 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v6)+10)) = uint8(v98)
															v104 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[3]))
															v105 = F_Int64GetDatum(m, v104)
															mBase = m.M
															v106 = m.ExcPending
															if v106 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v105
																v109 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[4])))
																if v109 == int32(0) {
																	v112 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)) = uint8(v112)
																	v119 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[5]))
																	if v119 == int64(0) {
																		v122 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v122)
																		v128 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[6]))
																		if v128 == int64(0) {
																			v131 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v131)
																			v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																			mBase = m.M
																			v141 = m.ExcPending
																			if v141 != 0 {
																				return int32(0)
																			} else {
																				v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																				v143 = F_HeapTupleHeaderGetDatum(m, v142)
																				mBase = m.M
																				v144 = m.ExcPending
																				if v144 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v6 + int32(48)
																					return v143
																				}
																			}
																		} else {
																			v133 = F_Int64GetDatum(m, v128)
																			mBase = m.M
																			v134 = m.ExcPending
																			if v134 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v133
																				v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return int32(0)
																				} else {
																					v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																					v143 = F_HeapTupleHeaderGetDatum(m, v142)
																					mBase = m.M
																					v144 = m.ExcPending
																					if v144 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v6 + int32(48)
																						return v143
																					}
																				}
																			}
																		}
																	} else {
																		v124 = F_Int64GetDatum(m, v119)
																		mBase = m.M
																		v125 = m.ExcPending
																		if v125 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v124
																			v128 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[6]))
																			if v128 == int64(0) {
																				v131 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v131)
																				v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return int32(0)
																				} else {
																					v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																					v143 = F_HeapTupleHeaderGetDatum(m, v142)
																					mBase = m.M
																					v144 = m.ExcPending
																					if v144 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v6 + int32(48)
																						return v143
																					}
																				}
																			} else {
																				v133 = F_Int64GetDatum(m, v128)
																				mBase = m.M
																				v134 = m.ExcPending
																				if v134 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v133
																					v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																						v143 = F_HeapTupleHeaderGetDatum(m, v142)
																						mBase = m.M
																						v144 = m.ExcPending
																						if v144 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v6 + int32(48)
																							return v143
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	v115 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_archiver_7))
																	mBase = m.M
																	v116 = m.ExcPending
																	if v116 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v115
																		v119 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[5]))
																		if v119 == int64(0) {
																			v122 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v122)
																			v128 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[6]))
																			if v128 == int64(0) {
																				v131 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v131)
																				v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return int32(0)
																				} else {
																					v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																					v143 = F_HeapTupleHeaderGetDatum(m, v142)
																					mBase = m.M
																					v144 = m.ExcPending
																					if v144 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v6 + int32(48)
																						return v143
																					}
																				}
																			} else {
																				v133 = F_Int64GetDatum(m, v128)
																				mBase = m.M
																				v134 = m.ExcPending
																				if v134 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v133
																					v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																						v143 = F_HeapTupleHeaderGetDatum(m, v142)
																						mBase = m.M
																						v144 = m.ExcPending
																						if v144 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v6 + int32(48)
																							return v143
																						}
																					}
																				}
																			}
																		} else {
																			v124 = F_Int64GetDatum(m, v119)
																			mBase = m.M
																			v125 = m.ExcPending
																			if v125 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v124
																				v128 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[6]))
																				if v128 == int64(0) {
																					v131 = int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v131)
																					v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																						v143 = F_HeapTupleHeaderGetDatum(m, v142)
																						mBase = m.M
																						v144 = m.ExcPending
																						if v144 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v6 + int32(48)
																							return v143
																						}
																					}
																				} else {
																					v133 = F_Int64GetDatum(m, v128)
																					mBase = m.M
																					v134 = m.ExcPending
																					if v134 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v133
																						v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return int32(0)
																						} else {
																							v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																							v143 = F_HeapTupleHeaderGetDatum(m, v142)
																							mBase = m.M
																							v144 = m.ExcPending
																							if v144 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v6 + int32(48)
																								return v143
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
															v100 = F_Int64GetDatum(m, v95)
															mBase = m.M
															v101 = m.ExcPending
															if v101 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v100
																v104 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[3]))
																v105 = F_Int64GetDatum(m, v104)
																mBase = m.M
																v106 = m.ExcPending
																if v106 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v105
																	v109 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[4])))
																	if v109 == int32(0) {
																		v112 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)) = uint8(v112)
																		v119 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[5]))
																		if v119 == int64(0) {
																			v122 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v122)
																			v128 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[6]))
																			if v128 == int64(0) {
																				v131 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v131)
																				v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return int32(0)
																				} else {
																					v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																					v143 = F_HeapTupleHeaderGetDatum(m, v142)
																					mBase = m.M
																					v144 = m.ExcPending
																					if v144 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v6 + int32(48)
																						return v143
																					}
																				}
																			} else {
																				v133 = F_Int64GetDatum(m, v128)
																				mBase = m.M
																				v134 = m.ExcPending
																				if v134 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v133
																					v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																						v143 = F_HeapTupleHeaderGetDatum(m, v142)
																						mBase = m.M
																						v144 = m.ExcPending
																						if v144 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v6 + int32(48)
																							return v143
																						}
																					}
																				}
																			}
																		} else {
																			v124 = F_Int64GetDatum(m, v119)
																			mBase = m.M
																			v125 = m.ExcPending
																			if v125 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v124
																				v128 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[6]))
																				if v128 == int64(0) {
																					v131 = int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v131)
																					v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																						v143 = F_HeapTupleHeaderGetDatum(m, v142)
																						mBase = m.M
																						v144 = m.ExcPending
																						if v144 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v6 + int32(48)
																							return v143
																						}
																					}
																				} else {
																					v133 = F_Int64GetDatum(m, v128)
																					mBase = m.M
																					v134 = m.ExcPending
																					if v134 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v133
																						v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return int32(0)
																						} else {
																							v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																							v143 = F_HeapTupleHeaderGetDatum(m, v142)
																							mBase = m.M
																							v144 = m.ExcPending
																							if v144 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v6 + int32(48)
																								return v143
																							}
																						}
																					}
																				}
																			}
																		}
																	} else {
																		v115 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_archiver_7))
																		mBase = m.M
																		v116 = m.ExcPending
																		if v116 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v115
																			v119 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[5]))
																			if v119 == int64(0) {
																				v122 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v122)
																				v128 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[6]))
																				if v128 == int64(0) {
																					v131 = int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v131)
																					v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																						v143 = F_HeapTupleHeaderGetDatum(m, v142)
																						mBase = m.M
																						v144 = m.ExcPending
																						if v144 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v6 + int32(48)
																							return v143
																						}
																					}
																				} else {
																					v133 = F_Int64GetDatum(m, v128)
																					mBase = m.M
																					v134 = m.ExcPending
																					if v134 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v133
																						v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return int32(0)
																						} else {
																							v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																							v143 = F_HeapTupleHeaderGetDatum(m, v142)
																							mBase = m.M
																							v144 = m.ExcPending
																							if v144 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v6 + int32(48)
																								return v143
																							}
																						}
																					}
																				}
																			} else {
																				v124 = F_Int64GetDatum(m, v119)
																				mBase = m.M
																				v125 = m.ExcPending
																				if v125 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v124
																					v128 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[6]))
																					if v128 == int64(0) {
																						v131 = int32(1)
																						*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v131)
																						v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return int32(0)
																						} else {
																							v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																							v143 = F_HeapTupleHeaderGetDatum(m, v142)
																							mBase = m.M
																							v144 = m.ExcPending
																							if v144 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v6 + int32(48)
																								return v143
																							}
																						}
																					} else {
																						v133 = F_Int64GetDatum(m, v128)
																						mBase = m.M
																						v134 = m.ExcPending
																						if v134 != 0 {
																							return int32(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v133
																							v140 = F_heap_form_tuple(m, v21, v6+int32(16), v6+int32(8))
																							mBase = m.M
																							v141 = m.ExcPending
																							if v141 != 0 {
																								return int32(0)
																							} else {
																								v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
																								v143 = F_HeapTupleHeaderGetDatum(m, v142)
																								mBase = m.M
																								v144 = m.ExcPending
																								if v144 != 0 {
																									return int32(0)
																								} else {
																									m.G0 = v6 + int32(48)
																									return v143
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
func F_pg_stat_get_backend_activity(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pgstat_get_beentry_by_proc_number(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			v27 = int32(_a_F_pg_stat_get_backend_activity_0)
			v28 = F_pgstat_clip_activity(m, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = F_cstring_to_text(m, v28)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v28)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						return v30
					}
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_activity[0]))
			v14 = F_has_privs_of_role(m, v12, int32(3375))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 != 0 {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v5)+216))
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
					if v24 != 0 {
						v25 = v22
					} else {
						v25 = int32(_a_F_pg_stat_get_backend_activity_1)
					}
					v27 = v25
					v28 = F_pgstat_clip_activity(m, v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = F_cstring_to_text(m, v28)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v28)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								return v30
							}
						}
					}
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_activity[0]))
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)+52))
					v19 = F_has_privs_of_role(m, v17, v18)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						if v19 != 0 {
							v22 = *(*int32)(unsafe.Add(mBase, uint32(v5)+216))
							v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
							if v24 != 0 {
								v25 = v22
							} else {
								v25 = int32(_a_F_pg_stat_get_backend_activity_1)
							}
							v27 = v25
						} else {
							v27 = int32(_a_F_pg_stat_get_backend_activity_2)
						}
						v28 = F_pgstat_clip_activity(m, v27)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v30 = F_cstring_to_text(m, v28)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v28)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									return v30
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_stat_get_backend_client_port(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v20 int32
	_ = v20
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
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
	var v172 int32
	_ = v172
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pgstat_get_beentry_by_proc_number(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return v172
L2:
	;
	return int32(0)
L3:
	;
	if v14 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
	v172 = v2
	goto L1
L5:
	;
	goto L6
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_client_port[0]))
	v25 = F_has_privs_of_role(m, v23, int32(3375))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v35 = v14 + int32(56)
	v39 = (int32(-56) - v14) & int32(3)
	if v39 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	if v25 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_client_port[0]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v30 = F_has_privs_of_role(m, v28, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	if v30 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v32 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
	v172 = v2
	goto L1
L12:
	;
	v168 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v168)
	v172 = int32(0)
	goto L1
L13:
	;
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
	switch v146 - int32(1) {
	case 0:
		v172 = int32(-1)
		goto L1
	case 1, 9:
		goto L42
	default:
		goto L43
	}
L14:
	;
	v52 = v39 + v35
	v56 = (v14 + int32(188)) & int32(-4)
	v58 = v56 - int32(28)
	if base.Ui32(v52) < base.Ui32(v58) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v42 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v39 == int32(1) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+57)))
	if v45 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if v39 == int32(2) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+58)))
	if v48|base.B2i32(v39 != int32(3)) != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L14
L21:
	;
	v61 = v52
	v62 = v39
	goto L24
L22:
	;
	v89 = v39
	goto L23
L23:
	;
	v95 = v89 + v35
	if base.Ui32(v95) < base.Ui32(v56) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v68|(v69|(v70|(v71|(v72|(v73|(v74|v75)))))) != 0 {
		goto L13
	} else {
		goto L26
	}
L25:
	;
	v89 = v84
	goto L23
L26:
	;
	v84 = v62 + int32(32)
	v85 = v35 + v84
	if base.Ui32(v85) < base.Ui32(v58) {
		v61 = v85
		v62 = v84
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v98 = v95
	v99 = v89
	goto L31
L29:
	;
	v112 = v89
	goto L30
L30:
	;
	v118 = int32(132)
	if base.Ui32(v112) <= base.Ui32(v118) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v105 != 0 {
		goto L13
	} else {
		goto L33
	}
L32:
	;
	v112 = v107
	goto L30
L33:
	;
	v107 = v99 + int32(4)
	v108 = v35 + v107
	if base.Ui32(v108) < base.Ui32(v56) {
		v98 = v108
		v99 = v107
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v121 = v118
	goto L37
L36:
	;
	v121 = v112
	goto L37
L37:
	;
	v124 = v112
	goto L38
L38:
	;
	if v121 == v124 {
		goto L12
	} else {
		goto L40
	}
L39:
	;
	goto L13
L40:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+v35))))
	if v134 == int32(0) {
		v124 = v124 + int32(1)
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v152 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v152)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v14)+184))
	v160 = F_pg_getnameinfo_all(m, v35, v155, v152, v152, v11, int32(32), int32(3))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L44
	}
L43:
	;
	v149 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v149)
	v172 = int32(0)
	goto L1
L44:
	;
	if v160 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v162 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v162)
	v172 = v152
	goto L1
L46:
	;
	goto L47
L47:
	;
	v166 = F_DirectFunctionCall1Coll(m, int32(1408), int32(0), v11)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	v172 = v166
	goto L1
}
func F_pg_stat_get_backend_pid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pgstat_get_beentry_by_proc_number(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
			return v14
		}
	}
}
func F_pg_stat_get_backend_wait_event_type(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pgstat_get_beentry_by_proc_number(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v95 = F_cstring_to_text(m, v93)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L38
	}
L2:
	;
	return int32(0)
L3:
	;
	if v5 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v93 = int32(_a_F_pg_stat_get_backend_wait_event_type_0)
	goto L1
L5:
	;
	goto L6
L6:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_wait_event_type[0]))
	v15 = F_has_privs_of_role(m, v13, int32(3375))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v27 = F_BackendPidGetProc(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L14
	}
L8:
	;
	if v15 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_wait_event_type[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v5)+52))
	v20 = F_has_privs_of_role(m, v18, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	if v20 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v23 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_backend_wait_event_type_1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	return v23
L13:
	;
	v89 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
	return int32(0)
L14:
	;
	if v27 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v32 = int32(0)
	if v31 == v32 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v70 = v27
	goto L17
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+548))
	if v71 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L18:
	;
	if v67 == int32(0) {
		goto L13
	} else {
		goto L29
	}
L19:
	;
	v67 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_wait_event_type[1]))
	v44 = v32
	goto L24
L22:
	;
	v67 = v61
	goto L18
L23:
	;
	v61 = v51 + int32(640)
	goto L22
L24:
	;
	v47 = v44 * int32(640)
	v48 = v40 + v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+44))
	if v49 == v31 {
		v61 = v48
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v67 = int32(0)
	goto L18
L26:
	;
	v51 = v40 + v47
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+684))
	if v52 == v31 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v55 = v44 + int32(2)
	if v55 != int32(38) {
		v44 = v55
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v70 = v67
	goto L17
L30:
	;
	if v86 != 0 {
		v93 = v86
		goto L1
	} else {
		goto L37
	}
L31:
	;
	v86 = int32(0)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v76 = v71 - int32(16777216)
	if base.Ui32(int32(184549375)) < base.Ui32(v76) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v86 = int32(_a_F_pg_stat_get_backend_wait_event_type_2)
	goto L30
L35:
	;
	goto L36
L36:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v76)>>(uint(int32(22))%32))&int32(1020))+uint32(_c_F_pg_stat_get_backend_wait_event_type[2])))
	v86 = v84
	goto L30
L37:
	;
	goto L13
L38:
	;
	return v95
}
func F_pg_stat_get_checkpointer_num_timed(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = F_pgstat_fetch_stat_checkpointer(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
		v7 = F_Int64GetDatum(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_pg_stat_get_db_checksum_last_failure(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_db_checksum_last_failure[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+252))
	if base.B2i32(v7 != int32(0)) == int32(0) {
		v12 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
		return int32(0)
	} else {
		v16 = F_pgstat_fetch_stat_dbentry(m, v4)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v16 != 0 {
				v20 = *(*int64)(unsafe.Add(mBase, uint32(v16)+160))
				if v20 != int64(0) {
					v28 = F_Int64GetDatum(m, v20)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						return v28
					}
				} else {
					v24 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
					return int32(0)
				}
			} else {
				v24 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
				return int32(0)
			}
		}
	}
}
func F_pg_stat_get_db_numbackends(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pgstat_fetch_stat_numbackends(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if int32(0) < v7 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v13 = int32(1)
	v14 = v2
	goto L6
L4:
	;
	v26 = v2
	goto L5
L5:
	;
	return v26
L6:
	;
	v17 = F_pgstat_get_local_beentry_by_index(m, v13)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v26 = v21
	goto L5
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	v21 = v14 + base.B2i32(v19 == v5)
	v23 = v13 + int32(1)
	if v23 <= v7 {
		v13 = v23
		v14 = v21
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
}
func F_pg_stat_get_db_tuples_fetched(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+40))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_db_tuples_updated(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+56))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_db_xact_rollback(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+8))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_function_calls(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pgstat_fetch_stat_funcentry(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int32(0)
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
			v15 = F_Int64GetDatum(m, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_pg_stat_get_live_tuples(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+72))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_mod_since_analyze(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+88))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_numscans(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_tuples_inserted(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+32))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_xact_tuples_deleted(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_find_tabstat_entry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+56))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_xact_tuples_hot_updated(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_find_tabstat_entry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+64))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_get_xact_tuples_newpage_updated(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_find_tabstat_entry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Int64GetDatum(m, int64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+72))
			v14 = F_Int64GetDatum(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_stat_reset_shared(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
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
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
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
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int64
	_ = v200
	var v203 int32
	_ = v203
	var v204 int64
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v7 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L6
	} else {
		goto L85
	}
L2:
	;
	m.G0 = v5 + int32(16)
	return int32(0)
L3:
	;
	F_pgstat_reset_of_kind(m, int32(7))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v59 = F_pg_detoast_datum_packed(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L14
	}
L6:
	;
	return int32(0)
L7:
	;
	F_pgstat_reset_of_kind(m, int32(8))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_pgstat_reset_of_kind(m, int32(9))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	F_pgstat_reset_of_kind(m, int32(10))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v24 = int32(_a_F_pg_stat_reset_shared_0)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	v26 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v26
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+16)) = v30
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+24)) = v30
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+32)) = v30
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+40)) = v30
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v49)+48)) = v30
	goto L11
L11:
	;
	F_pgstat_reset_of_kind(m, int32(11))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	F_pgstat_reset_of_kind(m, int32(12))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	goto L2
L14:
	;
	v61 = F_text_to_cstring(m, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v63 = int32(_a_F_pg_stat_reset_shared_1)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[1])))
	if base.B2i32(v66 == int32(0))|base.B2i32(v66 != v69) != 0 {
		v87 = v66
		v88 = v69
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v87-v88 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	goto L16
L18:
	;
	v72 = v61
	v73 = v63
	goto L19
L19:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	if v77 == int32(0) {
		v87 = v77
		v88 = v76
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v87 = v77
	v88 = v76
	goto L17
L21:
	;
	v80 = int32(1)
	if v77 == v76 {
		v72 = v72 + v80
		v73 = v73 + v80
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	F_pgstat_reset_of_kind(m, int32(7))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L6
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v95 = int32(_a_F_pg_stat_reset_shared_2)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[2])))
	if base.B2i32(v98 == int32(0))|base.B2i32(v98 != v101) != 0 {
		v119 = v98
		v120 = v101
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L2
L27:
	;
	if v119-v120 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	goto L27
L29:
	;
	v104 = v61
	v105 = v95
	goto L30
L30:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+1)))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
	if v109 == int32(0) {
		v119 = v109
		v120 = v108
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v119 = v109
	v120 = v108
	goto L28
L32:
	;
	v112 = int32(1)
	if v109 == v108 {
		v104 = v104 + v112
		v105 = v105 + v112
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	F_pgstat_reset_of_kind(m, int32(8))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v127 = int32(_a_F_pg_stat_reset_shared_3)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[3])))
	if base.B2i32(v130 == int32(0))|base.B2i32(v130 != v133) != 0 {
		v151 = v130
		v152 = v133
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L2
L38:
	;
	if v151-v152 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L39:
	;
	goto L38
L40:
	;
	v136 = v61
	v137 = v127
	goto L41
L41:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+1)))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	if v141 == int32(0) {
		v151 = v141
		v152 = v140
		goto L39
	} else {
		goto L43
	}
L42:
	;
	v151 = v141
	v152 = v140
	goto L39
L43:
	;
	v144 = int32(1)
	if v141 == v140 {
		v136 = v136 + v144
		v137 = v137 + v144
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	F_pgstat_reset_of_kind(m, int32(9))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L6
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v159 != int32(105) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L2
L49:
	;
	v169 = int32(_a_F_pg_stat_reset_shared_4)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[4])))
	if base.B2i32(v172 == int32(0))|base.B2i32(v172 != v175) != 0 {
		v193 = v172
		v194 = v175
		goto L55
	} else {
		goto L56
	}
L50:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v162 != int32(111) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+2)))
	if v165 != 0 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	F_pgstat_reset_of_kind(m, int32(10))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	goto L2
L54:
	;
	if v193-v194 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L55:
	;
	goto L54
L56:
	;
	v178 = v61
	v179 = v169
	goto L57
L57:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+1)))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
	if v183 == int32(0) {
		v193 = v183
		v194 = v182
		goto L55
	} else {
		goto L59
	}
L58:
	;
	v193 = v183
	v194 = v182
	goto L55
L59:
	;
	v186 = int32(1)
	if v183 == v182 {
		v178 = v178 + v186
		v179 = v179 + v186
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v198 = int32(_a_F_pg_stat_reset_shared_0)
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	v200 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v199))) = v200
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	v204 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v203)+8)) = v204
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v207)+16)) = v204
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v211)+24)) = v204
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v215)+32)) = v204
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v219)+40)) = v204
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v223)+48)) = v204
	goto L64
L62:
	;
	goto L63
L63:
	;
	v226 = int32(_a_F_pg_stat_reset_shared_5)
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v232 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[5])))
	if base.B2i32(v229 == int32(0))|base.B2i32(v229 != v232) != 0 {
		v250 = v229
		v251 = v232
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L2
L65:
	;
	if v250-v251 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L66:
	;
	goto L65
L67:
	;
	v235 = v61
	v236 = v226
	goto L68
L68:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+1)))
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+1)))
	if v240 == int32(0) {
		v250 = v240
		v251 = v239
		goto L66
	} else {
		goto L70
	}
L69:
	;
	v250 = v240
	v251 = v239
	goto L66
L70:
	;
	v243 = int32(1)
	if v240 == v239 {
		v235 = v235 + v243
		v236 = v236 + v243
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	F_pgstat_reset_of_kind(m, int32(11))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L6
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v258 = int32(_a_F_pg_stat_reset_shared_6)
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[6])))
	if base.B2i32(v261 == int32(0))|base.B2i32(v261 != v264) != 0 {
		v282 = v261
		v283 = v264
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L2
L76:
	;
	if v282-v283 != 0 {
		goto L1
	} else {
		goto L83
	}
L77:
	;
	goto L76
L78:
	;
	v267 = v61
	v268 = v258
	goto L79
L79:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+1)))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267)+1)))
	if v272 == int32(0) {
		v282 = v272
		v283 = v271
		goto L77
	} else {
		goto L81
	}
L80:
	;
	v282 = v272
	v283 = v271
	goto L77
L81:
	;
	v275 = int32(1)
	if v272 == v271 {
		v267 = v267 + v275
		v268 = v268 + v275
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	F_pgstat_reset_of_kind(m, int32(12))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	goto L2
L85:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v61
	F_errmsg(m, int32(_a_F_pg_stat_reset_shared_7), v5)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	F_errhint(m, int32(_a_F_pg_stat_reset_shared_8), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_pg_stat_reset_shared_9), int32(1922), int32(_a_F_pg_stat_reset_shared_10))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_stat_reset_single_function_counters(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_single_function_counters[0]))
	v5 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+20)))
	F_pgstat_reset(m, int32(3), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pg_stat_statements_1_8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pg_stat_statements_internal(m, l0, int32(4), base.B2i32(v3 != int32(0)))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pg_stat_statements_info(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
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
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v2)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_info[0]))
	if v14 == v2 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v73 = m.ExcPending
		if v73 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_stat_statements_info_0), int32(0))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_stat_statements_info_1), int32(2029), int32(_a_F_pg_stat_statements_info_2))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
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
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_info[1]))
		if v18 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_stat_statements_info_0), int32(0))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_stat_statements_info_1), int32(2029), int32(_a_F_pg_stat_statements_info_2))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
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
			v24 = F_get_call_result_type(m, l0, int32(0), v7+int32(28))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v24 != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_pg_stat_statements_info_3), int32(0))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_stat_statements_info_1), int32(2033), int32(_a_F_pg_stat_statements_info_2))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_info[0]))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(1)
					if v32 != 0 {
						v36 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_info[0]))
						F_s_lock(m, v36+int32(20), int32(_a_F_pg_stat_statements_info_1), int32(2036), int32(_a_F_pg_stat_statements_info_2))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_info[0]))
							*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = int32(0)
							v48 = *(*int64)(unsafe.Add(mBase, uint32(v45)+48))
							v49 = *(*int64)(unsafe.Add(mBase, uint32(v45)+40))
							v50 = F_Int64GetDatum(m, v49)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v50
								v53 = F_Int64GetDatum(m, v48)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v53
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
									v61 = F_heap_form_tuple(m, v56, v7+int32(16), v7+int32(14))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
										v64 = F_HeapTupleHeaderGetDatum(m, v63)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											m.G0 = v7 + int32(32)
											return v64
										}
									}
								}
							}
						}
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_info[0]))
						*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = int32(0)
						v48 = *(*int64)(unsafe.Add(mBase, uint32(v45)+48))
						v49 = *(*int64)(unsafe.Add(mBase, uint32(v45)+40))
						v50 = F_Int64GetDatum(m, v49)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v50
							v53 = F_Int64GetDatum(m, v48)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v53
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
								v61 = F_heap_form_tuple(m, v56, v7+int32(16), v7+int32(14))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
									v64 = F_HeapTupleHeaderGetDatum(m, v63)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										m.G0 = v7 + int32(32)
										return v64
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
