package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_stat_get_backend_idset(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	if v6 == int32(0) {
		v9 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
			v15 = F_MemoryContextAlloc(m, v13, int32(4))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v15
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(0)
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v27 = v25 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v24))) = v27
				v29 = F_pgstat_fetch_stat_numbackends(m)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					if v27 <= v29 {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
						v33 = F_pgstat_get_local_beentry_by_index(m, v32)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
							*(*int64)(unsafe.Add(mBase, uint32(v23))) = v35 + int64(1)
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = int32(1)
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v33)+408))
							return v42
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = int32(2)
							v49 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v49)
							return int32(0)
						}
					}
				}
			}
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
		v27 = v25 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v24))) = v27
		v29 = F_pgstat_fetch_stat_numbackends(m)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			if v27 <= v29 {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v33 = F_pgstat_get_local_beentry_by_index(m, v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
					*(*int64)(unsafe.Add(mBase, uint32(v23))) = v35 + int64(1)
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = int32(1)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v33)+408))
					return v42
				}
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = int32(2)
					v49 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v49)
					return int32(0)
				}
			}
		}
	}
}
func F_pg_stat_get_backend_xact_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pgstat_get_beentry_by_proc_number(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[168]))
			v18 = F_has_privs_of_role(m, v16, int32(3375))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 != 0 {
					v29 = *(*int64)(unsafe.Add(mBase, uint32(v5)+24))
					if v29 == int64(0) {
						v32 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
						return int32(0)
					} else {
						v36 = F_Int64GetDatum(m, v29)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							return v36
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, _consts[168]))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v5)+52))
					v23 = F_has_privs_of_role(m, v21, v22)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						if v23 != 0 {
							v29 = *(*int64)(unsafe.Add(mBase, uint32(v5)+24))
							if v29 == int64(0) {
								v32 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
								return int32(0)
							} else {
								v36 = F_Int64GetDatum(m, v29)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									return v36
								}
							}
						} else {
							v25 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v25)
							return int32(0)
						}
					}
				}
			}
		}
	}
}
func F_pg_stat_get_bgwriter_buf_written_clean(m *base.Module, l0 int32) int32 {
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
func F_pg_stat_get_bgwriter_maxwritten_clean(m *base.Module, l0 int32) int32 {
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
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+8))
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
func F_pg_stat_get_db_blk_write_time(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+176))
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
func F_pg_stat_get_db_checksum_failures(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, _consts[284]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+252))
	if base.B2i32(v6 != int32(0)) == int32(0) {
		v11 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
		return int32(0)
	} else {
		v15 = F_pgstat_fetch_stat_dbentry(m, v3)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 == int32(0) {
				v22 = F_Int64GetDatum(m, int64(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					return v22
				}
			} else {
				v25 = *(*int64)(unsafe.Add(mBase, uint32(v15)+152))
				v26 = F_Int64GetDatum(m, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					return v26
				}
			}
		}
	}
}
func F_pg_stat_get_db_conflict_snapshot(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+96))
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
func F_pg_stat_get_db_idle_in_transaction_time(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+208))
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
func F_pg_stat_get_db_parallel_workers_to_launch(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+240))
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
func F_pg_stat_get_db_sessions(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+184))
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
func F_pg_stat_get_db_tuples_deleted(m *base.Module, l0 int32) int32 {
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
func F_pg_stat_get_dead_tuples(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+80))
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
func F_pg_stat_get_progress_info(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v323 int32
	_ = v323
	var v324 int64
	_ = v324
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v373 int64
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int64
	_ = v381
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v416 int32
	_ = v416
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	v11 = m.G0
	v13 = v11 - int32(144)
	m.G0 = v13
	v15 = F_pgstat_fetch_stat_numbackends(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum_packed(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = F_text_to_cstring(m, v20)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = v22
	v30 = int32(531325)
	goto L8
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L124
	}
L6:
	;
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L99
	}
L7:
	;
	if v67 == int32(0) {
		v293 = int32(1)
		goto L6
	} else {
		goto L20
	}
L8:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v33 == v34 {
		v56 = v33
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v67 = int32(0)
	goto L7
L10:
	;
	v58 = int32(1)
	if v56 != 0 {
		v29 = v29 + v58
		v30 = v30 + v58
		goto L8
	} else {
		goto L19
	}
L11:
	;
	if base.Ui32((v33-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v44 = v33 | int32(32)
	goto L14
L13:
	;
	v44 = v33
	goto L14
L14:
	;
	if base.Ui32((v34-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v53 = v34 | int32(32)
	goto L17
L16:
	;
	v53 = v34
	goto L17
L17:
	;
	if v44 == v53 {
		v56 = v44
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v67 = v44 - v53
	goto L7
L19:
	;
	goto L9
L20:
	;
	v73 = v22
	v74 = int32(537443)
	goto L22
L21:
	;
	if v111 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v77 == v78 {
		v100 = v77
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v111 = int32(0)
	goto L21
L24:
	;
	v102 = int32(1)
	if v100 != 0 {
		v73 = v73 + v102
		v74 = v74 + v102
		goto L22
	} else {
		goto L33
	}
L25:
	;
	if base.Ui32((v77-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v88 = v77 | int32(32)
	goto L28
L27:
	;
	v88 = v77
	goto L28
L28:
	;
	if base.Ui32((v78-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v97 = v78 | int32(32)
	goto L31
L30:
	;
	v97 = v78
	goto L31
L31:
	;
	if v88 == v97 {
		v100 = v88
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v111 = v88 - v97
	goto L21
L33:
	;
	goto L23
L34:
	;
	v293 = int32(2)
	goto L6
L35:
	;
	goto L36
L36:
	;
	v118 = v22
	v119 = int32(525390)
	goto L38
L37:
	;
	if v156 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v122 == v123 {
		v145 = v122
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v156 = int32(0)
	goto L37
L40:
	;
	v147 = int32(1)
	if v145 != 0 {
		v118 = v118 + v147
		v119 = v119 + v147
		goto L38
	} else {
		goto L49
	}
L41:
	;
	if base.Ui32((v122-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v133 = v122 | int32(32)
	goto L44
L43:
	;
	v133 = v122
	goto L44
L44:
	;
	if base.Ui32((v123-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v142 = v123 | int32(32)
	goto L47
L46:
	;
	v142 = v123
	goto L47
L47:
	;
	if v133 == v142 {
		v145 = v133
		goto L40
	} else {
		goto L48
	}
L48:
	;
	v156 = v133 - v142
	goto L37
L49:
	;
	goto L39
L50:
	;
	v293 = int32(3)
	goto L6
L51:
	;
	goto L52
L52:
	;
	v163 = v22
	v164 = int32(510195)
	goto L54
L53:
	;
	if v201 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L54:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v167 == v168 {
		v190 = v167
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v201 = int32(0)
	goto L53
L56:
	;
	v192 = int32(1)
	if v190 != 0 {
		v163 = v163 + v192
		v164 = v164 + v192
		goto L54
	} else {
		goto L65
	}
L57:
	;
	if base.Ui32((v167-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v178 = v167 | int32(32)
	goto L60
L59:
	;
	v178 = v167
	goto L60
L60:
	;
	if base.Ui32((v168-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v187 = v168 | int32(32)
	goto L63
L62:
	;
	v187 = v168
	goto L63
L63:
	;
	if v178 == v187 {
		v190 = v178
		goto L56
	} else {
		goto L64
	}
L64:
	;
	v201 = v178 - v187
	goto L53
L65:
	;
	goto L55
L66:
	;
	v293 = int32(4)
	goto L6
L67:
	;
	goto L68
L68:
	;
	v208 = v22
	v209 = int32(526638)
	goto L70
L69:
	;
	if v246 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L70:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if v212 == v213 {
		v235 = v212
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v246 = int32(0)
	goto L69
L72:
	;
	v237 = int32(1)
	if v235 != 0 {
		v208 = v208 + v237
		v209 = v209 + v237
		goto L70
	} else {
		goto L81
	}
L73:
	;
	if base.Ui32((v212-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v223 = v212 | int32(32)
	goto L76
L75:
	;
	v223 = v212
	goto L76
L76:
	;
	if base.Ui32((v213-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v232 = v213 | int32(32)
	goto L79
L78:
	;
	v232 = v213
	goto L79
L79:
	;
	if v223 == v232 {
		v235 = v223
		goto L72
	} else {
		goto L80
	}
L80:
	;
	v246 = v223 - v232
	goto L69
L81:
	;
	goto L71
L82:
	;
	v293 = int32(5)
	goto L6
L83:
	;
	goto L84
L84:
	;
	v253 = v22
	v254 = int32(509245)
	goto L86
L85:
	;
	if v291 != 0 {
		goto L5
	} else {
		goto L98
	}
L86:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	if v257 == v258 {
		v280 = v257
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v291 = int32(0)
	goto L85
L88:
	;
	v282 = int32(1)
	if v280 != 0 {
		v253 = v253 + v282
		v254 = v254 + v282
		goto L86
	} else {
		goto L97
	}
L89:
	;
	if base.Ui32((v257-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v268 = v257 | int32(32)
	goto L92
L91:
	;
	v268 = v257
	goto L92
L92:
	;
	if base.Ui32((v258-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v277 = v258 | int32(32)
	goto L95
L94:
	;
	v277 = v258
	goto L95
L95:
	;
	if v268 == v277 {
		v280 = v268
		goto L88
	} else {
		goto L96
	}
L96:
	;
	v291 = v268 - v277
	goto L85
L97:
	;
	goto L87
L98:
	;
	v293 = int32(6)
	goto L6
L99:
	;
	if int32(0) < v15 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v302 = v13 + int32(16) | int32(2)
	v312 = int32(1)
	goto L103
L101:
	;
	goto L102
L102:
	;
	m.G0 = v13 + int32(144)
	return int32(0)
L103:
	;
	v323 = F__emscripten_memset_bulkmem(m, v13+int32(48), base.I32_extend8_s(int32(0)), int32(92))
	mBase = m.M
	goto L105
L104:
	;
	goto L102
L105:
	;
	v324 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+31)) = v324
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v324
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v324
	v330 = F_pgstat_get_local_beentry_by_index(m, v312)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v330)+220))
	if v293 == v332 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v334
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v330)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v336
	v339 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	v341 = F_has_privs_of_role(m, v339, int32(3375))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L112
	}
L108:
	;
	goto L109
L109:
	;
	v416 = v312 + int32(1)
	if v416 <= v15 {
		v312 = v416
		goto L103
	} else {
		goto L123
	}
L110:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	F_tuplestore_putvalues(m, v397, v398, v13+int32(48), v13+int32(16))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L122
	}
L111:
	;
	v381 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v302))) = v381
	*(*int64)(unsafe.Add(mBase, uint32(v302+int32(13)))) = v381
	*(*int64)(unsafe.Add(mBase, uint32(v302+int32(8)))) = v381
	goto L110
L112:
	;
	if v341 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v330)+52))
	v348 = F_has_privs_of_role(m, v346, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v330)+224))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v352
	v357 = int32(0)
	goto L118
L116:
	;
	if v348 == int32(0) {
		goto L111
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v330+int32(232)+v357<<(uint(int32(3))%32))))
	v374 = F_Int64GetDatum(m, v373)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L120
	}
L119:
	;
	goto L110
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357<<(uint(int32(2))%32)+v13)+60)) = v374
	v378 = v357 + int32(1)
	if v378 != int32(20) {
		v357 = v378
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	goto L109
L123:
	;
	goto L104
L124:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v22
	F_errmsg(m, int32(727294), v13)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(494750), int32(280), int32(241720))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_stat_get_slru(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int64
	_ = v129
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_pgstat_snapshot_fixed(m, int32(11))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = v19 + int32(-16)
	v35 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v35
	v38 = v19 + int32(-24)
	v39 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = v39
	v42 = v19 + int32(-32)
	*(*int64)(unsafe.Add(mBase, uint32(v42))) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v39
	v50 = v19 + int32(-56)
	*(*uint8)(unsafe.Add(mBase, uint32(v50))) = uint8(v35)
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v39
	goto L5
L4:
	;
	if v63 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[1323]))
	goto L7
L7:
	;
	goto L4
L8:
	;
	v71 = v63
	v72 = v2
	goto L11
L9:
	;
	goto L10
L10:
	;
	m.G0 = v21 - int32(-64)
	return int32(0)
L11:
	;
	v84 = v72 << (uint(int32(6)) % 32)
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[1324])))
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[1325])))
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[1326])))
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[1327])))
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[1328])))
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[1329])))
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[1330])))
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[1331])))
	v94 = F_cstring_to_text(m, v71)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v94
	v97 = F_Int64GetDatum(m, v93)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v97
	v100 = F_Int64GetDatum(m, v92)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v100
	v103 = F_Int64GetDatum(m, v91)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v103
	v106 = F_Int64GetDatum(m, v90)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v106
	v109 = F_Int64GetDatum(m, v89)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v109
	v112 = F_Int64GetDatum(m, v88)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v112
	v115 = F_Int64GetDatum(m, v87)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v115
	v118 = F_Int64GetDatum(m, v86)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v118
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	F_tuplestore_putvalues(m, v121, v122, v19+int32(-48), v21)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v127 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v127
	v129 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = v129
	*(*int64)(unsafe.Add(mBase, uint32(v42))) = v129
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v129
	*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v129
	*(*uint8)(unsafe.Add(mBase, uint32(v50))) = uint8(v127)
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v129
	v142 = v72 + int32(1)
	if base.Ui32(v142) <= base.Ui32(int32(7)) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v151 != 0 {
		v71 = v151
		v72 = v142
		goto L11
	} else {
		goto L27
	}
L24:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v142<<(uint(int32(2))%32))+uint32(_consts[1323])))
	v151 = v150
	goto L26
L25:
	;
	v151 = v127
	goto L26
L26:
	;
	goto L23
L27:
	;
	goto L12
}
func F_pg_stat_get_wal_receiver(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v423 int64
	_ = v423
	var v424 int64
	_ = v424
	var v426 int32
	_ = v426
	var v427 int64
	_ = v427
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	v17 = m.G0
	v19 = v17 - int32(1360)
	m.G0 = v19
	v22 = *(*int32)(unsafe.Add(mBase, _consts[866]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+1456))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+1456)) = int32(1)
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[866]))
	F_s_lock(m, v27+int32(1456), int32(495327), int32(1416), int32(215017))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[866]))
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)+96))
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v38)+88))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v38)+80))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v38)+72))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)+56))
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v38)+48))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)+40))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v38)+32))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1453)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v51 = v19 + int32(1024)
	v53 = v38 + int32(1388)
	goto L9
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v170 = v19 + int32(1088)
	v172 = *(*int32)(unsafe.Add(mBase, _consts[866]))
	v174 = v172 + int32(1128)
	goto L41
L7:
	;
	v166 = F_strlen(m, v155)
	mBase = m.M
	goto L6
L9:
	;
	goto L10
L10:
	;
	v60 = int32(63)
	if (v51^v53)&int32(3) != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v159 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v159)
	goto L7
L12:
	;
	v140 = v135
	v141 = v136
	v142 = v137
	goto L34
L13:
	;
	if v130 == int32(0) {
		v155 = v128
		v156 = v129
		goto L11
	} else {
		goto L33
	}
L14:
	;
	v128 = v53
	v129 = v51
	v130 = v60
	goto L13
L15:
	;
	goto L16
L16:
	;
	if v53&int32(3) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v97 == int32(0) {
		v155 = v94
		v156 = v95
		goto L11
	} else {
		goto L26
	}
L18:
	;
	v94 = v53
	v95 = v51
	v96 = v60
	v97 = int32(1)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v73 = v53
	v74 = v51
	v75 = v60
	goto L21
L21:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v77)
	if v77 == int32(0) {
		v135 = v73
		v136 = v74
		v137 = v75
		goto L12
	} else {
		goto L23
	}
L22:
	;
	v94 = v88
	v95 = v82
	v96 = v84
	v97 = v86
	goto L17
L23:
	;
	v81 = int32(1)
	v82 = v74 + v81
	v84 = v75 - v81
	v85 = int32(0)
	v86 = base.B2i32(v84 != v85)
	v88 = v73 + v81
	if v88&int32(3) == v85 {
		v94 = v88
		v95 = v82
		v96 = v84
		v97 = v86
		goto L17
	} else {
		goto L24
	}
L24:
	;
	if v84 != 0 {
		v73 = v88
		v74 = v82
		v75 = v84
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	if v100 == int32(0) {
		v128 = v94
		v129 = v95
		v130 = v96
		goto L13
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32(v96) < base.Ui32(int32(4)) {
		v128 = v94
		v129 = v95
		v130 = v96
		goto L13
	} else {
		goto L28
	}
L28:
	;
	v106 = v94
	v107 = v95
	v108 = v96
	goto L29
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v114 = int32(-2139062144)
	if (int32(16843008)-v111|v111)&v114 != v114 {
		v135 = v106
		v136 = v107
		v137 = v108
		goto L12
	} else {
		goto L31
	}
L30:
	;
	v128 = v122
	v129 = v120
	v130 = v124
	goto L13
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v111
	v119 = int32(4)
	v120 = v107 + v119
	v122 = v106 + v119
	v124 = v108 - v119
	if base.Ui32(int32(3)) < base.Ui32(v124) {
		v106 = v122
		v107 = v120
		v108 = v124
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v135 = v128
	v136 = v129
	v137 = v130
	goto L12
L34:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v144)
	if v144 == int32(0) {
		v155 = v140
		v156 = v141
		goto L11
	} else {
		goto L36
	}
L35:
	;
	v155 = v151
	v156 = v149
	goto L11
L36:
	;
	v148 = int32(1)
	v149 = v141 + v148
	v151 = v140 + v148
	v153 = v142 - v148
	if v153 != 0 {
		v140 = v151
		v141 = v149
		v142 = v153
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _consts[866]))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+1384))
	v294 = v291 + int32(104)
	goto L73
L39:
	;
	v287 = F_strlen(m, v276)
	mBase = m.M
	goto L38
L41:
	;
	goto L42
L42:
	;
	v181 = int32(254)
	if (v170^v174)&int32(3) != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v280 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v277))) = uint8(v280)
	goto L39
L44:
	;
	v261 = v256
	v262 = v257
	v263 = v258
	goto L66
L45:
	;
	if v251 == int32(0) {
		v276 = v249
		v277 = v250
		goto L43
	} else {
		goto L65
	}
L46:
	;
	v249 = v174
	v250 = v170
	v251 = v181
	goto L45
L47:
	;
	goto L48
L48:
	;
	if v174&int32(3) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v218 == int32(0) {
		v276 = v215
		v277 = v216
		goto L43
	} else {
		goto L58
	}
L50:
	;
	v215 = v174
	v216 = v170
	v217 = v181
	v218 = int32(1)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v194 = v174
	v195 = v170
	v196 = v181
	goto L53
L53:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v198)
	if v198 == int32(0) {
		v256 = v194
		v257 = v195
		v258 = v196
		goto L44
	} else {
		goto L55
	}
L54:
	;
	v215 = v209
	v216 = v203
	v217 = v205
	v218 = v207
	goto L49
L55:
	;
	v202 = int32(1)
	v203 = v195 + v202
	v205 = v196 - v202
	v206 = int32(0)
	v207 = base.B2i32(v205 != v206)
	v209 = v194 + v202
	if v209&int32(3) == v206 {
		v215 = v209
		v216 = v203
		v217 = v205
		v218 = v207
		goto L49
	} else {
		goto L56
	}
L56:
	;
	if v205 != 0 {
		v194 = v209
		v195 = v203
		v196 = v205
		goto L53
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if v221 == int32(0) {
		v249 = v215
		v250 = v216
		v251 = v217
		goto L45
	} else {
		goto L59
	}
L59:
	;
	if base.Ui32(v217) < base.Ui32(int32(4)) {
		v249 = v215
		v250 = v216
		v251 = v217
		goto L45
	} else {
		goto L60
	}
L60:
	;
	v227 = v215
	v228 = v216
	v229 = v217
	goto L61
L61:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	v235 = int32(-2139062144)
	if (int32(16843008)-v232|v232)&v235 != v235 {
		v256 = v227
		v257 = v228
		v258 = v229
		goto L44
	} else {
		goto L63
	}
L62:
	;
	v249 = v243
	v250 = v241
	v251 = v245
	goto L45
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v232
	v240 = int32(4)
	v241 = v228 + v240
	v243 = v227 + v240
	v245 = v229 - v240
	if base.Ui32(int32(3)) < base.Ui32(v245) {
		v227 = v243
		v228 = v241
		v229 = v245
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v256 = v249
	v257 = v250
	v258 = v251
	goto L44
L66:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	*(*uint8)(unsafe.Add(mBase, uint32(v262))) = uint8(v265)
	if v265 == int32(0) {
		v276 = v261
		v277 = v262
		goto L43
	} else {
		goto L68
	}
L67:
	;
	v276 = v272
	v277 = v270
	goto L43
L68:
	;
	v269 = int32(1)
	v270 = v262 + v269
	v272 = v261 + v269
	v274 = v263 - v269
	if v274 != 0 {
		v261 = v272
		v262 = v270
		v263 = v274
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v411 = *(*int32)(unsafe.Add(mBase, _consts[866]))
	v412 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v411)+1456)) = v412
	if v48&int32(1) != 0 {
		goto L104
	} else {
		goto L105
	}
L71:
	;
	v407 = F_strlen(m, v396)
	mBase = m.M
	goto L70
L73:
	;
	goto L74
L74:
	;
	v301 = int32(1023)
	if (v19^v294)&int32(3) != 0 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v400 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v397))) = uint8(v400)
	goto L71
L76:
	;
	v381 = v376
	v382 = v377
	v383 = v378
	goto L98
L77:
	;
	if v371 == int32(0) {
		v396 = v369
		v397 = v370
		goto L75
	} else {
		goto L97
	}
L78:
	;
	v369 = v294
	v370 = v19
	v371 = v301
	goto L77
L79:
	;
	goto L80
L80:
	;
	if v294&int32(3) == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	if v338 == int32(0) {
		v396 = v335
		v397 = v336
		goto L75
	} else {
		goto L90
	}
L82:
	;
	v335 = v294
	v336 = v19
	v337 = v301
	v338 = int32(1)
	goto L81
L83:
	;
	goto L84
L84:
	;
	v314 = v294
	v315 = v19
	v316 = v301
	goto L85
L85:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314))))
	*(*uint8)(unsafe.Add(mBase, uint32(v315))) = uint8(v318)
	if v318 == int32(0) {
		v376 = v314
		v377 = v315
		v378 = v316
		goto L76
	} else {
		goto L87
	}
L86:
	;
	v335 = v329
	v336 = v323
	v337 = v325
	v338 = v327
	goto L81
L87:
	;
	v322 = int32(1)
	v323 = v315 + v322
	v325 = v316 - v322
	v326 = int32(0)
	v327 = base.B2i32(v325 != v326)
	v329 = v314 + v322
	if v329&int32(3) == v326 {
		v335 = v329
		v336 = v323
		v337 = v325
		v338 = v327
		goto L81
	} else {
		goto L88
	}
L88:
	;
	if v325 != 0 {
		v314 = v329
		v315 = v323
		v316 = v325
		goto L85
	} else {
		goto L89
	}
L89:
	;
	goto L86
L90:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	if v341 == int32(0) {
		v369 = v335
		v370 = v336
		v371 = v337
		goto L77
	} else {
		goto L91
	}
L91:
	;
	if base.Ui32(v337) < base.Ui32(int32(4)) {
		v369 = v335
		v370 = v336
		v371 = v337
		goto L77
	} else {
		goto L92
	}
L92:
	;
	v347 = v335
	v348 = v336
	v349 = v337
	goto L93
L93:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	v355 = int32(-2139062144)
	if (int32(16843008)-v352|v352)&v355 != v355 {
		v376 = v347
		v377 = v348
		v378 = v349
		goto L76
	} else {
		goto L95
	}
L94:
	;
	v369 = v363
	v370 = v361
	v371 = v365
	goto L77
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v348))) = v352
	v360 = int32(4)
	v361 = v348 + v360
	v363 = v347 + v360
	v365 = v349 - v360
	if base.Ui32(int32(3)) < base.Ui32(v365) {
		v347 = v363
		v348 = v361
		v349 = v365
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v376 = v369
	v377 = v370
	v378 = v371
	goto L76
L98:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	*(*uint8)(unsafe.Add(mBase, uint32(v382))) = uint8(v385)
	if v385 == int32(0) {
		v396 = v381
		v397 = v382
		goto L75
	} else {
		goto L100
	}
L99:
	;
	v396 = v392
	v397 = v390
	goto L75
L100:
	;
	v389 = int32(1)
	v390 = v382 + v389
	v392 = v381 + v389
	v394 = v383 - v389
	if v394 != 0 {
		v381 = v392
		v382 = v390
		v383 = v394
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L4
	} else {
		goto L182
	}
L103:
	;
	m.G0 = v19 + int32(1360)
	return v567
L104:
	;
	v417 = v49
	goto L106
L105:
	;
	v417 = v412
	goto L106
L106:
	;
	if v417 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v420 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v420)
	v567 = int32(0)
	goto L103
L108:
	;
	goto L109
L109:
	;
	v423 = int64(0)
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v411)+1464))
	v426 = base.B2i32(v424 == v423)
	if v424 == v423 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v427 = v423
	goto L112
L111:
	;
	v427 = v424
	goto L112
L112:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v411)+1464)) = v427
	v432 = F_get_call_result_type(m, l0, int32(0), v19+int32(1356))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	if v432 != int32(1) {
		goto L102
	} else {
		goto L114
	}
L114:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1356))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)))
	v440 = F_palloc0(m, v437<<(uint(int32(2))%32))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1356))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	v444 = F_palloc0(m, v443)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = v49
	v448 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	v450 = F_has_privs_of_role(m, v448, int32(3375))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L4
	} else {
		goto L118
	}
L117:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1356))
	v558 = F_heap_form_tuple(m, v557, v440, v444)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L4
	} else {
		goto L180
	}
L118:
	;
	if v450 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v454 = int32(1)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1356))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	v462 = F__emscripten_memset_bulkmem(m, v444+v454, base.I32_extend8_s(v454), v458-v454)
	mBase = m.M
	goto L122
L120:
	;
	goto L121
L121:
	;
	if base.Ui32(v47) <= base.Ui32(int32(5)) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	goto L117
L123:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v47<<(uint(int32(2))%32))+uint32(_consts[888])))
	v471 = v469
	goto L125
L124:
	;
	v471 = int32(528048)
	goto L125
L125:
	;
	v472 = F_cstring_to_text(m, v471)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+4)) = v472
	if v46 == int64(0) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+12)) = v45
	if v424 == v423 {
		goto L133
	} else {
		goto L134
	}
L128:
	;
	v477 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v444)+2)) = uint8(v477)
	goto L127
L129:
	;
	goto L130
L130:
	;
	v479 = F_Int64GetDatum(m, v46)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+8)) = v479
	goto L127
L132:
	;
	if v44 == int64(0) {
		goto L138
	} else {
		goto L139
	}
L133:
	;
	v483 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v444)+4)) = uint8(v483)
	goto L132
L134:
	;
	goto L135
L135:
	;
	v485 = F_Int64GetDatum(m, v424)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+16)) = v485
	goto L132
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+24)) = v43
	if v42 == int64(0) {
		goto L143
	} else {
		goto L144
	}
L138:
	;
	v490 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v444)+5)) = uint8(v490)
	goto L137
L139:
	;
	goto L140
L140:
	;
	v492 = F_Int64GetDatum(m, v44)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+20)) = v492
	goto L137
L142:
	;
	if v41 == int64(0) {
		goto L148
	} else {
		goto L149
	}
L143:
	;
	v498 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v444)+7)) = uint8(v498)
	goto L142
L144:
	;
	goto L145
L145:
	;
	v500 = F_Int64GetDatum(m, v42)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+28)) = v500
	goto L142
L147:
	;
	if v40 == int64(0) {
		goto L153
	} else {
		goto L154
	}
L148:
	;
	v505 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v444)+8)) = uint8(v505)
	goto L147
L149:
	;
	goto L150
L150:
	;
	v507 = F_Int64GetDatum(m, v41)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+32)) = v507
	goto L147
L152:
	;
	if v39 == int64(0) {
		goto L158
	} else {
		goto L159
	}
L153:
	;
	v512 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v444)+9)) = uint8(v512)
	goto L152
L154:
	;
	goto L155
L155:
	;
	v514 = F_Int64GetDatum(m, v40)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+36)) = v514
	goto L152
L157:
	;
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1024)))
	if v524 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L158:
	;
	v519 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v444)+10)) = uint8(v519)
	goto L157
L159:
	;
	goto L160
L160:
	;
	v521 = F_Int64GetDatum(m, v39)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+40)) = v521
	goto L157
L162:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1088)))
	if v534 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L163:
	;
	v527 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v444)+11)) = uint8(v527)
	goto L162
L164:
	;
	goto L165
L165:
	;
	v531 = F_cstring_to_text(m, v19+int32(1024))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+44)) = v531
	goto L162
L167:
	;
	if v292 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L168:
	;
	v537 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v444)+12)) = uint8(v537)
	goto L167
L169:
	;
	goto L170
L170:
	;
	v541 = F_cstring_to_text(m, v19+int32(1088))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L4
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+48)) = v541
	goto L167
L172:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v549 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L173:
	;
	v546 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v444)+13)) = uint8(v546)
	goto L172
L174:
	;
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+52)) = v292
	goto L172
L176:
	;
	v552 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v444)+14)) = uint8(v552)
	goto L117
L177:
	;
	goto L178
L178:
	;
	v554 = F_cstring_to_text(m, v19)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+56)) = v554
	goto L117
L180:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v558)+16))
	v561 = F_HeapTupleHeaderGetDatum(m, v560)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	v567 = v561
	goto L103
L182:
	;
	F_errmsg_internal(m, int32(367510), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	F_errfinish(m, int32(495327), int32(1451), int32(215017))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_stat_reset_replication_slot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v4 == int32(1) {
		F_pgstat_reset_of_kind(m, int32(4))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = F_text_to_cstring(m, v15)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = m.G0
				v21 = v19 - int32(16)
				m.G0 = v21
				v24 = *(*int32)(unsafe.Add(mBase, _consts[86]))
				v28 = F_LWLockAcquire(m, v24+int32(4736), int32(1))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v31 = F_SearchNamedReplicationSlot(m, v17, int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if v31 != 0 {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+88))
							if v33 != 0 {
								v37 = *(*int32)(unsafe.Add(mBase, _consts[853]))
								v40 = base.I32_div_s(v31-v37, int32(288))
								F_pgstat_reset(m, int32(4), int32(0), base.I64_extend_i32_s(v40))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, _consts[86]))
									F_LWLockRelease(m, v45+int32(4736))
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										m.G0 = v21 + int32(16)
										return int32(0)
									}
								}
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, _consts[86]))
								F_LWLockRelease(m, v45+int32(4736))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									m.G0 = v21 + int32(16)
									return int32(0)
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v21))) = v17
									F_errmsg(m, int32(70656), v21)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(493165), int32(57), int32(84688))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
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
		}
	}
}
func F_pg_stat_reset_subscription_stats(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v7 == int32(1) {
		F_pgstat_reset_of_kind(m, int32(5))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			m.G0 = v5 + int32(16)
			return int32(0)
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v15 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(0)
					F_errmsg(m, int32(55893), v5)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(494750), int32(2039), int32(125880))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
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
			F_pgstat_reset(m, int32(5), int32(0), base.I64_extend_i32_u(v15))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				m.G0 = v5 + int32(16)
				return int32(0)
			}
		}
	}
}
