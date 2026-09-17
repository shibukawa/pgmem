package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_stat_file(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int64
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int64
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	v5 = m.G0
	v7 = v5 - int32(144)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v14 == int32(2) {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v20 = base.B2i32(v17 != int32(0))
		} else {
			v20 = int32(0)
		}
		v21 = F_convert_and_check_filename(m, v10)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v27 = F___fstatat(m, int32(-100), v21, v7+int32(48), int32(0))
			mBase = m.M
			if v27 < int32(0) {
				if v20 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v21
							F_errmsg(m, int32(_a_F_pg_stat_file_0), v7)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_pg_stat_file_1), int32(437), int32(_a_F_pg_stat_file_2))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
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
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_file[0]))
					if v33 != int32(44) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v21
								F_errmsg(m, int32(_a_F_pg_stat_file_0), v7)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_stat_file_1), int32(437), int32(_a_F_pg_stat_file_2))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
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
						v36 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v36)
						v154 = int32(0)
						m.G0 = v7 + int32(144)
						return v154
					}
				}
			} else {
				v55 = F_CreateTemplateTupleDesc(m, int32(6))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					F_TupleDescInitEntry(m, v55, int32(1), int32(_a_F_pg_stat_file_3), int32(20), int32(-1), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						F_TupleDescInitEntry(m, v55, int32(2), int32(_a_F_pg_stat_file_4), int32(1184), int32(-1), int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							F_TupleDescInitEntry(m, v55, int32(3), int32(_a_F_pg_stat_file_5), int32(1184), int32(-1), int32(0))
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								F_TupleDescInitEntry(m, v55, int32(4), int32(_a_F_pg_stat_file_6), int32(1184), int32(-1), int32(0))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									F_TupleDescInitEntry(m, v55, int32(5), int32(_a_F_pg_stat_file_7), int32(1184), int32(-1), int32(0))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										F_TupleDescInitEntry(m, v55, int32(6), int32(_a_F_pg_stat_file_8), int32(16), int32(-1), int32(0))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											v99 = F_BlessTupleDesc(m, v55)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return int32(0)
											} else {
												v101 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v101)
												*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v101
												v105 = *(*int64)(unsafe.Add(mBase, uint32(v7)+72))
												v106 = F_Int64GetDatum(m, v105)
												mBase = m.M
												v107 = m.ExcPending
												if v107 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v106
													v109 = *(*int64)(unsafe.Add(mBase, uint32(v7)+88))
													v114 = F_Int64GetDatum(m, v109*int64(1000000)-int64(946684800000000))
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v114
														v117 = *(*int64)(unsafe.Add(mBase, uint32(v7)+104))
														v122 = F_Int64GetDatum(m, v117*int64(1000000)-int64(946684800000000))
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v122
															v125 = *(*int64)(unsafe.Add(mBase, uint32(v7)+120))
															v130 = F_Int64GetDatum(m, v125*int64(1000000)-int64(946684800000000))
															mBase = m.M
															v131 = m.ExcPending
															if v131 != 0 {
																return int32(0)
															} else {
																v132 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v132)
																*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v130
																v135 = *(*int32)(unsafe.Add(mBase, uint32(v7)+52))
																*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = base.B2i32(v135&int32(_a_F_pg_stat_file_9) == int32(_a_F_pg_stat_file_10))
																v145 = F_heap_form_tuple(m, v55, v7+int32(16), v7+int32(8))
																mBase = m.M
																v146 = m.ExcPending
																if v146 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v21)
																	mBase = m.M
																	v148 = m.ExcPending
																	if v148 != 0 {
																		return int32(0)
																	} else {
																		v149 = *(*int32)(unsafe.Add(mBase, uint32(v145)+16))
																		v150 = F_HeapTupleHeaderGetDatum(m, v149)
																		mBase = m.M
																		v151 = m.ExcPending
																		if v151 != 0 {
																			return int32(0)
																		} else {
																			v154 = v150
																			m.G0 = v7 + int32(144)
																			return v154
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
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
			v31 = v2
			m.G0 = v8 + int32(32)
			return v31
		} else {
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v12)+2912))
			*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v21
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v12)+2904))
			*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v23
			v25 = *(*int64)(unsafe.Add(mBase, uint32(v12)+2896))
			*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v25
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v12)+2888))
			*(*int64)(unsafe.Add(mBase, uint32(v8))) = v27
			v29 = F_pg_stat_wal_build_tuple(m, v8, v20)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = v29
				m.G0 = v8 + int32(32)
				return v31
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v69 int64
	_ = v69
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
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
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	v2 = int32(0)
	v22 = m.G0
	v24 = v22 + int32(-64)
	m.G0 = v24
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v26 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v30 = v29
	goto L3
L2:
	;
	v30 = v2
	goto L3
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_subscription[0]))
	v42 = F_LWLockAcquire(m, v38+int32(_a_F_pg_stat_get_subscription_0), int32(1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_subscription[1]))
	if v45 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v352 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_subscription[0]))
	F_LWLockRelease(m, v352+int32(_a_F_pg_stat_get_subscription_0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L69
	}
L8:
	;
	v55 = v2
	goto L9
L9:
	;
	v69 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+48)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v24)+40)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v24)+32)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v24)+24)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v24)+16)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = v69
	v81 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+8)) = uint16(v81)
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_subscription[2]))
	v87 = v84 + v55*int32(112)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+36))
	if v88 == v81 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L7
L11:
	;
	v326 = v55 + int32(1)
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_subscription[1]))
	if v326 < v328 {
		v55 = v326
		goto L9
	} else {
		goto L68
	}
L12:
	;
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v87)+120))
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v87)+112))
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v87)+104))
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v87)+96))
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v87)+88))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+80))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v87)+52))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v87)+48))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+32)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	v101 = int32(0)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v88)+44))
	if v103 == v101 {
		v208 = v101
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v209 = int32(0)
	if base.B2i32(v208 == v209)|base.B2i32(base.B2i32(v30 == v209)|base.B2i32(v30 == v98) == v209) != 0 {
		goto L11
	} else {
		goto L25
	}
L14:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_subscription[0]))
	v111 = F_LWLockAcquire(m, v107+int32(512), int32(1))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_subscription[3]))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if int32(0) < v115 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_subscription[4]))
	v122 = v101
	goto L20
L17:
	;
	v160 = v101
	goto L18
L18:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_subscription[0]))
	F_LWLockRelease(m, v182+int32(512))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L24
	}
L19:
	;
	v160 = base.B2i32(v157 != int32(0))
	goto L18
L20:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v114+int32(36)+v122<<(uint(int32(2))%32))))
	v149 = v121 + v146*int32(640)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+44))
	if v150 == v103 {
		v157 = v149
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v157 = int32(0)
	goto L19
L22:
	;
	v153 = v122 + int32(1)
	if v153 != v115 {
		v122 = v153
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v208 = v160
	goto L13
L25:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v88)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v98
	v220 = int32(1)
	v221 = v99 & v220
	v222 = int32(0)
	if base.B2i32(v221 == v222)|base.B2i32(v100 != v220) == v222 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	if v95 == int64(0) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	v240 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+3)) = uint8(v240)
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v97
	goto L27
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v218
	v232 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)) = uint8(v232)
	if base.B2i32(v221 == int32(0))|base.B2i32(v100 != int32(3)) != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v96
	goto L26
L32:
	;
	if v94 == int64(0) {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v244 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)) = uint8(v244)
	goto L32
L34:
	;
	goto L35
L35:
	;
	v246 = F_Int64GetDatum(m, v95)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v246
	goto L32
L37:
	;
	if v93 == int64(0) {
		goto L43
	} else {
		goto L44
	}
L38:
	;
	v251 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+5)) = uint8(v251)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v253 = F_Int64GetDatum(m, v94)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v253
	goto L37
L42:
	;
	if v92 == int64(0) {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	v258 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+6)) = uint8(v258)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v260 = F_Int64GetDatum(m, v93)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v260
	goto L42
L47:
	;
	if v91 == int64(0) {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v265 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+7)) = uint8(v265)
	goto L47
L49:
	;
	goto L50
L50:
	;
	v267 = F_Int64GetDatum(m, v92)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v267
	goto L47
L52:
	;
	switch v100 {
	case 0:
		goto L60
	case 1:
		goto L61
	case 2:
		v293 = int32(_a_F_pg_stat_get_subscription_1)
		goto L58
	case 3:
		goto L59
	default:
		goto L57
	}
L53:
	;
	v272 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)) = uint8(v272)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v274 = F_Int64GetDatum(m, v91)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v274
	goto L52
L57:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
	F_tuplestore_putvalues(m, v298, v299, v22+int32(-48), v24)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L4
	} else {
		goto L66
	}
L58:
	;
	v294 = F_cstring_to_text(m, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L4
	} else {
		goto L65
	}
L59:
	;
	v293 = int32(_a_F_pg_stat_get_subscription_2)
	goto L58
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L4
	} else {
		goto L62
	}
L61:
	;
	v293 = int32(_a_F_pg_stat_get_subscription_3)
	goto L58
L62:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_get_subscription_4), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_pg_stat_get_subscription_5), int32(1377), int32(_a_F_pg_stat_get_subscription_6))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+52)) = v294
	goto L57
L66:
	;
	if v30 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	goto L11
L68:
	;
	goto L10
L69:
	;
	m.G0 = v24 - int32(-64)
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
	var v18 int32
	_ = v18
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v71 int32
	_ = v71
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
	v48 = v4
	goto L6
L6:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_backend_stats[0]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v54 = base.I32_div_s(v48-v51, int32(640))
	v55 = F_pgstat_get_beentry_by_proc_number(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L19
	}
L7:
	;
	if v45 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L8:
	;
	v45 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_backend_stats[1]))
	v22 = v10
	goto L13
L11:
	;
	v45 = v39
	goto L7
L12:
	;
	v39 = v29 + int32(640)
	goto L11
L13:
	;
	v25 = v22 * int32(640)
	v26 = v18 + v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+44))
	if v27 == v3 {
		v39 = v26
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v45 = int32(0)
	goto L7
L15:
	;
	v29 = v18 + v25
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+684))
	if v30 == v3 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v33 = v22 + int32(2)
	if v33 != int32(38) {
		v22 = v33
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v48 = v45
	goto L6
L19:
	;
	if v55 == int32(0) {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	goto L21
L21:
	;
	if int32(base.Ui32(int32(_a_F_pg_stat_reset_backend_stats_0))>>(uint(v59)%32))&base.B2i32(base.Ui32(v59) < base.Ui32(int32(17))) == int32(0) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_pgstat_reset(m, int32(6), int32(0), base.I64_extend_i32_s(v54))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
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
	var v84 int64
	_ = v84
	var v100 int32
	_ = v100
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
				F_gettimeofday(m, v25)
				mBase = m.M
				v28 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
				v29 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25)+8)))
				m.G0 = v25 + v24
				v39 = F_strcmp(m, int32(_a_F_pg_stat_reset_slru_0), v18)
				mBase = m.M
				if v39 == int32(0) {
					v73 = int32(0)
				} else {
					v44 = F_strcmp(m, int32(_a_F_pg_stat_reset_slru_1), v18)
					mBase = m.M
					if v44 == int32(0) {
						v73 = int32(1)
					} else {
						v49 = F_strcmp(m, int32(_a_F_pg_stat_reset_slru_2), v18)
						mBase = m.M
						if v49 == int32(0) {
							v73 = int32(2)
						} else {
							v54 = F_strcmp(m, int32(_a_F_pg_stat_reset_slru_3), v18)
							mBase = m.M
							if v54 == int32(0) {
								v73 = int32(3)
							} else {
								v59 = F_strcmp(m, int32(_a_F_pg_stat_reset_slru_4), v18)
								mBase = m.M
								if v59 == int32(0) {
									v73 = int32(4)
								} else {
									v64 = F_strcmp(m, int32(_a_F_pg_stat_reset_slru_5), v18)
									mBase = m.M
									if v64 == int32(0) {
										v73 = int32(5)
									} else {
										v71 = F_strcmp(m, int32(_a_F_pg_stat_reset_slru_6), v18)
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
				v75 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_slru[0]))
				v77 = v75 + int32(_a_F_pg_stat_reset_slru_7)
				v79 = F_LWLockAcquire(m, v77, int32(0))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					v83 = v75 + v73<<(uint(int32(6))%32)
					v84 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_pg_stat_reset_slru[1]))) = v84
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_pg_stat_reset_slru[2]))) = v84
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_pg_stat_reset_slru[3]))) = v84
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_pg_stat_reset_slru[4]))) = v84
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_pg_stat_reset_slru[5]))) = v84
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_pg_stat_reset_slru[6]))) = v84
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_pg_stat_reset_slru[7]))) = v84
					*(*int64)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_pg_stat_reset_slru[8]))) = v29 + v28*int64(1000000) - int64(946684800000000)
					F_LWLockRelease(m, v77)
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				}
			}
		}
	}
}
func F_pg_stat_statements_1_12(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pg_stat_statements_internal(m, l0, int32(8), base.B2i32(v3 != int32(0)))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pg_stat_statements_1_9(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pg_stat_statements_internal(m, l0, int32(5), base.B2i32(v3 != int32(0)))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pg_stat_statements_reset(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v2 = int32(0)
	v6 = F_entry_reset(m, v2, v2, int64(0), v2)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pg_stat_wal_build_tuple(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(320)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+304)) = v3
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+284)) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+280)) = v3
	v17 = F_CreateTemplateTupleDesc(m, int32(5))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		F_TupleDescInitEntry(m, v17, int32(1), int32(_a_F_pg_stat_wal_build_tuple_0), int32(20), int32(-1), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			F_TupleDescInitEntry(m, v17, int32(2), int32(_a_F_pg_stat_wal_build_tuple_1), int32(20), int32(-1), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				F_TupleDescInitEntry(m, v17, int32(3), int32(_a_F_pg_stat_wal_build_tuple_2), int32(1700), int32(-1), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					F_TupleDescInitEntry(m, v17, int32(4), int32(_a_F_pg_stat_wal_build_tuple_3), int32(20), int32(-1), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						F_TupleDescInitEntry(m, v17, int32(5), int32(_a_F_pg_stat_wal_build_tuple_4), int32(1184), int32(-1), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v56 = F_BlessTupleDesc(m, v17)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								v58 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
								v59 = F_Int64GetDatum(m, v58)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+288)) = v59
									v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
									v63 = F_Int64GetDatum(m, v62)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+292)) = v63
										v66 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v8))) = v66
										v69 = v8 + int32(16)
										v72 = F_pg_snprintf(m, v69, int32(256), int32(_a_F_pg_stat_wal_build_tuple_5), v8)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											v75 = int32(0)
											v78 = F_DirectFunctionCall3Coll(m, int32(408), v75, v69, v75, int32(-1))
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8)+296)) = v78
												v81 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
												v82 = F_Int64GetDatum(m, v81)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+300)) = v82
													if l1 != int64(0) {
														v87 = F_Int64GetDatum(m, l1)
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v8)+304)) = v87
															v96 = F_heap_form_tuple(m, v17, v8+int32(288), v8+int32(280))
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return int32(0)
															} else {
																v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
																v99 = F_HeapTupleHeaderGetDatum(m, v98)
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v8 + int32(320)
																	return v99
																}
															}
														}
													} else {
														v90 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v8)+284)) = uint8(v90)
														v96 = F_heap_form_tuple(m, v17, v8+int32(288), v8+int32(280))
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return int32(0)
														} else {
															v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
															v99 = F_HeapTupleHeaderGetDatum(m, v98)
															mBase = m.M
															v100 = m.ExcPending
															if v100 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(320)
																return v99
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
