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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
	if l0 == int32(0) {
		if l1 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v16 = int32(8)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v19 = v17 + int32(4)
			if v19 <= v16 {
				v22 = v16
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
				*(*int32)(unsafe.Add(mBase, uint32(v36))) = v13
				v44 = v36 + int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v44
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v48 = v17 << (uint(int32(2)) % 32)
				if v48 != 0 {
					v49 = F__emscripten_memcpy_bulkmem(m, v44, v46, v48)
					mBase = m.M
				} else {
				}
				return v36
			}
		}
	} else {
		if l1 != 0 {
			v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v54 = v52 + v53
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v55 < v54 {
				v59 = int32(16)
				if v54 <= v59 {
					v62 = v59
				} else {
					v62 = v54
				}
				if v62&(v62-int32(1)) != 0 {
					v69 = int32(1) << (uint(int32(32)-base.I32_clz(v62)) % 32)
				} else {
					v69 = v62
				}
				v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v72 = l0 + int32(16)
				if v70 == v72 {
					v74 = F_GetMemoryChunkContext(m, l0)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						v78 = F_MemoryContextAlloc(m, v74, v69<<(uint(int32(2))%32))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v78
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v83 = v81 << (uint(int32(2)) % 32)
							if v83 != 0 {
								v84 = F__emscripten_memcpy_bulkmem(m, v78, v72, v83)
								mBase = m.M
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v69
							v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v95 = v94
							v96 = v93
							v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v99 = int32(2)
							v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
							v104 = v95 << (uint(v99) % 32)
							if v104 != 0 {
								v105 = F__emscripten_memcpy_bulkmem(m, v98+v96<<(uint(v99)%32), v102, v104)
								mBase = m.M
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v54
							return l0
						}
					}
				} else {
					v88 = F_repalloc(m, v70, v69<<(uint(int32(2))%32))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v88
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v69
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v95 = v94
						v96 = v93
						v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v99 = int32(2)
						v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v104 = v95 << (uint(v99) % 32)
						if v104 != 0 {
							v105 = F__emscripten_memcpy_bulkmem(m, v98+v96<<(uint(v99)%32), v102, v104)
							mBase = m.M
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v54
						return l0
					}
				}
			} else {
				v95 = v52
				v96 = v53
				v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v99 = int32(2)
				v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v104 = v95 << (uint(v99) % 32)
				if v104 != 0 {
					v105 = F__emscripten_memcpy_bulkmem(m, v98+v96<<(uint(v99)%32), v102, v104)
					mBase = m.M
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v54
				return l0
			}
		} else {
			return l0
		}
	}
}
func F_list_copy_head(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v3 = int32(0)
	if l0 == v3 {
		v51 = v3
		return v51
	} else {
		if l1 <= int32(0) {
			v51 = v3
			return v51
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v14 < l1 {
				v16 = v14
			} else {
				v16 = l1
			}
			v18 = v16 + int32(4)
			if v18 <= int32(8) {
				v21 = int32(8)
			} else {
				v21 = v18
			}
			if v21&(v21-int32(1)) != 0 {
				v28 = int32(1) << (uint(int32(32)-base.I32_clz(v21)) % 32)
			} else {
				v28 = v21
			}
			v30 = v28 - int32(4)
			v35 = F_palloc(m, v30<<(uint(int32(2))%32)+int32(16))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v30
				*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v16
				*(*int32)(unsafe.Add(mBase, uint32(v35))) = v10
				v43 = v35 + int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v43
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v47 = v16 << (uint(int32(2)) % 32)
				if v47 != 0 {
					v48 = F__emscripten_memcpy_bulkmem(m, v43, v45, v47)
					mBase = m.M
				} else {
				}
				v51 = v35
				return v51
			}
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
	var v22 int32
	_ = v22
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
			v22 = l0
			return v22
		}
	} else {
		v22 = int32(0)
		return v22
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
	var v23 int32
	_ = v23
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
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v199 int32
	_ = v199
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
	v23 = v3
	goto L9
L9:
	;
	v29 = v19 + v23<<(uint(int32(2))%32)
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
	v33 = v23 + int32(1)
	if v18 != v33 {
		v23 = v33
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
	v49 = v29 + int32(4)
	v54 = (v12 + (v23 ^ int32(-1))) << (uint(int32(2)) % 32)
	if v29 == v49 {
		goto L25
	} else {
		goto L26
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
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v199 - int32(1)
	goto L4
L25:
	;
	goto L24
L26:
	;
	v58 = v29 + v54
	if base.Ui32(v49-v58) <= base.Ui32(int32(0)-v54<<(uint(int32(1))%32)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v65 = F___memcpy(m, v29, v49, v54)
	mBase = m.M
	goto L24
L28:
	;
	goto L29
L29:
	;
	v68 = (v29 ^ v49) & int32(3)
	if base.Ui32(v29) < base.Ui32(v49) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	if v170 == int32(0) {
		goto L25
	} else {
		goto L66
	}
L31:
	;
	if base.Ui32(v148) <= base.Ui32(int32(3)) {
		v169 = v147
		v170 = v148
		v171 = v149
		goto L30
	} else {
		goto L62
	}
L32:
	;
	if v68 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	if v68 != 0 {
		v130 = v54
		goto L45
	} else {
		goto L46
	}
L35:
	;
	v169 = v49
	v170 = v54
	v171 = v29
	goto L30
L36:
	;
	goto L37
L37:
	;
	if v29&int32(3) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v147 = v49
	v148 = v54
	v149 = v29
	goto L31
L39:
	;
	goto L40
L40:
	;
	v75 = v49
	v76 = v54
	v77 = v29
	goto L41
L41:
	;
	if v76 == int32(0) {
		goto L25
	} else {
		goto L43
	}
L42:
	;
	v147 = v84
	v148 = v86
	v149 = v88
	goto L31
L43:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v81)
	v83 = int32(1)
	v84 = v75 + v83
	v86 = v76 - v83
	v88 = v77 + v83
	if v88&int32(3) != 0 {
		v75 = v84
		v76 = v86
		v77 = v88
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	if v130 == int32(0) {
		goto L25
	} else {
		goto L58
	}
L46:
	;
	if v58&int32(3) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v95 = v54
	goto L50
L48:
	;
	v110 = v54
	goto L49
L49:
	;
	if base.Ui32(v110) <= base.Ui32(int32(3)) {
		v130 = v110
		goto L45
	} else {
		goto L54
	}
L50:
	;
	if v95 == int32(0) {
		goto L25
	} else {
		goto L52
	}
L51:
	;
	v110 = v101
	goto L49
L52:
	;
	v101 = v95 - int32(1)
	v102 = v29 + v101
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+v101))))
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v104)
	if v102&int32(3) != 0 {
		v95 = v101
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v117 = v110
	goto L55
L55:
	;
	v121 = v117 - int32(4)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v49+v121)))
	*(*int32)(unsafe.Add(mBase, uint32(v29+v121))) = v124
	if base.Ui32(int32(3)) < base.Ui32(v121) {
		v117 = v121
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v130 = v121
	goto L45
L57:
	;
	goto L56
L58:
	;
	v137 = v130
	goto L59
L59:
	;
	v141 = v137 - int32(1)
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+v141))))
	*(*uint8)(unsafe.Add(mBase, uint32(v29+v141))) = uint8(v144)
	if v141 != 0 {
		v137 = v141
		goto L59
	} else {
		goto L61
	}
L60:
	;
	goto L25
L61:
	;
	goto L60
L62:
	;
	v154 = v147
	v155 = v148
	v156 = v149
	goto L63
L63:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v158
	v160 = int32(4)
	v161 = v154 + v160
	v163 = v156 + v160
	v165 = v155 - v160
	if base.Ui32(int32(3)) < base.Ui32(v165) {
		v154 = v161
		v155 = v165
		v156 = v163
		goto L63
	} else {
		goto L65
	}
L64:
	;
	v169 = v161
	v170 = v165
	v171 = v163
	goto L30
L65:
	;
	goto L64
L66:
	;
	v176 = v169
	v177 = v170
	v178 = v171
	goto L67
L67:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	*(*uint8)(unsafe.Add(mBase, uint32(v178))) = uint8(v180)
	v182 = int32(1)
	v187 = v177 - v182
	if v187 != 0 {
		v176 = v176 + v182
		v177 = v187
		v178 = v178 + v182
		goto L67
	} else {
		goto L69
	}
L68:
	;
	goto L25
L69:
	;
	goto L68
}
func F_list_insert_nth(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
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
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = F_palloc(m, int32(32))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v24 <= v23 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(4294967297)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v10 + int32(16)
	return v10
L6:
	;
	v26 = int32(1)
	v28 = int32(16)
	v30 = v23 + v26
	if v30 <= v28 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v65 = v23
	goto L8
L8:
	;
	if l1 < v65 {
		goto L26
	} else {
		goto L27
	}
L9:
	;
	v33 = v28
	goto L11
L10:
	;
	v33 = v30
	goto L11
L11:
	;
	if v33&(v33-int32(1)) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v40 = v26 << (uint(int32(32)-base.I32_clz(v33)) % 32)
	goto L14
L13:
	;
	v40 = v33
	goto L14
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v43 = l0 + int32(16)
	if v41 == v43 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v40
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v65 = v64
	goto L8
L16:
	;
	v45 = F_GetMemoryChunkContext(m, l0)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v59 = F_repalloc(m, v41, v40<<(uint(int32(2))%32))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L25
	}
L19:
	;
	v49 = F_MemoryContextAlloc(m, v45, v40<<(uint(int32(2))%32))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v49
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v54 = v52 << (uint(int32(2)) % 32)
	if v54 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L15
L22:
	;
	v55 = F__emscripten_memcpy_bulkmem(m, v49, v43, v54)
	mBase = m.M
	goto L24
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v59
	goto L15
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v70 = int32(2)
	v72 = v69 + l1<<(uint(v70)%32)
	v74 = v72 + int32(4)
	v77 = (v65 - l1) << (uint(v70) % 32)
	if v74 == v72 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v224 = v65
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v224 + int32(1)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v228+l1<<(uint(int32(2))%32)))) = l2
	return l0
L29:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v224 = v222
	goto L28
L30:
	;
	goto L29
L31:
	;
	v81 = v74 + v77
	if base.Ui32(v72-v81) <= base.Ui32(int32(0)-v77<<(uint(int32(1))%32)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v88 = F___memcpy(m, v74, v72, v77)
	mBase = m.M
	goto L29
L33:
	;
	goto L34
L34:
	;
	v91 = (v74 ^ v72) & int32(3)
	if base.Ui32(v74) < base.Ui32(v72) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	if v193 == int32(0) {
		goto L30
	} else {
		goto L71
	}
L36:
	;
	if base.Ui32(v171) <= base.Ui32(int32(3)) {
		v192 = v170
		v193 = v171
		v194 = v172
		goto L35
	} else {
		goto L67
	}
L37:
	;
	if v91 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	if v91 != 0 {
		v153 = v77
		goto L50
	} else {
		goto L51
	}
L40:
	;
	v192 = v72
	v193 = v77
	v194 = v74
	goto L35
L41:
	;
	goto L42
L42:
	;
	if v74&int32(3) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v170 = v72
	v171 = v77
	v172 = v74
	goto L36
L44:
	;
	goto L45
L45:
	;
	v98 = v72
	v99 = v77
	v100 = v74
	goto L46
L46:
	;
	if v99 == int32(0) {
		goto L30
	} else {
		goto L48
	}
L47:
	;
	v170 = v107
	v171 = v109
	v172 = v111
	goto L36
L48:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v104)
	v106 = int32(1)
	v107 = v98 + v106
	v109 = v99 - v106
	v111 = v100 + v106
	if v111&int32(3) != 0 {
		v98 = v107
		v99 = v109
		v100 = v111
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	if v153 == int32(0) {
		goto L30
	} else {
		goto L63
	}
L51:
	;
	if v81&int32(3) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v118 = v77
	goto L55
L53:
	;
	v133 = v77
	goto L54
L54:
	;
	if base.Ui32(v133) <= base.Ui32(int32(3)) {
		v153 = v133
		goto L50
	} else {
		goto L59
	}
L55:
	;
	if v118 == int32(0) {
		goto L30
	} else {
		goto L57
	}
L56:
	;
	v133 = v124
	goto L54
L57:
	;
	v124 = v118 - int32(1)
	v125 = v74 + v124
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v124))))
	*(*uint8)(unsafe.Add(mBase, uint32(v125))) = uint8(v127)
	if v125&int32(3) != 0 {
		v118 = v124
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v140 = v133
	goto L60
L60:
	;
	v144 = v140 - int32(4)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v72+v144)))
	*(*int32)(unsafe.Add(mBase, uint32(v74+v144))) = v147
	if base.Ui32(int32(3)) < base.Ui32(v144) {
		v140 = v144
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v153 = v144
	goto L50
L62:
	;
	goto L61
L63:
	;
	v160 = v153
	goto L64
L64:
	;
	v164 = v160 - int32(1)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v164))))
	*(*uint8)(unsafe.Add(mBase, uint32(v74+v164))) = uint8(v167)
	if v164 != 0 {
		v160 = v164
		goto L64
	} else {
		goto L66
	}
L65:
	;
	goto L30
L66:
	;
	goto L65
L67:
	;
	v177 = v170
	v178 = v171
	v179 = v172
	goto L68
L68:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	*(*int32)(unsafe.Add(mBase, uint32(v179))) = v181
	v183 = int32(4)
	v184 = v177 + v183
	v186 = v179 + v183
	v188 = v178 - v183
	if base.Ui32(int32(3)) < base.Ui32(v188) {
		v177 = v184
		v178 = v188
		v179 = v186
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v192 = v184
	v193 = v188
	v194 = v186
	goto L35
L70:
	;
	goto L69
L71:
	;
	v199 = v192
	v200 = v193
	v201 = v194
	goto L72
L72:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	*(*uint8)(unsafe.Add(mBase, uint32(v201))) = uint8(v203)
	v205 = int32(1)
	v210 = v200 - v205
	if v210 != 0 {
		v199 = v199 + v205
		v200 = v210
		v201 = v201 + v205
		goto L72
	} else {
		goto L74
	}
L73:
	;
	goto L30
L74:
	;
	goto L73
}
func F_list_truncate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v3 = int32(0)
	if l0 == v3 {
		v11 = v3
	} else {
		if l1 <= int32(0) {
			v11 = v3
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if l1 < v8 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
			} else {
			}
			v11 = l0
		}
	}
	return v11
}
