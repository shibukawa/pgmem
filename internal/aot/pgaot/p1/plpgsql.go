package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_plpgsql_build_recfield(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if int32(0) <= v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v113
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _consts[1273]))
	v13 = v6
	goto L5
L3:
	;
	goto L4
L4:
	;
	v57 = F_palloc0(m, int32(48))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v10+v13<<(uint(int32(2))%32))))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v24 == int32(0) {
		v43 = v23
		v44 = v24
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L4
L7:
	;
	if v44-v43 == int32(0) {
		v113 = v19
		goto L1
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	if v23 != v24 {
		v43 = v23
		v44 = v24
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v28 = v20
	v29 = l1
	goto L11
L11:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	if v33 == int32(0) {
		v43 = v32
		v44 = v33
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v43 = v32
	v44 = v33
	goto L8
L13:
	;
	v36 = int32(1)
	if v32 == v33 {
		v28 = v28 + v36
		v29 = v29 + v36
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if int32(0) <= v48 {
		v13 = v48
		goto L5
	} else {
		goto L16
	}
L16:
	;
	goto L6
L17:
	;
	return int32(0)
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = int32(3)
	v63 = F_pstrdup(m, l1)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v63
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v57)+24)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v66
	v73 = *(*int32)(unsafe.Add(mBase, _consts[1273]))
	v76 = *(*int32)(unsafe.Add(mBase, _consts[1274]))
	v79 = *(*int32)(unsafe.Add(mBase, _consts[1275]))
	if v76 == v79 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1275])) = v76 << (uint(int32(1)) % 32)
	v91 = F_repalloc(m, v73, v76<<(uint(int32(3))%32))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L17
	} else {
		goto L23
	}
L21:
	;
	v97 = v76
	v98 = v73
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v97
	*(*int32)(unsafe.Add(mBase, _consts[1274])) = v97 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v98+v97<<(uint(int32(2))%32)))) = v57
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+16)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v97
	v113 = v57
	goto L1
L23:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1273])) = v91
	v96 = *(*int32)(unsafe.Add(mBase, _consts[1274]))
	v97 = v96
	v98 = v91
	goto L22
}
func F_plpgsql_compile(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
	v9 = F_cached_function_compile(m, l0, v4, int32(6729), int32(4752), int32(536), int32(0), l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v9
		return v9
	}
}
func F_plpgsql_create_econtext(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
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
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v5 == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[1276]))
		if v9 == int32(0) {
			v12 = int32(4489152)
			v13 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v15 = *(*int32)(unsafe.Add(mBase, _consts[190]))
			*(*int32)(unsafe.Add(mBase, _consts[3])) = v15
			v17 = F_CreateExecutorState(m)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v13
				*(*int32)(unsafe.Add(mBase, _consts[1276])) = v17
				v22 = v17
				*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v22
				v26 = v22
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
				if v29 == int32(0) {
					v33 = *(*int32)(unsafe.Add(mBase, _consts[1277]))
					if v33 == int32(0) {
						v40 = *(*int32)(unsafe.Add(mBase, _consts[176]))
						v43 = F_ResourceOwnerCreate(m, v40, int32(146116))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[1277])) = v43
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
							v47 = v46
							v48 = v43
							*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v48
							v50 = v47
							v52 = F_CreateExprContext(m, v50)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v52
								v56 = *(*int32)(unsafe.Add(mBase, _consts[190]))
								v58 = F_MemoryContextAlloc(m, v56, int32(12))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
									*(*int32)(unsafe.Add(mBase, uint32(v58))) = v60
									v63 = *(*int32)(unsafe.Add(mBase, _consts[4]))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v64
									v66 = int32(4582368)
									v67 = *(*int32)(unsafe.Add(mBase, _consts[1278]))
									*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v67
									*(*int32)(unsafe.Add(mBase, _consts[1278])) = v58
									return
								}
							}
						}
					} else {
						v47 = v26
						v48 = v33
						*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v48
						v50 = v47
						v52 = F_CreateExprContext(m, v50)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v52
							v56 = *(*int32)(unsafe.Add(mBase, _consts[190]))
							v58 = F_MemoryContextAlloc(m, v56, int32(12))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
								*(*int32)(unsafe.Add(mBase, uint32(v58))) = v60
								v63 = *(*int32)(unsafe.Add(mBase, _consts[4]))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v64
								v66 = int32(4582368)
								v67 = *(*int32)(unsafe.Add(mBase, _consts[1278]))
								*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v67
								*(*int32)(unsafe.Add(mBase, _consts[1278])) = v58
								return
							}
						}
					}
				} else {
					v50 = v26
					v52 = F_CreateExprContext(m, v50)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v52
						v56 = *(*int32)(unsafe.Add(mBase, _consts[190]))
						v58 = F_MemoryContextAlloc(m, v56, int32(12))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
							*(*int32)(unsafe.Add(mBase, uint32(v58))) = v60
							v63 = *(*int32)(unsafe.Add(mBase, _consts[4]))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v64
							v66 = int32(4582368)
							v67 = *(*int32)(unsafe.Add(mBase, _consts[1278]))
							*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v67
							*(*int32)(unsafe.Add(mBase, _consts[1278])) = v58
							return
						}
					}
				}
			}
		} else {
			v22 = v9
			*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v22
			v26 = v22
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
			if v29 == int32(0) {
				v33 = *(*int32)(unsafe.Add(mBase, _consts[1277]))
				if v33 == int32(0) {
					v40 = *(*int32)(unsafe.Add(mBase, _consts[176]))
					v43 = F_ResourceOwnerCreate(m, v40, int32(146116))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[1277])) = v43
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
						v47 = v46
						v48 = v43
						*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v48
						v50 = v47
						v52 = F_CreateExprContext(m, v50)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v52
							v56 = *(*int32)(unsafe.Add(mBase, _consts[190]))
							v58 = F_MemoryContextAlloc(m, v56, int32(12))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
								*(*int32)(unsafe.Add(mBase, uint32(v58))) = v60
								v63 = *(*int32)(unsafe.Add(mBase, _consts[4]))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v64
								v66 = int32(4582368)
								v67 = *(*int32)(unsafe.Add(mBase, _consts[1278]))
								*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v67
								*(*int32)(unsafe.Add(mBase, _consts[1278])) = v58
								return
							}
						}
					}
				} else {
					v47 = v26
					v48 = v33
					*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v48
					v50 = v47
					v52 = F_CreateExprContext(m, v50)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v52
						v56 = *(*int32)(unsafe.Add(mBase, _consts[190]))
						v58 = F_MemoryContextAlloc(m, v56, int32(12))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
							*(*int32)(unsafe.Add(mBase, uint32(v58))) = v60
							v63 = *(*int32)(unsafe.Add(mBase, _consts[4]))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v64
							v66 = int32(4582368)
							v67 = *(*int32)(unsafe.Add(mBase, _consts[1278]))
							*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v67
							*(*int32)(unsafe.Add(mBase, _consts[1278])) = v58
							return
						}
					}
				}
			} else {
				v50 = v26
				v52 = F_CreateExprContext(m, v50)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v52
					v56 = *(*int32)(unsafe.Add(mBase, _consts[190]))
					v58 = F_MemoryContextAlloc(m, v56, int32(12))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						*(*int32)(unsafe.Add(mBase, uint32(v58))) = v60
						v63 = *(*int32)(unsafe.Add(mBase, _consts[4]))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v64
						v66 = int32(4582368)
						v67 = *(*int32)(unsafe.Add(mBase, _consts[1278]))
						*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v67
						*(*int32)(unsafe.Add(mBase, _consts[1278])) = v58
						return
					}
				}
			}
		}
	} else {
		v26 = v5
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		if v29 == int32(0) {
			v33 = *(*int32)(unsafe.Add(mBase, _consts[1277]))
			if v33 == int32(0) {
				v40 = *(*int32)(unsafe.Add(mBase, _consts[176]))
				v43 = F_ResourceOwnerCreate(m, v40, int32(146116))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[1277])) = v43
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
					v47 = v46
					v48 = v43
					*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v48
					v50 = v47
					v52 = F_CreateExprContext(m, v50)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v52
						v56 = *(*int32)(unsafe.Add(mBase, _consts[190]))
						v58 = F_MemoryContextAlloc(m, v56, int32(12))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
							*(*int32)(unsafe.Add(mBase, uint32(v58))) = v60
							v63 = *(*int32)(unsafe.Add(mBase, _consts[4]))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v64
							v66 = int32(4582368)
							v67 = *(*int32)(unsafe.Add(mBase, _consts[1278]))
							*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v67
							*(*int32)(unsafe.Add(mBase, _consts[1278])) = v58
							return
						}
					}
				}
			} else {
				v47 = v26
				v48 = v33
				*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v48
				v50 = v47
				v52 = F_CreateExprContext(m, v50)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v52
					v56 = *(*int32)(unsafe.Add(mBase, _consts[190]))
					v58 = F_MemoryContextAlloc(m, v56, int32(12))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						*(*int32)(unsafe.Add(mBase, uint32(v58))) = v60
						v63 = *(*int32)(unsafe.Add(mBase, _consts[4]))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v64
						v66 = int32(4582368)
						v67 = *(*int32)(unsafe.Add(mBase, _consts[1278]))
						*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v67
						*(*int32)(unsafe.Add(mBase, _consts[1278])) = v58
						return
					}
				}
			}
		} else {
			v50 = v26
			v52 = F_CreateExprContext(m, v50)
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v52
				v56 = *(*int32)(unsafe.Add(mBase, _consts[190]))
				v58 = F_MemoryContextAlloc(m, v56, int32(12))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
					*(*int32)(unsafe.Add(mBase, uint32(v58))) = v60
					v63 = *(*int32)(unsafe.Add(mBase, _consts[4]))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v64
					v66 = int32(4582368)
					v67 = *(*int32)(unsafe.Add(mBase, _consts[1278]))
					*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v67
					*(*int32)(unsafe.Add(mBase, _consts[1278])) = v58
					return
				}
			}
		}
	}
}
func F_plpgsql_exec_event_trigger(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(160)
	m.G0 = v8
	F_plpgsql_estate_setup(m, v8+int32(16), l0, v3, v3, v3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(6734)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8)+148)) = int32(351676)
	v22 = int32(4482056)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v8 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v23
	v29 = v8 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v29
	F_copy_plpgsql_datums(m, v29, l0)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[1281]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v37 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v76 != 0 {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+148)) = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+140)) = v42
	v74 = v42
	goto L4
L6:
	;
	goto L7
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v44 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	if v66 == int32(0) {
		v74 = v65
		goto L4
	} else {
		goto L14
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+148)) = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+140)) = v49
	v64 = v37
	v65 = v49
	goto L8
L10:
	;
	goto L11
L11:
	;
	m.T0[v44].(func(*base.Module, int32, int32))(m, v8+int32(16), l0)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[1281]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v58 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+148)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+140)) = v60
	if v57 == v58 {
		v74 = v60
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v64 = v57
	v65 = v60
	goto L8
L14:
	;
	m.T0[v66].(func(*base.Module, int32, int32))(m, v8+int32(16), v65)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v74 = v65
	goto L4
L16:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v82 = F_exec_stmt_block(m, v8+int32(16), v74)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _consts[1281]))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v85 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+140)) = int32(0)
	if v82 == int32(2) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	if v88 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	m.T0[v88].(func(*base.Module, int32, int32))(m, v8+int32(16), v74)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+148)) = int32(99326)
	v103 = *(*int32)(unsafe.Add(mBase, _consts[1281]))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v104 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+148)) = int32(0)
	F_errstart_cold(m, int32(21), int32(547943))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L39
	}
L28:
	;
	v115 = int32(4582368)
	v116 = *(*int32)(unsafe.Add(mBase, _consts[1278]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	F_pfree(m, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L32
	}
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	if v107 == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	m.T0[v107].(func(*base.Module, int32, int32))(m, v8+int32(16), l0)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1278])) = v117
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v8)+136))
	F_FreeExprContext(m, v121, int32(1))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v125 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+136)) = v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v8)+120))
	if v127 == v125 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v142
	m.G0 = v8 + int32(160)
	return
L35:
	;
	F_SPI_freetuptable(m, v127)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v132 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+120)) = v132
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v8)+136))
	if v134 == v132 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
	F_MemoryContextReset(m, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	F_errcode(m, int32(83887490))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errmsg(m, int32(526490), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(498212), int32(1217), int32(223176))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_plpgsql_extra_checks_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v215 int32
	_ = v215
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v347 int32
	_ = v347
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = v13
	v18 = int32(304442)
	goto L4
L1:
	;
	m.G0 = v11 + int32(32)
	return v441
L2:
	;
	v433 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L33
	} else {
		goto L137
	}
L3:
	;
	if v55 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v21 == v22 {
		v44 = v21
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v55 = int32(0)
	goto L3
L6:
	;
	v46 = int32(1)
	if v44 != 0 {
		v17 = v17 + v46
		v18 = v18 + v46
		goto L4
	} else {
		goto L15
	}
L7:
	;
	if base.Ui32((v21-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v32 = v21 | int32(32)
	goto L10
L9:
	;
	v32 = v21
	goto L10
L10:
	;
	if base.Ui32((v22-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v41 = v22 | int32(32)
	goto L13
L12:
	;
	v41 = v22
	goto L13
L13:
	;
	if v32 == v41 {
		v44 = v32
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v55 = v32 - v41
	goto L3
L15:
	;
	goto L5
L16:
	;
	v429 = int32(-1)
	goto L2
L17:
	;
	goto L18
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v63 = v59
	v64 = int32(371483)
	goto L20
L19:
	;
	if v101 == int32(0) {
		v429 = v4
		goto L2
	} else {
		goto L32
	}
L20:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v67 == v68 {
		v90 = v67
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v101 = int32(0)
	goto L19
L22:
	;
	v92 = int32(1)
	if v90 != 0 {
		v63 = v63 + v92
		v64 = v64 + v92
		goto L20
	} else {
		goto L31
	}
L23:
	;
	if base.Ui32((v67-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v78 = v67 | int32(32)
	goto L26
L25:
	;
	v78 = v67
	goto L26
L26:
	;
	if base.Ui32((v68-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v87 = v68 | int32(32)
	goto L29
L28:
	;
	v87 = v68
	goto L29
L29:
	;
	if v78 == v87 {
		v90 = v78
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v101 = v78 - v87
	goto L19
L31:
	;
	goto L21
L32:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v105 = F_pstrdup(m, v104)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	return int32(0)
L34:
	;
	v112 = F_SplitIdentifierString(m, v105, int32(44), v11+int32(28))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	if v112 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	*(*int32)(unsafe.Add(mBase, _consts[425])) = v117
	goto L39
L37:
	;
	goto L38
L38:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if v132 == int32(0) {
		v416 = v4
		goto L43
	} else {
		goto L44
	}
L39:
	;
	v120 = int32(0)
	v124 = F_format_elog_string(m, int32(629205), v120)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L33
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, _consts[825])) = v124
	F_pfree(m, v105)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L33
	} else {
		goto L41
	}
L41:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_list_free(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L33
	} else {
		goto L42
	}
L42:
	;
	v441 = v120
	goto L1
L43:
	;
	F_pfree(m, v105)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L33
	} else {
		goto L135
	}
L44:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v135 <= int32(0) {
		v416 = v4
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v139 = int32(0)
	v144 = v4
	goto L46
L46:
	;
	v147 = int32(2)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148+v139<<(uint(v147)%32))))
	v156 = v152
	v157 = int32(166363)
	goto L50
L47:
	;
	v416 = v406
	goto L43
L48:
	;
	v406 = v144 | v405
	v408 = v139 + int32(1)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v408 < v409 {
		v139 = v408
		v144 = v406
		goto L46
	} else {
		goto L134
	}
L49:
	;
	if v194 == int32(0) {
		v405 = v147
		goto L48
	} else {
		goto L62
	}
L50:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v160 == v161 {
		v183 = v160
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v194 = int32(0)
	goto L49
L52:
	;
	v185 = int32(1)
	if v183 != 0 {
		v156 = v156 + v185
		v157 = v157 + v185
		goto L50
	} else {
		goto L61
	}
L53:
	;
	if base.Ui32((v160-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v171 = v160 | int32(32)
	goto L56
L55:
	;
	v171 = v160
	goto L56
L56:
	;
	if base.Ui32((v161-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v180 = v161 | int32(32)
	goto L59
L58:
	;
	v180 = v161
	goto L59
L59:
	;
	if v171 == v180 {
		v183 = v171
		goto L52
	} else {
		goto L60
	}
L60:
	;
	v194 = v171 - v180
	goto L49
L61:
	;
	goto L51
L62:
	;
	v200 = v152
	v201 = int32(113273)
	goto L64
L63:
	;
	if v238 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L64:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	if v204 == v205 {
		v227 = v204
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v238 = int32(0)
	goto L63
L66:
	;
	v229 = int32(1)
	if v227 != 0 {
		v200 = v200 + v229
		v201 = v201 + v229
		goto L64
	} else {
		goto L75
	}
L67:
	;
	if base.Ui32((v204-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v215 = v204 | int32(32)
	goto L70
L69:
	;
	v215 = v204
	goto L70
L70:
	;
	if base.Ui32((v205-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v224 = v205 | int32(32)
	goto L73
L72:
	;
	v224 = v205
	goto L73
L73:
	;
	if v215 == v224 {
		v227 = v215
		goto L66
	} else {
		goto L74
	}
L74:
	;
	v238 = v215 - v224
	goto L63
L75:
	;
	goto L65
L76:
	;
	v405 = int32(4)
	goto L48
L77:
	;
	goto L78
L78:
	;
	v245 = v152
	v246 = int32(94652)
	goto L80
L79:
	;
	if v283 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L80:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	if v249 == v250 {
		v272 = v249
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v283 = int32(0)
	goto L79
L82:
	;
	v274 = int32(1)
	if v272 != 0 {
		v245 = v245 + v274
		v246 = v246 + v274
		goto L80
	} else {
		goto L91
	}
L83:
	;
	if base.Ui32((v249-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v260 = v249 | int32(32)
	goto L86
L85:
	;
	v260 = v249
	goto L86
L86:
	;
	if base.Ui32((v250-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v269 = v250 | int32(32)
	goto L89
L88:
	;
	v269 = v250
	goto L89
L89:
	;
	if v260 == v269 {
		v272 = v260
		goto L82
	} else {
		goto L90
	}
L90:
	;
	v283 = v260 - v269
	goto L79
L91:
	;
	goto L81
L92:
	;
	v405 = int32(8)
	goto L48
L93:
	;
	goto L94
L94:
	;
	v290 = v152
	v291 = int32(304442)
	goto L97
L95:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	*(*int32)(unsafe.Add(mBase, _consts[425])) = v389
	goto L130
L96:
	;
	if v328 != 0 {
		goto L109
	} else {
		goto L110
	}
L97:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290))))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	if v294 == v295 {
		v317 = v294
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v328 = int32(0)
	goto L96
L99:
	;
	v319 = int32(1)
	if v317 != 0 {
		v290 = v290 + v319
		v291 = v291 + v319
		goto L97
	} else {
		goto L108
	}
L100:
	;
	if base.Ui32((v294-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v305 = v294 | int32(32)
	goto L103
L102:
	;
	v305 = v294
	goto L103
L103:
	;
	if base.Ui32((v295-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v314 = v295 | int32(32)
	goto L106
L105:
	;
	v314 = v295
	goto L106
L106:
	;
	if v305 == v314 {
		v317 = v305
		goto L99
	} else {
		goto L107
	}
L107:
	;
	v328 = v305 - v314
	goto L96
L108:
	;
	goto L98
L109:
	;
	v332 = v152
	v333 = int32(371483)
	goto L113
L110:
	;
	goto L111
L111:
	;
	v373 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	*(*int32)(unsafe.Add(mBase, _consts[425])) = v373
	goto L126
L112:
	;
	if v370 != 0 {
		goto L95
	} else {
		goto L125
	}
L113:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	if v336 == v337 {
		v359 = v336
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v370 = int32(0)
	goto L112
L115:
	;
	v361 = int32(1)
	if v359 != 0 {
		v332 = v332 + v361
		v333 = v333 + v361
		goto L113
	} else {
		goto L124
	}
L116:
	;
	if base.Ui32((v336-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v347 = v336 | int32(32)
	goto L119
L118:
	;
	v347 = v336
	goto L119
L119:
	;
	if base.Ui32((v337-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v356 = v337 | int32(32)
	goto L122
L121:
	;
	v356 = v337
	goto L122
L122:
	;
	if v347 == v356 {
		v359 = v347
		goto L115
	} else {
		goto L123
	}
L123:
	;
	v370 = v347 - v356
	goto L112
L124:
	;
	goto L114
L125:
	;
	goto L111
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v152
	v379 = F_format_elog_string(m, int32(585310), v11)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L33
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, _consts[825])) = v379
	F_pfree(m, v105)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L33
	} else {
		goto L128
	}
L128:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_list_free(m, v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L33
	} else {
		goto L129
	}
L129:
	;
	v441 = int32(0)
	goto L1
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v152
	v397 = F_format_elog_string(m, int32(650910), v11+int32(16))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L33
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, _consts[825])) = v397
	F_pfree(m, v105)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L33
	} else {
		goto L132
	}
L132:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_list_free(m, v402)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L33
	} else {
		goto L133
	}
L133:
	;
	v441 = int32(0)
	goto L1
L134:
	;
	goto L47
L135:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_list_free(m, v421)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L33
	} else {
		goto L136
	}
L136:
	;
	v429 = v416
	goto L2
L137:
	;
	if v433 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v441 = int32(0)
	goto L1
L139:
	;
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v433))) = v429
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v433
	v441 = int32(1)
	goto L1
}
func F_plpgsql_free_function_memory(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+504))
	if v2 < v11 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(547943))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L13
	} else {
		goto L27
	}
L2:
	;
	v16 = v2
	goto L5
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+504)) = int32(0)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	F_free_stmt(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L13
	} else {
		goto L22
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+508))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v16<<(uint(int32(2))%32))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	switch v25 {
	case 0, 4:
		goto L9
	case 1, 3:
		goto L7
	case 2:
		goto L8
	default:
		goto L1
	}
L6:
	;
	goto L4
L7:
	;
	v61 = v16 + int32(1)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+504))
	if v61 < v62 {
		v16 = v61
		goto L5
	} else {
		goto L21
	}
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if v47 == int32(0) {
		goto L7
	} else {
		goto L18
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if v26 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v37 == int32(0) {
		goto L7
	} else {
		goto L15
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	if v29 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	F_SPI_freeplan(m, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = int32(0)
	goto L10
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	if v40 == int32(0) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	F_SPI_freeplan(m, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = int32(0)
	goto L7
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	if v50 == int32(0) {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	F_SPI_freeplan(m, v50)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = int32(0)
	goto L7
L21:
	;
	goto L6
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+516)) = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_MemoryContextDelete(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L13
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	m.G0 = v9 + int32(16)
	return
L26:
	;
	goto L25
L27:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v91
	F_errmsg_internal(m, int32(484385), v9)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L13
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(493160), int32(751), int32(13660))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_plpgsql_getdiag_kindname(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	if base.Ui32(int32(12)) < base.Ui32(l0) {
		return int32(243120)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[1283])))
		return v10
	}
}
func F_plpgsql_ns_top(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, _consts[1282]))
	return v2
}
func F_plpgsql_param_eval_var_check(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
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
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16+v17<<(uint(int32(2))%32)-int32(4))))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+44)))
	if v25 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return
L2:
	;
	switch v132 - int32(1) {
	case 0:
		goto L41
	case 1:
		goto L40
	case 2:
		goto L39
	default:
		goto L1
	}
L3:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v132 = v131
	goto L2
L4:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v69 == int32(27) {
		goto L25
	} else {
		goto L26
	}
L5:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v24
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v25)
	goto L1
L6:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v26 != int32(1) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v29 != int32(3) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	if v34 != 0 {
		v132 = v34
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v33)+48)) = int64(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+20)))
	if v38 != int32(1) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v17
	if v37 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v46 == int32(8) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = int32(2)
	goto L3
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v49 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v55 = F_expression_tree_walker_impl(m, v37, int32(6747), v12+int32(8))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	if v50 == v17 {
		v61 = v37
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L4
L18:
	;
	return
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v57 != int32(1) {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v61 = v60
	goto L12
L21:
	;
	v109 = F_get_func_support(m, v106)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L18
	} else {
		goto L35
	}
L22:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v84 = F_SearchSysCache1(m, int32(82), v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L18
	} else {
		goto L29
	}
L23:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	v106 = v80
	v108 = v81
	goto L21
L24:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	v106 = v78
	v108 = v79
	goto L21
L25:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v74 = v72
	v75 = v73
	goto L27
L26:
	;
	v74 = v37
	v75 = v69
	goto L27
L27:
	;
	switch v75 - int32(14) {
	case 0:
		goto L22
	case 1:
		goto L24
	default:
		goto L3
	case 3:
		goto L23
	}
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v74)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v74)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v97
	v103 = F_list_make2_impl(m, v12+int32(4), v12)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L18
	} else {
		goto L34
	}
L29:
	;
	if v84 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+22)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v86+v87)+88))
	F_ReleaseCatCache(m, v84)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L18
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v94 = int32(0)
	goto L28
L33:
	;
	v94 = v89
	goto L28
L34:
	;
	v106 = v94
	v108 = v103
	goto L21
L35:
	;
	if v109 == int32(0) {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(464)
	v121 = F_OidFunctionCall1Coll(m, v109, int32(0), v12+int32(8))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L18
	} else {
		goto L37
	}
L37:
	;
	if v121 == int32(0) {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+52)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = int32(3)
	goto L3
L39:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	if v210 == v32 {
		goto L54
	} else {
		goto L55
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(6748)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+68))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v174+v175<<(uint(int32(2))%32)-int32(4))))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+40))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+44)))
	if v183 != 0 {
		goto L49
	} else {
		goto L50
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(6742)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+68))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v141+v142<<(uint(int32(2))%32)-int32(4))))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+40))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+44)))
	if v150 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v153 != int32(1) {
		v162 = v149
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v164 = v150
	v165 = v149
	goto L44
L44:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v165
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v168))) = uint8(v164)
	goto L1
L45:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+44)))
	v164 = v163
	v165 = v162
	goto L44
L46:
	;
	goto L45
L47:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
	if v156 != int32(3) {
		v162 = v149
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v149)+2))
	v162 = v159 + int32(18)
	goto L46
L49:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v182
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v208))) = uint8(v183)
	goto L1
L50:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if v184 != int32(1) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
	if v187 != int32(3) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v173)+120))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+20))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v182)+2))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	F_MemoryContextSetParent(m, v193, v191)
	mBase = m.M
	goto L53
L53:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v192 + int32(12)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v200 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v199))) = uint8(v200)
	v202 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v181)+44)) = uint16(v202)
	*(*int32)(unsafe.Add(mBase, uint32(v181)+40)) = v200
	goto L1
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(6743)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+68))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v217+v218<<(uint(int32(2))%32)-int32(4))))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v227))) = uint8(v228)
	goto L1
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(6742)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+68))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v235+v236<<(uint(int32(2))%32)-int32(4))))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+40))
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+44)))
	if v244 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L1
L58:
	;
	v247 = F_MakeExpandedObjectReadOnlyInternal(m, v243)
	mBase = m.M
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+44)))
	v249 = v247
	v250 = v248
	goto L60
L59:
	;
	v249 = v243
	v250 = v244
	goto L60
L60:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v249
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v253))) = uint8(v250)
	goto L57
}
func F_plpgsql_scanner_init(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v6 = F_palloc0(m, int32(200))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v14 = F_scanner_init(m, l0, v6, int32(4158836), int32(2154304))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = l0
			v19 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[1284])) = v19
			*(*int64)(unsafe.Add(mBase, uint32(v6)+68)) = int64(0)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = int32(1)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
			*(*int32)(unsafe.Add(mBase, uint32(v23)+188)) = v26
			v28 = int32(10)
			v29 = F___strchrnul(m, v26, v28)
			mBase = m.M
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
			if v31 == v28 {
				v35 = v29
			} else {
				v35 = v19
			}
			*(*int32)(unsafe.Add(mBase, uint32(v23)+192)) = v35
			return v14
		}
	}
}
func F_plpgsql_sql_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = F_plpgsql_scanner_errposition(m, v3, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = F_geterrposition(m)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			if v7 <= int32(0) {
				v22 = F_errposition(m, int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					return
				}
			} else {
				v11 = F_getinternalerrposition(m)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					if v11 <= int32(0) {
						v22 = F_errposition(m, int32(0))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							return
						}
					} else {
						F_internalerrposition(m, v7+v11-int32(1))
						mBase = m.M
						v19 = m.ExcPending
						if v19 != 0 {
							return
						} else {
							v22 = F_errposition(m, int32(0))
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
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
