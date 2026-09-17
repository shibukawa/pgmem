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
	var v55 int32
	_ = v55
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
	v27 = v10 + (v22+int32(1))&int32(_a_F_GinDataLeafPageGetItems_0)
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
	v55 = v29
	goto L10
L10:
	;
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+14)))
	v63 = v55 + (v58+int32(1))&int32(_a_F_GinDataLeafPageGetItems_0)
	v65 = v63 + int32(8)
	if base.Ui32(v30) <= base.Ui32(v65) {
		v81 = v55
		goto L7
	} else {
		goto L12
	}
L11:
	;
	v81 = v55
	goto L7
L12:
	;
	v67 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v63)+12)))
	v68 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v63)+10)))
	v71 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v63)+8)))
	if base.Ui64(v67|(v68<<(uint(int64(32))%64)|v71<<(uint(int64(48))%64))) <= base.Ui64(v39) {
		v53 = v63
		v55 = v65
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
		goto L20
	} else {
		goto L21
	}
L20:
	;
	base.MemoryCopy(m, v108, v10, v112)
	goto L22
L21:
	;
	goto L22
L22:
	;
	return v108
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v91 int32
	_ = v91
	if l0 < int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_GinInitMetabuffer[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6+(l0^int32(-1))<<(uint(int32(2))%32))))
		v20 = v12
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_GinInitMetabuffer[1]))
		v20 = v14 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	v21 = int32(_a_F_GinInitMetabuffer_0)
	v23 = int32(0)
	if v23|(v20&int32(3)|int32(1)) == v23 {
		v39 = v20 + v21
		v41 = v20 + int32(4)
		if base.Ui32(v41) < base.Ui32(v39) {
			v43 = v39
		} else {
			v43 = v41
		}
		v48 = (v20^int32(-1)+v43)&int32(-4) + int32(4)
		if v48 == int32(0) {
		} else {
			base.MemoryFill(m, v20, int32(0), v48)
		}
	} else {
		base.MemoryFill(m, v20, int32(0), v21)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v20)+10)) = int32(_a_F_GinInitMetabuffer_1)
	v62 = int32(_a_F_GinInitMetabuffer_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+18)) = uint16(v62)
	v68 = int32(_a_F_GinInitMetabuffer_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)) = uint16(v68)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+14)) = uint16(v68)
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	v72 = v20 + v71
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(-1)
	v75 = int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+6)) = uint16(v75)
	v77 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+64)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = int32(2)
	v91 = int32(80)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+12)) = uint16(v91)
	return
}
func F_GinInitPage(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	v2 = l1
	v7 = int32(3)
	if l2&v7|(l0&v7|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(l2))) == int32(0) {
		if l2 == int32(0) {
		} else {
			v21 = l0 + l2
			v23 = l0 + int32(4)
			if base.Ui32(v23) < base.Ui32(v21) {
				v25 = v21
			} else {
				v25 = v23
			}
			v30 = (l0^int32(-1)+v25)&int32(-4) + int32(4)
			if v30 == int32(0) {
			} else {
				base.MemoryFill(m, l0, int32(0), v30)
			}
		}
	} else {
		if l2 == int32(0) {
		} else {
			base.MemoryFill(m, l0, int32(0), l2)
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(_a_F_GinInitPage_0)
	v44 = l2 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v44)
	v50 = l2 - int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v50)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v50)
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v54 = l0 + v53
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+6)) = uint16(v2)
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
			v10 = F_palloc(m, int32(_a_F_ginAllocEntryAccumulator_0))
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
		v10 = F_palloc(m, int32(_a_F_ginAllocEntryAccumulator_0))
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
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(base.B2i32(v14 == int32(_a_F_ginBeginBAScan_0)))
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
					F_errmsg(m, int32(_a_F_ginCombineData_0), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						F_errhint(m, int32(_a_F_ginCombineData_1), int32(0))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ginCombineData_2), int32(45), int32(_a_F_ginCombineData_3))
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
						v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
						*(*uint16)(unsafe.Add(mBase, uint32(v70)+4)) = uint16(v72)
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
						*(*int32)(unsafe.Add(mBase, uint32(v70))) = v74
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
		v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
		*(*uint16)(unsafe.Add(mBase, uint32(v70)+4)) = uint16(v72)
		v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
		*(*int32)(unsafe.Add(mBase, uint32(v70))) = v74
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
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
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
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	v3 = l2
	v5 = l4
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v8) {
		v18 = int32(base.Ui32(v8+int32(_a_F_ginEntryFillRoot_0))>>(uint(int32(2))%32)) & int32(_a_F_ginEntryFillRoot_1)
	} else {
		v18 = int32(0)
	}
	v19 = int32(2)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l3+v18<<(uint(v19)%32))+20))
	v25 = l3 + v22&int32(_a_F_ginEntryFillRoot_2)
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+16)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v26)+6)))
	if v28&v19 == int32(0) {
		v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)))
		v58 = F_palloc(m, v55&int32(_a_F_ginEntryFillRoot_3))
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return
		} else {
			v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)))
			v62 = v60 & int32(_a_F_ginEntryFillRoot_3)
			if v62 != 0 {
				base.MemoryCopy(m, v58, v25, v62)
			} else {
			}
			v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+6)))
			v65 = v64
			v66 = v58
			v68 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v68)
			*(*uint16)(unsafe.Add(mBase, uint32(v66)+2)) = uint16(v3)
			v72 = int32(base.Ui32(v3) >> (uint(int32(16)) % 32))
			*(*uint16)(unsafe.Add(mBase, uint32(v66))) = uint16(v72)
			v78 = F_PageAddItemExtended(m, l1, v66, v65&int32(_a_F_ginEntryFillRoot_3), v68, v68)
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return
			} else {
				if v78 != 0 {
					F_pfree(m, v66)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+12)))
						if base.Ui32(int32(25)) <= base.Ui32(v82) {
							v92 = int32(base.Ui32(v82+int32(_a_F_ginEntryFillRoot_0))>>(uint(int32(2))%32)) & int32(_a_F_ginEntryFillRoot_1)
						} else {
							v92 = int32(0)
						}
						v93 = int32(2)
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l5+v92<<(uint(v93)%32))+20))
						v99 = l5 + v96&int32(_a_F_ginEntryFillRoot_2)
						v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+16)))
						v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v100)+6)))
						if v102&v93 == int32(0) {
							v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
							v132 = F_palloc(m, v129&int32(_a_F_ginEntryFillRoot_3))
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
								return
							} else {
								v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
								v136 = v134 & int32(_a_F_ginEntryFillRoot_3)
								if v136 != 0 {
									base.MemoryCopy(m, v132, v99, v136)
								} else {
								}
								v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+6)))
								v140 = v132
								v141 = v138
								v142 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v142)
								*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
								v146 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
								*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v146)
								v152 = F_PageAddItemExtended(m, l1, v140, v141&int32(_a_F_ginEntryFillRoot_3), v142, v142)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									if v152 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v174 = m.ExcPending
										if v174 != 0 {
											return
										} else {
											F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
											mBase = m.M
											v178 = m.ExcPending
											if v178 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(736), int32(_a_F_ginEntryFillRoot_6))
												mBase = m.M
												v183 = m.ExcPending
												if v183 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										F_pfree(m, v140)
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						} else {
							v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+4)))
							if v107 == int32(_a_F_ginEntryFillRoot_1) {
								v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
								v132 = F_palloc(m, v129&int32(_a_F_ginEntryFillRoot_3))
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
									return
								} else {
									v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
									v136 = v134 & int32(_a_F_ginEntryFillRoot_3)
									if v136 != 0 {
										base.MemoryCopy(m, v132, v99, v136)
									} else {
									}
									v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+6)))
									v140 = v132
									v141 = v138
									v142 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v142)
									*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
									v146 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
									*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v146)
									v152 = F_PageAddItemExtended(m, l1, v140, v141&int32(_a_F_ginEntryFillRoot_3), v142, v142)
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return
									} else {
										if v152 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return
											} else {
												F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(736), int32(_a_F_ginEntryFillRoot_6))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											F_pfree(m, v140)
											mBase = m.M
											v157 = m.ExcPending
											if v157 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							} else {
								v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+2)))
								v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99))))
								v120 = (v110 | v111<<(uint(int32(16))%32)&int32(2147418112) + int32(7)) & int32(-8)
								v121 = F_palloc(m, v120)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return
								} else {
									if v120 != 0 {
										base.MemoryCopy(m, v121, v99, v120)
									} else {
									}
									v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121)+6)))
									v127 = v124&int32(-8192) | v120
									*(*uint16)(unsafe.Add(mBase, uint32(v121)+6)) = uint16(v127)
									v140 = v121
									v141 = v127
									v142 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v142)
									*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
									v146 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
									*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v146)
									v152 = F_PageAddItemExtended(m, l1, v140, v141&int32(_a_F_ginEntryFillRoot_3), v142, v142)
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return
									} else {
										if v152 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return
											} else {
												F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(736), int32(_a_F_ginEntryFillRoot_6))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											F_pfree(m, v140)
											mBase = m.M
											v157 = m.ExcPending
											if v157 != 0 {
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
					v161 = m.ExcPending
					if v161 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
						mBase = m.M
						v165 = m.ExcPending
						if v165 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(731), int32(_a_F_ginEntryFillRoot_6))
							mBase = m.M
							v170 = m.ExcPending
							if v170 != 0 {
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
		v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
		if v33 == int32(_a_F_ginEntryFillRoot_1) {
			v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)))
			v58 = F_palloc(m, v55&int32(_a_F_ginEntryFillRoot_3))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)))
				v62 = v60 & int32(_a_F_ginEntryFillRoot_3)
				if v62 != 0 {
					base.MemoryCopy(m, v58, v25, v62)
				} else {
				}
				v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+6)))
				v65 = v64
				v66 = v58
				v68 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v68)
				*(*uint16)(unsafe.Add(mBase, uint32(v66)+2)) = uint16(v3)
				v72 = int32(base.Ui32(v3) >> (uint(int32(16)) % 32))
				*(*uint16)(unsafe.Add(mBase, uint32(v66))) = uint16(v72)
				v78 = F_PageAddItemExtended(m, l1, v66, v65&int32(_a_F_ginEntryFillRoot_3), v68, v68)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return
				} else {
					if v78 != 0 {
						F_pfree(m, v66)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return
						} else {
							v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+12)))
							if base.Ui32(int32(25)) <= base.Ui32(v82) {
								v92 = int32(base.Ui32(v82+int32(_a_F_ginEntryFillRoot_0))>>(uint(int32(2))%32)) & int32(_a_F_ginEntryFillRoot_1)
							} else {
								v92 = int32(0)
							}
							v93 = int32(2)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(l5+v92<<(uint(v93)%32))+20))
							v99 = l5 + v96&int32(_a_F_ginEntryFillRoot_2)
							v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+16)))
							v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v100)+6)))
							if v102&v93 == int32(0) {
								v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
								v132 = F_palloc(m, v129&int32(_a_F_ginEntryFillRoot_3))
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
									return
								} else {
									v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
									v136 = v134 & int32(_a_F_ginEntryFillRoot_3)
									if v136 != 0 {
										base.MemoryCopy(m, v132, v99, v136)
									} else {
									}
									v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+6)))
									v140 = v132
									v141 = v138
									v142 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v142)
									*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
									v146 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
									*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v146)
									v152 = F_PageAddItemExtended(m, l1, v140, v141&int32(_a_F_ginEntryFillRoot_3), v142, v142)
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return
									} else {
										if v152 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return
											} else {
												F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(736), int32(_a_F_ginEntryFillRoot_6))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											F_pfree(m, v140)
											mBase = m.M
											v157 = m.ExcPending
											if v157 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							} else {
								v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+4)))
								if v107 == int32(_a_F_ginEntryFillRoot_1) {
									v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
									v132 = F_palloc(m, v129&int32(_a_F_ginEntryFillRoot_3))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return
									} else {
										v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
										v136 = v134 & int32(_a_F_ginEntryFillRoot_3)
										if v136 != 0 {
											base.MemoryCopy(m, v132, v99, v136)
										} else {
										}
										v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+6)))
										v140 = v132
										v141 = v138
										v142 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v142)
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
										v146 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
										*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v146)
										v152 = F_PageAddItemExtended(m, l1, v140, v141&int32(_a_F_ginEntryFillRoot_3), v142, v142)
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return
										} else {
											if v152 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v174 = m.ExcPending
												if v174 != 0 {
													return
												} else {
													F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(736), int32(_a_F_ginEntryFillRoot_6))
														mBase = m.M
														v183 = m.ExcPending
														if v183 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												F_pfree(m, v140)
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								} else {
									v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+2)))
									v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99))))
									v120 = (v110 | v111<<(uint(int32(16))%32)&int32(2147418112) + int32(7)) & int32(-8)
									v121 = F_palloc(m, v120)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return
									} else {
										if v120 != 0 {
											base.MemoryCopy(m, v121, v99, v120)
										} else {
										}
										v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121)+6)))
										v127 = v124&int32(-8192) | v120
										*(*uint16)(unsafe.Add(mBase, uint32(v121)+6)) = uint16(v127)
										v140 = v121
										v141 = v127
										v142 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v142)
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
										v146 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
										*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v146)
										v152 = F_PageAddItemExtended(m, l1, v140, v141&int32(_a_F_ginEntryFillRoot_3), v142, v142)
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return
										} else {
											if v152 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v174 = m.ExcPending
												if v174 != 0 {
													return
												} else {
													F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(736), int32(_a_F_ginEntryFillRoot_6))
														mBase = m.M
														v183 = m.ExcPending
														if v183 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												F_pfree(m, v140)
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
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
						v161 = m.ExcPending
						if v161 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
							mBase = m.M
							v165 = m.ExcPending
							if v165 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(731), int32(_a_F_ginEntryFillRoot_6))
								mBase = m.M
								v170 = m.ExcPending
								if v170 != 0 {
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
			v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+2)))
			v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
			v46 = (v36 | v37<<(uint(int32(16))%32)&int32(2147418112) + int32(7)) & int32(-8)
			v47 = F_palloc(m, v46)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				if v46 != 0 {
					base.MemoryCopy(m, v47, v25, v46)
				} else {
				}
				v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+6)))
				v53 = v50&int32(-8192) | v46
				*(*uint16)(unsafe.Add(mBase, uint32(v47)+6)) = uint16(v53)
				v65 = v53
				v66 = v47
				v68 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v68)
				*(*uint16)(unsafe.Add(mBase, uint32(v66)+2)) = uint16(v3)
				v72 = int32(base.Ui32(v3) >> (uint(int32(16)) % 32))
				*(*uint16)(unsafe.Add(mBase, uint32(v66))) = uint16(v72)
				v78 = F_PageAddItemExtended(m, l1, v66, v65&int32(_a_F_ginEntryFillRoot_3), v68, v68)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return
				} else {
					if v78 != 0 {
						F_pfree(m, v66)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return
						} else {
							v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+12)))
							if base.Ui32(int32(25)) <= base.Ui32(v82) {
								v92 = int32(base.Ui32(v82+int32(_a_F_ginEntryFillRoot_0))>>(uint(int32(2))%32)) & int32(_a_F_ginEntryFillRoot_1)
							} else {
								v92 = int32(0)
							}
							v93 = int32(2)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(l5+v92<<(uint(v93)%32))+20))
							v99 = l5 + v96&int32(_a_F_ginEntryFillRoot_2)
							v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+16)))
							v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v100)+6)))
							if v102&v93 == int32(0) {
								v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
								v132 = F_palloc(m, v129&int32(_a_F_ginEntryFillRoot_3))
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
									return
								} else {
									v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
									v136 = v134 & int32(_a_F_ginEntryFillRoot_3)
									if v136 != 0 {
										base.MemoryCopy(m, v132, v99, v136)
									} else {
									}
									v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+6)))
									v140 = v132
									v141 = v138
									v142 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v142)
									*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
									v146 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
									*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v146)
									v152 = F_PageAddItemExtended(m, l1, v140, v141&int32(_a_F_ginEntryFillRoot_3), v142, v142)
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return
									} else {
										if v152 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return
											} else {
												F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(736), int32(_a_F_ginEntryFillRoot_6))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											F_pfree(m, v140)
											mBase = m.M
											v157 = m.ExcPending
											if v157 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							} else {
								v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+4)))
								if v107 == int32(_a_F_ginEntryFillRoot_1) {
									v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
									v132 = F_palloc(m, v129&int32(_a_F_ginEntryFillRoot_3))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return
									} else {
										v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
										v136 = v134 & int32(_a_F_ginEntryFillRoot_3)
										if v136 != 0 {
											base.MemoryCopy(m, v132, v99, v136)
										} else {
										}
										v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+6)))
										v140 = v132
										v141 = v138
										v142 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v142)
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
										v146 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
										*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v146)
										v152 = F_PageAddItemExtended(m, l1, v140, v141&int32(_a_F_ginEntryFillRoot_3), v142, v142)
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return
										} else {
											if v152 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v174 = m.ExcPending
												if v174 != 0 {
													return
												} else {
													F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(736), int32(_a_F_ginEntryFillRoot_6))
														mBase = m.M
														v183 = m.ExcPending
														if v183 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												F_pfree(m, v140)
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								} else {
									v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+2)))
									v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99))))
									v120 = (v110 | v111<<(uint(int32(16))%32)&int32(2147418112) + int32(7)) & int32(-8)
									v121 = F_palloc(m, v120)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return
									} else {
										if v120 != 0 {
											base.MemoryCopy(m, v121, v99, v120)
										} else {
										}
										v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121)+6)))
										v127 = v124&int32(-8192) | v120
										*(*uint16)(unsafe.Add(mBase, uint32(v121)+6)) = uint16(v127)
										v140 = v121
										v141 = v127
										v142 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v142)
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
										v146 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
										*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v146)
										v152 = F_PageAddItemExtended(m, l1, v140, v141&int32(_a_F_ginEntryFillRoot_3), v142, v142)
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return
										} else {
											if v152 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v174 = m.ExcPending
												if v174 != 0 {
													return
												} else {
													F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(736), int32(_a_F_ginEntryFillRoot_6))
														mBase = m.M
														v183 = m.ExcPending
														if v183 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												F_pfree(m, v140)
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
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
						v161 = m.ExcPending
						if v161 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
							mBase = m.M
							v165 = m.ExcPending
							if v165 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(731), int32(_a_F_ginEntryFillRoot_6))
								mBase = m.M
								v170 = m.ExcPending
								if v170 != 0 {
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
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
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
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int64
	_ = v260
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	v2 = l1
	v4 = l3
	v8 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v8)
	v20 = v15 + int32(16)
	base.MemoryFill(m, v20, v8, int32(68))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = int32(43)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = int32(44)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(45)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = int32(46)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = int32(47)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = int32(49)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(50)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+60)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = l2
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+54)) = uint16(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+52)) = uint16(v8)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+36)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = int32(51)
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+69)) = uint8(base.B2i32(l6 != v55))
	v60 = F_ginFindLeafPage(m, v20, v55, v55)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		return
	} else {
		v62 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
		if v62 < int32(0) {
			v66 = *(*int32)(unsafe.Add(mBase, _c_F_ginEntryInsert[0]))
			v72 = *(*int32)(unsafe.Add(mBase, uint32(v66+(v62^int32(-1))<<(uint(int32(2))%32))))
			v80 = v72
		} else {
			v74 = *(*int32)(unsafe.Add(mBase, _c_F_ginEntryInsert[1]))
			v80 = v74 + v62<<(uint(int32(13))%32) + int32(-8192)
		}
		v83 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
		v84 = m.T0[v83].(func(*base.Module, int32, int32) int32)(m, v15+int32(16), v60)
		mBase = m.M
		v85 = m.ExcPending
		if v85 != 0 {
			return
		} else {
			if v84 != 0 {
				v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+8)))
				v90 = *(*int32)(unsafe.Add(mBase, uint32(v80+v86<<(uint(int32(2))%32))+20))
				v93 = v80 + v90&int32(_a_F_ginEntryInsert_0)
				v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+4)))
				if v94 == int32(_a_F_ginEntryInsert_1) {
					v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+2)))
					v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93))))
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
					F_LockBuffer(m, v99, int32(0))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return
					} else {
						F_freeGinBtreeStack(m, v60)
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return
						} else {
							v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							F_ginInsertItemPointers(m, v105, v97|v98<<(uint(int32(16))%32), l4, l5, l6)
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return
							} else {
								m.G0 = v15 + int32(96)
								return
							}
						}
					}
				} else {
					v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v112 = int32(0)
					v113 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
					if v113 < v112 {
						v117 = *(*int32)(unsafe.Add(mBase, _c_F_ginEntryInsert[2]))
						v123 = *(*int32)(unsafe.Add(mBase, uint32(v117+(v113^int32(-1))<<(uint(int32(6))%32))+16))
						v132 = v123
					} else {
						v125 = *(*int32)(unsafe.Add(mBase, _c_F_ginEntryInsert[3]))
						v131 = *(*int32)(unsafe.Add(mBase, uint32(v125+v113<<(uint(int32(6))%32)+int32(-64))+16))
						v132 = v131
					}
					F_CheckForSerializableConflictIn(m, v111, v112, v132)
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						v135 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
						v136 = F_gintuple_get_attrnum(m, l0, v93)
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return
						} else {
							v140 = F_gintuple_get_key(m, l0, v93, v15+int32(95))
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								v144 = F_ginReadTuple(m, v93, v15+int32(88))
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
									return
								} else {
									v146 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
									v149 = F_ginMergeItemPointers(m, l4, l5, v144, v146, v15+int32(84))
									mBase = m.M
									v150 = m.ExcPending
									if v150 != 0 {
										return
									} else {
										v151 = *(*int32)(unsafe.Add(mBase, uint32(v15)+84))
										v154 = F_ginCompressPostingList(m, v149, v151, int32(2712), int32(0))
										mBase = m.M
										v155 = m.ExcPending
										if v155 != 0 {
											return
										} else {
											F_pfree(m, v149)
											mBase = m.M
											v157 = m.ExcPending
											if v157 != 0 {
												return
											} else {
												if v154 != 0 {
													v158 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+95)))
													v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+6)))
													v166 = *(*int32)(unsafe.Add(mBase, uint32(v15)+84))
													v168 = F_GinFormTuple(m, l0, v136, v140, v158, v154, (v159+int32(1))&int32(_a_F_ginEntryInsert_2)+int32(8), v166, int32(0))
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
														return
													} else {
														F_pfree(m, v154)
														mBase = m.M
														v171 = m.ExcPending
														if v171 != 0 {
															return
														} else {
															if v168 != 0 {
																v194 = v168
																F_pfree(m, v144)
																mBase = m.M
																v196 = m.ExcPending
																if v196 != 0 {
																	return
																} else {
																	v197 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v197)
																	v268 = v194
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v268
																	F_ginInsertValue(m, v15+int32(16), v60, v15+int32(8), l6)
																	mBase = m.M
																	v277 = m.ExcPending
																	if v277 != 0 {
																		return
																	} else {
																		F_pfree(m, v268)
																		mBase = m.M
																		v279 = m.ExcPending
																		if v279 != 0 {
																			return
																		} else {
																			m.G0 = v15 + int32(96)
																			return
																		}
																	}
																}
															} else {
																v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																v174 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
																v175 = F_createPostingTree(m, v173, v144, v174, l6, v135)
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
																	return
																} else {
																	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																	F_ginInsertItemPointers(m, v177, v175, l4, l5, l6)
																	mBase = m.M
																	v179 = m.ExcPending
																	if v179 != 0 {
																		return
																	} else {
																		v180 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+95)))
																		v181 = int32(0)
																		v185 = F_GinFormTuple(m, l0, v136, v140, v180, v181, v181, v181, int32(1))
																		mBase = m.M
																		v186 = m.ExcPending
																		if v186 != 0 {
																			return
																		} else {
																			*(*uint16)(unsafe.Add(mBase, uint32(v185)+2)) = uint16(v175)
																			v189 = int32(base.Ui32(v175) >> (uint(int32(16)) % 32))
																			*(*uint16)(unsafe.Add(mBase, uint32(v185))) = uint16(v189)
																			v191 = int32(_a_F_ginEntryInsert_1)
																			*(*uint16)(unsafe.Add(mBase, uint32(v185)+4)) = uint16(v191)
																			v194 = v185
																			F_pfree(m, v144)
																			mBase = m.M
																			v196 = m.ExcPending
																			if v196 != 0 {
																				return
																			} else {
																				v197 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v197)
																				v268 = v194
																				*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v268
																				F_ginInsertValue(m, v15+int32(16), v60, v15+int32(8), l6)
																				mBase = m.M
																				v277 = m.ExcPending
																				if v277 != 0 {
																					return
																				} else {
																					F_pfree(m, v268)
																					mBase = m.M
																					v279 = m.ExcPending
																					if v279 != 0 {
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
													v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v174 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
													v175 = F_createPostingTree(m, v173, v144, v174, l6, v135)
													mBase = m.M
													v176 = m.ExcPending
													if v176 != 0 {
														return
													} else {
														v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														F_ginInsertItemPointers(m, v177, v175, l4, l5, l6)
														mBase = m.M
														v179 = m.ExcPending
														if v179 != 0 {
															return
														} else {
															v180 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+95)))
															v181 = int32(0)
															v185 = F_GinFormTuple(m, l0, v136, v140, v180, v181, v181, v181, int32(1))
															mBase = m.M
															v186 = m.ExcPending
															if v186 != 0 {
																return
															} else {
																*(*uint16)(unsafe.Add(mBase, uint32(v185)+2)) = uint16(v175)
																v189 = int32(base.Ui32(v175) >> (uint(int32(16)) % 32))
																*(*uint16)(unsafe.Add(mBase, uint32(v185))) = uint16(v189)
																v191 = int32(_a_F_ginEntryInsert_1)
																*(*uint16)(unsafe.Add(mBase, uint32(v185)+4)) = uint16(v191)
																v194 = v185
																F_pfree(m, v144)
																mBase = m.M
																v196 = m.ExcPending
																if v196 != 0 {
																	return
																} else {
																	v197 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v197)
																	v268 = v194
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v268
																	F_ginInsertValue(m, v15+int32(16), v60, v15+int32(8), l6)
																	mBase = m.M
																	v277 = m.ExcPending
																	if v277 != 0 {
																		return
																	} else {
																		F_pfree(m, v268)
																		mBase = m.M
																		v279 = m.ExcPending
																		if v279 != 0 {
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
				v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v200 = int32(0)
				v201 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
				if v201 < v200 {
					v205 = *(*int32)(unsafe.Add(mBase, _c_F_ginEntryInsert[2]))
					v211 = *(*int32)(unsafe.Add(mBase, uint32(v205+(v201^int32(-1))<<(uint(int32(6))%32))+16))
					v220 = v211
				} else {
					v213 = *(*int32)(unsafe.Add(mBase, _c_F_ginEntryInsert[3]))
					v219 = *(*int32)(unsafe.Add(mBase, uint32(v213+v201<<(uint(int32(6))%32)+int32(-64))+16))
					v220 = v219
				}
				F_CheckForSerializableConflictIn(m, v199, v200, v220)
				mBase = m.M
				v222 = m.ExcPending
				if v222 != 0 {
					return
				} else {
					v223 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
					v226 = F_ginCompressPostingList(m, l4, l5, int32(2712), int32(0))
					mBase = m.M
					v227 = m.ExcPending
					if v227 != 0 {
						return
					} else {
						if v226 != 0 {
							v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v226)+6)))
							v236 = F_GinFormTuple(m, l0, v2, l2, v4, v226, (v228+int32(1))&int32(_a_F_ginEntryInsert_2)+int32(8), l5, int32(0))
							mBase = m.M
							v237 = m.ExcPending
							if v237 != 0 {
								return
							} else {
								F_pfree(m, v226)
								mBase = m.M
								v239 = m.ExcPending
								if v239 != 0 {
									return
								} else {
									if v236 != 0 {
										v257 = v236
										if l6 == int32(0) {
											v268 = v257
										} else {
											v260 = *(*int64)(unsafe.Add(mBase, uint32(l6)+16))
											*(*int64)(unsafe.Add(mBase, uint32(l6)+16)) = v260 + int64(1)
											v268 = v257
										}
										*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v268
										F_ginInsertValue(m, v15+int32(16), v60, v15+int32(8), l6)
										mBase = m.M
										v277 = m.ExcPending
										if v277 != 0 {
											return
										} else {
											F_pfree(m, v268)
											mBase = m.M
											v279 = m.ExcPending
											if v279 != 0 {
												return
											} else {
												m.G0 = v15 + int32(96)
												return
											}
										}
									} else {
										v241 = int32(0)
										v245 = F_GinFormTuple(m, l0, v2, l2, v4, v241, v241, v241, int32(1))
										mBase = m.M
										v246 = m.ExcPending
										if v246 != 0 {
											return
										} else {
											v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v248 = F_createPostingTree(m, v247, l4, l5, l6, v223)
											mBase = m.M
											v249 = m.ExcPending
											if v249 != 0 {
												return
											} else {
												*(*uint16)(unsafe.Add(mBase, uint32(v245)+2)) = uint16(v248)
												v251 = int32(_a_F_ginEntryInsert_1)
												*(*uint16)(unsafe.Add(mBase, uint32(v245)+4)) = uint16(v251)
												v254 = int32(base.Ui32(v248) >> (uint(int32(16)) % 32))
												*(*uint16)(unsafe.Add(mBase, uint32(v245))) = uint16(v254)
												v257 = v245
												if l6 == int32(0) {
													v268 = v257
												} else {
													v260 = *(*int64)(unsafe.Add(mBase, uint32(l6)+16))
													*(*int64)(unsafe.Add(mBase, uint32(l6)+16)) = v260 + int64(1)
													v268 = v257
												}
												*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v268
												F_ginInsertValue(m, v15+int32(16), v60, v15+int32(8), l6)
												mBase = m.M
												v277 = m.ExcPending
												if v277 != 0 {
													return
												} else {
													F_pfree(m, v268)
													mBase = m.M
													v279 = m.ExcPending
													if v279 != 0 {
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
							v241 = int32(0)
							v245 = F_GinFormTuple(m, l0, v2, l2, v4, v241, v241, v241, int32(1))
							mBase = m.M
							v246 = m.ExcPending
							if v246 != 0 {
								return
							} else {
								v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v248 = F_createPostingTree(m, v247, l4, l5, l6, v223)
								mBase = m.M
								v249 = m.ExcPending
								if v249 != 0 {
									return
								} else {
									*(*uint16)(unsafe.Add(mBase, uint32(v245)+2)) = uint16(v248)
									v251 = int32(_a_F_ginEntryInsert_1)
									*(*uint16)(unsafe.Add(mBase, uint32(v245)+4)) = uint16(v251)
									v254 = int32(base.Ui32(v248) >> (uint(int32(16)) % 32))
									*(*uint16)(unsafe.Add(mBase, uint32(v245))) = uint16(v254)
									v257 = v245
									if l6 == int32(0) {
										v268 = v257
									} else {
										v260 = *(*int64)(unsafe.Add(mBase, uint32(l6)+16))
										*(*int64)(unsafe.Add(mBase, uint32(l6)+16)) = v260 + int64(1)
										v268 = v257
									}
									*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v268
									F_ginInsertValue(m, v15+int32(16), v60, v15+int32(8), l6)
									mBase = m.M
									v277 = m.ExcPending
									if v277 != 0 {
										return
									} else {
										F_pfree(m, v268)
										mBase = m.M
										v279 = m.ExcPending
										if v279 != 0 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFreeScanKeys[0])))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFreeScanKeys[1])))
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
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFreeScanKeys[2])))
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFreeScanKeys[3])))
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
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFreeScanKeys[1])))
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
	*(*int64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFreeScanKeys[3]))) = v41
	*(*int64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFreeScanKeys[0]))) = v41
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
	var v40 int32
	_ = v40
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
			if base.B2i32(v29 != int32(1))|base.B2i32(base.Ui32(v27) < base.Ui32(int32(2))) == int32(0) {
				F_pg_qsort(m, v26, v27, int32(6), int32(34))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					return v26
				}
			} else {
				return v26
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
				if base.B2i32(v29 != int32(1))|base.B2i32(base.Ui32(v27) < base.Ui32(int32(2))) == int32(0) {
					F_pg_qsort(m, v26, v27, int32(6), int32(34))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						return v26
					}
				} else {
					return v26
				}
			}
		}
	}
}
func F_ginInsertCleanup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v511 int32
	_ = v511
	var v515 int64
	_ = v515
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v541 int64
	_ = v541
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int64
	_ = v590
	var v591 int64
	_ = v591
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v611 int32
	_ = v611
	var v619 int64
	_ = v619
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v644 int64
	_ = v644
	var v647 int32
	_ = v647
	var v654 int64
	_ = v654
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v669 int32
	_ = v669
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v731 int32
	_ = v731
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v773 int32
	_ = v773
	var v796 int32
	_ = v796
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v835 int64
	_ = v835
	var v837 int64
	_ = v837
	var v839 int64
	_ = v839
	var v841 int64
	_ = v841
	var v843 int64
	_ = v843
	var v845 int64
	_ = v845
	var v847 int64
	_ = v847
	var v853 int32
	_ = v853
	var v856 int64
	_ = v856
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v872 int32
	_ = v872
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v930 int32
	_ = v930
	var v949 int32
	_ = v949
	var v955 int32
	_ = v955
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v1004 int32
	_ = v1004
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1096 int32
	_ = v1096
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1144 int32
	_ = v1144
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1185 int32
	_ = v1185
	v26 = m.G0
	v28 = v26 - int32(256)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l3 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v28 + int32(256)
	return
L2:
	;
	v75 = F_ReadBuffer(m, v30, int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L16
	}
L3:
	;
	F_LockPage(m, v30, int32(0), int32(7))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v47 = m.G0
	v49 = v47 - int32(16)
	m.G0 = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v30)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v30)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = int32(16973824)
	v56 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v53
	v62 = F_LockAcquire(m, v49, int32(7), v56, int32(1))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L6
	} else {
		goto L14
	}
L6:
	;
	return
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[0]))
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[1]))
	if v36 != int32(-1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v41 = v36
	goto L10
L9:
	;
	v41 = v38
	goto L10
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[2]))
	if v43 == int32(4) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v46 = v41
	goto L13
L12:
	;
	v46 = v38
	goto L13
L13:
	;
	v73 = v46
	goto L2
L14:
	;
	m.G0 = v49 + int32(16)
	if v62 == int32(0) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[3]))
	v73 = v70
	goto L2
L16:
	;
	F_LockBuffer(m, v75, int32(1))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	if v75 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
	if v98 == int32(-1) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[4]))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v83+(v75^int32(-1))<<(uint(int32(2))%32))))
	v97 = v89
	goto L18
L20:
	;
	goto L21
L21:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[5]))
	v97 = v91 + v75<<(uint(int32(13))%32) + int32(-8192)
	goto L18
L22:
	;
	F_UnlockReleaseBuffer(m, v75)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v97)+28))
	v108 = F_ReadBuffer(m, v30, v98)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	F_UnlockPage(m, v30, int32(0), int32(7))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	goto L1
L27:
	;
	F_LockBuffer(m, v108, int32(1))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	if v108 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	F_LockBuffer(m, v75, int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L6
	} else {
		goto L33
	}
L30:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[4]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116+(v108^int32(-1))<<(uint(int32(2))%32))))
	v130 = v122
	goto L29
L31:
	;
	goto L32
L32:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[5]))
	v130 = v124 + v108<<(uint(int32(13))%32) + int32(-8192)
	goto L29
L33:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[6]))
	v140 = F_AllocSetContextCreateInternal(m, v135, int32(_a_F_ginInsertCleanup_0), int32(0), int32(_a_F_ginInsertCleanup_1), int32(_a_F_ginInsertCleanup_2))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	v142 = int32(_a_F_ginInsertCleanup_3)
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[6])) = v140
	v147 = F_palloc(m, int32(512))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v147
	v151 = F_palloc(m, int32(128))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+20)) = int64(549755813888)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v151
	F_ginInitBA(m, v28+int32(28))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = l0
	v177 = v130
	v178 = v108
	v180 = v98
	v185 = int32(0)
	goto L38
L38:
	;
	F_processPendingPage(m, v28+int32(28), v28+int32(12), v177, int32(1))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v208 = base.B2i32(v180 == v107)&(l1^int32(1)) | v185
	v209 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v177)+16)))
	v210 = v177 + v209
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	if v211 != int32(-1) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L6
	} else {
		goto L183
	}
L43:
	;
	F_UnlockReleaseBuffer(m, v178)
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L6
	} else {
		goto L182
	}
L44:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+6)))
	if v214&int32(32) == int32(0) {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v177)+12)))
	F_LockBuffer(m, v178, int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L6
	} else {
		goto L49
	}
L47:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	if base.Ui32(v219) < base.Ui32(v73<<(uint(int32(10))%32)) {
		goto L43
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v221) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v232 = int32(base.Ui32(v221+int32(_a_F_ginInsertCleanup_4)) >> (uint(int32(2)) % 32))
	goto L52
L51:
	;
	v232 = int32(0)
	goto L52
L52:
	;
	v234 = v28 + int32(28)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v234)+16))
	v238 = m.G0
	v239 = int32(16)
	v240 = v238 - v239
	m.G0 = v240
	v243 = v28 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v243)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v243))) = v237
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	*(*uint8)(unsafe.Add(mBase, uint32(v243)+12)) = uint8(base.B2i32(v247 == int32(_a_F_ginInsertCleanup_5)))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+4)) = int32(790)
	m.G0 = v240 + v239
	goto L53
L53:
	;
	v262 = F_ginGetBAEntry(m, v234, v28, v28+int32(4), v28+int32(3), v28+int32(8))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	if v262 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v267 = v262
	goto L58
L56:
	;
	goto L57
L57:
	;
	F_LockBuffer(m, v75, int32(2))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L6
	} else {
		goto L64
	}
L58:
	;
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28))))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v291 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28)+3)))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	F_ginEntryInsert(m, l0, v289, v290, v291, v267, v292, int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
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
	v298 = m.ExcPending
	if v298 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	v307 = F_ginGetBAEntry(m, v28+int32(28), v28, v28+int32(4), v28+int32(3), v28+int32(8))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	if v307 != 0 {
		v267 = v307
		goto L58
	} else {
		goto L63
	}
L63:
	;
	goto L59
L64:
	;
	F_LockBuffer(m, v178, int32(1))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	v340 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v177)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v340) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v464 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v177)+16)))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v177+v464)))
	F_UnlockReleaseBuffer(m, v178)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L6
	} else {
		goto L81
	}
L67:
	;
	v348 = int32(base.Ui32(v340+int32(_a_F_ginInsertCleanup_4)) >> (uint(int32(2)) % 32))
	goto L69
L68:
	;
	v348 = int32(0)
	goto L69
L69:
	;
	v349 = int32(_a_F_ginInsertCleanup_6)
	if v348&v349 == v232&v349 {
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v355 = v28 + int32(28)
	F_ginInitBA(m, v355)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	F_processPendingPage(m, v355, v28+int32(12), v177, (v232+int32(1))&int32(_a_F_ginInsertCleanup_6))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L6
	} else {
		goto L72
	}
L72:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v355)+16))
	v369 = m.G0
	v370 = int32(16)
	v371 = v369 - v370
	m.G0 = v371
	v374 = v28 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v374)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v374))) = v368
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v368)))
	*(*uint8)(unsafe.Add(mBase, uint32(v374)+12)) = uint8(base.B2i32(v378 == int32(_a_F_ginInsertCleanup_5)))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+4)) = int32(790)
	m.G0 = v371 + v370
	goto L73
L73:
	;
	v393 = F_ginGetBAEntry(m, v355, v28, v28+int32(4), v28+int32(3), v28+int32(8))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	if v393 == int32(0) {
		goto L66
	} else {
		goto L75
	}
L75:
	;
	v400 = v393
	goto L76
L76:
	;
	v422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28))))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v424 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28)+3)))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	F_ginEntryInsert(m, l0, v422, v423, v424, v400, v425, int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L6
	} else {
		goto L78
	}
L77:
	;
	goto L66
L78:
	;
	v437 = F_ginGetBAEntry(m, v28+int32(28), v28, v28+int32(4), v28+int32(3), v28+int32(8))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	if v437 != 0 {
		v400 = v437
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	if v75 < int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v482 = v480 + int32(32)
	v484 = v480 + int32(24)
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v480)+24))
	v487 = v485
	goto L86
L83:
	;
	v472 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[4]))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v472+(v75^int32(-1))<<(uint(int32(2))%32))))
	v480 = v474
	goto L82
L84:
	;
	goto L85
L85:
	;
	v476 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[5]))
	v480 = v476 + v75<<(uint(int32(13))%32) + int32(-8192)
	goto L82
L86:
	;
	v511 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+248)) = v511
	v515 = int64(0)
	if v487 == v466 {
		v596 = v466
		v601 = v511
		v611 = v511
		v619 = v515
		goto L88
	} else {
		goto L89
	}
L87:
	;
	F_LockBuffer(m, v75, int32(0))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L6
	} else {
		goto L167
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
	v520 = v487
	v523 = v511
	v541 = v515
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28-int32(-64)+v523<<(uint(int32(2))%32)))) = v520
	v548 = F_ReadBuffer(m, v30, v520)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L6
	} else {
		goto L92
	}
L91:
	;
	v596 = v588
	v601 = v584
	v611 = v589
	v619 = v591
	goto L88
L92:
	;
	v551 = v28 + int32(128)
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	v553 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v551+v552<<(uint(v553)%32)))) = v548
	F_LockBuffer(m, v548, v553)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v551+v560<<(uint(int32(2))%32))))
	if v564 < int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v584 = v560 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+248)) = v584
	v586 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v582)+16)))
	v587 = v582 + v586
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	v589 = base.B2i32(v588 != v466)
	v590 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v587)+4)))
	v591 = v541 + v590
	if int32(15) < v584 {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v568 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[4]))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v568+(v564^int32(-1))<<(uint(int32(2))%32))))
	v582 = v574
	goto L94
L96:
	;
	goto L97
L97:
	;
	v576 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[5]))
	v582 = v576 + v564<<(uint(int32(13))%32) + int32(-8192)
	goto L94
L98:
	;
	v596 = v588
	v601 = v584
	v611 = v589
	v619 = v591
	goto L88
L99:
	;
	goto L100
L100:
	;
	if v588 != v466 {
		v520 = v588
		v523 = v584
		v541 = v591
		goto L90
	} else {
		goto L101
	}
L101:
	;
	goto L91
L102:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+28)) = v620 + v601
	goto L104
L103:
	;
	goto L104
L104:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623)+118)))
	if v624 != int32(112) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v636 = int32(_a_F_ginInsertCleanup_7)
	v638 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[7])) = v638 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v480)+24)) = v596
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	v644 = *(*int64)(unsafe.Add(mBase, uint32(v480)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v480)+40)) = v644 - v619
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v480)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v480)+36)) = v647 - v643
	if v596 == int32(-1) {
		goto L113
	} else {
		goto L114
	}
L106:
	;
	v628 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[8]))
	if v628 <= int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
	if v631 != 0 {
		goto L105
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	F_XLogEnsureRecordSpace(m, v601, int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L6
	} else {
		goto L112
	}
L110:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v30)+40))
	if v632 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v480)+28)) = int32(-1)
	v654 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v482)+8)) = v654
	*(*int64)(unsafe.Add(mBase, uint32(v482))) = v654
	goto L115
L114:
	;
	goto L115
L115:
	;
	v658 = int32(80)
	*(*uint16)(unsafe.Add(mBase, uint32(v480)+12)) = uint16(v658)
	F_MarkBufferDirty(m, v75)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	v662 = int32(0)
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	if v662 < v663 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v669 = v662
	goto L120
L118:
	;
	v731 = v663
	goto L119
L119:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+118)))
	if v751 != int32(112) {
		v930 = v731
		goto L130
	} else {
		goto L131
	}
L120:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(128)+v669<<(uint(int32(2))%32))))
	if v696 < int32(0) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v731 = v723
	goto L119
L122:
	;
	v715 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v714)+16)))
	v717 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v715+v714)+6)) = uint16(v717)
	F_MarkBufferDirty(m, v696)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L6
	} else {
		goto L126
	}
L123:
	;
	v700 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[4]))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v700+(v696^int32(-1))<<(uint(int32(2))%32))))
	v714 = v706
	goto L122
L124:
	;
	goto L125
L125:
	;
	v708 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[5]))
	v714 = v708 + v696<<(uint(int32(13))%32) + int32(-8192)
	goto L122
L126:
	;
	v722 = v669 + int32(1)
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	if v722 < v723 {
		v669 = v722
		goto L120
	} else {
		goto L127
	}
L127:
	;
	goto L121
L128:
	;
	if v611 != 0 {
		v487 = v596
		goto L86
	} else {
		goto L166
	}
L129:
	;
	v1063 = int32(_a_F_ginInsertCleanup_7)
	v1065 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[7])) = v1065 - int32(1)
	goto L128
L130:
	;
	v949 = int32(0)
	if v930 <= v949 {
		goto L129
	} else {
		goto L156
	}
L131:
	;
	v755 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[8]))
	if v755 <= int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
	if v758 != 0 {
		v930 = v731
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
	v761 = m.ExcPending
	if v761 != 0 {
		goto L6
	} else {
		goto L137
	}
L135:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v30)+40))
	if v759 != 0 {
		v930 = v731
		goto L130
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	v762 = int32(0)
	F_XLogRegisterBuffer(m, v762, v75, int32(14))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	if int32(0) < v767 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v773 = v762
	goto L142
L140:
	;
	goto L141
L141:
	;
	v835 = *(*int64)(unsafe.Add(mBase, uint32(v484)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+240)) = v835
	v837 = *(*int64)(unsafe.Add(mBase, uint32(v484)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+232)) = v837
	v839 = *(*int64)(unsafe.Add(mBase, uint32(v484)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+224)) = v839
	v841 = *(*int64)(unsafe.Add(mBase, uint32(v484)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+216)) = v841
	v843 = *(*int64)(unsafe.Add(mBase, uint32(v484)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+208)) = v843
	v845 = *(*int64)(unsafe.Add(mBase, uint32(v484)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+200)) = v845
	v847 = *(*int64)(unsafe.Add(mBase, uint32(v484)))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+192)) = v847
	F_XLogRegisterData(m, v28+int32(192), int32(64))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L6
	} else {
		goto L146
	}
L142:
	;
	v796 = v773 + int32(1)
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(128)+v773<<(uint(int32(2))%32))))
	F_XLogRegisterBuffer(m, v796&int32(255), v804, int32(6))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L6
	} else {
		goto L144
	}
L143:
	;
	goto L141
L144:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	if v796 < v808 {
		v773 = v796
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v856 = F_XLogInsert(m, int32(13), int32(128))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L6
	} else {
		goto L147
	}
L147:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v480))) = base.I64_rotr(v856, int64(32))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	if v861 <= int32(0) {
		goto L129
	} else {
		goto L148
	}
L148:
	;
	v872 = int32(0)
	goto L149
L149:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(128)+v872<<(uint(int32(2))%32))))
	if v899 < int32(0) {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	v930 = v922
	goto L130
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v917)+4)) = base.I32_wrap_i64(v856)
	*(*int32)(unsafe.Add(mBase, uint32(v917))) = base.I32_wrap_i64(int64(base.Ui64(v856) >> (uint(int64(32)) % 64)))
	v921 = v872 + int32(1)
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	if v921 < v922 {
		v872 = v921
		goto L149
	} else {
		goto L155
	}
L152:
	;
	v903 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[4]))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v903+(v899^int32(-1))<<(uint(int32(2))%32))))
	v917 = v909
	goto L151
L153:
	;
	goto L154
L154:
	;
	v911 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[5]))
	v917 = v911 + v899<<(uint(int32(13))%32) + int32(-8192)
	goto L151
L155:
	;
	goto L150
L156:
	;
	v955 = v949
	goto L157
L157:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(128)+v955<<(uint(int32(2))%32))))
	F_UnlockReleaseBuffer(m, v982)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L6
	} else {
		goto L159
	}
L158:
	;
	v989 = int32(0)
	v990 = int32(_a_F_ginInsertCleanup_7)
	v992 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[7])) = v992 - int32(1)
	if base.B2i32(l2 == v989)|base.B2i32(v987 <= v989) != 0 {
		goto L128
	} else {
		goto L161
	}
L159:
	;
	v986 = v955 + int32(1)
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	if v986 < v987 {
		v955 = v986
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	v1004 = v989
	goto L162
L162:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v28-int32(-64)+v1004<<(uint(int32(2))%32))))
	F_RecordFreeIndexPage(m, v30, v1031)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L6
	} else {
		goto L164
	}
L163:
	;
	goto L128
L164:
	;
	v1035 = v1004 + int32(1)
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	if v1035 < v1036 {
		v1004 = v1035
		goto L162
	} else {
		goto L165
	}
L165:
	;
	goto L163
L166:
	;
	goto L87
L167:
	;
	if (base.B2i32(v466 == int32(-1))|v208)&int32(1) == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	F_MemoryContextReset(m, v140)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L6
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	F_UnlockPage(m, v30, int32(0), int32(7))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L6
	} else {
		goto L175
	}
L171:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v1109 = F_palloc(m, v1106<<(uint(int32(2))%32))
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L6
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v1109
	v1112 = F_palloc(m, v1106)
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L6
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v1112
	F_ginInitBA(m, v28+int32(28))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L6
	} else {
		goto L174
	}
L174:
	;
	v1144 = v466
	goto L42
L175:
	;
	F_ReleaseBuffer(m, v75)
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L6
	} else {
		goto L176
	}
L176:
	;
	if l2 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	F_FreeSpaceMapVacuum(m, v30)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L6
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[6])) = v143
	F_MemoryContextDelete(m, v140)
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L6
	} else {
		goto L181
	}
L180:
	;
	goto L179
L181:
	;
	goto L1
L182:
	;
	v1144 = v211
	goto L42
L183:
	;
	v1163 = F_ReadBuffer(m, v30, v1144)
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L6
	} else {
		goto L184
	}
L184:
	;
	F_LockBuffer(m, v1163, int32(1))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L6
	} else {
		goto L185
	}
L185:
	;
	if v1163 < int32(0) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[4]))
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1171+(v1163^int32(-1))<<(uint(int32(2))%32))))
	v1185 = v1177
	goto L188
L187:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[5]))
	v1185 = v1179 + v1163<<(uint(int32(13))%32) + int32(-8192)
	goto L188
L188:
	;
	v177 = v1185
	v178 = v1163
	v180 = v1144
	v185 = v208
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
	var v87 int32
	_ = v87
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
	var v205 int64
	_ = v205
	var v209 int64
	_ = v209
	var v214 int64
	_ = v214
	var v215 int32
	_ = v215
	var v217 int64
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v723 int64
	_ = v723
	var v724 int32
	_ = v724
	var v725 int64
	_ = v725
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v758 int32
	_ = v758
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v772 int32
	_ = v772
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v805 int32
	_ = v805
	var v812 int32
	_ = v812
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
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
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[2]))
	v51 = F_AllocSetContextCreateInternal(m, v46, int32(_a_F_ginPlaceToPage_0), v39, int32(_a_F_ginPlaceToPage_1), int32(_a_F_ginPlaceToPage_2))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[0]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24+(v20^int32(-1))<<(uint(int32(2))%32))))
	v38 = v30
	goto L1
L3:
	;
	goto L4
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[1]))
	v38 = v32 + v20<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	v55 = int32(_a_F_ginPlaceToPage_3)
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[2])) = v51
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
	v87 = int32(1)
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
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[0]))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v71+(l4^int32(-1))<<(uint(int32(2))%32))))
	v85 = v63
	v86 = v77
	goto L7
L12:
	;
	goto L13
L13:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[1]))
	v85 = v63
	v86 = v79 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L7
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L5
	} else {
		goto L194
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[2])) = v56
	F_MemoryContextDelete(m, v51)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L5
	} else {
		goto L193
	}
L16:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v234 = F_GinNewBuffer(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L5
	} else {
		goto L60
	}
L17:
	;
	v98 = int32(_a_F_ginPlaceToPage_4)
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[3])) = v100 + int32(1)
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
		v805 = v87
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
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[4]))
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
	v129 = v127 & int32(_a_F_ginPlaceToPage_5)
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
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[4]))
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
	v227 = int32(_a_F_ginPlaceToPage_4)
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[3])) = v229 - int32(1)
	v805 = v87
	goto L15
L41:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[4]))
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
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+4)) = v222
	goto L40
L50:
	;
	if l4 < int32(0) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	goto L52
L52:
	;
	v214 = F_XLogInsert(m, int32(13), int32(32))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L5
	} else {
		goto L59
	}
L53:
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
		goto L57
	}
L54:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v171+(l4^int32(-1))<<(uint(int32(6))%32))+16))
	v186 = v177
	goto L53
L55:
	;
	goto L56
L56:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v179+l4<<(uint(int32(6))%32)+int32(-64))+16))
	v186 = v185
	goto L53
L57:
	;
	v203 = F_XLogInsert(m, int32(13), int32(32))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	v205 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = base.I64_rotr(v203, v205)
	v209 = int64(base.Ui64(v203) >> (uint(v205) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v86))) = uint32(v209)
	v220 = v86
	v222 = base.I32_wrap_i64(v203)
	goto L49
L59:
	;
	v217 = int64(base.Ui64(v214) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v38))) = uint32(v217)
	v220 = v38
	v222 = base.I32_wrap_i64(v214)
	goto L49
L60:
	;
	if l5 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v38+v249)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v252)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v252)+8))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+40)) = uint16(v85)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v254
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v253
	if l4 != 0 {
		goto L67
	} else {
		goto L68
	}
L62:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v238 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v241 + int32(1)
	goto L61
L64:
	;
	goto L65
L65:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v245 + int32(1)
	goto L61
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v284
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v286 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L67:
	;
	if l4 < int32(0) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	goto L69
L69:
	;
	v281 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v281
	v284 = v281
	goto L66
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v276
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+16)))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v86+v278)))
	v284 = v280
	goto L66
L71:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v261+(l4^int32(-1))<<(uint(int32(6))%32))+16))
	v276 = v267
	goto L70
L72:
	;
	goto L73
L73:
	;
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v269+l4<<(uint(int32(6))%32)+int32(-64))+16))
	v276 = v275
	goto L70
L74:
	;
	v584 = int32(_a_F_ginPlaceToPage_4)
	v586 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[3])) = v586 + int32(1)
	F_MarkBufferDirty(m, v234)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L5
	} else {
		goto L133
	}
L75:
	;
	v579 = v575
	v583 = v577
	goto L74
L76:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v533 < int32(0) {
		goto L125
	} else {
		goto L126
	}
L77:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v290 = F_GinNewBuffer(m, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L5
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v251
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	v468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v467)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v467+v468))) = v251
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v472 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v471)+16)))
	v473 = v471 + v472
	v474 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v473)+6)))
	v476 = v474 | int32(64)
	*(*uint16)(unsafe.Add(mBase, uint32(v473)+6)) = uint16(v476)
	if v234 < int32(0) {
		goto L116
	} else {
		goto L117
	}
L80:
	;
	if l5 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v305 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v305
	v308 = v85 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+40)) = uint16(v308)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	v311 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v310)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v310+v311))) = v305
	if v234 < int32(0) {
		goto L87
	} else {
		goto L88
	}
L82:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v294 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v297 + int32(1)
	goto L81
L84:
	;
	goto L85
L85:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v301 + int32(1)
	goto L81
L86:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v335 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v334)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v334+v335))) = v333
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	v339 = F_PageGetTempPage(m, v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L5
	} else {
		goto L90
	}
L87:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v318+(v234^int32(-1))<<(uint(int32(6))%32))+16))
	v333 = v324
	goto L86
L88:
	;
	goto L89
L89:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v326+v234<<(uint(int32(6))%32)+int32(-64))+16))
	v333 = v332
	goto L86
L90:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v341)+16)))
	v344 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v341+v342)+6)))
	v346 = v344 & int32(_a_F_ginPlaceToPage_6)
	F_PageInit(m, v339, int32(_a_F_ginPlaceToPage_1), int32(8))
	mBase = m.M
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v339)+16)))
	v351 = v339 + v350
	*(*int32)(unsafe.Add(mBase, uint32(v351))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v351)+6)) = uint16(v346)
	goto L91
L91:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v290 < int32(0) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	if v234 < int32(0) {
		goto L97
	} else {
		goto L98
	}
L93:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v359+(v290^int32(-1))<<(uint(int32(6))%32))+16))
	v374 = v365
	goto L92
L94:
	;
	goto L95
L95:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v367+v290<<(uint(int32(6))%32)+int32(-64))+16))
	v374 = v373
	goto L92
L96:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	m.T0[v355].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l0, v339, v374, v375, v394, v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L5
	} else {
		goto L100
	}
L97:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v379+(v234^int32(-1))<<(uint(int32(6))%32))+16))
	v394 = v385
	goto L96
L98:
	;
	goto L99
L99:
	;
	v387 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v387+v234<<(uint(int32(6))%32)+int32(-64))+16))
	v394 = v393
	goto L96
L100:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v398 < int32(0) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v417 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v416)+16)))
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417+v416)+6)))
	if v419&int32(2) == int32(0) {
		v575 = v290
		v577 = v339
		goto L75
	} else {
		goto L105
	}
L102:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[0]))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v402+(v398^int32(-1))<<(uint(int32(2))%32))))
	v416 = v408
	goto L101
L103:
	;
	goto L104
L104:
	;
	v410 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[1]))
	v416 = v410 + v398<<(uint(int32(13))%32) + int32(-8192)
	goto L101
L105:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v398 < int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	if v290 < int32(0) {
		goto L111
	} else {
		goto L112
	}
L107:
	;
	v428 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v428+(v398^int32(-1))<<(uint(int32(6))%32))+16))
	v443 = v434
	goto L106
L108:
	;
	goto L109
L109:
	;
	v436 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v436+v398<<(uint(int32(6))%32)+int32(-64))+16))
	v443 = v442
	goto L106
L110:
	;
	F_PredicateLockPageSplit(m, v424, v443, v462)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L5
	} else {
		goto L114
	}
L111:
	;
	v447 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v447+(v290^int32(-1))<<(uint(int32(6))%32))+16))
	v462 = v453
	goto L110
L112:
	;
	goto L113
L113:
	;
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v455+v290<<(uint(int32(6))%32)+int32(-64))+16))
	v462 = v461
	goto L110
L114:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v530 = v290
	v532 = v339
	v533 = v465
	goto L76
L115:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v498 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v497)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v497+v498))) = v496
	v501 = int32(0)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v504 < v501 {
		goto L120
	} else {
		goto L121
	}
L116:
	;
	v481 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v481+(v234^int32(-1))<<(uint(int32(6))%32))+16))
	v496 = v487
	goto L115
L117:
	;
	goto L118
L118:
	;
	v489 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v489+v234<<(uint(int32(6))%32)+int32(-64))+16))
	v496 = v495
	goto L115
L119:
	;
	v523 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v522)+16)))
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523+v522)+6)))
	if v525&int32(2) == int32(0) {
		v579 = v501
		v583 = v501
		goto L74
	} else {
		goto L123
	}
L120:
	;
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[0]))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v508+(v504^int32(-1))<<(uint(int32(2))%32))))
	v522 = v514
	goto L119
L121:
	;
	goto L122
L122:
	;
	v516 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[1]))
	v522 = v516 + v504<<(uint(int32(13))%32) + int32(-8192)
	goto L119
L123:
	;
	v530 = v501
	v532 = v501
	v533 = v504
	goto L76
L124:
	;
	if v234 < int32(0) {
		goto L129
	} else {
		goto L130
	}
L125:
	;
	v538 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v538+(v533^int32(-1))<<(uint(int32(6))%32))+16))
	v553 = v544
	goto L124
L126:
	;
	goto L127
L127:
	;
	v546 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v546+v533<<(uint(int32(6))%32)+int32(-64))+16))
	v553 = v552
	goto L124
L128:
	;
	F_PredicateLockPageSplit(m, v534, v553, v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L5
	} else {
		goto L132
	}
L129:
	;
	v557 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v557+(v234^int32(-1))<<(uint(int32(6))%32))+16))
	v572 = v563
	goto L128
L130:
	;
	goto L131
L131:
	;
	v565 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v565+v234<<(uint(int32(6))%32)+int32(-64))+16))
	v572 = v571
	goto L128
L132:
	;
	v575 = v530
	v577 = v532
	goto L75
L133:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_MarkBufferDirty(m, v592)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L5
	} else {
		goto L134
	}
L134:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v595 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	base.MemoryCopy(m, v660, v661, int32(_a_F_ginPlaceToPage_1))
	if l4 != 0 {
		goto L150
	} else {
		goto L151
	}
L136:
	;
	F_MarkBufferDirty(m, v579)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L5
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	base.MemoryCopy(m, v38, v640, int32(_a_F_ginPlaceToPage_1))
	if v234 < int32(0) {
		goto L147
	} else {
		goto L148
	}
L139:
	;
	base.MemoryCopy(m, v38, v583, int32(_a_F_ginPlaceToPage_1))
	if v579 < int32(0) {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	base.MemoryCopy(m, v619, v620, int32(_a_F_ginPlaceToPage_1))
	if v234 < int32(0) {
		goto L144
	} else {
		goto L145
	}
L141:
	;
	v605 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[0]))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v605+(v579^int32(-1))<<(uint(int32(2))%32))))
	v619 = v611
	goto L140
L142:
	;
	goto L143
L143:
	;
	v613 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[1]))
	v619 = v613 + v579<<(uint(int32(13))%32) + int32(-8192)
	goto L140
L144:
	;
	v626 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[0]))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v626+(v234^int32(-1))<<(uint(int32(2))%32))))
	v660 = v632
	goto L135
L145:
	;
	goto L146
L146:
	;
	v634 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[1]))
	v660 = v634 + v234<<(uint(int32(13))%32) + int32(-8192)
	goto L135
L147:
	;
	v646 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[0]))
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v646+(v234^int32(-1))<<(uint(int32(2))%32))))
	v660 = v652
	goto L135
L148:
	;
	goto L149
L149:
	;
	v654 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[1]))
	v660 = v654 + v234<<(uint(int32(13))%32) + int32(-8192)
	goto L135
L150:
	;
	v664 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+16)))
	v665 = v86 + v664
	v666 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v665)+6)))
	v668 = v666 & int32(_a_F_ginPlaceToPage_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v665)+6)) = uint16(v668)
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L5
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v673)+48))
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674)+118)))
	if v675 != int32(112) {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	goto L152
L154:
	;
	v785 = int32(_a_F_ginPlaceToPage_4)
	v787 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[3])) = v787 - int32(1)
	F_UnlockReleaseBuffer(m, v234)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L5
	} else {
		goto L190
	}
L155:
	;
	v679 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[4]))
	if v679 <= int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v673)+32))
	if v682 != 0 {
		goto L154
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
	if v684 != 0 {
		goto L154
	} else {
		goto L161
	}
L159:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v673)+40))
	if v683 != 0 {
		goto L154
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L5
	} else {
		goto L162
	}
L162:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v687 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	if l4 != 0 {
		goto L172
	} else {
		goto L173
	}
L164:
	;
	F_XLogRegisterBuffer(m, int32(0), v579, int32(9))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L5
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_XLogRegisterBuffer(m, int32(0), v704, int32(9))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L5
	} else {
		goto L170
	}
L167:
	;
	F_XLogRegisterBuffer(m, int32(1), v234, int32(9))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L5
	} else {
		goto L168
	}
L168:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_XLogRegisterBuffer(m, int32(2), v699, int32(9))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L5
	} else {
		goto L169
	}
L169:
	;
	goto L163
L170:
	;
	F_XLogRegisterBuffer(m, int32(1), v234, int32(9))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L5
	} else {
		goto L171
	}
L171:
	;
	goto L163
L172:
	;
	F_XLogRegisterBuffer(m, int32(3), l4, int32(8))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L5
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	F_XLogRegisterData(m, v16+int32(-48), int32(28))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L5
	} else {
		goto L176
	}
L175:
	;
	goto L174
L176:
	;
	v723 = F_XLogInsert(m, int32(13), int32(48))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L5
	} else {
		goto L177
	}
L177:
	;
	v725 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = base.I64_rotr(v723, v725)
	v728 = base.I32_wrap_i64(v723)
	v731 = base.I32_wrap_i64(int64(base.Ui64(v723) >> (uint(v725) % 64)))
	if v234 < int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v749)+4)) = v728
	*(*int32)(unsafe.Add(mBase, uint32(v749))) = v731
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v752 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L179:
	;
	v735 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[0]))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v735+(v234^int32(-1))<<(uint(int32(2))%32))))
	v749 = v741
	goto L178
L180:
	;
	goto L181
L181:
	;
	v743 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[1]))
	v749 = v743 + v234<<(uint(int32(13))%32) + int32(-8192)
	goto L178
L182:
	;
	if v579 < int32(0) {
		goto L186
	} else {
		goto L187
	}
L183:
	;
	goto L184
L184:
	;
	if l4 == int32(0) {
		goto L154
	} else {
		goto L189
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v772)+4)) = v728
	*(*int32)(unsafe.Add(mBase, uint32(v772))) = v731
	goto L184
L186:
	;
	v758 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[0]))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v758+(v579^int32(-1))<<(uint(int32(2))%32))))
	v772 = v764
	goto L185
L187:
	;
	goto L188
L188:
	;
	v766 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[1]))
	v772 = v766 + v579<<(uint(int32(13))%32) + int32(-8192)
	goto L185
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v728
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v731
	goto L154
L190:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v793 != 0 {
		v805 = int32(0)
		goto L15
	} else {
		goto L191
	}
L191:
	;
	F_UnlockReleaseBuffer(m, v579)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L5
	} else {
		goto L192
	}
L192:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v805 = base.B2i32(v796 == int32(0))
	goto L15
L193:
	;
	m.G0 = v18 - int32(-64)
	return v805
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v96
	F_errmsg_internal(m, int32(_a_F_ginPlaceToPage_7), v18)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L5
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(_a_F_ginPlaceToPage_8), int32(650), int32(_a_F_ginPlaceToPage_9))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L5
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_cmp_tslexeme(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v15 = v8 + int32(1)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			v21 = v19 & int32(1)
			if v21 != 0 {
				v22 = v15
			} else {
				v22 = v8 + int32(4)
			}
			if v19 == int32(1) {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v28 == int32(18) {
					v31 = int32(16)
				} else {
					v31 = int32(0)
				}
				if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v38 = int32(4)
				} else {
					v38 = v31
				}
				v49 = v38
			} else {
				v39 = int32(1)
				if v21 != 0 {
					v49 = int32(base.Ui32(v19)>>(uint(v39)%32)) - v39
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v50 = int32(1)
			v51 = v17 + v50
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			v56 = v54 & v50
			if v56 != 0 {
				v57 = v51
			} else {
				v57 = v17 + int32(4)
			}
			if v54 == int32(1) {
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
				if v63 == int32(18) {
					v66 = int32(16)
				} else {
					v66 = int32(0)
				}
				if base.Ui32((v63-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v73 = int32(4)
				} else {
					v73 = v66
				}
				v84 = v73
			} else {
				v74 = int32(1)
				if v56 != 0 {
					v84 = int32(base.Ui32(v54)>>(uint(v74)%32)) - v74
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v84 = int32(base.Ui32(v78)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if v49 == int32(0) {
				v90 = int32(0)
				if v90 < v84 {
					v93 = int32(-1)
				} else {
					v93 = v90
				}
				v110 = v93
			} else {
				if v84 == int32(0) {
					v110 = base.B2i32(int32(0) < v49)
				} else {
					if base.Ui32(v49) < base.Ui32(v84) {
						v99 = v49
					} else {
						v99 = v84
					}
					v100 = F_memcmp(m, v22, v57, v99)
					mBase = m.M
					if v100 != 0 {
						v108 = v100
						v110 = v108
					} else {
						if v49 == v84 {
							v110 = int32(0)
						} else {
							if v49 < v84 {
								v107 = int32(-1)
							} else {
								v107 = int32(1)
							}
							v108 = v107
							v110 = v108
						}
					}
				}
			}
			v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v111 != v8 {
				F_pfree(m, v8)
				mBase = m.M
				v114 = m.ExcPending
				if v114 != 0 {
					return int32(0)
				} else {
					v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v115 != v17 {
						F_pfree(m, v17)
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return int32(0)
						} else {
							return v110
						}
					} else {
						return v110
					}
				}
			} else {
				v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v115 != v17 {
					F_pfree(m, v17)
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						return v110
					}
				} else {
					return v110
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
	if base.Ui32((v16-int32(1))&int32(_a_F_gin_compare_prefix_int2_0)) < base.Ui32(int32(2)) {
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
				F_errmsg_internal(m, int32(_a_F_gin_compare_prefix_int2_1), v6)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_gin_compare_prefix_int2_2), int32(166), int32(_a_F_gin_compare_prefix_int2_3))
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13907(m, l0, int32(64), int32(2103))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
			F_errmsg_internal(m, int32(_a_F_gin_extract_tsvector_2args_0), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_gin_extract_tsvector_2args_1), int32(307), int32(_a_F_gin_extract_tsvector_2args_2))
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v91 int32
	_ = v91
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
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v250 int64
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = F_superuser(m)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			if v23 != 0 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
				if v26 == int32(0) {
					v29 = F_init_MultiFuncCall(m, l0)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = int32(_a_F_gin_leafpage_items_0)
						v32 = *(*int32)(unsafe.Add(mBase, _c_F_gin_leafpage_items[0]))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
						*(*int32)(unsafe.Add(mBase, _c_F_gin_leafpage_items[0])) = v34
						v36 = F_get_page_from_raw(m, v19)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+14)))
							if v38 == int32(0) {
								*(*int32)(unsafe.Add(mBase, _c_F_gin_leafpage_items[0])) = v32
								v269 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v269)
								v273 = int32(0)
								m.G0 = v16 + int32(48)
								return v273
							} else {
								v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+19)))
								v44 = int32(8)
								v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+16)))
								if (v43<<(uint(v44)%32)-v46)&int32(_a_F_gin_leafpage_items_1) != v44 {
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
											F_errmsg(m, int32(_a_F_gin_leafpage_items_2), int32(0))
											mBase = m.M
											v315 = m.ExcPending
											if v315 != 0 {
												return int32(0)
											} else {
												v316 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+16)))
												v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+19)))
												v318 = int32(8)
												*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v318
												*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = (v317<<(uint(v318)%32) - v316) & int32(_a_F_gin_leafpage_items_1)
												F_errdetail(m, int32(_a_F_gin_leafpage_items_3), v16+int32(16))
												mBase = m.M
												v330 = m.ExcPending
												if v330 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_gin_leafpage_items_4), int32(214), int32(_a_F_gin_leafpage_items_5))
													mBase = m.M
													v335 = m.ExcPending
													if v335 != 0 {
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
									v52 = v36 + v46
									v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+6)))
									if v53 != int32(131) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v339 = m.ExcPending
										if v339 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v342 = m.ExcPending
											if v342 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_gin_leafpage_items_6), int32(0))
												mBase = m.M
												v346 = m.ExcPending
												if v346 != 0 {
													return int32(0)
												} else {
													v347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+6)))
													*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(131)
													*(*int32)(unsafe.Add(mBase, uint32(v16))) = v347
													F_errdetail(m, int32(_a_F_gin_leafpage_items_7), v16)
													mBase = m.M
													v353 = m.ExcPending
													if v353 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_gin_leafpage_items_4), int32(223), int32(_a_F_gin_leafpage_items_5))
														mBase = m.M
														v358 = m.ExcPending
														if v358 != 0 {
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
											v62 = F_get_call_result_type(m, l0, int32(0), v16+int32(36))
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												if v62 != int32(1) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v362 = m.ExcPending
													if v362 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_gin_leafpage_items_8), int32(0))
														mBase = m.M
														v366 = m.ExcPending
														if v366 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_gin_leafpage_items_4), int32(229), int32(_a_F_gin_leafpage_items_5))
															mBase = m.M
															v371 = m.ExcPending
															if v371 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v66 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
													*(*int32)(unsafe.Add(mBase, uint32(v57))) = v66
													v68 = int32(32)
													v69 = v36 + v68
													*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v69
													v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+12)))
													*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v69 + v71 - v68
													*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v57
													*(*int32)(unsafe.Add(mBase, _c_F_gin_leafpage_items[0])) = v32
													v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
													v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
													v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
													if v87 != v88 {
														*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v87
														v91 = int32(0)
														*(*uint16)(unsafe.Add(mBase, uint32(v16)+32)) = uint16(v91)
														*(*uint8)(unsafe.Add(mBase, uint32(v16)+34)) = uint8(v91)
														v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+6)))
														*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v96
														v100 = F_ginPostingListDecode(m, v87, v16+int32(28))
														mBase = m.M
														v101 = m.ExcPending
														if v101 != 0 {
															return int32(0)
														} else {
															v102 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
															v105 = F_palloc(m, v102<<(uint(int32(2))%32))
															mBase = m.M
															v106 = m.ExcPending
															if v106 != 0 {
																return int32(0)
															} else {
																v107 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
																if v107 <= int32(0) {
																} else {
																	v111 = v107 & int32(3)
																	if base.Ui32(int32(4)) <= base.Ui32(v107) {
																		v118 = v91
																		v123 = int32(0)
																		for {
																			v130 = int32(2)
																			v133 = int32(6)
																			*(*int32)(unsafe.Add(mBase, uint32(v105+v118<<(uint(v130)%32)))) = v100 + v118*v133
																			v138 = v118 | int32(1)
																			*(*int32)(unsafe.Add(mBase, uint32(v105+v138<<(uint(v130)%32)))) = v100 + v138*v133
																			v147 = v118 | v130
																			*(*int32)(unsafe.Add(mBase, uint32(v105+v147<<(uint(v130)%32)))) = v100 + v147*v133
																			v156 = v118 | int32(3)
																			*(*int32)(unsafe.Add(mBase, uint32(v105+v156<<(uint(v130)%32)))) = v100 + v156*v133
																			v164 = int32(4)
																			v165 = v118 + v164
																			v167 = v123 + v164
																			if v167 != v107&int32(2147483644) {
																				v118 = v165
																				v123 = v167
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v111 == int32(0) {
																		} else {
																			v172 = v165
																			v186 = v172
																			v191 = int32(0)
																			for {
																				*(*int32)(unsafe.Add(mBase, uint32(v105+v186<<(uint(int32(2))%32)))) = v100 + v186*int32(6)
																				v205 = int32(1)
																				v208 = v191 + v205
																				if v208 != v111 {
																					v186 = v186 + v205
																					v191 = v208
																					continue
																				} else {
																					break
																				}
																				break
																			}
																		}
																	} else {
																		v172 = v91
																		v186 = v172
																		v191 = int32(0)
																		for {
																			*(*int32)(unsafe.Add(mBase, uint32(v105+v186<<(uint(int32(2))%32)))) = v100 + v186*int32(6)
																			v205 = int32(1)
																			v208 = v191 + v205
																			if v208 != v111 {
																				v186 = v186 + v205
																				v191 = v208
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																}
																v224 = F_construct_array_builtin(m, v105, v107, int32(27))
																mBase = m.M
																v225 = m.ExcPending
																if v225 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v224
																	F_pfree(m, v105)
																	mBase = m.M
																	v228 = m.ExcPending
																	if v228 != 0 {
																		return int32(0)
																	} else {
																		F_pfree(m, v100)
																		mBase = m.M
																		v230 = m.ExcPending
																		if v230 != 0 {
																			return int32(0)
																		} else {
																			v231 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
																			v236 = F_heap_form_tuple(m, v231, v16+int32(36), v16+int32(32))
																			mBase = m.M
																			v237 = m.ExcPending
																			if v237 != 0 {
																				return int32(0)
																			} else {
																				v238 = *(*int32)(unsafe.Add(mBase, uint32(v236)+16))
																				v239 = F_HeapTupleHeaderGetDatum(m, v238)
																				mBase = m.M
																				v240 = m.ExcPending
																				if v240 != 0 {
																					return int32(0)
																				} else {
																					v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+6)))
																					v242 = int32(1)
																					*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v87 + (v241+v242)&int32(_a_F_gin_leafpage_items_9) + int32(8)
																					v250 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
																					*(*int64)(unsafe.Add(mBase, uint32(v85))) = v250 + int64(1)
																					v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																					*(*int32)(unsafe.Add(mBase, uint32(v254)+20)) = v242
																					v273 = v239
																					m.G0 = v16 + int32(48)
																					return v273
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
														v258 = m.ExcPending
														if v258 != 0 {
															return int32(0)
														} else {
															v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v259)+20)) = int32(2)
															v269 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v269)
															v273 = int32(0)
															m.G0 = v16 + int32(48)
															return v273
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
						*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v87
						v91 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v16)+32)) = uint16(v91)
						*(*uint8)(unsafe.Add(mBase, uint32(v16)+34)) = uint8(v91)
						v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+6)))
						*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v96
						v100 = F_ginPostingListDecode(m, v87, v16+int32(28))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
							v105 = F_palloc(m, v102<<(uint(int32(2))%32))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
								if v107 <= int32(0) {
								} else {
									v111 = v107 & int32(3)
									if base.Ui32(int32(4)) <= base.Ui32(v107) {
										v118 = v91
										v123 = int32(0)
										for {
											v130 = int32(2)
											v133 = int32(6)
											*(*int32)(unsafe.Add(mBase, uint32(v105+v118<<(uint(v130)%32)))) = v100 + v118*v133
											v138 = v118 | int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v105+v138<<(uint(v130)%32)))) = v100 + v138*v133
											v147 = v118 | v130
											*(*int32)(unsafe.Add(mBase, uint32(v105+v147<<(uint(v130)%32)))) = v100 + v147*v133
											v156 = v118 | int32(3)
											*(*int32)(unsafe.Add(mBase, uint32(v105+v156<<(uint(v130)%32)))) = v100 + v156*v133
											v164 = int32(4)
											v165 = v118 + v164
											v167 = v123 + v164
											if v167 != v107&int32(2147483644) {
												v118 = v165
												v123 = v167
												continue
											} else {
												break
											}
											break
										}
										if v111 == int32(0) {
										} else {
											v172 = v165
											v186 = v172
											v191 = int32(0)
											for {
												*(*int32)(unsafe.Add(mBase, uint32(v105+v186<<(uint(int32(2))%32)))) = v100 + v186*int32(6)
												v205 = int32(1)
												v208 = v191 + v205
												if v208 != v111 {
													v186 = v186 + v205
													v191 = v208
													continue
												} else {
													break
												}
												break
											}
										}
									} else {
										v172 = v91
										v186 = v172
										v191 = int32(0)
										for {
											*(*int32)(unsafe.Add(mBase, uint32(v105+v186<<(uint(int32(2))%32)))) = v100 + v186*int32(6)
											v205 = int32(1)
											v208 = v191 + v205
											if v208 != v111 {
												v186 = v186 + v205
												v191 = v208
												continue
											} else {
												break
											}
											break
										}
									}
								}
								v224 = F_construct_array_builtin(m, v105, v107, int32(27))
								mBase = m.M
								v225 = m.ExcPending
								if v225 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v224
									F_pfree(m, v105)
									mBase = m.M
									v228 = m.ExcPending
									if v228 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v100)
										mBase = m.M
										v230 = m.ExcPending
										if v230 != 0 {
											return int32(0)
										} else {
											v231 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
											v236 = F_heap_form_tuple(m, v231, v16+int32(36), v16+int32(32))
											mBase = m.M
											v237 = m.ExcPending
											if v237 != 0 {
												return int32(0)
											} else {
												v238 = *(*int32)(unsafe.Add(mBase, uint32(v236)+16))
												v239 = F_HeapTupleHeaderGetDatum(m, v238)
												mBase = m.M
												v240 = m.ExcPending
												if v240 != 0 {
													return int32(0)
												} else {
													v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+6)))
													v242 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v87 + (v241+v242)&int32(_a_F_gin_leafpage_items_9) + int32(8)
													v250 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
													*(*int64)(unsafe.Add(mBase, uint32(v85))) = v250 + int64(1)
													v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v254)+20)) = v242
													v273 = v239
													m.G0 = v16 + int32(48)
													return v273
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
						v258 = m.ExcPending
						if v258 != 0 {
							return int32(0)
						} else {
							v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v259)+20)) = int32(2)
							v269 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v269)
							v273 = int32(0)
							m.G0 = v16 + int32(48)
							return v273
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v292 = m.ExcPending
				if v292 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v295 = m.ExcPending
					if v295 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_gin_leafpage_items_10), int32(0))
						mBase = m.M
						v299 = m.ExcPending
						if v299 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_gin_leafpage_items_4), int32(188), int32(_a_F_gin_leafpage_items_5))
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
			v17 = F_DirectFunctionCall2Coll(m, int32(1327), int32(0), v4, v3)
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int64
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = F_superuser(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L66
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L61
	}
L5:
	;
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v17 = F_get_page_from_raw(m, v11)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
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
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L57
	}
L9:
	;
	m.G0 = v8 + int32(112)
	return v194
L10:
	;
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+14)))
	if v19 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
	v194 = int32(0)
	goto L9
L12:
	;
	goto L13
L13:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+19)))
	v26 = int32(8)
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)))
	if (v25<<(uint(v26)%32)-v28)&int32(_a_F_gin_page_opaque_info_0) != v26 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v37 = F_get_call_result_type(m, l0, int32(0), v8+int32(108))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v37 != int32(1) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v42 = v8 + int32(16)
	v43 = v17 + v28
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+6)))
	if v44&int32(1) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v50 = F_cstring_to_text(m, int32(_a_F_gin_page_opaque_info_6))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v54 = v42
	v55 = int32(0)
	goto L19
L19:
	;
	if v44&int32(2) != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v50
	v54 = v42 | int32(4)
	v55 = int32(1)
	goto L19
L21:
	;
	v59 = F_cstring_to_text(m, int32(_a_F_gin_page_opaque_info_7))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	v64 = v55
	goto L23
L23:
	;
	if v44&int32(4) != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v59
	v64 = v55 + int32(1)
	goto L23
L25:
	;
	v73 = F_cstring_to_text(m, int32(_a_F_gin_page_opaque_info_8))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v78 = v64
	goto L27
L27:
	;
	if v44&int32(8) != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(16)|v64<<(uint(int32(2))%32)))) = v73
	v78 = v64 + int32(1)
	goto L27
L29:
	;
	v87 = F_cstring_to_text(m, int32(_a_F_gin_page_opaque_info_9))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	v92 = v78
	goto L31
L31:
	;
	if v44&int32(16) != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(16)+v78<<(uint(int32(2))%32)))) = v87
	v92 = v78 + int32(1)
	goto L31
L33:
	;
	v101 = F_cstring_to_text(m, int32(_a_F_gin_page_opaque_info_10))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	v106 = v92
	goto L35
L35:
	;
	if v44&int32(32) != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(16)+v92<<(uint(int32(2))%32)))) = v101
	v106 = v92 + int32(1)
	goto L35
L37:
	;
	v115 = F_cstring_to_text(m, int32(_a_F_gin_page_opaque_info_11))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	v120 = v106
	goto L39
L39:
	;
	if v44&int32(64) != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(16)+v106<<(uint(int32(2))%32)))) = v115
	v120 = v106 + int32(1)
	goto L39
L41:
	;
	v129 = F_cstring_to_text(m, int32(_a_F_gin_page_opaque_info_12))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	v134 = v120
	goto L43
L43:
	;
	if v44&int32(128) != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(16)+v120<<(uint(int32(2))%32)))) = v129
	v134 = v120 + int32(1)
	goto L43
L45:
	;
	v143 = F_cstring_to_text(m, int32(_a_F_gin_page_opaque_info_13))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	v148 = v134
	goto L47
L47:
	;
	v150 = v44 & int32(_a_F_gin_page_opaque_info_14)
	if v150 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(16)+v134<<(uint(int32(2))%32)))) = v143
	v148 = v134 + int32(1)
	goto L47
L49:
	;
	v158 = F_DirectFunctionCall1Coll(m, int32(2895), int32(0), v150)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	v163 = v148
	goto L51
L51:
	;
	v164 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+94)) = uint8(v164)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+92)) = uint16(v164)
	v168 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v43))))
	v169 = F_Int64GetDatum(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(16)+v148<<(uint(int32(2))%32)))) = v158
	v163 = v148 + int32(1)
	goto L51
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = v169
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+100)) = v172
	v177 = F_construct_array_builtin(m, v8+int32(16), v163, int32(25))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+104)) = v177
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v8)+108))
	v185 = F_heap_form_tuple(m, v180, v8+int32(96), v8+int32(92))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v185)+16))
	v188 = F_HeapTupleHeaderGetDatum(m, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v194 = v188
	goto L9
L57:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(_a_F_gin_page_opaque_info_15), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_gin_page_opaque_info_3), int32(112), int32(_a_F_gin_page_opaque_info_4))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
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
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errmsg(m, int32(_a_F_gin_page_opaque_info_1), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+19)))
	v228 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = (v227<<(uint(v228)%32) - v226) & int32(_a_F_gin_page_opaque_info_0)
	F_errdetail(m, int32(_a_F_gin_page_opaque_info_2), v8)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_gin_page_opaque_info_3), int32(125), int32(_a_F_gin_page_opaque_info_4))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_gin_page_opaque_info_5), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_gin_page_opaque_info_3), int32(131), int32(_a_F_gin_page_opaque_info_4))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
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
