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
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int64
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	v2 = int32(0)
	v4 = int64(0)
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v4
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v4
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v7)+11)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v2
	v22 = F_CreateTemplateTupleDesc(m, int32(7))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		F_TupleDescInitEntry(m, v22, int32(1), int32(84601), int32(20), int32(-1), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			F_TupleDescInitEntry(m, v22, int32(2), int32(297533), int32(25), int32(-1), int32(0))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				F_TupleDescInitEntry(m, v22, int32(3), int32(362444), int32(1184), int32(-1), int32(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					F_TupleDescInitEntry(m, v22, int32(4), int32(84616), int32(20), int32(-1), int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						F_TupleDescInitEntry(m, v22, int32(5), int32(297551), int32(25), int32(-1), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_TupleDescInitEntry(m, v22, int32(6), int32(362463), int32(1184), int32(-1), int32(0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								F_TupleDescInitEntry(m, v22, int32(7), int32(100351), int32(1184), int32(-1), int32(0))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									v75 = F_BlessTupleDesc(m, v22)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										F_pgstat_snapshot_fixed(m, int32(7))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											v80 = int32(4379152)
											v81 = *(*int64)(unsafe.Add(mBase, _consts[821]))
											v82 = F_Int64GetDatum(m, v81)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v82
												v85 = int32(*(*uint8)(unsafe.Add(mBase, _consts[822])))
												if v85 == int32(0) {
													v88 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v7)+9)) = uint8(v88)
													v95 = *(*int64)(unsafe.Add(mBase, _consts[823]))
													if v95 == int64(0) {
														v98 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v7)+10)) = uint8(v98)
														v103 = *(*int64)(unsafe.Add(mBase, _consts[824]))
														v104 = F_Int64GetDatum(m, v103)
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v104
															v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[825])))
															if v107 == int32(0) {
																v110 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v110)
																v117 = *(*int64)(unsafe.Add(mBase, _consts[826]))
																if v117 == int64(0) {
																	v120 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v120)
																	v125 = *(*int64)(unsafe.Add(mBase, _consts[827]))
																	if v125 == int64(0) {
																		v128 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v128)
																		v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																		mBase = m.M
																		v138 = m.ExcPending
																		if v138 != 0 {
																			return int32(0)
																		} else {
																			v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																			v140 = F_HeapTupleHeaderGetDatum(m, v139)
																			mBase = m.M
																			v141 = m.ExcPending
																			if v141 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v7 + int32(48)
																				return v140
																			}
																		}
																	} else {
																		v130 = F_Int64GetDatum(m, v125)
																		mBase = m.M
																		v131 = m.ExcPending
																		if v131 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v130
																			v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																			mBase = m.M
																			v138 = m.ExcPending
																			if v138 != 0 {
																				return int32(0)
																			} else {
																				v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																				v140 = F_HeapTupleHeaderGetDatum(m, v139)
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v7 + int32(48)
																					return v140
																				}
																			}
																		}
																	}
																} else {
																	v122 = F_Int64GetDatum(m, v117)
																	mBase = m.M
																	v123 = m.ExcPending
																	if v123 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v122
																		v125 = *(*int64)(unsafe.Add(mBase, _consts[827]))
																		if v125 == int64(0) {
																			v128 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v128)
																			v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																			mBase = m.M
																			v138 = m.ExcPending
																			if v138 != 0 {
																				return int32(0)
																			} else {
																				v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																				v140 = F_HeapTupleHeaderGetDatum(m, v139)
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v7 + int32(48)
																					return v140
																				}
																			}
																		} else {
																			v130 = F_Int64GetDatum(m, v125)
																			mBase = m.M
																			v131 = m.ExcPending
																			if v131 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v130
																				v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																				mBase = m.M
																				v138 = m.ExcPending
																				if v138 != 0 {
																					return int32(0)
																				} else {
																					v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																					v140 = F_HeapTupleHeaderGetDatum(m, v139)
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v7 + int32(48)
																						return v140
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																v114 = F_cstring_to_text(m, int32(4379224))
																mBase = m.M
																v115 = m.ExcPending
																if v115 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v114
																	v117 = *(*int64)(unsafe.Add(mBase, _consts[826]))
																	if v117 == int64(0) {
																		v120 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v120)
																		v125 = *(*int64)(unsafe.Add(mBase, _consts[827]))
																		if v125 == int64(0) {
																			v128 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v128)
																			v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																			mBase = m.M
																			v138 = m.ExcPending
																			if v138 != 0 {
																				return int32(0)
																			} else {
																				v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																				v140 = F_HeapTupleHeaderGetDatum(m, v139)
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v7 + int32(48)
																					return v140
																				}
																			}
																		} else {
																			v130 = F_Int64GetDatum(m, v125)
																			mBase = m.M
																			v131 = m.ExcPending
																			if v131 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v130
																				v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																				mBase = m.M
																				v138 = m.ExcPending
																				if v138 != 0 {
																					return int32(0)
																				} else {
																					v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																					v140 = F_HeapTupleHeaderGetDatum(m, v139)
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v7 + int32(48)
																						return v140
																					}
																				}
																			}
																		}
																	} else {
																		v122 = F_Int64GetDatum(m, v117)
																		mBase = m.M
																		v123 = m.ExcPending
																		if v123 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v122
																			v125 = *(*int64)(unsafe.Add(mBase, _consts[827]))
																			if v125 == int64(0) {
																				v128 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v128)
																				v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																				mBase = m.M
																				v138 = m.ExcPending
																				if v138 != 0 {
																					return int32(0)
																				} else {
																					v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																					v140 = F_HeapTupleHeaderGetDatum(m, v139)
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v7 + int32(48)
																						return v140
																					}
																				}
																			} else {
																				v130 = F_Int64GetDatum(m, v125)
																				mBase = m.M
																				v131 = m.ExcPending
																				if v131 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v130
																					v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																					mBase = m.M
																					v138 = m.ExcPending
																					if v138 != 0 {
																						return int32(0)
																					} else {
																						v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																						v140 = F_HeapTupleHeaderGetDatum(m, v139)
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v7 + int32(48)
																							return v140
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
															*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v100
															v103 = *(*int64)(unsafe.Add(mBase, _consts[824]))
															v104 = F_Int64GetDatum(m, v103)
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v104
																v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[825])))
																if v107 == int32(0) {
																	v110 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v110)
																	v117 = *(*int64)(unsafe.Add(mBase, _consts[826]))
																	if v117 == int64(0) {
																		v120 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v120)
																		v125 = *(*int64)(unsafe.Add(mBase, _consts[827]))
																		if v125 == int64(0) {
																			v128 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v128)
																			v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																			mBase = m.M
																			v138 = m.ExcPending
																			if v138 != 0 {
																				return int32(0)
																			} else {
																				v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																				v140 = F_HeapTupleHeaderGetDatum(m, v139)
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v7 + int32(48)
																					return v140
																				}
																			}
																		} else {
																			v130 = F_Int64GetDatum(m, v125)
																			mBase = m.M
																			v131 = m.ExcPending
																			if v131 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v130
																				v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																				mBase = m.M
																				v138 = m.ExcPending
																				if v138 != 0 {
																					return int32(0)
																				} else {
																					v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																					v140 = F_HeapTupleHeaderGetDatum(m, v139)
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v7 + int32(48)
																						return v140
																					}
																				}
																			}
																		}
																	} else {
																		v122 = F_Int64GetDatum(m, v117)
																		mBase = m.M
																		v123 = m.ExcPending
																		if v123 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v122
																			v125 = *(*int64)(unsafe.Add(mBase, _consts[827]))
																			if v125 == int64(0) {
																				v128 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v128)
																				v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																				mBase = m.M
																				v138 = m.ExcPending
																				if v138 != 0 {
																					return int32(0)
																				} else {
																					v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																					v140 = F_HeapTupleHeaderGetDatum(m, v139)
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v7 + int32(48)
																						return v140
																					}
																				}
																			} else {
																				v130 = F_Int64GetDatum(m, v125)
																				mBase = m.M
																				v131 = m.ExcPending
																				if v131 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v130
																					v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																					mBase = m.M
																					v138 = m.ExcPending
																					if v138 != 0 {
																						return int32(0)
																					} else {
																						v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																						v140 = F_HeapTupleHeaderGetDatum(m, v139)
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v7 + int32(48)
																							return v140
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	v114 = F_cstring_to_text(m, int32(4379224))
																	mBase = m.M
																	v115 = m.ExcPending
																	if v115 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v114
																		v117 = *(*int64)(unsafe.Add(mBase, _consts[826]))
																		if v117 == int64(0) {
																			v120 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v120)
																			v125 = *(*int64)(unsafe.Add(mBase, _consts[827]))
																			if v125 == int64(0) {
																				v128 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v128)
																				v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																				mBase = m.M
																				v138 = m.ExcPending
																				if v138 != 0 {
																					return int32(0)
																				} else {
																					v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																					v140 = F_HeapTupleHeaderGetDatum(m, v139)
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v7 + int32(48)
																						return v140
																					}
																				}
																			} else {
																				v130 = F_Int64GetDatum(m, v125)
																				mBase = m.M
																				v131 = m.ExcPending
																				if v131 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v130
																					v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																					mBase = m.M
																					v138 = m.ExcPending
																					if v138 != 0 {
																						return int32(0)
																					} else {
																						v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																						v140 = F_HeapTupleHeaderGetDatum(m, v139)
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v7 + int32(48)
																							return v140
																						}
																					}
																				}
																			}
																		} else {
																			v122 = F_Int64GetDatum(m, v117)
																			mBase = m.M
																			v123 = m.ExcPending
																			if v123 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v122
																				v125 = *(*int64)(unsafe.Add(mBase, _consts[827]))
																				if v125 == int64(0) {
																					v128 = int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v128)
																					v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																					mBase = m.M
																					v138 = m.ExcPending
																					if v138 != 0 {
																						return int32(0)
																					} else {
																						v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																						v140 = F_HeapTupleHeaderGetDatum(m, v139)
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v7 + int32(48)
																							return v140
																						}
																					}
																				} else {
																					v130 = F_Int64GetDatum(m, v125)
																					mBase = m.M
																					v131 = m.ExcPending
																					if v131 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v130
																						v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																						mBase = m.M
																						v138 = m.ExcPending
																						if v138 != 0 {
																							return int32(0)
																						} else {
																							v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																							v140 = F_HeapTupleHeaderGetDatum(m, v139)
																							mBase = m.M
																							v141 = m.ExcPending
																							if v141 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v7 + int32(48)
																								return v140
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
													v92 = F_cstring_to_text(m, int32(4379160))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v92
														v95 = *(*int64)(unsafe.Add(mBase, _consts[823]))
														if v95 == int64(0) {
															v98 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v7)+10)) = uint8(v98)
															v103 = *(*int64)(unsafe.Add(mBase, _consts[824]))
															v104 = F_Int64GetDatum(m, v103)
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v104
																v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[825])))
																if v107 == int32(0) {
																	v110 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v110)
																	v117 = *(*int64)(unsafe.Add(mBase, _consts[826]))
																	if v117 == int64(0) {
																		v120 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v120)
																		v125 = *(*int64)(unsafe.Add(mBase, _consts[827]))
																		if v125 == int64(0) {
																			v128 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v128)
																			v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																			mBase = m.M
																			v138 = m.ExcPending
																			if v138 != 0 {
																				return int32(0)
																			} else {
																				v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																				v140 = F_HeapTupleHeaderGetDatum(m, v139)
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v7 + int32(48)
																					return v140
																				}
																			}
																		} else {
																			v130 = F_Int64GetDatum(m, v125)
																			mBase = m.M
																			v131 = m.ExcPending
																			if v131 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v130
																				v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																				mBase = m.M
																				v138 = m.ExcPending
																				if v138 != 0 {
																					return int32(0)
																				} else {
																					v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																					v140 = F_HeapTupleHeaderGetDatum(m, v139)
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v7 + int32(48)
																						return v140
																					}
																				}
																			}
																		}
																	} else {
																		v122 = F_Int64GetDatum(m, v117)
																		mBase = m.M
																		v123 = m.ExcPending
																		if v123 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v122
																			v125 = *(*int64)(unsafe.Add(mBase, _consts[827]))
																			if v125 == int64(0) {
																				v128 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v128)
																				v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																				mBase = m.M
																				v138 = m.ExcPending
																				if v138 != 0 {
																					return int32(0)
																				} else {
																					v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																					v140 = F_HeapTupleHeaderGetDatum(m, v139)
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v7 + int32(48)
																						return v140
																					}
																				}
																			} else {
																				v130 = F_Int64GetDatum(m, v125)
																				mBase = m.M
																				v131 = m.ExcPending
																				if v131 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v130
																					v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																					mBase = m.M
																					v138 = m.ExcPending
																					if v138 != 0 {
																						return int32(0)
																					} else {
																						v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																						v140 = F_HeapTupleHeaderGetDatum(m, v139)
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v7 + int32(48)
																							return v140
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	v114 = F_cstring_to_text(m, int32(4379224))
																	mBase = m.M
																	v115 = m.ExcPending
																	if v115 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v114
																		v117 = *(*int64)(unsafe.Add(mBase, _consts[826]))
																		if v117 == int64(0) {
																			v120 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v120)
																			v125 = *(*int64)(unsafe.Add(mBase, _consts[827]))
																			if v125 == int64(0) {
																				v128 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v128)
																				v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																				mBase = m.M
																				v138 = m.ExcPending
																				if v138 != 0 {
																					return int32(0)
																				} else {
																					v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																					v140 = F_HeapTupleHeaderGetDatum(m, v139)
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v7 + int32(48)
																						return v140
																					}
																				}
																			} else {
																				v130 = F_Int64GetDatum(m, v125)
																				mBase = m.M
																				v131 = m.ExcPending
																				if v131 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v130
																					v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																					mBase = m.M
																					v138 = m.ExcPending
																					if v138 != 0 {
																						return int32(0)
																					} else {
																						v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																						v140 = F_HeapTupleHeaderGetDatum(m, v139)
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v7 + int32(48)
																							return v140
																						}
																					}
																				}
																			}
																		} else {
																			v122 = F_Int64GetDatum(m, v117)
																			mBase = m.M
																			v123 = m.ExcPending
																			if v123 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v122
																				v125 = *(*int64)(unsafe.Add(mBase, _consts[827]))
																				if v125 == int64(0) {
																					v128 = int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v128)
																					v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																					mBase = m.M
																					v138 = m.ExcPending
																					if v138 != 0 {
																						return int32(0)
																					} else {
																						v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																						v140 = F_HeapTupleHeaderGetDatum(m, v139)
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v7 + int32(48)
																							return v140
																						}
																					}
																				} else {
																					v130 = F_Int64GetDatum(m, v125)
																					mBase = m.M
																					v131 = m.ExcPending
																					if v131 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v130
																						v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																						mBase = m.M
																						v138 = m.ExcPending
																						if v138 != 0 {
																							return int32(0)
																						} else {
																							v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																							v140 = F_HeapTupleHeaderGetDatum(m, v139)
																							mBase = m.M
																							v141 = m.ExcPending
																							if v141 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v7 + int32(48)
																								return v140
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
																*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v100
																v103 = *(*int64)(unsafe.Add(mBase, _consts[824]))
																v104 = F_Int64GetDatum(m, v103)
																mBase = m.M
																v105 = m.ExcPending
																if v105 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v104
																	v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[825])))
																	if v107 == int32(0) {
																		v110 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v110)
																		v117 = *(*int64)(unsafe.Add(mBase, _consts[826]))
																		if v117 == int64(0) {
																			v120 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v120)
																			v125 = *(*int64)(unsafe.Add(mBase, _consts[827]))
																			if v125 == int64(0) {
																				v128 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v128)
																				v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																				mBase = m.M
																				v138 = m.ExcPending
																				if v138 != 0 {
																					return int32(0)
																				} else {
																					v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																					v140 = F_HeapTupleHeaderGetDatum(m, v139)
																					mBase = m.M
																					v141 = m.ExcPending
																					if v141 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v7 + int32(48)
																						return v140
																					}
																				}
																			} else {
																				v130 = F_Int64GetDatum(m, v125)
																				mBase = m.M
																				v131 = m.ExcPending
																				if v131 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v130
																					v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																					mBase = m.M
																					v138 = m.ExcPending
																					if v138 != 0 {
																						return int32(0)
																					} else {
																						v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																						v140 = F_HeapTupleHeaderGetDatum(m, v139)
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v7 + int32(48)
																							return v140
																						}
																					}
																				}
																			}
																		} else {
																			v122 = F_Int64GetDatum(m, v117)
																			mBase = m.M
																			v123 = m.ExcPending
																			if v123 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v122
																				v125 = *(*int64)(unsafe.Add(mBase, _consts[827]))
																				if v125 == int64(0) {
																					v128 = int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v128)
																					v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																					mBase = m.M
																					v138 = m.ExcPending
																					if v138 != 0 {
																						return int32(0)
																					} else {
																						v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																						v140 = F_HeapTupleHeaderGetDatum(m, v139)
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v7 + int32(48)
																							return v140
																						}
																					}
																				} else {
																					v130 = F_Int64GetDatum(m, v125)
																					mBase = m.M
																					v131 = m.ExcPending
																					if v131 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v130
																						v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																						mBase = m.M
																						v138 = m.ExcPending
																						if v138 != 0 {
																							return int32(0)
																						} else {
																							v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																							v140 = F_HeapTupleHeaderGetDatum(m, v139)
																							mBase = m.M
																							v141 = m.ExcPending
																							if v141 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v7 + int32(48)
																								return v140
																							}
																						}
																					}
																				}
																			}
																		}
																	} else {
																		v114 = F_cstring_to_text(m, int32(4379224))
																		mBase = m.M
																		v115 = m.ExcPending
																		if v115 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v114
																			v117 = *(*int64)(unsafe.Add(mBase, _consts[826]))
																			if v117 == int64(0) {
																				v120 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v120)
																				v125 = *(*int64)(unsafe.Add(mBase, _consts[827]))
																				if v125 == int64(0) {
																					v128 = int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v128)
																					v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																					mBase = m.M
																					v138 = m.ExcPending
																					if v138 != 0 {
																						return int32(0)
																					} else {
																						v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																						v140 = F_HeapTupleHeaderGetDatum(m, v139)
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v7 + int32(48)
																							return v140
																						}
																					}
																				} else {
																					v130 = F_Int64GetDatum(m, v125)
																					mBase = m.M
																					v131 = m.ExcPending
																					if v131 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v130
																						v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																						mBase = m.M
																						v138 = m.ExcPending
																						if v138 != 0 {
																							return int32(0)
																						} else {
																							v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																							v140 = F_HeapTupleHeaderGetDatum(m, v139)
																							mBase = m.M
																							v141 = m.ExcPending
																							if v141 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v7 + int32(48)
																								return v140
																							}
																						}
																					}
																				}
																			} else {
																				v122 = F_Int64GetDatum(m, v117)
																				mBase = m.M
																				v123 = m.ExcPending
																				if v123 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v122
																					v125 = *(*int64)(unsafe.Add(mBase, _consts[827]))
																					if v125 == int64(0) {
																						v128 = int32(1)
																						*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v128)
																						v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																						mBase = m.M
																						v138 = m.ExcPending
																						if v138 != 0 {
																							return int32(0)
																						} else {
																							v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																							v140 = F_HeapTupleHeaderGetDatum(m, v139)
																							mBase = m.M
																							v141 = m.ExcPending
																							if v141 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v7 + int32(48)
																								return v140
																							}
																						}
																					} else {
																						v130 = F_Int64GetDatum(m, v125)
																						mBase = m.M
																						v131 = m.ExcPending
																						if v131 != 0 {
																							return int32(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v130
																							v137 = F_heap_form_tuple(m, v22, v7+int32(16), v7+int32(8))
																							mBase = m.M
																							v138 = m.ExcPending
																							if v138 != 0 {
																								return int32(0)
																							} else {
																								v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
																								v140 = F_HeapTupleHeaderGetDatum(m, v139)
																								mBase = m.M
																								v141 = m.ExcPending
																								if v141 != 0 {
																									return int32(0)
																								} else {
																									m.G0 = v7 + int32(48)
																									return v140
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
			v27 = int32(528466)
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
			v12 = *(*int32)(unsafe.Add(mBase, _consts[4]))
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
						v25 = int32(528537)
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
					v17 = *(*int32)(unsafe.Add(mBase, _consts[4]))
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
								v25 = int32(528537)
							}
							v27 = v25
						} else {
							v27 = int32(528502)
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
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
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
	return v175
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
	v175 = v2
	goto L1
L5:
	;
	goto L6
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[4]))
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
	if v35&int32(3) == int32(0) {
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
	v28 = *(*int32)(unsafe.Add(mBase, _consts[4]))
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
	v175 = v2
	goto L1
L12:
	;
	v171 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v171)
	v175 = int32(0)
	goto L1
L13:
	;
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
	switch v149 - int32(1) {
	case 0:
		v175 = int32(-1)
		goto L1
	case 1, 9:
		goto L40
	default:
		goto L41
	}
L14:
	;
	v63 = v35 + (int32(-56)-v14)&int32(3)
	v65 = v14 + int32(188)
	v67 = v65 & int32(-4)
	v69 = v67 - int32(28)
	if base.Ui32(v63) < base.Ui32(v69) {
		goto L22
	} else {
		goto L23
	}
L15:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+56)))
	if v40 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if (v14+int32(57))&int32(3) == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+57)))
	if v47 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if (v14+int32(58))&int32(3) == int32(0) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+58)))
	if v54 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	if (v14-int32(1))&int32(3) != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L14
L22:
	;
	v72 = v63
	goto L25
L23:
	;
	v98 = v63
	goto L24
L24:
	;
	if base.Ui32(v98) < base.Ui32(v67) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v72)+24))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v79|(v80|(v81|(v82|(v83|(v84|(v85|v86)))))) != 0 {
		goto L13
	} else {
		goto L27
	}
L26:
	;
	v98 = v95
	goto L24
L27:
	;
	v95 = v72 + int32(32)
	if base.Ui32(v95) < base.Ui32(v69) {
		v72 = v95
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v107 = v98
	goto L32
L30:
	;
	v119 = v98
	goto L31
L31:
	;
	v127 = v119
	goto L36
L32:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v114 != 0 {
		goto L13
	} else {
		goto L34
	}
L33:
	;
	v119 = v116
	goto L31
L34:
	;
	v116 = v107 + int32(4)
	if base.Ui32(v116) < base.Ui32(v67) {
		v107 = v116
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	if base.Ui32(v65) <= base.Ui32(v127) {
		goto L12
	} else {
		goto L38
	}
L37:
	;
	goto L13
L38:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v135 == int32(0) {
		v127 = v127 + int32(1)
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v155 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v155)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v14)+184))
	v163 = F_pg_getnameinfo_all(m, v35, v158, v155, v155, v11, int32(32), int32(3))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L42
	}
L41:
	;
	v152 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v152)
	v175 = int32(0)
	goto L1
L42:
	;
	if v163 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v165 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v165)
	v175 = v155
	goto L1
L44:
	;
	goto L45
L45:
	;
	v169 = F_DirectFunctionCall1Coll(m, int32(1423), int32(0), v11)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v175 = v169
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
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
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
	v96 = F_cstring_to_text(m, v94)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L2
	} else {
		goto L37
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
	v94 = int32(528466)
	goto L1
L5:
	;
	goto L6
L6:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[4]))
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
	v18 = *(*int32)(unsafe.Add(mBase, _consts[4]))
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
	v23 = F_cstring_to_text(m, int32(528502))
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
	v90 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v90)
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
	v69 = v27
	goto L17
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+548))
	if v70 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L18:
	;
	if v66 == int32(0) {
		goto L13
	} else {
		goto L28
	}
L19:
	;
	v66 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[693]))
	v41 = v32
	goto L23
L22:
	;
	v66 = v61
	goto L18
L23:
	;
	v46 = v39 + v41*int32(640)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v47 == v31 {
		v61 = v46
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v66 = int32(0)
	goto L18
L25:
	;
	v53 = v39 + (v41|int32(1))*int32(640)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+44))
	if v54 == v31 {
		v61 = v53
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v57 = v41 + int32(2)
	if v57 != int32(38) {
		v41 = v57
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v69 = v66
	goto L17
L29:
	;
	if v87 != 0 {
		v94 = v87
		goto L1
	} else {
		goto L36
	}
L30:
	;
	v87 = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v75 = v70 - int32(16777216)
	if base.Ui32(int32(184549375)) < base.Ui32(v75) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v87 = int32(528115)
	goto L29
L34:
	;
	goto L35
L35:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v75)>>(uint(int32(22))%32))&int32(1020))+uint32(_consts[813])))
	v87 = v85
	goto L29
L36:
	;
	goto L13
L37:
	;
	return v96
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
	v6 = *(*int32)(unsafe.Add(mBase, _consts[178]))
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
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v199 int32
	_ = v199
	var v200 int64
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
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
	v291 = m.ExcPending
	if v291 != 0 {
		goto L6
	} else {
		goto L91
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
	v24 = int32(4351084)
	v25 = *(*int32)(unsafe.Add(mBase, _consts[814]))
	v26 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v26
	v29 = *(*int32)(unsafe.Add(mBase, _consts[814]))
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	v33 = *(*int32)(unsafe.Add(mBase, _consts[814]))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+16)) = v30
	v37 = *(*int32)(unsafe.Add(mBase, _consts[814]))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+24)) = v30
	v41 = *(*int32)(unsafe.Add(mBase, _consts[814]))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+32)) = v30
	v45 = *(*int32)(unsafe.Add(mBase, _consts[814]))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+40)) = v30
	v49 = *(*int32)(unsafe.Add(mBase, _consts[814]))
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
	v63 = int32(207243)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, _consts[815])))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v67 == int32(0) {
		v86 = v66
		v87 = v67
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v87-v86 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	goto L16
L18:
	;
	if v66 != v67 {
		v86 = v66
		v87 = v67
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v71 = v61
	v72 = v63
	goto L20
L20:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	if v76 == int32(0) {
		v86 = v75
		v87 = v76
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v86 = v75
	v87 = v76
	goto L17
L22:
	;
	v79 = int32(1)
	if v75 == v76 {
		v71 = v71 + v79
		v72 = v72 + v79
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	F_pgstat_reset_of_kind(m, int32(7))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v94 = int32(208148)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, _consts[816])))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v98 == int32(0) {
		v117 = v97
		v118 = v98
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L2
L28:
	;
	if v118-v117 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	goto L28
L30:
	;
	if v97 != v98 {
		v117 = v97
		v118 = v98
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v102 = v61
	v103 = v94
	goto L32
L32:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	if v107 == int32(0) {
		v117 = v106
		v118 = v107
		goto L29
	} else {
		goto L34
	}
L33:
	;
	v117 = v106
	v118 = v107
	goto L29
L34:
	;
	v110 = int32(1)
	if v106 == v107 {
		v102 = v102 + v110
		v103 = v103 + v110
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	F_pgstat_reset_of_kind(m, int32(8))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L6
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v125 = int32(207570)
	v128 = int32(*(*uint8)(unsafe.Add(mBase, _consts[817])))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v129 == int32(0) {
		v148 = v128
		v149 = v129
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L2
L40:
	;
	if v149-v148 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L41:
	;
	goto L40
L42:
	;
	if v128 != v129 {
		v148 = v128
		v149 = v129
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v133 = v61
	v134 = v125
	goto L44
L44:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+1)))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+1)))
	if v138 == int32(0) {
		v148 = v137
		v149 = v138
		goto L41
	} else {
		goto L46
	}
L45:
	;
	v148 = v137
	v149 = v138
	goto L41
L46:
	;
	v141 = int32(1)
	if v137 == v138 {
		v133 = v133 + v141
		v134 = v134 + v141
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	F_pgstat_reset_of_kind(m, int32(9))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L6
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v156 != int32(105) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L2
L52:
	;
	v166 = int32(312421)
	v169 = int32(*(*uint8)(unsafe.Add(mBase, _consts[818])))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v170 == int32(0) {
		v189 = v169
		v190 = v170
		goto L58
	} else {
		goto L59
	}
L53:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v159 != int32(111) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+2)))
	if v162 != 0 {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	F_pgstat_reset_of_kind(m, int32(10))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	goto L2
L57:
	;
	if v190-v189 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L58:
	;
	goto L57
L59:
	;
	if v169 != v170 {
		v189 = v169
		v190 = v170
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v174 = v61
	v175 = v166
	goto L61
L61:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
	if v179 == int32(0) {
		v189 = v178
		v190 = v179
		goto L58
	} else {
		goto L63
	}
L62:
	;
	v189 = v178
	v190 = v179
	goto L58
L63:
	;
	v182 = int32(1)
	if v178 == v179 {
		v174 = v174 + v182
		v175 = v175 + v182
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v194 = int32(4351084)
	v195 = *(*int32)(unsafe.Add(mBase, _consts[814]))
	v196 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v195))) = v196
	v199 = *(*int32)(unsafe.Add(mBase, _consts[814]))
	v200 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v199)+8)) = v200
	v203 = *(*int32)(unsafe.Add(mBase, _consts[814]))
	*(*int64)(unsafe.Add(mBase, uint32(v203)+16)) = v200
	v207 = *(*int32)(unsafe.Add(mBase, _consts[814]))
	*(*int64)(unsafe.Add(mBase, uint32(v207)+24)) = v200
	v211 = *(*int32)(unsafe.Add(mBase, _consts[814]))
	*(*int64)(unsafe.Add(mBase, uint32(v211)+32)) = v200
	v215 = *(*int32)(unsafe.Add(mBase, _consts[814]))
	*(*int64)(unsafe.Add(mBase, uint32(v215)+40)) = v200
	v219 = *(*int32)(unsafe.Add(mBase, _consts[814]))
	*(*int64)(unsafe.Add(mBase, uint32(v219)+48)) = v200
	goto L68
L66:
	;
	goto L67
L67:
	;
	v222 = int32(36178)
	v225 = int32(*(*uint8)(unsafe.Add(mBase, _consts[819])))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v226 == int32(0) {
		v245 = v225
		v246 = v226
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L2
L69:
	;
	if v246-v245 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L70:
	;
	goto L69
L71:
	;
	if v225 != v226 {
		v245 = v225
		v246 = v226
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v230 = v61
	v231 = v222
	goto L73
L73:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+1)))
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	if v235 == int32(0) {
		v245 = v234
		v246 = v235
		goto L70
	} else {
		goto L75
	}
L74:
	;
	v245 = v234
	v246 = v235
	goto L70
L75:
	;
	v238 = int32(1)
	if v234 == v235 {
		v230 = v230 + v238
		v231 = v231 + v238
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	F_pgstat_reset_of_kind(m, int32(11))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L6
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v253 = int32(297568)
	v256 = int32(*(*uint8)(unsafe.Add(mBase, _consts[820])))
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v257 == int32(0) {
		v276 = v256
		v277 = v257
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L2
L81:
	;
	if v277-v276 != 0 {
		goto L1
	} else {
		goto L89
	}
L82:
	;
	goto L81
L83:
	;
	if v256 != v257 {
		v276 = v256
		v277 = v257
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v261 = v61
	v262 = v253
	goto L85
L85:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+1)))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+1)))
	if v266 == int32(0) {
		v276 = v265
		v277 = v266
		goto L82
	} else {
		goto L87
	}
L86:
	;
	v276 = v265
	v277 = v266
	goto L82
L87:
	;
	v269 = int32(1)
	if v265 == v266 {
		v261 = v261 + v269
		v262 = v262 + v269
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	F_pgstat_reset_of_kind(m, int32(12))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L6
	} else {
		goto L90
	}
L90:
	;
	goto L2
L91:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L6
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v61
	F_errmsg(m, int32(690213), v5)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	F_errhint(m, int32(633759), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(477856), int32(1922), int32(435987))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[226]))
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
