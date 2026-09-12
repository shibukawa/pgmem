package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_UpdateChangedParamSet(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+68))
	v5 = F_bms_intersect(m, v4, l1)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v8 = F_bms_join(m, v7, v5)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v8
			return
		}
	}
}
func F_unlink_span(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
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
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v7 == int32(0) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v44 = v10
		if v44 != 0 {
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+1468))
			if v48 != v50 {
				v55 = F_LWLockAcquire(m, v49+int32(1476), int32(0))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					F_check_for_freed_segments_locked(m, l0)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						F_LWLockRelease(m, v59+int32(1476))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							v67 = int32(base.Ui32(v44) >> (uint(int32(27)) % 32))
							v72 = l0 + v67*int32(20) + int32(12)
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
							if v73 != 0 {
								v77 = v73
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
								return
							} else {
								v74 = F_get_segment_by_index(m, l0, v67)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
									v77 = v76
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
									return
								}
							}
						}
					}
				}
			} else {
				v67 = int32(base.Ui32(v44) >> (uint(int32(27)) % 32))
				v72 = l0 + v67*int32(20) + int32(12)
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
				if v73 != 0 {
					v77 = v73
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
					return
				} else {
					v74 = F_get_segment_by_index(m, l0, v67)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
						v77 = v76
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
						return
					}
				}
			}
		} else {
			v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
			v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+1468))
			if v82 != v84 {
				v89 = F_LWLockAcquire(m, v83+int32(1476), int32(0))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return
				} else {
					F_check_for_freed_segments_locked(m, l0)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return
					} else {
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						F_LWLockRelease(m, v93+int32(1476))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							v101 = int32(base.Ui32(v81) >> (uint(int32(27)) % 32))
							v106 = l0 + v101*int32(20) + int32(12)
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
							if v107 != 0 {
								v111 = v107
								v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
								v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
								return
							} else {
								v108 = F_get_segment_by_index(m, l0, v101)
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return
								} else {
									v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
									v111 = v110
									v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
									v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
									return
								}
							}
						}
					}
				}
			} else {
				v101 = int32(base.Ui32(v81) >> (uint(int32(27)) % 32))
				v106 = l0 + v101*int32(20) + int32(12)
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
				if v107 != 0 {
					v111 = v107
					v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
					return
				} else {
					v108 = F_get_segment_by_index(m, l0, v101)
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return
					} else {
						v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
						v111 = v110
						v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
						v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
						return
					}
				}
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1468))
		if v11 != v13 {
			v18 = F_LWLockAcquire(m, v12+int32(1476), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				F_check_for_freed_segments_locked(m, l0)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					F_LWLockRelease(m, v22+int32(1476))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v30 = int32(base.Ui32(v7) >> (uint(int32(27)) % 32))
						v35 = l0 + v30*int32(20) + int32(12)
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
						if v36 != 0 {
							v40 = v36
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v40+v7&int32(134217727))+4)) = v42
							v44 = v42
							if v44 != 0 {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+1468))
								if v48 != v50 {
									v55 = F_LWLockAcquire(m, v49+int32(1476), int32(0))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										F_check_for_freed_segments_locked(m, l0)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return
										} else {
											v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											F_LWLockRelease(m, v59+int32(1476))
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return
											} else {
												v67 = int32(base.Ui32(v44) >> (uint(int32(27)) % 32))
												v72 = l0 + v67*int32(20) + int32(12)
												v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
												if v73 != 0 {
													v77 = v73
													v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
													return
												} else {
													v74 = F_get_segment_by_index(m, l0, v67)
													mBase = m.M
													v75 = m.ExcPending
													if v75 != 0 {
														return
													} else {
														v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
														v77 = v76
														v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
														return
													}
												}
											}
										}
									}
								} else {
									v67 = int32(base.Ui32(v44) >> (uint(int32(27)) % 32))
									v72 = l0 + v67*int32(20) + int32(12)
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
									if v73 != 0 {
										v77 = v73
										v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
										return
									} else {
										v74 = F_get_segment_by_index(m, l0, v67)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
											v77 = v76
											v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
											return
										}
									}
								}
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
								v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+1468))
								if v82 != v84 {
									v89 = F_LWLockAcquire(m, v83+int32(1476), int32(0))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										F_check_for_freed_segments_locked(m, l0)
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return
										} else {
											v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											F_LWLockRelease(m, v93+int32(1476))
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
											} else {
												v101 = int32(base.Ui32(v81) >> (uint(int32(27)) % 32))
												v106 = l0 + v101*int32(20) + int32(12)
												v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
												if v107 != 0 {
													v111 = v107
													v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
													v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
													return
												} else {
													v108 = F_get_segment_by_index(m, l0, v101)
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return
													} else {
														v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
														v111 = v110
														v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
														v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
														return
													}
												}
											}
										}
									}
								} else {
									v101 = int32(base.Ui32(v81) >> (uint(int32(27)) % 32))
									v106 = l0 + v101*int32(20) + int32(12)
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
									if v107 != 0 {
										v111 = v107
										v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
										v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
										return
									} else {
										v108 = F_get_segment_by_index(m, l0, v101)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
											v111 = v110
											v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
											v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
											return
										}
									}
								}
							}
						} else {
							v37 = F_get_segment_by_index(m, l0, v30)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
								v40 = v39
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v40+v7&int32(134217727))+4)) = v42
								v44 = v42
								if v44 != 0 {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
									v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+1468))
									if v48 != v50 {
										v55 = F_LWLockAcquire(m, v49+int32(1476), int32(0))
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return
										} else {
											F_check_for_freed_segments_locked(m, l0)
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return
											} else {
												v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_LWLockRelease(m, v59+int32(1476))
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return
												} else {
													v67 = int32(base.Ui32(v44) >> (uint(int32(27)) % 32))
													v72 = l0 + v67*int32(20) + int32(12)
													v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
													if v73 != 0 {
														v77 = v73
														v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
														return
													} else {
														v74 = F_get_segment_by_index(m, l0, v67)
														mBase = m.M
														v75 = m.ExcPending
														if v75 != 0 {
															return
														} else {
															v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
															v77 = v76
															v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
															return
														}
													}
												}
											}
										}
									} else {
										v67 = int32(base.Ui32(v44) >> (uint(int32(27)) % 32))
										v72 = l0 + v67*int32(20) + int32(12)
										v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
										if v73 != 0 {
											v77 = v73
											v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
											return
										} else {
											v74 = F_get_segment_by_index(m, l0, v67)
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return
											} else {
												v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
												v77 = v76
												v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
												return
											}
										}
									}
								} else {
									v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
									v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+1468))
									if v82 != v84 {
										v89 = F_LWLockAcquire(m, v83+int32(1476), int32(0))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											F_check_for_freed_segments_locked(m, l0)
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_LWLockRelease(m, v93+int32(1476))
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return
												} else {
													v101 = int32(base.Ui32(v81) >> (uint(int32(27)) % 32))
													v106 = l0 + v101*int32(20) + int32(12)
													v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
													if v107 != 0 {
														v111 = v107
														v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
														v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
														return
													} else {
														v108 = F_get_segment_by_index(m, l0, v101)
														mBase = m.M
														v109 = m.ExcPending
														if v109 != 0 {
															return
														} else {
															v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
															v111 = v110
															v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
															v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
															return
														}
													}
												}
											}
										}
									} else {
										v101 = int32(base.Ui32(v81) >> (uint(int32(27)) % 32))
										v106 = l0 + v101*int32(20) + int32(12)
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
										if v107 != 0 {
											v111 = v107
											v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
											v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
											return
										} else {
											v108 = F_get_segment_by_index(m, l0, v101)
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return
											} else {
												v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
												v111 = v110
												v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
												v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
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
		} else {
			v30 = int32(base.Ui32(v7) >> (uint(int32(27)) % 32))
			v35 = l0 + v30*int32(20) + int32(12)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
			if v36 != 0 {
				v40 = v36
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v40+v7&int32(134217727))+4)) = v42
				v44 = v42
				if v44 != 0 {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+1468))
					if v48 != v50 {
						v55 = F_LWLockAcquire(m, v49+int32(1476), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							F_check_for_freed_segments_locked(m, l0)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								F_LWLockRelease(m, v59+int32(1476))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									v67 = int32(base.Ui32(v44) >> (uint(int32(27)) % 32))
									v72 = l0 + v67*int32(20) + int32(12)
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
									if v73 != 0 {
										v77 = v73
										v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
										return
									} else {
										v74 = F_get_segment_by_index(m, l0, v67)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
											v77 = v76
											v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
											return
										}
									}
								}
							}
						}
					} else {
						v67 = int32(base.Ui32(v44) >> (uint(int32(27)) % 32))
						v72 = l0 + v67*int32(20) + int32(12)
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
						if v73 != 0 {
							v77 = v73
							v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
							return
						} else {
							v74 = F_get_segment_by_index(m, l0, v67)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
								v77 = v76
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
								return
							}
						}
					}
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+1468))
					if v82 != v84 {
						v89 = F_LWLockAcquire(m, v83+int32(1476), int32(0))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return
						} else {
							F_check_for_freed_segments_locked(m, l0)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return
							} else {
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								F_LWLockRelease(m, v93+int32(1476))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return
								} else {
									v101 = int32(base.Ui32(v81) >> (uint(int32(27)) % 32))
									v106 = l0 + v101*int32(20) + int32(12)
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
									if v107 != 0 {
										v111 = v107
										v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
										v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
										return
									} else {
										v108 = F_get_segment_by_index(m, l0, v101)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
											v111 = v110
											v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
											v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
											return
										}
									}
								}
							}
						}
					} else {
						v101 = int32(base.Ui32(v81) >> (uint(int32(27)) % 32))
						v106 = l0 + v101*int32(20) + int32(12)
						v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
						if v107 != 0 {
							v111 = v107
							v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
							v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
							return
						} else {
							v108 = F_get_segment_by_index(m, l0, v101)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
								v111 = v110
								v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
								v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
								return
							}
						}
					}
				}
			} else {
				v37 = F_get_segment_by_index(m, l0, v30)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
					v40 = v39
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v40+v7&int32(134217727))+4)) = v42
					v44 = v42
					if v44 != 0 {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+1468))
						if v48 != v50 {
							v55 = F_LWLockAcquire(m, v49+int32(1476), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								F_check_for_freed_segments_locked(m, l0)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									F_LWLockRelease(m, v59+int32(1476))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										v67 = int32(base.Ui32(v44) >> (uint(int32(27)) % 32))
										v72 = l0 + v67*int32(20) + int32(12)
										v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
										if v73 != 0 {
											v77 = v73
											v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
											return
										} else {
											v74 = F_get_segment_by_index(m, l0, v67)
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return
											} else {
												v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
												v77 = v76
												v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
												return
											}
										}
									}
								}
							}
						} else {
							v67 = int32(base.Ui32(v44) >> (uint(int32(27)) % 32))
							v72 = l0 + v67*int32(20) + int32(12)
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
							if v73 != 0 {
								v77 = v73
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
								return
							} else {
								v74 = F_get_segment_by_index(m, l0, v67)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
									v77 = v76
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v77+v44&int32(134217727))+8)) = v79
									return
								}
							}
						}
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+1468))
						if v82 != v84 {
							v89 = F_LWLockAcquire(m, v83+int32(1476), int32(0))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								F_check_for_freed_segments_locked(m, l0)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return
								} else {
									v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									F_LWLockRelease(m, v93+int32(1476))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return
									} else {
										v101 = int32(base.Ui32(v81) >> (uint(int32(27)) % 32))
										v106 = l0 + v101*int32(20) + int32(12)
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
										if v107 != 0 {
											v111 = v107
											v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
											v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
											return
										} else {
											v108 = F_get_segment_by_index(m, l0, v101)
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return
											} else {
												v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
												v111 = v110
												v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
												v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
												return
											}
										}
									}
								}
							}
						} else {
							v101 = int32(base.Ui32(v81) >> (uint(int32(27)) % 32))
							v106 = l0 + v101*int32(20) + int32(12)
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
							if v107 != 0 {
								v111 = v107
								v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
								v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
								return
							} else {
								v108 = F_get_segment_by_index(m, l0, v101)
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return
								} else {
									v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
									v111 = v110
									v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
									v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v111+v81&int32(134217727)+v113<<(uint(int32(2))%32))+16)) = v117
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
func F_upc_cast_from_ean13(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	F_ean2isn(m, v8, v5+int32(8), int32(6))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
		v17 = F_Int64GetDatum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			m.G0 = v5 + int32(16)
			return v17
		}
	}
}
func F_upc_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = F_string2ean(m, v7, v8, v5+int32(8), int32(6))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v24 = int32(0)
			m.G0 = v5 + int32(16)
			return v24
		} else {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
			v22 = F_Int64GetDatum(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = v22
				m.G0 = v5 + int32(16)
				return v24
			}
		}
	}
}
func F_updateAclDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	var v11 int32
	_ = v11
	F_updateAclDependenciesWorker(m, l0, l1, l2, l3, int32(97), l4, l5, l6, l7)
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		return
	}
}
func F_updateAclDependenciesWorker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v431 int32
	_ = v431
	var v446 int32
	_ = v446
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v485 int32
	_ = v485
	var v497 int32
	_ = v497
	var v513 int32
	_ = v513
	var v541 int32
	_ = v541
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v607 int32
	_ = v607
	var v618 int32
	_ = v618
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v664 int32
	_ = v664
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	v10 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(48)
	m.G0 = v22
	if l5 <= v10 {
		v91 = v10
		v92 = v10
		v94 = v10
		v95 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v91 < l5 {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v36 = v10
	v37 = v10
	v39 = v10
	v40 = v10
	goto L3
L3:
	;
	if l7 <= v37 {
		v91 = v36
		v92 = v37
		v94 = v39
		v95 = v40
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v91 = v76
	v92 = v77
	v94 = v78
	v95 = v79
	goto L1
L5:
	;
	v46 = int32(2)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l6+v36<<(uint(v46)%32))))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l8+v37<<(uint(v46)%32))))
	if v49 == v53 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v76 < l5 {
		v36 = v76
		v37 = v77
		v39 = v78
		v40 = v79
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v55 = int32(1)
	v76 = v36 + v55
	v77 = v37 + v55
	v78 = v39
	v79 = v40
	goto L6
L8:
	;
	goto L9
L9:
	;
	if base.Ui32(v49) < base.Ui32(v53) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6+v40<<(uint(int32(2))%32)))) = v49
	v64 = int32(1)
	v76 = v36 + v64
	v77 = v37
	v78 = v39
	v79 = v40 + v64
	goto L6
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8+v39<<(uint(int32(2))%32)))) = v53
	v72 = int32(1)
	v76 = v36
	v77 = v37 + v72
	v78 = v39 + v72
	v79 = v40
	goto L6
L13:
	;
	goto L4
L14:
	;
	v104 = (l5 - v91) & int32(3)
	if v104 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v248 = v95
	goto L16
L16:
	;
	if v92 < l7 {
		goto L30
	} else {
		goto L31
	}
L17:
	;
	v161 = l5 + v95 - v91
	if base.Ui32(v91-l5) <= base.Ui32(int32(-4)) {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v151 = v91
	v156 = v95
	goto L17
L19:
	;
	goto L20
L20:
	;
	v117 = v91
	v120 = int32(0)
	v122 = v95
	goto L21
L21:
	;
	v127 = int32(2)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l6+v117<<(uint(v127)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l6+v122<<(uint(v127)%32)))) = v133
	v135 = int32(1)
	v136 = v122 + v135
	v138 = v117 + v135
	v140 = v120 + v135
	if v140 != v104 {
		v117 = v138
		v120 = v140
		v122 = v136
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v151 = v138
	v156 = v136
	goto L17
L23:
	;
	goto L22
L24:
	;
	v166 = l6 + int32(12)
	v168 = l6 + int32(8)
	v170 = l6 + int32(4)
	v180 = v151
	v185 = v156
	goto L27
L25:
	;
	goto L26
L26:
	;
	v248 = v161
	goto L16
L27:
	;
	v190 = int32(2)
	v191 = v185 << (uint(v190) % 32)
	v194 = v180 << (uint(v190) % 32)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l6+v194)))
	*(*int32)(unsafe.Add(mBase, uint32(l6+v191))) = v196
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v194+v170)))
	*(*int32)(unsafe.Add(mBase, uint32(v191+v170))) = v200
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v194+v168)))
	*(*int32)(unsafe.Add(mBase, uint32(v191+v168))) = v204
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v166+v194)))
	*(*int32)(unsafe.Add(mBase, uint32(v166+v191))) = v208
	v210 = int32(4)
	v213 = v185 + v210
	if v213 != v161 {
		v180 = v180 + v210
		v185 = v213
		goto L27
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	goto L28
L30:
	;
	v257 = (l7 - v92) & int32(3)
	if v257 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v400 = v94
	goto L32
L32:
	;
	v406 = int32(0)
	if base.B2i32(v248 <= v406)&base.B2i32(v400 <= v406) == v406 {
		goto L46
	} else {
		goto L47
	}
L33:
	;
	v314 = l7 + v94 - v92
	if base.Ui32(v92-l7) <= base.Ui32(int32(-4)) {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	v304 = v92
	v308 = v94
	goto L33
L35:
	;
	goto L36
L36:
	;
	v270 = v92
	v271 = int32(0)
	v274 = v94
	goto L37
L37:
	;
	v280 = int32(2)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l8+v270<<(uint(v280)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l8+v274<<(uint(v280)%32)))) = v286
	v288 = int32(1)
	v289 = v274 + v288
	v291 = v270 + v288
	v293 = v271 + v288
	if v293 != v257 {
		v270 = v291
		v271 = v293
		v274 = v289
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v304 = v291
	v308 = v289
	goto L33
L39:
	;
	goto L38
L40:
	;
	v319 = l8 + int32(12)
	v321 = l8 + int32(8)
	v323 = l8 + int32(4)
	v333 = v304
	v337 = v308
	goto L43
L41:
	;
	goto L42
L42:
	;
	v400 = v314
	goto L32
L43:
	;
	v343 = int32(2)
	v344 = v337 << (uint(v343) % 32)
	v347 = v333 << (uint(v343) % 32)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l8+v347)))
	*(*int32)(unsafe.Add(mBase, uint32(l8+v344))) = v349
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v347+v323)))
	*(*int32)(unsafe.Add(mBase, uint32(v344+v323))) = v353
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v321+v347)))
	*(*int32)(unsafe.Add(mBase, uint32(v321+v344))) = v357
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v347+v319)))
	*(*int32)(unsafe.Add(mBase, uint32(v344+v319))) = v361
	v363 = int32(4)
	v366 = v337 + v363
	if v366 != v314 {
		v333 = v333 + v363
		v337 = v366
		goto L43
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	goto L44
L46:
	;
	v415 = F_table_open(m, int32(1214), int32(3))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	if l6 != 0 {
		goto L114
	} else {
		goto L115
	}
L49:
	;
	return
L50:
	;
	if int32(0) < v400 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v431 = int32(0)
	goto L54
L52:
	;
	goto L53
L53:
	;
	if int32(0) < v248 {
		goto L99
	} else {
		goto L100
	}
L54:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l8+v431<<(uint(int32(2))%32))))
	if base.B2i32(base.B2i32(l4 != int32(97)) == int32(0))&base.B2i32(v446 == l3) != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	v568 = v431 + int32(1)
	if v568 != v400 {
		v431 = v568
		goto L54
	} else {
		goto L98
	}
L57:
	;
	goto L60
L58:
	;
	if v463 != 0 {
		goto L56
	} else {
		goto L62
	}
L59:
	;
	goto L58
L60:
	;
	if base.Ui32(int32(11999)) < base.Ui32(v446) {
		v463 = int32(0)
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v456 = int32(1)
	v463 = (v456 | base.B2i32(v446 != int32(2200))) & v456
	goto L59
L62:
	;
	F_shdepLockAndCheckObject(m, int32(1260), v446)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L49
	} else {
		goto L63
	}
L63:
	;
	v467 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+11)) = v467
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v467
	v473 = int32(1)
	if l0 <= int32(3591) {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = int32(1260)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = l0
	v551 = *(*int32)(unsafe.Add(mBase, _consts[223]))
	if v541 != 0 {
		goto L92
	} else {
		goto L93
	}
L65:
	;
	goto L64
L66:
	;
	v541 = int32(0)
	goto L65
L67:
	;
	if base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)) {
		v541 = v473
		goto L65
	} else {
		goto L90
	}
L68:
	;
	if l0 <= int32(2670) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	if l0 <= int32(5999) {
		goto L79
	} else {
		goto L80
	}
L71:
	;
	switch l0 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v541 = v473
		goto L65
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L66
	default:
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v485 = l0 - int32(2671)
	if base.Ui32(int32(27)) < base.Ui32(v485) {
		goto L67
	} else {
		goto L76
	}
L74:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l0-int32(2396)) {
		goto L66
	} else {
		goto L75
	}
L75:
	;
	v541 = v473
	goto L65
L76:
	;
	if int32(1)<<(uint(v485)%32)&int32(226492515) == int32(0) {
		goto L67
	} else {
		goto L77
	}
L77:
	;
	v541 = v473
	goto L65
L78:
	;
	if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
		v541 = v473
		goto L65
	} else {
		goto L88
	}
L79:
	;
	v497 = l0 - int32(4177)
	if base.Ui32(int32(9)) < base.Ui32(v497) {
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	switch l0 - int32(6243) {
	case 0, 1, 2, 3, 4, 59, 60:
		v541 = v473
		goto L65
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L66
	default:
		goto L84
	}
L82:
	;
	if int32(1)<<(uint(v497)%32)&int32(963) == int32(0) {
		goto L78
	} else {
		goto L83
	}
L83:
	;
	v541 = v473
	goto L65
L84:
	;
	if base.Ui32(l0-int32(6000)) < base.Ui32(int32(3)) {
		v541 = v473
		goto L65
	} else {
		goto L85
	}
L85:
	;
	v513 = l0 - int32(6100)
	if base.Ui32(int32(15)) < base.Ui32(v513) {
		goto L66
	} else {
		goto L86
	}
L86:
	;
	if int32(1)<<(uint(v513)%32)&int32(49153) != 0 {
		v541 = v473
		goto L65
	} else {
		goto L87
	}
L87:
	;
	goto L66
L88:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
		goto L66
	} else {
		goto L89
	}
L89:
	;
	v541 = v473
	goto L65
L90:
	;
	if base.Ui32(l0-int32(2846)) < base.Ui32(int32(2)) {
		v541 = v473
		goto L65
	} else {
		goto L91
	}
L91:
	;
	goto L66
L92:
	;
	v552 = int32(0)
	goto L94
L93:
	;
	v552 = v551
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v552
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v415)+52))
	v559 = F_heap_form_tuple(m, v554, v22+int32(16), v22+int32(8))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L49
	} else {
		goto L95
	}
L95:
	;
	F_CatalogTupleInsert(m, v415, v559)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L49
	} else {
		goto L96
	}
L96:
	;
	F_pfree(m, v559)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L49
	} else {
		goto L97
	}
L97:
	;
	goto L56
L98:
	;
	goto L55
L99:
	;
	v607 = int32(0)
	goto L102
L100:
	;
	goto L101
L101:
	;
	F_sequence_close(m, v415, int32(3))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L49
	} else {
		goto L113
	}
L102:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l6+v607<<(uint(int32(2))%32))))
	if base.B2i32(base.B2i32(l4 != int32(97)) == int32(0))&base.B2i32(v618 == l3) != 0 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	goto L101
L104:
	;
	v641 = v607 + int32(1)
	if v641 != v248 {
		v607 = v641
		goto L102
	} else {
		goto L112
	}
L105:
	;
	goto L108
L106:
	;
	if v635 != 0 {
		goto L104
	} else {
		goto L110
	}
L107:
	;
	goto L106
L108:
	;
	if base.Ui32(int32(11999)) < base.Ui32(v618) {
		v635 = int32(0)
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v628 = int32(1)
	v635 = (v628 | base.B2i32(v618 != int32(2200))) & v628
	goto L107
L110:
	;
	F_shdepDropDependency(m, v415, l0, l1, l2, int32(0), int32(1260), v618, l4)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L49
	} else {
		goto L111
	}
L111:
	;
	goto L104
L112:
	;
	goto L103
L113:
	;
	goto L48
L114:
	;
	F_pfree(m, l6)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L49
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	if l8 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L116
L118:
	;
	F_pfree(m, l8)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L49
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	m.G0 = v22 + int32(48)
	return
L121:
	;
	goto L120
}
func F_update_metainfo_datafile(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	v4 = m.G0
	v6 = v4 - int32(144)
	m.G0 = v6
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[475])))
	if v9&int32(25) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v6 + int32(144)
	return
L2:
	;
	v15 = F_unlink(m, int32(175575))
	mBase = m.M
	if int32(0) <= v15 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[476]))
	v43 = int32(4452672)
	v44 = *(*int32)(unsafe.Add(mBase, _consts[477]))
	*(*int32)(unsafe.Add(mBase, _consts[477])) = v41
	v47 = F___syscall_ret(m, v44)
	mBase = m.M
	goto L13
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	if v19 == int32(44) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v24 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	if v24 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(175575)
	F_errmsg(m, int32(314011), v6)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(520822), int32(1490), int32(405992))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L1
L13:
	;
	v50 = F_fopen(m, int32(247843), int32(34101))
	mBase = m.M
	v52 = int32(4452672)
	v53 = *(*int32)(unsafe.Add(mBase, _consts[477]))
	*(*int32)(unsafe.Add(mBase, _consts[477])) = v47
	v56 = F___syscall_ret(m, v53)
	mBase = m.M
	goto L14
L14:
	;
	if v50 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _consts[479]))
	if v125 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+80)) = int32(-1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	if v59 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	v106 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L7
	} else {
		goto L35
	}
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[478]))
	if v67 == int32(0) {
		goto L15
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+80)) = int32(10)
	goto L22
L21:
	;
	goto L22
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v62 | int32(64)
	goto L19
L23:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, _consts[475])))
	if v71&int32(1) == int32(0) {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+128)) = v67
	v80 = F_pg_fprintf(m, v50, int32(784160), v6+int32(128))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	if int32(0) <= v80 {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	v86 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	if v86 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L7
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v102 = F_fclose(m, v50)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L7
	} else {
		goto L34
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+112)) = int32(247843)
	F_errmsg(m, int32(314042), v6+int32(112))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(520822), int32(1524), int32(405992))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	goto L1
L35:
	;
	if v106 == int32(0) {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(247843)
	F_errmsg(m, int32(313457), v6+int32(16))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(520822), int32(1513), int32(405992))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	goto L1
L40:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _consts[480]))
	if v163 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L41:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, _consts[475])))
	if v129&int32(8) == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+96)) = v125
	v138 = F_pg_fprintf(m, v50, int32(784171), v6+int32(96))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	if int32(0) <= v138 {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v144 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	if v144 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L7
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v160 = F_fclose(m, v50)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L7
	} else {
		goto L52
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = int32(247843)
	F_errmsg(m, int32(314042), v6+int32(80))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(520822), int32(1537), int32(405992))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	goto L48
L52:
	;
	goto L1
L53:
	;
	v200 = F_fclose(m, v50)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L7
	} else {
		goto L66
	}
L54:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, _consts[475])))
	if v167&int32(16) == int32(0) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+64)) = v163
	v176 = F_pg_fprintf(m, v50, int32(784182), v6-int32(-64))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	if int32(0) <= v176 {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v182 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	if v182 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L7
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v198 = F_fclose(m, v50)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L7
	} else {
		goto L65
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = int32(247843)
	F_errmsg(m, int32(314042), v6+int32(48))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(520822), int32(1550), int32(405992))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	goto L1
L66:
	;
	v204 = F_rename(m, int32(247843), int32(175575))
	mBase = m.M
	if v204 == int32(0) {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v209 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	if v209 == int32(0) {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = int32(175575)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(247843)
	F_errmsg(m, int32(312168), v6+int32(32))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(520822), int32(1561), int32(405992))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	goto L1
}
func F_utime(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v10 = m.Env.X__syscall_utimensat(m, int32(-100), l0, v2, v2)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v10) {
		*(*int32)(unsafe.Add(mBase, _consts[163])) = int32(0) - v10
	} else {
	}
	m.G0 = v5 + int32(32)
	return
}
