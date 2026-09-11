package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bmsToString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_initStringInfo(m, v5)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		F_outBitmapset(m, v5, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			m.G0 = v5 + int32(16)
			return v13
		}
	}
}
func F_bms_add_members(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v166 int32
	_ = v166
	v3 = int32(0)
	if l0 == v3 {
		if l1 == int32(0) {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v23 = v19<<(uint(int32(2))%32) + int32(8)
			v24 = F_palloc(m, v23)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v23 != 0 {
					v28 = F__emscripten_memcpy_bulkmem(m, v24, l1, v23)
					mBase = m.M
					v29 = v28
				} else {
					v29 = v24
				}
				return v29
			}
		}
	} else {
		if l1 == int32(0) {
			return l0
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v34 <= v35 {
				v46 = l1
				v47 = v34
				v49 = l0
				v50 = int32(8)
				v51 = v49 + v50
				v53 = v46 + v50
				v54 = int32(1)
				if v47 <= v54 {
					v57 = v54
				} else {
					v57 = v47
				}
				v59 = v57 & int32(3)
				v60 = int32(0)
				if int32(4) <= v47 {
					v67 = v60
					v71 = int32(0)
					for {
						v79 = v67 << (uint(int32(2)) % 32)
						v80 = v51 + v79
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v53)))
						*(*int32)(unsafe.Add(mBase, uint32(v80))) = v81 | v83
						v86 = int32(4)
						v87 = v79 | v86
						v88 = v51 + v87
						v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v53)))
						*(*int32)(unsafe.Add(mBase, uint32(v88))) = v89 | v91
						v95 = v79 | int32(8)
						v96 = v51 + v95
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
						v99 = *(*int32)(unsafe.Add(mBase, uint32(v95+v53)))
						*(*int32)(unsafe.Add(mBase, uint32(v96))) = v97 | v99
						v103 = v79 | int32(12)
						v104 = v51 + v103
						v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
						v107 = *(*int32)(unsafe.Add(mBase, uint32(v103+v53)))
						*(*int32)(unsafe.Add(mBase, uint32(v104))) = v105 | v107
						v111 = v67 + v86
						v113 = v71 + v86
						if v113 != v57&int32(2147483644) {
							v67 = v111
							v71 = v113
							continue
						} else {
							break
						}
						break
					}
					v116 = v111
				} else {
					v116 = v60
				}
				if v59 != 0 {
					v128 = v116
					v137 = v3
					for {
						v140 = v128 << (uint(int32(2)) % 32)
						v141 = v51 + v140
						v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
						v144 = *(*int32)(unsafe.Add(mBase, uint32(v140+v53)))
						*(*int32)(unsafe.Add(mBase, uint32(v141))) = v142 | v144
						v147 = int32(1)
						v150 = v137 + v147
						if v150 != v59 {
							v128 = v128 + v147
							v137 = v150
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				if l0 != v49 {
					F_pfree(m, l0)
					mBase = m.M
					v166 = m.ExcPending
					if v166 != 0 {
						return int32(0)
					} else {
						return v49
					}
				} else {
					return v49
				}
			} else {
				v40 = v34<<(uint(int32(2))%32) + int32(8)
				v41 = F_palloc(m, v40)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					if v40 != 0 {
						v43 = F__emscripten_memcpy_bulkmem(m, v41, l1, v40)
						mBase = m.M
					} else {
					}
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v46 = l0
					v47 = v45
					v49 = v41
					v50 = int32(8)
					v51 = v49 + v50
					v53 = v46 + v50
					v54 = int32(1)
					if v47 <= v54 {
						v57 = v54
					} else {
						v57 = v47
					}
					v59 = v57 & int32(3)
					v60 = int32(0)
					if int32(4) <= v47 {
						v67 = v60
						v71 = int32(0)
						for {
							v79 = v67 << (uint(int32(2)) % 32)
							v80 = v51 + v79
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v53)))
							*(*int32)(unsafe.Add(mBase, uint32(v80))) = v81 | v83
							v86 = int32(4)
							v87 = v79 | v86
							v88 = v51 + v87
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v53)))
							*(*int32)(unsafe.Add(mBase, uint32(v88))) = v89 | v91
							v95 = v79 | int32(8)
							v96 = v51 + v95
							v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v95+v53)))
							*(*int32)(unsafe.Add(mBase, uint32(v96))) = v97 | v99
							v103 = v79 | int32(12)
							v104 = v51 + v103
							v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v103+v53)))
							*(*int32)(unsafe.Add(mBase, uint32(v104))) = v105 | v107
							v111 = v67 + v86
							v113 = v71 + v86
							if v113 != v57&int32(2147483644) {
								v67 = v111
								v71 = v113
								continue
							} else {
								break
							}
							break
						}
						v116 = v111
					} else {
						v116 = v60
					}
					if v59 != 0 {
						v128 = v116
						v137 = v3
						for {
							v140 = v128 << (uint(int32(2)) % 32)
							v141 = v51 + v140
							v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
							v144 = *(*int32)(unsafe.Add(mBase, uint32(v140+v53)))
							*(*int32)(unsafe.Add(mBase, uint32(v141))) = v142 | v144
							v147 = int32(1)
							v150 = v137 + v147
							if v150 != v59 {
								v128 = v128 + v147
								v137 = v150
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					if l0 != v49 {
						F_pfree(m, l0)
						mBase = m.M
						v166 = m.ExcPending
						if v166 != 0 {
							return int32(0)
						} else {
							return v49
						}
					} else {
						return v49
					}
				}
			}
		}
	}
}
func F_bms_compare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	if l1 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v10 = int32(-1)
	goto L6
L5:
	;
	v10 = int32(0)
	goto L6
L6:
	;
	return v10
L7:
	;
	return int32(1)
L8:
	;
	goto L9
L9:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v16 != v17 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if v17 < v16 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v24 = int32(8)
	v31 = v16 - int32(1)
	goto L17
L13:
	;
	v22 = int32(1)
	goto L15
L14:
	;
	v22 = int32(-1)
	goto L15
L15:
	;
	return v22
L16:
	;
	if base.Ui32(v40) < base.Ui32(v38) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v36 = v31 << (uint(int32(2)) % 32)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0+v24+v36)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+(l1+v24))))
	if v38 != v40 {
		goto L16
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	v43 = v31 - int32(1)
	if int32(0) <= v43 {
		v31 = v43
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v51 = int32(1)
	goto L23
L22:
	;
	v51 = int32(-1)
	goto L23
L23:
	;
	return v51
}
func F_bms_overlap(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v3 = int32(0)
	if l0 == v3 {
		v44 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v44
L2:
	;
	if l1 == int32(0) {
		v44 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v12 < v13 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v15 = v12
	goto L6
L5:
	;
	v15 = v13
	goto L6
L6:
	;
	if v15 <= int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v18 = int32(1)
	goto L9
L8:
	;
	v18 = v15
	goto L9
L9:
	;
	v19 = int32(8)
	v24 = int32(0)
	goto L10
L10:
	;
	v31 = v24 << (uint(int32(2)) % 32)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1+v19+v31)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+(l0+v19))))
	v36 = v33 & v35
	v38 = base.B2i32(v36 != int32(0))
	if v36 != 0 {
		v44 = v38
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v44 = v38
	goto L1
L12:
	;
	v40 = v24 + int32(1)
	if v40 != v18 {
		v24 = v40
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
}
