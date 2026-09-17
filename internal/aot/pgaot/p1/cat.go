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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	v3 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v15 = v13 - int32(1)
	if v3 <= v15 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = int32(48)
	v24 = v15
	goto L4
L2:
	;
	goto L3
L3:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v144)+4)) = v145
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v147
	v149 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+38)))
	if int32(0) < v149 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1+v18+v24<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+60)) = int32(0)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+36)))
	if v40 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	if int32(0) < v24 {
		v24 = v24 - int32(1)
		goto L4
	} else {
		goto L21
	}
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	if v43 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v47
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+37)))
	if v49 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_pfree(m, v37)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L17
	} else {
		goto L20
	}
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v52 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v61 = int32(0)
	goto L12
L12:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v76 = v61 << (uint(int32(2)) % 32)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0+v18+v76)))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+v71<<(uint(int32(4))%32)+v78*int32(100))+2)))
	if v82 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L9
L14:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(8)+v76)))
	F_pfree(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v90 = v61 + int32(1)
	if v90 != v52 {
		v61 = v90
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
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v107 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v106 - v107
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_CatCacheRemoveCList[0]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v112 - v107
	goto L6
L21:
	;
	goto L5
L22:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v164 = v3
	goto L25
L23:
	;
	goto L24
L24:
	;
	F_pfree(m, l1)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L17
	} else {
		goto L32
	}
L25:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v174 = v164 << (uint(int32(2)) % 32)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(48)+v174)))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+v169<<(uint(int32(4))%32)+v176*int32(100))+2)))
	if v180 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L24
L27:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(16)+v174)))
	F_pfree(m, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L17
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v188 = v164 + int32(1)
	if v188 != v149 {
		v164 = v188
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
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v204 - int32(1)
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseCatCacheList[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v5 - int32(1)
	if v4 != 0 {
		F_ResourceOwnerForget(m, v4, l0, int32(_a_F_ReleaseCatCacheList_0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
			if v12 != int32(1) {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				if v15 != 0 {
					return
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					F_CatCacheRemoveCList(m, v16, l0)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
		if v12 != int32(1) {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			if v15 != 0 {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				F_CatCacheRemoveCList(m, v16, l0)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
