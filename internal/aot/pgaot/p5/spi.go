package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SPI_connect_ext(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int64
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v82 int64
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	v6 = *(*int32)(unsafe.Add(mBase, _consts[359]))
	if v6 == int32(0) {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[360]))
		if v10 != int32(-1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v134 = m.ExcPending
			if v134 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(427703), int32(0))
				mBase = m.M
				v138 = m.ExcPending
				if v138 != 0 {
					return
				} else {
					F_errfinish(m, int32(478474), int32(108), int32(61239))
					mBase = m.M
					v143 = m.ExcPending
					if v143 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _consts[361]))
			if v14 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(427703), int32(0))
					mBase = m.M
					v138 = m.ExcPending
					if v138 != 0 {
						return
					} else {
						F_errfinish(m, int32(478474), int32(108), int32(61239))
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, _consts[146]))
				v19 = F_MemoryContextAlloc(m, v17, int32(1024))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v38 = int32(16)
					v39 = v19
					*(*int32)(unsafe.Add(mBase, _consts[361])) = v38
					*(*int32)(unsafe.Add(mBase, _consts[359])) = v39
					v44 = v39
					v47 = int32(4068516)
					v49 = *(*int32)(unsafe.Add(mBase, _consts[360]))
					v51 = v49 + int32(1)
					*(*int32)(unsafe.Add(mBase, _consts[360])) = v51
					v56 = v44 + v51<<(uint(int32(6))%32)
					*(*int32)(unsafe.Add(mBase, _consts[362])) = v56
					*(*int32)(unsafe.Add(mBase, uint32(v56)+24)) = int32(0)
					v60 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v56)+16)) = v60
					*(*int64)(unsafe.Add(mBase, uint32(v56)+8)) = v60
					*(*int64)(unsafe.Add(mBase, uint32(v56))) = v60
					v67 = *(*int32)(unsafe.Add(mBase, _consts[39]))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
					v70 = *(*int32)(unsafe.Add(mBase, _consts[362]))
					v71 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v70)+41)) = uint8(v71)
					v74 = l0 & int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v70)+40)) = uint8(base.B2i32(v74 == v71))
					*(*int32)(unsafe.Add(mBase, uint32(v70)+36)) = v71
					*(*int32)(unsafe.Add(mBase, uint32(v70)+32)) = v68
					v82 = *(*int64)(unsafe.Add(mBase, _consts[363]))
					*(*int64)(unsafe.Add(mBase, uint32(v70)+48)) = v82
					v85 = *(*int32)(unsafe.Add(mBase, _consts[364]))
					*(*int32)(unsafe.Add(mBase, uint32(v70)+56)) = v85
					v88 = *(*int32)(unsafe.Add(mBase, _consts[365]))
					*(*int32)(unsafe.Add(mBase, uint32(v70)+60)) = v88
					if v74 != 0 {
						v92 = int32(4449552)
					} else {
						v92 = int32(4449544)
					}
					v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
					v98 = F_AllocSetContextCreateInternal(m, v93, int32(470154), int32(0), int32(8192), int32(8388608))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return
					} else {
						v101 = *(*int32)(unsafe.Add(mBase, _consts[362]))
						*(*int32)(unsafe.Add(mBase, uint32(v101)+20)) = v98
						v104 = *(*int32)(unsafe.Add(mBase, _consts[318]))
						v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+40)))
						if v105 != 0 {
							v106 = v104
						} else {
							v106 = v98
						}
						v111 = F_AllocSetContextCreateInternal(m, v106, int32(472633), int32(0), int32(8192), int32(8388608))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return
						} else {
							v114 = *(*int32)(unsafe.Add(mBase, _consts[362]))
							*(*int32)(unsafe.Add(mBase, uint32(v114)+24)) = v111
							v116 = int32(4449520)
							v117 = *(*int32)(unsafe.Add(mBase, _consts[9]))
							v119 = *(*int32)(unsafe.Add(mBase, uint32(v114)+20))
							*(*int32)(unsafe.Add(mBase, _consts[9])) = v119
							*(*int32)(unsafe.Add(mBase, uint32(v114)+28)) = v117
							*(*int64)(unsafe.Add(mBase, _consts[363])) = int64(0)
							v126 = int32(0)
							*(*int32)(unsafe.Add(mBase, _consts[364])) = v126
							*(*int32)(unsafe.Add(mBase, _consts[365])) = v126
							return
						}
					}
				}
			}
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _consts[361]))
		if v22 <= int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v148 = m.ExcPending
			if v148 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(427703), int32(0))
				mBase = m.M
				v152 = m.ExcPending
				if v152 != 0 {
					return
				} else {
					F_errfinish(m, int32(478474), int32(118), int32(61239))
					mBase = m.M
					v157 = m.ExcPending
					if v157 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, _consts[360]))
			if v22 <= v26 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v148 = m.ExcPending
				if v148 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(427703), int32(0))
					mBase = m.M
					v152 = m.ExcPending
					if v152 != 0 {
						return
					} else {
						F_errfinish(m, int32(478474), int32(118), int32(61239))
						mBase = m.M
						v157 = m.ExcPending
						if v157 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				if v22 != v26+int32(1) {
					v44 = v6
					v47 = int32(4068516)
					v49 = *(*int32)(unsafe.Add(mBase, _consts[360]))
					v51 = v49 + int32(1)
					*(*int32)(unsafe.Add(mBase, _consts[360])) = v51
					v56 = v44 + v51<<(uint(int32(6))%32)
					*(*int32)(unsafe.Add(mBase, _consts[362])) = v56
					*(*int32)(unsafe.Add(mBase, uint32(v56)+24)) = int32(0)
					v60 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v56)+16)) = v60
					*(*int64)(unsafe.Add(mBase, uint32(v56)+8)) = v60
					*(*int64)(unsafe.Add(mBase, uint32(v56))) = v60
					v67 = *(*int32)(unsafe.Add(mBase, _consts[39]))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
					v70 = *(*int32)(unsafe.Add(mBase, _consts[362]))
					v71 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v70)+41)) = uint8(v71)
					v74 = l0 & int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v70)+40)) = uint8(base.B2i32(v74 == v71))
					*(*int32)(unsafe.Add(mBase, uint32(v70)+36)) = v71
					*(*int32)(unsafe.Add(mBase, uint32(v70)+32)) = v68
					v82 = *(*int64)(unsafe.Add(mBase, _consts[363]))
					*(*int64)(unsafe.Add(mBase, uint32(v70)+48)) = v82
					v85 = *(*int32)(unsafe.Add(mBase, _consts[364]))
					*(*int32)(unsafe.Add(mBase, uint32(v70)+56)) = v85
					v88 = *(*int32)(unsafe.Add(mBase, _consts[365]))
					*(*int32)(unsafe.Add(mBase, uint32(v70)+60)) = v88
					if v74 != 0 {
						v92 = int32(4449552)
					} else {
						v92 = int32(4449544)
					}
					v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
					v98 = F_AllocSetContextCreateInternal(m, v93, int32(470154), int32(0), int32(8192), int32(8388608))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return
					} else {
						v101 = *(*int32)(unsafe.Add(mBase, _consts[362]))
						*(*int32)(unsafe.Add(mBase, uint32(v101)+20)) = v98
						v104 = *(*int32)(unsafe.Add(mBase, _consts[318]))
						v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+40)))
						if v105 != 0 {
							v106 = v104
						} else {
							v106 = v98
						}
						v111 = F_AllocSetContextCreateInternal(m, v106, int32(472633), int32(0), int32(8192), int32(8388608))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return
						} else {
							v114 = *(*int32)(unsafe.Add(mBase, _consts[362]))
							*(*int32)(unsafe.Add(mBase, uint32(v114)+24)) = v111
							v116 = int32(4449520)
							v117 = *(*int32)(unsafe.Add(mBase, _consts[9]))
							v119 = *(*int32)(unsafe.Add(mBase, uint32(v114)+20))
							*(*int32)(unsafe.Add(mBase, _consts[9])) = v119
							*(*int32)(unsafe.Add(mBase, uint32(v114)+28)) = v117
							*(*int64)(unsafe.Add(mBase, _consts[363])) = int64(0)
							v126 = int32(0)
							*(*int32)(unsafe.Add(mBase, _consts[364])) = v126
							*(*int32)(unsafe.Add(mBase, _consts[365])) = v126
							return
						}
					}
				} else {
					v35 = F_repalloc(m, v6, v22<<(uint(int32(7))%32))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						v38 = v22 << (uint(int32(1)) % 32)
						v39 = v35
						*(*int32)(unsafe.Add(mBase, _consts[361])) = v38
						*(*int32)(unsafe.Add(mBase, _consts[359])) = v39
						v44 = v39
						v47 = int32(4068516)
						v49 = *(*int32)(unsafe.Add(mBase, _consts[360]))
						v51 = v49 + int32(1)
						*(*int32)(unsafe.Add(mBase, _consts[360])) = v51
						v56 = v44 + v51<<(uint(int32(6))%32)
						*(*int32)(unsafe.Add(mBase, _consts[362])) = v56
						*(*int32)(unsafe.Add(mBase, uint32(v56)+24)) = int32(0)
						v60 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v56)+16)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v56)+8)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v56))) = v60
						v67 = *(*int32)(unsafe.Add(mBase, _consts[39]))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
						v70 = *(*int32)(unsafe.Add(mBase, _consts[362]))
						v71 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v70)+41)) = uint8(v71)
						v74 = l0 & int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v70)+40)) = uint8(base.B2i32(v74 == v71))
						*(*int32)(unsafe.Add(mBase, uint32(v70)+36)) = v71
						*(*int32)(unsafe.Add(mBase, uint32(v70)+32)) = v68
						v82 = *(*int64)(unsafe.Add(mBase, _consts[363]))
						*(*int64)(unsafe.Add(mBase, uint32(v70)+48)) = v82
						v85 = *(*int32)(unsafe.Add(mBase, _consts[364]))
						*(*int32)(unsafe.Add(mBase, uint32(v70)+56)) = v85
						v88 = *(*int32)(unsafe.Add(mBase, _consts[365]))
						*(*int32)(unsafe.Add(mBase, uint32(v70)+60)) = v88
						if v74 != 0 {
							v92 = int32(4449552)
						} else {
							v92 = int32(4449544)
						}
						v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
						v98 = F_AllocSetContextCreateInternal(m, v93, int32(470154), int32(0), int32(8192), int32(8388608))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							v101 = *(*int32)(unsafe.Add(mBase, _consts[362]))
							*(*int32)(unsafe.Add(mBase, uint32(v101)+20)) = v98
							v104 = *(*int32)(unsafe.Add(mBase, _consts[318]))
							v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+40)))
							if v105 != 0 {
								v106 = v104
							} else {
								v106 = v98
							}
							v111 = F_AllocSetContextCreateInternal(m, v106, int32(472633), int32(0), int32(8192), int32(8388608))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return
							} else {
								v114 = *(*int32)(unsafe.Add(mBase, _consts[362]))
								*(*int32)(unsafe.Add(mBase, uint32(v114)+24)) = v111
								v116 = int32(4449520)
								v117 = *(*int32)(unsafe.Add(mBase, _consts[9]))
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v114)+20))
								*(*int32)(unsafe.Add(mBase, _consts[9])) = v119
								*(*int32)(unsafe.Add(mBase, uint32(v114)+28)) = v117
								*(*int64)(unsafe.Add(mBase, _consts[363])) = int64(0)
								v126 = int32(0)
								*(*int32)(unsafe.Add(mBase, _consts[364])) = v126
								*(*int32)(unsafe.Add(mBase, _consts[365])) = v126
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_SPI_keepplan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v5 != int32(569278163) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v8 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v9 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v10 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v10)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v14 = *(*int32)(unsafe.Add(mBase, _consts[366]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v18 != v14 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v48 == int32(0) {
		goto L1
	} else {
		goto L23
	}
L7:
	;
	if v18 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	goto L6
L10:
	;
	if v14 != 0 {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if v23 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v22 == int32(0) {
		goto L10
	} else {
		goto L16
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v22
	goto L12
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v22
	goto L12
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v28
	goto L10
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v14
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v35
	if v35 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
	goto L9
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v12
	goto L22
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v12
	goto L6
L23:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v51 <= int32(0) {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v55 = int32(0)
	goto L25
L25:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v55<<(uint(int32(2))%32))))
	F_SaveCachedPlan(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L1
L27:
	;
	return
L28:
	;
	v65 = v55 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v65 < v66 {
		v55 = v65
		goto L25
	} else {
		goto L29
	}
L29:
	;
	goto L26
}
func F_spi_dest_startup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v6 = *(*int32)(unsafe.Add(mBase, _consts[362]))
	if v6 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		if v7 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(221754), int32(0))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return
				} else {
					F_errfinish(m, int32(478474), int32(2133), int32(221771))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v8 = int32(4449520)
			v9 = *(*int32)(unsafe.Add(mBase, _consts[9]))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
			*(*int32)(unsafe.Add(mBase, _consts[9])) = v11
			v18 = F_AllocSetContextCreateInternal(m, v11, int32(381418), int32(0), int32(8192), int32(8388608))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[9])) = v18
				v22 = F_palloc0(m, int32(40))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _consts[362]))
					*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v22
					*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v18
					v29 = *(*int32)(unsafe.Add(mBase, _consts[39]))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v30
					v33 = *(*int32)(unsafe.Add(mBase, _consts[362]))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v22 + int32(28)
					*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = int64(128)
					v42 = F_palloc(m, int32(512))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v42
						v47 = F_CreateTupleDescCopy(m, l2)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v22))) = v47
							*(*int32)(unsafe.Add(mBase, _consts[9])) = v9
							return
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(514340), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				F_errfinish(m, int32(478474), int32(2130), int32(221771))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
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
