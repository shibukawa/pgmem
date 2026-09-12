package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_stat_file(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int64
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int64
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	v6 = m.G0
	v8 = v6 - int32(144)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v15 == int32(2) {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v21 = base.B2i32(v18 != int32(0))
		} else {
			v21 = int32(0)
		}
		v22 = F_convert_and_check_filename(m, v11)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v28 = F___fstatat(m, int32(-100), v22, v8+int32(48), int32(0))
			mBase = m.M
			if v28 < int32(0) {
				if v21 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v22
							F_errmsg(m, int32(294607), v8)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(493492), int32(437), int32(382259))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
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
					v34 = *(*int32)(unsafe.Add(mBase, _consts[86]))
					if v34 != int32(44) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v22
								F_errmsg(m, int32(294607), v8)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(493492), int32(437), int32(382259))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
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
						v37 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v37)
						v158 = int32(0)
						m.G0 = v8 + int32(144)
						return v158
					}
				}
			} else {
				v56 = F_CreateTemplateTupleDesc(m, int32(6))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					F_TupleDescInitEntry(m, v56, int32(1), int32(337698), int32(20), int32(-1), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_TupleDescInitEntry(m, v56, int32(2), int32(129033), int32(1184), int32(-1), int32(0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							F_TupleDescInitEntry(m, v56, int32(3), int32(264741), int32(1184), int32(-1), int32(0))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								F_TupleDescInitEntry(m, v56, int32(4), int32(398854), int32(1184), int32(-1), int32(0))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									F_TupleDescInitEntry(m, v56, int32(5), int32(262175), int32(1184), int32(-1), int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										F_TupleDescInitEntry(m, v56, int32(6), int32(211337), int32(16), int32(-1), int32(0))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											v100 = F_BlessTupleDesc(m, v56)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return int32(0)
											} else {
												v103 = v8 + int32(12)
												v104 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(v103))) = uint16(v104)
												*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v104
												v108 = *(*int64)(unsafe.Add(mBase, uint32(v8)+72))
												v109 = F_Int64GetDatum(m, v108)
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v109
													v112 = *(*int64)(unsafe.Add(mBase, uint32(v8)+88))
													v117 = F_Int64GetDatum(m, v112*int64(1000000)-int64(946684800000000))
													mBase = m.M
													v118 = m.ExcPending
													if v118 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v117
														v120 = *(*int64)(unsafe.Add(mBase, uint32(v8)+104))
														v125 = F_Int64GetDatum(m, v120*int64(1000000)-int64(946684800000000))
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v125
															v128 = *(*int64)(unsafe.Add(mBase, uint32(v8)+120))
															v133 = F_Int64GetDatum(m, v128*int64(1000000)-int64(946684800000000))
															mBase = m.M
															v134 = m.ExcPending
															if v134 != 0 {
																return int32(0)
															} else {
																v135 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v135)
																*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v133
																v138 = *(*int32)(unsafe.Add(mBase, uint32(v8)+52))
																*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = base.B2i32(v138&int32(61440) == int32(16384))
																v148 = F_heap_form_tuple(m, v56, v8+int32(16), v8+int32(8))
																mBase = m.M
																v149 = m.ExcPending
																if v149 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v22)
																	mBase = m.M
																	v151 = m.ExcPending
																	if v151 != 0 {
																		return int32(0)
																	} else {
																		v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+16))
																		v153 = F_HeapTupleHeaderGetDatum(m, v152)
																		mBase = m.M
																		v154 = m.ExcPending
																		if v154 != 0 {
																			return int32(0)
																		} else {
																			v158 = v153
																			m.G0 = v8 + int32(144)
																			return v158
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
func F_pg_stat_get_backend_wal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
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
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v27 int64
	_ = v27
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pgstat_fetch_stat_backend_by_pid(m, v10, v2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v37 = v2
			m.G0 = v8 + int32(32)
			return v37
		} else {
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(2912))))
			*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v23
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(2904))))
			*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v27
			v31 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(2896))))
			*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v31
			v33 = *(*int64)(unsafe.Add(mBase, uint32(v12)+2888))
			*(*int64)(unsafe.Add(mBase, uint32(v8))) = v33
			v35 = F_pg_stat_wal_build_tuple(m, v8, v20)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v37 = v35
				m.G0 = v8 + int32(32)
				return v37
			}
		}
	}
}
func F_pg_stat_get_bgwriter_stat_reset_time(m *base.Module, l0 int32) int32 {
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
	v2 = F_pgstat_fetch_stat_bgwriter(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+24))
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
func F_pg_stat_get_checkpointer_buffers_written(m *base.Module, l0 int32) int32 {
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
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+64))
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
func F_pg_stat_get_checkpointer_restartpoints_requested(m *base.Module, l0 int32) int32 {
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
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+32))
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
func F_pg_stat_get_checkpointer_sync_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = F_pgstat_fetch_stat_checkpointer(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+56))
		v8 = F_Float8GetDatum(m, base.F64_convert_i64_s(v6))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_pg_stat_get_db_active_time(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v10 = F_Float8GetDatum(m, float64(0))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+200))
			v17 = F_Float8GetDatum(m, base.F64_div(base.F64_convert_i64_s(v13), float64(1000)))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v17
			}
		}
	}
}
func F_pg_stat_get_db_blocks_fetched(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+16))
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
func F_pg_stat_get_db_conflict_all(m *base.Module, l0 int32) int32 {
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
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+120))
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v3)+112))
			v15 = *(*int64)(unsafe.Add(mBase, uint32(v3)+104))
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v3)+96))
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v3)+88))
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v3)+80))
			v24 = F_Int64GetDatum(m, v13+(v14+(v15+(v16+(v17+v18)))))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				return v24
			}
		}
	}
}
func F_pg_stat_get_db_conflict_bufferpin(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+112))
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
func F_pg_stat_get_db_parallel_workers_launched(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+248))
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
func F_pg_stat_get_function_total_time(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v4)+8))
			v18 = F_Float8GetDatum(m, base.F64_div(base.F64_convert_i64_s(v14), float64(1000)))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				return v18
			}
		}
	}
}
func F_pg_stat_get_last_analyze_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pgstat_fetch_stat_tabentry(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)+152))
			if v9 != int64(0) {
				v17 = F_Int64GetDatum(m, v9)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return v17
				}
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				return int32(0)
			}
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int32(0)
		}
	}
}
func F_pg_stat_get_last_autovacuum_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pgstat_fetch_stat_tabentry(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)+136))
			if v9 != int64(0) {
				v17 = F_Int64GetDatum(m, v9)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return v17
				}
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				return int32(0)
			}
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int32(0)
		}
	}
}
func F_pg_stat_get_subscription(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v62 int32
	_ = v62
	var v77 int64
	_ = v77
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	v2 = int32(0)
	v24 = m.G0
	v26 = v24 + int32(-64)
	m.G0 = v26
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v28 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v32 = v31
	goto L3
L2:
	;
	v32 = v2
	goto L3
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v44 = F_LWLockAcquire(m, v40+int32(5504), int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[512]))
	if v47 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v365 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v365+int32(5504))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L4
	} else {
		goto L71
	}
L8:
	;
	v62 = v2
	goto L9
L9:
	;
	v77 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v24+int32(-24)))) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v26)+24)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = v77
	v87 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(-56)))) = uint16(v87)
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = v77
	v92 = *(*int32)(unsafe.Add(mBase, _consts[513]))
	v95 = v92 + v62*int32(112)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+36))
	if v96 == v87 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L7
L11:
	;
	v337 = v62 + int32(1)
	v339 = *(*int32)(unsafe.Add(mBase, _consts[512]))
	if v337 < v339 {
		v62 = v337
		goto L9
	} else {
		goto L70
	}
L12:
	;
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v95)+120))
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v95)+112))
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v95)+104))
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v95)+96))
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v95)+88))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v95)+80))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v95)+52))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+32)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
	v109 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v96)+44))
	if v111 == v109 {
		v222 = v109
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v222 == int32(0) {
		goto L11
	} else {
		goto L25
	}
L14:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v119 = F_LWLockAcquire(m, v115+int32(512), int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _consts[515]))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if int32(0) < v123 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	v133 = v109
	goto L20
L17:
	;
	v173 = v109
	goto L18
L18:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v194+int32(512))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L24
	}
L19:
	;
	v173 = base.B2i32(v167 != int32(0))
	goto L18
L20:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v122+int32(36)+v133<<(uint(int32(2))%32))))
	v159 = v129 + v156*int32(640)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+44))
	if v160 == v111 {
		v167 = v159
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v167 = int32(0)
	goto L19
L22:
	;
	v163 = v133 + int32(1)
	if v163 != v123 {
		v133 = v163
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v222 = v173
	goto L13
L25:
	;
	v225 = int32(0)
	if base.B2i32(v32 == v225)|base.B2i32(v32 == v106) == v225 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v96)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v106
	v234 = v107 & int32(1)
	if v234 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	if v103 == int64(0) {
		goto L35
	} else {
		goto L36
	}
L28:
	;
	v249 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+3)) = uint8(v249)
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v231
	v242 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)) = uint8(v242)
	if v234 == int32(0) {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	if v108 != int32(1) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v105
	goto L28
L32:
	;
	if v108 != int32(3) {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v104
	goto L27
L34:
	;
	if v102 == int64(0) {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	v253 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+4)) = uint8(v253)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v255 = F_Int64GetDatum(m, v103)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v255
	goto L34
L39:
	;
	if v101 == int64(0) {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v260 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+5)) = uint8(v260)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v262 = F_Int64GetDatum(m, v102)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v262
	goto L39
L44:
	;
	if v100 == int64(0) {
		goto L50
	} else {
		goto L51
	}
L45:
	;
	v267 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+6)) = uint8(v267)
	goto L44
L46:
	;
	goto L47
L47:
	;
	v269 = F_Int64GetDatum(m, v101)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v269
	goto L44
L49:
	;
	if v99 == int64(0) {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	v274 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+7)) = uint8(v274)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v276 = F_Int64GetDatum(m, v100)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v276
	goto L49
L54:
	;
	switch v108 {
	case 0:
		goto L62
	case 1:
		goto L63
	case 2:
		v302 = int32(19325)
		goto L60
	case 3:
		goto L61
	default:
		goto L59
	}
L55:
	;
	v281 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+8)) = uint8(v281)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v283 = F_Int64GetDatum(m, v99)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v283
	goto L54
L59:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
	F_tuplestore_putvalues(m, v307, v308, v24+int32(-48), v26)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L4
	} else {
		goto L68
	}
L60:
	;
	v303 = F_cstring_to_text(m, v302)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L67
	}
L61:
	;
	v302 = int32(19316)
	goto L60
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L4
	} else {
		goto L64
	}
L63:
	;
	v302 = int32(256180)
	goto L60
L64:
	;
	F_errmsg_internal(m, int32(364257), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(490041), int32(1377), int32(244550))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v303
	goto L59
L68:
	;
	if v32 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	goto L11
L70:
	;
	goto L10
L71:
	;
	m.G0 = v26 - int32(-64)
	return int32(0)
}
func F_pg_stat_get_tuples_fetched(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+24))
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
func F_pg_stat_get_tuples_hot_updated(m *base.Module, l0 int32) int32 {
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
func F_pg_stat_get_tuples_updated(m *base.Module, l0 int32) int32 {
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
func F_pg_stat_get_vacuum_count(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+128))
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
func F_pg_stat_get_xact_blocks_hit(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+120))
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
func F_pg_stat_get_xact_numscans(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+16))
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
func F_pg_stat_reset_backend_stats(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
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
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_BackendPidGetProc(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return int32(0)
L2:
	;
	return int32(0)
L3:
	;
	if v4 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v10 = int32(0)
	if v3 == v10 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v47 = v4
	goto L6
L6:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[101]))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v53 = base.I32_div_s(v47-v50, int32(640))
	v54 = F_pgstat_get_beentry_by_proc_number(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L18
	}
L7:
	;
	if v44 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L8:
	;
	v44 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[1096]))
	v19 = v10
	goto L12
L11:
	;
	v44 = v39
	goto L7
L12:
	;
	v24 = v17 + v19*int32(640)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v25 == v3 {
		v39 = v24
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v44 = int32(0)
	goto L7
L14:
	;
	v31 = v17 + (v19|int32(1))*int32(640)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	if v32 == v3 {
		v39 = v31
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v35 = v19 + int32(2)
	if v35 != int32(38) {
		v19 = v35
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v47 = v44
	goto L6
L18:
	;
	if v54 == int32(0) {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	goto L20
L20:
	;
	if int32(base.Ui32(int32(115186))>>(uint(v58)%32))&base.B2i32(base.Ui32(v58) < base.Ui32(int32(17))) == int32(0) {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_pgstat_reset(m, int32(6), int32(0), base.I64_extend_i32_s(v53))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	goto L1
}
func F_pg_stat_reset_slru(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v116 int32
	_ = v116
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v5 == int32(1) {
		F_pgstat_reset_of_kind(m, int32(11))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = F_text_to_cstring(m, v16)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v23 = m.G0
				v24 = int32(16)
				v25 = v23 - v24
				m.G0 = v25
				F___gettimeofday(m, v25)
				mBase = m.M
				v28 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
				v29 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25)+8)))
				m.G0 = v25 + v24
				v39 = F_strcmp(m, int32(234320), v18)
				mBase = m.M
				if v39 == int32(0) {
					v73 = int32(0)
				} else {
					v44 = F_strcmp(m, int32(226281), v18)
					mBase = m.M
					if v44 == int32(0) {
						v73 = int32(1)
					} else {
						v49 = F_strcmp(m, int32(104554), v18)
						mBase = m.M
						if v49 == int32(0) {
							v73 = int32(2)
						} else {
							v54 = F_strcmp(m, int32(20623), v18)
							mBase = m.M
							if v54 == int32(0) {
								v73 = int32(3)
							} else {
								v59 = F_strcmp(m, int32(387195), v18)
								mBase = m.M
								if v59 == int32(0) {
									v73 = int32(4)
								} else {
									v64 = F_strcmp(m, int32(253431), v18)
									mBase = m.M
									if v64 == int32(0) {
										v73 = int32(5)
									} else {
										v71 = F_strcmp(m, int32(255376), v18)
										mBase = m.M
										if v71 != 0 {
											v72 = int32(7)
										} else {
											v72 = int32(6)
										}
										v73 = v72
									}
								}
							}
						}
					}
				}
				v75 = *(*int32)(unsafe.Add(mBase, _consts[561]))
				v77 = v75 + int32(52744)
				v79 = F_LWLockAcquire(m, v77, int32(0))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					v83 = v75 + v73<<(uint(int32(6))%32)
					v86 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_consts[1097]))) = v86
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_consts[1098]))) = v86
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_consts[1099]))) = v86
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_consts[1100]))) = v86
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_consts[1101]))) = v86
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_consts[1102]))) = v86
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_consts[1103]))) = v86
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_consts[1104]))) = v29 + v28*int64(1000000) - int64(946684800000000)
					F_LWLockRelease(m, v77)
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				}
			}
		}
	}
}
func F_pg_stat_wal_build_tuple(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int64
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int64
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(320)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+304)) = v3
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+284)) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+280)) = v3
	v16 = F_CreateTemplateTupleDesc(m, int32(5))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		F_TupleDescInitEntry(m, v16, int32(1), int32(170354), int32(20), int32(-1), int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			F_TupleDescInitEntry(m, v16, int32(2), int32(316458), int32(20), int32(-1), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_TupleDescInitEntry(m, v16, int32(3), int32(157355), int32(1700), int32(-1), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					F_TupleDescInitEntry(m, v16, int32(4), int32(300298), int32(20), int32(-1), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_TupleDescInitEntry(m, v16, int32(5), int32(104771), int32(1184), int32(-1), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v55 = F_BlessTupleDesc(m, v16)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								v57 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
								v58 = F_Int64GetDatum(m, v57)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+288)) = v58
									v61 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
									v62 = F_Int64GetDatum(m, v61)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+292)) = v62
										v65 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v7))) = v65
										v71 = F_pg_snprintf(m, v7+int32(16), int32(256), int32(38064), v7)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											v74 = int32(0)
											v79 = F_DirectFunctionCall3Coll(m, int32(408), v74, v7+int32(16), v74, int32(-1))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7)+296)) = v79
												v82 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
												v83 = F_Int64GetDatum(m, v82)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v7)+300)) = v83
													if l1 != int64(0) {
														v88 = F_Int64GetDatum(m, l1)
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v7)+304)) = v88
															v97 = F_heap_form_tuple(m, v16, v7+int32(288), v7+int32(280))
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return int32(0)
															} else {
																v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
																v100 = F_HeapTupleHeaderGetDatum(m, v99)
																mBase = m.M
																v101 = m.ExcPending
																if v101 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v7 + int32(320)
																	return v100
																}
															}
														}
													} else {
														v91 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v7)+284)) = uint8(v91)
														v97 = F_heap_form_tuple(m, v16, v7+int32(288), v7+int32(280))
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return int32(0)
														} else {
															v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
															v100 = F_HeapTupleHeaderGetDatum(m, v99)
															mBase = m.M
															v101 = m.ExcPending
															if v101 != 0 {
																return int32(0)
															} else {
																m.G0 = v7 + int32(320)
																return v100
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
