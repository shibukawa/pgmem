package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__jumbleDropTableSpaceStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int64
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v152 int64
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v199 int64
	_ = v199
	var v201 int32
	_ = v201
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != 0 {
		v5 = F_strlen(m, v4)
		mBase = m.M
		v7 = v5 + int32(1)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v13 == int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v70 = v16
		} else {
			v17 = int32(4)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if base.Ui32(v19-int32(1021)) < base.Ui32(v17) {
				v29 = v19
				v31 = v17
				v33 = l0 + int32(28)
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v29) {
						v38 = F_hash_bytes_extended(m, v18, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v18))) = v38
						v41 = int32(8)
					} else {
						v41 = v29
					}
					v43 = int32(1024) - v41
					if base.Ui32(v31) < base.Ui32(v43) {
						v45 = v31
					} else {
						v45 = v43
					}
					if v45 != 0 {
						base.MemoryCopy(m, v41+v18, v33, v45)
					} else {
					}
					v49 = v41 + v45
					v50 = v31 - v45
					if v50 != 0 {
						v29 = v49
						v31 = v50
						v33 = v45 + v33
						continue
					} else {
						break
					}
					break
				}
				v59 = v49
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19+v18))) = v13
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v59 = v53 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v59
			v70 = v59
		}
		v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(int32(1024)-v70) < base.Ui32(v7) {
			v80 = v4
			v81 = v7
			v82 = v70
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v82) {
					v91 = F_hash_bytes_extended(m, v75, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v75))) = v91
					v94 = int32(8)
				} else {
					v94 = v82
				}
				v96 = int32(1024) - v94
				if base.Ui32(v81) < base.Ui32(v96) {
					v98 = v81
				} else {
					v98 = v96
				}
				if v98 != 0 {
					base.MemoryCopy(m, v94+v75, v80, v98)
				} else {
				}
				v102 = v94 + v98
				v103 = v81 - v98
				if v103 != 0 {
					v80 = v80 + v98
					v81 = v103
					v82 = v102
					continue
				} else {
					break
				}
				break
			}
			v111 = v102
		} else {
			if v7 != 0 {
				base.MemoryCopy(m, v70+v75, v4, v7)
			} else {
			}
			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v111 = v106 + v7
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v111
	} else {
		v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v117 + int32(1)
	}
	v122 = l1 + int32(8)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v128 == int32(0) {
		v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v182 = v131
	} else {
		v132 = int32(4)
		v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if base.Ui32(v134-int32(1021)) < base.Ui32(v132) {
			v143 = v134
			v144 = v132
			v147 = l0 + int32(28)
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v143) {
					v152 = F_hash_bytes_extended(m, v133, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v133))) = v152
					v155 = int32(8)
				} else {
					v155 = v143
				}
				v157 = int32(1024) - v155
				if base.Ui32(v144) < base.Ui32(v157) {
					v159 = v144
				} else {
					v159 = v157
				}
				if v159 != 0 {
					base.MemoryCopy(m, v155+v133, v147, v159)
				} else {
				}
				v163 = v155 + v159
				v164 = v144 - v159
				if v164 != 0 {
					v143 = v163
					v144 = v164
					v147 = v159 + v147
					continue
				} else {
					break
				}
				break
			}
			v172 = v163
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v134+v133))) = v128
			v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v172 = v167 + int32(4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v172
		v182 = v172
	}
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v182 != int32(1024) {
		v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
		*(*uint8)(unsafe.Add(mBase, uint32(v182+v187))) = uint8(v191)
		v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v193 + int32(1)
	} else {
		v199 = F_hash_bytes_extended(m, v187, int32(1024), int64(0))
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, uint32(v187))) = v199
		v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
		*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)) = uint8(v201)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(9)
	}
	return
}
func F__jumbleRangeTableSample(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F__jumbleNode(m, l0, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			F__jumbleNode(m, l0, v9)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				F__jumbleNode(m, l0, v12)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F__jumbleRelabelType(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int64
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v7 = l1 + int32(8)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v13 == int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v67 = v16
		} else {
			v17 = int32(4)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if base.Ui32(v19-int32(1021)) < base.Ui32(v17) {
				v28 = v19
				v30 = v17
				v32 = l0 + int32(28)
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v28) {
						v37 = F_hash_bytes_extended(m, v18, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v18))) = v37
						v40 = int32(8)
					} else {
						v40 = v28
					}
					v42 = int32(1024) - v40
					if base.Ui32(v30) < base.Ui32(v42) {
						v44 = v30
					} else {
						v44 = v42
					}
					if v44 != 0 {
						base.MemoryCopy(m, v40+v18, v32, v44)
					} else {
					}
					v48 = v40 + v44
					v49 = v30 - v44
					if v49 != 0 {
						v28 = v48
						v30 = v49
						v32 = v44 + v32
						continue
					} else {
						break
					}
					break
				}
				v57 = v48
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19+v18))) = v13
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v57 = v52 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v57
			v67 = v57
		}
		v72 = int32(4)
		v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(v67-int32(1021)) < base.Ui32(v72) {
			v79 = v7
			v80 = v67
			v82 = v72
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v80) {
					v89 = F_hash_bytes_extended(m, v73, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v73))) = v89
					v92 = int32(8)
				} else {
					v92 = v80
				}
				v94 = int32(1024) - v92
				if base.Ui32(v82) < base.Ui32(v94) {
					v96 = v82
				} else {
					v96 = v94
				}
				if v96 != 0 {
					base.MemoryCopy(m, v92+v73, v79, v96)
				} else {
				}
				v100 = v92 + v96
				v101 = v82 - v96
				if v101 != 0 {
					v79 = v79 + v96
					v80 = v100
					v82 = v101
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v100
		} else {
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			*(*int32)(unsafe.Add(mBase, uint32(v67+v73))) = v104
			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106 + int32(4)
		}
		return
	}
}
func F__jumbleResTarget(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int64
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != 0 {
		v5 = F_strlen(m, v4)
		mBase = m.M
		v7 = v5 + int32(1)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v13 == int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v70 = v16
		} else {
			v17 = int32(4)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if base.Ui32(v19-int32(1021)) < base.Ui32(v17) {
				v29 = v19
				v31 = v17
				v33 = l0 + int32(28)
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v29) {
						v38 = F_hash_bytes_extended(m, v18, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v18))) = v38
						v41 = int32(8)
					} else {
						v41 = v29
					}
					v43 = int32(1024) - v41
					if base.Ui32(v31) < base.Ui32(v43) {
						v45 = v31
					} else {
						v45 = v43
					}
					if v45 != 0 {
						base.MemoryCopy(m, v41+v18, v33, v45)
					} else {
					}
					v49 = v41 + v45
					v50 = v31 - v45
					if v50 != 0 {
						v29 = v49
						v31 = v50
						v33 = v45 + v33
						continue
					} else {
						break
					}
					break
				}
				v59 = v49
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19+v18))) = v13
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v59 = v53 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v59
			v70 = v59
		}
		v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(int32(1024)-v70) < base.Ui32(v7) {
			v80 = v4
			v81 = v7
			v82 = v70
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v82) {
					v91 = F_hash_bytes_extended(m, v75, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v75))) = v91
					v94 = int32(8)
				} else {
					v94 = v82
				}
				v96 = int32(1024) - v94
				if base.Ui32(v81) < base.Ui32(v96) {
					v98 = v81
				} else {
					v98 = v96
				}
				if v98 != 0 {
					base.MemoryCopy(m, v94+v75, v80, v98)
				} else {
				}
				v102 = v94 + v98
				v103 = v81 - v98
				if v103 != 0 {
					v80 = v80 + v98
					v81 = v103
					v82 = v102
					continue
				} else {
					break
				}
				break
			}
			v111 = v102
		} else {
			if v7 != 0 {
				base.MemoryCopy(m, v70+v75, v4, v7)
			} else {
			}
			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v111 = v106 + v7
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v111
	} else {
		v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v117 + int32(1)
	}
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F__jumbleNode(m, l0, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		return
	} else {
		v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		F__jumbleNode(m, l0, v124)
		mBase = m.M
		v126 = m.ExcPending
		if v126 != 0 {
			return
		} else {
			return
		}
	}
}
func F__jumbleRoleSpec(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int64
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int64
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int64
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v201 int64
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	v4 = l1 + int32(4)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v10 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v64 = v13
	} else {
		v14 = int32(4)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if base.Ui32(v16-int32(1021)) < base.Ui32(v14) {
			v25 = v16
			v27 = v14
			v29 = l0 + int32(28)
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v25) {
					v34 = F_hash_bytes_extended(m, v15, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v15))) = v34
					v37 = int32(8)
				} else {
					v37 = v25
				}
				v39 = int32(1024) - v37
				if base.Ui32(v27) < base.Ui32(v39) {
					v41 = v27
				} else {
					v41 = v39
				}
				if v41 != 0 {
					base.MemoryCopy(m, v37+v15, v29, v41)
				} else {
				}
				v45 = v37 + v41
				v46 = v27 - v41
				if v46 != 0 {
					v25 = v45
					v27 = v46
					v29 = v41 + v29
					continue
				} else {
					break
				}
				break
			}
			v54 = v45
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16+v15))) = v10
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v54 = v49 + int32(4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v54
		v64 = v54
	}
	v69 = int32(4)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v64-int32(1021)) < base.Ui32(v69) {
		v76 = v4
		v77 = v64
		v79 = v69
		for {
			if base.Ui32(int32(1024)) <= base.Ui32(v77) {
				v86 = F_hash_bytes_extended(m, v70, int32(1024), int64(0))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v70))) = v86
				v89 = int32(8)
			} else {
				v89 = v77
			}
			v91 = int32(1024) - v89
			if base.Ui32(v79) < base.Ui32(v91) {
				v93 = v79
			} else {
				v93 = v91
			}
			if v93 != 0 {
				base.MemoryCopy(m, v89+v70, v76, v93)
			} else {
			}
			v97 = v89 + v93
			v98 = v79 - v93
			if v98 != 0 {
				v76 = v76 + v93
				v77 = v97
				v79 = v98
				continue
			} else {
				break
			}
			break
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v97
	} else {
		v101 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		*(*int32)(unsafe.Add(mBase, uint32(v64+v70))) = v101
		v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103 + int32(4)
	}
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v114 != 0 {
		v115 = F_strlen(m, v114)
		mBase = m.M
		v117 = v115 + int32(1)
		v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v123 == int32(0) {
			v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v180 = v126
		} else {
			v127 = int32(4)
			v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if base.Ui32(v129-int32(1021)) < base.Ui32(v127) {
				v139 = v129
				v141 = v127
				v143 = l0 + int32(28)
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v139) {
						v148 = F_hash_bytes_extended(m, v128, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v128))) = v148
						v151 = int32(8)
					} else {
						v151 = v139
					}
					v153 = int32(1024) - v151
					if base.Ui32(v141) < base.Ui32(v153) {
						v155 = v141
					} else {
						v155 = v153
					}
					if v155 != 0 {
						base.MemoryCopy(m, v151+v128, v143, v155)
					} else {
					}
					v159 = v151 + v155
					v160 = v141 - v155
					if v160 != 0 {
						v139 = v159
						v141 = v160
						v143 = v155 + v143
						continue
					} else {
						break
					}
					break
				}
				v169 = v159
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v129+v128))) = v123
				v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v169 = v163 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v169
			v180 = v169
		}
		v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(int32(1024)-v180) < base.Ui32(v117) {
			v190 = v114
			v191 = v117
			v192 = v180
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v192) {
					v201 = F_hash_bytes_extended(m, v185, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v185))) = v201
					v204 = int32(8)
				} else {
					v204 = v192
				}
				v206 = int32(1024) - v204
				if base.Ui32(v191) < base.Ui32(v206) {
					v208 = v191
				} else {
					v208 = v206
				}
				if v208 != 0 {
					base.MemoryCopy(m, v204+v185, v190, v208)
				} else {
				}
				v212 = v204 + v208
				v213 = v191 - v208
				if v213 != 0 {
					v190 = v190 + v208
					v191 = v213
					v192 = v212
					continue
				} else {
					break
				}
				break
			}
			v221 = v212
		} else {
			if v117 != 0 {
				base.MemoryCopy(m, v180+v185, v114, v117)
			} else {
			}
			v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v221 = v216 + v117
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v221
		return
	} else {
		v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v227 + int32(1)
		return
	}
}
func F__jumbleUpdateStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F__jumbleNode(m, l0, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			F__jumbleNode(m, l0, v9)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				F__jumbleNode(m, l0, v12)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					F__jumbleNode(m, l0, v15)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						F__jumbleNode(m, l0, v18)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
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
