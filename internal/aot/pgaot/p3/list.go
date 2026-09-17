package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_assign_list_collations(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if l1 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v12 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = v3
	goto L4
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v18<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l0
	v32 = F_assign_collations_walker(m, v24, v8+int32(8))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	v35 = v18 + int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v35 < v36 {
		v18 = v35
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F_list_append_unique(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v3 = int32(0)
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v30
L2:
	;
	v28 = F_lappend(m, l0, l1)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L11
	}
L3:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v11 = v3
	goto L5
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12+v11<<(uint(int32(2))%32))))
	v17 = F_equal(m, v16, l1)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L2
L7:
	;
	return int32(0)
L8:
	;
	if v17 != 0 {
		v30 = l0
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v22 = v11 + int32(1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v22 < v23 {
		v11 = v22
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L6
L11:
	;
	v30 = v28
	goto L1
}
func F_list_concat(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	if l0 == int32(0) {
		if l1 == int32(0) {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v17 = int32(8)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v20 = v18 + int32(4)
			if v20 <= v17 {
				v23 = v17
			} else {
				v23 = v20
			}
			if v23&(v23-int32(1)) != 0 {
				v30 = int32(1) << (uint(int32(32)-base.I32_clz(v23)) % 32)
			} else {
				v30 = v23
			}
			v32 = v30 - int32(4)
			v37 = F_palloc(m, v32<<(uint(int32(2))%32)+int32(16))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v32
				*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v18
				*(*int32)(unsafe.Add(mBase, uint32(v37))) = v14
				v45 = v37 + int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v45
				v48 = v18 << (uint(int32(2)) % 32)
				if v48 == int32(0) {
					v115 = v37
					return v115
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					base.MemoryCopy(m, v45, v51, v48)
					return v37
				}
			}
		}
	} else {
		if l1 == int32(0) {
			v115 = l0
			return v115
		} else {
			v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v58 = v56 + v57
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v59 < v58 {
				v63 = int32(16)
				if v58 <= v63 {
					v66 = v63
				} else {
					v66 = v58
				}
				if v66&(v66-int32(1)) != 0 {
					v73 = int32(1) << (uint(int32(32)-base.I32_clz(v66)) % 32)
				} else {
					v73 = v66
				}
				v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v76 = l0 + int32(16)
				if v74 == v76 {
					v78 = F_GetMemoryChunkContext(m, l0)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						v82 = F_MemoryContextAlloc(m, v78, v73<<(uint(int32(2))%32))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v82
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v87 = v85 << (uint(int32(2)) % 32)
							if v87 == int32(0) {
							} else {
								base.MemoryCopy(m, v82, v76, v87)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v73
							v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v101 = v99
							v105 = v100
							v107 = v105 << (uint(int32(2)) % 32)
							if v107 != 0 {
								v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
								base.MemoryCopy(m, v108+v101<<(uint(int32(2))%32), v112, v107)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v58
							v115 = l0
							return v115
						}
					}
				} else {
					v93 = F_repalloc(m, v74, v73<<(uint(int32(2))%32))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v93
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v73
						v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v101 = v99
						v105 = v100
						v107 = v105 << (uint(int32(2)) % 32)
						if v107 != 0 {
							v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
							base.MemoryCopy(m, v108+v101<<(uint(int32(2))%32), v112, v107)
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v58
						v115 = l0
						return v115
					}
				}
			} else {
				v101 = v57
				v105 = v56
				v107 = v105 << (uint(int32(2)) % 32)
				if v107 != 0 {
					v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					base.MemoryCopy(m, v108+v101<<(uint(int32(2))%32), v112, v107)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v58
				v115 = l0
				return v115
			}
		}
	}
}
func F_list_copy_head(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	v3 = int32(0)
	if base.B2i32(l0 == v3)|base.B2i32(l1 <= v3) != 0 {
		v53 = v3
		return v53
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v15 < l1 {
			v17 = v15
		} else {
			v17 = l1
		}
		v19 = v17 + int32(4)
		if v19 <= int32(8) {
			v22 = int32(8)
		} else {
			v22 = v19
		}
		if v22&(v22-int32(1)) != 0 {
			v29 = int32(1) << (uint(int32(32)-base.I32_clz(v22)) % 32)
		} else {
			v29 = v22
		}
		v31 = v29 - int32(4)
		v36 = F_palloc(m, v31<<(uint(int32(2))%32)+int32(16))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v31
			*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v17
			*(*int32)(unsafe.Add(mBase, uint32(v36))) = v11
			v44 = v36 + int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v44
			v47 = v17 << (uint(int32(2)) % 32)
			if v47 == int32(0) {
				v53 = v36
			} else {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				base.MemoryCopy(m, v44, v50, v47)
				v53 = v36
			}
			return v53
		}
	}
}
func F_list_delete_last(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	if l0 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v3 <= int32(1) {
			v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v6 != l0+int32(16) {
				F_pfree(m, v6)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return int32(0)
				} else {
					F_pfree(m, l0)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				}
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3 - int32(1)
			v23 = l0
			return v23
		}
	} else {
		v23 = int32(0)
		return v23
	}
}
func F_list_delete_ptr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return l0
L5:
	;
	v15 = int32(0)
	if v15 < v12 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v18 = v12
	goto L8
L7:
	;
	v18 = v15
	goto L8
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v24 = v3
	goto L9
L9:
	;
	v29 = v19 + v24<<(uint(int32(2))%32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if l1 != v30 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v12 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v33 = v24 + int32(1)
	if v18 != v33 {
		v24 = v33
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L10
L14:
	;
	goto L4
L15:
	;
	if l0+int32(16) != v19 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v52 = (v12 + (v24 ^ int32(-1))) << (uint(int32(2)) % 32)
	if v52 != 0 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	F_pfree(m, v19)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	F_pfree(m, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L21
	} else {
		goto L23
	}
L21:
	;
	return int32(0)
L22:
	;
	goto L20
L23:
	;
	return int32(0)
L24:
	;
	base.MemoryCopy(m, v29, v29+int32(4), v52)
	goto L26
L25:
	;
	goto L26
L26:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v56 - int32(1)
	goto L4
}
func F_list_insert_nth(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	if l0 == int32(0) {
		v11 = F_palloc(m, int32(32))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(4)
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(4294967297)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v11 + int32(16)
			return v11
		}
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v25 <= v24 {
			v27 = int32(1)
			v29 = int32(16)
			v31 = v24 + v27
			if v31 <= v29 {
				v34 = v29
			} else {
				v34 = v31
			}
			if v34&(v34-int32(1)) != 0 {
				v41 = v27 << (uint(int32(32)-base.I32_clz(v34)) % 32)
			} else {
				v41 = v34
			}
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v44 = l0 + int32(16)
			if v42 == v44 {
				v46 = F_GetMemoryChunkContext(m, l0)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v50 = F_MemoryContextAlloc(m, v46, v41<<(uint(int32(2))%32))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v50
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v55 = v53 << (uint(int32(2)) % 32)
						if v55 == int32(0) {
						} else {
							base.MemoryCopy(m, v50, v44, v55)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v41
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v68 = v67
						if l1 < v68 {
							v75 = (v68 - l1) << (uint(int32(2)) % 32)
							if v75 != 0 {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v79 = v76 + l1<<(uint(int32(2))%32)
								base.MemoryCopy(m, v79+int32(4), v79, v75)
							} else {
							}
							v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v87 = v84
						} else {
							v87 = v68
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v87 + int32(1)
						v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v91+l1<<(uint(int32(2))%32)))) = l2
						return l0
					}
				}
			} else {
				v61 = F_repalloc(m, v42, v41<<(uint(int32(2))%32))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v61
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v41
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v68 = v67
					if l1 < v68 {
						v75 = (v68 - l1) << (uint(int32(2)) % 32)
						if v75 != 0 {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v79 = v76 + l1<<(uint(int32(2))%32)
							base.MemoryCopy(m, v79+int32(4), v79, v75)
						} else {
						}
						v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v87 = v84
					} else {
						v87 = v68
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v87 + int32(1)
					v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v91+l1<<(uint(int32(2))%32)))) = l2
					return l0
				}
			}
		} else {
			v68 = v24
			if l1 < v68 {
				v75 = (v68 - l1) << (uint(int32(2)) % 32)
				if v75 != 0 {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v79 = v76 + l1<<(uint(int32(2))%32)
					base.MemoryCopy(m, v79+int32(4), v79, v75)
				} else {
				}
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v87 = v84
			} else {
				v87 = v68
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v87 + int32(1)
			v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v91+l1<<(uint(int32(2))%32)))) = l2
			return l0
		}
	}
}
func F_list_truncate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v3 = int32(0)
	if base.B2i32(l0 == v3)|base.B2i32(l1 <= v3) != 0 {
		v13 = int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if l1 < v10 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
		} else {
		}
		v13 = l0
	}
	return v13
}
