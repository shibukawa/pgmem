package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CatCacheRemoveCList(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	v3 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v14 = v12 - int32(1)
	if v3 <= v14 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = v14
	goto L4
L2:
	;
	goto L3
L3:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+4)) = v137
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v139
	v141 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+54)))
	if int32(0) < v141 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1-int32(-64)+v23<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+76)) = int32(0)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+52)))
	if v38 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	if int32(0) < v23 {
		v23 = v23 - int32(1)
		goto L4
	} else {
		goto L21
	}
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
	if v41 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v45
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+53)))
	if v47 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_pfree(m, v35)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L17
	} else {
		goto L20
	}
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v50 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v59 = int32(0)
	goto L12
L12:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(48)+v59<<(uint(int32(2))%32))))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v71<<(uint(int32(3))%32))+24)))
	if v75 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L9
L14:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v35+int32(16)+v59<<(uint(int32(3))%32))))
	F_pfree(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v85 = v59 + int32(1)
	if v85 != v50 {
		v59 = v85
		goto L12
	} else {
		goto L19
	}
L17:
	;
	return
L18:
	;
	goto L16
L19:
	;
	goto L13
L20:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v101 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v100 - v101
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_CatCacheRemoveCList[0]))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v106 - v101
	goto L6
L21:
	;
	goto L5
L22:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v155 = v3
	goto L25
L23:
	;
	goto L24
L24:
	;
	F_pfree(m, l1)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L17
	} else {
		goto L32
	}
L25:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(48)+v155<<(uint(int32(2))%32))))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148+v163<<(uint(int32(3))%32))+24)))
	if v167 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L24
L27:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(16)+v155<<(uint(int32(3))%32))))
	F_pfree(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L17
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v177 = v155 + int32(1)
	if v177 != v141 {
		v155 = v177
		goto L25
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	goto L26
L32:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v192 - int32(1)
	return
}
func F_InitCatCachePhase2(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3 == int32(0) {
		F_CatalogCacheInitializeCache(m, l0)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			if l1 == int32(0) {
				return
			} else {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if base.Ui32(v10-int32(1)) < base.Ui32(int32(2)) {
					return
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					F_LockRelationOid(m, v15, int32(1))
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						v21 = F_index_open(m, v19, int32(1))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							F_relation_close(m, v21, int32(1))
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
								F_UnlockRelationOid(m, v26, int32(1))
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
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
		if l1 == int32(0) {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if base.Ui32(v10-int32(1)) < base.Ui32(int32(2)) {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
				F_LockRelationOid(m, v15, int32(1))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
					v21 = F_index_open(m, v19, int32(1))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						F_relation_close(m, v21, int32(1))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							F_UnlockRelationOid(m, v26, int32(1))
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
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
}
func F_ReleaseCatCacheList(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseCatCacheList[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v5 - int32(1)
	if v4 != 0 {
		F_ResourceOwnerForget(m, v4, base.I64_extend_i32_u(l0), int32(_a_F_ReleaseCatCacheList_0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
			if v13 != int32(1) {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if v16 != 0 {
					return
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
					F_CatCacheRemoveCList(m, v17, l0)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
		if v13 != int32(1) {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if v16 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
				F_CatCacheRemoveCList(m, v17, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
