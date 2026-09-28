package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ValidateDate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	if l1|base.B2i32(l0&int32(4) == int32(0)) != 0 {
		if l0&int32(_a_F_ValidateDate_0) != 0 {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
			v45 = v40 + int32(_a_F_ValidateDate_1)
			v47 = base.I32_div_s(v45, int32(4))
			v50 = base.I32_div_s(v45, int32(-100))
			v53 = base.I32_div_s(v45, int32(400))
			v54 = v39 + v40*int32(365) + v47 + v50 + v53
			v56 = v54 + int32(_a_F_ValidateDate_2)
			v57 = int32(_a_F_ValidateDate_3)
			v58 = base.I32_div_u_s(v56, v57)
			v59 = int32(3)
			v65 = int32(2)
			v70 = base.I32_div_u_s((v58*int32(1073595727)+v56)<<(uint(v65)%32)|v59, v57)
			v73 = v54 + v58*v59 + v70 + int32(_a_F_ValidateDate_4)
			v74 = int32(1461)
			v75 = base.I32_div_u_s(v73, v74)
			v78 = v75*int32(-1461) + v73
			v80 = v78 << (uint(v65) % 32)
			if base.Ui32(v74) <= base.Ui32(v80) {
				v86 = base.I32_rem_u_s(v78+int32(305), int32(365))
				v91 = v86
			} else {
				v90 = base.I32_rem_u_s(v78+int32(306), int32(366))
				v91 = v90
			}
			v93 = base.I32_div_u_s(v80, int32(1461))
			*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v93 + v75<<(uint(int32(2))%32) - int32(_a_F_ValidateDate_5)
			v101 = v91 + int32(123)
			v105 = int32(base.Ui32(v101*int32(2141)) >> (uint(int32(16)) % 32))
			v109 = base.I32_rem_u_s(v105+int32(10), int32(12))
			*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v109 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v101 - int32(base.Ui32(v105*int32(_a_F_ValidateDate_6))>>(uint(int32(8))%32))
		} else {
		}
		if l0&int32(2) == int32(0) {
			if l0&int32(8) == int32(0) {
				v144 = int32(14)
				if l0&v144 != v144 {
					return int32(0)
				} else {
					v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
					v150 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
					if v150&int32(3) != 0 {
						v160 = int32(0)
					} else {
						v155 = base.I32_rem_s(v150, int32(100))
						if v155 != 0 {
							v160 = int32(1)
						} else {
							v157 = base.I32_rem_s(v150, int32(400))
							v160 = base.B2i32(v157 == int32(0))
						}
					}
					v163 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					v169 = *(*int32)(unsafe.Add(mBase, uint32(v160*int32(52)+v163<<(uint(int32(2))%32))+uint32(_c_F_ValidateDate[0])))
					if v148 <= v169 {
						return int32(0)
					} else {
						return int32(-2)
					}
				}
			} else {
				v137 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
				if base.Ui32(int32(-31)) <= base.Ui32(v137-int32(32)) {
					v144 = int32(14)
					if l0&v144 != v144 {
						return int32(0)
					} else {
						v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
						v150 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
						if v150&int32(3) != 0 {
							v160 = int32(0)
						} else {
							v155 = base.I32_rem_s(v150, int32(100))
							if v155 != 0 {
								v160 = int32(1)
							} else {
								v157 = base.I32_rem_s(v150, int32(400))
								v160 = base.B2i32(v157 == int32(0))
							}
						}
						v163 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						v169 = *(*int32)(unsafe.Add(mBase, uint32(v160*int32(52)+v163<<(uint(int32(2))%32))+uint32(_c_F_ValidateDate[0])))
						if v148 <= v169 {
							return int32(0)
						} else {
							return int32(-2)
						}
					}
				} else {
					return int32(-3)
				}
			}
		} else {
			v126 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
			if base.Ui32(int32(-12)) <= base.Ui32(v126-int32(13)) {
				if l0&int32(8) == int32(0) {
					v144 = int32(14)
					if l0&v144 != v144 {
						return int32(0)
					} else {
						v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
						v150 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
						if v150&int32(3) != 0 {
							v160 = int32(0)
						} else {
							v155 = base.I32_rem_s(v150, int32(100))
							if v155 != 0 {
								v160 = int32(1)
							} else {
								v157 = base.I32_rem_s(v150, int32(400))
								v160 = base.B2i32(v157 == int32(0))
							}
						}
						v163 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						v169 = *(*int32)(unsafe.Add(mBase, uint32(v160*int32(52)+v163<<(uint(int32(2))%32))+uint32(_c_F_ValidateDate[0])))
						if v148 <= v169 {
							return int32(0)
						} else {
							return int32(-2)
						}
					}
				} else {
					v137 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
					if base.Ui32(int32(-31)) <= base.Ui32(v137-int32(32)) {
						v144 = int32(14)
						if l0&v144 != v144 {
							return int32(0)
						} else {
							v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
							v150 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
							if v150&int32(3) != 0 {
								v160 = int32(0)
							} else {
								v155 = base.I32_rem_s(v150, int32(100))
								if v155 != 0 {
									v160 = int32(1)
								} else {
									v157 = base.I32_rem_s(v150, int32(400))
									v160 = base.B2i32(v157 == int32(0))
								}
							}
							v163 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
							v169 = *(*int32)(unsafe.Add(mBase, uint32(v160*int32(52)+v163<<(uint(int32(2))%32))+uint32(_c_F_ValidateDate[0])))
							if v148 <= v169 {
								return int32(0)
							} else {
								return int32(-2)
							}
						}
					} else {
						return int32(-3)
					}
				}
			} else {
				return int32(-3)
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
		if l3 != 0 {
			if int32(0) < v11 {
				v34 = int32(1) - v11
				*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v34
				if l0&int32(_a_F_ValidateDate_0) != 0 {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
					v45 = v40 + int32(_a_F_ValidateDate_1)
					v47 = base.I32_div_s(v45, int32(4))
					v50 = base.I32_div_s(v45, int32(-100))
					v53 = base.I32_div_s(v45, int32(400))
					v54 = v39 + v40*int32(365) + v47 + v50 + v53
					v56 = v54 + int32(_a_F_ValidateDate_2)
					v57 = int32(_a_F_ValidateDate_3)
					v58 = base.I32_div_u_s(v56, v57)
					v59 = int32(3)
					v65 = int32(2)
					v70 = base.I32_div_u_s((v58*int32(1073595727)+v56)<<(uint(v65)%32)|v59, v57)
					v73 = v54 + v58*v59 + v70 + int32(_a_F_ValidateDate_4)
					v74 = int32(1461)
					v75 = base.I32_div_u_s(v73, v74)
					v78 = v75*int32(-1461) + v73
					v80 = v78 << (uint(v65) % 32)
					if base.Ui32(v74) <= base.Ui32(v80) {
						v86 = base.I32_rem_u_s(v78+int32(305), int32(365))
						v91 = v86
					} else {
						v90 = base.I32_rem_u_s(v78+int32(306), int32(366))
						v91 = v90
					}
					v93 = base.I32_div_u_s(v80, int32(1461))
					*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v93 + v75<<(uint(int32(2))%32) - int32(_a_F_ValidateDate_5)
					v101 = v91 + int32(123)
					v105 = int32(base.Ui32(v101*int32(2141)) >> (uint(int32(16)) % 32))
					v109 = base.I32_rem_u_s(v105+int32(10), int32(12))
					*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v109 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v101 - int32(base.Ui32(v105*int32(_a_F_ValidateDate_6))>>(uint(int32(8))%32))
				} else {
				}
				if l0&int32(2) == int32(0) {
					if l0&int32(8) == int32(0) {
						v144 = int32(14)
						if l0&v144 != v144 {
							return int32(0)
						} else {
							v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
							v150 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
							if v150&int32(3) != 0 {
								v160 = int32(0)
							} else {
								v155 = base.I32_rem_s(v150, int32(100))
								if v155 != 0 {
									v160 = int32(1)
								} else {
									v157 = base.I32_rem_s(v150, int32(400))
									v160 = base.B2i32(v157 == int32(0))
								}
							}
							v163 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
							v169 = *(*int32)(unsafe.Add(mBase, uint32(v160*int32(52)+v163<<(uint(int32(2))%32))+uint32(_c_F_ValidateDate[0])))
							if v148 <= v169 {
								return int32(0)
							} else {
								return int32(-2)
							}
						}
					} else {
						v137 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
						if base.Ui32(int32(-31)) <= base.Ui32(v137-int32(32)) {
							v144 = int32(14)
							if l0&v144 != v144 {
								return int32(0)
							} else {
								v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
								v150 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
								if v150&int32(3) != 0 {
									v160 = int32(0)
								} else {
									v155 = base.I32_rem_s(v150, int32(100))
									if v155 != 0 {
										v160 = int32(1)
									} else {
										v157 = base.I32_rem_s(v150, int32(400))
										v160 = base.B2i32(v157 == int32(0))
									}
								}
								v163 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
								v169 = *(*int32)(unsafe.Add(mBase, uint32(v160*int32(52)+v163<<(uint(int32(2))%32))+uint32(_c_F_ValidateDate[0])))
								if v148 <= v169 {
									return int32(0)
								} else {
									return int32(-2)
								}
							}
						} else {
							return int32(-3)
						}
					}
				} else {
					v126 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					if base.Ui32(int32(-12)) <= base.Ui32(v126-int32(13)) {
						if l0&int32(8) == int32(0) {
							v144 = int32(14)
							if l0&v144 != v144 {
								return int32(0)
							} else {
								v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
								v150 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
								if v150&int32(3) != 0 {
									v160 = int32(0)
								} else {
									v155 = base.I32_rem_s(v150, int32(100))
									if v155 != 0 {
										v160 = int32(1)
									} else {
										v157 = base.I32_rem_s(v150, int32(400))
										v160 = base.B2i32(v157 == int32(0))
									}
								}
								v163 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
								v169 = *(*int32)(unsafe.Add(mBase, uint32(v160*int32(52)+v163<<(uint(int32(2))%32))+uint32(_c_F_ValidateDate[0])))
								if v148 <= v169 {
									return int32(0)
								} else {
									return int32(-2)
								}
							}
						} else {
							v137 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
							if base.Ui32(int32(-31)) <= base.Ui32(v137-int32(32)) {
								v144 = int32(14)
								if l0&v144 != v144 {
									return int32(0)
								} else {
									v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
									v150 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
									if v150&int32(3) != 0 {
										v160 = int32(0)
									} else {
										v155 = base.I32_rem_s(v150, int32(100))
										if v155 != 0 {
											v160 = int32(1)
										} else {
											v157 = base.I32_rem_s(v150, int32(400))
											v160 = base.B2i32(v157 == int32(0))
										}
									}
									v163 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
									v169 = *(*int32)(unsafe.Add(mBase, uint32(v160*int32(52)+v163<<(uint(int32(2))%32))+uint32(_c_F_ValidateDate[0])))
									if v148 <= v169 {
										return int32(0)
									} else {
										return int32(-2)
									}
								}
							} else {
								return int32(-3)
							}
						}
					} else {
						return int32(-3)
					}
				}
			} else {
				return int32(-2)
			}
		} else {
			if l2 != 0 {
				if v11 < int32(0) {
					return int32(-2)
				} else {
					if base.Ui32(v11) <= base.Ui32(int32(69)) {
						v34 = v11 + int32(2000)
						*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v34
					} else {
						if base.Ui32(int32(99)) < base.Ui32(v11) {
						} else {
							v34 = v11 + int32(1900)
							*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v34
						}
					}
					if l0&int32(_a_F_ValidateDate_0) != 0 {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
						v45 = v40 + int32(_a_F_ValidateDate_1)
						v47 = base.I32_div_s(v45, int32(4))
						v50 = base.I32_div_s(v45, int32(-100))
						v53 = base.I32_div_s(v45, int32(400))
						v54 = v39 + v40*int32(365) + v47 + v50 + v53
						v56 = v54 + int32(_a_F_ValidateDate_2)
						v57 = int32(_a_F_ValidateDate_3)
						v58 = base.I32_div_u_s(v56, v57)
						v59 = int32(3)
						v65 = int32(2)
						v70 = base.I32_div_u_s((v58*int32(1073595727)+v56)<<(uint(v65)%32)|v59, v57)
						v73 = v54 + v58*v59 + v70 + int32(_a_F_ValidateDate_4)
						v74 = int32(1461)
						v75 = base.I32_div_u_s(v73, v74)
						v78 = v75*int32(-1461) + v73
						v80 = v78 << (uint(v65) % 32)
						if base.Ui32(v74) <= base.Ui32(v80) {
							v86 = base.I32_rem_u_s(v78+int32(305), int32(365))
							v91 = v86
						} else {
							v90 = base.I32_rem_u_s(v78+int32(306), int32(366))
							v91 = v90
						}
						v93 = base.I32_div_u_s(v80, int32(1461))
						*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v93 + v75<<(uint(int32(2))%32) - int32(_a_F_ValidateDate_5)
						v101 = v91 + int32(123)
						v105 = int32(base.Ui32(v101*int32(2141)) >> (uint(int32(16)) % 32))
						v109 = base.I32_rem_u_s(v105+int32(10), int32(12))
						*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v109 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v101 - int32(base.Ui32(v105*int32(_a_F_ValidateDate_6))>>(uint(int32(8))%32))
					} else {
					}
					if l0&int32(2) == int32(0) {
						if l0&int32(8) == int32(0) {
							v144 = int32(14)
							if l0&v144 != v144 {
								return int32(0)
							} else {
								v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
								v150 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
								if v150&int32(3) != 0 {
									v160 = int32(0)
								} else {
									v155 = base.I32_rem_s(v150, int32(100))
									if v155 != 0 {
										v160 = int32(1)
									} else {
										v157 = base.I32_rem_s(v150, int32(400))
										v160 = base.B2i32(v157 == int32(0))
									}
								}
								v163 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
								v169 = *(*int32)(unsafe.Add(mBase, uint32(v160*int32(52)+v163<<(uint(int32(2))%32))+uint32(_c_F_ValidateDate[0])))
								if v148 <= v169 {
									return int32(0)
								} else {
									return int32(-2)
								}
							}
						} else {
							v137 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
							if base.Ui32(int32(-31)) <= base.Ui32(v137-int32(32)) {
								v144 = int32(14)
								if l0&v144 != v144 {
									return int32(0)
								} else {
									v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
									v150 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
									if v150&int32(3) != 0 {
										v160 = int32(0)
									} else {
										v155 = base.I32_rem_s(v150, int32(100))
										if v155 != 0 {
											v160 = int32(1)
										} else {
											v157 = base.I32_rem_s(v150, int32(400))
											v160 = base.B2i32(v157 == int32(0))
										}
									}
									v163 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
									v169 = *(*int32)(unsafe.Add(mBase, uint32(v160*int32(52)+v163<<(uint(int32(2))%32))+uint32(_c_F_ValidateDate[0])))
									if v148 <= v169 {
										return int32(0)
									} else {
										return int32(-2)
									}
								}
							} else {
								return int32(-3)
							}
						}
					} else {
						v126 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						if base.Ui32(int32(-12)) <= base.Ui32(v126-int32(13)) {
							if l0&int32(8) == int32(0) {
								v144 = int32(14)
								if l0&v144 != v144 {
									return int32(0)
								} else {
									v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
									v150 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
									if v150&int32(3) != 0 {
										v160 = int32(0)
									} else {
										v155 = base.I32_rem_s(v150, int32(100))
										if v155 != 0 {
											v160 = int32(1)
										} else {
											v157 = base.I32_rem_s(v150, int32(400))
											v160 = base.B2i32(v157 == int32(0))
										}
									}
									v163 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
									v169 = *(*int32)(unsafe.Add(mBase, uint32(v160*int32(52)+v163<<(uint(int32(2))%32))+uint32(_c_F_ValidateDate[0])))
									if v148 <= v169 {
										return int32(0)
									} else {
										return int32(-2)
									}
								}
							} else {
								v137 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
								if base.Ui32(int32(-31)) <= base.Ui32(v137-int32(32)) {
									v144 = int32(14)
									if l0&v144 != v144 {
										return int32(0)
									} else {
										v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
										v150 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
										if v150&int32(3) != 0 {
											v160 = int32(0)
										} else {
											v155 = base.I32_rem_s(v150, int32(100))
											if v155 != 0 {
												v160 = int32(1)
											} else {
												v157 = base.I32_rem_s(v150, int32(400))
												v160 = base.B2i32(v157 == int32(0))
											}
										}
										v163 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
										v169 = *(*int32)(unsafe.Add(mBase, uint32(v160*int32(52)+v163<<(uint(int32(2))%32))+uint32(_c_F_ValidateDate[0])))
										if v148 <= v169 {
											return int32(0)
										} else {
											return int32(-2)
										}
									}
								} else {
									return int32(-3)
								}
							}
						} else {
							return int32(-3)
						}
					}
				}
			} else {
				if int32(0) < v11 {
					if l0&int32(_a_F_ValidateDate_0) != 0 {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
						v45 = v40 + int32(_a_F_ValidateDate_1)
						v47 = base.I32_div_s(v45, int32(4))
						v50 = base.I32_div_s(v45, int32(-100))
						v53 = base.I32_div_s(v45, int32(400))
						v54 = v39 + v40*int32(365) + v47 + v50 + v53
						v56 = v54 + int32(_a_F_ValidateDate_2)
						v57 = int32(_a_F_ValidateDate_3)
						v58 = base.I32_div_u_s(v56, v57)
						v59 = int32(3)
						v65 = int32(2)
						v70 = base.I32_div_u_s((v58*int32(1073595727)+v56)<<(uint(v65)%32)|v59, v57)
						v73 = v54 + v58*v59 + v70 + int32(_a_F_ValidateDate_4)
						v74 = int32(1461)
						v75 = base.I32_div_u_s(v73, v74)
						v78 = v75*int32(-1461) + v73
						v80 = v78 << (uint(v65) % 32)
						if base.Ui32(v74) <= base.Ui32(v80) {
							v86 = base.I32_rem_u_s(v78+int32(305), int32(365))
							v91 = v86
						} else {
							v90 = base.I32_rem_u_s(v78+int32(306), int32(366))
							v91 = v90
						}
						v93 = base.I32_div_u_s(v80, int32(1461))
						*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v93 + v75<<(uint(int32(2))%32) - int32(_a_F_ValidateDate_5)
						v101 = v91 + int32(123)
						v105 = int32(base.Ui32(v101*int32(2141)) >> (uint(int32(16)) % 32))
						v109 = base.I32_rem_u_s(v105+int32(10), int32(12))
						*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v109 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v101 - int32(base.Ui32(v105*int32(_a_F_ValidateDate_6))>>(uint(int32(8))%32))
					} else {
					}
					if l0&int32(2) == int32(0) {
						if l0&int32(8) == int32(0) {
							v144 = int32(14)
							if l0&v144 != v144 {
								return int32(0)
							} else {
								v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
								v150 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
								if v150&int32(3) != 0 {
									v160 = int32(0)
								} else {
									v155 = base.I32_rem_s(v150, int32(100))
									if v155 != 0 {
										v160 = int32(1)
									} else {
										v157 = base.I32_rem_s(v150, int32(400))
										v160 = base.B2i32(v157 == int32(0))
									}
								}
								v163 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
								v169 = *(*int32)(unsafe.Add(mBase, uint32(v160*int32(52)+v163<<(uint(int32(2))%32))+uint32(_c_F_ValidateDate[0])))
								if v148 <= v169 {
									return int32(0)
								} else {
									return int32(-2)
								}
							}
						} else {
							v137 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
							if base.Ui32(int32(-31)) <= base.Ui32(v137-int32(32)) {
								v144 = int32(14)
								if l0&v144 != v144 {
									return int32(0)
								} else {
									v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
									v150 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
									if v150&int32(3) != 0 {
										v160 = int32(0)
									} else {
										v155 = base.I32_rem_s(v150, int32(100))
										if v155 != 0 {
											v160 = int32(1)
										} else {
											v157 = base.I32_rem_s(v150, int32(400))
											v160 = base.B2i32(v157 == int32(0))
										}
									}
									v163 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
									v169 = *(*int32)(unsafe.Add(mBase, uint32(v160*int32(52)+v163<<(uint(int32(2))%32))+uint32(_c_F_ValidateDate[0])))
									if v148 <= v169 {
										return int32(0)
									} else {
										return int32(-2)
									}
								}
							} else {
								return int32(-3)
							}
						}
					} else {
						v126 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						if base.Ui32(int32(-12)) <= base.Ui32(v126-int32(13)) {
							if l0&int32(8) == int32(0) {
								v144 = int32(14)
								if l0&v144 != v144 {
									return int32(0)
								} else {
									v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
									v150 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
									if v150&int32(3) != 0 {
										v160 = int32(0)
									} else {
										v155 = base.I32_rem_s(v150, int32(100))
										if v155 != 0 {
											v160 = int32(1)
										} else {
											v157 = base.I32_rem_s(v150, int32(400))
											v160 = base.B2i32(v157 == int32(0))
										}
									}
									v163 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
									v169 = *(*int32)(unsafe.Add(mBase, uint32(v160*int32(52)+v163<<(uint(int32(2))%32))+uint32(_c_F_ValidateDate[0])))
									if v148 <= v169 {
										return int32(0)
									} else {
										return int32(-2)
									}
								}
							} else {
								v137 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
								if base.Ui32(int32(-31)) <= base.Ui32(v137-int32(32)) {
									v144 = int32(14)
									if l0&v144 != v144 {
										return int32(0)
									} else {
										v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
										v150 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
										if v150&int32(3) != 0 {
											v160 = int32(0)
										} else {
											v155 = base.I32_rem_s(v150, int32(100))
											if v155 != 0 {
												v160 = int32(1)
											} else {
												v157 = base.I32_rem_s(v150, int32(400))
												v160 = base.B2i32(v157 == int32(0))
											}
										}
										v163 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
										v169 = *(*int32)(unsafe.Add(mBase, uint32(v160*int32(52)+v163<<(uint(int32(2))%32))+uint32(_c_F_ValidateDate[0])))
										if v148 <= v169 {
											return int32(0)
										} else {
											return int32(-2)
										}
									}
								} else {
									return int32(-3)
								}
							}
						} else {
							return int32(-3)
						}
					}
				} else {
					return int32(-2)
				}
			}
		}
	}
}
func F_date_cmp_timestamp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = F_date_cmp_timestamp_internal(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_s(v4)
	}
}
func F_date_cmp_timestamptz_internal(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_date_cmp_timestamptz_internal[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v10
	v13 = *(*int64)(unsafe.Add(mBase, _c_F_date_cmp_timestamptz_internal[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v13
	v15 = F_date2timestamptz_safe(m, l0, v7)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v19 != int32(1) {
			v39 = base.B2i32(l1 < v15) - base.B2i32(v15 < l1)
		} else {
			if v15 != int64(-9223372036854775807-1) {
				if v15 != int64(9223372036854775807) {
					v39 = base.B2i32(l1 < v15) - base.B2i32(v15 < l1)
				} else {
					if l1 == int64(9223372036854775807) {
						v30 = int32(-1)
					} else {
						v30 = int32(1)
					}
					v39 = v30
				}
			} else {
				if l1 == int64(-9223372036854775807-1) {
					v35 = int32(1)
				} else {
					v35 = int32(-1)
				}
				v39 = v35
			}
		}
		m.G0 = v7 + int32(16)
		return v39
	}
}
func F_date_finite(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	return base.I64_extend_i32_u(base.B2i32(base.Ui32(v2+int32(2147483647)) < base.Ui32(int32(-2))))
}
func F_date_ge_timestamp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = F_date_cmp_timestamp_internal(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(int32(0) <= v4))
	}
}
func F_date_gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_u(base.B2i32(v3 < v2))
}
func F_date_mii(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = base.I32_wrap_i64(v5)
	if base.Ui32(v6-int32(2147483647)) <= base.Ui32(int32(1)) {
		return base.I64_extend32_s(v5)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v14 = v6 - v13
		if int32(0) <= v13 {
			if v14 <= v6 {
				if base.Ui32(int32(2147483494)) <= base.Ui32(v14+int32(_a_F_date_mii_0)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_date_mii_1), int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_date_mii_2), int32(607), int32(_a_F_date_mii_3))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					return base.I64_extend_i32_s(v14)
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_date_mii_1), int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_date_mii_2), int32(607), int32(_a_F_date_mii_3))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			if v14 < v6 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_date_mii_1), int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_date_mii_2), int32(607), int32(_a_F_date_mii_3))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				if base.Ui32(int32(2147483494)) <= base.Ui32(v14+int32(_a_F_date_mii_0)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_date_mii_1), int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_date_mii_2), int32(607), int32(_a_F_date_mii_3))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					return base.I64_extend_i32_s(v14)
				}
			}
		}
	}
}
