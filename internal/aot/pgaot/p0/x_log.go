package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	v14 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+440)) = int32(1)
	if v15 != 0 {
		v19 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		F_s_lock(m, v19+int32(440), int32(523959), int32(3730), int32(462465))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			*(*int32)(unsafe.Add(mBase, uint32(v28)+440)) = int32(0)
			v31 = *(*int64)(unsafe.Add(mBase, uint32(v28)+232))
			if base.Ui64(l0) <= base.Ui64(v31) {
				v36 = *(*int32)(unsafe.Add(mBase, _consts[164]))
				F_XLogFileName(m, v9+int32(16), l1, l0, v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[163])) = v12
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(16)
							F_errmsg(m, int32(462364), v9)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								F_errfinish(m, int32(523959), int32(3743), int32(462465))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
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
				*(*int32)(unsafe.Add(mBase, _consts[163])) = v12
				m.G0 = v9 + int32(80)
				return
			}
		}
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		*(*int32)(unsafe.Add(mBase, uint32(v28)+440)) = int32(0)
		v31 = *(*int64)(unsafe.Add(mBase, uint32(v28)+232))
		if base.Ui64(l0) <= base.Ui64(v31) {
			v36 = *(*int32)(unsafe.Add(mBase, _consts[164]))
			F_XLogFileName(m, v9+int32(16), l1, l0, v36)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[163])) = v12
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(16)
						F_errmsg(m, int32(462364), v9)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							F_errfinish(m, int32(523959), int32(3743), int32(462465))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
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
			*(*int32)(unsafe.Add(mBase, _consts[163])) = v12
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	v6 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(1152)
	m.G0 = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l4
	v19 = int64(*(*int32)(unsafe.Add(mBase, _consts[164])))
	v20 = base.I64_div_u_s(int64(4294967296), v19)
	v21 = base.I64_div_u_s(v15, v20)
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+20)) = uint32(v21)
	v24 = v15 - v20*v21
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+24)) = uint32(v24)
	v32 = F_pg_snprintf(m, v13+int32(128), int32(1024), int32(537023), v13+int32(16))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v41 = F_LWLockAcquire(m, v37+int32(1152), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+320)))
	if v45 != int32(1) {
		v123 = v6
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _consts[29]))
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
	v113 = F_durable_rename(m, l1, v13+int32(128), int32(15))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L19
	}
L7:
	;
	v54 = F___fstatat(m, int32(-100), v13+int32(128), v13+int32(32), int32(0))
	mBase = m.M
	goto L10
L8:
	;
	goto L9
L9:
	;
	v98 = F_durable_unlink(m, v13+int32(128), int32(14))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L18
	}
L10:
	;
	if v54 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l4
	v73 = int64(*(*int32)(unsafe.Add(mBase, _consts[164])))
	v74 = base.I64_div_u_s(int64(4294967296), v73)
	v75 = base.I64_div_u_s(v68, v74)
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+4)) = uint32(v75)
	v78 = v68 - v74*v75
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+8)) = uint32(v78)
	v84 = F_pg_snprintf(m, v13+int32(128), int32(1024), int32(537023), v13)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v92 = F___fstatat(m, int32(-100), v13+int32(128), v13+int32(32), int32(0))
	mBase = m.M
	goto L16
L16:
	;
	if v92 == int32(0) {
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
	v123 = base.B2i32(v113 == int32(0))
	goto L4
L20:
	;
	m.G0 = v13 + int32(1152)
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[179]))
	if int32(0) <= v3 {
		v28 = base.B2i32(v3 != int32(0))
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[2])))
		if v9 == int32(1) {
			v14 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+316))
			v17 = base.B2i32(v15 != int32(2))
			*(*uint8)(unsafe.Add(mBase, _consts[2])) = uint8(v17)
			if v15 != int32(2) {
				v26 = int32(0)
			} else {
				v22 = int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[179])) = v22
				v26 = v22
			}
		} else {
			v22 = int32(1)
			*(*int32)(unsafe.Add(mBase, _consts[179])) = v22
			v26 = v22
		}
		v28 = v26
	}
	if v28 != 0 {
		v30 = int32(*(*uint8)(unsafe.Add(mBase, _consts[180])))
		if v30 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(476323), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					F_errfinish(m, int32(517557), int32(160), int32(88269))
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
			*(*uint8)(unsafe.Add(mBase, _consts[180])) = uint8(v32)
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(14723), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				F_errfinish(m, int32(517557), int32(157), int32(88269))
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
	var v19 int64
	_ = v19
	var v21 int32
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
	var v64 int32
	_ = v64
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
			v64 = int32(0)
			m.G0 = v8 + int32(16)
			return v64
		} else {
			v19 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			*(*int64)(unsafe.Add(mBase, uint32(v12))) = v19
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v21
			v25 = F_palloc_extended(m, int32(8192), int32(2))
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
						v64 = int32(0)
						m.G0 = v8 + int32(16)
						return v64
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
									v64 = int32(0)
									m.G0 = v8 + int32(16)
									return v64
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
									v54 = F_palloc(m, int32(40960))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+1248)) = int32(40960)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+1244)) = v54
										v64 = v12
										m.G0 = v8 + int32(16)
										return v64
									}
								}
							} else {
								v54 = F_palloc(m, int32(40960))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+1248)) = int32(40960)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+1244)) = v54
									v64 = v12
									m.G0 = v8 + int32(16)
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
func F_XLogRecGetBlockTagExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v21 int32
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
		v15 = v8 + l1*int32(52) + int32(76)
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
		if v16 != int32(1) {
			v33 = v7
		} else {
			if l2 != 0 {
				v19 = *(*int64)(unsafe.Add(mBase, uint32(v15)+4))
				*(*int64)(unsafe.Add(mBase, uint32(l2))) = v19
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v21
			} else {
			}
			if l3 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v23
			} else {
			}
			if l4 != 0 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = v25
			} else {
			}
			v27 = int32(1)
			if l5 == int32(0) {
				v33 = v27
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
				*(*int32)(unsafe.Add(mBase, uint32(l5))) = v30
				v33 = v27
			}
		}
	}
	return v33
}
