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
					F_errmsg_internal(m, int32(57923), int32(0))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(469336), int32(5606), int32(332322))
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
				v51 = int32(4425280)
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
				F_errmsg_internal(m, int32(57923), int32(0))
				mBase = m.M
				v163 = m.ExcPending
				if v163 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(469336), int32(5606), int32(332322))
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
			v51 = int32(4425280)
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
					v56 = int32(4425280)
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
							F_errmsg_internal(m, int32(57923), int32(0))
							mBase = m.M
							v163 = m.ExcPending
							if v163 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(469336), int32(5606), int32(332322))
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
			F_errmsg_internal(m, int32(57923), int32(0))
			mBase = m.M
			v150 = m.ExcPending
			if v150 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(469336), int32(5946), int32(351345))
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
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v237 int32
	_ = v237
	var v239 float64
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
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
	var v252 int64
	_ = v252
	var v253 int32
	_ = v253
	var v257 int64
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v520 int32
	_ = v520
	var v532 int32
	_ = v532
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
	return v532
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
	v59 = F_cstring_to_text(m, int32(706478))
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
	v532 = v59
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
	v454 = v65 + int32(4)
	F_NUM_processor(m, v71, v14+int32(12), v454, v443, int32(0), v445, v446, int32(1))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L3
	} else {
		goto L137
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
	v443 = v83
	v445 = v2
	v446 = v2
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
		goto L76
	} else {
		goto L77
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
	v443 = v90
	v445 = v2
	v446 = v2
	goto L21
L35:
	;
	goto L36
L36:
	;
	if v90&int32(3) == int32(0) {
		v118 = v90
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v154 = F_palloc(m, v151+int32(2))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L3
	} else {
		goto L54
	}
L38:
	;
	v151 = v143 - v90
	goto L37
L39:
	;
	v122 = v118
	goto L48
L40:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v102 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v151 = int32(0)
	goto L37
L42:
	;
	goto L43
L43:
	;
	v107 = v90
	goto L44
L44:
	;
	v111 = v107 + int32(1)
	if v111&int32(3) == int32(0) {
		v118 = v111
		goto L39
	} else {
		goto L46
	}
L45:
	;
	v143 = v111
	goto L38
L46:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	if v116 != 0 {
		v107 = v111
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v131 = int32(-2139062144)
	if (int32(16843008)-v128|v128)&v131 == v131 {
		v122 = v122 + int32(4)
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v137 = v122
	goto L51
L50:
	;
	goto L49
L51:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v141 != 0 {
		v137 = v137 + int32(1)
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v143 = v137
	goto L38
L53:
	;
	goto L52
L54:
	;
	v156 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v156)
	v159 = v154 + int32(1)
	if (v90^v159)&int32(3) != 0 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v443 = v154
	v445 = v2
	v446 = v2
	goto L21
L56:
	;
	goto L55
L57:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v214))) = uint8(v213)
	if v213&int32(255) == int32(0) {
		goto L56
	} else {
		goto L72
	}
L58:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	v212 = v90
	v213 = v165
	v214 = v159
	goto L57
L59:
	;
	goto L60
L60:
	;
	if v90&int32(3) != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v169 = v90
	v171 = v159
	goto L64
L62:
	;
	v183 = v90
	v185 = v159
	goto L63
L63:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v190 = int32(-2139062144)
	if (int32(16843008)-v187|v187)&v190 != v190 {
		v212 = v183
		v213 = v187
		v214 = v185
		goto L57
	} else {
		goto L68
	}
L64:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	*(*uint8)(unsafe.Add(mBase, uint32(v171))) = uint8(v172)
	if v172 == int32(0) {
		goto L56
	} else {
		goto L66
	}
L65:
	;
	v183 = v179
	v185 = v177
	goto L63
L66:
	;
	v176 = int32(1)
	v177 = v171 + v176
	v179 = v169 + v176
	if v179&int32(3) != 0 {
		v169 = v179
		v171 = v177
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v195 = v183
	v196 = v187
	v197 = v185
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v196
	v199 = int32(4)
	v200 = v197 + v199
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v203 = v195 + v199
	v207 = int32(-2139062144)
	if (v201|(int32(16843008)-v201))&v207 == v207 {
		v195 = v203
		v196 = v201
		v197 = v200
		goto L69
	} else {
		goto L71
	}
L70:
	;
	v212 = v203
	v213 = v201
	v214 = v200
	goto L57
L71:
	;
	goto L70
L72:
	;
	v221 = v212
	v223 = v214
	goto L73
L73:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)) = uint8(v224)
	v226 = int32(1)
	if v224 != 0 {
		v221 = v221 + v226
		v223 = v223 + v226
		goto L73
	} else {
		goto L75
	}
L74:
	;
	goto L56
L75:
	;
	goto L74
L76:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v239 = F_pow(m, float64(10), base.F64_convert_i32_s(v237))
	mBase = m.M
	v242 = F_Int64GetDatum(m, v17)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L3
	} else {
		goto L79
	}
L77:
	;
	v257 = v17
	goto L78
L78:
	;
	v261 = F_Int64GetDatum(m, v257)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L3
	} else {
		goto L83
	}
L79:
	;
	v246 = F_Float8GetDatum(m, v239)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L3
	} else {
		goto L80
	}
L80:
	;
	v248 = F_DirectFunctionCall1Coll(m, int32(1314), int32(0), v246)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L3
	} else {
		goto L81
	}
L81:
	;
	v250 = F_DirectFunctionCall2Coll(m, int32(1278), int32(0), v242, v248)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L3
	} else {
		goto L82
	}
L82:
	;
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v250)))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v237 + v253
	v257 = v252
	goto L78
L83:
	;
	v263 = F_DirectFunctionCall1Coll(m, int32(1315), int32(0), v261)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L3
	} else {
		goto L84
	}
L84:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	v267 = base.B2i32(v265 == int32(45))
	v268 = v263 + v267
	if v268&int32(3) == int32(0) {
		v292 = v268
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v326 != 0 {
		goto L102
	} else {
		goto L103
	}
L86:
	;
	v325 = v317 - v268
	goto L85
L87:
	;
	v296 = v292
	goto L96
L88:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	if v276 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v325 = int32(0)
	goto L85
L90:
	;
	goto L91
L91:
	;
	v281 = v268
	goto L92
L92:
	;
	v285 = v281 + int32(1)
	if v285&int32(3) == int32(0) {
		v292 = v285
		goto L87
	} else {
		goto L94
	}
L93:
	;
	v317 = v285
	goto L86
L94:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	if v290 != 0 {
		v281 = v285
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	v305 = int32(-2139062144)
	if (int32(16843008)-v302|v302)&v305 == v305 {
		v296 = v296 + int32(4)
		goto L96
	} else {
		goto L98
	}
L97:
	;
	v311 = v296
	goto L99
L98:
	;
	goto L97
L99:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	if v315 != 0 {
		v311 = v311 + int32(1)
		goto L99
	} else {
		goto L101
	}
L100:
	;
	v317 = v311
	goto L86
L101:
	;
	goto L100
L102:
	;
	v330 = F_palloc(m, v325+v326+int32(2))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L3
	} else {
		goto L105
	}
L103:
	;
	v417 = v268
	goto L104
L104:
	;
	if v265 == int32(45) {
		goto L128
	} else {
		goto L129
	}
L105:
	;
	if (v268^v330)&int32(3) != 0 {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	v406 = v330 + v325
	v407 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v406))) = uint8(v407)
	v413 = F__emscripten_memset_bulkmem(m, v406+int32(1), base.I32_extend8_s(int32(48)), v326)
	mBase = m.M
	goto L127
L107:
	;
	goto L106
L108:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v386))) = uint8(v385)
	if v385&int32(255) == int32(0) {
		goto L107
	} else {
		goto L123
	}
L109:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	v384 = v268
	v385 = v337
	v386 = v330
	goto L108
L110:
	;
	goto L111
L111:
	;
	if v268&int32(3) != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v341 = v268
	v343 = v330
	goto L115
L113:
	;
	v355 = v268
	v357 = v330
	goto L114
L114:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v362 = int32(-2139062144)
	if (int32(16843008)-v359|v359)&v362 != v362 {
		v384 = v355
		v385 = v359
		v386 = v357
		goto L108
	} else {
		goto L119
	}
L115:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	*(*uint8)(unsafe.Add(mBase, uint32(v343))) = uint8(v344)
	if v344 == int32(0) {
		goto L107
	} else {
		goto L117
	}
L116:
	;
	v355 = v351
	v357 = v349
	goto L114
L117:
	;
	v348 = int32(1)
	v349 = v343 + v348
	v351 = v341 + v348
	if v351&int32(3) != 0 {
		v341 = v351
		v343 = v349
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v367 = v355
	v368 = v359
	v369 = v357
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = v368
	v371 = int32(4)
	v372 = v369 + v371
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	v375 = v367 + v371
	v379 = int32(-2139062144)
	if (v373|(int32(16843008)-v373))&v379 == v379 {
		v367 = v375
		v368 = v373
		v369 = v372
		goto L120
	} else {
		goto L122
	}
L121:
	;
	v384 = v375
	v385 = v373
	v386 = v372
	goto L108
L122:
	;
	goto L121
L123:
	;
	v393 = v384
	v395 = v386
	goto L124
L124:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+1)) = uint8(v396)
	v398 = int32(1)
	if v396 != 0 {
		v393 = v393 + v398
		v395 = v395 + v398
		goto L124
	} else {
		goto L126
	}
L125:
	;
	goto L107
L126:
	;
	goto L125
L127:
	;
	v415 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v406+v326)+1)) = uint8(v415)
	v417 = v330
	goto L104
L128:
	;
	v421 = int32(45)
	goto L130
L129:
	;
	v421 = int32(43)
	goto L130
L130:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v325 < v422 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v443 = v417
	v445 = v422 - v325
	v446 = v421
	goto L21
L132:
	;
	goto L133
L133:
	;
	v425 = int32(0)
	if v325 <= v422 {
		v443 = v417
		v445 = v425
		v446 = v421
		goto L21
	} else {
		goto L134
	}
L134:
	;
	v427 = v326 + v422
	v430 = F_palloc(m, v427+int32(2))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L3
	} else {
		goto L135
	}
L135:
	;
	v434 = v427 + int32(1)
	v436 = F__emscripten_memset_bulkmem(m, v430, base.I32_extend8_s(int32(35)), v434)
	mBase = m.M
	goto L136
L136:
	;
	v438 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v436+v434))) = uint8(v438)
	v441 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v436+v422))) = uint8(v441)
	v443 = v430
	v445 = v425
	v446 = v421
	goto L21
L137:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)))
	if v459 == int32(1) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	F_pfree(m, v71)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L3
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	if v454&int32(3) == int32(0) {
		v487 = v454
		goto L144
	} else {
		goto L145
	}
L141:
	;
	goto L140
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v520<<(uint(int32(2))%32) + int32(16)
	v532 = v65
	goto L1
L143:
	;
	v520 = v512 - v454
	goto L142
L144:
	;
	v491 = v487
	goto L153
L145:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454))))
	if v471 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v520 = int32(0)
	goto L142
L147:
	;
	goto L148
L148:
	;
	v476 = v454
	goto L149
L149:
	;
	v480 = v476 + int32(1)
	if v480&int32(3) == int32(0) {
		v487 = v480
		goto L144
	} else {
		goto L151
	}
L150:
	;
	v512 = v480
	goto L143
L151:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	if v485 != 0 {
		v476 = v480
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	v500 = int32(-2139062144)
	if (int32(16843008)-v497|v497)&v500 == v500 {
		v491 = v491 + int32(4)
		goto L153
	} else {
		goto L155
	}
L154:
	;
	v506 = v491
	goto L156
L155:
	;
	goto L154
L156:
	;
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506))))
	if v510 != 0 {
		v506 = v506 + int32(1)
		goto L156
	} else {
		goto L158
	}
L157:
	;
	v512 = v506
	goto L143
L158:
	;
	goto L157
}
