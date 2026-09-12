package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bms_del_members(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	v3 = int32(0)
	if l0 == v3 {
		return int32(0)
	} else {
		if l1 == int32(0) {
			return l0
		} else {
			v14 = int32(8)
			v15 = l0 + v14
			v17 = l1 + v14
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v19 < v18 {
				v23 = v3
				for {
					v29 = v23 << (uint(int32(2)) % 32)
					v30 = v15 + v29
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v17)))
					*(*int32)(unsafe.Add(mBase, uint32(v30))) = v31 & (v33 ^ int32(-1))
					v39 = v23 + int32(1)
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v39 < v40 {
						v23 = v39
						continue
					} else {
						break
					}
					break
				}
				return l0
			} else {
				v45 = v3
				v46 = int32(-1)
				for {
					v51 = v45 << (uint(int32(2)) % 32)
					v52 = v15 + v51
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+v17)))
					v58 = v53 & (v55 ^ int32(-1))
					*(*int32)(unsafe.Add(mBase, uint32(v52))) = v58
					if v58 != 0 {
						v60 = v45
					} else {
						v60 = v46
					}
					v62 = v45 + int32(1)
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v62 < v63 {
						v45 = v62
						v46 = v60
						continue
					} else {
						break
					}
					break
				}
				if v60 == int32(-1) {
					F_pfree(m, l0)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v60 + int32(1)
					return l0
				}
			}
		}
	}
}
func F_bms_difference(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l1 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = v13<<(uint(int32(2))%32) + int32(8)
	v20 = F_palloc(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v27 < v13 {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	return int32(0)
L8:
	;
	if v19 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	return v25
L10:
	;
	v24 = F__emscripten_memcpy_bulkmem(m, v20, l0, v19)
	mBase = m.M
	v25 = v24
	goto L12
L11:
	;
	v25 = v20
	goto L12
L12:
	;
	goto L9
L13:
	;
	v67 = int32(8)
	v68 = l1 + v67
	v72 = v13<<(uint(int32(2))%32) + v67
	v73 = F_palloc(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L7
	} else {
		goto L22
	}
L14:
	;
	v29 = int32(1)
	if v13 <= v29 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v32 = v29
	goto L17
L16:
	;
	v32 = v13
	goto L17
L17:
	;
	v33 = int32(8)
	v39 = v3
	goto L18
L18:
	;
	v46 = v39 << (uint(int32(2)) % 32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+v33+v46)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+(l1+v33))))
	if v48&(v50^int32(-1)) != 0 {
		goto L13
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	v55 = v39 + int32(1)
	if v55 != v32 {
		v39 = v55
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	if v72 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v78 = v76 + int32(8)
	v79 = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v81 < v80 {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v75 = F__emscripten_memcpy_bulkmem(m, v73, l0, v72)
	mBase = m.M
	v76 = v75
	goto L26
L25:
	;
	v76 = v73
	goto L26
L26:
	;
	goto L23
L27:
	;
	return v76
L28:
	;
	v85 = v79
	goto L31
L29:
	;
	goto L30
L30:
	;
	v108 = v79
	v109 = int32(-1)
	goto L34
L31:
	;
	v92 = v85 << (uint(int32(2)) % 32)
	v93 = v78 + v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v68)))
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v94 & (v96 ^ int32(-1))
	v102 = v85 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v102 < v103 {
		v85 = v102
		goto L31
	} else {
		goto L33
	}
L33:
	;
	goto L27
L34:
	;
	v115 = v108 << (uint(int32(2)) % 32)
	v116 = v78 + v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115+v68)))
	v122 = v117 & (v119 ^ int32(-1))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v122
	if v122 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = v124 + int32(1)
	goto L27
L36:
	;
	v124 = v108
	goto L38
L37:
	;
	v124 = v109
	goto L38
L38:
	;
	v126 = v108 + int32(1)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v126 < v127 {
		v108 = v126
		v109 = v124
		goto L34
	} else {
		goto L39
	}
L39:
	;
	goto L35
}
func F_bms_get_singleton_member(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
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
	v12 = int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 <= v12 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = v12
	goto L6
L5:
	;
	v16 = v13
	goto L6
L6:
	;
	v21 = int32(0)
	v24 = int32(-1)
	goto L8
L7:
	;
	return v49
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(8)+v21<<(uint(int32(2))%32))))
	if v31 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v41
	v49 = int32(1)
	goto L7
L10:
	;
	if int32(0) <= v24 {
		v49 = v3
		goto L7
	} else {
		goto L13
	}
L11:
	;
	v41 = v24
	goto L12
L12:
	;
	v43 = v21 + int32(1)
	if v43 != v16 {
		v21 = v43
		v24 = v41
		goto L8
	} else {
		goto L15
	}
L13:
	;
	if base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v31)) {
		v49 = v3
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v41 = base.I32_ctz(v31) | v21<<(uint(int32(5))%32)
	goto L12
L15:
	;
	goto L9
}
func F_bms_intersect(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v137 int32
	_ = v137
	v3 = int32(0)
	if l0 == v3 {
		v137 = v3
		return v137
	} else {
		if l1 == int32(0) {
			v137 = v3
			return v137
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v19 < v20 {
				v22 = v19
			} else {
				v22 = v20
			}
			v26 = v22<<(uint(int32(2))%32) + int32(8)
			v27 = F_palloc(m, v26)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = base.B2i32(v20 < v19)
				if v20 < v19 {
					v32 = l1
				} else {
					v32 = l0
				}
				if v26 != 0 {
					v33 = F__emscripten_memcpy_bulkmem(m, v27, v32, v26)
					mBase = m.M
					v34 = v33
				} else {
					v34 = v27
				}
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
				if v35 <= int32(1) {
					v38 = int32(1)
				} else {
					v38 = v35
				}
				if v20 < v19 {
					v41 = l0
				} else {
					v41 = l1
				}
				v42 = int32(8)
				v43 = v41 + v42
				v45 = v34 + v42
				if v35 < int32(2) {
					v93 = int32(0)
					v94 = int32(-1)
				} else {
					v52 = int32(0)
					v55 = v52
					v56 = int32(-1)
					v58 = v52
					for {
						v68 = int32(2)
						v69 = v55 << (uint(v68) % 32)
						v70 = v45 + v69
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v43+v69)))
						v74 = v71 & v73
						*(*int32)(unsafe.Add(mBase, uint32(v70))) = v74
						v77 = v55 | int32(1)
						v79 = v77 << (uint(v68) % 32)
						v80 = v45 + v79
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v43+v79)))
						v84 = v81 & v83
						*(*int32)(unsafe.Add(mBase, uint32(v80))) = v84
						if v74 != 0 {
							v86 = v55
						} else {
							v86 = v56
						}
						if v84 != 0 {
							v87 = v77
						} else {
							v87 = v86
						}
						v88 = int32(2)
						v89 = v55 + v88
						v91 = v58 + v88
						if v91 != v38&int32(2147483646) {
							v55 = v89
							v56 = v87
							v58 = v91
							continue
						} else {
							break
						}
						break
					}
					v93 = v89
					v94 = v87
				}
				if v38&int32(1) != 0 {
					v107 = v93 << (uint(int32(2)) % 32)
					v108 = v45 + v107
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v43+v107)))
					v112 = v109 & v111
					*(*int32)(unsafe.Add(mBase, uint32(v108))) = v112
					if v112 != 0 {
						v114 = v93
					} else {
						v114 = v94
					}
					v115 = v114
				} else {
					v115 = v94
				}
				if v115 == int32(-1) {
					F_pfree(m, v34)
					mBase = m.M
					v122 = m.ExcPending
					if v122 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v115 + int32(1)
					v137 = v27
					return v137
				}
			}
		}
	}
}
func F_bms_member_index(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v9 = int32(-1)
	if int32(0) <= l1 {
		if l0 == int32(0) {
			v72 = v9
		} else {
			v15 = int32(base.Ui32(l1) >> (uint(int32(5)) % 32))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v16 <= v15 {
				v72 = v9
			} else {
				v19 = l0 + int32(8)
				v22 = v19 + v15<<(uint(int32(2))%32)
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				if int32(base.Ui32(v23)>>(uint(l1)%32))&int32(1) == int32(0) {
					v72 = v9
				} else {
					if base.Ui32(l1) < base.Ui32(int32(32)) {
						v57 = int32(0)
						v59 = v23
					} else {
						v36 = int32(0)
						v38 = int32(0)
						for {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v19+v36<<(uint(int32(2))%32))))
							if v47 != 0 {
								v50 = v38 + base.I32_popcnt(v47)
							} else {
								v50 = v38
							}
							v52 = v36 + int32(1)
							if v52 != v15 {
								v36 = v52
								v38 = v50
								continue
							} else {
								break
							}
							break
						}
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
						v57 = v50
						v59 = v54
					}
					v63 = int32(-1)
					v72 = v57 + base.I32_popcnt(v59&(v63<<(uint(l1&int32(31))%32)^v63))
				}
			}
		}
		return v72
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v84 = m.ExcPending
		if v84 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(439534), int32(0))
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(493245), int32(519), int32(228491))
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
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
func F_bms_union(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
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
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	v3 = int32(0)
	if l0 == v3 {
		if l1 != 0 {
			v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v137 = l1
			v138 = v136
			v142 = v138<<(uint(int32(2))%32) + int32(8)
			v143 = F_palloc(m, v142)
			mBase = m.M
			v144 = m.ExcPending
			if v144 != 0 {
				return int32(0)
			} else {
				if v142 != 0 {
					v145 = F__emscripten_memcpy_bulkmem(m, v143, v137, v142)
					mBase = m.M
				} else {
				}
				v154 = v143
				return v154
			}
		} else {
			return int32(0)
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if l1 == int32(0) {
			v137 = l0
			v138 = v16
			v142 = v138<<(uint(int32(2))%32) + int32(8)
			v143 = F_palloc(m, v142)
			mBase = m.M
			v144 = m.ExcPending
			if v144 != 0 {
				return int32(0)
			} else {
				if v142 != 0 {
					v145 = F__emscripten_memcpy_bulkmem(m, v143, v137, v142)
					mBase = m.M
				} else {
				}
				v154 = v143
				return v154
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v20 = base.B2i32(v19 < v16)
			if v19 < v16 {
				v21 = v16
			} else {
				v21 = v19
			}
			v25 = v21<<(uint(int32(2))%32) + int32(8)
			v26 = F_palloc(m, v25)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				if v19 < v16 {
					v30 = l0
				} else {
					v30 = l1
				}
				if v25 != 0 {
					v31 = F__emscripten_memcpy_bulkmem(m, v26, v30, v25)
					mBase = m.M
					v32 = v31
				} else {
					v32 = v26
				}
				if v19 < v16 {
					v34 = l1
				} else {
					v34 = l0
				}
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
				if v35 <= int32(1) {
					v38 = int32(1)
				} else {
					v38 = v35
				}
				v40 = v38 & int32(3)
				v41 = int32(8)
				v42 = v34 + v41
				v44 = v32 + v41
				v45 = int32(0)
				if int32(4) <= v35 {
					v52 = v45
					v56 = int32(0)
					for {
						v63 = v52 << (uint(int32(2)) % 32)
						v64 = v44 + v63
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v42+v63)))
						*(*int32)(unsafe.Add(mBase, uint32(v64))) = v65 | v67
						v70 = int32(4)
						v71 = v63 | v70
						v72 = v44 + v71
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v42+v71)))
						*(*int32)(unsafe.Add(mBase, uint32(v72))) = v73 | v75
						v79 = v63 | int32(8)
						v80 = v44 + v79
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v42+v79)))
						*(*int32)(unsafe.Add(mBase, uint32(v80))) = v81 | v83
						v87 = v63 | int32(12)
						v88 = v44 + v87
						v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v42+v87)))
						*(*int32)(unsafe.Add(mBase, uint32(v88))) = v89 | v91
						v95 = v52 + v70
						v97 = v56 + v70
						if v97 != v38&int32(2147483644) {
							v52 = v95
							v56 = v97
							continue
						} else {
							break
						}
						break
					}
					v100 = v95
				} else {
					v100 = v45
				}
				if v40 == int32(0) {
					v154 = v26
				} else {
					v113 = v100
					v121 = v3
					for {
						v124 = v113 << (uint(int32(2)) % 32)
						v125 = v44 + v124
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
						v128 = *(*int32)(unsafe.Add(mBase, uint32(v42+v124)))
						*(*int32)(unsafe.Add(mBase, uint32(v125))) = v126 | v128
						v131 = int32(1)
						v134 = v121 + v131
						if v134 != v40 {
							v113 = v113 + v131
							v121 = v134
							continue
						} else {
							break
						}
						break
					}
					v154 = v26
				}
				return v154
			}
		}
	}
}
