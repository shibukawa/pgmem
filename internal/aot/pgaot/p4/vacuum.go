package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_VacuumUpdateCosts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 float64
	_ = v13
	var v17 float64
	_ = v17
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 float64
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 float64
	_ = v59
	var v64 float64
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 float64
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[718]))
	if v11 != 0 {
		v13 = *(*float64)(unsafe.Add(mBase, _consts[719]))
		if base.F64_ge(v13, float64(0)) != 0 {
			v22 = v13
		} else {
			v17 = *(*float64)(unsafe.Add(mBase, _consts[720]))
			if base.F64_ge(v17, float64(0)) != 0 {
				v22 = v17
			} else {
				v21 = *(*float64)(unsafe.Add(mBase, _consts[721]))
				v22 = v21
			}
		}
		*(*float64)(unsafe.Add(mBase, _consts[722])) = v22
		v26 = *(*int32)(unsafe.Add(mBase, _consts[723]))
		if int32(0) < v26 {
			v57 = v26
			v59 = v22
			*(*int32)(unsafe.Add(mBase, _consts[724])) = v57
			v64 = v59
			v66 = int32(*(*uint8)(unsafe.Add(mBase, _consts[725])))
			if v66 != 0 {
			} else {
				if base.F64_gt(v64, float64(0)) != 0 {
					v70 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _consts[726])) = uint8(v70)
				} else {
					v73 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[471])) = v73
					*(*uint8)(unsafe.Add(mBase, _consts[726])) = uint8(v73)
				}
			}
			if v11 == int32(0) {
				m.G0 = v8 + int32(32)
				return
			} else {
				v80 = int32(13)
				v87 = *(*int32)(unsafe.Add(mBase, _consts[224]))
				if v87 == int32(15) {
					v99 = int32(0)
					v103 = *(*int32)(unsafe.Add(mBase, _consts[225]))
					if v103 != int32(2) {
						v114 = v99
					} else {
						v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[226])))
						if v107 != 0 {
							v114 = v99
						} else {
							v111 = *(*int32)(unsafe.Add(mBase, _consts[227]))
							v114 = int32(0) | base.B2i32(v111 <= v80)
						}
					}
				} else {
					if v87 <= v80 {
						v114 = int32(1)
					} else {
						v99 = int32(0)
						v103 = *(*int32)(unsafe.Add(mBase, _consts[225]))
						if v103 != int32(2) {
							v114 = v99
						} else {
							v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[226])))
							if v107 != 0 {
								v114 = v99
							} else {
								v111 = *(*int32)(unsafe.Add(mBase, _consts[227]))
								v114 = int32(0) | base.B2i32(v111 <= v80)
							}
						}
					}
				}
				if v114 == int32(0) {
					m.G0 = v8 + int32(32)
					return
				} else {
					v119 = *(*int32)(unsafe.Add(mBase, _consts[86]))
					v123 = F_LWLockAcquire(m, v119+int32(2816), int32(1))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return
					} else {
						v126 = *(*int32)(unsafe.Add(mBase, _consts[718]))
						v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
						v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
						v130 = *(*int32)(unsafe.Add(mBase, _consts[86]))
						F_LWLockRelease(m, v130+int32(2816))
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return
						} else {
							v137 = F_errstart(m, int32(13), int32(0))
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return
							} else {
								if v137 == int32(0) {
									m.G0 = v8 + int32(32)
									return
								} else {
									v142 = *(*int32)(unsafe.Add(mBase, _consts[718]))
									v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+32))
									v145 = *(*float64)(unsafe.Add(mBase, _consts[722]))
									*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v145
									v150 = int32(*(*uint8)(unsafe.Add(mBase, _consts[725])))
									if v150 != 0 {
										v151 = int32(167702)
									} else {
										v151 = int32(253839)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v151
									if base.F64_gt(v145, float64(0)) != 0 {
										v157 = int32(167702)
									} else {
										v157 = int32(253839)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v157
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v128
									*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v127
									if v143 != 0 {
										v163 = int32(167702)
									} else {
										v163 = int32(253839)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v163
									v166 = *(*int32)(unsafe.Add(mBase, _consts[724]))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v166
									F_errmsg_internal(m, int32(708702), v8)
									mBase = m.M
									v170 = m.ExcPending
									if v170 != 0 {
										return
									} else {
										F_errfinish(m, int32(523274), int32(1710), int32(123678))
										mBase = m.M
										v175 = m.ExcPending
										if v175 != 0 {
											return
										} else {
											m.G0 = v8 + int32(32)
											return
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, _consts[727]))
			v33 = *(*int32)(unsafe.Add(mBase, _consts[728]))
			if int32(0) < v31 {
				v36 = v31
			} else {
				v36 = v33
			}
			*(*int32)(unsafe.Add(mBase, _consts[724])) = v36
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
			if v38 == int32(0) {
				v64 = v22
				v66 = int32(*(*uint8)(unsafe.Add(mBase, _consts[725])))
				if v66 != 0 {
				} else {
					if base.F64_gt(v64, float64(0)) != 0 {
						v70 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _consts[726])) = uint8(v70)
					} else {
						v73 = int32(0)
						*(*int32)(unsafe.Add(mBase, _consts[471])) = v73
						*(*uint8)(unsafe.Add(mBase, _consts[726])) = uint8(v73)
					}
				}
				if v11 == int32(0) {
					m.G0 = v8 + int32(32)
					return
				} else {
					v80 = int32(13)
					v87 = *(*int32)(unsafe.Add(mBase, _consts[224]))
					if v87 == int32(15) {
						v99 = int32(0)
						v103 = *(*int32)(unsafe.Add(mBase, _consts[225]))
						if v103 != int32(2) {
							v114 = v99
						} else {
							v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[226])))
							if v107 != 0 {
								v114 = v99
							} else {
								v111 = *(*int32)(unsafe.Add(mBase, _consts[227]))
								v114 = int32(0) | base.B2i32(v111 <= v80)
							}
						}
					} else {
						if v87 <= v80 {
							v114 = int32(1)
						} else {
							v99 = int32(0)
							v103 = *(*int32)(unsafe.Add(mBase, _consts[225]))
							if v103 != int32(2) {
								v114 = v99
							} else {
								v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[226])))
								if v107 != 0 {
									v114 = v99
								} else {
									v111 = *(*int32)(unsafe.Add(mBase, _consts[227]))
									v114 = int32(0) | base.B2i32(v111 <= v80)
								}
							}
						}
					}
					if v114 == int32(0) {
						m.G0 = v8 + int32(32)
						return
					} else {
						v119 = *(*int32)(unsafe.Add(mBase, _consts[86]))
						v123 = F_LWLockAcquire(m, v119+int32(2816), int32(1))
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, _consts[718]))
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
							v130 = *(*int32)(unsafe.Add(mBase, _consts[86]))
							F_LWLockRelease(m, v130+int32(2816))
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
								return
							} else {
								v137 = F_errstart(m, int32(13), int32(0))
								mBase = m.M
								v138 = m.ExcPending
								if v138 != 0 {
									return
								} else {
									if v137 == int32(0) {
										m.G0 = v8 + int32(32)
										return
									} else {
										v142 = *(*int32)(unsafe.Add(mBase, _consts[718]))
										v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+32))
										v145 = *(*float64)(unsafe.Add(mBase, _consts[722]))
										*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v145
										v150 = int32(*(*uint8)(unsafe.Add(mBase, _consts[725])))
										if v150 != 0 {
											v151 = int32(167702)
										} else {
											v151 = int32(253839)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v151
										if base.F64_gt(v145, float64(0)) != 0 {
											v157 = int32(167702)
										} else {
											v157 = int32(253839)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v157
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v128
										*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v127
										if v143 != 0 {
											v163 = int32(167702)
										} else {
											v163 = int32(253839)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v163
										v166 = *(*int32)(unsafe.Add(mBase, _consts[724]))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v166
										F_errmsg_internal(m, int32(708702), v8)
										mBase = m.M
										v170 = m.ExcPending
										if v170 != 0 {
											return
										} else {
											F_errfinish(m, int32(523274), int32(1710), int32(123678))
											mBase = m.M
											v175 = m.ExcPending
											if v175 != 0 {
												return
											} else {
												m.G0 = v8 + int32(32)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, _consts[708]))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_consts[716])))
				if v43 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v186 = m.ExcPending
					if v186 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(599842), int32(0))
						mBase = m.M
						v190 = m.ExcPending
						if v190 != 0 {
							return
						} else {
							F_errfinish(m, int32(523274), int32(1754), int32(108084))
							mBase = m.M
							v195 = m.ExcPending
							if v195 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v46 = int32(1)
					v47 = base.I32_div_s(v36, v43)
					if v47 <= v46 {
						v50 = v46
					} else {
						v50 = v47
					}
					v57 = v50
					v59 = v22
					*(*int32)(unsafe.Add(mBase, _consts[724])) = v57
					v64 = v59
					v66 = int32(*(*uint8)(unsafe.Add(mBase, _consts[725])))
					if v66 != 0 {
					} else {
						if base.F64_gt(v64, float64(0)) != 0 {
							v70 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[726])) = uint8(v70)
						} else {
							v73 = int32(0)
							*(*int32)(unsafe.Add(mBase, _consts[471])) = v73
							*(*uint8)(unsafe.Add(mBase, _consts[726])) = uint8(v73)
						}
					}
					if v11 == int32(0) {
						m.G0 = v8 + int32(32)
						return
					} else {
						v80 = int32(13)
						v87 = *(*int32)(unsafe.Add(mBase, _consts[224]))
						if v87 == int32(15) {
							v99 = int32(0)
							v103 = *(*int32)(unsafe.Add(mBase, _consts[225]))
							if v103 != int32(2) {
								v114 = v99
							} else {
								v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[226])))
								if v107 != 0 {
									v114 = v99
								} else {
									v111 = *(*int32)(unsafe.Add(mBase, _consts[227]))
									v114 = int32(0) | base.B2i32(v111 <= v80)
								}
							}
						} else {
							if v87 <= v80 {
								v114 = int32(1)
							} else {
								v99 = int32(0)
								v103 = *(*int32)(unsafe.Add(mBase, _consts[225]))
								if v103 != int32(2) {
									v114 = v99
								} else {
									v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[226])))
									if v107 != 0 {
										v114 = v99
									} else {
										v111 = *(*int32)(unsafe.Add(mBase, _consts[227]))
										v114 = int32(0) | base.B2i32(v111 <= v80)
									}
								}
							}
						}
						if v114 == int32(0) {
							m.G0 = v8 + int32(32)
							return
						} else {
							v119 = *(*int32)(unsafe.Add(mBase, _consts[86]))
							v123 = F_LWLockAcquire(m, v119+int32(2816), int32(1))
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
								return
							} else {
								v126 = *(*int32)(unsafe.Add(mBase, _consts[718]))
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
								v130 = *(*int32)(unsafe.Add(mBase, _consts[86]))
								F_LWLockRelease(m, v130+int32(2816))
								mBase = m.M
								v134 = m.ExcPending
								if v134 != 0 {
									return
								} else {
									v137 = F_errstart(m, int32(13), int32(0))
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
										return
									} else {
										if v137 == int32(0) {
											m.G0 = v8 + int32(32)
											return
										} else {
											v142 = *(*int32)(unsafe.Add(mBase, _consts[718]))
											v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+32))
											v145 = *(*float64)(unsafe.Add(mBase, _consts[722]))
											*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v145
											v150 = int32(*(*uint8)(unsafe.Add(mBase, _consts[725])))
											if v150 != 0 {
												v151 = int32(167702)
											} else {
												v151 = int32(253839)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v151
											if base.F64_gt(v145, float64(0)) != 0 {
												v157 = int32(167702)
											} else {
												v157 = int32(253839)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v157
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v128
											*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v127
											if v143 != 0 {
												v163 = int32(167702)
											} else {
												v163 = int32(253839)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v163
											v166 = *(*int32)(unsafe.Add(mBase, _consts[724]))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v166
											F_errmsg_internal(m, int32(708702), v8)
											mBase = m.M
											v170 = m.ExcPending
											if v170 != 0 {
												return
											} else {
												F_errfinish(m, int32(523274), int32(1710), int32(123678))
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return
												} else {
													m.G0 = v8 + int32(32)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v53 = *(*float64)(unsafe.Add(mBase, _consts[721]))
		*(*float64)(unsafe.Add(mBase, _consts[722])) = v53
		v56 = *(*int32)(unsafe.Add(mBase, _consts[728]))
		v57 = v56
		v59 = v53
		*(*int32)(unsafe.Add(mBase, _consts[724])) = v57
		v64 = v59
		v66 = int32(*(*uint8)(unsafe.Add(mBase, _consts[725])))
		if v66 != 0 {
		} else {
			if base.F64_gt(v64, float64(0)) != 0 {
				v70 = int32(1)
				*(*uint8)(unsafe.Add(mBase, _consts[726])) = uint8(v70)
			} else {
				v73 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[471])) = v73
				*(*uint8)(unsafe.Add(mBase, _consts[726])) = uint8(v73)
			}
		}
		if v11 == int32(0) {
			m.G0 = v8 + int32(32)
			return
		} else {
			v80 = int32(13)
			v87 = *(*int32)(unsafe.Add(mBase, _consts[224]))
			if v87 == int32(15) {
				v99 = int32(0)
				v103 = *(*int32)(unsafe.Add(mBase, _consts[225]))
				if v103 != int32(2) {
					v114 = v99
				} else {
					v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[226])))
					if v107 != 0 {
						v114 = v99
					} else {
						v111 = *(*int32)(unsafe.Add(mBase, _consts[227]))
						v114 = int32(0) | base.B2i32(v111 <= v80)
					}
				}
			} else {
				if v87 <= v80 {
					v114 = int32(1)
				} else {
					v99 = int32(0)
					v103 = *(*int32)(unsafe.Add(mBase, _consts[225]))
					if v103 != int32(2) {
						v114 = v99
					} else {
						v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[226])))
						if v107 != 0 {
							v114 = v99
						} else {
							v111 = *(*int32)(unsafe.Add(mBase, _consts[227]))
							v114 = int32(0) | base.B2i32(v111 <= v80)
						}
					}
				}
			}
			if v114 == int32(0) {
				m.G0 = v8 + int32(32)
				return
			} else {
				v119 = *(*int32)(unsafe.Add(mBase, _consts[86]))
				v123 = F_LWLockAcquire(m, v119+int32(2816), int32(1))
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return
				} else {
					v126 = *(*int32)(unsafe.Add(mBase, _consts[718]))
					v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
					v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
					v130 = *(*int32)(unsafe.Add(mBase, _consts[86]))
					F_LWLockRelease(m, v130+int32(2816))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						v137 = F_errstart(m, int32(13), int32(0))
						mBase = m.M
						v138 = m.ExcPending
						if v138 != 0 {
							return
						} else {
							if v137 == int32(0) {
								m.G0 = v8 + int32(32)
								return
							} else {
								v142 = *(*int32)(unsafe.Add(mBase, _consts[718]))
								v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+32))
								v145 = *(*float64)(unsafe.Add(mBase, _consts[722]))
								*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v145
								v150 = int32(*(*uint8)(unsafe.Add(mBase, _consts[725])))
								if v150 != 0 {
									v151 = int32(167702)
								} else {
									v151 = int32(253839)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v151
								if base.F64_gt(v145, float64(0)) != 0 {
									v157 = int32(167702)
								} else {
									v157 = int32(253839)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v157
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v128
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v127
								if v143 != 0 {
									v163 = int32(167702)
								} else {
									v163 = int32(253839)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v163
								v166 = *(*int32)(unsafe.Add(mBase, _consts[724]))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v166
								F_errmsg_internal(m, int32(708702), v8)
								mBase = m.M
								v170 = m.ExcPending
								if v170 != 0 {
									return
								} else {
									F_errfinish(m, int32(523274), int32(1710), int32(123678))
									mBase = m.M
									v175 = m.ExcPending
									if v175 != 0 {
										return
									} else {
										m.G0 = v8 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_vacuum_is_permitted_for_relation(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	v14 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	v15 = F_object_ownercheck(m, int32(1262), v12, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 != 0 {
			v19 = int32(1)
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+117)))
			if v20 != v19 {
				v71 = v19
				m.G0 = v8 + int32(16)
				return v71
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _consts[168]))
				v27 = F_pg_class_aclcheck(m, l0, v25, int64(16384))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if v27 == int32(0) {
						v71 = int32(1)
						m.G0 = v8 + int32(16)
						return v71
					} else {
						if l2&int32(1) != 0 {
							v34 = int32(0)
							v37 = F_errstart(m, int32(19), v34)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								if v37 == int32(0) {
									v71 = v34
									m.G0 = v8 + int32(16)
									return v71
								} else {
									v56 = int32(760)
									v57 = int32(111245)
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1 + int32(4)
									F_errmsg(m, v57, v8)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(523302), v56, int32(276758))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											v71 = int32(0)
											m.G0 = v8 + int32(16)
											return v71
										}
									}
								}
							}
						} else {
							if l2&int32(2) == int32(0) {
								v71 = int32(0)
								m.G0 = v8 + int32(16)
								return v71
							} else {
								v47 = int32(0)
								v50 = F_errstart(m, int32(19), v47)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									if v50 == int32(0) {
										v71 = v47
										m.G0 = v8 + int32(16)
										return v71
									} else {
										v56 = int32(773)
										v57 = int32(111291)
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1 + int32(4)
										F_errmsg(m, v57, v8)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(523302), v56, int32(276758))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v71 = int32(0)
												m.G0 = v8 + int32(16)
												return v71
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, _consts[168]))
			v27 = F_pg_class_aclcheck(m, l0, v25, int64(16384))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v27 == int32(0) {
					v71 = int32(1)
					m.G0 = v8 + int32(16)
					return v71
				} else {
					if l2&int32(1) != 0 {
						v34 = int32(0)
						v37 = F_errstart(m, int32(19), v34)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							if v37 == int32(0) {
								v71 = v34
								m.G0 = v8 + int32(16)
								return v71
							} else {
								v56 = int32(760)
								v57 = int32(111245)
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1 + int32(4)
								F_errmsg(m, v57, v8)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(523302), v56, int32(276758))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										v71 = int32(0)
										m.G0 = v8 + int32(16)
										return v71
									}
								}
							}
						}
					} else {
						if l2&int32(2) == int32(0) {
							v71 = int32(0)
							m.G0 = v8 + int32(16)
							return v71
						} else {
							v47 = int32(0)
							v50 = F_errstart(m, int32(19), v47)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								if v50 == int32(0) {
									v71 = v47
									m.G0 = v8 + int32(16)
									return v71
								} else {
									v56 = int32(773)
									v57 = int32(111291)
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1 + int32(4)
									F_errmsg(m, v57, v8)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(523302), v56, int32(276758))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											v71 = int32(0)
											m.G0 = v8 + int32(16)
											return v71
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
