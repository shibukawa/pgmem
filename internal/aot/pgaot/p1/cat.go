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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	v3 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v16 = v14 - int32(1)
	if v3 <= v16 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = int32(48)
	v25 = v16
	goto L4
L2:
	;
	goto L3
L3:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v152)+4)) = v153
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v155
	v157 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+38)))
	if int32(0) < v157 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1+v19+v25<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = int32(0)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+36)))
	if v42 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	if int32(0) < v25 {
		v25 = v25 - int32(1)
		goto L4
	} else {
		goto L21
	}
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	if v45 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v49
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+37)))
	if v51 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_pfree(m, v39)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L17
	} else {
		goto L20
	}
L10:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v54 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v65 = int32(0)
	goto L12
L12:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v81 = v65 << (uint(int32(2)) % 32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0+v19+v81)))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+int32(2)+v76<<(uint(int32(4))%32)+v83*int32(100)))))
	if v87 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L9
L14:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v81+(v39+int32(8)))))
	F_pfree(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v95 = v65 + int32(1)
	if v95 != v54 {
		v65 = v95
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
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v113 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v112 - v113
	v117 = *(*int32)(unsafe.Add(mBase, _consts[1146]))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+4)) = v118 - v113
	goto L6
L21:
	;
	goto L5
L22:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v172 = v3
	goto L25
L23:
	;
	goto L24
L24:
	;
	F_pfree(m, l1)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L17
	} else {
		goto L32
	}
L25:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	v185 = v172 << (uint(int32(2)) % 32)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(48)+v185)))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+int32(2)+v180<<(uint(int32(4))%32)+v187*int32(100)))))
	if v191 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L24
L27:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v185+(l1+int32(16)))))
	F_pfree(m, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L17
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v199 = v172 + int32(1)
	if v199 != v157 {
		v172 = v199
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
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v216 - int32(1)
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v5 - int32(1)
	if v4 != 0 {
		F_ResourceOwnerForget(m, v4, l0, int32(1767832))
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
