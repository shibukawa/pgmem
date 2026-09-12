package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GinDataLeafPageGetItems(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	v10 = l0 + int32(32)
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v12 = l0 + v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+6)))
	if v13&int32(128) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v17 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	if v17 == int64(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v104
	v108 = F_palloc(m, v104*int32(6))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L17
	} else {
		goto L19
	}
L4:
	;
	v90 = v10
	v94 = v16 - int32(32)
	goto L6
L5:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+38)))
	v27 = v10 + (v22+int32(1))&int32(131070)
	v29 = v27 + int32(8)
	v30 = l0 + v16
	if base.Ui32(v30) <= base.Ui32(v29) {
		v81 = v10
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v94 != 0 {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v90 = v81
	v94 = v30 - v81
	goto L6
L8:
	;
	v32 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	v33 = int64(48)
	v36 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
	v37 = int64(32)
	v39 = v32<<(uint(v33)%64) | v17 | v36<<(uint(v37)%64)
	v40 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v27)+12)))
	v41 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v27)+10)))
	v44 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v27)+8)))
	if base.Ui64(v39) < base.Ui64(v40|(v41<<(uint(v37)%64)|v44<<(uint(v33)%64))) {
		v81 = v10
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v53 = v27
	v56 = v29
	goto L10
L10:
	;
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+14)))
	v63 = v56 + (v58+int32(1))&int32(131070)
	v65 = v63 + int32(8)
	if base.Ui32(v30) <= base.Ui32(v65) {
		v81 = v56
		goto L7
	} else {
		goto L12
	}
L11:
	;
	v81 = v56
	goto L7
L12:
	;
	v67 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v63)+12)))
	v68 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v63)+10)))
	v71 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v63)+8)))
	if base.Ui64(v67|(v68<<(uint(int64(32))%64)|v71<<(uint(int64(48))%64))) <= base.Ui64(v39) {
		v53 = v63
		v56 = v65
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v95 = F_ginPostingListDecodeAllSegments(m, v90, v94, l1)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v100 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v100
	return v100
L17:
	;
	return int32(0)
L18:
	;
	return v95
L19:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v112 = v110 * int32(6)
	if v112 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	return v114
L21:
	;
	v113 = F__emscripten_memcpy_bulkmem(m, v108, v10, v112)
	mBase = m.M
	v114 = v113
	goto L23
L22:
	;
	v114 = v108
	goto L23
L23:
	;
	goto L20
}
func F_GinInitMetabuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int64
	_ = v70
	var v84 int32
	_ = v84
	if l0 < int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, _consts[12]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6+(l0^int32(-1))<<(uint(int32(2))%32))))
		v20 = v12
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[13]))
		v20 = v14 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	if v20&int32(3) != 0 {
	} else {
	}
	v47 = F___memset(m, v20, int32(0), int32(8192))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v20)+10)) = int32(1572864)
	v53 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+18)) = uint16(v53)
	v59 = int32(8184)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)) = uint16(v59)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+14)) = uint16(v59)
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	v63 = v20 + v62
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(-1)
	v66 = int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v63)+6)) = uint16(v66)
	v70 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20-int32(-64)))) = v70
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = v70
	*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = v70
	*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = int32(2)
	v84 = int32(80)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+12)) = uint16(v84)
	return
}
func F_GinInitPage(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v2 = l1
	if l0&int32(3) != 0 {
		v29 = F___memset(m, l0, int32(0), l2)
		mBase = m.M
	} else {
		if base.Ui32(int32(1024)) < base.Ui32(l2) {
			v29 = F___memset(m, l0, int32(0), l2)
			mBase = m.M
		} else {
			if l2&int32(3) != 0 {
				v29 = F___memset(m, l0, int32(0), l2)
				mBase = m.M
			} else {
				v13 = l0 + l2
				if base.Ui32(v13) <= base.Ui32(l0) {
				} else {
					v19 = l0 + int32(4)
					if base.Ui32(v19) < base.Ui32(v13) {
						v21 = v13
					} else {
						v21 = v19
					}
					v27 = F___memset(m, l0, int32(0), (l0^int32(-1)+v21)&int32(-4)+int32(4))
					mBase = m.M
				}
			}
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(1572864)
	v35 = l2 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v35)
	v41 = l2 - int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v41)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v41)
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v45 = l0 + v44
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+6)) = uint16(v2)
	return
}
func F_ginAllocEntryAccumulator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if base.Ui32(v5) < base.Ui32(int32(2048)) {
			v22 = v5
			v23 = v4
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v22 + int32(1)
			return v23 + v22*int32(40)
		} else {
			v10 = F_palloc(m, int32(81920))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v10
				v15 = F_GetMemoryChunkSpace(m, v10)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v15 + v17
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v22 = int32(0)
					v23 = v20
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v22 + int32(1)
					return v23 + v22*int32(40)
				}
			}
		}
	} else {
		v10 = F_palloc(m, int32(81920))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v10
			v15 = F_GetMemoryChunkSpace(m, v10)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v15 + v17
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v22 = int32(0)
				v23 = v20
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v22 + int32(1)
				return v23 + v22*int32(40)
			}
		}
	}
}
func F_ginBeginBAScan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = m.G0
	v6 = int32(16)
	v7 = v5 - v6
	m.G0 = v7
	v10 = l0 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v4
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(base.B2i32(v14 == int32(4159512)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(790)
	m.G0 = v7 + v6
	return
}
func F_ginCombineData(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
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
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v58 int64
	_ = v58
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if base.Ui32(v6) <= base.Ui32(v7) {
		if v6 < int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					F_errmsg(m, int32(343775), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						F_errhint(m, int32(697148), int32(0))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							F_errfinish(m, int32(521179), int32(45), int32(529965))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
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
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v12 = F_GetMemoryChunkSpace(m, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v14 - v12
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v17 << (uint(int32(1)) % 32)
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v24 = F_repalloc_huge(m, v21, v17*int32(12))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v24
					v27 = F_GetMemoryChunkSpace(m, v24)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v27 + v29
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
						if v35 != 0 {
						} else {
							v36 = int32(6)
							v38 = v34 + v33*v36
							v41 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v38-int32(4)))))
							v42 = int64(32)
							v46 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v38-v36))))
							v47 = int64(48)
							v52 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v38-int32(2)))))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
							v55 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v54)+2)))
							v58 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v54))))
							v62 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v54)+4)))
							if base.Ui64(v41<<(uint(v42)%64)|v46<<(uint(v47)%64)|v52) <= base.Ui64(v55<<(uint(v42)%64)|v58<<(uint(v47)%64)|v62) {
							} else {
								v65 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v65)
							}
						}
						v70 = v34 + v33*int32(6)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
						*(*int32)(unsafe.Add(mBase, uint32(v70))) = v72
						v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
						*(*uint16)(unsafe.Add(mBase, uint32(v70)+4)) = uint16(v74)
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v76 + int32(1)
						return
					}
				}
			}
		}
	} else {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v35 != 0 {
		} else {
			v36 = int32(6)
			v38 = v34 + v33*v36
			v41 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v38-int32(4)))))
			v42 = int64(32)
			v46 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v38-v36))))
			v47 = int64(48)
			v52 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v38-int32(2)))))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			v55 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v54)+2)))
			v58 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v54))))
			v62 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v54)+4)))
			if base.Ui64(v41<<(uint(v42)%64)|v46<<(uint(v47)%64)|v52) <= base.Ui64(v55<<(uint(v42)%64)|v58<<(uint(v47)%64)|v62) {
			} else {
				v65 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v65)
			}
		}
		v70 = v34 + v33*int32(6)
		v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
		*(*int32)(unsafe.Add(mBase, uint32(v70))) = v72
		v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
		*(*uint16)(unsafe.Add(mBase, uint32(v70)+4)) = uint16(v74)
		v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v76 + int32(1)
		return
	}
}
func F_ginEntryFillRoot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
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
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	v3 = l2
	v5 = l4
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+12)))
	if base.Ui32(v9) < base.Ui32(int32(25)) {
		v20 = int32(-1)
	} else {
		v20 = int32(base.Ui32(v9+int32(262120))>>(uint(int32(2))%32))&int32(65535) - int32(1)
	}
	v21 = int32(2)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l3+v20<<(uint(v21)%32))+24))
	v27 = l3 + v24&int32(32767)
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+16)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v28)+6)))
	if v30&v21 == int32(0) {
		v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+6)))
		v61 = F_palloc(m, v58&int32(8191))
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return
		} else {
			v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+6)))
			v65 = v63 & int32(8191)
			if v65 != 0 {
				v66 = F__emscripten_memcpy_bulkmem(m, v61, v27, v65)
				mBase = m.M
				v67 = v66
			} else {
				v67 = v61
			}
			v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+6)))
			v69 = v68
			v70 = v61
			v72 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v70)+4)) = uint16(v72)
			*(*uint16)(unsafe.Add(mBase, uint32(v70)+2)) = uint16(v3)
			v76 = int32(base.Ui32(v3) >> (uint(int32(16)) % 32))
			*(*uint16)(unsafe.Add(mBase, uint32(v70))) = uint16(v76)
			v82 = F_PageAddItemExtended(m, l1, v70, v69&int32(8191), v72, v72)
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return
			} else {
				if v82 != 0 {
					F_pfree(m, v70)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+12)))
						if base.Ui32(v87) < base.Ui32(int32(25)) {
							v98 = int32(-1)
						} else {
							v98 = int32(base.Ui32(v87+int32(262120))>>(uint(int32(2))%32))&int32(65535) - int32(1)
						}
						v99 = int32(2)
						v102 = *(*int32)(unsafe.Add(mBase, uint32(l5+v98<<(uint(v99)%32))+24))
						v105 = l5 + v102&int32(32767)
						v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+16)))
						v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v106)+6)))
						if v108&v99 == int32(0) {
							v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+6)))
							v139 = F_palloc(m, v136&int32(8191))
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return
							} else {
								v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+6)))
								v143 = v141 & int32(8191)
								if v143 != 0 {
									v144 = F__emscripten_memcpy_bulkmem(m, v139, v105, v143)
									mBase = m.M
									v145 = v144
								} else {
									v145 = v139
								}
								v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+6)))
								v148 = v139
								v149 = v146
								v150 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(v148)+4)) = uint16(v150)
								*(*uint16)(unsafe.Add(mBase, uint32(v148)+2)) = uint16(v5)
								v154 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
								*(*uint16)(unsafe.Add(mBase, uint32(v148))) = uint16(v154)
								v160 = F_PageAddItemExtended(m, l1, v148, v149&int32(8191), v150, v150)
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return
								} else {
									if v160 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v182 = m.ExcPending
										if v182 != 0 {
											return
										} else {
											F_errmsg_internal(m, int32(426262), int32(0))
											mBase = m.M
											v186 = m.ExcPending
											if v186 != 0 {
												return
											} else {
												F_errfinish(m, int32(522889), int32(736), int32(90715))
												mBase = m.M
												v191 = m.ExcPending
												if v191 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										F_pfree(m, v148)
										mBase = m.M
										v165 = m.ExcPending
										if v165 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						} else {
							v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+4)))
							if v113 == int32(65535) {
								v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+6)))
								v139 = F_palloc(m, v136&int32(8191))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return
								} else {
									v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+6)))
									v143 = v141 & int32(8191)
									if v143 != 0 {
										v144 = F__emscripten_memcpy_bulkmem(m, v139, v105, v143)
										mBase = m.M
										v145 = v144
									} else {
										v145 = v139
									}
									v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+6)))
									v148 = v139
									v149 = v146
									v150 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v148)+4)) = uint16(v150)
									*(*uint16)(unsafe.Add(mBase, uint32(v148)+2)) = uint16(v5)
									v154 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
									*(*uint16)(unsafe.Add(mBase, uint32(v148))) = uint16(v154)
									v160 = F_PageAddItemExtended(m, l1, v148, v149&int32(8191), v150, v150)
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return
									} else {
										if v160 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v182 = m.ExcPending
											if v182 != 0 {
												return
											} else {
												F_errmsg_internal(m, int32(426262), int32(0))
												mBase = m.M
												v186 = m.ExcPending
												if v186 != 0 {
													return
												} else {
													F_errfinish(m, int32(522889), int32(736), int32(90715))
													mBase = m.M
													v191 = m.ExcPending
													if v191 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											F_pfree(m, v148)
											mBase = m.M
											v165 = m.ExcPending
											if v165 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							} else {
								v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+2)))
								v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105))))
								v126 = (v116 | v117<<(uint(int32(16))%32)&int32(2147418112) + int32(7)) & int32(-8)
								v127 = F_palloc(m, v126)
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return
								} else {
									if v126 != 0 {
										v129 = F__emscripten_memcpy_bulkmem(m, v127, v105, v126)
										mBase = m.M
										v130 = v129
									} else {
										v130 = v127
									}
									v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+6)))
									v134 = v131&int32(-8192) | v126
									*(*uint16)(unsafe.Add(mBase, uint32(v130)+6)) = uint16(v134)
									v148 = v127
									v149 = v134
									v150 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v148)+4)) = uint16(v150)
									*(*uint16)(unsafe.Add(mBase, uint32(v148)+2)) = uint16(v5)
									v154 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
									*(*uint16)(unsafe.Add(mBase, uint32(v148))) = uint16(v154)
									v160 = F_PageAddItemExtended(m, l1, v148, v149&int32(8191), v150, v150)
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return
									} else {
										if v160 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v182 = m.ExcPending
											if v182 != 0 {
												return
											} else {
												F_errmsg_internal(m, int32(426262), int32(0))
												mBase = m.M
												v186 = m.ExcPending
												if v186 != 0 {
													return
												} else {
													F_errfinish(m, int32(522889), int32(736), int32(90715))
													mBase = m.M
													v191 = m.ExcPending
													if v191 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											F_pfree(m, v148)
											mBase = m.M
											v165 = m.ExcPending
											if v165 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v169 = m.ExcPending
					if v169 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(426262), int32(0))
						mBase = m.M
						v173 = m.ExcPending
						if v173 != 0 {
							return
						} else {
							F_errfinish(m, int32(522889), int32(731), int32(90715))
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
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
		}
	} else {
		v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+4)))
		if v35 == int32(65535) {
			v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+6)))
			v61 = F_palloc(m, v58&int32(8191))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return
			} else {
				v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+6)))
				v65 = v63 & int32(8191)
				if v65 != 0 {
					v66 = F__emscripten_memcpy_bulkmem(m, v61, v27, v65)
					mBase = m.M
					v67 = v66
				} else {
					v67 = v61
				}
				v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+6)))
				v69 = v68
				v70 = v61
				v72 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v70)+4)) = uint16(v72)
				*(*uint16)(unsafe.Add(mBase, uint32(v70)+2)) = uint16(v3)
				v76 = int32(base.Ui32(v3) >> (uint(int32(16)) % 32))
				*(*uint16)(unsafe.Add(mBase, uint32(v70))) = uint16(v76)
				v82 = F_PageAddItemExtended(m, l1, v70, v69&int32(8191), v72, v72)
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return
				} else {
					if v82 != 0 {
						F_pfree(m, v70)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+12)))
							if base.Ui32(v87) < base.Ui32(int32(25)) {
								v98 = int32(-1)
							} else {
								v98 = int32(base.Ui32(v87+int32(262120))>>(uint(int32(2))%32))&int32(65535) - int32(1)
							}
							v99 = int32(2)
							v102 = *(*int32)(unsafe.Add(mBase, uint32(l5+v98<<(uint(v99)%32))+24))
							v105 = l5 + v102&int32(32767)
							v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+16)))
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v106)+6)))
							if v108&v99 == int32(0) {
								v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+6)))
								v139 = F_palloc(m, v136&int32(8191))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return
								} else {
									v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+6)))
									v143 = v141 & int32(8191)
									if v143 != 0 {
										v144 = F__emscripten_memcpy_bulkmem(m, v139, v105, v143)
										mBase = m.M
										v145 = v144
									} else {
										v145 = v139
									}
									v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+6)))
									v148 = v139
									v149 = v146
									v150 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v148)+4)) = uint16(v150)
									*(*uint16)(unsafe.Add(mBase, uint32(v148)+2)) = uint16(v5)
									v154 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
									*(*uint16)(unsafe.Add(mBase, uint32(v148))) = uint16(v154)
									v160 = F_PageAddItemExtended(m, l1, v148, v149&int32(8191), v150, v150)
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return
									} else {
										if v160 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v182 = m.ExcPending
											if v182 != 0 {
												return
											} else {
												F_errmsg_internal(m, int32(426262), int32(0))
												mBase = m.M
												v186 = m.ExcPending
												if v186 != 0 {
													return
												} else {
													F_errfinish(m, int32(522889), int32(736), int32(90715))
													mBase = m.M
													v191 = m.ExcPending
													if v191 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											F_pfree(m, v148)
											mBase = m.M
											v165 = m.ExcPending
											if v165 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							} else {
								v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+4)))
								if v113 == int32(65535) {
									v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+6)))
									v139 = F_palloc(m, v136&int32(8191))
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return
									} else {
										v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+6)))
										v143 = v141 & int32(8191)
										if v143 != 0 {
											v144 = F__emscripten_memcpy_bulkmem(m, v139, v105, v143)
											mBase = m.M
											v145 = v144
										} else {
											v145 = v139
										}
										v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+6)))
										v148 = v139
										v149 = v146
										v150 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v148)+4)) = uint16(v150)
										*(*uint16)(unsafe.Add(mBase, uint32(v148)+2)) = uint16(v5)
										v154 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
										*(*uint16)(unsafe.Add(mBase, uint32(v148))) = uint16(v154)
										v160 = F_PageAddItemExtended(m, l1, v148, v149&int32(8191), v150, v150)
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return
										} else {
											if v160 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v182 = m.ExcPending
												if v182 != 0 {
													return
												} else {
													F_errmsg_internal(m, int32(426262), int32(0))
													mBase = m.M
													v186 = m.ExcPending
													if v186 != 0 {
														return
													} else {
														F_errfinish(m, int32(522889), int32(736), int32(90715))
														mBase = m.M
														v191 = m.ExcPending
														if v191 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												F_pfree(m, v148)
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								} else {
									v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+2)))
									v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105))))
									v126 = (v116 | v117<<(uint(int32(16))%32)&int32(2147418112) + int32(7)) & int32(-8)
									v127 = F_palloc(m, v126)
									mBase = m.M
									v128 = m.ExcPending
									if v128 != 0 {
										return
									} else {
										if v126 != 0 {
											v129 = F__emscripten_memcpy_bulkmem(m, v127, v105, v126)
											mBase = m.M
											v130 = v129
										} else {
											v130 = v127
										}
										v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+6)))
										v134 = v131&int32(-8192) | v126
										*(*uint16)(unsafe.Add(mBase, uint32(v130)+6)) = uint16(v134)
										v148 = v127
										v149 = v134
										v150 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v148)+4)) = uint16(v150)
										*(*uint16)(unsafe.Add(mBase, uint32(v148)+2)) = uint16(v5)
										v154 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
										*(*uint16)(unsafe.Add(mBase, uint32(v148))) = uint16(v154)
										v160 = F_PageAddItemExtended(m, l1, v148, v149&int32(8191), v150, v150)
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return
										} else {
											if v160 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v182 = m.ExcPending
												if v182 != 0 {
													return
												} else {
													F_errmsg_internal(m, int32(426262), int32(0))
													mBase = m.M
													v186 = m.ExcPending
													if v186 != 0 {
														return
													} else {
														F_errfinish(m, int32(522889), int32(736), int32(90715))
														mBase = m.M
														v191 = m.ExcPending
														if v191 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												F_pfree(m, v148)
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								}
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(426262), int32(0))
							mBase = m.M
							v173 = m.ExcPending
							if v173 != 0 {
								return
							} else {
								F_errfinish(m, int32(522889), int32(731), int32(90715))
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
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
			}
		} else {
			v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+2)))
			v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27))))
			v48 = (v38 | v39<<(uint(int32(16))%32)&int32(2147418112) + int32(7)) & int32(-8)
			v49 = F_palloc(m, v48)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				if v48 != 0 {
					v51 = F__emscripten_memcpy_bulkmem(m, v49, v27, v48)
					mBase = m.M
					v52 = v51
				} else {
					v52 = v49
				}
				v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+6)))
				v56 = v53&int32(-8192) | v48
				*(*uint16)(unsafe.Add(mBase, uint32(v52)+6)) = uint16(v56)
				v69 = v56
				v70 = v49
				v72 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v70)+4)) = uint16(v72)
				*(*uint16)(unsafe.Add(mBase, uint32(v70)+2)) = uint16(v3)
				v76 = int32(base.Ui32(v3) >> (uint(int32(16)) % 32))
				*(*uint16)(unsafe.Add(mBase, uint32(v70))) = uint16(v76)
				v82 = F_PageAddItemExtended(m, l1, v70, v69&int32(8191), v72, v72)
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return
				} else {
					if v82 != 0 {
						F_pfree(m, v70)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+12)))
							if base.Ui32(v87) < base.Ui32(int32(25)) {
								v98 = int32(-1)
							} else {
								v98 = int32(base.Ui32(v87+int32(262120))>>(uint(int32(2))%32))&int32(65535) - int32(1)
							}
							v99 = int32(2)
							v102 = *(*int32)(unsafe.Add(mBase, uint32(l5+v98<<(uint(v99)%32))+24))
							v105 = l5 + v102&int32(32767)
							v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+16)))
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v106)+6)))
							if v108&v99 == int32(0) {
								v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+6)))
								v139 = F_palloc(m, v136&int32(8191))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return
								} else {
									v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+6)))
									v143 = v141 & int32(8191)
									if v143 != 0 {
										v144 = F__emscripten_memcpy_bulkmem(m, v139, v105, v143)
										mBase = m.M
										v145 = v144
									} else {
										v145 = v139
									}
									v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+6)))
									v148 = v139
									v149 = v146
									v150 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v148)+4)) = uint16(v150)
									*(*uint16)(unsafe.Add(mBase, uint32(v148)+2)) = uint16(v5)
									v154 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
									*(*uint16)(unsafe.Add(mBase, uint32(v148))) = uint16(v154)
									v160 = F_PageAddItemExtended(m, l1, v148, v149&int32(8191), v150, v150)
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return
									} else {
										if v160 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v182 = m.ExcPending
											if v182 != 0 {
												return
											} else {
												F_errmsg_internal(m, int32(426262), int32(0))
												mBase = m.M
												v186 = m.ExcPending
												if v186 != 0 {
													return
												} else {
													F_errfinish(m, int32(522889), int32(736), int32(90715))
													mBase = m.M
													v191 = m.ExcPending
													if v191 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											F_pfree(m, v148)
											mBase = m.M
											v165 = m.ExcPending
											if v165 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							} else {
								v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+4)))
								if v113 == int32(65535) {
									v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+6)))
									v139 = F_palloc(m, v136&int32(8191))
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return
									} else {
										v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+6)))
										v143 = v141 & int32(8191)
										if v143 != 0 {
											v144 = F__emscripten_memcpy_bulkmem(m, v139, v105, v143)
											mBase = m.M
											v145 = v144
										} else {
											v145 = v139
										}
										v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+6)))
										v148 = v139
										v149 = v146
										v150 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v148)+4)) = uint16(v150)
										*(*uint16)(unsafe.Add(mBase, uint32(v148)+2)) = uint16(v5)
										v154 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
										*(*uint16)(unsafe.Add(mBase, uint32(v148))) = uint16(v154)
										v160 = F_PageAddItemExtended(m, l1, v148, v149&int32(8191), v150, v150)
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return
										} else {
											if v160 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v182 = m.ExcPending
												if v182 != 0 {
													return
												} else {
													F_errmsg_internal(m, int32(426262), int32(0))
													mBase = m.M
													v186 = m.ExcPending
													if v186 != 0 {
														return
													} else {
														F_errfinish(m, int32(522889), int32(736), int32(90715))
														mBase = m.M
														v191 = m.ExcPending
														if v191 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												F_pfree(m, v148)
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								} else {
									v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+2)))
									v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105))))
									v126 = (v116 | v117<<(uint(int32(16))%32)&int32(2147418112) + int32(7)) & int32(-8)
									v127 = F_palloc(m, v126)
									mBase = m.M
									v128 = m.ExcPending
									if v128 != 0 {
										return
									} else {
										if v126 != 0 {
											v129 = F__emscripten_memcpy_bulkmem(m, v127, v105, v126)
											mBase = m.M
											v130 = v129
										} else {
											v130 = v127
										}
										v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+6)))
										v134 = v131&int32(-8192) | v126
										*(*uint16)(unsafe.Add(mBase, uint32(v130)+6)) = uint16(v134)
										v148 = v127
										v149 = v134
										v150 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v148)+4)) = uint16(v150)
										*(*uint16)(unsafe.Add(mBase, uint32(v148)+2)) = uint16(v5)
										v154 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
										*(*uint16)(unsafe.Add(mBase, uint32(v148))) = uint16(v154)
										v160 = F_PageAddItemExtended(m, l1, v148, v149&int32(8191), v150, v150)
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return
										} else {
											if v160 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v182 = m.ExcPending
												if v182 != 0 {
													return
												} else {
													F_errmsg_internal(m, int32(426262), int32(0))
													mBase = m.M
													v186 = m.ExcPending
													if v186 != 0 {
														return
													} else {
														F_errfinish(m, int32(522889), int32(736), int32(90715))
														mBase = m.M
														v191 = m.ExcPending
														if v191 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												F_pfree(m, v148)
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								}
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(426262), int32(0))
							mBase = m.M
							v173 = m.ExcPending
							if v173 != 0 {
								return
							} else {
								F_errfinish(m, int32(522889), int32(731), int32(90715))
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
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
			}
		}
	}
}
func F_ginEntryInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int64
	_ = v262
	var v270 int32
	_ = v270
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	v2 = l1
	v4 = l3
	v8 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v8)
	v24 = F___memset(m, v15+int32(16), v8, int32(68))
	mBase = m.M
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = int32(43)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = int32(44)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = int32(45)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = int32(46)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(47)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(49)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = int32(50)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+60)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+56)) = l2
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+54)) = uint16(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+52)) = uint16(v8)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+36)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = int32(51)
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+69)) = uint8(base.B2i32(l6 != v55))
	v62 = F_ginFindLeafPage(m, v15+int32(16), v55, v55)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		return
	} else {
		v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
		if v64 < int32(0) {
			v68 = *(*int32)(unsafe.Add(mBase, _consts[12]))
			v74 = *(*int32)(unsafe.Add(mBase, uint32(v68+(v64^int32(-1))<<(uint(int32(2))%32))))
			v82 = v74
		} else {
			v76 = *(*int32)(unsafe.Add(mBase, _consts[13]))
			v82 = v76 + v64<<(uint(int32(13))%32) + int32(-8192)
		}
		v85 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
		v86 = m.T0[v85].(func(*base.Module, int32, int32) int32)(m, v15+int32(16), v62)
		mBase = m.M
		v87 = m.ExcPending
		if v87 != 0 {
			return
		} else {
			if v86 != 0 {
				v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+8)))
				v92 = *(*int32)(unsafe.Add(mBase, uint32(v88<<(uint(int32(2))%32)+v82)+20))
				v95 = v82 + v92&int32(32767)
				v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+4)))
				if v96 == int32(65535) {
					v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+2)))
					v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95))))
					v101 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
					F_LockBuffer(m, v101, int32(0))
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return
					} else {
						F_freeGinBtreeStack(m, v62)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return
						} else {
							v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							F_ginInsertItemPointers(m, v107, v99|v100<<(uint(int32(16))%32), l4, l5, l6)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return
							} else {
								m.G0 = v15 + int32(96)
								return
							}
						}
					}
				} else {
					v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v114 = int32(0)
					v115 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
					if v115 < v114 {
						v119 = *(*int32)(unsafe.Add(mBase, _consts[16]))
						v125 = *(*int32)(unsafe.Add(mBase, uint32(v119+(v115^int32(-1))<<(uint(int32(6))%32))+16))
						v134 = v125
					} else {
						v127 = *(*int32)(unsafe.Add(mBase, _consts[17]))
						v133 = *(*int32)(unsafe.Add(mBase, uint32(v127+v115<<(uint(int32(6))%32)+int32(-64))+16))
						v134 = v133
					}
					F_CheckForSerializableConflictIn(m, v113, v114, v134)
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return
					} else {
						v137 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
						v138 = F_gintuple_get_attrnum(m, l0, v95)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							v142 = F_gintuple_get_key(m, l0, v95, v15+int32(95))
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return
							} else {
								v146 = F_ginReadTuple(m, v95, v15+int32(88))
								mBase = m.M
								v147 = m.ExcPending
								if v147 != 0 {
									return
								} else {
									v148 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
									v151 = F_ginMergeItemPointers(m, l4, l5, v146, v148, v15+int32(84))
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return
									} else {
										v153 = *(*int32)(unsafe.Add(mBase, uint32(v15)+84))
										v156 = F_ginCompressPostingList(m, v151, v153, int32(2712), int32(0))
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return
										} else {
											F_pfree(m, v151)
											mBase = m.M
											v159 = m.ExcPending
											if v159 != 0 {
												return
											} else {
												if v156 != 0 {
													v160 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+95)))
													v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+6)))
													v168 = *(*int32)(unsafe.Add(mBase, uint32(v15)+84))
													v170 = F_GinFormTuple(m, l0, v138, v142, v160, v156, (v161+int32(1))&int32(131070)+int32(8), v168, int32(0))
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return
													} else {
														F_pfree(m, v156)
														mBase = m.M
														v173 = m.ExcPending
														if v173 != 0 {
															return
														} else {
															if v170 != 0 {
																v196 = v170
																F_pfree(m, v146)
																mBase = m.M
																v198 = m.ExcPending
																if v198 != 0 {
																	return
																} else {
																	v199 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v199)
																	v270 = v196
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v270
																	F_ginInsertValue(m, v15+int32(16), v62, v15+int32(8), l6)
																	mBase = m.M
																	v279 = m.ExcPending
																	if v279 != 0 {
																		return
																	} else {
																		F_pfree(m, v270)
																		mBase = m.M
																		v281 = m.ExcPending
																		if v281 != 0 {
																			return
																		} else {
																			m.G0 = v15 + int32(96)
																			return
																		}
																	}
																}
															} else {
																v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																v176 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
																v177 = F_createPostingTree(m, v175, v146, v176, l6, v137)
																mBase = m.M
																v178 = m.ExcPending
																if v178 != 0 {
																	return
																} else {
																	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																	F_ginInsertItemPointers(m, v179, v177, l4, l5, l6)
																	mBase = m.M
																	v181 = m.ExcPending
																	if v181 != 0 {
																		return
																	} else {
																		v182 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+95)))
																		v183 = int32(0)
																		v187 = F_GinFormTuple(m, l0, v138, v142, v182, v183, v183, v183, int32(1))
																		mBase = m.M
																		v188 = m.ExcPending
																		if v188 != 0 {
																			return
																		} else {
																			*(*uint16)(unsafe.Add(mBase, uint32(v187)+2)) = uint16(v177)
																			v191 = int32(base.Ui32(v177) >> (uint(int32(16)) % 32))
																			*(*uint16)(unsafe.Add(mBase, uint32(v187))) = uint16(v191)
																			v193 = int32(65535)
																			*(*uint16)(unsafe.Add(mBase, uint32(v187)+4)) = uint16(v193)
																			v196 = v187
																			F_pfree(m, v146)
																			mBase = m.M
																			v198 = m.ExcPending
																			if v198 != 0 {
																				return
																			} else {
																				v199 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v199)
																				v270 = v196
																				*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v270
																				F_ginInsertValue(m, v15+int32(16), v62, v15+int32(8), l6)
																				mBase = m.M
																				v279 = m.ExcPending
																				if v279 != 0 {
																					return
																				} else {
																					F_pfree(m, v270)
																					mBase = m.M
																					v281 = m.ExcPending
																					if v281 != 0 {
																						return
																					} else {
																						m.G0 = v15 + int32(96)
																						return
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
													v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v176 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
													v177 = F_createPostingTree(m, v175, v146, v176, l6, v137)
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return
													} else {
														v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														F_ginInsertItemPointers(m, v179, v177, l4, l5, l6)
														mBase = m.M
														v181 = m.ExcPending
														if v181 != 0 {
															return
														} else {
															v182 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+95)))
															v183 = int32(0)
															v187 = F_GinFormTuple(m, l0, v138, v142, v182, v183, v183, v183, int32(1))
															mBase = m.M
															v188 = m.ExcPending
															if v188 != 0 {
																return
															} else {
																*(*uint16)(unsafe.Add(mBase, uint32(v187)+2)) = uint16(v177)
																v191 = int32(base.Ui32(v177) >> (uint(int32(16)) % 32))
																*(*uint16)(unsafe.Add(mBase, uint32(v187))) = uint16(v191)
																v193 = int32(65535)
																*(*uint16)(unsafe.Add(mBase, uint32(v187)+4)) = uint16(v193)
																v196 = v187
																F_pfree(m, v146)
																mBase = m.M
																v198 = m.ExcPending
																if v198 != 0 {
																	return
																} else {
																	v199 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v199)
																	v270 = v196
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v270
																	F_ginInsertValue(m, v15+int32(16), v62, v15+int32(8), l6)
																	mBase = m.M
																	v279 = m.ExcPending
																	if v279 != 0 {
																		return
																	} else {
																		F_pfree(m, v270)
																		mBase = m.M
																		v281 = m.ExcPending
																		if v281 != 0 {
																			return
																		} else {
																			m.G0 = v15 + int32(96)
																			return
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
			} else {
				v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v202 = int32(0)
				v203 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
				if v203 < v202 {
					v207 = *(*int32)(unsafe.Add(mBase, _consts[16]))
					v213 = *(*int32)(unsafe.Add(mBase, uint32(v207+(v203^int32(-1))<<(uint(int32(6))%32))+16))
					v222 = v213
				} else {
					v215 = *(*int32)(unsafe.Add(mBase, _consts[17]))
					v221 = *(*int32)(unsafe.Add(mBase, uint32(v215+v203<<(uint(int32(6))%32)+int32(-64))+16))
					v222 = v221
				}
				F_CheckForSerializableConflictIn(m, v201, v202, v222)
				mBase = m.M
				v224 = m.ExcPending
				if v224 != 0 {
					return
				} else {
					v225 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
					v228 = F_ginCompressPostingList(m, l4, l5, int32(2712), int32(0))
					mBase = m.M
					v229 = m.ExcPending
					if v229 != 0 {
						return
					} else {
						if v228 != 0 {
							v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228)+6)))
							v238 = F_GinFormTuple(m, l0, v2, l2, v4, v228, (v230+int32(1))&int32(131070)+int32(8), l5, int32(0))
							mBase = m.M
							v239 = m.ExcPending
							if v239 != 0 {
								return
							} else {
								F_pfree(m, v228)
								mBase = m.M
								v241 = m.ExcPending
								if v241 != 0 {
									return
								} else {
									if v238 != 0 {
										v259 = v238
										if l6 == int32(0) {
											v270 = v259
										} else {
											v262 = *(*int64)(unsafe.Add(mBase, uint32(l6)+16))
											*(*int64)(unsafe.Add(mBase, uint32(l6)+16)) = v262 + int64(1)
											v270 = v259
										}
										*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v270
										F_ginInsertValue(m, v15+int32(16), v62, v15+int32(8), l6)
										mBase = m.M
										v279 = m.ExcPending
										if v279 != 0 {
											return
										} else {
											F_pfree(m, v270)
											mBase = m.M
											v281 = m.ExcPending
											if v281 != 0 {
												return
											} else {
												m.G0 = v15 + int32(96)
												return
											}
										}
									} else {
										v243 = int32(0)
										v247 = F_GinFormTuple(m, l0, v2, l2, v4, v243, v243, v243, int32(1))
										mBase = m.M
										v248 = m.ExcPending
										if v248 != 0 {
											return
										} else {
											v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v250 = F_createPostingTree(m, v249, l4, l5, l6, v225)
											mBase = m.M
											v251 = m.ExcPending
											if v251 != 0 {
												return
											} else {
												*(*uint16)(unsafe.Add(mBase, uint32(v247)+2)) = uint16(v250)
												v253 = int32(65535)
												*(*uint16)(unsafe.Add(mBase, uint32(v247)+4)) = uint16(v253)
												v256 = int32(base.Ui32(v250) >> (uint(int32(16)) % 32))
												*(*uint16)(unsafe.Add(mBase, uint32(v247))) = uint16(v256)
												v259 = v247
												if l6 == int32(0) {
													v270 = v259
												} else {
													v262 = *(*int64)(unsafe.Add(mBase, uint32(l6)+16))
													*(*int64)(unsafe.Add(mBase, uint32(l6)+16)) = v262 + int64(1)
													v270 = v259
												}
												*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v270
												F_ginInsertValue(m, v15+int32(16), v62, v15+int32(8), l6)
												mBase = m.M
												v279 = m.ExcPending
												if v279 != 0 {
													return
												} else {
													F_pfree(m, v270)
													mBase = m.M
													v281 = m.ExcPending
													if v281 != 0 {
														return
													} else {
														m.G0 = v15 + int32(96)
														return
													}
												}
											}
										}
									}
								}
							}
						} else {
							v243 = int32(0)
							v247 = F_GinFormTuple(m, l0, v2, l2, v4, v243, v243, v243, int32(1))
							mBase = m.M
							v248 = m.ExcPending
							if v248 != 0 {
								return
							} else {
								v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v250 = F_createPostingTree(m, v249, l4, l5, l6, v225)
								mBase = m.M
								v251 = m.ExcPending
								if v251 != 0 {
									return
								} else {
									*(*uint16)(unsafe.Add(mBase, uint32(v247)+2)) = uint16(v250)
									v253 = int32(65535)
									*(*uint16)(unsafe.Add(mBase, uint32(v247)+4)) = uint16(v253)
									v256 = int32(base.Ui32(v250) >> (uint(int32(16)) % 32))
									*(*uint16)(unsafe.Add(mBase, uint32(v247))) = uint16(v256)
									v259 = v247
									if l6 == int32(0) {
										v270 = v259
									} else {
										v262 = *(*int64)(unsafe.Add(mBase, uint32(l6)+16))
										*(*int64)(unsafe.Add(mBase, uint32(l6)+16)) = v262 + int64(1)
										v270 = v259
									}
									*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v270
									F_ginInsertValue(m, v15+int32(16), v62, v15+int32(8), l6)
									mBase = m.M
									v279 = m.ExcPending
									if v279 != 0 {
										return
									} else {
										F_pfree(m, v270)
										mBase = m.M
										v281 = m.ExcPending
										if v281 != 0 {
											return
										} else {
											m.G0 = v15 + int32(96)
											return
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
func F_ginFreeScanKeys(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[22])))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[23])))
	if v6 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v10 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[24])))
	F_MemoryContextReset(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L27
	}
L7:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[25])))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11+v10<<(uint(int32(2))%32))))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	if v16 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	F_ReleaseBuffer(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+644))
	if v19 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	return
L13:
	;
	goto L11
L14:
	;
	F_pfree(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	if v22 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	F_pfree(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L12
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	if v25 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	F_tbm_free(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L12
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v29 = v10 + int32(1)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[23])))
	if base.Ui32(v29) < base.Ui32(v30) {
		v10 = v29
		goto L7
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	goto L8
L27:
	;
	v41 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[25]))) = v41
	*(*int64)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[22]))) = v41
	goto L3
}
func F_ginGetBAEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	v7 = l0 + int32(20)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)))
	if v8 != 0 {
		v15 = int32(0)
		if v15 == int32(0) {
			return int32(0)
		} else {
			v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+22)))
			*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v20)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v22
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)))
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v24)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v27
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+24)))
			if v29 != int32(1) {
				return v26
			} else {
				if base.Ui32(v27) < base.Ui32(int32(2)) {
					return v26
				} else {
					F_pg_qsort(m, v26, v27, int32(6), int32(34))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						return v26
					}
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		v11 = m.T0[v10].(func(*base.Module, int32) int32)(m, v7)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = v11
			if v15 == int32(0) {
				return int32(0)
			} else {
				v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+22)))
				*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v20)
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v22
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)))
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v24)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = v27
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+24)))
				if v29 != int32(1) {
					return v26
				} else {
					if base.Ui32(v27) < base.Ui32(int32(2)) {
						return v26
					} else {
						F_pg_qsort(m, v26, v27, int32(6), int32(34))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							return v26
						}
					}
				}
			}
		}
	}
}
func F_ginInsertCleanup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v375 int32
	_ = v375
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v527 int32
	_ = v527
	var v555 int32
	_ = v555
	var v559 int64
	_ = v559
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v589 int64
	_ = v589
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int64
	_ = v640
	var v641 int64
	_ = v641
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v658 int32
	_ = v658
	var v673 int64
	_ = v673
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v698 int64
	_ = v698
	var v701 int32
	_ = v701
	var v708 int64
	_ = v708
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v723 int32
	_ = v723
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v789 int32
	_ = v789
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v835 int32
	_ = v835
	var v862 int32
	_ = v862
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v905 int64
	_ = v905
	var v907 int64
	_ = v907
	var v909 int64
	_ = v909
	var v911 int64
	_ = v911
	var v913 int64
	_ = v913
	var v915 int64
	_ = v915
	var v917 int64
	_ = v917
	var v923 int32
	_ = v923
	var v926 int64
	_ = v926
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v942 int32
	_ = v942
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1004 int32
	_ = v1004
	var v1027 int32
	_ = v1027
	var v1033 int32
	_ = v1033
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1085 int32
	_ = v1085
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1189 int32
	_ = v1189
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1213 int32
	_ = v1213
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1238 int32
	_ = v1238
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	v30 = m.G0
	v32 = v30 - int32(256)
	m.G0 = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l3 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v32 + int32(256)
	return
L2:
	;
	v79 = F_ReadBuffer(m, v34, int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L16
	}
L3:
	;
	F_LockPage(m, v34, int32(0), int32(7))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v51 = m.G0
	v53 = v51 - int32(16)
	m.G0 = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v34)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+12)) = int32(16973824)
	v60 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v57
	v66 = F_LockAcquire(m, v53, int32(7), v60, int32(1))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L14
	}
L6:
	;
	return
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v42 = *(*int32)(unsafe.Add(mBase, _consts[19]))
	if v40 != int32(-1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v45 = v40
	goto L10
L9:
	;
	v45 = v42
	goto L10
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	if v47 == int32(4) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v50 = v45
	goto L13
L12:
	;
	v50 = v42
	goto L13
L13:
	;
	v77 = v50
	goto L2
L14:
	;
	m.G0 = v53 + int32(16)
	if v66 == int32(0) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _consts[21]))
	v77 = v74
	goto L2
L16:
	;
	F_LockBuffer(m, v79, int32(1))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	if v79 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+24))
	if v102 == int32(-1) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v87+(v79^int32(-1))<<(uint(int32(2))%32))))
	v101 = v93
	goto L18
L20:
	;
	goto L21
L21:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v101 = v95 + v79<<(uint(int32(13))%32) + int32(-8192)
	goto L18
L22:
	;
	F_UnlockReleaseBuffer(m, v79)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v101)+28))
	v112 = F_ReadBuffer(m, v34, v102)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	F_UnlockPage(m, v34, int32(0), int32(7))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	goto L1
L27:
	;
	F_LockBuffer(m, v112, int32(1))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	if v112 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	F_LockBuffer(m, v79, int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L6
	} else {
		goto L33
	}
L30:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v120+(v112^int32(-1))<<(uint(int32(2))%32))))
	v134 = v126
	goto L29
L31:
	;
	goto L32
L32:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v134 = v128 + v112<<(uint(int32(13))%32) + int32(-8192)
	goto L29
L33:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v144 = F_AllocSetContextCreateInternal(m, v139, int32(65211), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	v146 = int32(4554240)
	v147 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v144
	v151 = F_palloc(m, int32(512))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v151
	v155 = F_palloc(m, int32(128))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+20)) = int64(549755813888)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v155
	F_ginInitBA(m, v32+int32(28))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = l0
	v181 = v134
	v182 = v112
	v185 = v102
	v190 = int32(0)
	goto L38
L38:
	;
	F_processPendingPage(m, v32+int32(28), v32+int32(12), v181, int32(1))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v216 = base.B2i32(v185 == v111)&(l1^int32(1)) | v190
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+16)))
	v218 = v181 + v217
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	if v219 != int32(-1) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L6
	} else {
		goto L184
	}
L43:
	;
	F_UnlockReleaseBuffer(m, v182)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L6
	} else {
		goto L183
	}
L44:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+6)))
	if v222&int32(32) == int32(0) {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+12)))
	F_LockBuffer(m, v182, int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L6
	} else {
		goto L49
	}
L47:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
	if base.Ui32(v227) < base.Ui32(v77<<(uint(int32(10))%32)) {
		goto L43
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v229) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v240 = int32(base.Ui32(v229+int32(262120)) >> (uint(int32(2)) % 32))
	goto L52
L51:
	;
	v240 = int32(0)
	goto L52
L52:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(28))+16))
	v246 = m.G0
	v247 = int32(16)
	v248 = v246 - v247
	m.G0 = v248
	v251 = v32 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v251)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v245
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	*(*uint8)(unsafe.Add(mBase, uint32(v251)+12)) = uint8(base.B2i32(v255 == int32(4159512)))
	*(*int32)(unsafe.Add(mBase, uint32(v251)+4)) = int32(790)
	m.G0 = v248 + v247
	goto L53
L53:
	;
	v272 = F_ginGetBAEntry(m, v32+int32(28), v32, v32+int32(4), v32+int32(3), v32+int32(8))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	if v272 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v277 = v272
	goto L58
L56:
	;
	goto L57
L57:
	;
	F_LockBuffer(m, v79, int32(2))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L6
	} else {
		goto L64
	}
L58:
	;
	v303 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32))))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v305 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32)+3)))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	F_ginEntryInsert(m, l0, v303, v304, v305, v277, v306, int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L6
	} else {
		goto L60
	}
L59:
	;
	goto L57
L60:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	v321 = F_ginGetBAEntry(m, v32+int32(28), v32, v32+int32(4), v32+int32(3), v32+int32(8))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	if v321 != 0 {
		v277 = v321
		goto L58
	} else {
		goto L63
	}
L63:
	;
	goto L59
L64:
	;
	F_LockBuffer(m, v182, int32(1))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	v358 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v358) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v496 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+16)))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v181+v496)))
	F_UnlockReleaseBuffer(m, v182)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L6
	} else {
		goto L81
	}
L67:
	;
	v366 = int32(base.Ui32(v358+int32(262120)) >> (uint(int32(2)) % 32))
	goto L69
L68:
	;
	v366 = int32(0)
	goto L69
L69:
	;
	v367 = int32(65535)
	if v366&v367 == v240&v367 {
		goto L66
	} else {
		goto L70
	}
L70:
	;
	F_ginInitBA(m, v32+int32(28))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	F_processPendingPage(m, v32+int32(28), v32+int32(12), v181, (v240+int32(1))&int32(65535))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L6
	} else {
		goto L72
	}
L72:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(28))+16))
	v391 = m.G0
	v392 = int32(16)
	v393 = v391 - v392
	m.G0 = v393
	v396 = v32 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v396)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v396))) = v390
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	*(*uint8)(unsafe.Add(mBase, uint32(v396)+12)) = uint8(base.B2i32(v400 == int32(4159512)))
	*(*int32)(unsafe.Add(mBase, uint32(v396)+4)) = int32(790)
	m.G0 = v393 + v392
	goto L73
L73:
	;
	v417 = F_ginGetBAEntry(m, v32+int32(28), v32, v32+int32(4), v32+int32(3), v32+int32(8))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	if v417 == int32(0) {
		goto L66
	} else {
		goto L75
	}
L75:
	;
	v424 = v417
	goto L76
L76:
	;
	v450 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32))))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v452 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32)+3)))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	F_ginEntryInsert(m, l0, v450, v451, v452, v424, v453, int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L6
	} else {
		goto L78
	}
L77:
	;
	goto L66
L78:
	;
	v465 = F_ginGetBAEntry(m, v32+int32(28), v32, v32+int32(4), v32+int32(3), v32+int32(8))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	if v465 != 0 {
		v424 = v465
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	if v79 < int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v514 = v512 + int32(24)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v514)))
	v517 = v512 + int32(40)
	v519 = v512 + int32(36)
	v527 = v515
	goto L86
L83:
	;
	v504 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v504+(v79^int32(-1))<<(uint(int32(2))%32))))
	v512 = v506
	goto L82
L84:
	;
	goto L85
L85:
	;
	v508 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v512 = v508 + v79<<(uint(int32(13))%32) + int32(-8192)
	goto L82
L86:
	;
	v555 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+248)) = v555
	v559 = int64(0)
	if v527 == v498 {
		v646 = v498
		v651 = v555
		v658 = v555
		v673 = v559
		goto L88
	} else {
		goto L89
	}
L87:
	;
	F_LockBuffer(m, v79, int32(0))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L6
	} else {
		goto L168
	}
L88:
	;
	if l4 != 0 {
		goto L102
	} else {
		goto L103
	}
L89:
	;
	v564 = v527
	v567 = v555
	v589 = v559
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32-int32(-64)+v567<<(uint(int32(2))%32)))) = v564
	v596 = F_ReadBuffer(m, v34, v564)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L6
	} else {
		goto L92
	}
L91:
	;
	v646 = v638
	v651 = v634
	v658 = v639
	v673 = v641
	goto L88
L92:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v32)+248))
	v601 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(128)+v600<<(uint(v601)%32)))) = v596
	F_LockBuffer(m, v596, v601)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v32)+248))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(128)+v610<<(uint(int32(2))%32))))
	if v614 < int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v634 = v610 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+248)) = v634
	v636 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v632)+16)))
	v637 = v632 + v636
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	v639 = base.B2i32(v638 != v498)
	v640 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v637)+4)))
	v641 = v589 + v640
	if int32(15) < v634 {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v618 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v618+(v614^int32(-1))<<(uint(int32(2))%32))))
	v632 = v624
	goto L94
L96:
	;
	goto L97
L97:
	;
	v626 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v632 = v626 + v614<<(uint(int32(13))%32) + int32(-8192)
	goto L94
L98:
	;
	v646 = v638
	v651 = v634
	v658 = v639
	v673 = v641
	goto L88
L99:
	;
	goto L100
L100:
	;
	if v638 != v498 {
		v564 = v638
		v567 = v634
		v589 = v641
		goto L90
	} else {
		goto L101
	}
L101:
	;
	goto L91
L102:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+28)) = v674 + v651
	goto L104
L103:
	;
	goto L104
L104:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677)+118)))
	if v678 != int32(112) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v690 = int32(4548900)
	v692 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v692 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v514))) = v646
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v32)+248))
	v698 = *(*int64)(unsafe.Add(mBase, uint32(v517)))
	*(*int64)(unsafe.Add(mBase, uint32(v517))) = v698 - v673
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	*(*int32)(unsafe.Add(mBase, uint32(v519))) = v701 - v697
	if v646 == int32(-1) {
		goto L113
	} else {
		goto L114
	}
L106:
	;
	v682 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v682 <= int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v34)+32))
	if v685 != 0 {
		goto L105
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	F_XLogEnsureRecordSpace(m, v651, int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L6
	} else {
		goto L112
	}
L110:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	if v686 != 0 {
		goto L105
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	goto L105
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v512+int32(28)))) = int32(-1)
	v708 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v517))) = v708
	*(*int64)(unsafe.Add(mBase, uint32(v512+int32(32)))) = v708
	goto L115
L114:
	;
	goto L115
L115:
	;
	v712 = int32(80)
	*(*uint16)(unsafe.Add(mBase, uint32(v512)+12)) = uint16(v712)
	F_MarkBufferDirty(m, v79)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	v716 = int32(0)
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v32)+248))
	if v716 < v717 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v723 = v716
	goto L120
L118:
	;
	v789 = v717
	goto L119
L119:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v812)+118)))
	if v813 != int32(112) {
		v1004 = v789
		goto L130
	} else {
		goto L131
	}
L120:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(128)+v723<<(uint(int32(2))%32))))
	if v754 < int32(0) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v789 = v781
	goto L119
L122:
	;
	v773 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v772)+16)))
	v775 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v773+v772)+6)) = uint16(v775)
	F_MarkBufferDirty(m, v754)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L6
	} else {
		goto L126
	}
L123:
	;
	v758 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v758+(v754^int32(-1))<<(uint(int32(2))%32))))
	v772 = v764
	goto L122
L124:
	;
	goto L125
L125:
	;
	v766 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v772 = v766 + v754<<(uint(int32(13))%32) + int32(-8192)
	goto L122
L126:
	;
	v780 = v723 + int32(1)
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v32)+248))
	if v780 < v781 {
		v723 = v780
		goto L120
	} else {
		goto L127
	}
L127:
	;
	goto L121
L128:
	;
	if v658 != 0 {
		v527 = v646
		goto L86
	} else {
		goto L167
	}
L129:
	;
	v1152 = int32(4548900)
	v1154 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v1154 - int32(1)
	goto L128
L130:
	;
	v1027 = int32(0)
	if v1004 <= v1027 {
		goto L129
	} else {
		goto L156
	}
L131:
	;
	v817 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v817 <= int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v34)+32))
	if v820 != 0 {
		v1004 = v789
		goto L130
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L6
	} else {
		goto L137
	}
L135:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	if v821 != 0 {
		v1004 = v789
		goto L130
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	v824 = int32(0)
	F_XLogRegisterBuffer(m, v824, v79, int32(14))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v32)+248))
	if int32(0) < v829 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v835 = v824
	goto L142
L140:
	;
	goto L141
L141:
	;
	v905 = *(*int64)(unsafe.Add(mBase, uint32(v514)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+240)) = v905
	v907 = *(*int64)(unsafe.Add(mBase, uint32(v514)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+232)) = v907
	v909 = *(*int64)(unsafe.Add(mBase, uint32(v514)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+224)) = v909
	v911 = *(*int64)(unsafe.Add(mBase, uint32(v514)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+216)) = v911
	v913 = *(*int64)(unsafe.Add(mBase, uint32(v514)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+208)) = v913
	v915 = *(*int64)(unsafe.Add(mBase, uint32(v514)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+200)) = v915
	v917 = *(*int64)(unsafe.Add(mBase, uint32(v514)))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+192)) = v917
	F_XLogRegisterData(m, v32+int32(192), int32(64))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L6
	} else {
		goto L146
	}
L142:
	;
	v862 = v835 + int32(1)
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(128)+v835<<(uint(int32(2))%32))))
	F_XLogRegisterBuffer(m, v862&int32(255), v870, int32(6))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L6
	} else {
		goto L144
	}
L143:
	;
	goto L141
L144:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v32)+248))
	if v862 < v874 {
		v835 = v862
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v926 = F_XLogInsert(m, int32(13), int32(128))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L6
	} else {
		goto L147
	}
L147:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v512))) = base.I64_rotr(v926, int64(32))
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v32)+248))
	if v931 <= int32(0) {
		goto L129
	} else {
		goto L148
	}
L148:
	;
	v942 = int32(0)
	goto L149
L149:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(128)+v942<<(uint(int32(2))%32))))
	if v973 < int32(0) {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	v1004 = v996
	goto L130
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v991)+4)) = base.I32_wrap_i64(v926)
	*(*int32)(unsafe.Add(mBase, uint32(v991))) = base.I32_wrap_i64(int64(base.Ui64(v926) >> (uint(int64(32)) % 64)))
	v995 = v942 + int32(1)
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v32)+248))
	if v995 < v996 {
		v942 = v995
		goto L149
	} else {
		goto L155
	}
L152:
	;
	v977 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v977+(v973^int32(-1))<<(uint(int32(2))%32))))
	v991 = v983
	goto L151
L153:
	;
	goto L154
L154:
	;
	v985 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v991 = v985 + v973<<(uint(int32(13))%32) + int32(-8192)
	goto L151
L155:
	;
	goto L150
L156:
	;
	v1033 = v1027
	goto L157
L157:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(128)+v1033<<(uint(int32(2))%32))))
	F_UnlockReleaseBuffer(m, v1064)
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L6
	} else {
		goto L159
	}
L158:
	;
	v1071 = int32(0)
	v1072 = int32(4548900)
	v1074 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v1074 - int32(1)
	if l2 == v1071 {
		goto L128
	} else {
		goto L161
	}
L159:
	;
	v1068 = v1033 + int32(1)
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v32)+248))
	if v1068 < v1069 {
		v1033 = v1068
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	if v1069 <= int32(0) {
		goto L128
	} else {
		goto L162
	}
L162:
	;
	v1085 = v1071
	goto L163
L163:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v32-int32(-64)+v1085<<(uint(int32(2))%32))))
	F_RecordFreeIndexPage(m, v34, v1116)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L6
	} else {
		goto L165
	}
L164:
	;
	goto L128
L165:
	;
	v1120 = v1085 + int32(1)
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v32)+248))
	if v1120 < v1121 {
		v1085 = v1120
		goto L163
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	goto L87
L168:
	;
	if (base.B2i32(v498 == int32(-1))|v216)&int32(1) == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	F_MemoryContextReset(m, v144)
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L6
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	F_UnlockPage(m, v34, int32(0), int32(7))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L6
	} else {
		goto L176
	}
L172:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v1202 = F_palloc(m, v1199<<(uint(int32(2))%32))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L6
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v1202
	v1205 = F_palloc(m, v1199)
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L6
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v1205
	F_ginInitBA(m, v32+int32(28))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L6
	} else {
		goto L175
	}
L175:
	;
	v1238 = v498
	goto L42
L176:
	;
	F_ReleaseBuffer(m, v79)
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L6
	} else {
		goto L177
	}
L177:
	;
	if l2 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	F_FreeSpaceMapVacuum(m, v34)
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L6
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v147
	F_MemoryContextDelete(m, v144)
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L6
	} else {
		goto L182
	}
L181:
	;
	goto L180
L182:
	;
	goto L1
L183:
	;
	v1238 = v219
	goto L42
L184:
	;
	v1260 = F_ReadBuffer(m, v34, v1238)
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L6
	} else {
		goto L185
	}
L185:
	;
	F_LockBuffer(m, v1260, int32(1))
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L6
	} else {
		goto L186
	}
L186:
	;
	if v1260 < int32(0) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1268+(v1260^int32(-1))<<(uint(int32(2))%32))))
	v181 = v1274
	v182 = v1260
	v185 = v1238
	v190 = v216
	goto L38
L188:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v181 = v1276 + v1260<<(uint(int32(13))%32) + int32(-8192)
	v182 = v1260
	v185 = v1238
	v190 = v216
	goto L38
}
func F_ginPlaceToPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
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
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v203 int64
	_ = v203
	var v204 int32
	_ = v204
	var v212 int64
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int64
	_ = v244
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v716 int64
	_ = v716
	var v717 int32
	_ = v717
	var v718 int64
	_ = v718
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v751 int32
	_ = v751
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v765 int32
	_ = v765
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v797 int32
	_ = v797
	var v805 int32
	_ = v805
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v822 int32
	_ = v822
	v7 = int32(0)
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v20 < v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v39 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v39
	v46 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v51 = F_AllocSetContextCreateInternal(m, v46, int32(65491), v39, int32(8192), int32(8388608))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24+(v20^int32(-1))<<(uint(int32(2))%32))))
	v38 = v30
	goto L1
L3:
	;
	goto L4
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v38 = v32 + v20<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	v55 = int32(4554240)
	v56 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v51
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)))
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38+v59)+6)))
	v63 = v61 & int32(1)
	if v61&int32(2) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v96 = m.T0[v95].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, l0, v88, l1, l2, l3, v16+int32(-12), v16+int32(-4), v16+int32(-8))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L18
	}
L8:
	;
	v85 = v63 | int32(2)
	v86 = v7
	goto L7
L9:
	;
	goto L10
L10:
	;
	if l4 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v71+(l4^int32(-1))<<(uint(int32(2))%32))))
	v85 = v63
	v86 = v77
	goto L7
L12:
	;
	goto L13
L13:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v85 = v63
	v86 = v79 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L7
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L5
	} else {
		goto L209
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v56
	F_MemoryContextDelete(m, v51)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L5
	} else {
		goto L208
	}
L16:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v223 = F_GinNewBuffer(m, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L59
	}
L17:
	;
	v98 = int32(4548900)
	v100 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v100 + int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+48))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+118)))
	if v106 != int32(112) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	switch v96 {
	case 0:
		v797 = int32(1)
		goto L15
	case 1:
		goto L17
	case 2:
		goto L16
	default:
		goto L14
	}
L19:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	m.T0[v120].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l0, v118, l1, l2, l3, v119)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L28
	}
L20:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v110 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v104)+32))
	if v113 != 0 {
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
	if v115 != 0 {
		goto L19
	} else {
		goto L26
	}
L24:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v104)+40))
	if v114 != 0 {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	goto L19
L28:
	;
	if l4 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+48))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+118)))
	if v152 != int32(112) {
		goto L40
	} else {
		goto L41
	}
L30:
	;
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+16)))
	v126 = v86 + v125
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+6)))
	v129 = v127 & int32(65471)
	*(*uint16)(unsafe.Add(mBase, uint32(v126)+6)) = uint16(v129)
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+48))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+118)))
	if v135 != int32(112) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v139 <= int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v133)+32))
	if v142 != 0 {
		goto L29
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
	if v144 != 0 {
		goto L29
	} else {
		goto L38
	}
L36:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v133)+40))
	if v143 != 0 {
		goto L29
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	F_XLogRegisterBuffer(m, int32(1), l4, int32(8))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	goto L29
L40:
	;
	v215 = int32(4548900)
	v217 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v218 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v217 - v218
	v797 = v218
	goto L15
L41:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v156 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v150)+32))
	if v159 != 0 {
		goto L40
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
	if v161 != 0 {
		goto L40
	} else {
		goto L47
	}
L45:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v150)+40))
	if v160 != 0 {
		goto L40
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+50)) = uint16(v85)
	F_XLogRegisterData(m, v16+int32(-14), int32(2))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	if l4 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if l4 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	goto L51
L51:
	;
	v203 = F_XLogInsert(m, int32(13), int32(32))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L5
	} else {
		goto L57
	}
L52:
	;
	v187 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = base.I32_rotr(v186, v187)
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+16)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v86+v190)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = base.I32_rotr(v192, v187)
	F_XLogRegisterData(m, v16+int32(-48), int32(8))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L5
	} else {
		goto L56
	}
L53:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v171+(l4^int32(-1))<<(uint(int32(6))%32))+16))
	v186 = v177
	goto L52
L54:
	;
	goto L55
L55:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v179+l4<<(uint(int32(6))%32)+int32(-64))+16))
	v186 = v185
	goto L52
L56:
	;
	goto L51
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = base.I64_rotr(v203, int64(32))
	if l4 == int32(0) {
		goto L40
	} else {
		goto L58
	}
L58:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v86)+4)) = uint32(v203)
	v212 = int64(base.Ui64(v203) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v86))) = uint32(v212)
	goto L40
L59:
	;
	if l5 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v38+v238)))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v242
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v241)))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+40)) = uint16(v85)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v244
	if l4 != 0 {
		goto L66
	} else {
		goto L67
	}
L61:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v227 == int32(1) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v230 + int32(1)
	goto L60
L63:
	;
	goto L64
L64:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v234 + int32(1)
	goto L60
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v273
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v275 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L66:
	;
	if l4 < int32(0) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	goto L68
L68:
	;
	v270 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v270
	v273 = v270
	goto L65
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v265
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+16)))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v86+v267)))
	v273 = v269
	goto L65
L70:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v250+(l4^int32(-1))<<(uint(int32(6))%32))+16))
	v265 = v256
	goto L69
L71:
	;
	goto L72
L72:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v258+l4<<(uint(int32(6))%32)+int32(-64))+16))
	v265 = v264
	goto L69
L73:
	;
	v573 = int32(4548900)
	v575 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v575 + int32(1)
	F_MarkBufferDirty(m, v223)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L5
	} else {
		goto L132
	}
L74:
	;
	v568 = v564
	v572 = v566
	goto L73
L75:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v522 < int32(0) {
		goto L124
	} else {
		goto L125
	}
L76:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v279 = F_GinNewBuffer(m, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L5
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v240
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	v457 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v456)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v456+v457))) = v240
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v461 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v460)+16)))
	v462 = v460 + v461
	v463 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v462)+6)))
	v465 = v463 | int32(64)
	*(*uint16)(unsafe.Add(mBase, uint32(v462)+6)) = uint16(v465)
	if v223 < int32(0) {
		goto L115
	} else {
		goto L116
	}
L79:
	;
	if l5 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v294 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v294
	v297 = v85 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+40)) = uint16(v297)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	v300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v299)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v299+v300))) = v294
	if v223 < int32(0) {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v283 == int32(1) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v286 + int32(1)
	goto L80
L83:
	;
	goto L84
L84:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v290 + int32(1)
	goto L80
L85:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v324 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v323)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v323+v324))) = v322
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	v328 = F_PageGetTempPage(m, v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L5
	} else {
		goto L89
	}
L86:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v307+(v223^int32(-1))<<(uint(int32(6))%32))+16))
	v322 = v313
	goto L85
L87:
	;
	goto L88
L88:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v315+v223<<(uint(int32(6))%32)+int32(-64))+16))
	v322 = v321
	goto L85
L89:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v330)+16)))
	v333 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v330+v331)+6)))
	v335 = v333 & int32(65405)
	F_PageInit(m, v328, int32(8192), int32(8))
	mBase = m.M
	v339 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v328)+16)))
	v340 = v328 + v339
	*(*int32)(unsafe.Add(mBase, uint32(v340))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v340)+6)) = uint16(v335)
	goto L90
L90:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v279 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	if v223 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L92:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v348+(v279^int32(-1))<<(uint(int32(6))%32))+16))
	v363 = v354
	goto L91
L93:
	;
	goto L94
L94:
	;
	v356 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v356+v279<<(uint(int32(6))%32)+int32(-64))+16))
	v363 = v362
	goto L91
L95:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	m.T0[v344].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l0, v328, v363, v364, v383, v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L5
	} else {
		goto L99
	}
L96:
	;
	v368 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v368+(v223^int32(-1))<<(uint(int32(6))%32))+16))
	v383 = v374
	goto L95
L97:
	;
	goto L98
L98:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v376+v223<<(uint(int32(6))%32)+int32(-64))+16))
	v383 = v382
	goto L95
L99:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v387 < int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v406 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v405)+16)))
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406+v405)+6)))
	if v408&int32(2) == int32(0) {
		v564 = v279
		v566 = v328
		goto L74
	} else {
		goto L104
	}
L101:
	;
	v391 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v391+(v387^int32(-1))<<(uint(int32(2))%32))))
	v405 = v397
	goto L100
L102:
	;
	goto L103
L103:
	;
	v399 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v405 = v399 + v387<<(uint(int32(13))%32) + int32(-8192)
	goto L100
L104:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v387 < int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if v279 < int32(0) {
		goto L110
	} else {
		goto L111
	}
L106:
	;
	v417 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v417+(v387^int32(-1))<<(uint(int32(6))%32))+16))
	v432 = v423
	goto L105
L107:
	;
	goto L108
L108:
	;
	v425 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v425+v387<<(uint(int32(6))%32)+int32(-64))+16))
	v432 = v431
	goto L105
L109:
	;
	F_PredicateLockPageSplit(m, v413, v432, v451)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L5
	} else {
		goto L113
	}
L110:
	;
	v436 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v436+(v279^int32(-1))<<(uint(int32(6))%32))+16))
	v451 = v442
	goto L109
L111:
	;
	goto L112
L112:
	;
	v444 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v444+v279<<(uint(int32(6))%32)+int32(-64))+16))
	v451 = v450
	goto L109
L113:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v519 = v279
	v521 = v328
	v522 = v454
	goto L75
L114:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v487 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v486)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v486+v487))) = v485
	v490 = int32(0)
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v493 < v490 {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	v470 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v470+(v223^int32(-1))<<(uint(int32(6))%32))+16))
	v485 = v476
	goto L114
L116:
	;
	goto L117
L117:
	;
	v478 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v478+v223<<(uint(int32(6))%32)+int32(-64))+16))
	v485 = v484
	goto L114
L118:
	;
	v512 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v511)+16)))
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512+v511)+6)))
	if v514&int32(2) == int32(0) {
		v568 = v490
		v572 = v490
		goto L73
	} else {
		goto L122
	}
L119:
	;
	v497 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v497+(v493^int32(-1))<<(uint(int32(2))%32))))
	v511 = v503
	goto L118
L120:
	;
	goto L121
L121:
	;
	v505 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v511 = v505 + v493<<(uint(int32(13))%32) + int32(-8192)
	goto L118
L122:
	;
	v519 = v490
	v521 = v490
	v522 = v493
	goto L75
L123:
	;
	if v223 < int32(0) {
		goto L128
	} else {
		goto L129
	}
L124:
	;
	v527 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v527+(v522^int32(-1))<<(uint(int32(6))%32))+16))
	v542 = v533
	goto L123
L125:
	;
	goto L126
L126:
	;
	v535 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v535+v522<<(uint(int32(6))%32)+int32(-64))+16))
	v542 = v541
	goto L123
L127:
	;
	F_PredicateLockPageSplit(m, v523, v542, v561)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L5
	} else {
		goto L131
	}
L128:
	;
	v546 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v546+(v223^int32(-1))<<(uint(int32(6))%32))+16))
	v561 = v552
	goto L127
L129:
	;
	goto L130
L130:
	;
	v554 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v554+v223<<(uint(int32(6))%32)+int32(-64))+16))
	v561 = v560
	goto L127
L131:
	;
	v564 = v519
	v566 = v521
	goto L74
L132:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_MarkBufferDirty(m, v581)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L5
	} else {
		goto L133
	}
L133:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v584 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	goto L162
L135:
	;
	F_MarkBufferDirty(m, v568)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L5
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	goto L155
L138:
	;
	goto L140
L139:
	;
	if v568 < int32(0) {
		goto L144
	} else {
		goto L145
	}
L140:
	;
	v590 = F__emscripten_memcpy_bulkmem(m, v38, v572, int32(8192))
	mBase = m.M
	goto L142
L142:
	;
	goto L139
L143:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	goto L148
L144:
	;
	v595 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v595+(v568^int32(-1))<<(uint(int32(2))%32))))
	v609 = v601
	goto L143
L145:
	;
	goto L146
L146:
	;
	v603 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v609 = v603 + v568<<(uint(int32(13))%32) + int32(-8192)
	goto L143
L147:
	;
	if v223 < int32(0) {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	v612 = F__emscripten_memcpy_bulkmem(m, v609, v610, int32(8192))
	mBase = m.M
	goto L150
L150:
	;
	goto L147
L151:
	;
	v617 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v617+(v223^int32(-1))<<(uint(int32(2))%32))))
	v652 = v623
	goto L134
L152:
	;
	goto L153
L153:
	;
	v625 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v652 = v625 + v223<<(uint(int32(13))%32) + int32(-8192)
	goto L134
L154:
	;
	if v223 < int32(0) {
		goto L158
	} else {
		goto L159
	}
L155:
	;
	v633 = F__emscripten_memcpy_bulkmem(m, v38, v631, int32(8192))
	mBase = m.M
	goto L157
L157:
	;
	goto L154
L158:
	;
	v638 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v638+(v223^int32(-1))<<(uint(int32(2))%32))))
	v652 = v644
	goto L134
L159:
	;
	goto L160
L160:
	;
	v646 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v652 = v646 + v223<<(uint(int32(13))%32) + int32(-8192)
	goto L134
L161:
	;
	if l4 != 0 {
		goto L165
	} else {
		goto L166
	}
L162:
	;
	v655 = F__emscripten_memcpy_bulkmem(m, v652, v653, int32(8192))
	mBase = m.M
	goto L164
L164:
	;
	goto L161
L165:
	;
	v657 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+16)))
	v658 = v86 + v657
	v659 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v658)+6)))
	v661 = v659 & int32(65471)
	*(*uint16)(unsafe.Add(mBase, uint32(v658)+6)) = uint16(v661)
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L5
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v666)+48))
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+118)))
	if v668 != int32(112) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	goto L167
L169:
	;
	v778 = int32(4548900)
	v780 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v780 - int32(1)
	F_UnlockReleaseBuffer(m, v223)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L5
	} else {
		goto L205
	}
L170:
	;
	v672 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v672 <= int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v666)+32))
	if v675 != 0 {
		goto L169
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
	if v677 != 0 {
		goto L169
	} else {
		goto L176
	}
L174:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v666)+40))
	if v676 != 0 {
		goto L169
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L5
	} else {
		goto L177
	}
L177:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v680 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	if l4 != 0 {
		goto L187
	} else {
		goto L188
	}
L179:
	;
	F_XLogRegisterBuffer(m, int32(0), v568, int32(9))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L5
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_XLogRegisterBuffer(m, int32(0), v697, int32(9))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L5
	} else {
		goto L185
	}
L182:
	;
	F_XLogRegisterBuffer(m, int32(1), v223, int32(9))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L5
	} else {
		goto L183
	}
L183:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_XLogRegisterBuffer(m, int32(2), v692, int32(9))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L5
	} else {
		goto L184
	}
L184:
	;
	goto L178
L185:
	;
	F_XLogRegisterBuffer(m, int32(1), v223, int32(9))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L5
	} else {
		goto L186
	}
L186:
	;
	goto L178
L187:
	;
	F_XLogRegisterBuffer(m, int32(3), l4, int32(8))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L5
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	F_XLogRegisterData(m, v16+int32(-48), int32(28))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L5
	} else {
		goto L191
	}
L190:
	;
	goto L189
L191:
	;
	v716 = F_XLogInsert(m, int32(13), int32(48))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L5
	} else {
		goto L192
	}
L192:
	;
	v718 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = base.I64_rotr(v716, v718)
	v721 = base.I32_wrap_i64(v716)
	v724 = base.I32_wrap_i64(int64(base.Ui64(v716) >> (uint(v718) % 64)))
	if v223 < int32(0) {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v742)+4)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v742))) = v724
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v745 == int32(0) {
		goto L197
	} else {
		goto L198
	}
L194:
	;
	v728 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v728+(v223^int32(-1))<<(uint(int32(2))%32))))
	v742 = v734
	goto L193
L195:
	;
	goto L196
L196:
	;
	v736 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v742 = v736 + v223<<(uint(int32(13))%32) + int32(-8192)
	goto L193
L197:
	;
	if v568 < int32(0) {
		goto L201
	} else {
		goto L202
	}
L198:
	;
	goto L199
L199:
	;
	if l4 == int32(0) {
		goto L169
	} else {
		goto L204
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v765)+4)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v765))) = v724
	goto L199
L201:
	;
	v751 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v751+(v568^int32(-1))<<(uint(int32(2))%32))))
	v765 = v757
	goto L200
L202:
	;
	goto L203
L203:
	;
	v759 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v765 = v759 + v568<<(uint(int32(13))%32) + int32(-8192)
	goto L200
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v724
	goto L169
L205:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v786 != 0 {
		v797 = int32(0)
		goto L15
	} else {
		goto L206
	}
L206:
	;
	F_UnlockReleaseBuffer(m, v568)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L5
	} else {
		goto L207
	}
L207:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v797 = base.B2i32(v789 == int32(0))
	goto L15
L208:
	;
	m.G0 = v18 - int32(-64)
	return v797
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v96
	F_errmsg_internal(m, int32(508795), v18)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L5
	} else {
		goto L210
	}
L210:
	;
	F_errfinish(m, int32(523013), int32(650), int32(427949))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L5
	} else {
		goto L211
	}
L211:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_cmp_tslexeme(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v16 = v9 + int32(1)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum_packed(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			v21 = int32(1)
			v22 = v20 & v21
			if v20 == v21 {
				v25 = int32(4)
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
				if v27&int32(254) == int32(2) {
					v36 = v25
				} else {
					v36 = base.B2i32(v27 == int32(18)) << (uint(v25) % 32)
				}
				if v27 == int32(1) {
					v39 = v25
				} else {
					v39 = v36
				}
				v50 = v39
			} else {
				v40 = int32(1)
				if v22 != 0 {
					v50 = int32(base.Ui32(v20)>>(uint(v40)%32)) - v40
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if v22 != 0 {
				v51 = v16
			} else {
				v51 = v9 + int32(4)
			}
			v52 = int32(1)
			v53 = v18 + v52
			v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
			v58 = v56 & v52
			if v58 != 0 {
				v59 = v53
			} else {
				v59 = v18 + int32(4)
			}
			if v56 == int32(1) {
				v62 = int32(4)
				v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				if v64&int32(254) == int32(2) {
					v73 = v62
				} else {
					v73 = base.B2i32(v64 == int32(18)) << (uint(v62) % 32)
				}
				if v64 == int32(1) {
					v76 = v62
				} else {
					v76 = v73
				}
				v87 = v76
			} else {
				v77 = int32(1)
				if v58 != 0 {
					v87 = int32(base.Ui32(v56)>>(uint(v77)%32)) - v77
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
					v87 = int32(base.Ui32(v81)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if v50 == int32(0) {
				v93 = int32(0)
				if v93 < v87 {
					v96 = int32(-1)
				} else {
					v96 = v93
				}
				v113 = v96
			} else {
				if v87 == int32(0) {
					v113 = base.B2i32(int32(0) < v50)
				} else {
					if base.Ui32(v50) < base.Ui32(v87) {
						v102 = v50
					} else {
						v102 = v87
					}
					v103 = F_memcmp(m, v51, v59, v102)
					mBase = m.M
					if v103 != 0 {
						v111 = v103
						v113 = v111
					} else {
						if v50 == v87 {
							v113 = int32(0)
						} else {
							if v50 < v87 {
								v110 = int32(-1)
							} else {
								v110 = int32(1)
							}
							v111 = v110
							v113 = v111
						}
					}
				}
			}
			v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v114 != v9 {
				F_pfree(m, v9)
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
					return int32(0)
				} else {
					v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v118 != v18 {
						F_pfree(m, v18)
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return int32(0)
						} else {
							return v113
						}
					} else {
						return v113
					}
				}
			} else {
				v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v118 != v18 {
					F_pfree(m, v18)
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int32(0)
					} else {
						return v113
					}
				} else {
					return v113
				}
			}
		}
	}
}
func F_gin_compare_prefix_int2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8))))
	if base.Ui32((v16-int32(1))&int32(65535)) < base.Ui32(int32(2)) {
		v23 = v8 + int32(4)
	} else {
		v23 = l0 + int32(20)
	}
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = F_CallerFInfoFunctionCall2(m, v9, v10, v11, v24, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int32(0)
	} else {
		v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8))))
		switch v30 - int32(1) {
		case 0:
			v61 = base.B2i32(v26 <= int32(0))
			m.G0 = v6 + int32(16)
			return v61
		case 1:
			v61 = int32(base.Ui32(v26) >> (uint(int32(31)) % 32))
			m.G0 = v6 + int32(16)
			return v61
		case 2:
			v61 = base.B2i32(v26 != int32(0))
			m.G0 = v6 + int32(16)
			return v61
		case 3:
			v61 = base.B2i32(int32(0) < v26)
			m.G0 = v6 + int32(16)
			return v61
		case 4:
			v39 = int32(0)
			if v26 < v39 {
				v61 = v39
			} else {
				if v26 != 0 {
					v44 = int32(1)
				} else {
					v44 = int32(-1)
				}
				v61 = v44
			}
			m.G0 = v6 + int32(16)
			return v61
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8))))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v49
				F_errmsg_internal(m, int32(504037), v6)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(519738), int32(166), int32(27888))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
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
func F_gin_extract_jsonb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
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
	var v110 int32
	_ = v110
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
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
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v19 = v17 & int32(268435455)
	if v19 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v9 + int32(32)
	return v110
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(0)
	v110 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v26 = F_palloc(m, v19<<(uint(int32(3))%32))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v30 = F_JsonbIteratorInit(m, v12+int32(4))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v30
	v36 = int32(0)
	v37 = v19 << (uint(int32(1)) % 32)
	v38 = v26
	goto L9
L9:
	;
	v47 = F_JsonbIteratorNext(m, v9+int32(28), v9+int32(8), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L17
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v36
	v110 = v38
	goto L3
L11:
	;
	goto L10
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99+v36<<(uint(int32(2))%32)))) = v100
	v36 = v36 + int32(1)
	v37 = v98
	v38 = v99
	goto L9
L13:
	;
	v96 = F_palloc(m, int32(32))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L30
	}
L14:
	;
	v82 = F_make_scalar_key(m, v9+int32(8), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L26
	}
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v68 = F_make_scalar_key(m, v9+int32(8), base.B2i32(v65 == int32(1)))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L22
	}
L16:
	;
	v52 = F_make_scalar_key(m, v9+int32(8), int32(1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	switch v47 {
	case 0:
		goto L11
	case 1:
		goto L16
	case 2:
		goto L14
	case 3:
		goto L15
	default:
		goto L9
	}
L18:
	;
	if v36 < v37 {
		v98 = v37
		v99 = v38
		v100 = v52
		goto L12
	} else {
		goto L19
	}
L19:
	;
	if v37 == int32(0) {
		v93 = v52
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v59 = F_repalloc(m, v38, v37<<(uint(int32(3))%32))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v98 = v37 << (uint(int32(1)) % 32)
	v99 = v59
	v100 = v52
	goto L12
L22:
	;
	if v36 < v37 {
		v98 = v37
		v99 = v38
		v100 = v68
		goto L12
	} else {
		goto L23
	}
L23:
	;
	if v37 == int32(0) {
		v93 = v68
		goto L13
	} else {
		goto L24
	}
L24:
	;
	v75 = F_repalloc(m, v38, v37<<(uint(int32(3))%32))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v98 = v37 << (uint(int32(1)) % 32)
	v99 = v75
	v100 = v68
	goto L12
L26:
	;
	if v36 < v37 {
		v98 = v37
		v99 = v38
		v100 = v82
		goto L12
	} else {
		goto L27
	}
L27:
	;
	if v37 == int32(0) {
		v93 = v82
		goto L13
	} else {
		goto L28
	}
L28:
	;
	v89 = F_repalloc(m, v38, v37<<(uint(int32(3))%32))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v98 = v37 << (uint(int32(1)) % 32)
	v99 = v89
	v100 = v82
	goto L12
L30:
	;
	v98 = int32(8)
	v99 = v96
	v100 = v93
	goto L12
}
func F_gin_extract_query_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_palloc(m, int32(4))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = F_palloc(m, int32(16))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v26
			v29 = F_palloc(m, v26)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v29
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(2119)
				*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v15
				*(*uint16)(unsafe.Add(mBase, uint32(v24))) = uint16(v14)
				v41 = F_palloc(m, int32(4))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v24
					v46 = v14 & int32(65535)
					switch v46 - int32(1) {
					case 0, 1:
						v70 = F_palloc0(m, int32(64))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = v70
							v73 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v73)
							m.G0 = v11 + int32(16)
							return v19
						}
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					case 3, 4:
						v49 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v49)
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v46
							F_errmsg_internal(m, int32(504037), v11)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(519738), int32(97), int32(15868))
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
func F_gin_extract_tsvector_2args(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v2 <= int32(2) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(129616), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(514653), int32(307), int32(164962))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v20 = F_gin_extract_tsvector(m, l0)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			return v20
		}
	}
}
func F_gin_extract_value_inet(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_palloc(m, int32(4))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = F_pg_detoast_datum(m, v5)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v11
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(1)
			return v7
		}
	}
}
func F_gin_leafpage_items(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v246 int64
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v270 int32
	_ = v270
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = F_superuser(m)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if v22 != 0 {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
				if v25 == int32(0) {
					v28 = int32(4554240)
					v29 = F_init_MultiFuncCall(m, l0)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, _consts[10]))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
						*(*int32)(unsafe.Add(mBase, _consts[10])) = v32
						v34 = F_get_page_from_raw(m, v18)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+14)))
							if v36 == int32(0) {
								*(*int32)(unsafe.Add(mBase, _consts[10])) = v31
								v41 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
								v270 = int32(0)
								m.G0 = v15 + int32(48)
								return v270
							} else {
								v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+19)))
								v44 = int32(8)
								v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+16)))
								if (v43<<(uint(v44)%32)-v46)&int32(65535) != v44 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v308 = m.ExcPending
									if v308 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v311 = m.ExcPending
										if v311 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(426868), int32(0))
											mBase = m.M
											v317 = m.ExcPending
											if v317 != 0 {
												return int32(0)
											} else {
												v318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+16)))
												v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+19)))
												v320 = int32(8)
												*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v320
												*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = (v319<<(uint(v320)%32) - v318) & int32(65535)
												F_errdetail(m, int32(680593), v15+int32(16))
												mBase = m.M
												v333 = m.ExcPending
												if v333 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(517608), int32(214), int32(159898))
													mBase = m.M
													v340 = m.ExcPending
													if v340 != 0 {
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
								} else {
									v52 = v34 + v46
									v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+6)))
									if v53 != int32(131) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v344 = m.ExcPending
										if v344 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v347 = m.ExcPending
											if v347 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(426913), int32(0))
												mBase = m.M
												v353 = m.ExcPending
												if v353 != 0 {
													return int32(0)
												} else {
													v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+6)))
													*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(131)
													*(*int32)(unsafe.Add(mBase, uint32(v15))) = v354
													F_errdetail(m, int32(534957), v15)
													mBase = m.M
													v361 = m.ExcPending
													if v361 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(517608), int32(223), int32(159898))
														mBase = m.M
														v368 = m.ExcPending
														if v368 != 0 {
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
									} else {
										v57 = F_palloc(m, int32(12))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
										} else {
											v62 = F_get_call_result_type(m, l0, int32(0), v15+int32(36))
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												if v62 != int32(1) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v372 = m.ExcPending
													if v372 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(384879), int32(0))
														mBase = m.M
														v378 = m.ExcPending
														if v378 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(517608), int32(229), int32(159898))
															mBase = m.M
															v385 = m.ExcPending
															if v385 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v66 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
													*(*int32)(unsafe.Add(mBase, uint32(v57))) = v66
													v68 = int32(32)
													v69 = v34 + v68
													*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v69
													v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+12)))
													*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v69 + v71 - v68
													*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v57
													*(*int32)(unsafe.Add(mBase, _consts[10])) = v31
													v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
													v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
													v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
													if v87 != v88 {
														v90 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v15)+34)) = uint8(v90)
														*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v87
														*(*uint16)(unsafe.Add(mBase, uint32(v15)+32)) = uint16(v90)
														v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+6)))
														*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v96
														v100 = F_ginPostingListDecode(m, v87, v15+int32(28))
														mBase = m.M
														v101 = m.ExcPending
														if v101 != 0 {
															return int32(0)
														} else {
															v102 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
															v105 = F_palloc(m, v102<<(uint(int32(2))%32))
															mBase = m.M
															v106 = m.ExcPending
															if v106 != 0 {
																return int32(0)
															} else {
																v107 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
																if v107 <= int32(0) {
																} else {
																	if base.Ui32(int32(4)) <= base.Ui32(v107) {
																		v116 = v90
																		v121 = int32(0)
																		for {
																			v127 = int32(2)
																			v130 = int32(6)
																			*(*int32)(unsafe.Add(mBase, uint32(v105+v116<<(uint(v127)%32)))) = v100 + v116*v130
																			v135 = v116 | int32(1)
																			*(*int32)(unsafe.Add(mBase, uint32(v105+v135<<(uint(v127)%32)))) = v100 + v135*v130
																			v144 = v116 | v127
																			*(*int32)(unsafe.Add(mBase, uint32(v105+v144<<(uint(v127)%32)))) = v100 + v144*v130
																			v153 = v116 | int32(3)
																			*(*int32)(unsafe.Add(mBase, uint32(v105+v153<<(uint(v127)%32)))) = v100 + v153*v130
																			v161 = int32(4)
																			v162 = v116 + v161
																			v164 = v121 + v161
																			if v164 != v107&int32(2147483644) {
																				v116 = v162
																				v121 = v164
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		v167 = v162
																	} else {
																		v167 = v90
																	}
																	v179 = v107 & int32(3)
																	if v179 == int32(0) {
																	} else {
																		v184 = v167
																		v189 = int32(0)
																		for {
																			*(*int32)(unsafe.Add(mBase, uint32(v105+v184<<(uint(int32(2))%32)))) = v100 + v184*int32(6)
																			v202 = int32(1)
																			v205 = v189 + v202
																			if v205 != v179 {
																				v184 = v184 + v202
																				v189 = v205
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																}
																v220 = F_construct_array_builtin(m, v105, v107, int32(27))
																mBase = m.M
																v221 = m.ExcPending
																if v221 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v220
																	F_pfree(m, v105)
																	mBase = m.M
																	v224 = m.ExcPending
																	if v224 != 0 {
																		return int32(0)
																	} else {
																		F_pfree(m, v100)
																		mBase = m.M
																		v226 = m.ExcPending
																		if v226 != 0 {
																			return int32(0)
																		} else {
																			v227 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
																			v232 = F_heap_form_tuple(m, v227, v15+int32(36), v15+int32(32))
																			mBase = m.M
																			v233 = m.ExcPending
																			if v233 != 0 {
																				return int32(0)
																			} else {
																				v234 = *(*int32)(unsafe.Add(mBase, uint32(v232)+16))
																				v235 = F_HeapTupleHeaderGetDatum(m, v234)
																				mBase = m.M
																				v236 = m.ExcPending
																				if v236 != 0 {
																					return int32(0)
																				} else {
																					v237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+6)))
																					v238 = int32(1)
																					*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v87 + (v237+v238)&int32(131070) + int32(8)
																					v246 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
																					*(*int64)(unsafe.Add(mBase, uint32(v85))) = v246 + int64(1)
																					v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																					*(*int32)(unsafe.Add(mBase, uint32(v250)+20)) = v238
																					v270 = v235
																					m.G0 = v15 + int32(48)
																					return v270
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														F_end_MultiFuncCall(m, l0)
														mBase = m.M
														v254 = m.ExcPending
														if v254 != 0 {
															return int32(0)
														} else {
															v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v255)+20)) = int32(2)
															v258 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v258)
															v270 = int32(0)
															m.G0 = v15 + int32(48)
															return v270
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
					v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
					v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
					if v87 != v88 {
						v90 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v15)+34)) = uint8(v90)
						*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v87
						*(*uint16)(unsafe.Add(mBase, uint32(v15)+32)) = uint16(v90)
						v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+6)))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v96
						v100 = F_ginPostingListDecode(m, v87, v15+int32(28))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
							v105 = F_palloc(m, v102<<(uint(int32(2))%32))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
								if v107 <= int32(0) {
								} else {
									if base.Ui32(int32(4)) <= base.Ui32(v107) {
										v116 = v90
										v121 = int32(0)
										for {
											v127 = int32(2)
											v130 = int32(6)
											*(*int32)(unsafe.Add(mBase, uint32(v105+v116<<(uint(v127)%32)))) = v100 + v116*v130
											v135 = v116 | int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v105+v135<<(uint(v127)%32)))) = v100 + v135*v130
											v144 = v116 | v127
											*(*int32)(unsafe.Add(mBase, uint32(v105+v144<<(uint(v127)%32)))) = v100 + v144*v130
											v153 = v116 | int32(3)
											*(*int32)(unsafe.Add(mBase, uint32(v105+v153<<(uint(v127)%32)))) = v100 + v153*v130
											v161 = int32(4)
											v162 = v116 + v161
											v164 = v121 + v161
											if v164 != v107&int32(2147483644) {
												v116 = v162
												v121 = v164
												continue
											} else {
												break
											}
											break
										}
										v167 = v162
									} else {
										v167 = v90
									}
									v179 = v107 & int32(3)
									if v179 == int32(0) {
									} else {
										v184 = v167
										v189 = int32(0)
										for {
											*(*int32)(unsafe.Add(mBase, uint32(v105+v184<<(uint(int32(2))%32)))) = v100 + v184*int32(6)
											v202 = int32(1)
											v205 = v189 + v202
											if v205 != v179 {
												v184 = v184 + v202
												v189 = v205
												continue
											} else {
												break
											}
											break
										}
									}
								}
								v220 = F_construct_array_builtin(m, v105, v107, int32(27))
								mBase = m.M
								v221 = m.ExcPending
								if v221 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v220
									F_pfree(m, v105)
									mBase = m.M
									v224 = m.ExcPending
									if v224 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v100)
										mBase = m.M
										v226 = m.ExcPending
										if v226 != 0 {
											return int32(0)
										} else {
											v227 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
											v232 = F_heap_form_tuple(m, v227, v15+int32(36), v15+int32(32))
											mBase = m.M
											v233 = m.ExcPending
											if v233 != 0 {
												return int32(0)
											} else {
												v234 = *(*int32)(unsafe.Add(mBase, uint32(v232)+16))
												v235 = F_HeapTupleHeaderGetDatum(m, v234)
												mBase = m.M
												v236 = m.ExcPending
												if v236 != 0 {
													return int32(0)
												} else {
													v237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+6)))
													v238 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v87 + (v237+v238)&int32(131070) + int32(8)
													v246 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
													*(*int64)(unsafe.Add(mBase, uint32(v85))) = v246 + int64(1)
													v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v250)+20)) = v238
													v270 = v235
													m.G0 = v15 + int32(48)
													return v270
												}
											}
										}
									}
								}
							}
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v254 = m.ExcPending
						if v254 != 0 {
							return int32(0)
						} else {
							v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v255)+20)) = int32(2)
							v258 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v258)
							v270 = int32(0)
							m.G0 = v15 + int32(48)
							return v270
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v288 = m.ExcPending
				if v288 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v291 = m.ExcPending
					if v291 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(149499), int32(0))
						mBase = m.M
						v297 = m.ExcPending
						if v297 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(517608), int32(188), int32(159898))
							mBase = m.M
							v304 = m.ExcPending
							if v304 != 0 {
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
func F_gin_numeric_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == int32(0) {
		if v3 != 0 {
			v9 = int32(-1)
		} else {
			v9 = int32(0)
		}
		return v9
	} else {
		if v3 == int32(0) {
			return int32(1)
		} else {
			v17 = F_DirectFunctionCall2Coll(m, int32(1343), int32(0), v4, v3)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				return v17
			}
		}
	}
}
func F_gin_page_opaque_info(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int64
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
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
	v16 = F_superuser(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L66
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L61
	}
L5:
	;
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v18 = F_get_page_from_raw(m, v12)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L57
	}
L9:
	;
	m.G0 = v9 + int32(112)
	return v199
L10:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+14)))
	if v20 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v23 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
	v199 = int32(0)
	goto L9
L12:
	;
	goto L13
L13:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+19)))
	v27 = int32(8)
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+16)))
	if (v26<<(uint(v27)%32)-v29)&int32(65535) != v27 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v35 = int32(0)
	v39 = F_get_call_result_type(m, l0, v35, v9+int32(108))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v39 != int32(1) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v45 = v18 + v29
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+6)))
	if v46&int32(1) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v54 = F_cstring_to_text(m, int32(529729))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v58 = v35
	v59 = v9 + int32(16)
	goto L19
L19:
	;
	if v46&int32(2) != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v54
	v58 = int32(1)
	v59 = v9 + int32(16) | int32(4)
	goto L19
L21:
	;
	v63 = F_cstring_to_text(m, int32(356119))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	v68 = v58
	goto L23
L23:
	;
	if v46&int32(4) != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v63
	v68 = v58 + int32(1)
	goto L23
L25:
	;
	v77 = F_cstring_to_text(m, int32(467190))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v82 = v68
	goto L27
L27:
	;
	if v46&int32(8) != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)|v68<<(uint(int32(2))%32)))) = v77
	v82 = v68 + int32(1)
	goto L27
L29:
	;
	v91 = F_cstring_to_text(m, int32(528228))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	v96 = v82
	goto L31
L31:
	;
	if v46&int32(16) != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)+v82<<(uint(int32(2))%32)))) = v91
	v96 = v82 + int32(1)
	goto L31
L33:
	;
	v105 = F_cstring_to_text(m, int32(81571))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	v110 = v96
	goto L35
L35:
	;
	if v46&int32(32) != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)+v96<<(uint(int32(2))%32)))) = v105
	v110 = v96 + int32(1)
	goto L35
L37:
	;
	v119 = F_cstring_to_text(m, int32(31425))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	v124 = v110
	goto L39
L39:
	;
	if v46&int32(64) != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)+v110<<(uint(int32(2))%32)))) = v119
	v124 = v110 + int32(1)
	goto L39
L41:
	;
	v133 = F_cstring_to_text(m, int32(109070))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	v138 = v124
	goto L43
L43:
	;
	if v46&int32(128) != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)+v124<<(uint(int32(2))%32)))) = v133
	v138 = v124 + int32(1)
	goto L43
L45:
	;
	v147 = F_cstring_to_text(m, int32(470421))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	v152 = v138
	goto L47
L47:
	;
	v154 = v46 & int32(65280)
	if v154 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)+v138<<(uint(int32(2))%32)))) = v147
	v152 = v138 + int32(1)
	goto L47
L49:
	;
	v162 = F_DirectFunctionCall1Coll(m, int32(2911), int32(0), v154)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	v167 = v152
	goto L51
L51:
	;
	v168 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+94)) = uint8(v168)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+92)) = uint16(v168)
	v172 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v45))))
	v173 = F_Int64GetDatum(m, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(16)+v152<<(uint(int32(2))%32)))) = v162
	v167 = v152 + int32(1)
	goto L51
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v173
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+100)) = v176
	v181 = F_construct_array_builtin(m, v9+int32(16), v167, int32(25))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+104)) = v181
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v9)+108))
	v189 = F_heap_form_tuple(m, v184, v9+int32(96), v9+int32(92))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v189)+16))
	v192 = F_HeapTupleHeaderGetDatum(m, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v199 = v192
	goto L9
L57:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(149499), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(517608), int32(112), int32(253537))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errmsg(m, int32(426868), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+16)))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+19)))
	v239 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v239
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = (v238<<(uint(v239)%32) - v237) & int32(65535)
	F_errdetail(m, int32(680593), v9)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(517608), int32(125), int32(253537))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errmsg_internal(m, int32(384879), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(517608), int32(131), int32(253537))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
