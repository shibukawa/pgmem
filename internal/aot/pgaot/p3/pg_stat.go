package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_pg_stat_clear_snapshot(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	F_pgstat_clear_snapshot(m)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pg_stat_file_1arg(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_pg_stat_file(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_pg_stat_get_autoanalyze_count(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+176))
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
func F_pg_stat_get_backend_io(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v17 = F_pgstat_fetch_stat_backend_by_pid(m, v14, v6+int32(12))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v17 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
				F_pg_stat_io_build_tuples(m, v13, v17+int32(8), v21, v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return int32(0)
				}
			} else {
				m.G0 = v6 + int32(16)
				return int32(0)
			}
		}
	}
}
func F_pg_stat_get_backend_subxact(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+6)) = uint16(v2)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_CreateTemplateTupleDesc(m, int32(2))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		F_TupleDescInitEntry(m, v14, int32(1), int32(_a_F_pg_stat_get_backend_subxact_0), int32(23), int32(-1), int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			F_TupleDescInitEntry(m, v14, int32(2), int32(_a_F_pg_stat_get_backend_subxact_1), int32(16), int32(-1), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = F_BlessTupleDesc(m, v14)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = F_pgstat_get_beentry_by_proc_number(m, v12)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						if v34 != 0 {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+420))
							*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v36
							v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+424)))
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v38
						} else {
							v40 = int32(257)
							*(*uint16)(unsafe.Add(mBase, uint32(v6)+6)) = uint16(v40)
						}
						v46 = F_heap_form_tuple(m, v14, v6+int32(8), v6+int32(6))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
							v49 = F_HeapTupleHeaderGetDatum(m, v48)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								m.G0 = v6 + int32(16)
								return v49
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_stat_get_backend_userid(m *base.Module, l0 int32) int32 {
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
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)+52))
			return v14
		}
	}
}
func F_pg_stat_get_backend_wait_event(m *base.Module, l0 int32) int32 {
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
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
	v82 = F_cstring_to_text(m, v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L32
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
	v80 = int32(_a_F_pg_stat_get_backend_wait_event_0)
	goto L1
L5:
	;
	goto L6
L6:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_wait_event[0]))
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
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_wait_event[0]))
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
	v23 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_backend_wait_event_1))
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
	v76 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
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
	v72 = F_pgstat_get_wait_event(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L2
	} else {
		goto L30
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
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_wait_event[1]))
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
	if v72 != 0 {
		v80 = v72
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L13
L32:
	;
	return v82
}
func F_pg_stat_get_buf_alloc(m *base.Module, l0 int32) int32 {
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
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+16))
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
func F_pg_stat_get_checkpointer_write_time(m *base.Module, l0 int32) int32 {
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
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+48))
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
func F_pg_stat_get_db_blk_read_time(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+168))
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
func F_pg_stat_get_db_conflict_lock(m *base.Module, l0 int32) int32 {
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
func F_pg_stat_get_db_conflict_logicalslot(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+104))
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
func F_pg_stat_get_db_session_time(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+192))
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
func F_pg_stat_get_db_stat_reset_time(m *base.Module, l0 int32) int32 {
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
	v5 = F_pgstat_fetch_stat_dbentry(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)+256))
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
func F_pg_stat_get_db_temp_bytes(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+136))
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
func F_pg_stat_get_db_tuples_inserted(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+48))
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
func F_pg_stat_get_recovery_prefetch(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v6)+8)) = uint16(v14)
		v16 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v16
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery_prefetch[0]))
		v23 = base.AtomicRmwCmpxchg64(m, v19, v14, v16, v16)
		v24 = F_Int64GetDatum(m, v23)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v24
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery_prefetch[0]))
			v29 = int64(0)
			v32 = base.AtomicRmwCmpxchg64(m, v28, int32(8), v29, v29)
			v33 = F_Int64GetDatum(m, v32)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v33
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery_prefetch[0]))
				v38 = int64(0)
				v41 = base.AtomicRmwCmpxchg64(m, v37, int32(16), v38, v38)
				v42 = F_Int64GetDatum(m, v41)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v42
					v46 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery_prefetch[0]))
					v47 = int64(0)
					v50 = base.AtomicRmwCmpxchg64(m, v46, int32(24), v47, v47)
					v51 = F_Int64GetDatum(m, v50)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v51
						v55 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery_prefetch[0]))
						v56 = int64(0)
						v59 = base.AtomicRmwCmpxchg64(m, v55, int32(32), v56, v56)
						v60 = F_Int64GetDatum(m, v59)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v60
							v64 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery_prefetch[0]))
							v65 = int64(0)
							v68 = base.AtomicRmwCmpxchg64(m, v64, int32(40), v65, v65)
							v69 = F_Int64GetDatum(m, v68)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v69
								v73 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery_prefetch[0]))
								v74 = int64(0)
								v77 = base.AtomicRmwCmpxchg64(m, v73, int32(48), v74, v74)
								v78 = F_Int64GetDatum(m, v77)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v78
									v82 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_recovery_prefetch[0]))
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+56))
									*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v83
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+60))
									*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v85
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)+64))
									*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = v87
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
									F_tuplestore_putvalues(m, v89, v90, v4+int32(-48), v6)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										m.G0 = v6 - int32(-64)
										return int32(0)
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
func F_pg_stat_get_replication_slot(m *base.Module, l0 int32) int32 {
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
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
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
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
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int64
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int64
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int64
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int64
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int64
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int64
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int64
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int64
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int64
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(272)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v8)+192)) = v15
		*(*int64)(unsafe.Add(mBase, uint32(v8)+184)) = v15
		*(*int64)(unsafe.Add(mBase, uint32(v8)+176)) = v15
		*(*int64)(unsafe.Add(mBase, uint32(v8)+168)) = v15
		*(*int64)(unsafe.Add(mBase, uint32(v8)+160)) = v15
		v25 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v8)+152)) = uint16(v25)
		*(*int64)(unsafe.Add(mBase, uint32(v8)+144)) = v15
		v30 = F_CreateTemplateTupleDesc(m, int32(10))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			F_TupleDescInitEntry(m, v30, int32(1), int32(_a_F_pg_stat_get_replication_slot_0), int32(25), int32(-1), int32(0))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				F_TupleDescInitEntry(m, v30, int32(2), int32(_a_F_pg_stat_get_replication_slot_1), int32(20), int32(-1), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_TupleDescInitEntry(m, v30, int32(3), int32(_a_F_pg_stat_get_replication_slot_2), int32(20), int32(-1), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						F_TupleDescInitEntry(m, v30, int32(4), int32(_a_F_pg_stat_get_replication_slot_3), int32(20), int32(-1), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_TupleDescInitEntry(m, v30, int32(5), int32(_a_F_pg_stat_get_replication_slot_4), int32(20), int32(-1), int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_TupleDescInitEntry(m, v30, int32(6), int32(_a_F_pg_stat_get_replication_slot_5), int32(20), int32(-1), int32(0))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									F_TupleDescInitEntry(m, v30, int32(7), int32(_a_F_pg_stat_get_replication_slot_6), int32(20), int32(-1), int32(0))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										F_TupleDescInitEntry(m, v30, int32(8), int32(_a_F_pg_stat_get_replication_slot_7), int32(20), int32(-1), int32(0))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											F_TupleDescInitEntry(m, v30, int32(9), int32(_a_F_pg_stat_get_replication_slot_8), int32(20), int32(-1), int32(0))
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												F_TupleDescInitEntry(m, v30, int32(10), int32(_a_F_pg_stat_get_replication_slot_9), int32(1184), int32(-1), int32(0))
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return int32(0)
												} else {
													v102 = F_BlessTupleDesc(m, v30)
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return int32(0)
													} else {
														v106 = F_text_to_cstring(m, v11)
														mBase = m.M
														v107 = m.ExcPending
														if v107 != 0 {
															return int32(0)
														} else {
															v109 = F_strncpy(m, v8+int32(208), v106, int32(64))
															mBase = m.M
															v110 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v109)+63)) = uint8(v110)
															v112 = *(*int64)(unsafe.Add(mBase, uint32(v8)+264))
															*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = v112
															v114 = *(*int64)(unsafe.Add(mBase, uint32(v8)+256))
															*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = v114
															v116 = *(*int64)(unsafe.Add(mBase, uint32(v8)+248))
															*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = v116
															v118 = *(*int64)(unsafe.Add(mBase, uint32(v8)+240))
															*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v118
															v120 = *(*int64)(unsafe.Add(mBase, uint32(v8)+232))
															*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v120
															v122 = *(*int64)(unsafe.Add(mBase, uint32(v8)+224))
															*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v122
															v124 = *(*int64)(unsafe.Add(mBase, uint32(v8)+216))
															*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v124
															v126 = *(*int64)(unsafe.Add(mBase, uint32(v8)+208))
															*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v126
															v131 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_replication_slot[0]))
															v135 = F_LWLockAcquire(m, v131+int32(_a_F_pg_stat_get_replication_slot_10), int32(1))
															mBase = m.M
															v136 = m.ExcPending
															if v136 != 0 {
																return int32(0)
															} else {
																v138 = F_SearchNamedReplicationSlot(m, v8+int32(8), int32(0))
																mBase = m.M
																v139 = m.ExcPending
																if v139 != 0 {
																	return int32(0)
																} else {
																	if v138 == int32(0) {
																		v155 = v2
																		v157 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_replication_slot[0]))
																		F_LWLockRelease(m, v157+int32(_a_F_pg_stat_get_replication_slot_10))
																		mBase = m.M
																		v161 = m.ExcPending
																		if v161 != 0 {
																			return int32(0)
																		} else {
																			if v155 == int32(0) {
																				v164 = int32(72)
																				v165 = v8 + v164
																				base.MemoryFill(m, v165, int32(0), v164)
																				v170 = v165
																			} else {
																				v170 = v155
																			}
																			v173 = F_cstring_to_text(m, v8+int32(208))
																			mBase = m.M
																			v174 = m.ExcPending
																			if v174 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v8)+160)) = v173
																				v176 = *(*int64)(unsafe.Add(mBase, uint32(v170)))
																				v177 = F_Int64GetDatum(m, v176)
																				mBase = m.M
																				v178 = m.ExcPending
																				if v178 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v8)+164)) = v177
																					v180 = *(*int64)(unsafe.Add(mBase, uint32(v170)+8))
																					v181 = F_Int64GetDatum(m, v180)
																					mBase = m.M
																					v182 = m.ExcPending
																					if v182 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v8)+168)) = v181
																						v184 = *(*int64)(unsafe.Add(mBase, uint32(v170)+16))
																						v185 = F_Int64GetDatum(m, v184)
																						mBase = m.M
																						v186 = m.ExcPending
																						if v186 != 0 {
																							return int32(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v8)+172)) = v185
																							v188 = *(*int64)(unsafe.Add(mBase, uint32(v170)+24))
																							v189 = F_Int64GetDatum(m, v188)
																							mBase = m.M
																							v190 = m.ExcPending
																							if v190 != 0 {
																								return int32(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v8)+176)) = v189
																								v192 = *(*int64)(unsafe.Add(mBase, uint32(v170)+32))
																								v193 = F_Int64GetDatum(m, v192)
																								mBase = m.M
																								v194 = m.ExcPending
																								if v194 != 0 {
																									return int32(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v8)+180)) = v193
																									v196 = *(*int64)(unsafe.Add(mBase, uint32(v170)+40))
																									v197 = F_Int64GetDatum(m, v196)
																									mBase = m.M
																									v198 = m.ExcPending
																									if v198 != 0 {
																										return int32(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v197
																										v200 = *(*int64)(unsafe.Add(mBase, uint32(v170)+48))
																										v201 = F_Int64GetDatum(m, v200)
																										mBase = m.M
																										v202 = m.ExcPending
																										if v202 != 0 {
																											return int32(0)
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v201
																											v204 = *(*int64)(unsafe.Add(mBase, uint32(v170)+56))
																											v205 = F_Int64GetDatum(m, v204)
																											mBase = m.M
																											v206 = m.ExcPending
																											if v206 != 0 {
																												return int32(0)
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(v8)+192)) = v205
																												v208 = *(*int64)(unsafe.Add(mBase, uint32(v170)+64))
																												if v208 == int64(0) {
																													v211 = int32(1)
																													*(*uint8)(unsafe.Add(mBase, uint32(v8)+153)) = uint8(v211)
																													v220 = F_heap_form_tuple(m, v30, v8+int32(160), v8+int32(144))
																													mBase = m.M
																													v221 = m.ExcPending
																													if v221 != 0 {
																														return int32(0)
																													} else {
																														v222 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
																														v223 = F_HeapTupleHeaderGetDatum(m, v222)
																														mBase = m.M
																														v224 = m.ExcPending
																														if v224 != 0 {
																															return int32(0)
																														} else {
																															m.G0 = v8 + int32(272)
																															return v223
																														}
																													}
																												} else {
																													v213 = F_Int64GetDatum(m, v208)
																													mBase = m.M
																													v214 = m.ExcPending
																													if v214 != 0 {
																														return int32(0)
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(v8)+196)) = v213
																														v220 = F_heap_form_tuple(m, v30, v8+int32(160), v8+int32(144))
																														mBase = m.M
																														v221 = m.ExcPending
																														if v221 != 0 {
																															return int32(0)
																														} else {
																															v222 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
																															v223 = F_HeapTupleHeaderGetDatum(m, v222)
																															mBase = m.M
																															v224 = m.ExcPending
																															if v224 != 0 {
																																return int32(0)
																															} else {
																																m.G0 = v8 + int32(272)
																																return v223
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
																	} else {
																		v143 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_replication_slot[1]))
																		v146 = base.I32_div_s(v138-v143, int32(288))
																		if v146 == int32(-1) {
																			v155 = v2
																			v157 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_replication_slot[0]))
																			F_LWLockRelease(m, v157+int32(_a_F_pg_stat_get_replication_slot_10))
																			mBase = m.M
																			v161 = m.ExcPending
																			if v161 != 0 {
																				return int32(0)
																			} else {
																				if v155 == int32(0) {
																					v164 = int32(72)
																					v165 = v8 + v164
																					base.MemoryFill(m, v165, int32(0), v164)
																					v170 = v165
																				} else {
																					v170 = v155
																				}
																				v173 = F_cstring_to_text(m, v8+int32(208))
																				mBase = m.M
																				v174 = m.ExcPending
																				if v174 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v8)+160)) = v173
																					v176 = *(*int64)(unsafe.Add(mBase, uint32(v170)))
																					v177 = F_Int64GetDatum(m, v176)
																					mBase = m.M
																					v178 = m.ExcPending
																					if v178 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v8)+164)) = v177
																						v180 = *(*int64)(unsafe.Add(mBase, uint32(v170)+8))
																						v181 = F_Int64GetDatum(m, v180)
																						mBase = m.M
																						v182 = m.ExcPending
																						if v182 != 0 {
																							return int32(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v8)+168)) = v181
																							v184 = *(*int64)(unsafe.Add(mBase, uint32(v170)+16))
																							v185 = F_Int64GetDatum(m, v184)
																							mBase = m.M
																							v186 = m.ExcPending
																							if v186 != 0 {
																								return int32(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v8)+172)) = v185
																								v188 = *(*int64)(unsafe.Add(mBase, uint32(v170)+24))
																								v189 = F_Int64GetDatum(m, v188)
																								mBase = m.M
																								v190 = m.ExcPending
																								if v190 != 0 {
																									return int32(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v8)+176)) = v189
																									v192 = *(*int64)(unsafe.Add(mBase, uint32(v170)+32))
																									v193 = F_Int64GetDatum(m, v192)
																									mBase = m.M
																									v194 = m.ExcPending
																									if v194 != 0 {
																										return int32(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v8)+180)) = v193
																										v196 = *(*int64)(unsafe.Add(mBase, uint32(v170)+40))
																										v197 = F_Int64GetDatum(m, v196)
																										mBase = m.M
																										v198 = m.ExcPending
																										if v198 != 0 {
																											return int32(0)
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v197
																											v200 = *(*int64)(unsafe.Add(mBase, uint32(v170)+48))
																											v201 = F_Int64GetDatum(m, v200)
																											mBase = m.M
																											v202 = m.ExcPending
																											if v202 != 0 {
																												return int32(0)
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v201
																												v204 = *(*int64)(unsafe.Add(mBase, uint32(v170)+56))
																												v205 = F_Int64GetDatum(m, v204)
																												mBase = m.M
																												v206 = m.ExcPending
																												if v206 != 0 {
																													return int32(0)
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(v8)+192)) = v205
																													v208 = *(*int64)(unsafe.Add(mBase, uint32(v170)+64))
																													if v208 == int64(0) {
																														v211 = int32(1)
																														*(*uint8)(unsafe.Add(mBase, uint32(v8)+153)) = uint8(v211)
																														v220 = F_heap_form_tuple(m, v30, v8+int32(160), v8+int32(144))
																														mBase = m.M
																														v221 = m.ExcPending
																														if v221 != 0 {
																															return int32(0)
																														} else {
																															v222 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
																															v223 = F_HeapTupleHeaderGetDatum(m, v222)
																															mBase = m.M
																															v224 = m.ExcPending
																															if v224 != 0 {
																																return int32(0)
																															} else {
																																m.G0 = v8 + int32(272)
																																return v223
																															}
																														}
																													} else {
																														v213 = F_Int64GetDatum(m, v208)
																														mBase = m.M
																														v214 = m.ExcPending
																														if v214 != 0 {
																															return int32(0)
																														} else {
																															*(*int32)(unsafe.Add(mBase, uint32(v8)+196)) = v213
																															v220 = F_heap_form_tuple(m, v30, v8+int32(160), v8+int32(144))
																															mBase = m.M
																															v221 = m.ExcPending
																															if v221 != 0 {
																																return int32(0)
																															} else {
																																v222 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
																																v223 = F_HeapTupleHeaderGetDatum(m, v222)
																																mBase = m.M
																																v224 = m.ExcPending
																																if v224 != 0 {
																																	return int32(0)
																																} else {
																																	m.G0 = v8 + int32(272)
																																	return v223
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
																		} else {
																			v152 = F_pgstat_fetch_entry(m, int32(4), int32(0), base.I64_extend_i32_s(v146))
																			mBase = m.M
																			v153 = m.ExcPending
																			if v153 != 0 {
																				return int32(0)
																			} else {
																				v155 = v152
																				v157 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_replication_slot[0]))
																				F_LWLockRelease(m, v157+int32(_a_F_pg_stat_get_replication_slot_10))
																				mBase = m.M
																				v161 = m.ExcPending
																				if v161 != 0 {
																					return int32(0)
																				} else {
																					if v155 == int32(0) {
																						v164 = int32(72)
																						v165 = v8 + v164
																						base.MemoryFill(m, v165, int32(0), v164)
																						v170 = v165
																					} else {
																						v170 = v155
																					}
																					v173 = F_cstring_to_text(m, v8+int32(208))
																					mBase = m.M
																					v174 = m.ExcPending
																					if v174 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v8)+160)) = v173
																						v176 = *(*int64)(unsafe.Add(mBase, uint32(v170)))
																						v177 = F_Int64GetDatum(m, v176)
																						mBase = m.M
																						v178 = m.ExcPending
																						if v178 != 0 {
																							return int32(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v8)+164)) = v177
																							v180 = *(*int64)(unsafe.Add(mBase, uint32(v170)+8))
																							v181 = F_Int64GetDatum(m, v180)
																							mBase = m.M
																							v182 = m.ExcPending
																							if v182 != 0 {
																								return int32(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v8)+168)) = v181
																								v184 = *(*int64)(unsafe.Add(mBase, uint32(v170)+16))
																								v185 = F_Int64GetDatum(m, v184)
																								mBase = m.M
																								v186 = m.ExcPending
																								if v186 != 0 {
																									return int32(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v8)+172)) = v185
																									v188 = *(*int64)(unsafe.Add(mBase, uint32(v170)+24))
																									v189 = F_Int64GetDatum(m, v188)
																									mBase = m.M
																									v190 = m.ExcPending
																									if v190 != 0 {
																										return int32(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v8)+176)) = v189
																										v192 = *(*int64)(unsafe.Add(mBase, uint32(v170)+32))
																										v193 = F_Int64GetDatum(m, v192)
																										mBase = m.M
																										v194 = m.ExcPending
																										if v194 != 0 {
																											return int32(0)
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(v8)+180)) = v193
																											v196 = *(*int64)(unsafe.Add(mBase, uint32(v170)+40))
																											v197 = F_Int64GetDatum(m, v196)
																											mBase = m.M
																											v198 = m.ExcPending
																											if v198 != 0 {
																												return int32(0)
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v197
																												v200 = *(*int64)(unsafe.Add(mBase, uint32(v170)+48))
																												v201 = F_Int64GetDatum(m, v200)
																												mBase = m.M
																												v202 = m.ExcPending
																												if v202 != 0 {
																													return int32(0)
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v201
																													v204 = *(*int64)(unsafe.Add(mBase, uint32(v170)+56))
																													v205 = F_Int64GetDatum(m, v204)
																													mBase = m.M
																													v206 = m.ExcPending
																													if v206 != 0 {
																														return int32(0)
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(v8)+192)) = v205
																														v208 = *(*int64)(unsafe.Add(mBase, uint32(v170)+64))
																														if v208 == int64(0) {
																															v211 = int32(1)
																															*(*uint8)(unsafe.Add(mBase, uint32(v8)+153)) = uint8(v211)
																															v220 = F_heap_form_tuple(m, v30, v8+int32(160), v8+int32(144))
																															mBase = m.M
																															v221 = m.ExcPending
																															if v221 != 0 {
																																return int32(0)
																															} else {
																																v222 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
																																v223 = F_HeapTupleHeaderGetDatum(m, v222)
																																mBase = m.M
																																v224 = m.ExcPending
																																if v224 != 0 {
																																	return int32(0)
																																} else {
																																	m.G0 = v8 + int32(272)
																																	return v223
																																}
																															}
																														} else {
																															v213 = F_Int64GetDatum(m, v208)
																															mBase = m.M
																															v214 = m.ExcPending
																															if v214 != 0 {
																																return int32(0)
																															} else {
																																*(*int32)(unsafe.Add(mBase, uint32(v8)+196)) = v213
																																v220 = F_heap_form_tuple(m, v30, v8+int32(160), v8+int32(144))
																																mBase = m.M
																																v221 = m.ExcPending
																																if v221 != 0 {
																																	return int32(0)
																																} else {
																																	v222 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
																																	v223 = F_HeapTupleHeaderGetDatum(m, v222)
																																	mBase = m.M
																																	v224 = m.ExcPending
																																	if v224 != 0 {
																																		return int32(0)
																																	} else {
																																		m.G0 = v8 + int32(272)
																																		return v223
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
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_stat_get_snapshot_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[0])))
	if v10 != 0 {
		v12 = int64(0)
		*(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[1])) = v12
		*(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[2])) = v12
		*(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[3])) = v12
		v21 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[4])) = uint8(v21)
		*(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[5])) = v21
		*(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[6])) = v21
		v30 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[7]))
		if v30 != 0 {
			F_MemoryContextDelete(m, v30)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[7])) = int32(0)
				F_pgstat_clear_backend_activity_snapshot(m)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v41 = int32(0)
					*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[0])) = uint8(v41)
					v45 = v7 + int32(15)
					v47 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[6]))
					if v47 == int32(2) {
						v50 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v50)
						v53 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[8]))
						v57 = v53
					} else {
						v54 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v54)
						v57 = int64(0)
					}
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
					if v58 == int32(0) {
						v61 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v61)
						v66 = int32(0)
						m.G0 = v7 + int32(16)
						return v66
					} else {
						v64 = F_Int64GetDatum(m, v57)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							v66 = v64
							m.G0 = v7 + int32(16)
							return v66
						}
					}
				}
			}
		} else {
			F_pgstat_clear_backend_activity_snapshot(m)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v41 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[0])) = uint8(v41)
				v45 = v7 + int32(15)
				v47 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[6]))
				if v47 == int32(2) {
					v50 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v50)
					v53 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[8]))
					v57 = v53
				} else {
					v54 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v54)
					v57 = int64(0)
				}
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
				if v58 == int32(0) {
					v61 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v61)
					v66 = int32(0)
					m.G0 = v7 + int32(16)
					return v66
				} else {
					v64 = F_Int64GetDatum(m, v57)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						v66 = v64
						m.G0 = v7 + int32(16)
						return v66
					}
				}
			}
		}
	} else {
		v45 = v7 + int32(15)
		v47 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[6]))
		if v47 == int32(2) {
			v50 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v50)
			v53 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_snapshot_timestamp[8]))
			v57 = v53
		} else {
			v54 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v54)
			v57 = int64(0)
		}
		v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v58 == int32(0) {
			v61 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v61)
			v66 = int32(0)
			m.G0 = v7 + int32(16)
			return v66
		} else {
			v64 = F_Int64GetDatum(m, v57)
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				v66 = v64
				m.G0 = v7 + int32(16)
				return v66
			}
		}
	}
}
func F_pg_stat_get_subscription_stats(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
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
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int64
	_ = v116
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int64
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int64
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int64
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int64
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	v2 = int32(0)
	v5 = int64(0)
	v6 = m.G0
	v8 = v6 - int32(144)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+136)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v8)+128)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v8)+120)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v8)+112)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v8)+104)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v8)+96)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v8)+87)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v8)+80)) = v5
	v30 = F_pgstat_fetch_entry(m, int32(5), v2, base.I64_extend_i32_u(v10))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int32(0)
	} else {
		v35 = F_CreateTemplateTupleDesc(m, int32(11))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			F_TupleDescInitEntry(m, v35, int32(1), int32(_a_F_pg_stat_get_subscription_stats_0), int32(26), int32(-1), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				F_TupleDescInitEntry(m, v35, int32(2), int32(_a_F_pg_stat_get_subscription_stats_1), int32(20), int32(-1), int32(0))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_TupleDescInitEntry(m, v35, int32(3), int32(_a_F_pg_stat_get_subscription_stats_2), int32(20), int32(-1), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						F_TupleDescInitEntry(m, v35, int32(4), int32(_a_F_pg_stat_get_subscription_stats_3), int32(20), int32(-1), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							F_TupleDescInitEntry(m, v35, int32(5), int32(_a_F_pg_stat_get_subscription_stats_4), int32(20), int32(-1), int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								F_TupleDescInitEntry(m, v35, int32(6), int32(_a_F_pg_stat_get_subscription_stats_5), int32(20), int32(-1), int32(0))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									F_TupleDescInitEntry(m, v35, int32(7), int32(_a_F_pg_stat_get_subscription_stats_6), int32(20), int32(-1), int32(0))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int32(0)
									} else {
										F_TupleDescInitEntry(m, v35, int32(8), int32(_a_F_pg_stat_get_subscription_stats_7), int32(20), int32(-1), int32(0))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int32(0)
										} else {
											F_TupleDescInitEntry(m, v35, int32(9), int32(_a_F_pg_stat_get_subscription_stats_8), int32(20), int32(-1), int32(0))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return int32(0)
											} else {
												F_TupleDescInitEntry(m, v35, int32(10), int32(_a_F_pg_stat_get_subscription_stats_9), int32(20), int32(-1), int32(0))
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													F_TupleDescInitEntry(m, v35, int32(11), int32(_a_F_pg_stat_get_subscription_stats_10), int32(1184), int32(-1), int32(0))
													mBase = m.M
													v113 = m.ExcPending
													if v113 != 0 {
														return int32(0)
													} else {
														v114 = F_BlessTupleDesc(m, v35)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return int32(0)
														} else {
															if v30 != 0 {
																v116 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
																v120 = v30
																v121 = v116
															} else {
																base.MemoryFill(m, v8, int32(0), int32(80))
																v120 = v8
																v121 = v5
															}
															*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = v10
															v123 = F_Int64GetDatum(m, v121)
															mBase = m.M
															v124 = m.ExcPending
															if v124 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v8)+100)) = v123
																v126 = *(*int64)(unsafe.Add(mBase, uint32(v120)+8))
																v127 = F_Int64GetDatum(m, v126)
																mBase = m.M
																v128 = m.ExcPending
																if v128 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v8)+104)) = v127
																	v130 = *(*int64)(unsafe.Add(mBase, uint32(v120)+16))
																	v131 = F_Int64GetDatum(m, v130)
																	mBase = m.M
																	v132 = m.ExcPending
																	if v132 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v8)+108)) = v131
																		v134 = *(*int64)(unsafe.Add(mBase, uint32(v120)+24))
																		v135 = F_Int64GetDatum(m, v134)
																		mBase = m.M
																		v136 = m.ExcPending
																		if v136 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = v135
																			v138 = *(*int64)(unsafe.Add(mBase, uint32(v120)+32))
																			v139 = F_Int64GetDatum(m, v138)
																			mBase = m.M
																			v140 = m.ExcPending
																			if v140 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v8)+116)) = v139
																				v142 = *(*int64)(unsafe.Add(mBase, uint32(v120)+40))
																				v143 = F_Int64GetDatum(m, v142)
																				mBase = m.M
																				v144 = m.ExcPending
																				if v144 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v8)+120)) = v143
																					v146 = *(*int64)(unsafe.Add(mBase, uint32(v120)+48))
																					v147 = F_Int64GetDatum(m, v146)
																					mBase = m.M
																					v148 = m.ExcPending
																					if v148 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v8)+124)) = v147
																						v150 = *(*int64)(unsafe.Add(mBase, uint32(v120)+56))
																						v151 = F_Int64GetDatum(m, v150)
																						mBase = m.M
																						v152 = m.ExcPending
																						if v152 != 0 {
																							return int32(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v8)+128)) = v151
																							v154 = *(*int64)(unsafe.Add(mBase, uint32(v120)+64))
																							v155 = F_Int64GetDatum(m, v154)
																							mBase = m.M
																							v156 = m.ExcPending
																							if v156 != 0 {
																								return int32(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v8)+132)) = v155
																								v158 = *(*int64)(unsafe.Add(mBase, uint32(v120)+72))
																								if v158 == int64(0) {
																									v161 = int32(1)
																									*(*uint8)(unsafe.Add(mBase, uint32(v8)+90)) = uint8(v161)
																									v170 = F_heap_form_tuple(m, v35, v8+int32(96), v8+int32(80))
																									mBase = m.M
																									v171 = m.ExcPending
																									if v171 != 0 {
																										return int32(0)
																									} else {
																										v172 = *(*int32)(unsafe.Add(mBase, uint32(v170)+16))
																										v173 = F_HeapTupleHeaderGetDatum(m, v172)
																										mBase = m.M
																										v174 = m.ExcPending
																										if v174 != 0 {
																											return int32(0)
																										} else {
																											m.G0 = v8 + int32(144)
																											return v173
																										}
																									}
																								} else {
																									v163 = F_Int64GetDatum(m, v158)
																									mBase = m.M
																									v164 = m.ExcPending
																									if v164 != 0 {
																										return int32(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v8)+136)) = v163
																										v170 = F_heap_form_tuple(m, v35, v8+int32(96), v8+int32(80))
																										mBase = m.M
																										v171 = m.ExcPending
																										if v171 != 0 {
																											return int32(0)
																										} else {
																											v172 = *(*int32)(unsafe.Add(mBase, uint32(v170)+16))
																											v173 = F_HeapTupleHeaderGetDatum(m, v172)
																											mBase = m.M
																											v174 = m.ExcPending
																											if v174 != 0 {
																												return int32(0)
																											} else {
																												m.G0 = v8 + int32(144)
																												return v173
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
		}
	}
}
func F_pg_stat_get_tuples_deleted(m *base.Module, l0 int32) int32 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v3)+48))
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
func F_pg_stat_get_wal_senders(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
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
	var v87 int32
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
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int64
	_ = v216
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	v2 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(80)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = F_SyncRepGetCandidateStandbys(m, v22+int32(76))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_senders[0]))
	if int32(0) < v35 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v39 = v22 + int32(1)
	v44 = v2
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v22 + int32(80)
	return int32(0)
L7:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_senders[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = int64(0)
	v67 = v60 + v44*int32(96)
	v68 = int32(164)
	v69 = v67 + v68
	v72 = base.AtomicRmwXchg32(m, v67, v68, int32(1))
	if v72 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	F_s_lock(m, v69, int32(_a_F_pg_stat_get_wal_senders_0), int32(3950), int32(_a_F_pg_stat_get_wal_senders_1))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v79 = v67 + int32(88)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v80 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L11
L13:
	;
	v296 = v44 + int32(1)
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_senders[0]))
	if v296 < v298 {
		v44 = v296
		goto L7
	} else {
		goto L83
	}
L14:
	;
	v83 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v69))), uint32(v83))
	goto L13
L15:
	;
	goto L16
L16:
	;
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v79)+80))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v79)+72))
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v79)+64))
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v79)+56))
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v79)+48))
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v79)+40))
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v79)+32))
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v79)+24))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v79)+8))
	v96 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v79)+76)), uint32(v96))
	if v32 <= v96 {
		v137 = int32(1)
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v80
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_senders[2]))
	v158 = F_has_privs_of_role(m, v156, int32(3375))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L26
	}
L18:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	v104 = int32(0)
	goto L19
L19:
	;
	v125 = v103 + v104*int32(48)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+36))
	if v126 != v44 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v137 = v131
	goto L17
L21:
	;
	v131 = int32(1)
	v133 = v104 + v131
	if v133 != v32 {
		v104 = v133
		goto L19
	} else {
		goto L24
	}
L22:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	if v128 != v80 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v137 = int32(0)
	goto L17
L24:
	;
	goto L20
L25:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	F_tuplestore_putvalues(m, v270, v271, v22+int32(16), v22)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L82
	}
L26:
	;
	if v158 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+7)) = int32(16843009)
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = int64(72340172838076673)
	goto L25
L28:
	;
	goto L29
L29:
	;
	if base.Ui32(v94) <= base.Ui32(int32(4)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v94<<(uint(int32(2))%32))+uint32(_c_F_pg_stat_get_wal_senders[3])))
	v172 = v170
	goto L32
L31:
	;
	v172 = int32(_a_F_pg_stat_get_wal_senders_2)
	goto L32
L32:
	;
	v173 = F_cstring_to_text(m, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v173
	if v95 == int64(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v178 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)) = uint8(v178)
	goto L36
L35:
	;
	goto L36
L36:
	;
	v180 = F_Int64GetDatum(m, v95)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v180
	if v93 == int64(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v185 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)) = uint8(v185)
	goto L40
L39:
	;
	goto L40
L40:
	;
	v187 = F_Int64GetDatum(m, v93)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v187
	if v92 == int64(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v192 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)) = uint8(v192)
	goto L44
L43:
	;
	goto L44
L44:
	;
	v194 = F_Int64GetDatum(m, v92)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v194
	if v91 == int64(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v199 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)) = uint8(v199)
	goto L48
L47:
	;
	goto L48
L48:
	;
	v201 = F_Int64GetDatum(m, v91)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v201
	if v90 < int64(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v216 = int64(0)
	if v89 < v216 {
		goto L56
	} else {
		goto L57
	}
L51:
	;
	v206 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)) = uint8(v206)
	goto L50
L52:
	;
	goto L53
L53:
	;
	v209 = F_palloc(m, int32(16))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v209)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v209))) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v209
	goto L50
L55:
	;
	if v92 == v216 {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v220 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+7)) = uint8(v220)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v223 = F_palloc(m, int32(16))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v223)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v223))) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v223
	goto L55
L60:
	;
	v231 = int32(0)
	goto L62
L61:
	;
	v231 = v87
	goto L62
L62:
	;
	if v88 < int64(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v231
	if base.B2i32(v231 == int32(0))|v137 != 0 {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	v234 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)) = uint8(v234)
	goto L63
L65:
	;
	goto L66
L66:
	;
	v237 = F_palloc(m, int32(16))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v237)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v237))) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v237
	goto L63
L68:
	;
	if v231 != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_senders[4]))
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+8)))
	if v255 != 0 {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	v258 = F_cstring_to_text(m, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L77
	}
L71:
	;
	v250 = int32(_a_F_pg_stat_get_wal_senders_3)
	goto L73
L72:
	;
	v250 = int32(_a_F_pg_stat_get_wal_senders_4)
	goto L73
L73:
	;
	v257 = v250
	goto L70
L74:
	;
	v256 = int32(_a_F_pg_stat_get_wal_senders_5)
	goto L76
L75:
	;
	v256 = int32(_a_F_pg_stat_get_wal_senders_6)
	goto L76
L76:
	;
	v257 = v256
	goto L70
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v258
	if v86 == int64(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v263 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+11)) = uint8(v263)
	goto L25
L79:
	;
	goto L80
L80:
	;
	v265 = F_Int64GetDatum(m, v86)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v265
	goto L25
L82:
	;
	goto L13
L83:
	;
	goto L8
}
func F_pg_stat_get_xact_blocks_fetched(m *base.Module, l0 int32) int32 {
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
func F_pg_stat_have_stats(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v354 int32
	_ = v354
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v399 int32
	_ = v399
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v467 int32
	_ = v467
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v489 int32
	_ = v489
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v512 int32
	_ = v512
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v534 int32
	_ = v534
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v584 int32
	_ = v584
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v633 int32
	_ = v633
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v682 int32
	_ = v682
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v731 int32
	_ = v731
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v780 int32
	_ = v780
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v829 int32
	_ = v829
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v878 int32
	_ = v878
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v901 int32
	_ = v901
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v927 int32
	_ = v927
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v950 int32
	_ = v950
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v979 int32
	_ = v979
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v1002 int32
	_ = v1002
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1037 int32
	_ = v1037
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = F_text_to_cstring(m, v8)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v24 = v12
	v25 = int32(_a_F_pg_stat_have_stats_0)
	goto L6
L4:
	;
	m.G0 = v19 + int32(16)
	if base.Ui32(v1023-int32(1)) <= base.Ui32(int32(11)) {
		goto L370
	} else {
		goto L371
	}
L5:
	;
	if v62 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v28 == v29 {
		v51 = v28
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v62 = int32(0)
	goto L5
L8:
	;
	v53 = int32(1)
	if v51 != 0 {
		v24 = v24 + v53
		v25 = v25 + v53
		goto L6
	} else {
		goto L17
	}
L9:
	;
	if base.Ui32((v28-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v39 = v28 | int32(32)
	goto L12
L11:
	;
	v39 = v28
	goto L12
L12:
	;
	if base.Ui32((v29-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v48 = v29 | int32(32)
	goto L15
L14:
	;
	v48 = v29
	goto L15
L15:
	;
	if v39 == v48 {
		v51 = v39
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v62 = v39 - v48
	goto L5
L17:
	;
	goto L7
L18:
	;
	v1023 = int32(1)
	goto L4
L19:
	;
	goto L20
L20:
	;
	v69 = v12
	v70 = int32(_a_F_pg_stat_have_stats_1)
	goto L22
L21:
	;
	if v107 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v73 == v74 {
		v96 = v73
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v107 = int32(0)
	goto L21
L24:
	;
	v98 = int32(1)
	if v96 != 0 {
		v69 = v69 + v98
		v70 = v70 + v98
		goto L22
	} else {
		goto L33
	}
L25:
	;
	if base.Ui32((v73-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v84 = v73 | int32(32)
	goto L28
L27:
	;
	v84 = v73
	goto L28
L28:
	;
	if base.Ui32((v74-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v93 = v74 | int32(32)
	goto L31
L30:
	;
	v93 = v74
	goto L31
L31:
	;
	if v84 == v93 {
		v96 = v84
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v107 = v84 - v93
	goto L21
L33:
	;
	goto L23
L34:
	;
	v1023 = int32(2)
	goto L4
L35:
	;
	goto L36
L36:
	;
	v114 = v12
	v115 = int32(_a_F_pg_stat_have_stats_2)
	goto L38
L37:
	;
	if v152 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	if v118 == v119 {
		v141 = v118
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v152 = int32(0)
	goto L37
L40:
	;
	v143 = int32(1)
	if v141 != 0 {
		v114 = v114 + v143
		v115 = v115 + v143
		goto L38
	} else {
		goto L49
	}
L41:
	;
	if base.Ui32((v118-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v129 = v118 | int32(32)
	goto L44
L43:
	;
	v129 = v118
	goto L44
L44:
	;
	if base.Ui32((v119-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v138 = v119 | int32(32)
	goto L47
L46:
	;
	v138 = v119
	goto L47
L47:
	;
	if v129 == v138 {
		v141 = v129
		goto L40
	} else {
		goto L48
	}
L48:
	;
	v152 = v129 - v138
	goto L37
L49:
	;
	goto L39
L50:
	;
	v1023 = int32(3)
	goto L4
L51:
	;
	goto L52
L52:
	;
	v159 = v12
	v160 = int32(_a_F_pg_stat_have_stats_3)
	goto L54
L53:
	;
	if v197 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L54:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if v163 == v164 {
		v186 = v163
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v197 = int32(0)
	goto L53
L56:
	;
	v188 = int32(1)
	if v186 != 0 {
		v159 = v159 + v188
		v160 = v160 + v188
		goto L54
	} else {
		goto L65
	}
L57:
	;
	if base.Ui32((v163-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v174 = v163 | int32(32)
	goto L60
L59:
	;
	v174 = v163
	goto L60
L60:
	;
	if base.Ui32((v164-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v183 = v164 | int32(32)
	goto L63
L62:
	;
	v183 = v164
	goto L63
L63:
	;
	if v174 == v183 {
		v186 = v174
		goto L56
	} else {
		goto L64
	}
L64:
	;
	v197 = v174 - v183
	goto L53
L65:
	;
	goto L55
L66:
	;
	v1023 = int32(4)
	goto L4
L67:
	;
	goto L68
L68:
	;
	v204 = v12
	v205 = int32(_a_F_pg_stat_have_stats_4)
	goto L70
L69:
	;
	if v242 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L70:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if v208 == v209 {
		v231 = v208
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v242 = int32(0)
	goto L69
L72:
	;
	v233 = int32(1)
	if v231 != 0 {
		v204 = v204 + v233
		v205 = v205 + v233
		goto L70
	} else {
		goto L81
	}
L73:
	;
	if base.Ui32((v208-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v219 = v208 | int32(32)
	goto L76
L75:
	;
	v219 = v208
	goto L76
L76:
	;
	if base.Ui32((v209-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v228 = v209 | int32(32)
	goto L79
L78:
	;
	v228 = v209
	goto L79
L79:
	;
	if v219 == v228 {
		v231 = v219
		goto L72
	} else {
		goto L80
	}
L80:
	;
	v242 = v219 - v228
	goto L69
L81:
	;
	goto L71
L82:
	;
	v1023 = int32(5)
	goto L4
L83:
	;
	goto L84
L84:
	;
	v249 = v12
	v250 = int32(_a_F_pg_stat_have_stats_5)
	goto L86
L85:
	;
	if v287 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L86:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if v253 == v254 {
		v276 = v253
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v287 = int32(0)
	goto L85
L88:
	;
	v278 = int32(1)
	if v276 != 0 {
		v249 = v249 + v278
		v250 = v250 + v278
		goto L86
	} else {
		goto L97
	}
L89:
	;
	if base.Ui32((v253-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v264 = v253 | int32(32)
	goto L92
L91:
	;
	v264 = v253
	goto L92
L92:
	;
	if base.Ui32((v254-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v273 = v254 | int32(32)
	goto L95
L94:
	;
	v273 = v254
	goto L95
L95:
	;
	if v264 == v273 {
		v276 = v264
		goto L88
	} else {
		goto L96
	}
L96:
	;
	v287 = v264 - v273
	goto L85
L97:
	;
	goto L87
L98:
	;
	v1023 = int32(6)
	goto L4
L99:
	;
	goto L100
L100:
	;
	v294 = v12
	v295 = int32(_a_F_pg_stat_have_stats_6)
	goto L102
L101:
	;
	if v332 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L102:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	if v298 == v299 {
		v321 = v298
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v332 = int32(0)
	goto L101
L104:
	;
	v323 = int32(1)
	if v321 != 0 {
		v294 = v294 + v323
		v295 = v295 + v323
		goto L102
	} else {
		goto L113
	}
L105:
	;
	if base.Ui32((v298-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v309 = v298 | int32(32)
	goto L108
L107:
	;
	v309 = v298
	goto L108
L108:
	;
	if base.Ui32((v299-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v318 = v299 | int32(32)
	goto L111
L110:
	;
	v318 = v299
	goto L111
L111:
	;
	if v309 == v318 {
		v321 = v309
		goto L104
	} else {
		goto L112
	}
L112:
	;
	v332 = v309 - v318
	goto L101
L113:
	;
	goto L103
L114:
	;
	v1023 = int32(7)
	goto L4
L115:
	;
	goto L116
L116:
	;
	v339 = v12
	v340 = int32(_a_F_pg_stat_have_stats_7)
	goto L118
L117:
	;
	if v377 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L118:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340))))
	if v343 == v344 {
		v366 = v343
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v377 = int32(0)
	goto L117
L120:
	;
	v368 = int32(1)
	if v366 != 0 {
		v339 = v339 + v368
		v340 = v340 + v368
		goto L118
	} else {
		goto L129
	}
L121:
	;
	if base.Ui32((v343-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v354 = v343 | int32(32)
	goto L124
L123:
	;
	v354 = v343
	goto L124
L124:
	;
	if base.Ui32((v344-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v363 = v344 | int32(32)
	goto L127
L126:
	;
	v363 = v344
	goto L127
L127:
	;
	if v354 == v363 {
		v366 = v354
		goto L120
	} else {
		goto L128
	}
L128:
	;
	v377 = v354 - v363
	goto L117
L129:
	;
	goto L119
L130:
	;
	v1023 = int32(8)
	goto L4
L131:
	;
	goto L132
L132:
	;
	v384 = v12
	v385 = int32(_a_F_pg_stat_have_stats_8)
	goto L134
L133:
	;
	if v422 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L134:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	if v388 == v389 {
		v411 = v388
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v422 = int32(0)
	goto L133
L136:
	;
	v413 = int32(1)
	if v411 != 0 {
		v384 = v384 + v413
		v385 = v385 + v413
		goto L134
	} else {
		goto L145
	}
L137:
	;
	if base.Ui32((v388-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v399 = v388 | int32(32)
	goto L140
L139:
	;
	v399 = v388
	goto L140
L140:
	;
	if base.Ui32((v389-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v408 = v389 | int32(32)
	goto L143
L142:
	;
	v408 = v389
	goto L143
L143:
	;
	if v399 == v408 {
		v411 = v399
		goto L136
	} else {
		goto L144
	}
L144:
	;
	v422 = v399 - v408
	goto L133
L145:
	;
	goto L135
L146:
	;
	v1023 = int32(9)
	goto L4
L147:
	;
	goto L148
L148:
	;
	v429 = v12
	v430 = int32(_a_F_pg_stat_have_stats_9)
	goto L150
L149:
	;
	if v467 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L150:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429))))
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	if v433 == v434 {
		v456 = v433
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v467 = int32(0)
	goto L149
L152:
	;
	v458 = int32(1)
	if v456 != 0 {
		v429 = v429 + v458
		v430 = v430 + v458
		goto L150
	} else {
		goto L161
	}
L153:
	;
	if base.Ui32((v433-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v444 = v433 | int32(32)
	goto L156
L155:
	;
	v444 = v433
	goto L156
L156:
	;
	if base.Ui32((v434-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v453 = v434 | int32(32)
	goto L159
L158:
	;
	v453 = v434
	goto L159
L159:
	;
	if v444 == v453 {
		v456 = v444
		goto L152
	} else {
		goto L160
	}
L160:
	;
	v467 = v444 - v453
	goto L149
L161:
	;
	goto L151
L162:
	;
	v1023 = int32(10)
	goto L4
L163:
	;
	goto L164
L164:
	;
	v474 = v12
	v475 = int32(_a_F_pg_stat_have_stats_10)
	goto L166
L165:
	;
	if v512 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L166:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
	if v478 == v479 {
		v501 = v478
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v512 = int32(0)
	goto L165
L168:
	;
	v503 = int32(1)
	if v501 != 0 {
		v474 = v474 + v503
		v475 = v475 + v503
		goto L166
	} else {
		goto L177
	}
L169:
	;
	if base.Ui32((v478-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v489 = v478 | int32(32)
	goto L172
L171:
	;
	v489 = v478
	goto L172
L172:
	;
	if base.Ui32((v479-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v498 = v479 | int32(32)
	goto L175
L174:
	;
	v498 = v479
	goto L175
L175:
	;
	if v489 == v498 {
		v501 = v489
		goto L168
	} else {
		goto L176
	}
L176:
	;
	v512 = v489 - v498
	goto L165
L177:
	;
	goto L167
L178:
	;
	v1023 = int32(11)
	goto L4
L179:
	;
	goto L180
L180:
	;
	v519 = v12
	v520 = int32(_a_F_pg_stat_have_stats_11)
	goto L182
L181:
	;
	if v557 == int32(0) {
		goto L194
	} else {
		goto L195
	}
L182:
	;
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519))))
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520))))
	if v523 == v524 {
		v546 = v523
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v557 = int32(0)
	goto L181
L184:
	;
	v548 = int32(1)
	if v546 != 0 {
		v519 = v519 + v548
		v520 = v520 + v548
		goto L182
	} else {
		goto L193
	}
L185:
	;
	if base.Ui32((v523-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v534 = v523 | int32(32)
	goto L188
L187:
	;
	v534 = v523
	goto L188
L188:
	;
	if base.Ui32((v524-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v543 = v524 | int32(32)
	goto L191
L190:
	;
	v543 = v524
	goto L191
L191:
	;
	if v534 == v543 {
		v546 = v534
		goto L184
	} else {
		goto L192
	}
L192:
	;
	v557 = v534 - v543
	goto L181
L193:
	;
	goto L183
L194:
	;
	v1023 = int32(12)
	goto L4
L195:
	;
	goto L196
L196:
	;
	v562 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	if v562 == int32(0) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L1
	} else {
		goto L366
	}
L198:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v562)))
	if v565 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)+68))
	v569 = v12
	v570 = v566
	goto L203
L200:
	;
	v613 = v562
	goto L201
L201:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v613)+4))
	if v614 != 0 {
		goto L218
	} else {
		goto L219
	}
L202:
	;
	if v607 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L203:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569))))
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v570))))
	if v573 == v574 {
		v596 = v573
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v607 = int32(0)
	goto L202
L205:
	;
	v598 = int32(1)
	if v596 != 0 {
		v569 = v569 + v598
		v570 = v570 + v598
		goto L203
	} else {
		goto L214
	}
L206:
	;
	if base.Ui32((v573-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v584 = v573 | int32(32)
	goto L209
L208:
	;
	v584 = v573
	goto L209
L209:
	;
	if base.Ui32((v574-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v593 = v574 | int32(32)
	goto L212
L211:
	;
	v593 = v574
	goto L212
L212:
	;
	if v584 == v593 {
		v596 = v584
		goto L205
	} else {
		goto L213
	}
L213:
	;
	v607 = v584 - v593
	goto L202
L214:
	;
	goto L204
L215:
	;
	v1023 = int32(24)
	goto L4
L216:
	;
	goto L217
L217:
	;
	v612 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	v613 = v612
	goto L201
L218:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v614)+68))
	v618 = v12
	v619 = v615
	goto L222
L219:
	;
	v662 = v613
	goto L220
L220:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v662)+8))
	if v663 != 0 {
		goto L237
	} else {
		goto L238
	}
L221:
	;
	if v656 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L222:
	;
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618))))
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619))))
	if v622 == v623 {
		v645 = v622
		goto L224
	} else {
		goto L225
	}
L223:
	;
	v656 = int32(0)
	goto L221
L224:
	;
	v647 = int32(1)
	if v645 != 0 {
		v618 = v618 + v647
		v619 = v619 + v647
		goto L222
	} else {
		goto L233
	}
L225:
	;
	if base.Ui32((v622-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v633 = v622 | int32(32)
	goto L228
L227:
	;
	v633 = v622
	goto L228
L228:
	;
	if base.Ui32((v623-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v642 = v623 | int32(32)
	goto L231
L230:
	;
	v642 = v623
	goto L231
L231:
	;
	if v633 == v642 {
		v645 = v633
		goto L224
	} else {
		goto L232
	}
L232:
	;
	v656 = v633 - v642
	goto L221
L233:
	;
	goto L223
L234:
	;
	v1023 = int32(25)
	goto L4
L235:
	;
	goto L236
L236:
	;
	v661 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	v662 = v661
	goto L220
L237:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v663)+68))
	v667 = v12
	v668 = v664
	goto L241
L238:
	;
	v711 = v662
	goto L239
L239:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v711)+12))
	if v712 != 0 {
		goto L256
	} else {
		goto L257
	}
L240:
	;
	if v705 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L241:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667))))
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668))))
	if v671 == v672 {
		v694 = v671
		goto L243
	} else {
		goto L244
	}
L242:
	;
	v705 = int32(0)
	goto L240
L243:
	;
	v696 = int32(1)
	if v694 != 0 {
		v667 = v667 + v696
		v668 = v668 + v696
		goto L241
	} else {
		goto L252
	}
L244:
	;
	if base.Ui32((v671-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v682 = v671 | int32(32)
	goto L247
L246:
	;
	v682 = v671
	goto L247
L247:
	;
	if base.Ui32((v672-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v691 = v672 | int32(32)
	goto L250
L249:
	;
	v691 = v672
	goto L250
L250:
	;
	if v682 == v691 {
		v694 = v682
		goto L243
	} else {
		goto L251
	}
L251:
	;
	v705 = v682 - v691
	goto L240
L252:
	;
	goto L242
L253:
	;
	v1023 = int32(26)
	goto L4
L254:
	;
	goto L255
L255:
	;
	v710 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	v711 = v710
	goto L239
L256:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v712)+68))
	v716 = v12
	v717 = v713
	goto L260
L257:
	;
	v760 = v711
	goto L258
L258:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v760)+16))
	if v761 != 0 {
		goto L275
	} else {
		goto L276
	}
L259:
	;
	if v754 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L260:
	;
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v716))))
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v717))))
	if v720 == v721 {
		v743 = v720
		goto L262
	} else {
		goto L263
	}
L261:
	;
	v754 = int32(0)
	goto L259
L262:
	;
	v745 = int32(1)
	if v743 != 0 {
		v716 = v716 + v745
		v717 = v717 + v745
		goto L260
	} else {
		goto L271
	}
L263:
	;
	if base.Ui32((v720-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v731 = v720 | int32(32)
	goto L266
L265:
	;
	v731 = v720
	goto L266
L266:
	;
	if base.Ui32((v721-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v740 = v721 | int32(32)
	goto L269
L268:
	;
	v740 = v721
	goto L269
L269:
	;
	if v731 == v740 {
		v743 = v731
		goto L262
	} else {
		goto L270
	}
L270:
	;
	v754 = v731 - v740
	goto L259
L271:
	;
	goto L261
L272:
	;
	v1023 = int32(27)
	goto L4
L273:
	;
	goto L274
L274:
	;
	v759 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	v760 = v759
	goto L258
L275:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v761)+68))
	v765 = v12
	v766 = v762
	goto L279
L276:
	;
	v809 = v760
	goto L277
L277:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v809)+20))
	if v810 != 0 {
		goto L294
	} else {
		goto L295
	}
L278:
	;
	if v803 == int32(0) {
		goto L291
	} else {
		goto L292
	}
L279:
	;
	v769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765))))
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766))))
	if v769 == v770 {
		v792 = v769
		goto L281
	} else {
		goto L282
	}
L280:
	;
	v803 = int32(0)
	goto L278
L281:
	;
	v794 = int32(1)
	if v792 != 0 {
		v765 = v765 + v794
		v766 = v766 + v794
		goto L279
	} else {
		goto L290
	}
L282:
	;
	if base.Ui32((v769-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v780 = v769 | int32(32)
	goto L285
L284:
	;
	v780 = v769
	goto L285
L285:
	;
	if base.Ui32((v770-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v789 = v770 | int32(32)
	goto L288
L287:
	;
	v789 = v770
	goto L288
L288:
	;
	if v780 == v789 {
		v792 = v780
		goto L281
	} else {
		goto L289
	}
L289:
	;
	v803 = v780 - v789
	goto L278
L290:
	;
	goto L280
L291:
	;
	v1023 = int32(28)
	goto L4
L292:
	;
	goto L293
L293:
	;
	v808 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	v809 = v808
	goto L277
L294:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v810)+68))
	v814 = v12
	v815 = v811
	goto L298
L295:
	;
	v858 = v809
	goto L296
L296:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v858)+24))
	if v859 != 0 {
		goto L313
	} else {
		goto L314
	}
L297:
	;
	if v852 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L298:
	;
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v814))))
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815))))
	if v818 == v819 {
		v841 = v818
		goto L300
	} else {
		goto L301
	}
L299:
	;
	v852 = int32(0)
	goto L297
L300:
	;
	v843 = int32(1)
	if v841 != 0 {
		v814 = v814 + v843
		v815 = v815 + v843
		goto L298
	} else {
		goto L309
	}
L301:
	;
	if base.Ui32((v818-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v829 = v818 | int32(32)
	goto L304
L303:
	;
	v829 = v818
	goto L304
L304:
	;
	if base.Ui32((v819-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v838 = v819 | int32(32)
	goto L307
L306:
	;
	v838 = v819
	goto L307
L307:
	;
	if v829 == v838 {
		v841 = v829
		goto L300
	} else {
		goto L308
	}
L308:
	;
	v852 = v829 - v838
	goto L297
L309:
	;
	goto L299
L310:
	;
	v1023 = int32(29)
	goto L4
L311:
	;
	goto L312
L312:
	;
	v857 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	v858 = v857
	goto L296
L313:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v859)+68))
	v863 = v12
	v864 = v860
	goto L317
L314:
	;
	v907 = v858
	goto L315
L315:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v907)+28))
	if v908 != 0 {
		goto L332
	} else {
		goto L333
	}
L316:
	;
	if v901 == int32(0) {
		goto L329
	} else {
		goto L330
	}
L317:
	;
	v867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v863))))
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864))))
	if v867 == v868 {
		v890 = v867
		goto L319
	} else {
		goto L320
	}
L318:
	;
	v901 = int32(0)
	goto L316
L319:
	;
	v892 = int32(1)
	if v890 != 0 {
		v863 = v863 + v892
		v864 = v864 + v892
		goto L317
	} else {
		goto L328
	}
L320:
	;
	if base.Ui32((v867-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v878 = v867 | int32(32)
	goto L323
L322:
	;
	v878 = v867
	goto L323
L323:
	;
	if base.Ui32((v868-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v887 = v868 | int32(32)
	goto L326
L325:
	;
	v887 = v868
	goto L326
L326:
	;
	if v878 == v887 {
		v890 = v878
		goto L319
	} else {
		goto L327
	}
L327:
	;
	v901 = v878 - v887
	goto L316
L328:
	;
	goto L318
L329:
	;
	v1023 = int32(30)
	goto L4
L330:
	;
	goto L331
L331:
	;
	v906 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	v907 = v906
	goto L315
L332:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v908)+68))
	v912 = v12
	v913 = v909
	goto L336
L333:
	;
	v956 = v907
	goto L334
L334:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v956)+32))
	if v958 == int32(0) {
		goto L197
	} else {
		goto L351
	}
L335:
	;
	if v950 == int32(0) {
		goto L348
	} else {
		goto L349
	}
L336:
	;
	v916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v912))))
	v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v913))))
	if v916 == v917 {
		v939 = v916
		goto L338
	} else {
		goto L339
	}
L337:
	;
	v950 = int32(0)
	goto L335
L338:
	;
	v941 = int32(1)
	if v939 != 0 {
		v912 = v912 + v941
		v913 = v913 + v941
		goto L336
	} else {
		goto L347
	}
L339:
	;
	if base.Ui32((v916-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v927 = v916 | int32(32)
	goto L342
L341:
	;
	v927 = v916
	goto L342
L342:
	;
	if base.Ui32((v917-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v936 = v917 | int32(32)
	goto L345
L344:
	;
	v936 = v917
	goto L345
L345:
	;
	if v927 == v936 {
		v939 = v927
		goto L338
	} else {
		goto L346
	}
L346:
	;
	v950 = v927 - v936
	goto L335
L347:
	;
	goto L337
L348:
	;
	v1023 = int32(31)
	goto L4
L349:
	;
	goto L350
L350:
	;
	v955 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	v956 = v955
	goto L334
L351:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v958)+68))
	v964 = v12
	v965 = v961
	goto L353
L352:
	;
	if v1002 == int32(0) {
		v1023 = int32(32)
		goto L4
	} else {
		goto L365
	}
L353:
	;
	v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v964))))
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v965))))
	if v968 == v969 {
		v991 = v968
		goto L355
	} else {
		goto L356
	}
L354:
	;
	v1002 = int32(0)
	goto L352
L355:
	;
	v993 = int32(1)
	if v991 != 0 {
		v964 = v964 + v993
		v965 = v965 + v993
		goto L353
	} else {
		goto L364
	}
L356:
	;
	if base.Ui32((v968-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v979 = v968 | int32(32)
	goto L359
L358:
	;
	v979 = v968
	goto L359
L359:
	;
	if base.Ui32((v969-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v988 = v969 | int32(32)
	goto L362
L361:
	;
	v988 = v969
	goto L362
L362:
	;
	if v979 == v988 {
		v991 = v979
		goto L355
	} else {
		goto L363
	}
L363:
	;
	v1002 = v979 - v988
	goto L352
L364:
	;
	goto L354
L365:
	;
	goto L197
L366:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L1
	} else {
		goto L367
	}
L367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v12
	F_errmsg(m, int32(_a_F_pg_stat_have_stats_12), v19)
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L1
	} else {
		goto L368
	}
L368:
	;
	F_errfinish(m, int32(_a_F_pg_stat_have_stats_13), int32(1426), int32(_a_F_pg_stat_have_stats_14))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L1
	} else {
		goto L369
	}
L369:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L370:
	;
	v1044 = v1023*int32(72) + int32(_a_F_pg_stat_have_stats_15)
	goto L372
L371:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_have_stats[0]))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1037+v1023<<(uint(int32(2))%32)-int32(96))))
	v1044 = v1043
	goto L372
L372:
	;
	v1045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1044))))
	if v1045&int32(1) != 0 {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v1055 = int32(1)
	goto L375
L374:
	;
	v1049 = int32(0)
	v1051 = F_pgstat_get_entry_ref(m, v1023, v14, v16, v1049, v1049)
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L1
	} else {
		goto L376
	}
L375:
	;
	return v1055
L376:
	;
	v1055 = base.B2i32(v1051 != int32(0))
	goto L375
}
func F_pg_stat_statements(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	F_pg_stat_statements_internal(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pg_stat_statements_1_3(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pg_stat_statements_internal(m, l0, int32(3), base.B2i32(v3 != int32(0)))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pg_stat_statements_internal(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
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
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v395 int32
	_ = v395
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v535 int64
	_ = v535
	var v539 int32
	_ = v539
	var v544 int64
	_ = v544
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v660 int64
	_ = v660
	var v662 int64
	_ = v662
	var v665 int64
	_ = v665
	var v666 int64
	_ = v666
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v713 float64
	_ = v713
	var v718 float64
	_ = v718
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v790 int64
	_ = v790
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v818 int32
	_ = v818
	var v823 int32
	_ = v823
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v856 int32
	_ = v856
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v877 int32
	_ = v877
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	v4 = int32(0)
	v66 = m.G0
	v68 = v66 - int32(960)
	m.G0 = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v68)+956)) = v4
	v76 = F_has_privs_of_role(m, v72, int32(3375))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[1]))
	if v79 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[1]))
	if l2 != 0 {
		goto L84
	} else {
		goto L85
	}
L4:
	;
	v301 = v298
	v302 = v299
	v303 = v4
	v304 = v4
	v305 = v4
	v306 = v4
	v307 = v4
	v308 = int32(0)
	goto L3
L5:
	;
	v301 = v290
	v302 = v291
	v303 = v292
	v304 = v293
	v305 = v294
	v306 = v296
	v307 = v295
	v308 = int32(1)
	goto L3
L6:
	;
	v290 = v285
	v291 = v286
	v292 = v287
	v293 = v4
	v294 = v288
	v295 = v4
	v296 = int32(1)
	goto L5
L7:
	;
	v282 = int32(1)
	v285 = v282
	v286 = v282
	v287 = v281
	v288 = int32(0)
	goto L6
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L77
	}
L9:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[2]))
	if v83 == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v70)+28))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	switch v90 - int32(14) {
	case 0:
		goto L21
	default:
		goto L12
	case 4:
		goto L20
	case 5:
		goto L19
	case 9:
		goto L18
	case 18:
		goto L17
	case 19:
		goto L16
	case 29:
		goto L15
	case 35:
		goto L14
	case 38:
		goto L13
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L74
	}
L13:
	;
	if l1 == int32(8) {
		goto L68
	} else {
		goto L69
	}
L14:
	;
	if l1 == int32(7) {
		goto L62
	} else {
		goto L63
	}
L15:
	;
	if l1 == int32(6) {
		goto L56
	} else {
		goto L57
	}
L16:
	;
	if l1 == int32(5) {
		goto L50
	} else {
		goto L51
	}
L17:
	;
	if l1 == int32(4) {
		goto L44
	} else {
		goto L45
	}
L18:
	;
	if l1 == int32(3) {
		goto L38
	} else {
		goto L39
	}
L19:
	;
	if l1 == int32(2) {
		goto L32
	} else {
		goto L33
	}
L20:
	;
	if l1 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	if l1 == int32(0) {
		v298 = v4
		v299 = v4
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1699), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	v298 = int32(1)
	v299 = v4
	goto L4
L27:
	;
	goto L28
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1704), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
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
	v126 = int32(1)
	v298 = v126
	v299 = v126
	goto L4
L33:
	;
	goto L34
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1709), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	v143 = int32(1)
	v290 = v143
	v291 = v143
	v292 = v4
	v293 = v4
	v294 = v4
	v295 = v4
	v296 = int32(0)
	goto L5
L39:
	;
	goto L40
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1713), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	v281 = v4
	goto L7
L45:
	;
	goto L46
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1717), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v281 = int32(1)
	goto L7
L51:
	;
	goto L52
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1721), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v192 = int32(1)
	v285 = v192
	v286 = v192
	v287 = v192
	v288 = v192
	goto L6
L57:
	;
	goto L58
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1725), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v211 = int32(1)
	v290 = v211
	v291 = v211
	v292 = v211
	v293 = v211
	v294 = v211
	v295 = v4
	v296 = v211
	goto L5
L63:
	;
	goto L64
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1729), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	v232 = int32(1)
	v290 = v232
	v291 = v232
	v292 = v232
	v293 = v232
	v294 = v232
	v295 = v232
	v296 = v232
	goto L5
L69:
	;
	goto L70
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1733), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_statements_internal_0), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1736), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errmsg(m, int32(_a_F_pg_stat_statements_internal_3), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_pg_stat_statements_internal_1), int32(1686), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	v374 = v68 + int32(936)
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[2]))
	F_hash_seq_init(m, v374, v376)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L102
	}
L82:
	;
	F_emscripten_builtin_free(m, v361)
	mBase = m.M
	v366 = F_qtext_load_file(m, v68+int32(956))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L101
	}
L83:
	;
	v344 = F_qtext_load_file(m, v68+int32(956))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L94
	}
L84:
	;
	v313 = base.AtomicRmwXchg32(m, v310, int32(20), int32(1))
	if v313 != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L86
L86:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v340 = F_LWLockAcquire(m, v338, int32(1))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L93
	}
L87:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[1]))
	F_s_lock(m, v315+int32(20), int32(_a_F_pg_stat_statements_internal_1), int32(1754), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[1]))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+32))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v324)+28))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v324)+24))
	v328 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v324)+20)), uint32(v328))
	if v326 == v328 {
		goto L83
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v336 = F_LWLockAcquire(m, v334, int32(1))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v361 = v328
	goto L82
L93:
	;
	v370 = v4
	goto L81
L94:
	;
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[1]))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	v350 = F_LWLockAcquire(m, v348, int32(1))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	if v344 == int32(0) {
		v361 = v328
		goto L82
	} else {
		goto L96
	}
L96:
	;
	v355 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[1]))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+24))
	if v327 != v356 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v361 = v344
	goto L82
L98:
	;
	goto L99
L99:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v355)+32))
	if v358 == v325 {
		v370 = v344
		goto L81
	} else {
		goto L100
	}
L100:
	;
	v361 = v344
	goto L82
L101:
	;
	v370 = v366
	goto L81
L102:
	;
	v379 = F_hash_seq_search(m, v374)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	if v379 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v382 = v68 + int32(720)
	v395 = v68 + int32(312)
	v443 = v68 + int32(464)
	v445 = v68 + int32(456)
	v447 = v68 + int32(448)
	v449 = v68 + int32(432)
	v451 = v68 + int32(424)
	v453 = v68 + int32(416)
	v469 = v68 + int32(296)
	v470 = v379
	goto L107
L105:
	;
	goto L106
L106:
	;
	v954 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_internal[1]))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v954)))
	F_LWLockRelease(m, v955)
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L1
	} else {
		goto L189
	}
L107:
	;
	v535 = *(*int64)(unsafe.Add(mBase, uint32(v470)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v68)+272)) = v535
	v539 = int32(0)
	base.MemoryFill(m, v68+int32(720), v539, int32(208))
	*(*int32)(unsafe.Add(mBase, uint32(v68)+704)) = v539
	v544 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v68)+696)) = v544
	*(*int64)(unsafe.Add(mBase, uint32(v68)+688)) = v544
	*(*int64)(unsafe.Add(mBase, uint32(v68)+680)) = v544
	*(*int64)(unsafe.Add(mBase, uint32(v68)+672)) = v544
	*(*int64)(unsafe.Add(mBase, uint32(v68)+664)) = v544
	*(*int64)(unsafe.Add(mBase, uint32(v68)+656)) = v544
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	*(*int32)(unsafe.Add(mBase, uint32(v68)+720)) = v556
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v68)+724)) = v558
	if v303 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	goto L106
L109:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v68)+728)) = v561
	v564 = v382 | int32(12)
	v565 = int32(3)
	goto L111
L110:
	;
	v564 = v382 | int32(8)
	v565 = int32(2)
	goto L111
L111:
	;
	if v76|base.B2i32(v556 == v72) == int32(1) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v643 = base.AtomicRmwXchg32(m, v470, int32(424), int32(1))
	if v643 != 0 {
		goto L138
	} else {
		goto L139
	}
L113:
	;
	if v302 != 0 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	goto L115
L115:
	;
	if v302 != 0 {
		goto L131
	} else {
		goto L132
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v564))) = v68 + int32(272)
	v575 = v565 + int32(1)
	goto L118
L117:
	;
	v575 = v565
	goto L118
L118:
	;
	if l2 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	if v370 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L121
L121:
	;
	v613 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v68+int32(656)|v575))) = uint8(v613)
	v638 = v575
	goto L112
L122:
	;
	v608 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v68+int32(656)|v575))) = uint8(v608)
	v638 = v575
	goto L112
L123:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v470)+396))
	if v578 < int32(0) {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v470)+392))
	v582 = v578 + v581
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v68)+956))
	if base.Ui32(v583) <= base.Ui32(v582) {
		goto L122
	} else {
		goto L125
	}
L125:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370+v582))))
	if v586 != 0 {
		goto L122
	} else {
		goto L126
	}
L126:
	;
	v592 = v581 + v370
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v470)+400))
	v594 = F_pg_any_to_server(m, v592, v578, v593)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v596 = F_cstring_to_text(m, v594)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+int32(720)+v575<<(uint(int32(2))%32)))) = v596
	if v594 == v592 {
		v638 = v575
		goto L112
	} else {
		goto L129
	}
L129:
	;
	F_pfree(m, v594)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v638 = v575
	goto L112
L131:
	;
	v618 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v68+int32(656)|v565))) = uint8(v618)
	v622 = v565 + v618
	goto L133
L132:
	;
	v622 = v565
	goto L133
L133:
	;
	if l2 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v629 = F_cstring_to_text(m, int32(_a_F_pg_stat_statements_internal_4))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v635 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v68+int32(656)|v622))) = uint8(v635)
	v638 = v622
	goto L112
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+int32(720)+v622<<(uint(int32(2))%32)))) = v629
	v638 = v622
	goto L112
L138:
	;
	F_s_lock(m, v470+int32(424), int32(_a_F_pg_stat_statements_internal_1), int32(1873), int32(_a_F_pg_stat_statements_internal_2))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L1
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v652 = v68 + int32(288)
	base.MemoryCopy(m, v652, v470+int32(24), int32(368))
	v657 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v470)+424)), uint32(v657))
	v660 = *(*int64)(unsafe.Add(mBase, uint32(v68)+288))
	v662 = *(*int64)(unsafe.Add(mBase, uint32(v68)+296))
	if v660 != int64(0)-v662 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L140
L142:
	;
	v665 = *(*int64)(unsafe.Add(mBase, uint32(v470)+416))
	v666 = *(*int64)(unsafe.Add(mBase, uint32(v470)+408))
	v668 = v68 + int32(720)
	v670 = v638 + int32(1)
	v673 = v668 + v670<<(uint(int32(2))%32)
	if v306 != 0 {
		goto L147
	} else {
		goto L148
	}
L143:
	;
	goto L144
L144:
	;
	v886 = F_hash_seq_search(m, v68+int32(936))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L1
	} else {
		goto L187
	}
L145:
	;
	v731 = v668 + v728<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v731))) = v68 + int32(384)
	*(*int32)(unsafe.Add(mBase, uint32(v731)+8)) = v68 + int32(400)
	*(*int32)(unsafe.Add(mBase, uint32(v731)+4)) = v68 + int32(392)
	if v301 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L146:
	;
	v707 = v68 + int32(720) + v702<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v707))) = v68 + int32(328)
	*(*int32)(unsafe.Add(mBase, uint32(v707)+8)) = v68 + int32(360)
	*(*int32)(unsafe.Add(mBase, uint32(v707)+4)) = v68 + int32(344)
	if int64(2) <= v662 {
		goto L151
	} else {
		goto L152
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v673))) = v652
	*(*int32)(unsafe.Add(mBase, uint32(v673)+4)) = v68 + int32(304)
	v676 = int32(2)
	v678 = v638<<(uint(v676)%32) + v668
	*(*int32)(unsafe.Add(mBase, uint32(v678)+20)) = v68 + int32(352)
	*(*int32)(unsafe.Add(mBase, uint32(v678)+16)) = v68 + int32(336)
	*(*int32)(unsafe.Add(mBase, uint32(v678)+12)) = v68 + int32(320)
	*(*int32)(unsafe.Add(mBase, uint32(v678)+24)) = v68 + int32(280)
	v686 = v638 + int32(7)
	v689 = v686<<(uint(v676)%32) + v668
	*(*int32)(unsafe.Add(mBase, uint32(v689)+4)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v689))) = v469
	v701 = v686
	v702 = v638 + int32(9)
	goto L146
L148:
	;
	goto L149
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v673))) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v673)+4)) = v395
	v697 = v638 + int32(3)
	if v308|v306 == int32(0) {
		v728 = v697
		goto L145
	} else {
		goto L150
	}
L150:
	;
	v701 = v670
	v702 = v697
	goto L146
L151:
	;
	v713 = *(*float64)(unsafe.Add(mBase, uint32(v68)+376))
	v718 = base.F64_sqrt(base.F64_div(v713, base.F64_convert_i64_u(v662)))
	goto L153
L152:
	;
	v718 = float64(0)
	goto L153
L153:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v68)+280)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v707)+12)) = v68 + int32(280)
	v728 = v701 + int32(6)
	goto L145
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v757+v731))) = v755
	v760 = v756 + v728
	if v304 != 0 {
		goto L158
	} else {
		goto L159
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v731)+20)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v731)+16)) = v451
	*(*int32)(unsafe.Add(mBase, uint32(v731)+12)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v731)+28)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v731)+24)) = v447
	v755 = v443
	v756 = int32(9)
	v757 = int32(32)
	goto L154
L156:
	;
	goto L157
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v731)+12)) = v68 + int32(408)
	*(*int32)(unsafe.Add(mBase, uint32(v731)+44)) = v68 + int32(472)
	*(*int32)(unsafe.Add(mBase, uint32(v731)+40)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v731)+36)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v731)+32)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v731)+28)) = v68 + int32(440)
	*(*int32)(unsafe.Add(mBase, uint32(v731)+24)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v731)+20)) = v451
	*(*int32)(unsafe.Add(mBase, uint32(v731)+16)) = v453
	v755 = v68 + int32(480)
	v756 = int32(13)
	v757 = int32(48)
	goto L154
L158:
	;
	v763 = int32(2)
	v765 = v68 + int32(720) + v760<<(uint(v763)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v765))) = v68 + int32(488)
	*(*int32)(unsafe.Add(mBase, uint32(v765)+4)) = v68 + int32(496)
	v770 = v760 + v763
	goto L160
L159:
	;
	v770 = v760
	goto L160
L160:
	;
	if v305 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v774 = int32(2)
	v776 = v68 + int32(720) + v770<<(uint(v774)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v776))) = v68 + int32(504)
	*(*int32)(unsafe.Add(mBase, uint32(v776)+4)) = v68 + int32(512)
	v781 = v770 + v774
	goto L163
L162:
	;
	v781 = v770
	goto L163
L163:
	;
	if v306 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v787 = v68 + int32(720) + v781<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v787))) = v68 + int32(528)
	*(*int32)(unsafe.Add(mBase, uint32(v787)+4)) = v68 + int32(536)
	v790 = *(*int64)(unsafe.Add(mBase, uint32(v68)+544))
	*(*int64)(unsafe.Add(mBase, uint32(v68))) = v790
	v793 = v68 + int32(16)
	v796 = F_pg_snprintf(m, v793, int32(256), int32(_a_F_pg_stat_statements_internal_5), v68)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L167
	}
L165:
	;
	v807 = v781
	goto L166
L166:
	;
	if v307 != 0 {
		goto L169
	} else {
		goto L170
	}
L167:
	;
	v799 = int32(0)
	v802 = F_DirectFunctionCall3Coll(m, int32(408), v799, v793, v799, int32(-1))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v787)+8)) = v802
	v807 = v781 + int32(3)
	goto L166
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+int32(720)+v807<<(uint(int32(2))%32)))) = v68 + int32(552)
	v818 = v807 + int32(1)
	goto L171
L170:
	;
	v818 = v807
	goto L171
L171:
	;
	if v305 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v823 = v68 + int32(720) + v818<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v823))) = v68 + int32(560)
	*(*int32)(unsafe.Add(mBase, uint32(v823)+28)) = v68 + int32(632)
	*(*int32)(unsafe.Add(mBase, uint32(v823)+24)) = v68 + int32(624)
	*(*int32)(unsafe.Add(mBase, uint32(v823)+20)) = v68 + int32(616)
	*(*int32)(unsafe.Add(mBase, uint32(v823)+16)) = v68 + int32(608)
	*(*int32)(unsafe.Add(mBase, uint32(v823)+12)) = v68 + int32(600)
	*(*int32)(unsafe.Add(mBase, uint32(v823)+8)) = v68 + int32(576)
	*(*int32)(unsafe.Add(mBase, uint32(v823)+4)) = v68 + int32(568)
	v834 = v818 + int32(8)
	goto L174
L173:
	;
	v834 = v818
	goto L174
L174:
	;
	if v304 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v838 = int32(2)
	v840 = v68 + int32(720) + v834<<(uint(v838)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v840))) = v68 + int32(592)
	*(*int32)(unsafe.Add(mBase, uint32(v840)+4)) = v68 + int32(584)
	v845 = v834 + v838
	goto L177
L176:
	;
	v845 = v834
	goto L177
L177:
	;
	if v307 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v849 = int32(2)
	v851 = v68 + int32(720) + v845<<(uint(v849)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v851))) = v68 + int32(640)
	*(*int32)(unsafe.Add(mBase, uint32(v851)+4)) = v68 + int32(648)
	v856 = v845 + v849
	goto L180
L179:
	;
	v856 = v845
	goto L180
L180:
	;
	if v304 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v862 = v68 + int32(720) + v856<<(uint(int32(2))%32)
	v863 = F_Int64GetDatum(m, v666)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L1
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v70)+24))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v70)+28))
	F_tuplestore_putvalues(m, v870, v871, v68+int32(720), v68+int32(656))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L1
	} else {
		goto L186
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v862))) = v863
	v866 = F_Int64GetDatum(m, v665)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v862)+4)) = v866
	goto L183
L186:
	;
	goto L144
L187:
	;
	if v886 != 0 {
		v470 = v886
		goto L107
	} else {
		goto L188
	}
L188:
	;
	goto L108
L189:
	;
	F_emscripten_builtin_free(m, v370)
	mBase = m.M
	m.G0 = v68 + int32(960)
	return
}
func F_pg_stat_statements_reset_1_7(m *base.Module, l0 int32) int32 {
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
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v7 = F_entry_reset(m, v2, v3, v5, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
