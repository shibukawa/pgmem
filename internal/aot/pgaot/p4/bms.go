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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
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
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
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
				if v23 == int32(0) {
					v171 = v24
					return v171
				} else {
					base.MemoryCopy(m, v24, l1, v23)
					return v24
				}
			}
		}
	} else {
		if l1 == int32(0) {
			return l0
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v35 <= v36 {
				v46 = l1
				v47 = v35
				v48 = l0
				v49 = int32(8)
				v50 = v48 + v49
				v52 = v46 + v49
				v53 = int32(1)
				if v47 <= v53 {
					v56 = v53
				} else {
					v56 = v47
				}
				v58 = v56 & int32(3)
				v59 = int32(0)
				if int32(4) <= v47 {
					v66 = v59
					v70 = int32(0)
					for {
						v78 = v66 << (uint(int32(2)) % 32)
						v79 = v50 + v78
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v52)))
						*(*int32)(unsafe.Add(mBase, uint32(v79))) = v80 | v82
						v85 = int32(4)
						v86 = v78 | v85
						v87 = v50 + v86
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+v52)))
						*(*int32)(unsafe.Add(mBase, uint32(v87))) = v88 | v90
						v94 = v78 | int32(8)
						v95 = v50 + v94
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v52)))
						*(*int32)(unsafe.Add(mBase, uint32(v95))) = v96 | v98
						v102 = v78 | int32(12)
						v103 = v50 + v102
						v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
						v106 = *(*int32)(unsafe.Add(mBase, uint32(v102+v52)))
						*(*int32)(unsafe.Add(mBase, uint32(v103))) = v104 | v106
						v110 = v66 + v85
						v112 = v70 + v85
						if v112 != v56&int32(2147483644) {
							v66 = v110
							v70 = v112
							continue
						} else {
							break
						}
						break
					}
					if v58 == int32(0) {
					} else {
						v117 = v110
						v129 = v117
						v139 = v3
						for {
							v141 = v129 << (uint(int32(2)) % 32)
							v142 = v50 + v141
							v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
							v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v52)))
							*(*int32)(unsafe.Add(mBase, uint32(v142))) = v143 | v145
							v148 = int32(1)
							v151 = v139 + v148
							if v151 != v58 {
								v129 = v129 + v148
								v139 = v151
								continue
							} else {
								break
							}
							break
						}
					}
				} else {
					v117 = v59
					v129 = v117
					v139 = v3
					for {
						v141 = v129 << (uint(int32(2)) % 32)
						v142 = v50 + v141
						v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
						v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v52)))
						*(*int32)(unsafe.Add(mBase, uint32(v142))) = v143 | v145
						v148 = int32(1)
						v151 = v139 + v148
						if v151 != v58 {
							v129 = v129 + v148
							v139 = v151
							continue
						} else {
							break
						}
						break
					}
				}
				if l0 == v48 {
					v171 = v48
					return v171
				} else {
					F_pfree(m, l0)
					mBase = m.M
					v167 = m.ExcPending
					if v167 != 0 {
						return int32(0)
					} else {
						v171 = v48
						return v171
					}
				}
			} else {
				v41 = v35<<(uint(int32(2))%32) + int32(8)
				v42 = F_palloc(m, v41)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					if v41 != 0 {
						base.MemoryCopy(m, v42, l1, v41)
					} else {
					}
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v46 = l0
					v47 = v45
					v48 = v42
					v49 = int32(8)
					v50 = v48 + v49
					v52 = v46 + v49
					v53 = int32(1)
					if v47 <= v53 {
						v56 = v53
					} else {
						v56 = v47
					}
					v58 = v56 & int32(3)
					v59 = int32(0)
					if int32(4) <= v47 {
						v66 = v59
						v70 = int32(0)
						for {
							v78 = v66 << (uint(int32(2)) % 32)
							v79 = v50 + v78
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v52)))
							*(*int32)(unsafe.Add(mBase, uint32(v79))) = v80 | v82
							v85 = int32(4)
							v86 = v78 | v85
							v87 = v50 + v86
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+v52)))
							*(*int32)(unsafe.Add(mBase, uint32(v87))) = v88 | v90
							v94 = v78 | int32(8)
							v95 = v50 + v94
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v52)))
							*(*int32)(unsafe.Add(mBase, uint32(v95))) = v96 | v98
							v102 = v78 | int32(12)
							v103 = v50 + v102
							v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v102+v52)))
							*(*int32)(unsafe.Add(mBase, uint32(v103))) = v104 | v106
							v110 = v66 + v85
							v112 = v70 + v85
							if v112 != v56&int32(2147483644) {
								v66 = v110
								v70 = v112
								continue
							} else {
								break
							}
							break
						}
						if v58 == int32(0) {
						} else {
							v117 = v110
							v129 = v117
							v139 = v3
							for {
								v141 = v129 << (uint(int32(2)) % 32)
								v142 = v50 + v141
								v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
								v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v52)))
								*(*int32)(unsafe.Add(mBase, uint32(v142))) = v143 | v145
								v148 = int32(1)
								v151 = v139 + v148
								if v151 != v58 {
									v129 = v129 + v148
									v139 = v151
									continue
								} else {
									break
								}
								break
							}
						}
					} else {
						v117 = v59
						v129 = v117
						v139 = v3
						for {
							v141 = v129 << (uint(int32(2)) % 32)
							v142 = v50 + v141
							v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
							v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v52)))
							*(*int32)(unsafe.Add(mBase, uint32(v142))) = v143 | v145
							v148 = int32(1)
							v151 = v139 + v148
							if v151 != v58 {
								v129 = v129 + v148
								v139 = v151
								continue
							} else {
								break
							}
							break
						}
					}
					if l0 == v48 {
						v171 = v48
						return v171
					} else {
						F_pfree(m, l0)
						mBase = m.M
						v167 = m.ExcPending
						if v167 != 0 {
							return int32(0)
						} else {
							v171 = v48
							return v171
						}
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	v3 = int32(0)
	if base.B2i32(l0 == v3)|base.B2i32(l1 == v3) != 0 {
		v48 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v48
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v13 < v14 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v16 = v13
	goto L5
L4:
	;
	v16 = v14
	goto L5
L5:
	;
	if v16 <= int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v19 = int32(1)
	goto L8
L7:
	;
	v19 = v16
	goto L8
L8:
	;
	v20 = int32(8)
	v25 = int32(0)
	goto L9
L9:
	;
	v32 = v25 << (uint(int32(2)) % 32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1+v20+v32)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0+v20+v32)))
	v37 = v34 & v36
	v39 = base.B2i32(v37 != int32(0))
	if v37 != 0 {
		v48 = v39
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v48 = v39
	goto L1
L11:
	;
	v41 = v25 + int32(1)
	if v41 != v19 {
		v25 = v41
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
}
