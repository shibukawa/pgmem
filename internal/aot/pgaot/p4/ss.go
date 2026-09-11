package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SS_charge_for_initplans(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 float64
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 float64
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v44 int32
	_ = v44
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v48 float64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 float64
	_ = v69
	var v75 int32
	_ = v75
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v93 float64
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 float64
	_ = v118
	var v121 float64
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 float64
	_ = v168
	var v171 float64
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	v3 = int32(0)
	v10 = float64(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v11 == v3 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if v14 <= int32(0) {
			v88 = v3
			v93 = v10
		} else {
			v17 = int32(1)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
			if v14 == v17 {
				v60 = int32(0)
				v64 = v3
				v69 = v10
			} else {
				v26 = int32(0)
				v30 = v3
				v31 = v3
				v35 = v10
				for {
					v36 = int32(2)
					v38 = v19 + v26<<(uint(v36)%32)
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
					v40 = *(*float64)(unsafe.Add(mBase, uint32(v39)+56))
					v41 = *(*float64)(unsafe.Add(mBase, uint32(v39)+64))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					v45 = *(*float64)(unsafe.Add(mBase, uint32(v44)+56))
					v46 = *(*float64)(unsafe.Add(mBase, uint32(v44)+64))
					v48 = base.F64_add(base.F64_add(v35, base.F64_add(v40, v41)), base.F64_add(v45, v46))
					v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+38)))
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+38)))
					v54 = v49&v50 ^ int32(1) | v30
					v56 = v26 + v36
					v58 = v31 + v36
					if v58 != v14&int32(2147483646) {
						v26 = v56
						v30 = v54
						v31 = v58
						v35 = v48
						continue
					} else {
						break
					}
					break
				}
				v60 = v56
				v64 = v54
				v69 = v48
			}
			if v14&v17 == int32(0) {
				v88 = v64
				v93 = v69
			} else {
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v19+v60<<(uint(int32(2))%32))))
				v76 = *(*float64)(unsafe.Add(mBase, uint32(v75)+56))
				v77 = *(*float64)(unsafe.Add(mBase, uint32(v75)+64))
				v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+38)))
				v88 = v80 ^ int32(1) | v64
				v93 = base.F64_add(v69, base.F64_add(v76, v77))
			}
		}
		v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
		if v94 == int32(0) {
		} else {
			v97 = int32(0)
			v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
			if v98 <= v97 {
			} else {
				v105 = v97
				for {
					v113 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
					v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v105<<(uint(int32(2))%32))))
					v118 = *(*float64)(unsafe.Add(mBase, uint32(v117)+48))
					*(*float64)(unsafe.Add(mBase, uint32(v117)+48)) = base.F64_add(v93, v118)
					v121 = *(*float64)(unsafe.Add(mBase, uint32(v117)+56))
					*(*float64)(unsafe.Add(mBase, uint32(v117)+56)) = base.F64_add(v93, v121)
					if v88&int32(1) != 0 {
						v124 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v117)+21)) = uint8(v124)
					} else {
					}
					v127 = v105 + int32(1)
					v128 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
					if v127 < v128 {
						v105 = v127
						continue
					} else {
						break
					}
					break
				}
			}
		}
		if v88&int32(1) != 0 {
			v142 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v142)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v142
			return
		} else {
			v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			if v146 == int32(0) {
			} else {
				v149 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
				if v149 <= int32(0) {
				} else {
					v155 = int32(0)
					for {
						v163 = *(*int32)(unsafe.Add(mBase, uint32(v146)+12))
						v167 = *(*int32)(unsafe.Add(mBase, uint32(v163+v155<<(uint(int32(2))%32))))
						v168 = *(*float64)(unsafe.Add(mBase, uint32(v167)+48))
						*(*float64)(unsafe.Add(mBase, uint32(v167)+48)) = base.F64_add(v93, v168)
						v171 = *(*float64)(unsafe.Add(mBase, uint32(v167)+56))
						*(*float64)(unsafe.Add(mBase, uint32(v167)+56)) = base.F64_add(v93, v171)
						v175 = v155 + int32(1)
						v176 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
						if v175 < v176 {
							v155 = v175
							continue
						} else {
							break
						}
						break
					}
				}
			}
			return
		}
	}
}
func F_SS_compute_initplan_cost(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 float64
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 float64
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 float64
	_ = v47
	var v48 float64
	_ = v48
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v55 float64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 float64
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v105 float64
	_ = v105
	var v108 int32
	_ = v108
	v4 = int32(0)
	v12 = float64(0)
	if l0 == v4 {
		v98 = v4
		v105 = v12
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v15 <= int32(0) {
			v98 = v4
			v105 = v12
		} else {
			v18 = int32(0)
			if v18 < v15 {
				v21 = v15
			} else {
				v21 = v18
			}
			v22 = int32(1)
			if v15 == v22 {
				v70 = int32(0)
				v71 = v4
				v78 = v12
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v34 = int32(0)
				v35 = v4
				v38 = v4
				v42 = v12
				for {
					v43 = int32(2)
					v45 = v29 + v34<<(uint(v43)%32)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
					v47 = *(*float64)(unsafe.Add(mBase, uint32(v46)+56))
					v48 = *(*float64)(unsafe.Add(mBase, uint32(v46)+64))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
					v52 = *(*float64)(unsafe.Add(mBase, uint32(v51)+56))
					v53 = *(*float64)(unsafe.Add(mBase, uint32(v51)+64))
					v55 = base.F64_add(base.F64_add(v42, base.F64_add(v47, v48)), base.F64_add(v52, v53))
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+38)))
					v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+38)))
					v61 = v56&v57 ^ int32(1) | v35
					v63 = v34 + v43
					v65 = v38 + v43
					if v65 != v21&int32(2147483646) {
						v34 = v63
						v35 = v61
						v38 = v65
						v42 = v55
						continue
					} else {
						break
					}
					break
				}
				v70 = v63
				v71 = v61
				v78 = v55
			}
			if v21&v22 == int32(0) {
				v98 = v71
				v105 = v78
			} else {
				v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v81+v70<<(uint(int32(2))%32))))
				v86 = *(*float64)(unsafe.Add(mBase, uint32(v85)+56))
				v87 = *(*float64)(unsafe.Add(mBase, uint32(v85)+64))
				v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+38)))
				v98 = v90 ^ int32(1) | v71
				v105 = base.F64_add(v78, base.F64_add(v86, v87))
			}
		}
	}
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v105
	v108 = v98 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v108)
	return
}
func F_SS_identify_outer_params(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+64))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v10 != 0 {
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
	v13 = v2
	v15 = v10
	goto L7
L5:
	;
	v122 = v2
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v122
	goto L3
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v18 == int32(0) {
		v46 = v13
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v122 = v118
	goto L6
L9:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	if v51 == int32(0) {
		v108 = v46
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v21 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v22 <= v21 {
		v46 = v13
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v26 = v21
	v27 = v13
	goto L12
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v26<<(uint(int32(2))%32))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v38 = F_bms_add_member(m, v27, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v46 = v38
	goto L9
L14:
	;
	return
L15:
	;
	v41 = v26 + int32(1)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v41 < v42 {
		v26 = v41
		v27 = v38
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v15)+344))
	if int32(0) <= v113 {
		goto L30
	} else {
		goto L31
	}
L18:
	;
	v54 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v55 <= v54 {
		v108 = v46
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v60 = v46
	v64 = v54
	goto L20
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v64<<(uint(int32(2))%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+40))
	if v70 == int32(0) {
		v97 = v60
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v108 = v97
	goto L17
L22:
	;
	v103 = v64 + int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v103 < v104 {
		v60 = v97
		v64 = v103
		goto L20
	} else {
		goto L29
	}
L23:
	;
	v73 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v74 <= v73 {
		v97 = v60
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v78 = v73
	v79 = v60
	goto L25
L25:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84+v78<<(uint(int32(2))%32))))
	v89 = F_bms_add_member(m, v79, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L14
	} else {
		goto L27
	}
L26:
	;
	v97 = v89
	goto L22
L27:
	;
	v92 = v78 + int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v92 < v93 {
		v78 = v92
		v79 = v89
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	goto L21
L30:
	;
	v116 = F_bms_add_member(m, v108, v113)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L14
	} else {
		goto L33
	}
L31:
	;
	v118 = v108
	goto L32
L32:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v119 != 0 {
		v13 = v118
		v15 = v119
		goto L7
	} else {
		goto L34
	}
L33:
	;
	v118 = v116
	goto L32
L34:
	;
	goto L8
}
