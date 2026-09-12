package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int8_avg_accum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
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
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v93 int64
	_ = v93
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v104 int64
	_ = v104
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v120 int32
	_ = v120
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v131 int64
	_ = v131
	var v132 int64
	_ = v132
	var v134 int64
	_ = v134
	var v139 int32
	_ = v139
	var v142 int64
	_ = v142
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v12 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v15 != 0 {
			v65 = v15
			v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v67 == int32(0) {
				v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v71 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
				if v72 == int32(1) {
					v76 = v10 + int32(8)
					v80 = v71 * (v71 >> (uint(int64(63)) % 64))
					v83 = int64(32)
					v84 = int64(base.Ui64(v71) >> (uint(v83) % 64))
					v89 = int64(4294967295)
					v90 = v71 & v89
					v93 = v90 * v90
					v96 = v90 * v84
					v97 = int64(base.Ui64(v93)>>(uint(v83)%64)) + v96
					v104 = v96 + v97&v89
					*(*int64)(unsafe.Add(mBase, uint32(v76)+8)) = v80 + v80 + v84*v84 + int64(base.Ui64(v97)>>(uint(v83)%64)) + int64(base.Ui64(v104)>>(uint(v83)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v76))) = v93&v89 | v104<<(uint(v83)%64)
					v115 = *(*int64)(unsafe.Add(mBase, uint32(v65)+32))
					v116 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
					v117 = v115 + v116
					*(*int64)(unsafe.Add(mBase, uint32(v65)+32)) = v117
					v120 = v65 + int32(40)
					v123 = *(*int64)(unsafe.Add(mBase, uint32(v120)))
					v124 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v120))) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v117) < base.Ui64(v115))) + (v123 + v124)
				} else {
				}
				v131 = *(*int64)(unsafe.Add(mBase, uint32(v65)+16))
				v132 = v131 + v71
				*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v132
				v134 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v134 + int64(1)
				v139 = v65 + int32(24)
				v142 = *(*int64)(unsafe.Add(mBase, uint32(v139)))
				*(*int64)(unsafe.Add(mBase, uint32(v139))) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v132) < base.Ui64(v131))) + (v142 + v71>>(uint(int64(63))%64))
			} else {
			}
			m.G0 = v10 + int32(32)
			return v65
		} else {
			v18 = v10 + int32(28)
			v19 = int32(0)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v20 == v19 {
				v37 = int32(0)
				if v18 == v37 {
					v45 = v37
				} else {
					v40 = v37
					v41 = v19
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
					v45 = v41
				}
				v48 = v45
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				switch v23 - int32(429) {
				case 0:
					if v18 == int32(0) {
						v48 = int32(1)
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+168))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
						v40 = v30
						v41 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
						v45 = v41
						v48 = v45
					}
				case 1:
					if v18 == int32(0) {
						v48 = int32(2)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v20)+368))
						v40 = v35
						v41 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
						v45 = v41
						v48 = v45
					}
				default:
					v37 = int32(0)
					if v18 == v37 {
						v45 = v37
					} else {
						v40 = v37
						v41 = v19
						*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
						v45 = v41
					}
					v48 = v45
				}
			}
			if v48 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v159 = m.ExcPending
				if v159 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(61076), int32(0))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(499971), int32(5606), int32(353975))
						mBase = m.M
						v168 = m.ExcPending
						if v168 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v51 = int32(4515392)
				v52 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v54
				v57 = F_palloc0(m, int32(48))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					v61 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v61)
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v52
					v65 = v57
					v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v67 == int32(0) {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v71 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
						if v72 == int32(1) {
							v76 = v10 + int32(8)
							v80 = v71 * (v71 >> (uint(int64(63)) % 64))
							v83 = int64(32)
							v84 = int64(base.Ui64(v71) >> (uint(v83) % 64))
							v89 = int64(4294967295)
							v90 = v71 & v89
							v93 = v90 * v90
							v96 = v90 * v84
							v97 = int64(base.Ui64(v93)>>(uint(v83)%64)) + v96
							v104 = v96 + v97&v89
							*(*int64)(unsafe.Add(mBase, uint32(v76)+8)) = v80 + v80 + v84*v84 + int64(base.Ui64(v97)>>(uint(v83)%64)) + int64(base.Ui64(v104)>>(uint(v83)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v76))) = v93&v89 | v104<<(uint(v83)%64)
							v115 = *(*int64)(unsafe.Add(mBase, uint32(v65)+32))
							v116 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
							v117 = v115 + v116
							*(*int64)(unsafe.Add(mBase, uint32(v65)+32)) = v117
							v120 = v65 + int32(40)
							v123 = *(*int64)(unsafe.Add(mBase, uint32(v120)))
							v124 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v120))) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v117) < base.Ui64(v115))) + (v123 + v124)
						} else {
						}
						v131 = *(*int64)(unsafe.Add(mBase, uint32(v65)+16))
						v132 = v131 + v71
						*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v132
						v134 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v134 + int64(1)
						v139 = v65 + int32(24)
						v142 = *(*int64)(unsafe.Add(mBase, uint32(v139)))
						*(*int64)(unsafe.Add(mBase, uint32(v139))) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v132) < base.Ui64(v131))) + (v142 + v71>>(uint(int64(63))%64))
					} else {
					}
					m.G0 = v10 + int32(32)
					return v65
				}
			}
		}
	} else {
		v18 = v10 + int32(28)
		v19 = int32(0)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v20 == v19 {
			v37 = int32(0)
			if v18 == v37 {
				v45 = v37
			} else {
				v40 = v37
				v41 = v19
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
				v45 = v41
			}
			v48 = v45
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			switch v23 - int32(429) {
			case 0:
				if v18 == int32(0) {
					v48 = int32(1)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+168))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
					v40 = v30
					v41 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
					v45 = v41
					v48 = v45
				}
			case 1:
				if v18 == int32(0) {
					v48 = int32(2)
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v20)+368))
					v40 = v35
					v41 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
					v45 = v41
					v48 = v45
				}
			default:
				v37 = int32(0)
				if v18 == v37 {
					v45 = v37
				} else {
					v40 = v37
					v41 = v19
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
					v45 = v41
				}
				v48 = v45
			}
		}
		if v48 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v159 = m.ExcPending
			if v159 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(61076), int32(0))
				mBase = m.M
				v163 = m.ExcPending
				if v163 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(499971), int32(5606), int32(353975))
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v51 = int32(4515392)
			v52 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v54
			v57 = F_palloc0(m, int32(48))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				v61 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v61)
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v52
				v65 = v57
				v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v67 == int32(0) {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v71 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
					if v72 == int32(1) {
						v76 = v10 + int32(8)
						v80 = v71 * (v71 >> (uint(int64(63)) % 64))
						v83 = int64(32)
						v84 = int64(base.Ui64(v71) >> (uint(v83) % 64))
						v89 = int64(4294967295)
						v90 = v71 & v89
						v93 = v90 * v90
						v96 = v90 * v84
						v97 = int64(base.Ui64(v93)>>(uint(v83)%64)) + v96
						v104 = v96 + v97&v89
						*(*int64)(unsafe.Add(mBase, uint32(v76)+8)) = v80 + v80 + v84*v84 + int64(base.Ui64(v97)>>(uint(v83)%64)) + int64(base.Ui64(v104)>>(uint(v83)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v76))) = v93&v89 | v104<<(uint(v83)%64)
						v115 = *(*int64)(unsafe.Add(mBase, uint32(v65)+32))
						v116 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
						v117 = v115 + v116
						*(*int64)(unsafe.Add(mBase, uint32(v65)+32)) = v117
						v120 = v65 + int32(40)
						v123 = *(*int64)(unsafe.Add(mBase, uint32(v120)))
						v124 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v120))) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v117) < base.Ui64(v115))) + (v123 + v124)
					} else {
					}
					v131 = *(*int64)(unsafe.Add(mBase, uint32(v65)+16))
					v132 = v131 + v71
					*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v132
					v134 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v134 + int64(1)
					v139 = v65 + int32(24)
					v142 = *(*int64)(unsafe.Add(mBase, uint32(v139)))
					*(*int64)(unsafe.Add(mBase, uint32(v139))) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v132) < base.Ui64(v131))) + (v142 + v71>>(uint(int64(63))%64))
				} else {
				}
				m.G0 = v10 + int32(32)
				return v65
			}
		}
	}
}
func F_int8_avg_combine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v113 int64
	_ = v113
	var v116 int64
	_ = v116
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v125 int32
	_ = v125
	var v128 int64
	_ = v128
	var v133 int32
	_ = v133
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = v11 + int32(8)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 == v2 {
		v33 = int32(0)
		if v14 == v33 {
			v41 = v33
		} else {
			v36 = v33
			v37 = v2
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
			v41 = v37
		}
		v44 = v41
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		switch v19 - int32(429) {
		case 0:
			if v14 == int32(0) {
				v44 = int32(1)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+168))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
				v36 = v26
				v37 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
				v41 = v37
				v44 = v41
			}
		case 1:
			if v14 == int32(0) {
				v44 = int32(2)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+368))
				v36 = v31
				v37 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
				v41 = v37
				v44 = v41
			}
		default:
			v33 = int32(0)
			if v14 == v33 {
				v41 = v33
			} else {
				v36 = v33
				v37 = v2
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
				v41 = v37
			}
			v44 = v41
		}
	}
	if v44 != 0 {
		v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v45 == int32(0) {
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v49 = v48
		} else {
			v49 = v2
		}
		v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v50 != 0 {
			v133 = v49
			m.G0 = v11 + int32(16)
			return v133
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v51 == int32(0) {
				v133 = v49
				m.G0 = v11 + int32(16)
				return v133
			} else {
				if v49 == int32(0) {
					v56 = int32(4515392)
					v57 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v59
					v62 = v11 + int32(12)
					v63 = int32(0)
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v64 == v63 {
						v81 = int32(0)
						if v62 == v81 {
							v89 = v81
						} else {
							v84 = v81
							v85 = v63
							*(*int32)(unsafe.Add(mBase, uint32(v62))) = v84
							v89 = v85
						}
						v92 = v89
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
						switch v67 - int32(429) {
						case 0:
							if v62 == int32(0) {
								v92 = int32(1)
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v64)+168))
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
								v84 = v74
								v85 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v62))) = v84
								v89 = v85
								v92 = v89
							}
						case 1:
							if v62 == int32(0) {
								v92 = int32(2)
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+368))
								v84 = v79
								v85 = int32(2)
								*(*int32)(unsafe.Add(mBase, uint32(v62))) = v84
								v89 = v85
								v92 = v89
							}
						default:
							v81 = int32(0)
							if v62 == v81 {
								v89 = v81
							} else {
								v84 = v81
								v85 = v63
								*(*int32)(unsafe.Add(mBase, uint32(v62))) = v84
								v89 = v85
							}
							v92 = v89
						}
					}
					if v92 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v159 = m.ExcPending
						if v159 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(61076), int32(0))
							mBase = m.M
							v163 = m.ExcPending
							if v163 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(499971), int32(5606), int32(353975))
								mBase = m.M
								v168 = m.ExcPending
								if v168 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v96
						v99 = F_palloc0(m, int32(48))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int32(0)
						} else {
							v103 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v103)
							v105 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v99)+8)) = v105
							v107 = *(*int64)(unsafe.Add(mBase, uint32(v51)+16))
							v108 = *(*int64)(unsafe.Add(mBase, uint32(v51)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v99)+24)) = v108
							*(*int64)(unsafe.Add(mBase, uint32(v99)+16)) = v107
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v57
							v133 = v99
							m.G0 = v11 + int32(16)
							return v133
						}
					}
				} else {
					v113 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
					if v113 <= int64(0) {
						v133 = v49
					} else {
						v116 = *(*int64)(unsafe.Add(mBase, uint32(v49)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v116 + v113
						v119 = *(*int64)(unsafe.Add(mBase, uint32(v51)+24))
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v49)+16))
						v121 = *(*int64)(unsafe.Add(mBase, uint32(v51)+16))
						v122 = v120 + v121
						*(*int64)(unsafe.Add(mBase, uint32(v49)+16)) = v122
						v125 = v49 + int32(24)
						v128 = *(*int64)(unsafe.Add(mBase, uint32(v125)))
						*(*int64)(unsafe.Add(mBase, uint32(v125))) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v122) < base.Ui64(v120))) + (v119 + v128)
						v133 = v49
					}
					m.G0 = v11 + int32(16)
					return v133
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v146 = m.ExcPending
		if v146 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(61076), int32(0))
			mBase = m.M
			v150 = m.ExcPending
			if v150 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(499971), int32(5946), int32(373843))
				mBase = m.M
				v155 = m.ExcPending
				if v155 != 0 {
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
func F_int8_dist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = v9 - v5
	v13 = int32(0)
	if base.B2i32(base.B2i32(int64(0) < v5)^base.B2i32(v10 < v9) == v13)&base.B2i32(v10 != int64(-9223372036854775807-1)) == v13 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(402018), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(500135), int32(107), int32(75893))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v43 = v10 >> (uint(int64(63)) % 64)
		v46 = F_Int64GetDatum(m, v10^v43-v43)
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return int32(0)
		} else {
			return v46
		}
	}
}
func F_int8_to_char(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v181 int32
	_ = v181
	var v183 float64
	_ = v183
	var v186 int32
	_ = v186
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
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v197 int32
	_ = v197
	var v201 int64
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v364 int32
	_ = v364
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(48)
	return v364
L2:
	;
	if base.Ui32(int32(268435454)) <= base.Ui32(v53-int32(1)) {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	return int32(0)
L4:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v23 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = int32(4)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v28&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v41 = int32(1)
	if v23&v41 != 0 {
		v53 = int32(base.Ui32(v23)>>(uint(v41)%32)) - v41
		goto L2
	} else {
		goto L14
	}
L8:
	;
	v37 = v26
	goto L10
L9:
	;
	v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
	goto L10
L10:
	;
	if v28 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v40 = v26
	goto L13
L12:
	;
	v40 = v37
	goto L13
L13:
	;
	v53 = v40
	goto L2
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L2
L15:
	;
	v59 = F_cstring_to_text(m, int32(757269))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L3
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v65 = F_palloc0(m, v53<<(uint(int32(3))%32)|int32(5))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L3
	} else {
		goto L19
	}
L18:
	;
	v364 = v59
	goto L1
L19:
	;
	v71 = F_NUM_cache(m, v53, v14+int32(12), v19, v14+int32(11))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v73&int32(1024) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v342 = v65 + int32(4)
	F_NUM_processor(m, v71, v14+int32(12), v342, v331, int32(0), v333, v334, int32(1))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L3
	} else {
		goto L103
	}
L22:
	;
	if base.Ui64(int64(4294967296)) <= base.Ui64(v17+int64(2147483648)) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	if v73&int32(16384) != 0 {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v82 = int32(2147483647)
	goto L27
L26:
	;
	v82 = base.I32_wrap_i64(v17)
	goto L27
L27:
	;
	v83 = F_int_to_roman(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v331 = v83
	v333 = v2
	v334 = v2
	goto L21
L29:
	;
	v87 = F_int64_to_numeric(m, v17)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v73&int32(2048) != 0 {
		goto L59
	} else {
		goto L60
	}
L32:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v90 = F_numeric_out_sci(m, v87, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v92 == int32(45) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v331 = v90
	v333 = v2
	v334 = v2
	goto L21
L35:
	;
	goto L36
L36:
	;
	v95 = F_strlen(m, v90)
	mBase = m.M
	v98 = F_palloc(m, v95+int32(2))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v100 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v98))) = uint8(v100)
	v103 = v98 + int32(1)
	if (v90^v103)&int32(3) != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v331 = v98
	v333 = v2
	v334 = v2
	goto L21
L39:
	;
	goto L38
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v157)
	if v157&int32(255) == int32(0) {
		goto L39
	} else {
		goto L55
	}
L41:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	v156 = v90
	v157 = v109
	v158 = v103
	goto L40
L42:
	;
	goto L43
L43:
	;
	if v90&int32(3) != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v113 = v90
	v115 = v103
	goto L47
L45:
	;
	v127 = v90
	v129 = v103
	goto L46
L46:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v134 = int32(-2139062144)
	if (int32(16843008)-v131|v131)&v134 != v134 {
		v156 = v127
		v157 = v131
		v158 = v129
		goto L40
	} else {
		goto L51
	}
L47:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v116)
	if v116 == int32(0) {
		goto L39
	} else {
		goto L49
	}
L48:
	;
	v127 = v123
	v129 = v121
	goto L46
L49:
	;
	v120 = int32(1)
	v121 = v115 + v120
	v123 = v113 + v120
	if v123&int32(3) != 0 {
		v113 = v123
		v115 = v121
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v139 = v127
	v140 = v131
	v141 = v129
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v140
	v143 = int32(4)
	v144 = v141 + v143
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	v147 = v139 + v143
	v151 = int32(-2139062144)
	if (v145|(int32(16843008)-v145))&v151 == v151 {
		v139 = v147
		v140 = v145
		v141 = v144
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v156 = v147
	v157 = v145
	v158 = v144
	goto L40
L54:
	;
	goto L53
L55:
	;
	v165 = v156
	v167 = v158
	goto L56
L56:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)) = uint8(v168)
	v170 = int32(1)
	if v168 != 0 {
		v165 = v165 + v170
		v167 = v167 + v170
		goto L56
	} else {
		goto L58
	}
L57:
	;
	goto L39
L58:
	;
	goto L57
L59:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v183 = F_pow(m, float64(10), base.F64_convert_i32_s(v181))
	mBase = m.M
	v186 = F_Int64GetDatum(m, v17)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L3
	} else {
		goto L62
	}
L60:
	;
	v201 = v17
	goto L61
L61:
	;
	v205 = F_Int64GetDatum(m, v201)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L3
	} else {
		goto L66
	}
L62:
	;
	v190 = F_Float8GetDatum(m, v183)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L3
	} else {
		goto L63
	}
L63:
	;
	v192 = F_DirectFunctionCall1Coll(m, int32(1315), int32(0), v190)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L3
	} else {
		goto L64
	}
L64:
	;
	v194 = F_DirectFunctionCall2Coll(m, int32(1279), int32(0), v186, v192)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v194)))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v181 + v197
	v201 = v196
	goto L61
L66:
	;
	v207 = F_DirectFunctionCall1Coll(m, int32(1316), int32(0), v205)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	v211 = base.B2i32(v209 == int32(45))
	v212 = v207 + v211
	v213 = F_strlen(m, v212)
	mBase = m.M
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v214 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v218 = F_palloc(m, v213+v214+int32(2))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L3
	} else {
		goto L71
	}
L69:
	;
	v305 = v212
	goto L70
L70:
	;
	if v209 == int32(45) {
		goto L94
	} else {
		goto L95
	}
L71:
	;
	if (v212^v218)&int32(3) != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v294 = v218 + v213
	v295 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v294))) = uint8(v295)
	v301 = F__emscripten_memset_bulkmem(m, v294+int32(1), base.I32_extend8_s(int32(48)), v214)
	mBase = m.M
	goto L93
L73:
	;
	goto L72
L74:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v274))) = uint8(v273)
	if v273&int32(255) == int32(0) {
		goto L73
	} else {
		goto L89
	}
L75:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	v272 = v212
	v273 = v225
	v274 = v218
	goto L74
L76:
	;
	goto L77
L77:
	;
	if v212&int32(3) != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v229 = v212
	v231 = v218
	goto L81
L79:
	;
	v243 = v212
	v245 = v218
	goto L80
L80:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v250 = int32(-2139062144)
	if (int32(16843008)-v247|v247)&v250 != v250 {
		v272 = v243
		v273 = v247
		v274 = v245
		goto L74
	} else {
		goto L85
	}
L81:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	*(*uint8)(unsafe.Add(mBase, uint32(v231))) = uint8(v232)
	if v232 == int32(0) {
		goto L73
	} else {
		goto L83
	}
L82:
	;
	v243 = v239
	v245 = v237
	goto L80
L83:
	;
	v236 = int32(1)
	v237 = v231 + v236
	v239 = v229 + v236
	if v239&int32(3) != 0 {
		v229 = v239
		v231 = v237
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v255 = v243
	v256 = v247
	v257 = v245
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = v256
	v259 = int32(4)
	v260 = v257 + v259
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v263 = v255 + v259
	v267 = int32(-2139062144)
	if (v261|(int32(16843008)-v261))&v267 == v267 {
		v255 = v263
		v256 = v261
		v257 = v260
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v272 = v263
	v273 = v261
	v274 = v260
	goto L74
L88:
	;
	goto L87
L89:
	;
	v281 = v272
	v283 = v274
	goto L90
L90:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v283)+1)) = uint8(v284)
	v286 = int32(1)
	if v284 != 0 {
		v281 = v281 + v286
		v283 = v283 + v286
		goto L90
	} else {
		goto L92
	}
L91:
	;
	goto L73
L92:
	;
	goto L91
L93:
	;
	v303 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v294+v214)+1)) = uint8(v303)
	v305 = v218
	goto L70
L94:
	;
	v309 = int32(45)
	goto L96
L95:
	;
	v309 = int32(43)
	goto L96
L96:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v213 < v310 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v331 = v305
	v333 = v310 - v213
	v334 = v309
	goto L21
L98:
	;
	goto L99
L99:
	;
	v313 = int32(0)
	if v213 <= v310 {
		v331 = v305
		v333 = v313
		v334 = v309
		goto L21
	} else {
		goto L100
	}
L100:
	;
	v315 = v214 + v310
	v318 = F_palloc(m, v315+int32(2))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L3
	} else {
		goto L101
	}
L101:
	;
	v322 = v315 + int32(1)
	v324 = F__emscripten_memset_bulkmem(m, v318, base.I32_extend8_s(int32(35)), v322)
	mBase = m.M
	goto L102
L102:
	;
	v326 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v324+v322))) = uint8(v326)
	v329 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v324+v310))) = uint8(v329)
	v331 = v318
	v333 = v313
	v334 = v309
	goto L21
L103:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)))
	if v347 == int32(1) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	F_pfree(m, v71)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L3
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v352 = F_strlen(m, v342)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v352<<(uint(int32(2))%32) + int32(16)
	v364 = v65
	goto L1
L107:
	;
	goto L106
}
