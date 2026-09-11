package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_box_area(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_box_ar(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_Float8GetDatum(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_box_closept_point(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v27 float64
	_ = v27
	var v31 float64
	_ = v31
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v39 float64
	_ = v39
	var v41 float64
	_ = v41
	var v45 float64
	_ = v45
	var v48 int32
	_ = v48
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	var v57 float64
	_ = v57
	var v58 int32
	_ = v58
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v87 float64
	_ = v87
	var v88 int32
	_ = v88
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v107 float64
	_ = v107
	var v108 float64
	_ = v108
	var v110 float64
	_ = v110
	var v116 float64
	_ = v116
	var v117 int32
	_ = v117
	var v132 int64
	_ = v132
	var v134 int64
	_ = v134
	var v136 float64
	_ = v136
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v14 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v15 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_le(v14, v15) == int32(0) {
		v39 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v11))) = v13
		v41 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
		*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v39
		*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v13
		*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v41
		v45 = F_lseg_closept_point(m, l0, v11, l2)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return float64(0)
		} else {
			v49 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
			*(*float64)(unsafe.Add(mBase, uint32(v11))) = v49
			v51 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v39
			*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v13
			*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v51
			v57 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return float64(0)
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v57)&int64(9223372036854775807)) {
					v77 = v45
				} else {
					if base.B2i32(base.F64_gt(v45, v57) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v45)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312))) != 0 {
						v77 = v45
					} else {
						if l0 != 0 {
							v73 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
							*(*int64)(unsafe.Add(mBase, uint32(l0))) = v73
							v75 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v75
						} else {
						}
						v77 = v57
					}
				}
				v78 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
				v79 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
				v80 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
				*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v80
				*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v79
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v80
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v78
				v87 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return float64(0)
				} else {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v87)&int64(9223372036854775807)) {
						v107 = v77
					} else {
						if base.B2i32(base.F64_gt(v77, v87) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v77)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312))) != 0 {
							v107 = v77
						} else {
							if l0 != 0 {
								v103 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
								*(*int64)(unsafe.Add(mBase, uint32(l0))) = v103
								v105 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
							} else {
							}
							v107 = v87
						}
					}
					v108 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v108
					v110 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
					*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v80
					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v79
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v110
					v116 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return float64(0)
					} else {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v116)&int64(9223372036854775807)) {
							v136 = v107
						} else {
							if base.B2i32(base.F64_gt(v107, v116) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v107)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312))) != 0 {
								v136 = v107
							} else {
								if l0 != 0 {
									v132 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v132
									v134 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v134
								} else {
								}
								v136 = v116
							}
						}
						m.G0 = v11 + int32(48)
						return v136
					}
				}
			}
		}
	} else {
		if base.F64_ge(v14, v13) == int32(0) {
			v39 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			*(*float64)(unsafe.Add(mBase, uint32(v11))) = v13
			v41 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
			*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v39
			*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v13
			*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v41
			v45 = F_lseg_closept_point(m, l0, v11, l2)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return float64(0)
			} else {
				v49 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v49
				v51 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
				*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v39
				*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v13
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v51
				v57 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return float64(0)
				} else {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v57)&int64(9223372036854775807)) {
						v77 = v45
					} else {
						if base.B2i32(base.F64_gt(v45, v57) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v45)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312))) != 0 {
							v77 = v45
						} else {
							if l0 != 0 {
								v73 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
								*(*int64)(unsafe.Add(mBase, uint32(l0))) = v73
								v75 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v75
							} else {
							}
							v77 = v57
						}
					}
					v78 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
					v79 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
					v80 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
					*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v80
					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v79
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v80
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v78
					v87 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return float64(0)
					} else {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v87)&int64(9223372036854775807)) {
							v107 = v77
						} else {
							if base.B2i32(base.F64_gt(v77, v87) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v77)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312))) != 0 {
								v107 = v77
							} else {
								if l0 != 0 {
									v103 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v103
									v105 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
								} else {
								}
								v107 = v87
							}
						}
						v108 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v108
						v110 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v80
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v79
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v110
						v116 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return float64(0)
						} else {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v116)&int64(9223372036854775807)) {
								v136 = v107
							} else {
								if base.B2i32(base.F64_gt(v107, v116) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v107)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312))) != 0 {
									v136 = v107
								} else {
									if l0 != 0 {
										v132 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
										*(*int64)(unsafe.Add(mBase, uint32(l0))) = v132
										v134 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v134
									} else {
									}
									v136 = v116
								}
							}
							m.G0 = v11 + int32(48)
							return v136
						}
					}
				}
			}
		} else {
			v22 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
			v23 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			if base.F64_le(v22, v23) == int32(0) {
				v39 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v13
				v41 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
				*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v39
				*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v13
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v41
				v45 = F_lseg_closept_point(m, l0, v11, l2)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return float64(0)
				} else {
					v49 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v49
					v51 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
					*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v39
					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v13
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v51
					v57 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return float64(0)
					} else {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v57)&int64(9223372036854775807)) {
							v77 = v45
						} else {
							if base.B2i32(base.F64_gt(v45, v57) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v45)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312))) != 0 {
								v77 = v45
							} else {
								if l0 != 0 {
									v73 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v73
									v75 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v75
								} else {
								}
								v77 = v57
							}
						}
						v78 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
						v79 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
						v80 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v80
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v79
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v80
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v78
						v87 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return float64(0)
						} else {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v87)&int64(9223372036854775807)) {
								v107 = v77
							} else {
								if base.B2i32(base.F64_gt(v77, v87) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v77)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312))) != 0 {
									v107 = v77
								} else {
									if l0 != 0 {
										v103 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
										*(*int64)(unsafe.Add(mBase, uint32(l0))) = v103
										v105 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
									} else {
									}
									v107 = v87
								}
							}
							v108 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v108
							v110 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v80
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v79
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v110
							v116 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return float64(0)
							} else {
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v116)&int64(9223372036854775807)) {
									v136 = v107
								} else {
									if base.B2i32(base.F64_gt(v107, v116) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v107)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312))) != 0 {
										v136 = v107
									} else {
										if l0 != 0 {
											v132 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
											*(*int64)(unsafe.Add(mBase, uint32(l0))) = v132
											v134 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
											*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v134
										} else {
										}
										v136 = v116
									}
								}
								m.G0 = v11 + int32(48)
								return v136
							}
						}
					}
				}
			} else {
				v27 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
				if base.F64_le(v27, v22) == int32(0) {
					v39 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v13
					v41 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
					*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v39
					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v13
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v41
					v45 = F_lseg_closept_point(m, l0, v11, l2)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return float64(0)
					} else {
						v49 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v49
						v51 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v39
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v13
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v51
						v57 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return float64(0)
						} else {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v57)&int64(9223372036854775807)) {
								v77 = v45
							} else {
								if base.B2i32(base.F64_gt(v45, v57) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v45)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312))) != 0 {
									v77 = v45
								} else {
									if l0 != 0 {
										v73 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
										*(*int64)(unsafe.Add(mBase, uint32(l0))) = v73
										v75 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v75
									} else {
									}
									v77 = v57
								}
							}
							v78 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
							v79 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							v80 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v80
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v79
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v80
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v78
							v87 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return float64(0)
							} else {
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v87)&int64(9223372036854775807)) {
									v107 = v77
								} else {
									if base.B2i32(base.F64_gt(v77, v87) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v77)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312))) != 0 {
										v107 = v77
									} else {
										if l0 != 0 {
											v103 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
											*(*int64)(unsafe.Add(mBase, uint32(l0))) = v103
											v105 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
											*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
										} else {
										}
										v107 = v87
									}
								}
								v108 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								*(*float64)(unsafe.Add(mBase, uint32(v11))) = v108
								v110 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
								*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v80
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v79
								*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v110
								v116 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return float64(0)
								} else {
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v116)&int64(9223372036854775807)) {
										v136 = v107
									} else {
										if base.B2i32(base.F64_gt(v107, v116) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v107)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312))) != 0 {
											v136 = v107
										} else {
											if l0 != 0 {
												v132 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
												*(*int64)(unsafe.Add(mBase, uint32(l0))) = v132
												v134 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v134
											} else {
											}
											v136 = v116
										}
									}
									m.G0 = v11 + int32(48)
									return v136
								}
							}
						}
					}
				} else {
					v31 = float64(0)
					if l0 == int32(0) {
						v136 = v31
					} else {
						v34 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
						*(*int64)(unsafe.Add(mBase, uint32(l0))) = v34
						v36 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v36
						v136 = v31
					}
					m.G0 = v11 + int32(48)
					return v136
				}
			}
		}
	}
}
func F_box_cn(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v13 float64
	_ = v13
	var v22 float64
	_ = v22
	var v26 float64
	_ = v26
	var v27 float64
	_ = v27
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v40 float64
	_ = v40
	var v49 float64
	_ = v49
	var v53 float64
	_ = v53
	var v54 float64
	_ = v54
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v6 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v8 = base.F64_add(v6, v7)
	if base.F64_eq(base.F64_abs(v8), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v13 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v6), v13)&base.F64_ne(base.F64_abs(v7), v13) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v26 = base.F64_mul(v8, float64(0.5))
			v27 = float64(0)
			if base.F64_eq(v26, v27)&base.F64_ne(v8, v27) != 0 {
				F_float_underflow_error(m)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(l0))) = v26
				v33 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
				v34 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
				v35 = base.F64_add(v33, v34)
				if base.F64_eq(base.F64_abs(v35), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					v40 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_ne(base.F64_abs(v33), v40)&base.F64_ne(base.F64_abs(v34), v40) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v53 = base.F64_mul(v35, float64(0.5))
						v54 = float64(0)
						if base.F64_eq(v53, v54)&base.F64_ne(v35, v54) != 0 {
							F_float_underflow_error(m)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v53
							return
						}
					}
				} else {
					v49 = base.F64_mul(v35, float64(0.5))
					if base.F64_eq(base.F64_abs(v49), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v53 = v49
						v54 = float64(0)
						if base.F64_eq(v53, v54)&base.F64_ne(v35, v54) != 0 {
							F_float_underflow_error(m)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v53
							return
						}
					}
				}
			}
		}
	} else {
		v22 = base.F64_mul(v8, float64(0.5))
		if base.F64_eq(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v26 = v22
			v27 = float64(0)
			if base.F64_eq(v26, v27)&base.F64_ne(v8, v27) != 0 {
				F_float_underflow_error(m)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(l0))) = v26
				v33 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
				v34 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
				v35 = base.F64_add(v33, v34)
				if base.F64_eq(base.F64_abs(v35), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					v40 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_ne(base.F64_abs(v33), v40)&base.F64_ne(base.F64_abs(v34), v40) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v53 = base.F64_mul(v35, float64(0.5))
						v54 = float64(0)
						if base.F64_eq(v53, v54)&base.F64_ne(v35, v54) != 0 {
							F_float_underflow_error(m)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v53
							return
						}
					}
				} else {
					v49 = base.F64_mul(v35, float64(0.5))
					if base.F64_eq(base.F64_abs(v49), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v53 = v49
						v54 = float64(0)
						if base.F64_eq(v53, v54)&base.F64_ne(v35, v54) != 0 {
							F_float_underflow_error(m)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v53
							return
						}
					}
				}
			}
		}
	}
}
func F_box_contain_pt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v12 float64
	_ = v12
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v21 float64
	_ = v21
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	if base.F64_ge(v6, v8) == v2 {
		v23 = v2
	} else {
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v5)+16))
		if base.F64_le(v12, v8) == int32(0) {
			v23 = v2
		} else {
			v16 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v5)+8))
			if base.F64_le(v16, v17) == int32(0) {
				v23 = v2
			} else {
				v21 = *(*float64)(unsafe.Add(mBase, uint32(v5)+24))
				v23 = base.F64_le(v21, v16)
			}
		}
	}
	return v23
}
func F_box_copy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	v4 = F_palloc(m, int32(32))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int64)(unsafe.Add(mBase, uint32(v4)+24)) = v8
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v4)+16)) = v10
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = v12
		v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = v14
		return v4
	}
}
func F_box_diagonal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v13 float64
	_ = v13
	var v15 float64
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc(m, int32(32))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*float64)(unsafe.Add(mBase, uint32(v3)))
		*(*float64)(unsafe.Add(mBase, uint32(v5))) = v9
		v11 = *(*float64)(unsafe.Add(mBase, uint32(v3)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v5)+8)) = v11
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v3)+16))
		*(*float64)(unsafe.Add(mBase, uint32(v5)+16)) = v13
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v3)+24))
		*(*float64)(unsafe.Add(mBase, uint32(v5)+24)) = v15
		return v5
	}
}
func F_box_div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	v8 = m.G0
	v9 = int32(32)
	v10 = v8 - v9
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_palloc(m, v9)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		F_point_div_point(m, v10+int32(16), v13, v12)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_point_div_point(m, v10, v13+int32(16), v12)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
				v28 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
				if base.Ui64(base.I64_reinterpret_f64(v28)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
					if base.F64_lt(v28, v27) != 0 {
						v40 = v27
						v41 = v28
					} else {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v27)&int64(9223372036854775807)) {
							v40 = v27
							v41 = v28
						} else {
							v40 = v28
							v41 = v27
						}
					}
				} else {
					v40 = v28
					v41 = v27
				}
				*(*float64)(unsafe.Add(mBase, uint32(v15)+16)) = v41
				*(*float64)(unsafe.Add(mBase, uint32(v15))) = v40
				v44 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
				v45 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
				if base.Ui64(base.I64_reinterpret_f64(v45)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
					if base.F64_gt(v44, v45) != 0 {
						v57 = v45
						v58 = v44
					} else {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v44)&int64(9223372036854775807)) {
							v57 = v45
							v58 = v44
						} else {
							v57 = v44
							v58 = v45
						}
					}
				} else {
					v57 = v44
					v58 = v45
				}
				*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v57
				*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v58
				m.G0 = v10 + int32(32)
				return v15
			}
		}
	}
}
func F_box_overabove(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+24))
	return base.F64_le(v3, base.F64_add(v5, float64(1e-06)))
}
func F_box_overbelow(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
	return base.F64_le(v3, base.F64_add(v5, float64(1e-06)))
}
func F_box_penalty(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 float64
	_ = v11
	var v17 float64
	_ = v17
	var v23 float64
	_ = v23
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v29 float64
	_ = v29
	var v35 float64
	_ = v35
	var v41 float64
	_ = v41
	var v43 float64
	_ = v43
	var v45 float64
	_ = v45
	var v47 float64
	_ = v47
	var v48 float64
	_ = v48
	var v59 float64
	_ = v59
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	var v76 float64
	_ = v76
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v81 float64
	_ = v81
	var v84 int32
	_ = v84
	var v85 float64
	_ = v85
	var v86 int32
	_ = v86
	var v87 float64
	_ = v87
	var v98 int32
	_ = v98
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui64(base.I64_reinterpret_f64(v11)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v17 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v17)&int64(9223372036854775807)) {
			v23 = v17
		} else {
			v23 = v11
		}
		if base.F64_lt(v11, v17) != 0 {
			v25 = v17
		} else {
			v25 = v23
		}
		v27 = v25
	} else {
		v27 = v11
	}
	*(*float64)(unsafe.Add(mBase, uint32(v9))) = v27
	v29 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui64(base.I64_reinterpret_f64(v29)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v35 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v35)&int64(9223372036854775807)) {
			v41 = v35
		} else {
			v41 = v29
		}
		if base.F64_lt(v29, v35) != 0 {
			v43 = v35
		} else {
			v43 = v41
		}
		v45 = v43
	} else {
		v45 = v29
	}
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v45
	v47 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v48 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(base.I64_reinterpret_f64(v48)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v47)&int64(9223372036854775807)) {
			v59 = v48
		} else {
			v59 = v47
		}
		if base.F64_gt(v47, v48) != 0 {
			v61 = v48
		} else {
			v61 = v59
		}
		v62 = v61
	} else {
		v62 = v47
	}
	*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v62
	v64 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v65 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(base.I64_reinterpret_f64(v65)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v64)&int64(9223372036854775807)) {
			v76 = v65
		} else {
			v76 = v64
		}
		if base.F64_gt(v64, v65) != 0 {
			v78 = v65
		} else {
			v78 = v76
		}
		v79 = v78
	} else {
		v79 = v64
	}
	*(*float64)(unsafe.Add(mBase, uint32(v9)+24)) = v79
	v81 = F_size_box(m, v9)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		return float64(0)
	} else {
		v85 = F_size_box(m, l0)
		mBase = m.M
		v86 = m.ExcPending
		if v86 != 0 {
			return float64(0)
		} else {
			v87 = base.F64_sub(v81, v85)
			if base.F64_ne(base.F64_abs(v87), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				m.G0 = v9 + int32(32)
				return v87
			} else {
				if base.F64_eq(base.F64_abs(v81), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					m.G0 = v9 + int32(32)
					return v87
				} else {
					if base.F64_eq(base.F64_abs(v85), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						m.G0 = v9 + int32(32)
						return v87
					} else {
						F_float_overflow_error(m)
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return float64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
