package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
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
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v26 = v24 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = v26
				v28 = F_pgstat_fetch_stat_numbackends(m)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					if v26 <= v28 {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
						v32 = F_pgstat_get_local_beentry_by_index(m, v31)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
							*(*int64)(unsafe.Add(mBase, uint32(v22))) = v34 + int64(1)
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = int32(1)
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v32)+408))
							return v41
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = int32(2)
							v48 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v48)
							return int32(0)
						}
					}
				}
			}
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		v26 = v24 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v23))) = v26
		v28 = F_pgstat_fetch_stat_numbackends(m)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			if v26 <= v28 {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v32 = F_pgstat_get_local_beentry_by_index(m, v31)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
					*(*int64)(unsafe.Add(mBase, uint32(v22))) = v34 + int64(1)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = int32(1)
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v32)+408))
					return v41
				}
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = int32(2)
					v48 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v48)
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
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pgstat_get_beentry_by_proc_number(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			v28 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_xact_start[0]))
			v14 = F_has_privs_of_role(m, v12, int32(3375))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 != 0 {
					v21 = *(*int64)(unsafe.Add(mBase, uint32(v5)+24))
					if v21 == int64(0) {
						v28 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
						return int32(0)
					} else {
						v24 = F_Int64GetDatum(m, v21)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return v24
						}
					}
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_xact_start[0]))
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)+52))
					v19 = F_has_privs_of_role(m, v17, v18)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						if v19 != 0 {
							v21 = *(*int64)(unsafe.Add(mBase, uint32(v5)+24))
							if v21 == int64(0) {
								v28 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
								return int32(0)
							} else {
								v24 = F_Int64GetDatum(m, v21)
								mBase = m.M
								v25 = m.ExcPending
								if v25 != 0 {
									return int32(0)
								} else {
									return v24
								}
							}
						} else {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
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
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_db_checksum_failures[0]))
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v313 int64
	_ = v313
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
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
	var v343 int64
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int64
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int64
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int64
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int64
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int64
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int64
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int64
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int64
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int64
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int64
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int64
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int64
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int64
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int64
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int64
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int64
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int64
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int64
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int64
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int64
	_ = v423
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	v8 = m.G0
	v10 = v8 - int32(144)
	m.G0 = v10
	v12 = F_pgstat_fetch_stat_numbackends(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = F_text_to_cstring(m, v17)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v26 = v19
	v27 = int32(_a_F_pg_stat_get_progress_info_0)
	goto L8
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L139
	}
L6:
	;
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L99
	}
L7:
	;
	if v64 == int32(0) {
		v290 = int32(1)
		goto L6
	} else {
		goto L20
	}
L8:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v30 == v31 {
		v53 = v30
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v64 = int32(0)
	goto L7
L10:
	;
	v55 = int32(1)
	if v53 != 0 {
		v26 = v26 + v55
		v27 = v27 + v55
		goto L8
	} else {
		goto L19
	}
L11:
	;
	if base.Ui32((v30-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v41 = v30 | int32(32)
	goto L14
L13:
	;
	v41 = v30
	goto L14
L14:
	;
	if base.Ui32((v31-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v50 = v31 | int32(32)
	goto L17
L16:
	;
	v50 = v31
	goto L17
L17:
	;
	if v41 == v50 {
		v53 = v41
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v64 = v41 - v50
	goto L7
L19:
	;
	goto L9
L20:
	;
	v70 = v19
	v71 = int32(_a_F_pg_stat_get_progress_info_1)
	goto L22
L21:
	;
	if v108 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v74 == v75 {
		v97 = v74
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v108 = int32(0)
	goto L21
L24:
	;
	v99 = int32(1)
	if v97 != 0 {
		v70 = v70 + v99
		v71 = v71 + v99
		goto L22
	} else {
		goto L33
	}
L25:
	;
	if base.Ui32((v74-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v85 = v74 | int32(32)
	goto L28
L27:
	;
	v85 = v74
	goto L28
L28:
	;
	if base.Ui32((v75-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v94 = v75 | int32(32)
	goto L31
L30:
	;
	v94 = v75
	goto L31
L31:
	;
	if v85 == v94 {
		v97 = v85
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v108 = v85 - v94
	goto L21
L33:
	;
	goto L23
L34:
	;
	v290 = int32(2)
	goto L6
L35:
	;
	goto L36
L36:
	;
	v115 = v19
	v116 = int32(_a_F_pg_stat_get_progress_info_2)
	goto L38
L37:
	;
	if v153 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v119 == v120 {
		v142 = v119
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v153 = int32(0)
	goto L37
L40:
	;
	v144 = int32(1)
	if v142 != 0 {
		v115 = v115 + v144
		v116 = v116 + v144
		goto L38
	} else {
		goto L49
	}
L41:
	;
	if base.Ui32((v119-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v130 = v119 | int32(32)
	goto L44
L43:
	;
	v130 = v119
	goto L44
L44:
	;
	if base.Ui32((v120-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v139 = v120 | int32(32)
	goto L47
L46:
	;
	v139 = v120
	goto L47
L47:
	;
	if v130 == v139 {
		v142 = v130
		goto L40
	} else {
		goto L48
	}
L48:
	;
	v153 = v130 - v139
	goto L37
L49:
	;
	goto L39
L50:
	;
	v290 = int32(3)
	goto L6
L51:
	;
	goto L52
L52:
	;
	v160 = v19
	v161 = int32(_a_F_pg_stat_get_progress_info_3)
	goto L54
L53:
	;
	if v198 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L54:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v164 == v165 {
		v187 = v164
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v198 = int32(0)
	goto L53
L56:
	;
	v189 = int32(1)
	if v187 != 0 {
		v160 = v160 + v189
		v161 = v161 + v189
		goto L54
	} else {
		goto L65
	}
L57:
	;
	if base.Ui32((v164-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v175 = v164 | int32(32)
	goto L60
L59:
	;
	v175 = v164
	goto L60
L60:
	;
	if base.Ui32((v165-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v184 = v165 | int32(32)
	goto L63
L62:
	;
	v184 = v165
	goto L63
L63:
	;
	if v175 == v184 {
		v187 = v175
		goto L56
	} else {
		goto L64
	}
L64:
	;
	v198 = v175 - v184
	goto L53
L65:
	;
	goto L55
L66:
	;
	v290 = int32(4)
	goto L6
L67:
	;
	goto L68
L68:
	;
	v205 = v19
	v206 = int32(_a_F_pg_stat_get_progress_info_4)
	goto L70
L69:
	;
	if v243 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L70:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	if v209 == v210 {
		v232 = v209
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v243 = int32(0)
	goto L69
L72:
	;
	v234 = int32(1)
	if v232 != 0 {
		v205 = v205 + v234
		v206 = v206 + v234
		goto L70
	} else {
		goto L81
	}
L73:
	;
	if base.Ui32((v209-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v220 = v209 | int32(32)
	goto L76
L75:
	;
	v220 = v209
	goto L76
L76:
	;
	if base.Ui32((v210-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v229 = v210 | int32(32)
	goto L79
L78:
	;
	v229 = v210
	goto L79
L79:
	;
	if v220 == v229 {
		v232 = v220
		goto L72
	} else {
		goto L80
	}
L80:
	;
	v243 = v220 - v229
	goto L69
L81:
	;
	goto L71
L82:
	;
	v290 = int32(5)
	goto L6
L83:
	;
	goto L84
L84:
	;
	v250 = v19
	v251 = int32(_a_F_pg_stat_get_progress_info_5)
	goto L86
L85:
	;
	if v288 != 0 {
		goto L5
	} else {
		goto L98
	}
L86:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	if v254 == v255 {
		v277 = v254
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v288 = int32(0)
	goto L85
L88:
	;
	v279 = int32(1)
	if v277 != 0 {
		v250 = v250 + v279
		v251 = v251 + v279
		goto L86
	} else {
		goto L97
	}
L89:
	;
	if base.Ui32((v254-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v265 = v254 | int32(32)
	goto L92
L91:
	;
	v265 = v254
	goto L92
L92:
	;
	if base.Ui32((v255-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v274 = v255 | int32(32)
	goto L95
L94:
	;
	v274 = v255
	goto L95
L95:
	;
	if v265 == v274 {
		v277 = v265
		goto L88
	} else {
		goto L96
	}
L96:
	;
	v288 = v265 - v274
	goto L85
L97:
	;
	goto L87
L98:
	;
	v290 = int32(6)
	goto L6
L99:
	;
	if int32(0) < v12 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v299 = v10 + int32(16) | int32(2)
	v305 = int32(1)
	goto L103
L101:
	;
	goto L102
L102:
	;
	m.G0 = v10 + int32(144)
	return int32(0)
L103:
	;
	base.MemoryFill(m, v10+int32(48), int32(0), int32(92))
	v313 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+31)) = v313
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v313
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v313
	v319 = F_pgstat_get_local_beentry_by_index(m, v305)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L105
	}
L104:
	;
	goto L102
L105:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v319)+220))
	if v290 == v321 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v323
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v319)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v325
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_progress_info[0]))
	v330 = F_has_privs_of_role(m, v328, int32(3375))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L111
	}
L107:
	;
	goto L108
L108:
	;
	v438 = v305 + int32(1)
	if v438 <= v12 {
		v305 = v438
		goto L103
	} else {
		goto L138
	}
L109:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	F_tuplestore_putvalues(m, v429, v430, v10+int32(48), v10+int32(16))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L137
	}
L110:
	;
	v423 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v299)+13)) = v423
	*(*int64)(unsafe.Add(mBase, uint32(v299)+8)) = v423
	*(*int64)(unsafe.Add(mBase, uint32(v299))) = v423
	goto L109
L111:
	;
	if v330 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_progress_info[0]))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v319)+52))
	v337 = F_has_privs_of_role(m, v335, v336)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v319)+224))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v341
	v343 = *(*int64)(unsafe.Add(mBase, uint32(v319)+232))
	v344 = F_Int64GetDatum(m, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	if v337 == int32(0) {
		goto L110
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v344
	v347 = *(*int64)(unsafe.Add(mBase, uint32(v319)+240))
	v348 = F_Int64GetDatum(m, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v348
	v351 = *(*int64)(unsafe.Add(mBase, uint32(v319)+248))
	v352 = F_Int64GetDatum(m, v351)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v352
	v355 = *(*int64)(unsafe.Add(mBase, uint32(v319)+256))
	v356 = F_Int64GetDatum(m, v355)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v356
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v319)+264))
	v360 = F_Int64GetDatum(m, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+76)) = v360
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v319)+272))
	v364 = F_Int64GetDatum(m, v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v364
	v367 = *(*int64)(unsafe.Add(mBase, uint32(v319)+280))
	v368 = F_Int64GetDatum(m, v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = v368
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v319)+288))
	v372 = F_Int64GetDatum(m, v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = v372
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v319)+296))
	v376 = F_Int64GetDatum(m, v375)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+92)) = v376
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v319)+304))
	v380 = F_Int64GetDatum(m, v379)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v380
	v383 = *(*int64)(unsafe.Add(mBase, uint32(v319)+312))
	v384 = F_Int64GetDatum(m, v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = v384
	v387 = *(*int64)(unsafe.Add(mBase, uint32(v319)+320))
	v388 = F_Int64GetDatum(m, v387)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+104)) = v388
	v391 = *(*int64)(unsafe.Add(mBase, uint32(v319)+328))
	v392 = F_Int64GetDatum(m, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+108)) = v392
	v395 = *(*int64)(unsafe.Add(mBase, uint32(v319)+336))
	v396 = F_Int64GetDatum(m, v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v396
	v399 = *(*int64)(unsafe.Add(mBase, uint32(v319)+344))
	v400 = F_Int64GetDatum(m, v399)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+116)) = v400
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v319)+352))
	v404 = F_Int64GetDatum(m, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+120)) = v404
	v407 = *(*int64)(unsafe.Add(mBase, uint32(v319)+360))
	v408 = F_Int64GetDatum(m, v407)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+124)) = v408
	v411 = *(*int64)(unsafe.Add(mBase, uint32(v319)+368))
	v412 = F_Int64GetDatum(m, v411)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+128)) = v412
	v415 = *(*int64)(unsafe.Add(mBase, uint32(v319)+376))
	v416 = F_Int64GetDatum(m, v415)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = v416
	v419 = *(*int64)(unsafe.Add(mBase, uint32(v319)+384))
	v420 = F_Int64GetDatum(m, v419)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+136)) = v420
	goto L109
L137:
	;
	goto L108
L138:
	;
	goto L104
L139:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v19
	F_errmsg(m, int32(_a_F_pg_stat_get_progress_info_6), v10)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_pg_stat_get_progress_info_7), int32(280), int32(_a_F_pg_stat_get_progress_info_8))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
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
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v27
	v29 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v29
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)) = uint8(v27)
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v29
	goto L5
L4:
	;
	if v46 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_slru[0]))
	goto L7
L7:
	;
	goto L4
L8:
	;
	v52 = v46
	v53 = v2
	goto L11
L9:
	;
	goto L10
L10:
	;
	m.G0 = v16 - int32(-64)
	return int32(0)
L11:
	;
	v63 = v53 << (uint(int32(6)) % 32)
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_pg_stat_get_slru[1])))
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_pg_stat_get_slru[2])))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_pg_stat_get_slru[3])))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_pg_stat_get_slru[4])))
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_pg_stat_get_slru[5])))
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_pg_stat_get_slru[6])))
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_pg_stat_get_slru[7])))
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_pg_stat_get_slru[8])))
	v74 = F_cstring_to_text(m, v52)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v74
	v77 = F_Int64GetDatum(m, v73)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v77
	v80 = F_Int64GetDatum(m, v72)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v80
	v83 = F_Int64GetDatum(m, v71)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v83
	v86 = F_Int64GetDatum(m, v70)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v86
	v89 = F_Int64GetDatum(m, v69)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v89
	v92 = F_Int64GetDatum(m, v68)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v92
	v95 = F_Int64GetDatum(m, v67)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v95
	v98 = F_Int64GetDatum(m, v66)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v98
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	F_tuplestore_putvalues(m, v101, v102, v14+int32(-48), v16)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v107 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v107
	v109 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v109
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v109
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v109
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v109
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)) = uint8(v107)
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v109
	v122 = v53 + int32(1)
	if base.Ui32(v122) <= base.Ui32(int32(7)) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v129 != 0 {
		v52 = v129
		v53 = v122
		goto L11
	} else {
		goto L27
	}
L24:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122<<(uint(int32(2))%32))+uint32(_c_F_pg_stat_get_slru[0])))
	v129 = v127
	goto L26
L25:
	;
	v129 = int32(0)
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
	var v25 int32
	_ = v25
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
	var v64 int32
	_ = v64
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
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
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
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
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int64
	_ = v436
	var v439 int64
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	v17 = m.G0
	v19 = v17 - int32(1360)
	m.G0 = v19
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_receiver[0]))
	v25 = base.AtomicRmwXchg32(m, v22, int32(1456), int32(1))
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_receiver[0]))
	F_s_lock(m, v27+int32(1456), int32(_a_F_pg_stat_get_wal_receiver_0), int32(1416), int32(_a_F_pg_stat_get_wal_receiver_1))
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
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_receiver[0]))
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
	v174 = v19 + int32(1088)
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_receiver[0]))
	v178 = v176 + int32(1128)
	goto L40
L7:
	;
	v170 = F_strlen(m, v159)
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
	v163 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v160))) = uint8(v163)
	goto L7
L12:
	;
	v144 = v139
	v145 = v140
	v146 = v141
	goto L33
L13:
	;
	if v134 == int32(0) {
		v159 = v132
		v160 = v133
		goto L11
	} else {
		goto L32
	}
L14:
	;
	v132 = v53
	v133 = v51
	v134 = v60
	goto L13
L15:
	;
	goto L16
L16:
	;
	v64 = int32(0)
	if base.B2i32(v53&int32(3) == v64)|int32(0) == v64 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v100 == int32(0) {
		v159 = v97
		v160 = v98
		goto L11
	} else {
		goto L26
	}
L18:
	;
	v76 = v53
	v77 = v51
	v78 = v60
	goto L21
L19:
	;
	goto L20
L20:
	;
	v97 = v53
	v98 = v51
	v99 = v60
	v100 = int32(1)
	goto L17
L21:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v80)
	if v80 == int32(0) {
		v139 = v76
		v140 = v77
		v141 = v78
		goto L12
	} else {
		goto L23
	}
L22:
	;
	v97 = v91
	v98 = v85
	v99 = v87
	v100 = v89
	goto L17
L23:
	;
	v84 = int32(1)
	v85 = v77 + v84
	v87 = v78 - v84
	v88 = int32(0)
	v89 = base.B2i32(v87 != v88)
	v91 = v76 + v84
	if v91&int32(3) == v88 {
		v97 = v91
		v98 = v85
		v99 = v87
		v100 = v89
		goto L17
	} else {
		goto L24
	}
L24:
	;
	if v87 != 0 {
		v76 = v91
		v77 = v85
		v78 = v87
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if base.B2i32(v103 == int32(0))|base.B2i32(base.Ui32(v99) < base.Ui32(int32(4))) != 0 {
		v132 = v97
		v133 = v98
		v134 = v99
		goto L13
	} else {
		goto L27
	}
L27:
	;
	v110 = v97
	v111 = v98
	v112 = v99
	goto L28
L28:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v118 = int32(-2139062144)
	if (int32(16843008)-v115|v115)&v118 != v118 {
		v139 = v110
		v140 = v111
		v141 = v112
		goto L12
	} else {
		goto L30
	}
L29:
	;
	v132 = v126
	v133 = v124
	v134 = v128
	goto L13
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = v115
	v123 = int32(4)
	v124 = v111 + v123
	v126 = v110 + v123
	v128 = v112 - v123
	if base.Ui32(int32(3)) < base.Ui32(v128) {
		v110 = v126
		v111 = v124
		v112 = v128
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v139 = v132
	v140 = v133
	v141 = v134
	goto L12
L33:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	*(*uint8)(unsafe.Add(mBase, uint32(v145))) = uint8(v148)
	if v148 == int32(0) {
		v159 = v144
		v160 = v145
		goto L11
	} else {
		goto L35
	}
L34:
	;
	v159 = v155
	v160 = v153
	goto L11
L35:
	;
	v152 = int32(1)
	v153 = v145 + v152
	v155 = v144 + v152
	v157 = v146 - v152
	if v157 != 0 {
		v144 = v155
		v145 = v153
		v146 = v157
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_receiver[0]))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+1384))
	v302 = v299 + int32(104)
	goto L71
L38:
	;
	v295 = F_strlen(m, v284)
	mBase = m.M
	goto L37
L40:
	;
	goto L41
L41:
	;
	v185 = int32(254)
	if (v174^v178)&int32(3) != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v288 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v285))) = uint8(v288)
	goto L38
L43:
	;
	v269 = v264
	v270 = v265
	v271 = v266
	goto L64
L44:
	;
	if v259 == int32(0) {
		v284 = v257
		v285 = v258
		goto L42
	} else {
		goto L63
	}
L45:
	;
	v257 = v178
	v258 = v174
	v259 = v185
	goto L44
L46:
	;
	goto L47
L47:
	;
	v189 = int32(0)
	if base.B2i32(v178&int32(3) == v189)|int32(0) == v189 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v225 == int32(0) {
		v284 = v222
		v285 = v223
		goto L42
	} else {
		goto L57
	}
L49:
	;
	v201 = v178
	v202 = v174
	v203 = v185
	goto L52
L50:
	;
	goto L51
L51:
	;
	v222 = v178
	v223 = v174
	v224 = v185
	v225 = int32(1)
	goto L48
L52:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	*(*uint8)(unsafe.Add(mBase, uint32(v202))) = uint8(v205)
	if v205 == int32(0) {
		v264 = v201
		v265 = v202
		v266 = v203
		goto L43
	} else {
		goto L54
	}
L53:
	;
	v222 = v216
	v223 = v210
	v224 = v212
	v225 = v214
	goto L48
L54:
	;
	v209 = int32(1)
	v210 = v202 + v209
	v212 = v203 - v209
	v213 = int32(0)
	v214 = base.B2i32(v212 != v213)
	v216 = v201 + v209
	if v216&int32(3) == v213 {
		v222 = v216
		v223 = v210
		v224 = v212
		v225 = v214
		goto L48
	} else {
		goto L55
	}
L55:
	;
	if v212 != 0 {
		v201 = v216
		v202 = v210
		v203 = v212
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if base.B2i32(v228 == int32(0))|base.B2i32(base.Ui32(v224) < base.Ui32(int32(4))) != 0 {
		v257 = v222
		v258 = v223
		v259 = v224
		goto L44
	} else {
		goto L58
	}
L58:
	;
	v235 = v222
	v236 = v223
	v237 = v224
	goto L59
L59:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	v243 = int32(-2139062144)
	if (int32(16843008)-v240|v240)&v243 != v243 {
		v264 = v235
		v265 = v236
		v266 = v237
		goto L43
	} else {
		goto L61
	}
L60:
	;
	v257 = v251
	v258 = v249
	v259 = v253
	goto L44
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v236))) = v240
	v248 = int32(4)
	v249 = v236 + v248
	v251 = v235 + v248
	v253 = v237 - v248
	if base.Ui32(int32(3)) < base.Ui32(v253) {
		v235 = v251
		v236 = v249
		v237 = v253
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v264 = v257
	v265 = v258
	v266 = v259
	goto L43
L64:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	*(*uint8)(unsafe.Add(mBase, uint32(v270))) = uint8(v273)
	if v273 == int32(0) {
		v284 = v269
		v285 = v270
		goto L42
	} else {
		goto L66
	}
L65:
	;
	v284 = v280
	v285 = v278
	goto L42
L66:
	;
	v277 = int32(1)
	v278 = v270 + v277
	v280 = v269 + v277
	v282 = v271 - v277
	if v282 != 0 {
		v269 = v280
		v270 = v278
		v271 = v282
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_receiver[0]))
	v424 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v423)+1456)), uint32(v424))
	if v48&int32(1) != 0 {
		goto L101
	} else {
		goto L102
	}
L69:
	;
	v419 = F_strlen(m, v408)
	mBase = m.M
	goto L68
L71:
	;
	goto L72
L72:
	;
	v309 = int32(1023)
	if (v19^v302)&int32(3) != 0 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v412 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v409))) = uint8(v412)
	goto L69
L74:
	;
	v393 = v388
	v394 = v389
	v395 = v390
	goto L95
L75:
	;
	if v383 == int32(0) {
		v408 = v381
		v409 = v382
		goto L73
	} else {
		goto L94
	}
L76:
	;
	v381 = v302
	v382 = v19
	v383 = v309
	goto L75
L77:
	;
	goto L78
L78:
	;
	v313 = int32(0)
	if base.B2i32(v302&int32(3) == v313)|int32(0) == v313 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v349 == int32(0) {
		v408 = v346
		v409 = v347
		goto L73
	} else {
		goto L88
	}
L80:
	;
	v325 = v302
	v326 = v19
	v327 = v309
	goto L83
L81:
	;
	goto L82
L82:
	;
	v346 = v302
	v347 = v19
	v348 = v309
	v349 = int32(1)
	goto L79
L83:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325))))
	*(*uint8)(unsafe.Add(mBase, uint32(v326))) = uint8(v329)
	if v329 == int32(0) {
		v388 = v325
		v389 = v326
		v390 = v327
		goto L74
	} else {
		goto L85
	}
L84:
	;
	v346 = v340
	v347 = v334
	v348 = v336
	v349 = v338
	goto L79
L85:
	;
	v333 = int32(1)
	v334 = v326 + v333
	v336 = v327 - v333
	v337 = int32(0)
	v338 = base.B2i32(v336 != v337)
	v340 = v325 + v333
	if v340&int32(3) == v337 {
		v346 = v340
		v347 = v334
		v348 = v336
		v349 = v338
		goto L79
	} else {
		goto L86
	}
L86:
	;
	if v336 != 0 {
		v325 = v340
		v326 = v334
		v327 = v336
		goto L83
	} else {
		goto L87
	}
L87:
	;
	goto L84
L88:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346))))
	if base.B2i32(v352 == int32(0))|base.B2i32(base.Ui32(v348) < base.Ui32(int32(4))) != 0 {
		v381 = v346
		v382 = v347
		v383 = v348
		goto L75
	} else {
		goto L89
	}
L89:
	;
	v359 = v346
	v360 = v347
	v361 = v348
	goto L90
L90:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	v367 = int32(-2139062144)
	if (int32(16843008)-v364|v364)&v367 != v367 {
		v388 = v359
		v389 = v360
		v390 = v361
		goto L74
	} else {
		goto L92
	}
L91:
	;
	v381 = v375
	v382 = v373
	v383 = v377
	goto L75
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v360))) = v364
	v372 = int32(4)
	v373 = v360 + v372
	v375 = v359 + v372
	v377 = v361 - v372
	if base.Ui32(int32(3)) < base.Ui32(v377) {
		v359 = v375
		v360 = v373
		v361 = v377
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v388 = v381
	v389 = v382
	v390 = v383
	goto L74
L95:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393))))
	*(*uint8)(unsafe.Add(mBase, uint32(v394))) = uint8(v397)
	if v397 == int32(0) {
		v408 = v393
		v409 = v394
		goto L73
	} else {
		goto L97
	}
L96:
	;
	v408 = v404
	v409 = v402
	goto L73
L97:
	;
	v401 = int32(1)
	v402 = v394 + v401
	v404 = v393 + v401
	v406 = v395 - v401
	if v406 != 0 {
		v393 = v404
		v394 = v402
		v395 = v406
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L4
	} else {
		goto L176
	}
L100:
	;
	m.G0 = v19 + int32(1360)
	return v582
L101:
	;
	v430 = v49
	goto L103
L102:
	;
	v430 = v424
	goto L103
L103:
	;
	if v430 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v433 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v433)
	v582 = int32(0)
	goto L100
L105:
	;
	goto L106
L106:
	;
	v436 = int64(0)
	v439 = base.AtomicRmwCmpxchg64(m, v423, int32(1464), v436, v436)
	v443 = F_get_call_result_type(m, l0, int32(0), v19+int32(1356))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	if v443 != int32(1) {
		goto L99
	} else {
		goto L108
	}
L108:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1356))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	v451 = F_palloc0(m, v448<<(uint(int32(2))%32))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1356))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	v455 = F_palloc0(m, v454)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451))) = v49
	v459 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_receiver[1]))
	v461 = F_has_privs_of_role(m, v459, int32(3375))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L4
	} else {
		goto L112
	}
L111:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1356))
	v572 = F_heap_form_tuple(m, v571, v451, v455)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L4
	} else {
		goto L174
	}
L112:
	;
	if v461 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1356))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)))
	v468 = v466 - int32(1)
	if v468 == int32(0) {
		goto L111
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if base.Ui32(v47) <= base.Ui32(int32(5)) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v471 = int32(1)
	base.MemoryFill(m, v455+v471, v471, v468)
	goto L111
L117:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v47<<(uint(int32(2))%32))+uint32(_c_F_pg_stat_get_wal_receiver[2])))
	v483 = v481
	goto L119
L118:
	;
	v483 = int32(_a_F_pg_stat_get_wal_receiver_2)
	goto L119
L119:
	;
	v484 = F_cstring_to_text(m, v483)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+4)) = v484
	if v46 == int64(0) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+12)) = v45
	if v439 == int64(0) {
		goto L127
	} else {
		goto L128
	}
L122:
	;
	v489 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v455)+2)) = uint8(v489)
	goto L121
L123:
	;
	goto L124
L124:
	;
	v491 = F_Int64GetDatum(m, v46)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+8)) = v491
	goto L121
L126:
	;
	if v44 == int64(0) {
		goto L132
	} else {
		goto L133
	}
L127:
	;
	v495 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v455)+4)) = uint8(v495)
	goto L126
L128:
	;
	goto L129
L129:
	;
	v497 = F_Int64GetDatum(m, v439)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+16)) = v497
	goto L126
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+24)) = v43
	if v42 == int64(0) {
		goto L137
	} else {
		goto L138
	}
L132:
	;
	v502 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v455)+5)) = uint8(v502)
	goto L131
L133:
	;
	goto L134
L134:
	;
	v504 = F_Int64GetDatum(m, v44)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+20)) = v504
	goto L131
L136:
	;
	if v41 == int64(0) {
		goto L142
	} else {
		goto L143
	}
L137:
	;
	v510 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v455)+7)) = uint8(v510)
	goto L136
L138:
	;
	goto L139
L139:
	;
	v512 = F_Int64GetDatum(m, v42)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+28)) = v512
	goto L136
L141:
	;
	if v40 == int64(0) {
		goto L147
	} else {
		goto L148
	}
L142:
	;
	v517 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v455)+8)) = uint8(v517)
	goto L141
L143:
	;
	goto L144
L144:
	;
	v519 = F_Int64GetDatum(m, v41)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+32)) = v519
	goto L141
L146:
	;
	if v39 == int64(0) {
		goto L152
	} else {
		goto L153
	}
L147:
	;
	v524 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v455)+9)) = uint8(v524)
	goto L146
L148:
	;
	goto L149
L149:
	;
	v526 = F_Int64GetDatum(m, v40)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+36)) = v526
	goto L146
L151:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1024)))
	if v536 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L152:
	;
	v531 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v455)+10)) = uint8(v531)
	goto L151
L153:
	;
	goto L154
L154:
	;
	v533 = F_Int64GetDatum(m, v39)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+40)) = v533
	goto L151
L156:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1088)))
	if v546 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L157:
	;
	v539 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v455)+11)) = uint8(v539)
	goto L156
L158:
	;
	goto L159
L159:
	;
	v543 = F_cstring_to_text(m, v19+int32(1024))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+44)) = v543
	goto L156
L161:
	;
	if v300 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L162:
	;
	v549 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v455)+12)) = uint8(v549)
	goto L161
L163:
	;
	goto L164
L164:
	;
	v553 = F_cstring_to_text(m, v19+int32(1088))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+48)) = v553
	goto L161
L166:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v561 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L167:
	;
	v558 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v455)+13)) = uint8(v558)
	goto L166
L168:
	;
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+52)) = v300
	goto L166
L170:
	;
	v564 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v455)+14)) = uint8(v564)
	goto L111
L171:
	;
	goto L172
L172:
	;
	v566 = F_cstring_to_text(m, v19)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+56)) = v566
	goto L111
L174:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v572)+16))
	v575 = F_HeapTupleHeaderGetDatum(m, v574)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L4
	} else {
		goto L175
	}
L175:
	;
	v582 = v575
	goto L100
L176:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_get_wal_receiver_3), int32(0))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(_a_F_pg_stat_get_wal_receiver_0), int32(1451), int32(_a_F_pg_stat_get_wal_receiver_1))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
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
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_replication_slot[0]))
				v28 = F_LWLockAcquire(m, v24+int32(_a_F_pg_stat_reset_replication_slot_0), int32(1))
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
								v37 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_replication_slot[1]))
								v40 = base.I32_div_s(v31-v37, int32(288))
								F_pgstat_reset(m, int32(4), int32(0), base.I64_extend_i32_s(v40))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_replication_slot[0]))
									F_LWLockRelease(m, v45+int32(_a_F_pg_stat_reset_replication_slot_0))
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
								v45 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_replication_slot[0]))
								F_LWLockRelease(m, v45+int32(_a_F_pg_stat_reset_replication_slot_0))
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
									F_errmsg(m, int32(_a_F_pg_stat_reset_replication_slot_1), v21)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_pg_stat_reset_replication_slot_2), int32(57), int32(_a_F_pg_stat_reset_replication_slot_3))
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
					F_errmsg(m, int32(_a_F_pg_stat_reset_subscription_stats_0), v5)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_stat_reset_subscription_stats_1), int32(2039), int32(_a_F_pg_stat_reset_subscription_stats_2))
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
func F_pg_stat_statements_1_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pg_stat_statements_internal(m, l0, int32(2), base.B2i32(v3 != int32(0)))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pg_stat_statements_reset_1_11(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v9 = F_entry_reset(m, v2, v3, v5, base.B2i32(v6 != int32(0)))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = F_Int64GetDatum(m, v9)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return v13
		}
	}
}
