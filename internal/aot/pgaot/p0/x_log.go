package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_CheckXLogRemoved(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_CheckXLogRemoved[0]))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_CheckXLogRemoved[1]))
	v17 = base.AtomicRmwXchg32(m, v14, int32(440), int32(1))
	if v17 != 0 {
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_CheckXLogRemoved[1]))
		F_s_lock(m, v19+int32(440), int32(_a_F_CheckXLogRemoved_0), int32(3730), int32(_a_F_CheckXLogRemoved_1))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_CheckXLogRemoved[1]))
			v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)+232))
			v30 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v28)+440)), uint32(v30))
			if base.Ui64(l0) <= base.Ui64(v29) {
				v35 = v9 + int32(16)
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_CheckXLogRemoved[2]))
				F_XLogFileName(m, v35, l1, l0, v37)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_CheckXLogRemoved[0])) = v12
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v35
							F_errmsg(m, int32(_a_F_CheckXLogRemoved_2), v9)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_CheckXLogRemoved_0), int32(3743), int32(_a_F_CheckXLogRemoved_1))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_CheckXLogRemoved[0])) = v12
				m.G0 = v9 + int32(80)
				return
			}
		}
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, _c_F_CheckXLogRemoved[1]))
		v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)+232))
		v30 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v28)+440)), uint32(v30))
		if base.Ui64(l0) <= base.Ui64(v29) {
			v35 = v9 + int32(16)
			v37 = *(*int32)(unsafe.Add(mBase, _c_F_CheckXLogRemoved[2]))
			F_XLogFileName(m, v35, l1, l0, v37)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_CheckXLogRemoved[0])) = v12
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v35
						F_errmsg(m, int32(_a_F_CheckXLogRemoved_2), v9)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_CheckXLogRemoved_0), int32(3743), int32(_a_F_CheckXLogRemoved_1))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_CheckXLogRemoved[0])) = v12
			m.G0 = v9 + int32(80)
			return
		}
	}
}
func F_InstallXLogFileSegment(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
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
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(1152)
	m.G0 = v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l4
	v20 = int64(*(*int32)(unsafe.Add(mBase, _c_F_InstallXLogFileSegment[0])))
	v21 = base.I64_div_u_s(int64(4294967296), v20)
	v22 = base.I64_div_u_s(v16, v21)
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+20)) = uint32(v22)
	v25 = v16 - v21*v22
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+24)) = uint32(v25)
	v28 = v14 + int32(128)
	v33 = F_pg_snprintf(m, v28, int32(1024), int32(_a_F_InstallXLogFileSegment_0), v14+int32(16))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_InstallXLogFileSegment[1]))
	v42 = F_LWLockAcquire(m, v38+int32(1152), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_InstallXLogFileSegment[2]))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+320)))
	if v46 != int32(1) {
		v123 = v6
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_InstallXLogFileSegment[1]))
	F_LWLockRelease(m, v128+int32(1152))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L20
	}
L5:
	;
	if l2 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v112 = F_durable_rename(m, l1, v14+int32(128), int32(15))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L19
	}
L7:
	;
	v53 = F___fstatat(m, int32(-100), v28, v14+int32(32), int32(0))
	mBase = m.M
	goto L10
L8:
	;
	goto L9
L9:
	;
	v96 = F_durable_unlink(m, v14+int32(128), int32(14))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L18
	}
L10:
	;
	if v53 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L12
L12:
	;
	v65 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui64(l3) <= base.Ui64(v65) {
		v123 = v6
		goto L4
	} else {
		goto L14
	}
L13:
	;
	goto L6
L14:
	;
	v68 = v65 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l4
	v73 = int64(*(*int32)(unsafe.Add(mBase, _c_F_InstallXLogFileSegment[0])))
	v74 = base.I64_div_u_s(int64(4294967296), v73)
	v75 = base.I64_div_u_s(v68, v74)
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+4)) = uint32(v75)
	v78 = v68 - v74*v75
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+8)) = uint32(v78)
	v81 = v14 + int32(128)
	v84 = F_pg_snprintf(m, v81, int32(1024), int32(_a_F_InstallXLogFileSegment_0), v14)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v90 = F___fstatat(m, int32(-100), v81, v14+int32(32), int32(0))
	mBase = m.M
	goto L16
L16:
	;
	if v90 == int32(0) {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	goto L6
L19:
	;
	v123 = base.B2i32(v112 == int32(0))
	goto L4
L20:
	;
	m.G0 = v14 + int32(1152)
	return v123
}
func F_XLogBeginInsert(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBeginInsert[0]))
	if int32(0) <= v3 {
		v28 = base.B2i32(v3 != int32(0))
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogBeginInsert[1])))
		if v9 == int32(1) {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_XLogBeginInsert[2]))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+316))
			v17 = base.B2i32(v15 != int32(2))
			*(*uint8)(unsafe.Add(mBase, _c_F_XLogBeginInsert[1])) = uint8(v17)
			if v15 != int32(2) {
				v26 = int32(0)
			} else {
				v22 = int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_XLogBeginInsert[0])) = v22
				v26 = v22
			}
		} else {
			v22 = int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_XLogBeginInsert[0])) = v22
			v26 = v22
		}
		v28 = v26
	}
	if v28 != 0 {
		v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XLogBeginInsert[3])))
		if v30 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_XLogBeginInsert_0), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_XLogBeginInsert_1), int32(160), int32(_a_F_XLogBeginInsert_2))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v32 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_XLogBeginInsert[3])) = uint8(v32)
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_XLogBeginInsert_3), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_XLogBeginInsert_1), int32(157), int32(_a_F_XLogBeginInsert_2))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_XLogReaderAllocate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_palloc_extended(m, int32(1264), int32(6))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v65 = int32(0)
			m.G0 = v8 + int32(16)
			return v65
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v19
			v21 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			*(*int64)(unsafe.Add(mBase, uint32(v12))) = v21
			v25 = F_palloc_extended(m, int32(_a_F_XLogReaderAllocate_0), int32(2))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v25
				if v25 == int32(0) {
					F_pfree(m, v12)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v65 = int32(0)
						m.G0 = v8 + int32(16)
						return v65
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+1184)) = int32(0)
					*(*int64)(unsafe.Add(mBase, uint32(v12)+1176)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+1168)) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+1160)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = l2
					v40 = F_palloc_extended(m, int32(1001), int32(2))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+1252)) = v40
						if v40 == int32(0) {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)+128))
							F_pfree(m, v45)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v12)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v65 = int32(0)
									m.G0 = v8 + int32(16)
									return v65
								}
							}
						} else {
							v48 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v48)
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1244))
							if v50 != 0 {
								F_pfree(m, v50)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v54 = F_palloc(m, int32(_a_F_XLogReaderAllocate_1))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+1248)) = int32(_a_F_XLogReaderAllocate_1)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+1244)) = v54
										v65 = v12
										m.G0 = v8 + int32(16)
										return v65
									}
								}
							} else {
								v54 = F_palloc(m, int32(_a_F_XLogReaderAllocate_1))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+1248)) = int32(_a_F_XLogReaderAllocate_1)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+1244)) = v54
									v65 = v12
									m.G0 = v8 + int32(16)
									return v65
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_XLogRecGetBlockTagExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	v7 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
	if v9 < l1 {
		v33 = v7
	} else {
		v13 = v8 + l1*int32(52)
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)))
		if v14 != int32(1) {
			v33 = v7
		} else {
			v18 = v13 + int32(76)
			if l2 != 0 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v19
				v21 = *(*int64)(unsafe.Add(mBase, uint32(v18)+4))
				*(*int64)(unsafe.Add(mBase, uint32(l2))) = v21
			} else {
			}
			if l3 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v23
			} else {
			}
			if l4 != 0 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = v25
			} else {
			}
			v27 = int32(1)
			if l5 == int32(0) {
				v33 = v27
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
				*(*int32)(unsafe.Add(mBase, uint32(l5))) = v30
				v33 = v27
			}
		}
	}
	return v33
}
